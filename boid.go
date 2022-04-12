package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	Position Vector
	Velocity Vector
	ID       int
}

func (b *Boid) calcAcceleration() Vector {
	upper, lower := b.Position.AddV(viewRadius), b.Position.AddV(-viewRadius)
	avgPosition, avgVelocity, separetion := Vector{X: 0, Y: 0}, Vector{X: 0, Y: 0}, Vector{X: 0, Y: 0}
	count := 0.0
	//finding the average velocity of the set of boids inside the viewRadius
	rwLock.RLock()
	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.ID {
				if distance := boids[otherBoidId].Position.Distance(b.Position); distance < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].Velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].Position)
					separetion = separetion.Add(b.Position.Subtract(boids[otherBoidId].Position).DivisionV(distance))
				}
			}
		}
	}
	rwLock.RUnlock()
	accel := Vector{X: b.borderBounce(b.Position.X, screenWidth), Y: b.borderBounce(b.Position.Y, screenHeight)}
	if count > 0 {
		avgVelocity = avgVelocity.DivisionV(count)
		avgPosition = avgPosition.DivisionV(count)
		accelAlignment := avgVelocity.Subtract(b.Velocity).MultiplyV(adjRate)
		accelCohesion := avgPosition.Subtract(b.Position).MultiplyV(adjRate)
		accelSeparetion := separetion.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparetion)
	}
	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	rwLock.Lock()
	b.Velocity = b.Velocity.Add(acceleration).Limit(-1, 1) //update the velocity //!the calcAcceleration func tries to aqquire the lock as well which won't be possible
	boidMap[int(b.Position.X)][int(b.Position.Y)] = -1     //update the position in the 2d array to default
	b.Position = b.Position.Add(b.Velocity)                //move in the units of velocity

	rwLock.Unlock()
}

func (b *Boid) start() {
	//while true the boids move util the programme is stopped or killed
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		Position: Vector{X: rand.Float64() * screenWidth, Y: rand.Float64() * screenHeight},
		Velocity: Vector{X: (rand.Float64() * 2) - 1, Y: (rand.Float64() * 2) - 1},
		ID:       bid,
	}
	boids[bid] = &b
	boidMap[int(b.Position.X)][int(b.Position.Y)] = b.ID //at a certain position in the 2d space lies the boid with id as ID
	go b.start()                                         //each bid is excuted on a separte go routine
}

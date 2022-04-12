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
	avgVelocity := Vector{X: 0, Y: 0}
	count := 0.0
	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.ID {
				if distance := boids[otherBoidId].Position.Distance(b.Position); distance < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].Velocity)
				}
			}
		}
	}
	accel := Vector{X: 0, Y: 0}
	if count > 0 {
		avgVelocity = avgVelocity.DivisionV(count)
		accel = avgVelocity.Subtract(b.Velocity).MultiplyV(adjRate)
	}
	return accel
}

func (b *Boid) moveOne() {
	b.Velocity = b.Velocity.Add(b.calcAcceleration()).Limit(-1, 1)
	boidMap[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	boidMap[int(b.Position.X)][int(b.Position.Y)] = b.ID
	next := b.Position.Add(b.Velocity)
	if next.X >= screenWidth || next.X < 0 {
		b.Velocity = Vector{X: -b.Velocity.X, Y: b.Velocity.Y}
	}
	if next.Y >= screenHeight || next.Y < 0 {
		b.Velocity = Vector{X: b.Velocity.X, Y: -b.Velocity.Y}
	}
}

func (b *Boid) start() {
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
	boidMap[int(b.Position.X)][int(b.Position.Y)] = b.ID
	go b.start() //each bid is excuted on a separte go routine
}

package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	Position Vector
	Velocity Vector
	ID       int
}

func (b *Boid) moveOne() {
	b.Position = b.Position.Add(b.Velocity)
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
	go b.start() //each bid is excuted on a separte go routine
}

package particle

import (
	"math"
	"math/rand"
	"time"
)

type Particle struct {
	lifetime int
	speed    float64
	x        float64
	y        float64
}
type ParticleParams struct {
	MaxLife       int64
	MaxSpeed      float64
	X             int64
	Y             int64
	nextPosition  NextPosition
	Ascii         Ascii
	Reset         Reset
	ParticleCount int
}

type ParticleSystem struct {
	Particles []*Particle
	lastTime  int64
	ParticleParams
	// this is like ..
	place func(Particle *Particle, deltaMs int64)
}
type Updatefnc func(p *Particle)
type Reset func(p *Particle, params *ParticleParams)
type NextPosition func(p *Particle, deltaMS int64)
type Ascii func(row int, col int, count [][]int) rune

//this will return an ascii characater accorfing to the position

func NewParticleSystem(params ParticleParams) ParticleSystem {
	particle := make([]*Particle, 0)
	for i := 0; i < params.ParticleCount; i++ {
		particle = append(particle, &Particle{})
	}
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
		Particles:      particle,
	}
}

func (ps *ParticleSystem) Start() {

}

// func ascii(row int, col int, count [][]rune) {

// }

func (p *Particle) Reset(ps *ParticleParams) {
	x := rand.NormFloat64() * float64(ps.X)
	//-x to x
	y := rand.NormFloat64() * float64(ps.Y)
	// -y to y
	p.x = x
	p.y = y
}

// methos of Particle system

func (ps *ParticleSystem) Update() {
	now := time.Now().UnixMilli()
	deltaMs := now - ps.lastTime
	ps.lastTime = now
	//-x-> +x
	for _, p := range ps.Particles {
		if p.x >= float64(ps.X) || p.y >= float64(ps.Y) || p.x <= float64(-ps.X) || p.y <= float64(-ps.Y) || p.lifetime <= 0 {
			ps.Reset(p, &ps.ParticleParams)
		}
		ps.nextPosition(p, deltaMs)
	}
}

func (ps *ParticleSystem) Display() []string {
	rows := ps.Y
	cols := ps.X

	count := make([][]int, rows)
	for i := range count {
		count[i] = make([]int, cols)
	}

	for _, Particle := range ps.Particles {
		row := int(math.Floor(Particle.y))
		col := int(math.Floor(Particle.x))
		if row >= 0 && row < int(rows) && col >= 0 && col < int(cols) {
			count[row][col]++
		}
	}

	// Convert to ASCII representation
	out := make([][]rune, rows)
	for r := range count {
		outRow := make([]rune, cols)
		for c := range count[r] {
			outRow[c] = ps.Ascii(r, c, count)
		}
		out[r] = outRow
	}

	// Flip rows to match expected output order
	for r := 0; r < int(rows/2); r++ {
		out[r], out[int(rows)-r-1] = out[int(rows)-r-1], out[r]
	}

	var outStr []string
	for _, row := range out {
		outStr = append(outStr, string(row))
	}
	return outStr
}

package particle

import (
	"math"
	"math/rand"
)

type Coffe struct {
	ParticleSystem
}

func nextPosition(p *Particle, deltaMs int64) {
	p.lifetime -= int(deltaMs) / 1000.0
	p.y += (float64(deltaMs) * p.speed / 1000.0)
	// fmt.Printf("life %d,speed %f,pos %f", p.lifetime, p.speed, p.y)

}

func reset(p *Particle, params *ParticleParams) {

	p.lifetime = int(math.Floor(float64(9.0) * rand.Float64()))
	p.speed = float64(params.MaxSpeed) * rand.Float64()
	maxX := math.Floor(float64(params.X) / 2)
	//-maxX to maxX
	//0 to 2maxX
	//0 to X
	p.x = math.Max(-maxX, math.Min(rand.NormFloat64()*9, maxX)) + maxX
	p.y = 0
}

var dirs = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},

	{0, -1},
	{0, 1},

	{1, 0},
	{1, 1},
	{1, -1},
}

func countParticles(row, col int, counts [][]int) int {
	count := 0
	for _, dir := range dirs {
		r := row + dir[0]
		c := col + dir[1]
		if r < 0 || r >= len(counts) || c < 0 || c >= len(counts[0]) {
			continue
		}
		count += counts[row+dir[0]][col+dir[1]]
	}
	return count
}

func NewCoffe(width int, height int) Coffe {

	ascii := func(x int, y int, count [][]int) rune {
		counts := count[x][y]
		if counts == 0 {
			return ' '
		}
		if counts < 4 {
			return '░'
		}
		if counts < 6 {
			return '▒'
		}
		if counts < 9 {
			return '▓'
		}
		return '█'
	}

	particleParams := ParticleParams{
		MaxLife:       100,
		MaxSpeed:      0.9,
		X:             int64(width),
		Y:             int64(height),
		Reset:         reset,
		Ascii:         ascii,
		ParticleCount: 700,
		nextPosition:  nextPosition,
	}
	return Coffe{
		ParticleSystem: NewParticleSystem(
			particleParams,
		),
	}
}

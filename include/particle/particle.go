package particle

import (
	"math"

	"github.com/james0248/GoCyclone/include/core"
	errors "github.com/james0248/GoCyclone/include/error"
)

// Particle represents a single particle
type Particle struct {
	position, velocity, acceleration, extForce *core.Vector
	damping, inverseMass                       float64
}

// Integrate calculates the next particle's position and returns an error
func (p *Particle) Integrate(duration float64) error {
	if p.inverseMass <= float64(0.0) {
		return errors.ErrorIntegrateMassInfinite
	}
	if duration <= 0.0 {
		return errors.ErrorIntegrateDurationNegative
	}

	p.position.AddScaledVector(p.velocity, duration)          // dx = vdt
	p.acceleration.AddScaledVector(p.extForce, p.inverseMass) // a = F/m
	p.velocity.AddScaledVector(p.acceleration, duration)      // dv = adt
	p.velocity.MultiplyScalar(math.Pow(p.damping, duration))  // v' = v * damping^duration
	p.extForce.Clear()

	return nil
}

// AddForce adds a new force to the total external force
func (p *Particle) AddForce(force *core.Vector) {
	p.extForce.AddVector(force)
}

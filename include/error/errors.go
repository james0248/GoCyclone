package errors

import "errors"

var (
	// ErrorIntegrateDurationNegative is a error when integrating a negative duration
	ErrorIntegrateDurationNegative = errors.New("Cannot Integrate with a negative duration")
	// ErrorIntegrateMassInfinite is a error when integrating a infinite mass
	ErrorIntegrateMassInfinite = errors.New("Cannot Integrate a particle with infinite mass")
)
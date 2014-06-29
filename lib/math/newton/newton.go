package newton

const (
	SQRT_STEPS = 1000
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < SQRT_STEPS; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

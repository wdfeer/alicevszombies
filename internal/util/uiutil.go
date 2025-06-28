package util

// Returns a slice of 1d points with distance `spacing` from one another
func Space(spacing float32, count int) []float32 {
	points := make([]float32, count)
	for i := range count {
		points[i] = float32(i) * spacing
	}
	return points
}

// Returns a slice of 1d points with distance `spacing` from one another,
// with each point shifted by `offset`
func SpaceOffset(spacing float32, count int, offset float32) []float32 {
	points := Space(spacing, count)
	for i := range points {
		points[i] += offset
	}
	return points
}

// Returns a slice of 1d points with distance `spacing` from one another,
// centered around `center`
func SpaceCentered(spacing float32, count int, center float32) []float32 {
	offset := center - (spacing * float32(count-1) / 2)
	return SpaceOffset(spacing, count, offset)
}

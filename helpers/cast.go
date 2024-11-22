package helpers

func CastValue[T any](value interface{}) T {
	var casted T

	if value != nil {
		casted = value.(T)
	}
	return casted
}

func CastSlice[T any](slice []interface{}) []T {
	casted := make([]T, 0)

	if slice != nil {
		casted = make([]T, len(slice))

		for i, val := range slice {
			casted[i] = val.(T)
		}
	}
	return casted
}

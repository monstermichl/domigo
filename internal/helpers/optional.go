package helpers

func OptionalParams(params ...any) []any {
	collected := make([]interface{}, 0)

	for _, p := range params {
		if p == nil {
			break
		}
		collected = append(collected, &p)
	}
	return collected
}

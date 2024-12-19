package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func CheckIfSlice[T []any](val any) (T, error) {
	casted, ok := val.(T)

	if !ok {
		return []any{}, errors.New("value is not a slice")
	}
	return casted, nil
}

func CheckSliceTypeNames(val any, names []string) error {
	slice, err := CheckIfSlice[[]any](val)

	if err == nil {
		for _, e := range slice {
			err = CheckTypeNames(e, names)

			if err != nil {
				break
			}
		}
	}
	return err
}

func CheckTypeNames(val any, names []string) error {
	tn := reflect.TypeOf(val).Name()

	/* If names contains "number", check for all possible number types. */
	if slices.Contains(names, "number") {
		numTypes := []string{"byte", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64"}
		names = append(names, numTypes...)
	}

	if !slices.Contains(names, tn) {
		return fmt.Errorf("value is none of the following: %s", strings.Join(names, ", "))
	}
	return nil
}

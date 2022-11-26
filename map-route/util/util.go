package util

import "log"

func Must[T interface{}](fn func() (T, error), def T) T {
	v, err := fn()
	if err != nil {
		log.Println(err)
		return def
	}
	return v
}

func Normalise(val, min, max float32) float32 {
	return (val - min) / (max - min)
}

// math module only supports float64 and casting from 32<>64 bits results in conversion error
func MinFloat(i, j float32) float32 {
	if i < j {
		return i
	}
	return j
}

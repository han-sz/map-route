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

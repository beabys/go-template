package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// FindInSlice function to find string inside one slice
// if it's success will return true
func FindInSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// RandomString return a random string
func RandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// RandomInteger return a random Integer, between a range provided
func RandomInteger(min, max int) (int, error) {
	if max <= min {
		return 0, fmt.Errorf("%d can not be smaller than %d", max, min)
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min+1) + min, nil
}

func BindError(toBind, err error) error {
	return errors.Join(toBind, err)
}

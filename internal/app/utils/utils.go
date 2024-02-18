package utils

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"gitlab.com/beabys/quetzal"
)

// InterruptCh on system shut down sends a signal to the caller
func InterruptCh(l quetzal.Logger, serviceName string) chan interface{} {
	// signChan channel is used to transmit signal notifications.
	signChan := make(chan os.Signal, 1)
	// Catch and relay certain signal(s) to signChan channel.
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)

	idleChan := make(chan interface{}, 1)
	go func() {
		sig := <-signChan
		l.Info(fmt.Sprintf("[Shutdown] - %s shutdown: %v", serviceName, sig))
		idleChan <- sig
		close(idleChan)
	}()

	return idleChan
}

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

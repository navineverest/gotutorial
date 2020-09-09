package main

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (t DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, sleeper Sleeper) {

	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)

	}
	sleeper.Sleep()
	fmt.Fprint(w, "Go!")

}

/*
//ex1

package main

import (
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = (t4.Sub(t1)) - (t3.Sub(t2))
	theta = time.Duration((int64(t2.Sub(t1)) + int64(t3.Sub(t4))) / 2)
	return theta, delta
}

func measure(c_clock, s_clock Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3 * time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	
	t2 := s_clock.Now()

	time.Sleep(5 * time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	
	t3 := s_clock.Now()

	time.Sleep(3 * time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}


*/



//Ex2
/*

package main

import (
	"sort"
)

type LamportClock struct {
	counter int
}

func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() int {
	l.counter++
	return l.counter
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct {
	sender int
	ts     int
}

type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

*/

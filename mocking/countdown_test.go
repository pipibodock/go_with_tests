package main

import (
	"time"
	"reflect"
	"testing"
	"bytes"
)

const write = "write"
const sleep = "sleep"

//mock example
type MockCountdownOperations struct {
	Calls [] string
}
func (s *MockCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *MockCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type MockTime struct {
	durationSlept time.Duration
}
func (s *MockTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("test print count", func(t *testing.T){
		buffer := &bytes.Buffer{}
		Countdown(buffer, &MockCountdownOperations{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T){
		mockSleepPrinter := &MockCountdownOperations{}
		Countdown(mockSleepPrinter, mockSleepPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, mockSleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, mockSleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	mockTime := &MockTime{}
	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep: mockTime.Sleep,
	}
	sleeper.Sleep()

	if mockTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, mockTime.durationSlept)
	}
}
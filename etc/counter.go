package etc

import (
	"fmt"
	"sync"
)

type IntCounter struct {
	sync.Mutex
	last int
}

func NewIntCounter(last int) *IntCounter { return &IntCounter{last: last} }

func (i *IntCounter) Next() int {
	i.Lock()
	defer i.Unlock()

	i.last++
	return i.last
}

type Int64Counter struct {
	*IntCounter
}

func NewInt64Counter(last int64) *Int64Counter {
	return &Int64Counter{
		IntCounter: NewIntCounter(int(last)),
	}
}

func (i *Int64Counter) Next() int64 { return int64(i.IntCounter.Next()) }

type StringCounter struct {
	*IntCounter
	label string
}

func NewStringCounter(label string, start int) *StringCounter {
	return &StringCounter{NewIntCounter(start), label}
}

func (s *StringCounter) Next() string {
	return fmt.Sprintf("%s-%d", s.label, s.IntCounter.Next())
}

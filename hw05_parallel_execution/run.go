package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type errorsCounter struct {
	ErrLimit int

	errCount int
	exceeded atomic.Bool
	mx       sync.Mutex
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var result error
	ch := make(chan Task)
	done := make(chan struct{})
	ec := errorsCounter{ErrLimit: m}

	defer func() {
		close(ch)
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-done:
					return
				case task, ok := <-ch:
					if ec.isErrorsExceeded() || !ok {
						return
					}

					if err := task(); err != nil {
						ec.increaseErrorsAndMarkEx()
					}
				}
			}
		}()
	}

	for _, task := range tasks {
		if ec.isErrorsExceeded() {
			break
		}

		ch <- task
	}

	close(done)

	wg.Wait()

	if ec.isErrorsExceeded() {
		result = ErrErrorsLimitExceeded
	}

	return result
}

func (ec *errorsCounter) increaseErrorsAndMarkEx() {
	defer ec.mx.Unlock()
	ec.mx.Lock()

	ec.errCount++
	ec.checkAndMarkEx()
}

func (ec *errorsCounter) checkAndMarkEx() {
	// проверка в функции не защищена от рейса,
	// т.к. в данной реализации вызов checkAndMarkEx происходит
	// из функции где ставиться лок на входе и анлок на выходе
	if ec.ErrLimit == ec.errCount {
		ec.exceeded.Store(true)
	}
}

func (ec *errorsCounter) isErrorsExceeded() bool {
	return ec.exceeded.Load()
}

package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	input := in

	for _, stage := range stages {
		input = doStage(input, done, stage)
	}

	return input
}

func doStage(in In, done In, stage Stage) Out {
	ch := make(Bi)
	out := stage(ch)

	go func() {
		defer func() {
			close(ch)
			for range in { //nolint
			}
		}()

		for {
			select {
			case <-done:
				return
			case value, ok := <-in:
				if ok {
					ch <- value
				} else {
					return
				}
			}
		}
	}()

	return out
}

package async

import "sync"

func InParallel(tasks ...func()) {
	wg := sync.WaitGroup{}
	for _, t := range tasks {
		wg.Add(1)
		t := t // keep ref to task despite loop iterator
		go func() {
			t()
			wg.Done()
		}()
	}
	wg.Wait()
}

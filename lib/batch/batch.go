package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var mutex sync.Mutex
	var waitingGroup sync.WaitGroup
	sem := make(chan struct{}, pool)

	for i := int64(0); i < n; i++ {
		waitingGroup.Add(1)
		sem <- struct{}{}

		go func(i int64) {
			user := getOne(i)
			mutex.Lock()
			res = append(res, user)
			mutex.Unlock()
			<-sem
			waitingGroup.Done()
		}(i)
	}
	waitingGroup.Wait()
	return
}

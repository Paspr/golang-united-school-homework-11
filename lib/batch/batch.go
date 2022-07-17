package batch

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, n)
	errG, _ := errgroup.WithContext(context.Background())
	intPool := int(pool)
	errG.SetLimit(intPool)

	var i int64

	for i = 0; i < n; i++ {
		j := i
		errG.Go(func() error {
			user := getOne(j)
			res[j] = user
			return nil
		})
	}
	err := errG.Wait()
	fmt.Println(err)
	return res
}

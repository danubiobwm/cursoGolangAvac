package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d encerrado...\n", id)
			return
		default:
			fmt.Printf("Worker %d trabalhando...\n", id)
			time.Sleep(time.Millisecond * 500)
		}
	}

}
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(time.Second * 4)
	time.Sleep(time.Second * 5)

}

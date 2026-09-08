package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fazerRequest(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	time.Sleep(time.Second * 6)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil

}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := "https://jsonplaceholder.typicode.com/posts/1"

	dados, err := fazerRequest(ctx, url)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr == context.DeadlineExceeded {
			fmt.Println("Request cancelada devido ao timeout")
		}
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Dados recebidos:", dados)
	}

}

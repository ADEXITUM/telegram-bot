package checker

import (
	"fmt"
	"net/http"
)

func MakeQuery(link string, ch chan<- string) {
	resp, err := http.Get(link)
	if err != nil {
		ch <- fmt.Sprintf("err: %s\n", err)
	} else {
		ch <- fmt.Sprintf("url: %s, status: %d\n", link, resp.StatusCode)
		defer resp.Body.Close()
	}

}

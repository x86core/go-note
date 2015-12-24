package main

import (
	"fmt"
	// "io/ioutil"
	// "net/http"
)

type Op struct {
	IsOk   bool
	Action string
	Output string
}

func main() {
	var links = []string{"http://www.baidu.com", "http://www.sina.com.cn"}
	ch := make(chan Op)
	for _, link := range links {
		go GetRemote(link, ch)
	}

	for i := 0; i < len(links); i++ {
		li := <-ch

		fmt.Println(li.Output)
	}

	close(ch)
}

func GetRemote(link string, ch chan Op) {
	// resp, _ := http.Get(link)

	// output, _ := ioutil.ReadAll(resp.Body)
	ch <- Op{IsOk: true, Action: "getLink", Output: string(link)}
}

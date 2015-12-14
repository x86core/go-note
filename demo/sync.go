package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.baidu.com",
		"http://www.sina.com.cn",
		"http://dict.youdao.com",
	}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			// resp, _ := http.Get(url)
			http.Get(url)

			// out, _ := ioutil.ReadAll(resp.Body)
			// fmt.Println(string(out))

			fmt.Println(url)

		}(url)
	}

	wg.Wait()
	fmt.Println("over")
}

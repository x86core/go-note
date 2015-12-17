package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse form data

	if r.Method == "POST" {
		file, handle, err := r.FormFile("to_save")

		if err != nil {
			fmt.Println(err)
			return
		}

		sf, err := os.OpenFile("./img/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666) //image to write

		io.Copy(sf, file) // save file

		if err != nil {
			fmt.Println(err)
			return
		}

		defer sf.Close()
		defer file.Close()

		fmt.Println("upload OK...")
		w.Write([]byte("all finished")) // response
	} else {
		w.Write([]byte("method err"))
	}
	return
}

func main() {
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":12306", nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("server on ....")
}

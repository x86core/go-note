package main

//使用endless ，平滑重启webserver
import (
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gorilla/mux"

	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WORLD!"))
}

func main() {
	mux1 := mux.NewRouter()
	mux1.HandleFunc("/hello", handler).
		Methods("GET")

	srv := endless.NewServer("localhost:4242", mux1)
	srv.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		log.Printf("Current pid is %d", pid)

		f, err := os.OpenFile("pid.keep", os.O_RDWR|os.O_CREATE|os.SEEK_SET, 0777)

		if err != nil {
			log.Println(err)
			return
		}

		_, err = f.Write([]byte(strconv.Itoa(pid)))
		if err != nil {
			log.Println(err)
			return
		}
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 4242 stopped")

	os.Exit(0)
}

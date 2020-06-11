package main

import (
	"crypto/sha1"
	"fmt"
	logg "log"
	"net/http"
	"os"

	"github.com/tungyao/cedar"
)

var (
	log *logg.Logger
)

func init() {
	fs, err := os.OpenFile("./out.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 766)
	if err != nil {
		log.Fatalln(err)
	}
	log = logg.New(fs, "", logg.LstdFlags|logg.Lshortfile|logg.Ltime)
}
func sha(data string) string {
	t := sha1.New()
	t.Write([]byte(data))
	return fmt.Sprintf("%x", t.Sum(nil))
}
func main() {
	r := cedar.NewRouter()
	r.Get("/login/basic", func(writer http.ResponseWriter, request *http.Request) {
		// fmt.Println(sdk.Get([]byte("a")))
		writer.Write([]byte("hello"))
	}, nil)
	r.Post("/login/basic", func(writer http.ResponseWriter, request *http.Request) {

	}, nil)
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world x"))
	},nil)
	r.Get("/user/:id", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.Fragment)
		writer.Write([]byte(request.URL.Fragment))
	},nil)
	if err := http.ListenAndServe(":6001", r); err != nil {
		log.Fatalln(err)
	}
}

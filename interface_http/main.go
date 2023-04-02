package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(resp)

	//bodySlice := make([]byte, 99999)
	//resp.Body.Read(bodySlice)
	//fmt.Println(string(bodySlice))

	//      writer      reader
	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("length: ", len(bs))
	return len(bs), nil
}

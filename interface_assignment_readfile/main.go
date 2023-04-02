package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// approach 1:
	//arg := os.Args[1]
	//f, err := os.Open(arg)
	//bs := make([]byte, 9999)
	//n1, err := f.Read(bs)
	//if err != nil {
	//	log.Fatalf("unable to read file: %v", err)
	//}
	//fmt.Println(string(bs))
	//fmt.Println("length: ", n1)

	// approach 2:
	arg := os.Args[1]
	f, err := os.Open(arg)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	io.Copy(os.Stdout, f)

}

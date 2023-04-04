package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// approach 1:
	arg := os.Args[1]
	f, err1 := os.Open(arg)
	if err1 != nil {
		log.Fatalf("unable to read file: %v", err1)
	}
	bs := make([]byte, 9999)
	_, err2 := f.Read(bs)
	if err2 != nil {
		log.Fatalf("unable to read file: %v", err2)
	}
	fmt.Println(string(bs))
	//fmt.Println("length: ", n1)

	// approach 2:
	//arg := os.Args[1]
	//f, err := os.Open(arg)
	//if err != nil {
	//	log.Fatalf("unable to read file: %v", err)
	//}
	//
	//io.Copy(os.Stdout, f)

}

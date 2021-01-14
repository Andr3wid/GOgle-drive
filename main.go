package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://andref.xyz")

	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println("An error occured while trying to obtain resources, the error was: ")
		fmt.Println(err)
	}

}

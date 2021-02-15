package main

import (
	"net/http"
	"time"
	"fmt"
)

func main() {
	print_on()
}

func print_on() {
	fmt.Println("on")
	time.Sleep(time.Second)
	fmt.Println("off")
}

func curtain_on() {
	http.Get("http://192.168.127.17/ledon")
	time.Sleep(time.Second * 20)
	http.Get("http://192.168.127.17/ledoff")
}

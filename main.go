package main

import (
	"fmt"
	"study/http_server"
)

func main() {
	fmt.Println("Сервер начал работу!")
	err := http_server.StartHttpServer()
	if err != nil {
		fmt.Println("Во время работы сервера произошла ошибка:",err)
	} 
	fmt.Println("Сервер остановил свою работу!")
}
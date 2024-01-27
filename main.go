package main

import server "github.com/JrSchmidtt/go-react-ssr-boilerplate/src"

func main() {
	if err := server.Run(); err != nil {
		panic(err)
	}
}
package main

import (
	"flag"
	"fmt"
	"goServer/internal/flags"
)

func main(){
	c := flags.Config{}
	c.Setup()

	flag.Parse()

	fmt.Println(c.GetMessage())
}
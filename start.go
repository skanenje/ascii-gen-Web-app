package main

import (
	"os"



)

func flags1(){
	args := os.Args[1:]
	var filename string
	for i:=0 ; i< len(args) ; i++{
		words := args[i]
		if words[:9] == "--output="{
			filename = words[9:]
		}
	}

}
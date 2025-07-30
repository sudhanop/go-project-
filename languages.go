package main

import (
	"fmt"
	"strings"
)
func main(){
	var language string
	fmt.Println("enter the language you know:")
	fmt.Scanln(&language)
	language = strings.ToLower(language)
	//using a switch
	switch language {
	case "tamil":
		fmt.Println("vanakkam")
	case "english":
		fmt.Println("hello")
	case "hindi":
		fmt.Println("namaste")
		default :
		fmt.Println("sorry your language not available")
	}
	}

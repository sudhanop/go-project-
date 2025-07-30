package main

import (
	"fmt"
	"os"
)
func main() {
	content:= "this file should be created by go language"
	err:=os.WriteFile("sudhan1.txt", []byte(content), 0644)
if err != nil {
	fmt.Println("error while creating file",err)
	return
}
 fmt.Println("file created successfully")
 data,err:=os.ReadFile("sudhan1.txt")
 if err != nil{
	fmt.Print("error reading file",err)
	return
 }
 fmt.Println("file content :")
 fmt.Println(string(data))
}
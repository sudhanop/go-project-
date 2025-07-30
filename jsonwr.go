package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct{
	Name  string `json:"name"`
	Age int `json:"age"`
}

func main(){
	person := person{
		Name : "sudhan",
		Age : 20,
	}


jsondata,err:=json.MarshalIndent(person,"","  ")
if err != nil{
	fmt.Println("error in file json making",err)
	return
}
err =os.WriteFile("sudhan.json",jsondata,0644)
if err != nil{
	fmt.Println("error in wwriting the flie",err)
	return
}
fmt.Println("file sudhan.json successfully created")
}
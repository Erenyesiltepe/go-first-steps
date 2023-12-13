package main

import (
	"fmt"
	"plugin"
)

func main(){
    p, err := plugin.Open("example.so")
    if err != nil {
        panic(err)
    }
    f, err := p.Lookup("Add")
    if err != nil {
        panic(err)
    }
    var num1,num2 int
    fmt.Printf("Please enter numbers to sum up:")
    fmt.Scanf("%d %d",&num1,&num2)
    result:=f.(func(int,int)int)(num1,num2) 
    fmt.Println((result))
}
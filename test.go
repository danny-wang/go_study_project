package main

import (

	"fmt"

	"sort"
)
func triple(x int)(result int){
	defer func() {result += x}()
	return x
}

func main(){

	fmt.Println(32 << uint(^uint(0)>>63))
	fmt.Println(^uint(0))
	fmt.Println(^uint(0)>>63)
	fmt.Println(32<<1)
	sort.Strings()
}
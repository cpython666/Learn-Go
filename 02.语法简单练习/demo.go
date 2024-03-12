package main

import "fmt"

func main() {
	var mapInit = make(map[string]string, 1)
	mapInit["xiaoli"] = "湖南"
	mapInit["xiaoliu"] = "天津"
	mapInit["123"] = "456"
	l := len(mapInit)
	fmt.Println(l)
}

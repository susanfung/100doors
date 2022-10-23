package main

import "fmt"

func InitialArray() []string {
    doors := make([]string, 100)
		for i := range doors {
			doors[i] = "closed"
		}
		return doors
}

func Answer() []string {
	doors := InitialArray()
	return doors
}

func main() {
    fmt.Println(Answer())
}
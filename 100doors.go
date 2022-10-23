package main

import "fmt"

func InitialArray() []string {
    doors := make([]string, 100)
		for i := range doors {
			doors[i] = "closed"
		}
		return doors
}

func Answer(pass int) []string {
	doors := InitialArray()
	door := 0
	if pass > 0 {
		for door < len(doors) {
		if doors[door] == "closed" {
			doors[door] = "opened"
		} else {
			doors[door] = "closed"
		}
		door += pass
	}
	}
	return doors
}

func main() {
    fmt.Println(Answer(0))
}
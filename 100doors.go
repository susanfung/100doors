package main

import "fmt"

func InitialArray() []string {
    doors := make([]string, 100)
		for i := range doors {
			doors[i] = "closed"
		}
		return doors
}

func ToggleDoors(doors []string, pass int) []string {
	door := pass - 1

	for door < len(doors) {
		if doors[door] == "closed" {
			doors[door] = "opened"
		} else {
			doors[door] = "closed"
		}
		door += pass
	}

	return doors
}

func Answer(pass int) []string {
	doors := InitialArray()
	
	if pass > 0 {
		for i := 1; i <= pass; i++ {
			doors = ToggleDoors(doors, i)
		}
	}

	return doors
}

func main() {
    fmt.Println(Answer(0))
}
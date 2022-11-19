package main

import "fmt"

func main() {
	teams := map[string][]string{
		"Ajay":   {"Bechara", "22"},
		"Jaydip": {"Marvaniya", "22"},
		"Jeel":   {"Ranipa", "22"},
	}
	teams["new"] = []string{"1", "2", "3", "4"}
	fmt.Println(teams, len(teams))
	key, ok2 := teams["kkk"]
	fmt.Println(key, ok2)
	delete(teams, "Ajay")
	delete(teams, "lll")
	fmt.Println(teams, len(teams))

	ages := make(map[int][]string, 10)
	fmt.Println(ages, len(ages))
}

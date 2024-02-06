package main

import "fmt"

func main() {
	order := map[string][]string{
		"batch1": {
			"brownies",
			"nastar",
			"muffin",
			"muffin",
			"brownies",
			"pie",
			"nastar",
			"muffin",
		},
		"batch2": {
			"brownies",
			"pie",
			"pie",
			"nastar",
		},
	}

	countItemPerKey := make(map[string]map[string]int)
	for k, v := range order {
		fmt.Println("Processing: ", k)
		countItemPerKey[k] = make(map[string]int)
		for _, kr := range v {
			countItemPerKey[k][kr]++
		}

		cook("brownies", countItemPerKey[k]["brownies"])
		cook("nastar", countItemPerKey[k]["nastar"])
		if k == "batch1" {
			cook("muffin", countItemPerKey[k]["muffin"])
		}
		cook("pie", countItemPerKey[k]["pie"])
	}
}

func cook(cakeRecipe string, totalOrder int) {
	fmt.Println(fmt.Sprintf("%d %s is ready!", totalOrder, cakeRecipe))
}

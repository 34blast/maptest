package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("maptest.main() : starting exectuion")
	fmt.Println()

	var stateCapitalMap map[string]string

	/* create a map*/
	stateCapitalMap = make(map[string]string)

	stateCapitalMap["Texas"] = "Austin"
	stateCapitalMap["Illinois"] = "Springfield"
	stateCapitalMap["Lousianna"] = "Baton Rouge"
	stateCapitalMap["North Dakota"] = "Bismark"
	stateCapitalMap["South Dakota"] = "Pierre"
	stateCapitalMap["North Carolina"] = "Raleigh"
	stateCapitalMap["California"] = "Sacramento"

	var tableKeyMap map[int]string
	tableKeyMap = make(map[int]string)

	for idx := 0; idx < 5; idx++ {
		tableKeyMap[idx] = "test" + strconv.Itoa(idx)
	}

	fmt.Println("stateCapitalMap = ", stateCapitalMap)
	fmt.Println("tableKeyMap = ", tableKeyMap)
	fmt.Println()

	/* print map using keys*/
	for state := range stateCapitalMap {
		fmt.Println("Capital of", state, "is", stateCapitalMap[state])
	}
	fmt.Println()

	_, okay := stateCapitalMap["California"]
	if okay {
		delete(stateCapitalMap, "California")
	}
	_, okay = stateCapitalMap["abc"]
	if okay {
		delete(stateCapitalMap, "abc")
	}

	/* print map using keys*/
	for state := range stateCapitalMap {
		fmt.Println("Capital of", state, "is", stateCapitalMap[state])
	}
	fmt.Println()

	/* test if entry is present in the map or not*/
	capital, ok := stateCapitalMap["North Dakota"]

	/* if ok is true, entry is present otherwise entry is absent*/
	if ok {
		fmt.Println("Capital of North Dakota is", capital)
	} else {
		fmt.Println("Capital of North Dakota is not present")
	}
	fmt.Println()

	capital, ok = stateCapitalMap["New Mexico"]

	/* if ok is true, entry is present otherwise entry is absent*/
	if ok {
		fmt.Println("Capital of New Mexico is", capital)
	} else {
		fmt.Println("Capital of New Mexico is not present")
	}
	fmt.Println()

	fmt.Println("maptest.main() : ending  exectuion")

}

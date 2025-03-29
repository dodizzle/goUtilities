package maps

import "fmt"

func Maps() {
	// define map
	sbWins := map[string]int{}

	// add to map
	sbWins["Patriots"] = 7
	sbWins["Steelers"] = 4

	fmt.Println(sbWins)
	// Print specific match
	fmt.Println("Patriots=", sbWins["Patriots"])
	// delete from map
	delete(sbWins, "Steelers")
	fmt.Println(sbWins)
}

package structs

import "fmt"

func Structs() {

	type sbRings struct {
		name  string
		count int
	}

	tbRings := sbRings{
		name:  "Brady",
		count: 7,
	}

	jmRings := sbRings{}
	jmRings.name = "Montana"
	jmRings.count = 4

	fmt.Println(tbRings)
	fmt.Println(jmRings)
}

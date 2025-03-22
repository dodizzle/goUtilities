package main

import (
	"fmt"
	"math/rand"
)

const (
	one int = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	eleven
	twelve
	thirteen
	fourteen
	fifteen
	sixteen
)

func main() {
	round1 := map[int]int{
		one:   16,
		two:   15,
		three: 14,
		four:  13,
		five:  12,
		six:   11,
		seven: 10,
		eight: 9,
	}
	var regions []string
	regions = append(regions, "==SOUTH==", "==WEST==", "==EAST==", "==MIDWEST==")
	for _, v := range regions {
		fmt.Println(v)
		round1Results := getRound(round1, "Round 1")
		round2 := setRoundTwo(round1Results)
		round2Results := getRound(round2, "Round 2")
		round3 := setRoundThree(round2Results)
		round3Results := getRound(round3, "Round 3")
		round4 := setRoundFour(round3Results)
		getRound(round4, "Round 4")
	}
}

func setRoundFour(roundResults map[int]int) map[int]int {
	nextRound := make(map[int]int)
	var x []int
	for _, v := range roundResults {
		x = append(x, v)
	}
	nextRound[x[0]] = x[1]
	return nextRound
}

func setRoundThree(roundResults map[int]int) map[int]int {
	nextRound := make(map[int]int)
	var x []int
	var a, b, c, d int
	for _, v := range roundResults {
		x = append(x, v)
	}
	for _, v := range x {
		if v == 1 || v == 8 || v == 16 || v == 9 {
			a = v
		} else if v == 5 || v == 12 || v == 4 || v == 13 {
			b = v
		} else if v == 6 || v == 11 || v == 3 || v == 14 {
			c = v
		} else if v == 7 || v == 10 || v == 2 || v == 15 {
			d = v
		}
	}
	nextRound[a] = b
	nextRound[c] = d
	return nextRound
}

func setRoundTwo(round1Results map[int]int) map[int]int {
	nextRound := make(map[int]int)
	a := round1Results[1]
	b := round1Results[8]
	c := round1Results[5]
	d := round1Results[4]
	e := round1Results[6]
	f := round1Results[3]
	g := round1Results[7]
	h := round1Results[2]
	nextRound[a] = b
	nextRound[c] = d
	nextRound[e] = f
	nextRound[g] = h
	return nextRound
}

func getRound(round map[int]int, roundName string) map[int]int {
	fmt.Println(roundName)
	roundResults := make(map[int]int)

	for k, v := range round {
		winner := getWinner(k, v)
		fmt.Printf("%d vs %d = Winner is %d\n", k, v, winner)
		roundResults[k] = winner
	}
	return roundResults
}

func getWinner(team1 int, team2 int) (winner int) {
	sum := getSum(team1, team2)
	randomSeed := rand.Intn(sum)
	higherSeed, lowerSeed := rankSeeds(team1, team2)
	if randomSeed >= higherSeed {
		return higherSeed
	}
	return lowerSeed
}

func getSum(a int, b int) int {
	return a + b
}

func rankSeeds(a int, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

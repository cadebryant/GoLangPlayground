package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numTrials := 10000000
	fmt.Println(ComputeCrapsHouseEdge(numTrials))
}

func ComputeCrapsHouseEdge(numTrials int) float64 {
	count := 0
	for i := 0; i < numTrials; i++ {
		outcome := PlayCrapsOnce()
		if outcome {
			count++
		} else {
			count--
		}
	}
	return float64(count) / float64(numTrials)
}

func PlayCrapsOnce() bool {
	firstRoll := SumTwoDice()
	if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		return false
	}
	if firstRoll == 7 || firstRoll == 11 {
		return true
	}
	for {
		newRoll := SumTwoDice()
		if newRoll == firstRoll {
			return true
		}
		if newRoll == 7 {
			return false
		}
	}
}

func SumTwoDice() int {
	return SumDice(2)
}

func SumDice(numDice int) int {
	sum := 0
	for i := 0; i < numDice; i++ {
		sum += RollDie()
	}
	return sum
}

func RollDie() int {
	roll := rand.Intn(6)
	return roll + 1
}

func AddNoise(pollValue float64, marginOfError float64) float64 {
	randF := rand.NormFloat64()
	randF /= 2
	randF *= marginOfError
	return pollValue + randF
}

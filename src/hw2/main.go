package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// vals := make(map[int]int)
	// for i := 0; i < 1000; i++ {
	// 	roll := WeightedDie()
	// 	if roll == 0 {
	// 		continue
	// 	}
	// 	_, ok := vals[roll]
	// 	if ok {
	// 		vals[roll]++
	// 	} else {
	// 		vals[roll] = 1
	// 	}
	// }
	// fmt.Println(vals)

	// lowerBound := 100000
	// upperBound := 200000

	// x := rand.Intn(upperBound-lowerBound+1) + lowerBound
	// y := rand.Intn(upperBound-lowerBound+1) + lowerBound

	// //timing TrivialGCD()
	// start := time.Now()
	// TrivialGCD(x, y)
	// elapsed := time.Since(start)
	// log.Printf("TrivialGCD() took %s", elapsed)

	// //timing EuclidGCD()
	// start2 := time.Now()
	// EuclidGCD(x, y)
	// elapsed2 := time.Since(start2)
	// log.Printf("EuclidGCD() took %s", elapsed2)
	// numPoints := 10000000
	// fmt.Println(EstimatePi(numPoints))
	// a, b := 4, 9
	// fmt.Println(RelativelyPrime(a, b))
	// lowerBound := 1
	// upperBound := 1000000000000
	// numPairs := 1000000000
	// fmt.Println(RelativelyPrimeProbability(lowerBound, upperBound, numPairs))
	// numPeople := 100
	// numTrials := 100000
	// for i := 1; i <= numPeople; i++ {
	// 	fmt.Println(i, SharedBirthdayProbability(i, numTrials))
	// }
	// nums := []int{0, -1, -498212910, 155, 1000000, 999999}
	// for _, x := range nums {
	// 	fmt.Println(CountNumDigits(x))
	// }
	// seed := 1
	// numDigits := 2
	// count := 0
	// for ; seed <= 9999; seed++ {
	// 	seq := GenerateMiddleSquareSequence(seed, numDigits)
	// 	len := ComputePeriodLength(seq)
	// 	if len <= 10 {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)
	// a := []int{140, 278, 14, 28, 19, 28}
	// fmt.Println(ComputePeriodLength(a))
	// seed := 1
	// a := 2
	// c := 0
	// m := 9
	// fmt.Println(GenerateLinearCongruenceSequence(seed, a, c, m))
	seed := 1
	a := 5
	c := 1
	m := 8192
	seq := GenerateLinearCongruenceSequence(seed, a, c, m)
	fmt.Println(ComputePeriodLength(seq))
}

func GenerateLinearCongruenceSequence(seed, a, c, m int) []int {
	seq := []int{seed}
	for !HasRepeat(seq) {
		seed = (a*seed + c) % m
		seq = append(seq, seed)
	}

	return seq
}

func ComputePeriodLength(a []int) int {
	nums := make(map[int]int)
	len := len(a)
	for i := 0; i < len; i++ {
		foundIdx, ok := nums[a[i]]
		if ok {
			return i - foundIdx
		}
		nums[a[i]] = i
	}

	return 0
}

func GenerateMiddleSquareSequence(seed, numDigits int) []int {
	seq := []int{seed}
	for !HasRepeat(seq) {
		seed = SquareMiddle(seed, numDigits)
		seq = append(seq, seed)
	}
	return seq
}

func SquareMiddle(x, numDigits int) int {
	if numDigits%2 != 0 || x < 0 || numDigits < 1 || CountNumDigits(x) > numDigits {
		return -1
	}
	x *= x
	num := float64(x % int(math.Pow10((2*numDigits)-(numDigits/2))))
	denom := math.Pow10(numDigits / 2)
	return int(num / denom)
}

func CountNumDigits(x int) int {
	if x == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(x))))) + 1
}

func SharedBirthdayProbability(numPeople, numTrials int) float64 {
	sharedBirthdays := 0
	for i := 0; i < numTrials; i++ {
		if SimulateOneBirthdayTrial(numPeople) {
			sharedBirthdays += numPeople
		}
	}
	return float64(sharedBirthdays) / (float64(numTrials) * float64(numPeople))
}

func SimulateOneBirthdayTrial(num_people int) bool {
	birthdays := make([]int, num_people)
	for i := 0; i < num_people; i++ {
		birthdays[i] = rand.Intn(365) + 1
	}
	return HasRepeat(birthdays)
}

func HasRepeat(a []int) bool {
	nums := make(map[int]bool)
	for _, num := range a {
		_, ok := nums[num]
		if ok {
			return true
		}
		nums[num] = true
	}
	return false
}

func RelativelyPrimeProbability(lowerBound, upperBound, numPairs int) float64 {
	relPrimeCount := float64(0)

	for i := 0; i < numPairs; i++ {
		a := rand.Intn(upperBound-lowerBound+1) + lowerBound
		b := rand.Intn(upperBound-lowerBound+1) + lowerBound

		if RelativelyPrime(a, b) {
			relPrimeCount++
		}
	}

	return relPrimeCount / float64(numPairs)
}

func RelativelyPrime(a, b int) bool {
	return EuclidGCD(a, b) == 1
}

func EstimatePi(numPoints int) float64 {
	inCircCount := float64(0)
	for i := 0; i < numPoints; i++ {
		a := rand.Float64()*2 - 1
		b := rand.Float64()*2 - 1

		if (a*a)+(b*b) <= 1 {
			inCircCount++
		}
	}
	return (inCircCount / float64(numPoints)) * 4
}

func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func TrivialGCD(a, b int) int {
	d := 1
	m := Min2(a, b)
	for p := 1; p <= m; p++ {
		if a%p == 0 && b%p == 0 {
			d = p
		}
	}
	return d
}

func WeightedDie() int {
	valueWeights := map[int]float64{
		1: 0.1,
		2: 0.1,
		3: 0.5,
		4: 0.1,
		5: 0.1,
		6: 0.1,
	}

	rnd := rand.Float64()
	for v, w := range valueWeights {
		if rnd < w {
			return v
		}
		rnd -= w
	}

	return 0
}

func Min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

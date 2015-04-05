package prime

// init one bool slice which length is n and value is true
// from i in 0->n
// 	if i is prime
// make all number less than n but multiply of i false
func AllPrimeBelowN(n int) []int {
	primyMatrix := make([]bool, n-1, n-1)
	length := len(primyMatrix)
	for i := 0; i < length; i++ {
		primyMatrix[i] = true
	}
	primeSlice := make([]int, 0, n/10)
	for i := 0; i < length; i++ {
		if !primyMatrix[i] {
			continue
		}
		number := i + 2
		isPrime := judgePrimeByPrimyList(number, primeSlice)
		if !isPrime {
			primyMatrix[i] = false
			continue
		}
		primeSlice = append(primeSlice, number)
		for number <= n {
			primyMatrix[number-2] = false
			number += number
		}
	}
	return primeSlice
}

// is prime
// if n is not prime
// the number must be under log2(n)
func judgeIsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// all number can be create by prime,
// so best way to judge a prime ,is divide by all the prime less than it,
// notice primeSlice must be right
func judgePrimeByPrimyList(n int, primeSlice []int) bool {
	if n < 2 {
		return false
	}
	if n <= 4 {
		return true
	}
	for _, number := range primeSlice {
		if n%number == 0 {
			return false
		}
	}
	return true
}

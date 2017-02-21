package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	series := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	fmt.Println(GetMaxValueOfArray(SlicesOfNumber(series, 13)))

}

func GetMaxValueOfArray(numbers []int) int {
	max := numbers[0]
	for _, value := range numbers {
		if (value > max) {
			max = value
		}
	}
	return max
}

func FiboSumOfGivenRange(limit, divisibleBy int) int {
	f := fibo();
	result := 0;
	r := 0;
	for r < limit {
		r = f();
		if (r < limit && r % divisibleBy == 0) {
			result += r;
		}
	}
	return result;
}

func LargestPrimeFactor(num int) []int {
	slice := make([]int, 0)
	var d int = 2;
	for num > 1 {
		for num % d == 0 {
			num = num / d;
			slice = append(slice, d)
		}
		d = d + 1;
	}

	return slice;
}

func SumOfMultipleOf3Or5(upperLimit int) int {
	result := 0;
	for i := 0; i < upperLimit; i++ {
		if (i % 3 == 0 || i % 5 == 0) {
			result += i;
		}
	}
	return result;
}

func fibo() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a + b;
		return a;
	}
}

func NumdivisibeByGivenRange(lowerLimit int, upperLimit int) int {
	num := 11
	for num < 300000000 {
		if (isNumDivisibleByRange(lowerLimit, upperLimit, num)) {
			return num
		}
		num++
	}
	return num;
}

func isNumDivisibleByRange(lowerLimit int, upperLimit int, num int) bool {
	for i := lowerLimit; i < upperLimit; i++ {
		if (num % i != 0) {
			return false
		}
	}
	return true
}

func SumSquareDifference(upperLimit int) int {
	return int(math.Pow(float64(upperLimit * (upperLimit + 1) / 2), 2)) - upperLimit * (upperLimit + 1) * (2 * upperLimit + 1) / 6
}

func NthPrime(n int) int {
	counter := 0;
	for i := 1; i < 1000000; i++ {
		if (IsPrime(i)) {
			counter++
		}
		if (counter == n) {
			return i;
		}
	}
	return 0
}

func IsPrime(num int) bool {
	sqrtPlusOne := int(math.Floor(math.Sqrt(float64(num))) + 1)
	for i := 2; i <= sqrtPlusOne; i++ {
		if (num % i == 0) {
			return false;
		}
	}
	return true
}

func SlicesOfNumber(chunk string, rangeOfDigit int) []int {
	var slices []string
	for i := 0; i <= len(chunk) - rangeOfDigit; i++ {
		eachString := ""
		for j := i; j < i + rangeOfDigit; j++ {
			eachString += string(chunk[j])
		}
		slices = append(slices, eachString)
	}
	return MultiplicationOfDigits(slices, rangeOfDigit)
}

func MultiplicationOfDigits(chunk []string, rangeOfDigit int) []int {
	var multiplicationOfDigits []int
	var p int64 = 1
	for i := 0; i < len(chunk); i++ {
		for j := 0; j < rangeOfDigit; j++ {
			number, _ := strconv.ParseInt(string(chunk[i][j]), 10, 0)
			p *= number
		}
		multiplicationOfDigits = append(multiplicationOfDigits, int(p))
		p = 1
	}
	return multiplicationOfDigits
}



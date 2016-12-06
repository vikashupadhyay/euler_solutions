package main

func main() {
	//println(GetMaxValueOfArray(LargestPrimeFactor(600851475143)))
	isPalindrome(123)
}

func isPalindrome(number int) bool {
	a := string(number)
	println(a);
	return true;
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


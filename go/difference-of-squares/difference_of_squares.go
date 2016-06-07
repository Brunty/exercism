package diffsquares

func SquareOfSums(number int) int {
    var sum int

    for i := 0; i <= number; i++ {
        sum += i
    }

    return sum * sum
}

func SumOfSquares(number int) int {
    var total int

    for i := 0; i <= number; i++ {
        total += i * i
    }

    return total
}

func Difference(number int) int {
    return SquareOfSums(number) - SumOfSquares(number)
}


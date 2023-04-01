package math

func Sum(x int, y int) int {
	return x + y
}

func Diff(x int, y int) int {
	return x - y
}

func Product(x int, y int) int {
	return x * y
}

func Divide(x int, y int) int {
	return x / y
}

func Remainder(x int, y int) int {
	return x % y
}

// ******* Fibonacci btn start - end *******
func Fibonacci(start int, end int) []int {
	var seq []int // list to hold the febonaci sequence

	// generate sequence
	for seq = []int{0, 1}; seq[len(seq)-1] < end; seq = append(seq, seq[len(seq)-2]+seq[len(seq)-1]) {
	}

	seq = seq[:len(seq)-1]             // remove last to trim sequence to range
	seqCopy := (make([]int, len(seq))) // Declaring sequence to copy
	copy(seqCopy, seq)                 // make a copy of sequence

	for count := range seqCopy {
		if seqCopy[count] < start {
			seq = seq[1:] // remove first to trim sequence to range
		}
	}

	return seq
}

// ******* Add Fibonacci btn start - end *******
func AddFibonacci(start int, end int) int {
	var sum int = 0
	var sumSeq []int = Fibonacci(start, end)

	// Adding seq
	for _, value  := range sumSeq {
		sum += value
	}

	return sum
}

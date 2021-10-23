package piscine

func Fibonacci(index int) int {
	if index < 0 {
		print("\nJe passe ici\n", index)
		return -1
	}
	if index == 0 {
		print("\nJe passe ici\n", index)
		return 0
	}
	if index == 1 {
		print("\nJe passe ici\n", index)
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}

package fib

func Fib(n int) int {

	/*version 1 return 0
	 */

	/* version 2 if n == 0 {
		return 0
	}
	return 1*/

	//version 3

	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

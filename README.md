# Fibonacci utregning
F(0) = 0

#respons 1

.\fib_test.go:7:9: undefined: Fib
FAIL    github.com/OlePaulsen/fib/fib [build failed]
FAIL

#respons 2

=== RUN   TestFib
--- PASS: TestFib (0.00s)
PASS
ok      github.com/OlePaulsen/fib/fib   0.182s

#respons 3

=== RUN   TestFib
    fib_test.go:11: Fib(0): want 1, got 0
--- FAIL: TestFib (0.00s)
FAIL
FAIL    github.com/OlePaulsen/fib/fib   0.184s
FAIL

#respons 4

# github.com/OlePaulsen/fib/fib [github.com/OlePaulsen/fib/fib.test]
.\fib_test.go:7:2: undefined: input
.\fib_test.go:8:13: undefined: input
.\fib_test.go:12:40: undefined: input
FAIL    github.com/OlePaulsen/fib/fib [build failed]
FAIL

#respons 5

=== RUN   TestFib
--- PASS: TestFib (0.00s)
PASS
ok      github.com/OlePaulsen/fib/fib   0.213s

#respons 6 
=== RUN   TestFib
--- PASS: TestFib (0.00s)
PASS
ok      github.com/OlePaulsen/fib/fib   0.204s

#respons 7
=== RUN   TestFib
    fib_test.go:23: Fib(3): want 2, got 1
--- FAIL: TestFib (0.00s)
FAIL
FAIL    github.com/OlePaulsen/fib/fib   0.188s
FAIL

#respons 8 

=== RUN   TestFib
--- PASS: TestFib (0.00s)
PASS
ok      github.com/OlePaulsen/fib/fib   0.195s

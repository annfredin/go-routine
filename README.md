# GO ROUTINE:

The main function in the main package is the main goroutine. All goroutines are started from the main goroutine

Goroutines don’t have parents or children. When you start a goroutine it just executes alongside all other running goroutines. Each goroutine exits only when its function returns. The only exception to that is that all goroutines exit when the main goroutine (the one that runs function main) exits

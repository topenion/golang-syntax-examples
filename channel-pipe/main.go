package main

import "fmt"

func receiver(nums ...int) <-chan int {
	passer := make(chan int)
	go func() {
		for _, n := range nums {
			passer <- n
		}
		close(passer)
	}()
	return passer
}

func calculator(in <-chan int) <-chan int {
	passer := make(chan int)
	go func() {
		for n := range in { //single channel listening
			passer <- n * n
		}
		close(passer)
	}()
	return passer
}
func main() {
	// Set up the pipeline.
	valschan := receiver(2, 3, 78, 98998, 90909, 344, 89)
	calcchan := calculator(valschan)

	// Consume the output.
	fmt.Println(<-calcchan)
	fmt.Println(<-calcchan)
	fmt.Println(<-calcchan)
	fmt.Println(<-calcchan)
	fmt.Println(<-calcchan)
	fmt.Println(<-calcchan)
	fmt.Println(<-calcchan)
}

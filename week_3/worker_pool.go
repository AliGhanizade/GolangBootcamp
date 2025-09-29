package week_3

import (
	"fmt"

)

const(
	numWorkers int = 5
 	numJobs int = 5000
) 

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker start. id : %d job : %d \n", id, j)
		result := j^2 
		fmt.Printf("Worker finish. Id : %d job int : %d result : %d\n", id, j, result)
		
		results <- result
	}
}

func StartWorkerPool() {
	var jobs = make(chan int, numJobs) 
	var results = make(chan int, numJobs) 

	

	fmt.Printf("%d Worker ready for jobs.\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	fmt.Printf("%d Jobs.\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) 

	var totalResults int
	for r := 1; r <= numJobs; r++ {
		result := <-results 
		totalResults += result
	}

	close(results) 

	fmt.Printf("total : %d\n", totalResults)
	
}
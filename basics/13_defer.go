package main

import "fmt"

func cleanup(name string) {
	fmt.Println("Cleaning up resource : ", name)
}

func worker() {
	// some code to acquire the resource
	defer cleanup("A") // call defer after acquiring resource to handle cleanups

	// some code to acquire another resource
	defer cleanup("B")

	// perform tasks using the resource
	fmt.Println("Performing worker tasks...")
}

func main() {
	// go does garbage collection on resources after its usage is completed
	// its good practice to call defer when new resources are acquired, the garbage collector will call this method to during clean up process
	worker()

	// notice after worker tasks are completed, the defer calls were executed in the reverse order (stack LIFO model)
}

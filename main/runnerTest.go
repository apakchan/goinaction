package main

import (
	"../runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(id) * time.Second)
		log.Printf("Process - Task #%d Done.\n", id)
	}
}

func main() {
	log.Println("Starting Work")
	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended")
}

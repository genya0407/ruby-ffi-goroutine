package main

import "C"
import "fmt"
import "github.com/google/uuid"
import "time"

//export poll
func poll() {
	for i := 0; i < 2; i++ {
		go func() {
			id, err := uuid.NewRandom()
			if err != nil {
				panic(err)
			}

			for {
				fmt.Printf("hoge %v\n", id);
				time.Sleep(1 * time.Second)
			}
		}()
	}
}

func main() {
	poll()
	time.Sleep(100 * time.Second)
}
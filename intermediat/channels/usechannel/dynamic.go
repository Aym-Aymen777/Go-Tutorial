package usechannel

import (
	"fmt"
	"time"
)

func worker(name string , loops int ,delay time.Duration , done chan string){
   
	for i:=1; i<=loops; i++{
		fmt.Printf("%s is working... step %d\n",name,i)
		time.Sleep(delay)
	}
	done <- fmt.Sprintf("%s finished",name)
}

func Usechannel(){
	 done := make(chan string)

    go worker("Worker 1", 5, 300*time.Millisecond, done)
    go worker("Worker 2", 4, 500*time.Millisecond, done)
    go worker("Worker 3", 6, 200*time.Millisecond, done)

    // wait for 3 finish messages
    for i := 0; i < 3; i++ {
        msg := <-done
        fmt.Println(msg)
    }

    fmt.Println("All workers finished. Main over.")
}

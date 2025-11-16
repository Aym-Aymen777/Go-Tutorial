package usetimesleep

import (
	"fmt"
	"time"
)

func worker(name string,loops int,delay time.Duration){
	for i:=1;i<=loops;i++{
		fmt.Printf("%s working... step %d\n",name,i)
		time.Sleep(delay)
	}
	fmt.Printf("%s finished!\n",name)
}

func Usetimesleep(){
	start:= time.Now()

	go worker("Worker 1",5,300*time.Millisecond)
	go worker("Worker 2",4,500*time.Millisecond)
	go worker("Worker 3",6,200*time.Millisecond)

	time.Sleep(1*time.Second)

	fmt.Println("Usetimesleep over (but maybe workers not finished)")
    
	elapsed := time.Since(start)
    fmt.Println("Execution time:", elapsed)
}

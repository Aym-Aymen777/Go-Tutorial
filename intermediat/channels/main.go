package main
import (
	"fmt"
	"channels/usetimesleep"
	"channels/usechannel"
)

func main() {
	choice:=0
	fmt.Println("what do you want to try")
	fmt.Println("1. gouroutine with time.sleep")
	fmt.Println("2. gouroutine with channels synchronization")
	fmt.Println("3. exit")

	fmt.Scan(&choice)
	switch choice {
	case 1:usetimesleep.Usetimesleep()
	case 2:usechannel.Usechannel()
	case 3: fmt.Print("good luck")
	}

}
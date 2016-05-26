package main

import (
	"github.com/iron-io/iron_go3/mq"
	"fmt"
	"strconv"
)


func main()  {
	queue := mq.New("messages");

	for i := 0; i < 10000; i++  {
			fmt.Println("Pusing object: "+ strconv.Itoa(i))
			_, err := queue.PushString("Hello, World!")
			if err != nil {
				fmt.Printf(err.Error())
			}
		}
}
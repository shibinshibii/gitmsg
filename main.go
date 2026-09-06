package main

import (
	"fmt"
)

func main(){
	var msgtype int
	var msg string

	fmt.Print(`Choose a type:

1. feat
2. fix
3. refactor
4. docs
5. test
6. chore
`)
	_,err := fmt.Scan(&msgtype)
	if err != nil {
		fmt.Println("Some error occured", err)
	}

	switch msgtype {
	case 1:
		fmt.Println("feat selected!")
	case 2:
		fmt.Println("fix selected!")
	case 3:
		fmt.Println("refactor selected!")
	case 4:
		fmt.Println("docs selected!")
	case 5:
		fmt.Println("test selected!")
	case 6:
		fmt.Println("chore selected!")
	default:
		fmt.Println("Choose a number between 1-6!")
		fmt.Println("Closing...")
		return

	}

	_,err2 := fmt.Scan(&msg)
	if err2!= nil {
		fmt.Println("error occured", err2)
	}

	fmt.Println("Here is your git message")

	fmt.Println(msgtype,":",msg)
}

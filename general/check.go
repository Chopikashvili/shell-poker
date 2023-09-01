package general

import "fmt"

func Check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic("Program stopped. Try running again.")
	}
}

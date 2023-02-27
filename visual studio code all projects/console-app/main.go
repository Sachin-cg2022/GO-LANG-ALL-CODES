package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard" // we have obtained from 3rd party
)

func main() {
	// this will only execute when we can't open keyboard package
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	//whatever follows the defer keyword won't execute immediately.
	//Instead, it will execute as soon as the current function,
	defer func() {
		_ = keyboard.Close()

	}() //anonymous function

	//instruction
	fmt.Println("press any key on the keyboard. press ESC to quit")
	for {

		// instead of using bufio as eliza
		// i am going to call a function which gives 3 return arguments
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("you pressed ", char, key)
		} else {
			fmt.Println("you pressed ", char)
		}

		if key == keyboard.KeyEsc { //constant
			break
		}

	}
	fmt.Println("program exiting ...")
}

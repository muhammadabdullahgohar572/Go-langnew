package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {

	// fmt.Printf("HI my name is Abdullah .How Are  you ")
	// fmt.Printf("\n")
	// rander :=bufio.NewReader(os.Stdin);
    // fmt.Printf("Enter your age ");
       
	// age,_:=rander.ReadString('\n');
	// fmt.Printf("Your age is and %T %v",age)

        fmt.Print("HI my name is Abdullah.Hoe are you");
		fmt.Print("\n");
		render :=bufio.NewReader(os.Stdin);
		fmt.Printf("Enter your age ");
		name,_:=render.ReadString('\n')
		fmt.Printf("Your name is %s ,%T",name,name)
}
package main 

import (

	"bufio"
	"fmt"
	"os"

)


func main (){

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to this to dolist!!!!!")
	fmt.Println("test but we will create a real to do list in the future:")

     scanner.Scan()
	 task := scanner.Text()

	 fmt.Println("task of the day is:" , task)


}
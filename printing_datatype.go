package main
import "fmt"
func main(){

	//Printing a string
	var my_string="Welcome Aashish"
	fmt.Printf("The String is %s\n",my_string)

	//Printing a decimal number
	var my_number int=22
	fmt.Printf("My age is %d\n",my_number)

	//Printing string & number 
	const name, age = "Aashish", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	//Printing a floating-point number
	var my_data float32 = 3.142
  	fmt.Printf("The floating number is %g\n", my_data)

  	//Printing a binary representation
	var my_num int=8055
  	fmt.Printf("My number is %b",my_num)

}
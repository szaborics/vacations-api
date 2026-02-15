package controller

import (
	"fmt"
	"testing"
)


var controllerService = NewVacationController()


func TestGetFiltered (t *testing.T){
	
	t.Run("Test getting vacationg with no filter", func(t *testing.T) {

		myNumber:=100
		addressMyNumber := &myNumber
		

		fmt.Printf("%d is my number", myNumber)
		fmt.Printf("%d is my number's pointer to the memory address of myNumber &myNumber", *addressMyNumber)
		fmt.Printf("%d is the memory address of myNumber", &myNumber)


		
		

	})



}

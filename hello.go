package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main(){
	file,err := os.Open("input") 


	if err != nil {
		fmt.Println(err)
	}

	 elves := make([] int, 0)	

	scanner := bufio.NewScanner(file)

	elf := 0
	index := 0
	elves = append(elves, elf)
	for scanner.Scan(){
		
		line := scanner.Text()
		calorieCount, err := strconv.Atoi(line)

		if err != nil {
			elves[index] += elf
			elves = append(elves, elf)
			elf = 0	
			index++
		}

		elf += calorieCount

	}
}		


func bubbleSort(unsortedList []int) []int {
	sortedList :=  make([] int, 0)
	for elf := range unsortedList{
		if unsortedList[elf] > unsortedList[elf + 1]{
			tempContainer := 	
		}
	}		
	return sortedList
}

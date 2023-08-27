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
	elves = append(elves, elf)
	for scanner.Scan(){
		line := scanner.Text()
		calorieCount, err := strconv.Atoi(line)
		if calorieCount > 0{
			elf += calorieCount		
		}
		if err != nil {
			elves = append(elves, elf)
			elf = 0
		}
	}


	sortedElves := bubbleSort(elves)
	
	totalElves := len(sortedElves) 	
	bigThree := sortedElves[totalElves - 1] + sortedElves[totalElves - 2]	 + sortedElves[totalElves - 3]
	fmt.Println(sortedElves[totalElves -1])
	fmt.Println(bigThree)
}

func bubbleSort(unsortedList []int) []int {
		
	for index := len(unsortedList); index > 0; index--{
		for j:= 0; j < index; j++ {
			nextElf := j + 1

			if j == index - 1{
				
				break;
			}

			if unsortedList[j] > unsortedList[nextElf]{
				tempStorage := unsortedList[nextElf]
				unsortedList[nextElf] = unsortedList[j]
				unsortedList[j] = tempStorage

			}
		}
	}
	
	return unsortedList
}

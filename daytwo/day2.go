package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	input,err := os.Open("./input")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(input)
	partTwo(scanner)


}

func partOne(scanner *bufio.Scanner){


	totalScore := 0
	for scanner.Scan(){
		line := scanner.Text()
		round := strings.Split(line, " ")

		//this assumes the input is always in indentical format, being 3 charaters long, with the second always being a single space. eg "A Y", "B X", "C Z" 
		// Part 1  
		//	Your end score is your scores from all rounds totaled.
		//	Scores per round are a combination of whether you win 6/draw 3/lose 0 and your move
		//	X = Rock 1 Point
		//	Y = Paper 2 Points
		//	Z = Scissors 3 Points
		//	

		//	Opponent Input
		//	A = Rock
		//	B = Paper
		//	C = Scissors
		rock := 1
		paper := 2
		scissors := 3

		win := 6
		draw := 3
		lose := 0

		roundScore := 0
		theirInput := round[0]
		myInput := round[1]
		switch myInput {

		case "X":
			roundScore += rock
			switch theirInput {
			case "A":
				roundScore += draw	
			case "B":
				roundScore += lose	
			case "C":
				roundScore += win
			}
		case "Y":
			roundScore += paper 
			switch theirInput {
			case "A":
				roundScore += win
			case "B":
				roundScore += draw
			case "C":
				roundScore += lose
			}
		case "Z":
			roundScore += scissors
			switch theirInput {
			case "A":
				roundScore += lose
			case "B":
				roundScore += win
			case "C":
				roundScore += draw
			}
		}
	
		totalScore += roundScore
	}
	fmt.Println(totalScore)
	
}

func partTwo(scanner *bufio.Scanner){

	totalScore := 0

	for scanner.Scan() {

		line :=	scanner.Text()	
		round := strings.Split(line, " ")

		// Same round calculations as above, except the second input is if you should win lose or draw
		//this assumes the input is always in indentical format, being 3 charaters long, with the second always being a single space. eg "A Y", "B X", "C Z" 
		// Part 2  
		//	Your end score is your scores from all rounds totaled.
		//	Scores per round are a combination of whether you win 6/draw 3/lose 0 and your move
		//	X = Lose 0 Point
		//	Y = Draw 3 Points
		//	Z = Win 6 Points
		//	

		//	Opponent Input
		//	A = Rock
		//	B = Paper
		//	C = Scissors

		//	Playing Rock Gives 1 Point
		//	Playing Paper Gives 2 Points
		//	Playing Scissors Gives 3 Points

		rock := 1
		paper := 2
		scissors := 3

		win := 6
		draw := 3
		lose := 0

		roundScore := 0
		theirInput := round[0]
		roundResult := round[1]

		switch roundResult {

		case "X":
			roundScore += lose
			switch theirInput {
			case "A":
				roundScore += scissors
			case "B":
				roundScore += rock
			case "C":
				roundScore += paper
			}

		case "Y":
			roundScore += draw
			switch theirInput {
			case "A":
				roundScore += rock 
			case "B":
				roundScore += paper
			case "C":
				roundScore += scissors
			}
		case "Z":
			roundScore += win
			switch theirInput {
			case "A":
				roundScore += paper 
			case "B":
				roundScore += scissors 
			case "C":
				roundScore += rock 
			}
		}	
		totalScore += roundScore	
	}
	
	fmt.Println(totalScore)
}

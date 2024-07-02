package main

import (
	"fmt"
)

/*

	create a project that manage a football team. The program should allow the user to:
	1. add a new player - each player has a name, short number and number of goals
	2. Remove - remove a player by name
	3. list all players
	4. show the name of the player with the highest number of goals.
	5. Exit

*/

func showMenu (){	
	fmt.Println("1 - To add a new Player")
	fmt.Println("2 - To remove a player")
	fmt.Println("3 - To list all Players")
	fmt.Println("4 - To list the player with the highest score")
	fmt.Println("5 - To exit")

	fmt.Println("Enter your choice:")
	

}	

func highestScore (goalsNumber[]int, newPlayer[]string){

	max := goalsNumber[0]
	maxIndex := 0

	for i := 0; i < len(goalsNumber); i ++ {
		if goalsNumber[i] > max {
			max = goalsNumber[i]
			maxIndex = i
		}
	}
	fmt.Println("The player with the highest score is: ", newPlayer[maxIndex])


}


func listAllPlayers(pNewPlayer*[]string){

	fmt.Println("This is the list of all the players:")

	if len(*pNewPlayer) == 0 {
		fmt.Println("No players available.")
		return
	}

	for i := 0; i < len(*pNewPlayer); i++{
		fmt.Printf("%d, %s\n", i+1, *pNewPlayer)
	}

}

func removePlayer (pNewPlayer*[]string, pShirtNumber*[]int, pGoals*[]int){

	var remove string	
	fmt.Println("Enter the user you want to remove: ")
	fmt.Scan(&remove)

	for i := 0; i < len(*pNewPlayer); i++ {
		if (*pNewPlayer)[i] == remove {
			*pNewPlayer = append((*pNewPlayer)[:1], (*pNewPlayer)[i+1:]...)
			*pShirtNumber = append((*pShirtNumber)[:1], (*pShirtNumber)[i+1:]...)
			*pGoals = append((*pGoals)[:1], (*pGoals)[i+1:]...)
			
			fmt.Printf("%v has successfully been removed.", remove)

		}
	}
}

func addPlayer(pNewPlayer*[]string ,pShirtNumber*[]int,pGoals *[]int) {

	var newPlayer string
	fmt.Println("Please Enter the Players Name: ")
	fmt.Scan(&newPlayer)

	var shirtNumber int
	println("Please enter their shirt number: ")
	fmt.Scan(&shirtNumber)

	var goalsNumber int
	fmt.Println("Enter the number of goals they scored")
	fmt.Scan(&goalsNumber)


	*pNewPlayer = append(*pNewPlayer, newPlayer)
	*pShirtNumber = append(*pShirtNumber, shirtNumber)
	*pGoals = append(*pGoals, goalsNumber)
}

func main () {

	var newPlayer = [] string {}
	var shirtNumber = [] int {}
	var goalsNumber  = [] int {}

	for {
		showMenu()
		var choice int
		fmt.Scan(&choice)

		if choice == 1{
			addPlayer(&newPlayer, &shirtNumber, &goalsNumber)

		} else if choice == 2{

			removePlayer(&newPlayer, &shirtNumber, &goalsNumber)

		} else if choice == 3 {

			listAllPlayers(&newPlayer)

		} else if choice == 4 {

			highestScore(goalsNumber,newPlayer)

		} else if choice == 5 {

			fmt.Println("Exiting...")
			break
		} else {

			if choice < 1 || choice > 5 {
				fmt.Println("Invalid number, please enter a number between 1 - 5.")
			}
		}
	}






}




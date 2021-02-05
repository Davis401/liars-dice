// You can edit this code!
// Click here and start typing.
package main

 	

import (
    "fmt"
    "math/rand"
    
)

type Player struct { 
	DiceNum	int 
    Dice   [5]int
	Name	String
}

func roll(player *Player) Player{
	i := 0
	fmt.Println(player.Name," rolls: ")
	for i < player.DiceNum {
		var r int = rand.Intn(6);
		fmt.Print(r, ",")
		player.Dice[i] = r;
		i++
	}
	return player;
}


func main() {
	players := initPlayers(5);

	
}


func initPlayers(*numOfPlayers int) []Players{
	var players [numOfPlayers]Players;

	i := 0;
	j := 0;
	for i < numOfPlayers {
		var player Player;
		player.DiceNum = 5;
		player.Name = "Player " + (i + 1)
		player = roll(player);
		i++
	}
	return players;
}
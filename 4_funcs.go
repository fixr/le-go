package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Sin return
func shout(msg string) {
	fmt.Println(msg)
}

// Con return
func kill() string {
	players := []string{"@humanrender", "@cicloon", "@AlvarOS_anchez", "@eltercero", "@dadelmo", "@mefastidias", "@weapp", "@zaryel", "@oskarcalvo", "@pantulis", "@littlemove", "@fixr"}

	deadPlayer := players[rand.Intn(len(players))]

	return deadPlayer
}

// Con multiple named return
func shoot(bullets int) (dead int, alive int) {

	if bullets == 0 { bullets = 1 }

	for i := 0; i < bullets; i++ {
		fmt.Print("BAM! ")
	}
	fmt.Println()

	dead = rand.Intn(bullets)
	alive = bullets - dead

	return
}

func main() {

	rand.Seed(time.Now().Unix())

	// Closures
	// timer := func(secs int) {
	// 	fmt.Print("Shitstorm coming in... ")
	// 	for t := secs; t >= 1; t-- {
	// 		fmt.Printf("%d ", t)
	// 		time.Sleep(1 * time.Second)
	// 	}
	// 	fmt.Println()
	// }
	// timer(3)

	shout("FIRE IN THE HOLE!")

	// dead, alive := shoot(rand.Intn(10))

	// if dead > 0 {
	// 	player := kill()
	// 	fmt.Printf("%v was shot dead :(", player)
	// }

	// fmt.Println()
	// fmt.Printf("%d players were killed. %d survived", dead, alive)
	// fmt.Println()

}

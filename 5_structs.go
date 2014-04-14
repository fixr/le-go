package main

import (
  "fmt"
  "math/rand"
  "time"
)

type Zombie struct {
  Name string
  agility int // private
  hunger int  // private
}

// (z Zombie) sería como "self."
func (z Zombie) AteYou() bool {
  if z.agility > 4 && z.hunger > 3 {
    return true
  } else {
    return false
  }
}

// private method
func (z Zombie) toString() string {
  return fmt.Sprintf("%v (AGI %v, HGR %v)", z.Name, z.agility, z.hunger)
}

func main(){
  rand.Seed(time.Now().UTC().UnixNano())

  names := []string{"Arrghhho", "Oghozz", "Ghiaaazz"}
  
  z := Zombie{names[rand.Intn(len(names))], rand.Intn(10), rand.Intn(10)}

  if z.AteYou() {
    fmt.Printf("Mr. %v ate your brain.", z.toString())
    fmt.Println()
  } else {
    fmt.Println("You need brains to be eaten.")
  }
}
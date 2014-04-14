package main

import "fmt"

func main() {


  // FLOW CONTROL

  // matrix = true

  // if matrix
  //   puts "Took the red pill"
  // else
  //   puts "Blue pill it is. Boo!"
  // end

  fmt.Println("-----")
  matrix := true

  if matrix {
    fmt.Println("Took the red pill")
  } else {
    fmt.Println("Blue pill it is. Boo!")
  }


  // ITERADORES
  // all your FOR are belong to us

  // FOR normal

  // 5.times.do |n|
  //  puts "A la cuenta de... #{ n }"
  // end

  fmt.Println("-----")
  for i := 0; i < 5; i++ {
    fmt.Printf("A la cuenta de... %d\n", i)
  }

  // WHILE

  // mola = true
  // j = 0

  // while mola do
  //   puts "Mola!"
  //   j += 1
  //   mola = false if j > 2
  // end

  fmt.Println("-----")

  mola := true
  j := 0

  for mola {
    fmt.Println("Mola!")
    if j++; j > 2 {
      mola = false
    }
  }

  // EACH (_with_index)

  // puts "MISSING:"
  // found_people = ["Hurley", "Jack", "Kate", "Ben", "Sayid"]
  // found_people.each_with_index do |peep, index|
  //   puts "#{ index } #{ peep }"
  // end

  fmt.Println("-----")
  fmt.Println("MISSING:")

  foundPeople := []string{ "Hurley", "Jack", "Kate", "Ben", "Sayid" }
  for index, peep := range foundPeople {
    fmt.Println((index + 1), peep)
  }


  // SWITCH

  fmt.Println("-----")
  const answerToLife = 1
  switch answerToLife {
    case 1,2,3,4,5:
      fmt.Println("Nope.")
    case 42:
      fmt.Println("La respuesta a la vida es 42. ¿Cuál es la pregunta?")
  }

}
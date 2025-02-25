# Loops in Go
The basic loop in Go is written in standard C-like syntax:

for INITIAL; CONDITION; AFTER{
  // do something
}

For example:

for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9


# Omitting Conditions from a for Loop in Go
Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.

for INITIAL; ; AFTER {
  // do something forever
}

# There Is No While Loop in Go
Most programming languages have a concept of a while loop. Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.

for CONDITION {
  // do some stuff while CONDITION is true
}

For example:

plantHeight := 1
for plantHeight < 5 {
  fmt.Println("still growing! current height:", plantHeight)
  plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")

Continue & Break
Whenever we want to change the control flow of a loop we can use the continue and break keywords.

The continue keyword stops the current iteration of a loop and continues to the next iteration. continue is a powerful way to use the guard clause pattern within loops

The break keyword stops the current iteration of a loop and exits the loop.
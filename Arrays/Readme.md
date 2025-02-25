# Arrays in Go
Arrays are fixed-size groups of variables of the same type. For example, [4]string is an array of 4 values of type string.

To declare an array of 10 integers:

var myInts [10]int

or to declare an initialized literal:

primes := [6]int{2, 3, 5, 7, 11, 13}

# Slices in Go
99 times out of 100 you will use a slice instead of an array when working with ordered lists.

Arrays are fixed in size. Once you make an array like [10]int you can’t add an 11th element.

A slice is a dynamically-sized, flexible view of the elements of an array.

The zero value of slice is nil.

Non-nil slices always have an underlying array, though it isn’t always specified explicitly. To explicitly create a slice on top of an array we can do:

primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}

The syntax is:

arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]

Where lowIndex is inclusive and highIndex is exclusive.

# Make
Most of the time we don’t need to think about the underlying array of a slice. We can create a new slice using the make function:

// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)

Slices created with make will be filled with the zero value of the type.

If we want to create a slice with a specific set of values, we can use a slice literal:

mySlice := []string{"I", "love", "go"}

Capacity
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. It is accessed using the built-in cap() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3

Generally speaking, unless you’re hyper-optimizing the memory usage of your program, you don’t need to worry about the capacity of a slice because it will automatically grow as needed.

# Variadic
Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the “…” syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}

The familiar fmt.Println() and fmt.Sprintf() are variadic! fmt.Println() prints each element with space delimiters and a newline at the end.

func Println(a ...interface{}) (n int, err error)

# Spread Operator
The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}

# Append
The built-in append function is used to dynamically add elements to a slice:

func append(slice []Type, elems ...Type) []Type

f the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.

Notice that append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)


# Range
Go provides syntactic sugar to iterate easily over elements of a slice:

for INDEX, ELEMENT := range SLICE {
}

Note: the element is a copy of the value at that index.

For example:

fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape

# Slice of Slices
Slices can hold other slices, effectively creating a matrix, or a 2D slice.

rows := [][]int{}

# Tricky Slices
The append() function changes the underlying array of its parameter AND returns a new slice. This means that using append() on anything other than itself is usually a BAD idea.

// dont do this!
someSlice = append(otherSlice, element)

# Currying
Function currying is a concept from functional programming and involves partial application of functions. It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.

Let's simulate this behavior. For example:

func main() {
  squareFunc := selfMath(multiply)
  doubleFunc := selfMath(add)

  fmt.Println(squareFunc(5))
  // prints 25

  fmt.Println(doubleFunc(5))
  // prints 10
}

func multiply(x, y int) int {
  return x * y
}

func add(x, y int) int {
  return x + y
}

func selfMath(mathFunc func(int, int) int) func (int) int {
  return func(x int) int {
    return mathFunc(x, x)
  }
}
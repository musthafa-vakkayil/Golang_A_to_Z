# The Error Interface
Go programs express errors with error values. An Error is any type that implements the simple built-in error interface:

type error interface {
    Error() string
}

When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.

Because errors are just interfaces, you can build your own custom types that implement the error interface. Here’s an example of a userError struct that implements the error interface:

type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}

It can then be used as an error:

func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}

# The Errors Package
The Go standard library provides an “errors” package that makes it easy to deal with errors.

Read the godoc for the errors.New() function, but here’s a simple example:

var err error = errors.New("something went wrong")

# Panic
However, there is another way to deal with errors in Go: the panic function. When a function calls panic, the program crashes and prints a stack trace.

As a general rule, do not use panic!

The panic function yeets control out of the current function and up the call stack until it reaches a function that defers a recover. If no function calls recover, the goroutine (often the entire program) crashes.

func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}

Sometimes new Go developers look at panic/recover, and think, “This is like try/catch! I like this!” Don’t be like them.

I use error values for all “normal” error handling, and if I have a truly unrecoverable error, I use log.Fatal to print a message and exit the program.
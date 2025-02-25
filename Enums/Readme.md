# Lack of Enums
My least favorite part of Go? Glad you asked. It's Go's lack of enums, sum types, tagged unions, etc. Compared to other statically typed languages like:

Rust
TypeScript
OCaml
Go's type system just isn't as powerful. It's more similar to C's type system than it is to Rust's. It's more concerned with simplicity than it is with expressiveness.

# Error Handling
In Rust, like Go, errors are just values. In Go, we write something like this:

user, err := getUser()
if err != nil {
    return fmt.Errorf("failed to get user: %w", err)
}
// do something with user

In Rust, we can do something like this:

let user_result = get_user();
let user = match user_result {
    Ok(user) => user,
    Err(error) => return Err(format!("failed to get user: {}", error)),
};

In Rust, the get_user function returns a Result type: a type that is either an Ok or an Err. The compiler forces the developer to handle the error case before they can continue with the happy path (using the user data).

In Go, the developer can choose to happily ignore the error value if they choose and use the user data, even if it's invalid (probably nil or an empty struct).

The support for enums in Rust makes it easier to write bug-free code.

# Type Definitions
For all its faults, TypeScript does have a pretty incredible type system. Here's one of the things it can do that I often miss in Go:

type sendingChannel = "email" | "sms" | "phone";

function sendNotification(ch: sendingChannel, message: string) {
    // send the message
}

This sendingChannel type that we've created is a union type. It can only be one of the three strings that we've defined. That means when a developer calls sendNotification() they can't accidentally pass an invalid sendingChannel like "slack" or even a misspelled "emil". The TypeScript compiler will catch that mistake at compile time.

In Go, we don't have these nice things. We embrace the Grug mentality. The closest thing we have to a union type is a type definition:

type sendingChannel string

const (
    Email sendingChannel = "email"
    SMS   sendingChannel = "sms"
    Phone sendingChannel = "phone"
)

func sendNotification(ch sendingChannel, message string) {
    // send the message
}

It's a bit safer than using a plain old string in Go, but it's not completely safe. Go will stop us from doing this:

sendingCh := "slack"
sendNotification(sendingCh, "hello") // string is not sendingChannel

But it will not stop us from doing this:

// "slack" is automatically implied as a sendingChannel
sendNotification("slack", "hello")

And will also not stop us from doing this:

sendingCh := "slack"
convertedSendingCh := sendingChannel(sendingCh)
sendNotification(convertedSendingCh, "hello")

The sendingChannel type is just a wrapper for string, and because we made some constants of that type, most developers will just use those constants: we've made it easy to do the right thing. That said, Go still doesn't force us to do the right thing like TypeScript does.

# Iota
Go has a language feature, that when used with a type definition (and if you squint really hard), kinda looks like an enum (but it's not). It's called iota.

type sendingChannel int

const (
    Email sendingChannel = iota
    SMS
    Phone
)

The iota keyword is a special keyword in Go that creates a sequence of numbers. It starts at 0 and increments by 1 for each constant in the const block. So in the example above, Email is 0, SMS is 1, and Phone is 2.

Go developers sometimes use iota to create a sequence of constants to represent a set of related values, much like you would with an enum in other languages. But remember, it's not an enum. It's just a sequence of numbers.
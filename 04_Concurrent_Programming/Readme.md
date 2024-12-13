# Concurrency
Concurrency in Go is a fundamental aspect of the language's design, empowering developers to write efficient, scalable, and responsive software. Go's concurrency model revolves around goroutines and channels, providing a powerful yet simple way to handle concurrent tasks.

# Goroutines
Goroutines are lightweight threads managed by the Go runtime. They enable concurrent execution of functions, allowing multiple tasks to run simultaneously within a single Go program. Goroutines are incredibly cheap in terms of memory footprint and overhead, making it practical to spawn thousands of them within a single application.

To create a goroutine, you simply prefix a function call with the go keyword. For example:

func main() {
   // Start a new goroutine
   go sayHello()

   // Continue with main execution
   fmt.Println("Main function")

   // Sleep for a while to allow the goroutine to finish
   time.Sleep(3 * time.Second)
}

func sayHello() {
   fmt.Println("Hello from goroutine!")
}

Because goroutines operate concurrently in distinct paths, it's essential to acknowledge that if the main thread ends, all goroutines also terminate. To prevent premature termination of the main thread, a one-second delay is incorporated using time.Sleep. This pause enables the goroutine to finish its execution smoothly.


There are couple of shortcomings with adding sleep

1. When the CPU is busy, your program's execution might exceed the 3-second threshold, possibly causing goroutines to terminate prematurely within that 3 seconds time frame. This unpredictability could lead to inaccurate totals.

2. Even if your program finishes in less than 3 seconds, you still have to wait for 3 seconds for the output to generate.

In the next step, you will address all of the aforementioned issues and optimize the code.

# Wait Group
Wait groups are a synchronization mechanism provided by the sync package in Go. They allow you to wait for a collection of goroutines to finish their execution before proceeding further in the program. Wait groups are particularly useful when you have a dynamic number of goroutines and need to ensure they all complete their tasks before continuing.

# How Wait Groups Work
Wait groups are represented by the sync.WaitGroup type. You create a new wait group using var wg sync.WaitGroup, and then add the number of goroutines you want to wait for using the Add() method.

Each goroutine increments the wait group counter using wg.Add(1) before starting its task. When a goroutine completes its task, it decrements the counter using wg.Done().

In the below example, two goroutines are added to the wait group using wg.Add(2) before their execution begins.

Finally, the main goroutine, or any other goroutine waiting for the completion of the tasks, calls wg.Wait() to block until all goroutines have finished their tasks.

func main() {
   var wg sync.WaitGroup

   // Add two goroutine to the wait group
   wg.Add(2)

   // Start goroutines
   go doSomeWork(&wg)
	 go doSomeWork(&wg)

   // Wait for all goroutines to finish
   wg.Wait()
   fmt.Println("All goroutines have finished")
}

func doSomeWork(wg *sync.WaitGroup) {
   // Signal that the goroutine has finished
   defer wg.Done() 
   // Simulate some work
   fmt.Println("Goroutine: Working...")
   time.Sleep(time.Second)

   fmt.Println("Goroutine: Finished!")
}

In this example, defer wg.Done() is used inside the doSomeWork function. The defer ensures that no matter how the function exits, whether it returns normally or panics, it will always call wg.Done() before exiting. This helps ensure that the WaitGroup counter is decremented appropriately, allowing wg.Wait() in the main function to block until all workers are finished.

Since doSomeWork is invoked twice, it will decrement the counter twice using wg.Done() . Once the WaitGroup counter is decremented to 0, it will unblock the wg.Wait() call.

This pattern ensures that you properly wait for all goroutines to finish before proceeding, even if an error occurs within a goroutine.

It's possible for two goroutines to simultaneously update the totalSales variable, leading to inconsistent totals. To avoid this, Go offers synchronization primitive called mutex, which is used to control access to shared resources. It ensures that only one goroutine can access the shared resource at any given time, thus ensuring consistency.

# Mutex
A mutex, short for mutual exclusion, is a synchronization primitive used to control access to shared resources in concurrent programs. It ensures that only one goroutine can access a shared resource at a time, preventing data races and ensuring consistency.

In Go, a mutex is represented by the sync.Mutex type from the sync package. It provides two main methods:

1. Lock() : Acquires the mutex. If the mutex is already locked by another goroutine, Lock() will block until it becomes available.

2. Unlock() : Releases the mutex, allowing other goroutines to acquire it.

To use a mutex in your Go code, follow these steps:

Declare a variable of type sync.Mutex.
Use Lock() to acquire the mutex before accessing the shared resource, and Unlock() to release it afterwards.

var (
    counter int
    mutex   sync.Mutex
)

func increment() {
    mutex.Lock()    // Acquire the mutex before modifying the counter
    counter++       // Increment the counter
    mutex.Unlock()  // Unlock the mutex after modifying the counter
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            increment()
        }()
    }
    wg.Wait()

    fmt.Println("Counter:", counter)
}


Another powerful mechanism in Go for managing concurrency is channels.

There are certain scenarios where channels might be preferred over wait groups

Here are some common issues with wait groups that might lead to the use of channels instead:

1. Complexity:

Managing concurrency with wait groups can lead to intricate code, especially when dealing with nested or dynamic concurrency patterns

2. Error Handling

Wait groups lacks built-in error handling mechanisms, necessitating manual propagation of errors through other means, potentially leading to code complexity and increased error risk.

3. Unordered Completion

Wait groups doesn't ensure the ordered completion of goroutines. Processing results in a specific order requires additional synchronization mechanisms, adding complexity to the code.

4. Difficulty in Dynamic Scaling

Dynamically managing the number of goroutines with wait groups poses challenges, as careful addition and subtraction from the wait groups are required, risking race conditions if not executed correctly.

# Channels

Channels offer a more flexible and expressive way to manage concurrency:

1. Synchronization

Channels offer built-in synchronization, ensuring safe data transmission between goroutines without extra synchronization primitives.

2. Error Handling

Channels facilitate error propagation alongside data, simplifying error management and reducing error susceptibility.

3. Ordering

Channels enforce operation sequence, guaranteeing correct result processing without added synchronization overhead.

4. Dynamic Scaling

Channels excel in scenarios with unknown or dynamic goroutine counts. They enable on-demand goroutine creation and coordination via channels.

Channels serve as the primary means of communication and synchronization between goroutines in Go, providing a safe conduit for sending and receiving data. They enable concurrent processes to exchange values without the need for explicit locking or coordination, ensuring that operations proceed smoothly without interference or race conditions.

# Creating Channels
Channels are created using the make function with the chan keyword followed by the type of data that the channel will transmit. Here's how you create a channel:

// Creates an unbuffered channel of type int
ch := make(chan int) 

# Channel Operations
Channels support two main operations: sending and receiving values. These operations are performed using the <- operator:

ch <- value // Send value into the channel
value := <-ch // Receive value from the channel

# Buffered Channels
By default, channels are unbuffered, meaning they only accept a value if there's a corresponding receiver ready to receive it. Buffered channels, on the other hand, have a fixed capacity and can store a certain number of values without a corresponding receiver. Here's how you create a buffered channel:

ch := make(chan int, bufferSize) // Creates a buffered channel with capacity bufferSize

# Closing Channels
Channels can be closed to indicate that no more values will be sent. Receivers can use the second return value from a receive operation to determine if the channel has been closed:

close(ch) // Close the channel

# Channel Direction
You can specify the direction of a channel in its type signature to restrict its usage to sending or receiving operations. This helps enforce communication protocols and prevent misuse of channels. Here's how you specify channel direction:

func sendData(ch chan<- int) {
  // Send data into channel
}

func receiveData(ch <-chan int) {
  // Receive data from channel
}

Here's a simple example demonstrating how to use channels for communication between two goroutines. In this case, the sendData function sends integers to the channel, and the main function receives and prints those integers.

func sendData(ch chan<- int) {
 // Send data into the channel
 ch <- 10
 ch <- 20
 ch <- 30
 close(ch) // Close the channel after sending all values
}

func main() {
 // Create an unbuffered channel of type int
 ch := make(chan int)

 // Start a goroutine to send data into the channel
 go sendData(ch)

 // Receive data from the channel
 for {
   // Attempt to receive a value from the channel
   value, ok := <-ch
   if !ok {
   // Channel closed, exit the loop
   break
  }
  // Print the received value
  fmt.Println("Received:", value)
  }
}

In Go, channels can be used not only for communication between goroutines, but also for propagating errors. By sending error values through channels, goroutines can report errors to other parts of the program or to the main goroutine responsible for error handling:

// Create a channel for error communication
 errCh := make(chan error)

Example of error propagation using channels:

func doTask(resultCh chan int, errCh chan error) {
 resultCh <- 42              // Simulate a calculation
 errCh <- errors.New("something went wrong") // Simulate an error
}

func main() {
 resultCh := make(chan int)  // Channel for result
 errCh := make(chan error)   // Channel for error

go doTask(resultCh, errCh)

 select {
  case result := <-resultCh:
   fmt.Println("Result:", result)
  case err := <-errCh:
   fmt.Println("Error:", err)
  }
}



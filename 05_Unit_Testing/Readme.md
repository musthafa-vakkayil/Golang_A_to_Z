# Unit Testing in Go

As a Go developer, you are fortunate enough to work with a language that has good developer tooling. Go comes with many tooling out of the box, including built-in support for unit testing.

Any file that has a name ending in _test.go is considered a test file by the Go tooling. Within these files, any function that starts with Test is considered a test function. Test functions take the shape func TestXXX(t *testing.T) {}. They don't return anything.

go test ./*

go test -v ./*

Method	Description
t.Fail()	Marks the test as 'failed', but continues with the rest of the test function.
t.FailNow()	Marks the test as 'failed' and stops execution of the current test function.
t.Error(args ...any)	Logs the 'args' parameters to the console and then calls .Fail()
t.Fatal(args ...any)	Logs the 'args' parameters to the console and then calls .FailNow()

# Approaches

1. Arrange/Act/Assert
    * Arrange: set up the right preconditions for your test case
    * Act: execute the code you want to test
    * Assert: check the results and validate your assumptions about the outcome
2. Given/When/Then
    * Given <these parameters>, When <I execute this function>, Then <I expect the following outcome>


if you add a function func TestMain(m *testing.M), then only that function is executed, and no others. You would use it like this:

func TestMain(m *testing.M) {
    log.Println("Preparation")
    exitVal := m.Run()
    log.Println("Cleanup")

    os.Exit(exitVal) // Required to indicate whether the test run passed or failed
}

If you could group the tests according to the function they're testing, the readability of the test file and the test output would improve a lot. Fortunately, Go makes it easy to do so, by letting you define subtests. You can define subtests by inserting calls to t.Run() inside your test functions, so that

func TestMyFunctionWithValidInput(t *testing.T) {
    // arrange
    result := MyFunction("valid input") // act
    // assert
}

becomes

func TestMyFunction(t *testing.T) {
    t.Run("WithValidInput", func (t *testing.T) {
        // arrange
        result := MyFunction("valid input") // act
        // assert
    })
}

# Skipping tests
There can be situations in which you do not want to run all tests.

Go provides a few ways to control exactly which tests are executed. We will look at a few of them.

The simplest way to skip execution of a test is by using one of the following methods:

Method	Description
t.SkipNow()	Marks the test as 'skipped' and stops execution of the current test function.
t.Skip(args ...any)	Logs the 'args' parameters to the console and then calls .SkipNow()
t.Skipf(format string, args ...any)	Logs the parameters to the console (similar to how fmt.Printf() works) and then calls .SkipNow()

These methods can be called at any point during execution of your test, but if the test was already marked as failed earlier in the test, it will still be marked as failed in the output.

Skipping tests using the methods listed above can be useful, but the downside is that you actually have to adjust the test functions (by introducing these method calls). If you simply want to reduce the set of executed tests during development, you will have to resort to the command line. There are several options that you can pass to the go test command to select specific tests.

To zoom in on a particular package, you can pass the package path instead of ./..., e.g.:

go test getting-started-with-go-unit-testing/step5

To run specific tests matching a specific name, you can pass -run `MyFunction`. This will run all tests that contain the text MyFunction in their function name.

# Running tests in parallel

Although unit tests are generally fast, it can still take a while to execute all tests in a large code base. Luckily, Go has two tricks up its sleeve to speed things up:

 * It skip tests for code that has not changed; and
 * Tests in different packages are executed in parallel

That said, Go does not run tests within a package in parallel by default. The reason for this is that it's not uncommon for tests within a package to influence each other. It is up to you as the developer to mark tests within a package as safe for manual execution. To do so, the t *testing.T struct has one more useful method: t.Parallel(). This method must be called in each test that you want to run in parallel, preferably at the start of the test.
Welcome to the Guided: Debugging Go with the Delve CLI Code Lab!

In this lab, you'll learn how to debug Go applications using the Delve command-line debugger.

Delve is a powerful debugger built specifically for Go programs, enabling developers to inspect and control the execution of their code.

To start debugging a Go application with Delve, use the debug subcommand:

dlv debug <package>
<package>: The package or file to debug. If omitted, Delve will debug the package in the current directory.

Once you're in a Delve session (indicated by the (dlv) prompt), you can use a variety of commands to control execution and inspect the program state. Here are some fundamental commands:

help: Lists all available commands and their descriptions
break or b: Sets a breakpoint at a specific function or line
continue or c: Runs the program until it hits the next breakpoint or reaches the end
next or n: Moves execution to the next line in the current function
step or s: Steps into functions to debug them line by line
stepout or so: Continues execution until the current function returns
print or p: Evaluates and displays the value of variables or expressions
args: Lists the current function's arguments
locals: Displays local variables in the current scope
list or l: Displays the source code
exit, quit, or q: Exits the debugging session

# Setting Breakpoints
Breakpoints allow you to pause program execution at specific points and inspect the program state. For example, you can set breakpoints by function name:

break <function_name>
Or by file and line number:

break <file_name>:<line_number>

# Navigating Through Code
Controlling the flow of execution is key to understanding how your program works:

continue: Resumes execution until the next breakpoint or the end of the program
next: Moves to the next line in the current function, stepping over any function calls
step: Steps into the next line of code, including called functions
stepout: Runs the program until the current function returns

# Inspecting Program State
When the program is paused, you can inspect variables and the call stack using these commands:

print <expression>: Evaluates and displays the value of an expression or variable
args: Lists the current function's arguments and their values
locals: Shows local variables and their values in the current scope
list: Displays your current location and the surrounding source code
stack: Displays the current call stack, useful for understanding the sequence of function calls

# Exiting and Restarting the Debugger
When you're done debugging or want to restart the program, you can use the following commands:

restart: Restarts the program from the beginning while keeping breakpoints intact
exit, quit, or q: Exits the Delve debugger and returns you to the shell

# Conditional breakpoints
Conditional breakpoints are a powerful feature that allow you to pause execution only when specific conditions are met, saving time and simplifying the debugging process.

You can set a conditional breakpoint using the condition command, which lets you specify a boolean expression that must evaluate to true for the breakpoint to trigger. Here's how to do it:

1. Set a Breakpoint: First, set a breakpoint at the desired location using the break command. For example:

break main.go:10

2. Add a Condition: After setting the breakpoint, add a condition to it using the condition command. The syntax is:

condition <breakpoint ID> <boolean expression>

For example, if you want the breakpoint to trigger only when the variable x equals 5, you would use:

condition 1 x == 5

Here, 1 is the ID of the previously set breakpoint.

Alternatively, starting from Delve 1.23.0, you can combine these steps into one command by specifying the condition directly when setting the breakpoint. This is the syntax:

break <location> if <condition>

For example:

break main.go:10 if x == 5


# Further Exploration
Delve offers many more commands and features to enhance your debugging experience. Here are some additional commands you might find useful:

clear: Removes a breakpoint
goroutines: Lists all goroutines
threads: Lists all threads
sources: Lists source files
vars: Lists package variables
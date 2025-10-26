1. Why does Go not have exceptions?

We believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code.  
It also tends to encourage programmers to label too many ordinary errors, such as failing to open a file, as exceptional.  

Go takes a different approach. For plain error handling, Go's multi-value returns make it easy to report an error without    
overloading the return value. A canonical error type, coupled with Go's other features, makes error handling pleasant   
but quite different from that in other languages.  

Go also has a couple of built-in functions to signal and recover from truly exceptional conditions.  
The recovery mechanism is executed only as part of a function's state being torn down after an error, which is sufficient   
to handle catastrophe but requires no extra control structures and, when used well, can result in clean error-handling code.  

See the Defer, Panic, and Recover article for details.   
Also, the Errors are values blog post describes one approach to handling errors cleanly in Go by demonstrating that,  
since errors are just values, the full power of Go can be deployed in error handling.    

## Printing and logging
1. Write the code with errors before writing the code without errors. Always check for errors. Always, always, always.  
1.`fmt.Println()`  You can print out there.  
1. `log.Println()` You could write it to a file and we're going to see how to do that.   
1. `log.Fatalln()` Fatalln, mean, you're dead. shutting down program similar to ` os.Exit()`  
1. `log.Panicln() ` We're not dead yet. It's not fatal. It's a panic.  
1. So panic you can use recover.
   deferred functions run
   can use “recover”



   



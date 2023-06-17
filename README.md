# APL-go

## Introduction

### week 1 (simple program)

- [x] rewrite the Python in Golang

### week 2

- [x] Monolith (Four.go)
  - No abstractions
  - No use of lib functions
- [x] Cookbook (Five.go)
  - Larger problem decomposed in procedural abstractions
  - Larger problem solved as a sequence of commands, each corresponding to a procedure
- [x] Pipeline (Six.go)
  - Larger problem decomposed in functional abstractions. Functions, according to Mathematics, are relations from inputs to outputs.
  - Larger problem solved as a pipeline of function applications
- [x] Code Golf (Seven.go)
  - As few lines of code as possible

### week 3

- [x] Infinite Mirror (Eight.go)
  - All, or a significant part, of the problem is modelled by induction. That is, specify the base case (n_0) and then the n+1 rule
- [x] Kick Forward (Nine.go)
  - Each function takes an additional parameter, usually the last, which is another function
  - That function parameter is applied at the end of the current function
  - That function parameter is given as input what would be the output of the current function
  - Larger problem is solved as a pipeline of functions, but where the next function to be applied is given as parameter to the current function
- [x] The one (Ten.go)
  - Existence of an abstraction to which values can be converted.
  - This abstraction provides operations to (1) wrap around values, so that they become the abstraction; (2) bind itself to functions, so to establish sequences of functions; and (3) unwrap the value, so to examine the final result.
  - Larger problem is solved as a pipeline of functions bound together, with unwrapping happening at the end.
  - Particularly for The One style, the bind operation simply calls the given function, giving it the value that it holds, and holds on to the returned value.


## week 4

- [x] Letterbox (Twelve.go)
  - The larger problem is decomposed into 'things' that make sense for the problem domain
  - Each 'thing' is a capsule of data that exposes one single procedure, namely the ability to receive and dispatch messages that are sent to it
  - Message dispatch can result in sending the message to another capsule
- [x] Closed Maps (Thirteen.go)
  - The larger problem is decomposed into 'things' that make sense for the problem domain
  - Each 'thing' is a map from keys to values. Some values are procedures/functions.
- [x] Hollywood (Fifteen.go)
  - Larger problem is decomposed into entities using some form of abstraction (objects, modules or similar)
  - The entities are never called on directly for actions
  - The entities provide interfaces for other entities to be able to register callbacks 
  - At certain points of the computation, the entities call on the other entities that have registered for callbacks

## week 5

- [x] Constructivist (TwentyOne.go)
  - Every single procedure and function checks the sanity of its arguments and either returns something sensible when the arguments are unreasonable or assigns them reasonable values
  - All code blocks check for possible errors and escape the block when things go wrong, setting the state to something reasonable
- [x] Tantrum (TwentyTwo.go)
  - Every single procedure and function checks the sanity of its arguments and refuses to continue when the arguments are unreasonable
  - All code blocks check for all possible errors, possibly print out context-specific messages when errors occur, and pass the errors up the function call chain
- [x] Quarantine (TwentyFive.go)
  - Core program functions have no side effects of any kind, including IO
  - All IO actions must be contained in computation sequences that are clearly separated from the pure functions
  - All sequences that have IO must be called from the main program

## week 6

- [x] Introspective (Seventeen.java)
  - The problem is decomposed using some form of abstraction (procedures, functions, objects, etc.)
  - The abstractions have access to information about themselves, although they cannot modify that information
- [x] Plugins (Twenty.java)
  - The problem is decomposed using some form of abstraction (procedures, functions, objects, etc.)
  - All or some of those abstractions are physically encapsulated into their own, usually pre-compiled, packages. Main program and each of the packages are compiled independently. These packages are loaded dynamically by the main program, usually in the beginning (but not necessarily).
  - Main program uses functions/objects from the dynamically-loaded packages, without knowing which exact implementations will be used. New implementations can be used without having to adapt or recompile the main program.
  - External specification of which packages to load. This can be done by a configuration file, path conventions, user input or other mechanisms for external specification of code to be linked at run time.

## week 7

- [x] Spreadsheet (TwentySeven.go)
  - The problem is modeled like a spreadsheet, with columns of data and formulas
  - Some data depends on other data according to formulas. When data changes, the dependent data also changes automatically.
- [x] Lazy Rivers (TwentyEight.go)
  - Data comes to functions in streams, rather than as a complete whole all at at once
  - Functions are filters / transformers from one kind of data stream to another

## week 8

- [x] Actors (TwentyNine.java)
  - The larger problem is decomposed into 'things' that make sense for the problem domain
  - Each 'thing' has a queue meant for other \textit{things} to place messages in it
  - Each 'thing' is a capsule of data that exposes only its ability to receive messages via the queue
  - Each 'thing' has its own thread of execution independent of the others.
- [x] Dataspaces (Thirty.java)
  - Existence of one or more units that execute concurrently
  - Existence of one or more data spaces where concurrent units store and retrieve data
  - No direct data exchanges between the concurrent units, other than via the data spaces
- [x] Map-reduce (ThirtyTwo.java)
  - Input data is divided in chunks, similar to what an inverse multiplexer does to input signals
  - A map function applies a given worker function to each chunk of data, potentially in parallel
  - The results of the many worker functions are reshuffled in a way that allows for the reduce step to be also parallelized
  - The reshuffled chunks of data are given as input to a second map function that takes a reducible function as input

## week 9

- [x] arrays (one.py)
  - Using Python+numpy, or any other array programming language, implement a program that takes as input an array of characters (the characters from Pride and Prejudice, for example), normalizes to UPPERCASE, ignores words smaller than 2 characters, and replaces characters with their Leet  counterparts when there is a one-to-one mapping. Then it prints out the 5 most frequently occurring 2-grams. (2-grams are all 2 consecutive words in a sequence) Note that you should stick to the array programming style as much as possible. This means: avoid explicit iteration over the elements of the array. If you find yourself writing an iteration, think of how you could do it with one or more array operations. (Sometimes, it's not possible; but most often it is)


## LICENSE
- Code is based on MIT LICENSE
- Other files are based on CC-BY-NC-SA v4.0

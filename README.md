Practical approach with GO

1: Structures

2: Function

3: Methods

4: Interfaces

5: Routines

6: Channels

7: Producer/Consumer pattern

8: Flags

9: Mutex/sync


Structures:

	Used for the grouping of the arbitrary type named values(fields)

	Private/Public members

	Declaration syntaxes

	If a field is omitted in this kind of literal, it is set to the zero value for its type

	For efficiency, larger struct types are usually passed to or returned from functions indirectly	using a pointer

	Structs are comparable if all the fields of the struct are comparable and can be compared using "=" or "!="

	If struct is comparable then can be used as the key of the map

	Structs can be embedded. The


Functions:

	Declarartion

	Pass by value/Pass by refrence

	Anonymous function

	Receiving multiple result types

	Passing a function as argument also known as higher-order functions

	Deferring function calls using defer someting like finally

	Function panic

	Panic recovery


Methods:

	When a function is scoped to a type, or attached to the type, it is known as a method. The parameter attaches the function to the type of that parameter

	The paramenter is also called receiver, which can be the pointer as well(Pass by value / Pass by refrence)

	Nil is a valid receiver value. So it needs to be taken care as per the struct definition

	Encapsulation using methods


Interfaces:

	Implementing interfaces

	Generalization using subtyping with Go

	Multiple interfaces can be implemented by one struct

	Interfaces can be embedded

	Empty interface type : interface{}


Routines:

	Go routines are not thread

	Go routines submit an unit of execution

	Syntax for the execution of go routine

	Will not cover sync/mutex here


Channels:

	Go channels are used for the communication between go routines

    Send / Receive wih channels // <- chan , chan <-

    Channel blocks while receiving and writing. Any subsequent statement becomes unreachable

    Direction specific channels (Can be of send or receive directed)

	Send / receive deadlocks with channels

	Channels can be of two types:
		1: Synchronous	// ch := make(chan int)
		2: Buffered		// ch := make(chan int, 5)



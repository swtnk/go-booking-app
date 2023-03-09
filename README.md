## Booking App: Commandline application
### Learning basics of Go programming language.
---

## Initialize go project:
```sh
$ go mod init <project-name>
```

## Start Coding:
`Create file main.go in root directory of the project.`

* Create main function and greet user with a welcome message.
    ```go
    package main // Need to declare a package, in our case it is main.

    import "fmt" // fmt is a formatting package contains bunch of methods for printing to console.

    func main() { // Main function.

    }
    ```
* Data Containers: Variables and Constants
    ```
    ...
    <container_type> <container_name> <data_type> // It is necessary to specify data type in case of declaration only.
    <container_type> <container_name> = data // It is optional to specify data type in case of declaration and data assignment to it but in some case it is good practice to specify data types.
    <container_name> := data // In this way data types can not be defined in code.
    ...
    ```
    ```go
    package main

    import "fmt"

    func main() {
        var buyerName string
        var conferenceName = "Go Conference"
        const totalTickets uint = 100
        conferenceName  := "Go Conference"
    }
    ```

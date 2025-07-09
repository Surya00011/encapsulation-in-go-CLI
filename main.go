package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "encapsulation/user"
    "encapsulation/userlist"
)

func displayOptions() []string {
    return []string{
        "1. To Add User details",
        "2. To View Users list",
        "3. To delete user by email",
        "4. Exit",
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
  for{
    fmt.Println("===== Menu =====")
    for _, option := range displayOptions() {
        fmt.Println(option)
    }

    fmt.Print("Enter your choice: ")
    scanner.Scan()
    choice := strings.TrimSpace(scanner.Text())

    switch choice {
    case "1":
        fmt.Print("Enter name: ")
        scanner.Scan()
        name := strings.TrimSpace(scanner.Text())

        fmt.Print("Enter email: ")
        scanner.Scan()
        email := strings.TrimSpace(scanner.Text())

        newUser := user.UserConstructor(name, email)
        userlist.AddUser(newUser)

        fmt.Printf("User %s (%s) added successfully\n", name, email)

    case "2":
        fmt.Println("\n==== Users List ====")
        users := userlist.GetAllUsers()

        if len(users) == 0 {
            fmt.Println("No users found.")
            continue
        }

        for index, user:= range users {
            fmt.Printf("%d. Name: %s | Email: %s\n", index+1, user.GetName(), user.GetEmail())
        }
    case "3":
	 fmt.Println("Enter email id to delete user :")
	 scanner.Scan()
	 email:= strings.TrimSpace(scanner.Text())
         status := userlist.DeleteUser(email)
	 fmt.Println(status)
    case "4":
        fmt.Println("Exiting...")
        return

    default:
        fmt.Println("Invalid choice. Exiting.")
	continue
    }
  }
}


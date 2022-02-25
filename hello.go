// Hello Word in Go by Vivek Gite
package main

// Import OS and fmt packages
import (
	"fmt"
	"os"
)

// Let us start
func main() {
    fmt.Println("Hello, world!")  // Print simple text on screen
    fmt.Println(os.Getenv("USER"), ", Let's be friends!\n") // Read Linux $USER environment variable 
    fmt.Println("Hello from Docker on GitHub Container Registry!\nThis message shows that your installation appears to be working correctly.")
}



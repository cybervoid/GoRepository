package main

//A bare minimum go program used in the Docker image.
import (
    "fmt"
    "os"
)

//this program will take a name environment variable and print back the <name> variable value to the user , then it quit

func main() {
    fmt.Println(os.Getenv("NAME") + " came from an environment variable. Docker is exciting! ")
}

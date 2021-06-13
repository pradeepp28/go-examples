package main

import (
	"encoding/base64"
	"io"
	"log"
	"os"
)

const (
	base64encstr = `iVBORw0KGgoAAAANSUhEUgAAAagAAADrCAIAAACza5XhAAAKMWlDQ1BJQ0MgUHJvZmlsZQAASImd
lndUU9kWh8+9N71QkhCKlNBraFICSA29SJEuKjEJEErAkAAiNkRUcERRkaYIMijggKNDkbEiioUB
8b2kqeGaj4aTNftesu5mob4pr07ecMywRwLBvDCJOksqlUyldAZD7g9fxIZRWWPMvXRNJROJRBIG
Y7Vx0mva1HAwYqibdKONXye3dW4iUonhWFJnqK7OaanU1gGkErFYEgaj0cg8wK+zVPh2ziwnHy07
U8lYTNapezSzOuevRwLB7CFkqQQCwaJDiBQIBIJFhwh8AoFg0SHUqQUCASRJKkwkhMy/JfODWPEJ
BIJFhwh8AoFg0TFnQqQ55GtPFopcJsN97e1nYtNuIBYeGBgYCmYrmE3jZ05iaGAoMX0xzxkWz6Hv
yO7WvrlwzA0uLzrD+VkKqViwl9IfTBVNFMyc/x9alloiPPlqhQAAAABJRU5ErkJggg==`
)

func main() {

	f, err := os.Open("s.png")
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	defer f.Close()
	b, _ := io.ReadAll(f)
	log.Println(base64.StdEncoding.EncodeToString(b))
}

package main

import (
    "os"
    "NeoManager/handlers"
)

func main(){
    command := os.Args[1:]
    handlers.CommandHandle(command)
}

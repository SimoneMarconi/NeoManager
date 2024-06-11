package main

import (
    "os"
    "NeoManager/handlers"
)

func main(){
    check := len(os.Args)
    if check == 1 {
        handlers.Man()        
        os.Exit(0)
    }
    command := os.Args[1:]
    handlers.CommandHandle(command)
}

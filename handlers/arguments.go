package handlers

import (
	"fmt"
	"strconv"
)

func CommandHandle(command []string){
    switch command[0] {
    case "install":
        if len(command) < 2{
            fmt.Println("Version is needed")
        }else{
            version, err := strconv.Atoi(command[1])
            if err != nil {
                version := command[1]
                Install(version)

            }else{
                Install(version)
            }
        }
        break
    case "change":
        if len(command) < 2{
            fmt.Println("Version is needed")
        }else{
            version, err := strconv.Atoi(command[1])
            if err != nil {
                version := command[1]
                Change(version)

            }else{
                Change(version)
            }
        }
        break
    case "list":
        List()
        break
    case "init":
        Init()
        break
    case "update":
        Update()
        break
    case "help":
        Man()
        break
    default:
        fmt.Println("Command Not Found")
    }
    
}

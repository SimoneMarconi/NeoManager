package handlers

import (
    "time"
    "fmt"
)

func StartBuilding(ticker *time.Ticker){
    loadings := []string{"|", "/", "-", "|", "/", "-"}
    iter := 0
    for {
        select{
        case <- ticker.C:
            fmt.Printf("\x1bc")
            fmt.Print("building" + loadings[iter])
        iter = (iter + 1) % 6
        }
    }
}

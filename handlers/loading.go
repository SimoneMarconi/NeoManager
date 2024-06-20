package handlers

import (
    "time"
    "fmt"
)

func StartLoading(ticker *time.Ticker, msg string){
    loadings := []string{"|", "/", "-", "\\"}
    iter := 0
    for {
        select{
        case <- ticker.C:
            fmt.Printf("\x1bc")
            fmt.Print(msg + " " + loadings[iter])
        iter = (iter + 1) % len(loadings)
        }
    }
}

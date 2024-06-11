package handlers

import "fmt"

func Man(){
    text := `no action provided try:
    init = to set up the Neovim repository locally
    update = to update the Neovim repository
    install (version number) = to install the provided version
    change (version number) = to change the current version`
    fmt.Print(text)
}

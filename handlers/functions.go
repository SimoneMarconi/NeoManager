package handlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func Install(v interface{}){
    wd, _ := os.Getwd()
    dir := wd + "/source"
    if _, err := os.Stat(dir); os.IsNotExist(err){
        fmt.Println("NeoManager needs to be initiated, run NeoManager init")
    }
    build(v)
}

func Change(v interface{}){
    var version string
    switch v.(type){
    case int:
        version = fmt.Sprintf("v0.%d.0", v)
        break
    case string:
        version = v.(string)
        break
    }
    home, errHome:= os.UserHomeDir()
    if errHome != nil{
        log.Panic("Home not found")
    }
    dir := home + "/.NeoManager/versions/" + version
    log.Println(dir)

    if _, err := os.Stat(dir); os.IsNotExist(err){
        fmt.Println("Version not found")
        os.Exit(0)
    }
    exe := dir + "/bin/nvim"
    destination := home + "/.local/bin"
    cp := exec.Command("cp", exe, destination)
    out, errCp := cp.CombinedOutput()
    if errCp != nil{
        log.Println(string(out))
        log.Panic("Cp Error")
    }
    checkout := exec.Command("git", "checkout", version)
    checkout.Dir = "source"

    output, err := checkout.CombinedOutput()

    if err != nil {
        log.Panic(string(output))
        fmt.Printf("Error listing files: %s\n", err)
        return
    }

    fmt.Println(string(output))
}

func ChangeRepo(v string){
    checkout := exec.Command("git", "checkout", v)
    checkout.Dir = "source"

	output, err := checkout.CombinedOutput()

	if err != nil {
		fmt.Printf("Error changing Repository\n")
		return
	}

    fmt.Println(string(output))
}

func Testing(){
}

func build(v interface{}){
    home, _:= os.UserHomeDir()
    dir := home + "/.NeoManager"
    var ans string
    if _, err := os.Stat(dir); os.IsNotExist(err){
        fmt.Println("Directory .NeoManager not found, create it? (y/n)")
        fmt.Scan(&ans)
        if ans == "y"{
            mkdir := exec.Command("mkdir", ".NeoManager")
            mkdir.Dir = home
            err1 := mkdir.Run()
            if err1 != nil {
                fmt.Println(err)
            }
            vFolder:= exec.Command("mkdir", "versions")
            vFolder.Dir = dir
            err2 := vFolder.Run()
            if err2 != nil {
                log.Panic("Error creating versions directory")
            }
        }else{
            fmt.Println("Aborting")
        }
    }
    var vName string
    switch v.(type){
    case int:
        vName = fmt.Sprintf("v0.%d.0", v)
        break
    case string:
        if v != "latest"{
            fmt.Println("Wrong Version")
            os.Exit(0)
        }
        vName = v.(string)
    }
    //checking if the version is already installed
    verify := home + "/.NeoManager/versions/" + vName
    if _, err := os.Stat(verify); err == nil{
        checkBin := verify + "/bin"
        if _, err := os.Stat(checkBin); err == nil{
            fmt.Println("Version already installed")
            os.Exit(0)
        }
    }
        dBuild := exec.Command("mkdir", vName)
        dBuild.Dir = home + "/.NeoManager/versions"
        err := dBuild.Run()
        if err != nil {
            log.Panic(err)
    }
    //changing repo to execute the right build
    ChangeRepo(vName)
    buildTicker := time.NewTicker(500 * time.Millisecond)
    go StartBuilding(buildTicker)
    cBuild := exec.Command("make", "CMAKE_BUILD_TYPE=RelWithDebInfo")
    wd, errWd:= os.Getwd()
    if errWd != nil {
        log.Panic("Error Wd")
    }
    sourceDir := wd + "/source"
    cBuild.Dir = sourceDir
    out, errMake := cBuild.CombinedOutput()
    if errMake != nil{
        log.Println(string(out))
        log.Panic("Make error")
    }
    buildTicker.Stop()
    installTicker := time.NewTicker(500 * time.Millisecond)
    fmt.Print("\x1bc")
    fmt.Println("Successfull build")
    go StartInstalling(installTicker)
    home, errHome:= os.UserHomeDir()
    if errHome != nil{
        log.Panic("Home Error")
    }
    versionPath := home + fmt.Sprintf("/.NeoManager/versions/%s", vName)
    command := "CMAKE_INSTALL_PREFIX=" + versionPath
    cInstall := exec.Command("make", command, "install")
    cInstall.Dir = sourceDir
    out, errInstall := cInstall.CombinedOutput()
    if errInstall != nil{
        log.Println(string(out))
        log.Panic("Install Error")
    }
    fmt.Println("Installation finished")
    installTicker.Stop()
    //copying the compiled exe to the local bin directory
    exe := versionPath + "/bin/nvim"
    destination := home + "/.local/bin"
    cp := exec.Command("cp", exe, destination)
    out, errCp := cp.CombinedOutput()
    if errCp != nil{
        log.Println(string(out))
        log.Panic("Cp Error")
    }
}

func List(){
    home, errHome:= os.UserHomeDir()
    if errHome != nil{
        log.Panic("Home not found")
    }
    ls := exec.Command("ls")
    ls.Dir = home + "/.NeoManager/versions"
    out, errLs := ls.CombinedOutput()
    if errLs != nil{
        log.Panic("ls did not work")
    }
    fmt.Println(string(out))
}

func Init(){
    wd, _:= os.Getwd()
    dir := wd + "/source"
    if _, err := os.Stat(dir); os.IsNotExist(err){
        mkdir := exec.Command("mkdir", "source")
        mkdir.Dir = wd
        clone := exec.Command("git", "clone", "https://github.com/neovim/neovim.git", ".")
        clone.Dir = dir
        errDir := mkdir.Run()
        if errDir != nil {
            log.Panic("Error making source")
        }
        errClone := clone.Run()
        if errClone != nil{
            log.Panic("Error cloning repository")
        }
    }else{
        fmt.Println("NeoManager was already initialized")
    }
}

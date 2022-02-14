package main

import(
    "fmt"
    "os/exec"
)

func main(){    
    c := exec.Command("cmd", "/C", "del", "D:\\a.txt")

    if err := c.Run(); err != nil { 
        fmt.Println("Error: ", err)
    }   
}
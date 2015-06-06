package main
import (
    "fmt"
    "strings"
)

func main(){
    s:="/home/bino/cdn/bino/gobook.pdf"

    fmt.Println(strings.LastIndex(s,".pdf")==len(s)-4)
}

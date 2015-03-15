package main

import ("os"
        "fmt"
        "time"
        )
func f(){

    filename := "wangfakang.txt"
    //fout,err := os.Open(filename)
    //fout,err := os.Create(filename)
    fout,err := os.OpenFile(filename,os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
    if err != nil {
       fmt.Printf("file exit r create fail...\n")
    }
    defer fout.Close()

    if err != nil {
        fmt.Println("filename,err")
            return
    }
    for i :=0 ;i<10 ;i++{
        fout.WriteString("hello world\n")
        fout.Write([]byte("i can\n"))
    }


}
func main(){
    println("start")
    go f()
    println("end")
    time.Sleep(2e9)


}





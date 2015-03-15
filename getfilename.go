package main

import "fmt"
import "runtime"


func main(){
    pc,file,line,ok := runtime.Caller(0)
    if ok {
        fmt.Println("Func Name=" + runtime.FuncForPC(pc).Name())
        fmt.Printf("file :  %s  line = %d\n",file,line)
    }
}


package main

import "sync"
import "fmt"
import "time"
//import "os/exec"

func fu(){
   /* cmd := exec.Command("ls")
    out,err := cmd.Output()
    if err != nil {
        println(err)
    }
    cmd.Run()*/
    //fmt.Printf("%s\n",out)
    fmt.Printf("%s\n","default")
    return

}

func show(){
    for {
       select {
            case i := <-ch :
            fmt.Printf("sub produce %d \n",i)
            default:
                fu()
               // time.Sleep(1e9)
       }
    }
}

var ch = make(chan int)

var mutex sync.Mutex
func mshow(i int){
    fmt.Printf("LOCK \n")
    mutex.Lock()
    ch <- i
    mutex.Unlock()
    fmt.Printf("UNLOCK \n")
}
func ss(i int){
    fmt.Printf("go  go  go \n")
    go mshow(i)
}


func main(){
flag :=0
            go show()
    time.Sleep(3e9)
    for i:=0; i< 30 ;i++ {
        if flag == 0{
            flag++
        }
    time.Sleep(1e9)
        go mshow(i)
    }
    time.Sleep(100e9)
}



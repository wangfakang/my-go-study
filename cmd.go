package main
import "os/exec"
import "fmt"
//import "strings"
//import "os"
//import "io/ioutil"


/*
func checkmonitor(hostname string)string{
    cmd := exec.Command("/usr/bin/montool","-s"," "+hostname)
    out,err := cmd.Output()
    if err != nil {
        msg := fmt.Sprintf("exec montool : %s",err.Error())
        print(msg)
       // fileLogger.Error(msg)
    }
    err = cmd.Run()
    strout := string(out)
    status := strings.Split(strout,"\t")
    //if status[1] == "blocked" {
      //  return ""
    //}
    println(status[1])
    return hostname
}

func main(){
    checkmonitor(os.Args[1])
}

*/


func main(){

   // println(os.Args[0],os.Args[1],os.Args[2])
    //cmd := exec.Command("/usr/bin/montool","-s"," "+os.Args[1])
    //cmd := exec.Command("/home/users/wangfakang/montool","-s"," yf-nsop-dev01.yf01")
    cmd := exec.Command("/home/users/wangfakang/hello")
    out,err := cmd.Output()
    if err != nil {
        println("error")
     }
    err = cmd.Run()
    if err != nil {
        //println(err.Error())
    }
    str := string(out)
    fmt.Printf("%s\n",str)
}


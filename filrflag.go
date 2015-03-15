package main


import (
        "fmt"
        "strings"
        "os"
        "bufio"
        "io"
        )




func main() {
    f,err := os.Open("./wangfakang.txt")
    if err != nil {
        return
    }
    buff := bufio.NewReader(f)
    i := 0
    sum := ""
    state := 0
    for {
        line,err := buff.ReadString('\n')
        line = strings.TrimSpace(line)
        if line == "BEGIN"  {
            state = 1
            continue
        }

        if err != nil || io.EOF == err || line == "END"{
            break
        }
        if state == 1 {
            //println(len(line))
            if i != 0 && len(line) != 0{
                line += ","
            }
            i++
            sum += line
        }

    fmt.Printf("%s\n",sum)
    }
}



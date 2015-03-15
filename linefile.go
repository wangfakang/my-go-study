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
        if err != nil || io.EOF == err || line == "END"{
                break
        }

        if line == "BEGIN"  {
            state = 1

        }

        if state == 1 || state ==0 {
            if state == 1 {
                state += 1
            }
            continue
        }else if state == 2{
            println(len(line))
            if i != 0 && len(line) != 0{
                sum += ","
            }
            i++
            sum += line

        }
    }
    if len(sum)>0 && sum[len(sum)-1] == ','{
       // sum = sum[:len(sum)-1]
    }
    fmt.Printf("9%s9",sum)
}

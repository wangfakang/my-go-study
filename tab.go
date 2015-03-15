
package main


import "fmt"
import "bufio"
import "os"



func main(){
    cnt := 0
    var mycmd[50]string
    for {
        str := bufio.NewReader(os.Stdin)
       input,_ := str.ReadString('\n')
        
        ctype := string([]byte(input)[0])
        
        switch ctype[0] {
            case '\t' :

                break
            case '\n' :

                break
            default:
                mycmd[cnt] = ctype
                    cnt++
                break
        }

    }
        fmt.Printf("%s\n",mycmd)
}




/*
    c := '\t'
    fmt.Scanf("%c",&c)
    if c == '\t' {
        fmt.Println("tab")
    }
    if c == '\n' {
        fmt.Println("hang")
    }

}
*/

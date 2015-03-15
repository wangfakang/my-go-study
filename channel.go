
package  main
import "fmt"
import "time"

   var chs = make(chan int)
func Count() {
        time.Sleep(2e9)
    chs <- 1
    
        fmt.Println("Counting")
}


func main() {

        // for i := 0; i < 10; i++ {
          //   chs[i] = make(chan int)
                 go Count()
        // }
   //  for _, ch := range(chs) {
         fmt.Println(<-chs)
     //}
}




package main

import ("fmt" )

func prodect(ch chan int){
    for i:=0 ;i<10;i++{
//        println("start func p...")
        ch <- i 
        fmt.Printf("world\n")
    }
    return
}

func costume(ch chan int){
    for { 
  //      println("start func c...")
      //  v := <-ch
        fmt.Println( <-ch )
        fmt.Printf("hello\n")
    }
}

func add( a *int){
    *a++
}

func main(){
    var ch chan int
    ch <- 1
    //println("start...")
   //ch := make(chan int)
    //println("end...")
    go  prodect(ch)
    go  costume(ch)
    fmt.Println( <-ch )
    return 
}

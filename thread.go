
package main

import ("fmt" ; "time")

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
    //println("start...")
//    ch := make(chan int)
    //println("end...")
 //   go  prodect(ch)
   // go  costume(ch)
    a:=1
    go add(&a)
    go add(&a)
    time.Sleep(1e9)
    fmt.Printf("%d\n",a)
    return 
}

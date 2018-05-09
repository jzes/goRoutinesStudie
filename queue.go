package main


import(
        "fmt"
        //"io/ioutil"
)

func anySlowProcess(in chan int, out chan int){
        for m := range in{
                fmt.Println(m)
                out <- m*2
        }
}

func main(){
        var chanelA = make(chan int, 5)
        var chanelB = make(chan int, 5)
        chanelA <- 1
        chanelA <- 2
        chanelA <- 3
        chanelA <- 4
        chanelA <- 5
        close(chanelA)
        go anySlowProcess(chanelA, chanelB)
        for i := 0; i < 5; i++ {
                fmt.Println(<-chanelB)
        }
        
}
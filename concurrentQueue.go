package main


import(
        "fmt"
        "time"
        //"io/ioutil"
)

var queueChanel = make(chan int, 20)

func anySlowProcess( out chan int, workerName string, load int){
        for m := range queueChanel{
                time.Sleep(time.Duration(load) * time.Second)
                fmt.Println("Into Process of ",workerName,m)
                out <- m*2
        }
}

func loadQueue(){
        for i := 1; i <= 20; i++ {
                queueChanel <- i
        }
        close(queueChanel)
}

func main(){
        
        var chanelOutputA = make(chan int, 20)
        
        go anySlowProcess(chanelOutputA, "A", 1)
        go anySlowProcess(chanelOutputA, "B", 5)
        go anySlowProcess(chanelOutputA, "C", 2)
        loadQueue()
        for i := 0; i < 20; i++ {
                fmt.Println(<-chanelOutputA, "saida")
                
        }
}
package main
import(
        "fmt"
)



func multiplies(num int, chanelOut chan <- int){
        fmt.Println("Ini.:", num)
        retorno := 0
        for count := 0; count < num; count++ {
                retorno += count*2
        }
        fmt.Println("End.:", retorno)
        chanelOut <- retorno
}


func main(){
        var result int
        chanelB := make(chan int)
        
        for i := 2; i < 30; i++ {
                go multiplies(i, chanelB)        
        }
        
        for i := 2; i < 30; i++ {
                result = <- chanelB
                fmt.Println("Chanel out.:",result)
        }
}
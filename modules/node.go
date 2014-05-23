package modules

import(
    "fmt"
    "time"
)


func InitNode() {
    ip, ch := IpRequest(2)
    fmt.Println("I'm node ",ip)
    go func(){
        for{
            s := <-ch
            a,b := Pars(s)
            fmt.Println(a)
            fmt.Println(b)
            
            time.Sleep(150*time.Millisecond)
        }
    }()
}





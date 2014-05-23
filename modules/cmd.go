package modules

import(
    "fmt"
    "bufio"
    "time"
    "os"
)

func CmdInit(ro chan string) {
    fmt.Println("\n  Welcome!\n")
    bio := bufio.NewReader(os.Stdin)
    go func(){
        for {
            fmt.Print("dmn@date~ ")
            line,_,_ := bio.ReadLine()
            ro<-string(line)
            time.Sleep(50*time.Millisecond)
        }
    }()
}

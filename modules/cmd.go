package modules

import(
    "fmt"
    "bufio"
    "time"
    "os"
    "strings"
)

func CmdInit() {
    fmt.Println("\n  Welcome!\n")
    bio := bufio.NewReader(os.Stdin)
    go func(){
        for {
            fmt.Print("dmn@date~ ")
            line,_,_ := bio.ReadLine()
            parsCommand(string(line))
            time.Sleep(100*time.Millisecond)
        }
    }()
}

func parsCommand(s string){
    arr := strings.Split(s," ")
    switch arr[0]{
        case "test":
            fmt.Println("Hi, root!")
        case "send":
            if len(arr)<4 {
                fmt.Println("   usage : send [sender] [receiver] [msg]")
            }else{
                Send("0.0", arr[1], "snd:"+arr[2]+":"+arr[3])
            }
        case "snd":
            Send(arr[1],arr[2],arr[3])
        case "corrupt":
            Send("0.0",arr[1],"corr")
        default:
            fmt.Println("Wrong command.")
    }

}

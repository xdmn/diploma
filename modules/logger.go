package modules

import(
    "os"
)

var (
    log_chan chan string
)

func InitLog(LOGGING bool) func(string) {
    if LOGGING {
        log_file, err := os.Create("log.txt")
        if err!=nil {
            panic(err)
        }
        log_chan = make(chan string,10)
        go func(){
            for{
                s:=<-log_chan
                log_file.WriteString(s+"\n")
            }
        }()
        return log
    }else{
        return func (string){}
    }
}

func log(s string) {
    log_chan<-s
}

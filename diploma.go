package main

import (
    "fmt"
    "time"
    "flag"
    "strconv"
    "os"
    "os/signal"
    "github.com/xdmn/diploma/modules"
)

var (
    NODES   int  = 5
    SERVERS int  = 1
    CLIENTS int  = 1
    LOGGING bool = true
)

func init() {
    flag.IntVar(&NODES,   "n", NODES,   "количество промежуточных узлов")
    flag.IntVar(&SERVERS, "s", SERVERS, "количество серверов")
    flag.IntVar(&CLIENTS, "c", CLIENTS, "количество клиентов")
    flag.BoolVar(&LOGGING,"l", LOGGING, "запись в лог-файл")
}


func main() {
    flag.Parse()
    itoa := strconv.Itoa
//    atoi := strconv.Atoi


//Logging
    log := modules.LogInit(LOGGING)
    log("\n*** Log file ***")
    t:=time.Now().Local()
    const layout = "Jan 2, 2006 at 3:04pm"
    log(t.Format(layout))
    log("Network:")
    log("    "+itoa(CLIENTS)+" client(s);")
    log("    "+itoa(NODES)+  " node(s);")
    log("    "+itoa(SERVERS)+" server(s).")
//

//Interrupt chan
    interruptChan := make(chan os.Signal, 1)
    signal.Notify(interruptChan, os.Interrupt)


//Main loop
    for {
        select {
            case <-interruptChan :
                fmt.Println("\nCTRL=C: Stop.")
                log("--End--\n")
                time.Sleep(150*time.Millisecond)
                return
            default:
            }
    }



    modules.Cmd("Hello, xdmn!")
    modules.Client("Hello, client!")
    modules.Node("Hello, node!")
    modules.Server("Hello, server!")
    
    
    
}

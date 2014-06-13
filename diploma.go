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

//Logging
    log := modules.InitLog(LOGGING)
    log("\n*** Log file ***")
    t:=time.Now().Local()
    const layout = "Jan 2, 2006 at 3:04pm"
    log(t.Format(layout))
    log("Network:")
    log("    "+strconv.Itoa(CLIENTS)+" client(s);")
    log("    "+strconv.Itoa(NODES)+  " node(s);")
    log("    "+strconv.Itoa(SERVERS)+" server(s).")

//Command line
    modules.CmdInit()

//DHCP (classes: client, node, server)
    modules.InitDHCP(3)

//Nodes
    for i:=0;i<CLIENTS;i++ {
        modules.InitUnit(1)
    }
    for i:=0;i<NODES;i++ {
        modules.InitUnit(2)
    }
    for i:=0;i<SERVERS;i++ {
        modules.InitUnit(3)
    }

//Interrupt chan
    interruptChan := make(chan os.Signal, 1)
    signal.Notify(interruptChan, os.Interrupt)


//Main loop
    for {
        select {
            case <-interruptChan :
                fmt.Println("\nCTRL-C: Stop.")
                log("--End--\n")
                time.Sleep(150*time.Millisecond)
                return 
            default:
            }
    }
    
}

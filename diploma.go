package main

import (
//    "fmt"
    "time"
    "flag"
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

    //log is function
    log := modules.LogInit(LOGGING)

    log("*** Log file ***")

    modules.Cmd("Hello, xdmn!")
    modules.Client("Hello, client!")
    modules.Node("Hello, node!")
    modules.Server("Hello, server!")
    
    
    //tmp
    time.Sleep(150*time.Millisecond)
    
}

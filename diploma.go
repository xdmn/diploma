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
    log := modules.InitLog(LOGGING)
    log("\n*** Log file ***")
    t:=time.Now().Local()
    const layout = "Jan 2, 2006 at 3:04pm"
    log(t.Format(layout))
    log("Network:")
    log("    "+itoa(CLIENTS)+" client(s);")
    log("    "+itoa(NODES)+  " node(s);")
    log("    "+itoa(SERVERS)+" server(s).")

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


/*enc1:=modules.Encrypt("test",123)
enc2:=modules.Encrypt(enc1,124)
enc3:=modules.Encrypt(enc2,125)
fmt.Println(enc1)
fmt.Println(enc2)
fmt.Println(enc3)
fmt.Println(modules.Decrypt(modules.Decrypt(modules.Decrypt(enc3,125),124),123))
*/
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

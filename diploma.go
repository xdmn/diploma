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

//Routing chan
    routing := make(chan string, 2)

//Command line
//    modules.CmdInit(routing)

//DHCP (classes: client, node, server)
    modules.InitDHCP(3)

//Nodes
    modules.InitNode()
    modules.InitNode()
    //modules.Send("2.1","2.2","Hi, 2.2. I'm 2.1")
    //modules.Send("me1","2.1","test")
    //modules.Send("me2","2.2","test")
    modules.Send("2.1","2.2","test")
    //modules.Send("2.2","2.1","test")


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
        case s := <-routing :
                fmt.Println("yeah: ",s)
            case <-interruptChan :
                fmt.Println("\nCTRL-C: Stop.")
                log("--End--\n")
                time.Sleep(150*time.Millisecond)
                return 
            default:
            }
    }



    modules.Client("Hello, client!")
    modules.Server("Hello, server!")
    
    
    
}

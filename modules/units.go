package modules

import(
    "fmt"
    "time"
    "strings"
)


func InitUnit(class int) {
    ip, ch := IpRequest(class)
    //fmt.Println("I'm unit ",ip)
    go func(){
    var(
        keys map[string]string
        cook map[string]string
        corr bool 
    )
    switch class {
        case 1:
            keys = make(map[string]string)
        case 2:
            cook = make(map[string]string)
        case 3:
            corr = false
    }
        for{
            s := <-ch
            a,b := Pars(s)
            switch class {
                case 1:
                    clientTact(ip,keys,a,b)
                case 2:
                    nodeTact(ip,cook,a,b)
                case 3:
                    corr = serverTact(ip,corr,a,b)
            }
            time.Sleep(30*time.Millisecond)
        }
    }()
}

func clientTact(ip string, keys map[string]string,a string, b string){
    n1:="2.1"
    n2:="2.2"
    n3:="2.3"

    if a=="0.0" {
        arr := strings.Split(b, ":")
        switch arr[0] {
            case "snd":
                if len(arr)==3 {
                    Send(ip,n1,makeOnion(n1,n2,n3,arr[1],arr[2]))       
                    log("Client "+ip+" send message to "+arr[1]+" via "+n1+" "+n2+" "+n3)
                }
            }
    }else{
        if(b[0]=='('){
            b = destroyOnion(n1,n2,n3,b)
        }
        str:="Client "+ip+" recieved: "+ b+" from "+ a
        fmt.Println(str)
        log(str)
        arr := strings.Split(b,"!")
        if len(arr)>1{
            if arr[1] == "externallink"{
                str:="Client "+ip+" deanonimized by external request."
                fmt.Println(str)
                log(str)
            }
        }
    }
}

func nodeTact(ip string,cook map[string]string, a string, b string){
    if _,ok := cook[a]; !ok{
        dec := decrypt(b,"key"+ip)
        arr := strings.Split(dec,":")
        cook[arr[0]] = a
        Send(ip,arr[0],dec[len(arr[0])+1:])
    }else{
        enc := encrypt(b,"key"+ip)
        Send(ip,cook[a],enc)
    }
}

func serverTact(ip string,corr bool, a string, b string) bool{
    if a=="0.0" {
        if b  == "corr"{
            corr = true
            str:="Server "+ip+" corrupted!"
            fmt.Println(str)
            log(str)
        }   
        return corr

    }else{
        str:="Server "+ip+" recieved: "+ b+" from "+ a
        fmt.Println(str)
        log(str)
        ans := "answer"
        if corr {
            ans += "!externallink"
        }
        Send(ip,a,ans)
        return corr
    }
}



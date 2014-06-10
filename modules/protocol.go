package modules

import(
    "strconv"
    "strings"
//    "fmt"
)

var (
    classNum []int
    chans    map[string]chan string
)


func InitDHCP(classes int){
    chans = make(map[string]chan string)
    classNum = make([]int,classes+1)
}

func IpRequest(class int)(string, chan string){
    classNum[class]++
    ip := strconv.Itoa(class) + "." + strconv.Itoa(classNum[class])
 //   fmt.Println(ip,"start")
    chans[ip] = make(chan string,5)
    //chans[ip]<-"iii" 
    return ip,chans[ip]
}

func getCh(ip string)chan string{
    return chans[ip]
}

func encrypt(msg string, key string) string{
    return "("+key+"("+msg
}

func decrypt(msg string, key string) string{
    arr := strings.Split(msg,"(")
    if arr[0]=="" && arr[1]==key{
        return msg[len(key)+2:]
    }else{
        return "bad_decrypted_message"
    }

}

func makeOnion(n1 string, n2 string, n3 string, ser string, msg string) string{
    res := encrypt(ser+":"+msg, "key"+n3)
    res  = encrypt(n3 +":"+res, "key"+n2)
    res  = encrypt(n2 +":"+res, "key"+n1)
    return res
}

func destroyOnion(n1 string, n2 string, n3 string, oni string) string{
    msg := decrypt(oni,"key"+n1)
    msg  = decrypt(msg,"key"+n2) 
    msg  = decrypt(msg,"key"+n3) 
    return msg
}

func Send(sen string, rec string, msg string){
//    fmt.Println("Send: ",sen,rec,msg)
    getCh(rec)<-(":"+sen+":"+msg)
}


func Pars(msg string) (string, string){
    arr := strings.Split(msg,":")
//    fmt.Println("par: ",msg)
    return arr[1], msg[ len(arr[1])+2 : ]
}

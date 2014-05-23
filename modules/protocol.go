package modules

import(
    "strconv"
    "strings"
    "fmt"
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
    fmt.Println(ip,"start")
    chans[ip] = make(chan string,5)
    //chans[ip]<-"iii" 
    return ip,chans[ip]
}

func getCh(ip string)chan string{
    return chans[ip]
}

func Encrypt(msg string, key int) string{
    return "("+strconv.Itoa(key)+"("+msg
}

func Decrypt(msg string, key int) string{
    arr := strings.Split(msg,"(")
    s_key := strconv.Itoa(key) 
    if arr[0]=="" && arr[1]==s_key{
        return msg[len(s_key)+2:]
    }else{
        return "bad_decrypted_message"
    }

}

func Send(sen string, rec string, msg string){
    fmt.Println("Send: ",sen,rec,msg)
    getCh(rec)<-(":"+sen+":"+msg)
}


func Pars(msg string) (string, string){
    arr := strings.Split(msg,":")
    fmt.Println("par: ",msg)
    return arr[1], msg[ len(arr[1])+2 : ]
}

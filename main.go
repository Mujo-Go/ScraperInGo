package main
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
)
func main(){
    urls :=os.Args[1:]
    for _,url:= range urls{
        resp,_:=http.Get(url)
        defer resp.Body.Close()
        bytes,_:=ioutil.ReadAll(resp.Body)
        fmt.Print(string(bytes))

    }
}

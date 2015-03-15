package main

import (
            "fmt"
            "net/http"
  //          "io/ioutil"
//            "os"
//            "bytes"
            "encoding/json"
       )


func main() {
                resp, err := http.Get("http://machine.rms.baidu.com/cetus/index.php?r=restApi/StartOperate")
                if err != nil {
            
                    panic(err)
                                    }
                    //fmt.Printf("%s", resp.Body)

             //   body , _ := ioutil.ReadAll(resp.Body)
                defer resp.Body.Close() // 注意,这里并不读取resp.Body, 而resp.Body有大概700mb未读取
                
               // var out bytes.Buffer
               // err = json.Indent(&out, body, "", "\t")
                //if err != nil {
                    //     return
                  //   }

                type Hreq struct{
                    success bool
                    message string
                    param   string
                }
                
                var quest Hreq
                json.Unmarshal(resp.Body,&quest)

                fmt.printf("%+v",quest)

                                        
                        
}


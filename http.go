package main

import (
            "fmt"
            "net/http"
            "io/ioutil"
            "os"
            "bytes"
            "encoding/json"
       )

func main() {
                //resp, err := http.Get("http://api.rms.baidu.com/v1/unified_list/selfReboot")
                resp, err := http.Get("http://machine.rms.baidu.com/cetus/index.php?r=restApi/StartOperate")
                if err != nil {
                    panic(err)
                                    }
                    //fmt.Printf("%s", resp.Body)

                body , _ := ioutil.ReadAll(resp.Body)
                body1 := string(body)
                    fmt.Printf("%s", body1)
                resp.Body.Close() // 注意,这里并不读取resp.Body, 而resp.Body有大概700mb未读取
                if err != nil {
                                    return
                               }

                        var out bytes.Buffer
                        err = json.Indent(&out, body, "", "\t")
                        if err != nil {
                                   return
                                }

                            out.WriteTo(os.Stdout)
                           // str:=string(out)
                           //println(str)
                            
                        
        }


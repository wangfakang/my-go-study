package main


import (
    _"fmt"
    "os/exec"
    "strings"
)




func SendHi(msg ,number string){
    cmd := exec.Command("nc","hz01-nuomiop.hz01.baidu.com","15440")
    cmd.Stdin = strings.NewReader("send_msg  " + number  + "  " +  msg)
    cmd.Run()
}


func main(){
  // tag := "************************************ \r"
  // histr := "注意：下面机器已经重启成功: \r " + "successHost" + " \r请及时做相应处理,谢谢!"
  // histr = tag + histr + "\r" + tag
            tag := "************************************ \r"
            histr := "注意:下面机器重启失败:\r " + "hostlist[i]" + " \r失败原因：机器出故障，请到noah发故障单处理"
            ll:= "@王发康"
            histr = tag + histr + "\r" + tag
            histr += ll
   SendHi(histr,"SKY19928")

}

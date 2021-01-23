package main

// 需要你事先安装 github.com/beevik/ntp 第三方包
// 之所以使用网络的校准时间，是防止你篡改系统时间提取打开胶囊
// 每个月只有1天可以取出你的东西
import (
    "fmt"
    "ntp"
    "crypto/md5"
)

func password() {
    // 定义好你的 secret，然后hash后输出密码，当然也可以直接输出密码
    secret := ""
    h := md5.New()
    h.Write([]byte(secret))
    bs := h.Sum(nil)
    fmt.Printf("%x\n", bs)
}

func main() {

    time, err := ntp.Time("time.windows.com")
    if err != nil {
        panic(err)
    }
    day := time.Format("02")

    // 每个月15号可以输出一次
    if day == "23" {
        password()
    } else {
        fmt.Println("Not today. Wait till 15th.")
    }
}

package g

import (
	//"fmt"
	"os/exec"
	"strings"

	"github.com/axgle/mahonia"
	"encoding/base64"
	"github.com/yinheli/qqwry"
)

var (
	Dbug_log = true //调试状态
	IP_qqwry = qqwry.NewQQwry("qqwry.dat")
	coder = base64.NewEncoding("012345abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ6789+/") //bs64自定义编码表
)

//=====================================


func Cmdexec(cmd string, system string) string {
	defer Panic_Err() //异常处理
	var c *exec.Cmd
	var data string
	if system == "windows" {
		argArray := strings.Split("/c "+cmd, " ")
		c = exec.Command("cmd", argArray...)
	} else {
		c = exec.Command("/bin/sh", "-c", cmd)
	}
	out, _ := c.Output()
	data = string(out)
	if system == "windows" {
		dec := mahonia.NewDecoder("gbk")
		data = dec.ConvertString(data)
	}
	return data
}

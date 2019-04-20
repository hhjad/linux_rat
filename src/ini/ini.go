package ini

import (
	//	"fmt"
	//"html/template"
	//"log"
	//"net/http"
	//"time"
	//"os"
	//"os/signal"
	//"strconv"
	//"reflect"
	//"g"
	//"strings"
	//"time"
	//"www"
	//"seo_scan"
	//	"strings"

	//"github.com/widuu/goini"
	"github.com/Unknwon/goconfig"
)

func GetValue_ini(file_data, KEY1, KEY2 string) string { //读取INI配置文件
	c, err := goconfig.LoadConfigFile(file_data)
	if err != nil {
		return ""
	}
	datax, errx := c.GetValue(KEY1, KEY2)
	if errx != nil {
		return ""
	}
	return datax
	//	conf := goini.SetConfig(file_data) //goini.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置
	//	return conf.GetValue(KEY1, KEY2)   //database是你的[section]，username是你要获取值的key名称
}

//c, _ := goconfig.LoadConfigFile("server.ini")
//	data, err := c.GetValue("Abc", "path")
//	fmt.Println(err)
//	fmt.Println(data)
//	//s, _ := goconfig.LoadConfigFile("server.ini", "server.ini")
//	c.SetValue("Abc", "path", "pppppp")
//	//var dst bytes.Buffer
//	goconfig.SaveConfigFile(c, "server.ini")
//	data, err = c.GetValue("Abc", "path")
//	fmt.Println(err)
//	fmt.Println(data)

func SetValue_ini(file_data, section, KEY1, KEY2 string) bool { //设置INI
	c, err := goconfig.LoadConfigFile(file_data)
	if err != nil {
		return false
	}
	//errx:=
	c.SetValue(section, KEY1, KEY2)
	//	if errx!=nil{
	//		return false
	//	}
	errxx := goconfig.SaveConfigFile(c, file_data)
	if errxx != nil {
		return false
	}
	return true
	//	conf := goini.SetConfig(file_data)
	//	//return conf.SetValue(section, KEY1, KEY2)
	//	fmt.Printf("uuuu:%v %v %v %v\n", file_data, section, KEY1, KEY2)
	//	if conf.SetValue(section, KEY1, KEY2) {
	//		fmt.Println("okxx")
	//	} else {
	//		fmt.Println("noxx")
	//	}
}

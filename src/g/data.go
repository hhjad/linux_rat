package g

import (
	"fmt"
	"time"
	"path/filepath"
	"os"
	"net/url"
	//"log"
	"strings"
	//"encoding/base64"
	"strconv"
	"math/rand"
	//"os/exec"
	//"github.com/axgle/mahonia"
	"crypto/md5"
	"encoding/hex"
)

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func Url_Parse(cz string) string { //URL转码
	cz_data, cz_err := url.Parse(cz)
	if cz_err == nil {
		cz = cz_data.Path
	}
	return cz
}

func Wlwz_qqwry(ip string) string { //物理位置
	//defer Public_file.Panic_Err() //异常处理
	//sip := fmt.Sprintf("%v", ip)
	s := ""
	Try(func() {
		if len(ip) < 6 {
			s = "IP except"
		}
		IP_qqwry.Find(ip)
		//	q := qqwry.NewQQwry("qqwry.dat")
		//	q.Find("119.131.118.50")
		//	log.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)
		s = fmt.Sprintf("%v[%v %v]", ip, IP_qqwry.Country, IP_qqwry.City)
		//return s
	}, func(e interface{}) {
		s = ip
	})
	return s
}

//=====================================================
//func Log_msgbox(index int, msg string) {
//	log.Printf("%v---%v", index, msg)
//}
//
//func Path_Encode(www_path string) string { //转换路径
//	www_path = strings.Replace(www_path, "?", "#1#", -1)
//	www_path = strings.Replace(www_path, "、", "#2#", -1)
//	www_path = strings.Replace(www_path, "╲", "#3#", -1)
//	www_path = strings.Replace(www_path, "/", "#4#", -1)
//	www_path = strings.Replace(www_path, "*", "#5#", -1)
//	www_path = strings.Replace(www_path, "“", "#6#", -1)
//	www_path = strings.Replace(www_path, "”", "#7#", -1)
//	www_path = strings.Replace(www_path, "<", "#8#", -1)
//	www_path = strings.Replace(www_path, ">", "#9#", -1)
//	www_path = strings.Replace(www_path, "|", "#10#", -1)
//
//	return www_path
//}
//
//func Path_Decode(www_path string) string { //解密路径
//	www_path = strings.Replace(www_path, "#1#", "?", -1)
//	www_path = strings.Replace(www_path, "#2#", "、", -1)
//	www_path = strings.Replace(www_path, "#3#", "╲", -1)
//	www_path = strings.Replace(www_path, "#4#", "/", -1)
//	www_path = strings.Replace(www_path, "#5#", "*", -1)
//	www_path = strings.Replace(www_path, "#6#", "“", -1)
//	www_path = strings.Replace(www_path, "#7#", "”", -1)
//	www_path = strings.Replace(www_path, "#8#", "<", -1)
//	www_path = strings.Replace(www_path, "#9#", ">", -1)
//	www_path = strings.Replace(www_path, "#10#", "|", -1)
//
//	return www_path
//}
//
//func Base64_Encode(data string) string { //编码
//	return base64.StdEncoding.EncodeToString([]byte(data))
//}
//func Base64_Decode(data string) string { //解码
//	uDec, err := base64.URLEncoding.DecodeString(data) //base64.StdEncoding.DecodeString(data)
//	if err != nil {
//		ciphertext := strings.Replace(data, " ", "", -1)
//		k, err0 := base64.StdEncoding.DecodeString(ciphertext)
//		if err0 != nil {
//			return ""
//		}
//		return string(k)
//	}
//	return string(uDec)
//}
//
////自定义BS64
//func Base64_Encode_x(src string) string { //编码
//	return string([]byte(coder.EncodeToString([]byte(src))))
//}
//func Base64_Decode_x(src string) string { //解码
//	enbyte, err := coder.DecodeString(src)
//	if err != nil {
//		return ""
//	}
//	return string(enbyte)
//}

//num:=RandInt64(1500,2000)
//        fmt.Println(num)
func RandInt(minxx, maxxx string) string { //生成指定范围内随机数
	min, _ := strconv.ParseInt(minxx, 10, 64)
	max, _ := strconv.ParseInt(maxxx, 10, 64)
	if min == 0 {
		min = 1
	}
	if min >= max || min == 0 || max == 0 {
		//fmt.Printf("==minxx%v,maxxx%v=====min%v,max%v=====\n", minxx, maxxx, min, max)
		return strconv.FormatInt(max, 10)
	}

	return strconv.FormatInt(rand.Int63n(max-min)+min, 10)
	//	defer Panic_Err() //异常处理
	//	data := "0"
	//	for index := 0; index < 1; index++ {
	//		min, _ := strconv.ParseInt(minxx, 10, 64)
	//		max, _ := strconv.ParseInt(maxxx, 10, 64)

	//		if min >= max || min == 0 || max == 0 {
	//			//fmt.Printf("ppp%vppp\n", "oooooo")
	//			data = strconv.FormatInt(max, 10)
	//		} else {
	//			data = strconv.FormatInt(rand.Int63n(max-min)+min, 10)
	//		}

	//	}
}

//生成随机字符串
func GetRandomString(x int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < x; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func BFB(bx string) bool { //百分比计算比例
	randx1, _ := strconv.Atoi(RandInt("1", "100")) //抽取随机数
	randx2, _ := strconv.Atoi(bx)
	if randx2 >= randx1 {
		//fmt.Printf("==%v======%v==OK\n", randx1, randx2)
		return true
	}
	return false
}

//=====================================================
//func Cmdexec(cmd string, system string) string {
//	defer Panic_Err() //异常处理
//	var c *exec.Cmd
//	var data string
//	if system == "windows" {
//		argArray := strings.Split("/c "+cmd, " ")
//		c = exec.Command("cmd", argArray...)
//	} else {
//		c = exec.Command("/bin/sh", "-c", cmd)
//	}
//	out, _ := c.Output()
//	data = string(out)
//	if system == "windows" {
//		dec := mahonia.NewDecoder("gbk")
//		data = dec.ConvertString(data)
//	}
//	return data
//}

func Md5_hex(data string) string {
	h := md5.New()
	h.Write([]byte(data))                                    // 需要加密的字符串为 123456
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
}

func Url_http_deal(Host, Path_url, href_src string) string { //对HTTP路径进行处理
	//	Host := "127.0.0.1:500"                  //网址
	//	Path_url := "/qrqwer/weqr/qw/er/qwe/123" //当前请求路径
	//	//href_src := "..//skin/tuiedu/less/style.css"
	//	href_src := "style.css"
	//href_src地址为远程地址时处理
	if strings.Contains(href_src, "http://") || strings.Contains(href_src, "https://") {
		s := strings.Split(href_src, "/")
		s[2] = Host
		href_src = strings.Join(s, "/") //改变路径
		return href_src
	}
	//href_src 格式整理
	txd := "http://" //头 http   https
	if strings.Contains(href_src, "http://") {
		bb := []byte(href_src)
		txd = string(bb[0:7])
		href_src = string(bb[7:len(bb)])
	}
	if strings.Contains(href_src, "https://") {
		bb := []byte(href_src)
		txd = string(bb[0:8])
		href_src = string(bb[8:len(bb)])
	}
	ss := strings.NewReplacer(`///`, `/`, `//`, `/`)
	href_src = ss.Replace(href_src)

	//	bb := []byte(Path_url)
	//	if strings.Contains(string(bb[len(bb)-1:len(bb)]), "/") {
	//		Path_url = string(bb[0 : len(bb)-1])
	//	}
	//路径转换
	if strings.Contains(href_src, "..") {
		bb := []byte(href_src)
		sj := string(bb[0:2])
		if sj == ".." {
			s := strings.Split(Path_url, "/")
			if strings.Contains(s[len(s)-1], ".") { //删除文件名
				s = s[:len(s)-2]
			} else {
				s = s[:len(s)-1]
			}
			Path_url = strings.Join(s, "/") //改变路径
			href_src = string(bb[2:len(bb)])
		}
		return txd + Host + Path_url + href_src
	} else {
		ss := []byte(href_src)
		if string(ss[0:2]) == "./" {
			href_src = string(ss[1:len(ss)])
		}

		s := strings.Split(href_src, "/")
		if len(s) == 1 {
			return txd + Host + "/" + href_src
		}
		return txd + Host + href_src
	}
	return txd + Host + Path_url + href_src
}

func Url_src_deal(Host, Path_url, href_src string) string { //对相对路径进行处理
	//	Host := "baidu.com"                        //网址
	//	Path_url := "/1111/2222/3333/444/555/666/" //当前请求路径
	//	//href_src := "..//skin/tuiedu/less/style.css"
	//	href_src := "../123/ffff.css"
	//======================
	txd := "http://" //头 http   https
	if strings.Contains(href_src, "http://") {
		bb := []byte(href_src)
		txd = string(bb[0:7])
		href_src = string(bb[7:len(bb)])
	}
	if strings.Contains(href_src, "https://") {
		bb := []byte(href_src)
		txd = string(bb[0:8])
		href_src = string(bb[8:len(bb)])
	}
	bb := []byte(Path_url)
	if strings.Contains(string(bb[len(bb)-1:len(bb)]), "/") {
		Path_url = string(bb[0 : len(bb)-1])
	}

	//======================
	if strings.Contains(href_src, "..") {
		bb := []byte(href_src)
		sj := string(bb[0:2])
		if sj == ".." {
			s := strings.Split(Path_url, "/")
			if strings.Contains(s[len(s)-1], ".") { //删除文件名
				s = s[:len(s)-2]
			} else {
				s = s[:len(s)-1]
			}
			Path_url = strings.Join(s, "/") //改变路径
			href_src = string(bb[2:len(bb)])
		}
		//fmt.Printf("yeeeeeey:%s\n", txd+Host+Path_url+href_src)
		return txd + Host + Path_url + href_src
	} else {
		s := strings.Split(href_src, "/")
		if len(s) == 1 {
			//fmt.Printf("yaaaaaay:%v\n", txd+Host+"/"+href_src)
			return txd + Host + "/" + href_src
		}
		//fmt.Printf("yaaaaaay:%s\n", txd+Host+href_src)
		return txd + Host + href_src
	}
	return txd + Host + Path_url + href_src
}

func Unix_Time(Unix string) string { //时间戳转换成时间
	f, err := strconv.ParseInt(Unix, 10, 64)
	if err == nil {
		str_time := time.Unix(f, 0).Format("2006-01-02 15:04:05")
		return str_time
	} else { //转换失败则返回当前时间戳
		str2 := fmt.Sprintf("%v", time.Now().Unix())
		return str2
	}
	return "0"
}

func TTime_Unixx(Time string) string { //日期转换成时间戳
	//toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                       //转化所需模板
	loc, _ := time.LoadLocation("Local")                      //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, Time, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                      //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                      //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr) //打印输出时间戳 1420041600
	return fmt.Sprintf("%v", sr)
	//=====================================================
	//	the_time, err := time.Parse("2006-01-02 15:04:05", Time)
	//	if err == nil {
	//		unix_time := the_time.Unix()
	//		//fmt.Println(unix_time)
	//		str2 := fmt.Sprintf("%v", unix_time)
	//		return str2
	//	} else {
	//		//fmt.Println("cccc:", time.Now().Unix())
	//		str2 := fmt.Sprintf("%v", time.Now().Unix())
	//		return str2
	//	}
	//	return "0"
}

func Time_Unix() string { //获取当前时间戳
	the_time, err := time.Parse("2006-01-02 15:04:05", Get_time())
	if err == nil {
		unix_time := the_time.Unix()
		//fmt.Println(unix_time)
		str2 := fmt.Sprintf("%v", unix_time)
		return str2
	} else {
		str2 := fmt.Sprintf("%v", time.Now().Unix())
		return str2
	}
	return "99"
}

func Get_time() string { //获取当前时间
	now := time.Now()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d", year, mon, day, hour, min, sec)
}

func Get_time2() string { //获取当前时间
	now := time.Now()
	year, mon, day := now.Date()
	//	hour, min, sec := now.Clock()
	return fmt.Sprintf("%d-%d-%d", year, mon, day)
}

func Get_CurrentDirectory() string { //获取程序运行路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		//beego.Debug(err)
		return "/"
	}
	return dir //strings.Replace(dir, "\\", "/", -1)
}


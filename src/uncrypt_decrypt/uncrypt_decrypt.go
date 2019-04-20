package uncrypt_decrypt

//对内容加密   解密    进行了移位操作 和字符分割替换 防止人家破解
//QQ29295842      BLOG:http://blog.csdn.net/webxscan
import (
	//"io"
	//"os"
	b64 "encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

//加密   内容先进行BS64加密   然后在转换成ANSI编码    在对ANSI进行移位操作
//chinese := "简体中文"
//str1 := base64.StdEncoding.EncodeToString([]byte(chinese))
//fmt.Println(str1)
//str2, _ := base64.StdEncoding.DecodeString(str1)
//fmt.Println(string(str2))
func Uncrypt(data string) string {
	Uncrypt_data := ""
	Uncrypt := b64.StdEncoding.EncodeToString([]byte(data))
	for i := 0; i < len(Uncrypt); i++ {
		ascii_data := fmt.Sprintf("%d", int([]byte(string(Uncrypt[i]))[0])+5)
		Uncrypt_data = Uncrypt_data + "S" + ascii_data
		//fmt.Printf("%c%d\n", Uncrypt[i], int([]byte(string(Uncrypt[i]))[0])+5)

	}
	return string(Uncrypt_data)
}

//解密
func Decrypt(data string) string {
	Decrypt_data := ""
	canSplit := func(c rune) bool { return c == 'S' }
	lisit := strings.FieldsFunc(data, canSplit) //[58 215 20 30]
	for _, value := range lisit {

		str := string(value)
		input, e := strconv.Atoi(str) //转换成字符
		if e != nil {
			continue //跳过
		}
		//str := strconv.Atoi(string(value))
		//fmt.Printf("%d", input)
		//fmt.Println(string(input - 5))
		Decrypt_data = Decrypt_data + string(input-5)
	}
	//fmt.Printf(Decrypt_data)
	sDec, _ := b64.StdEncoding.DecodeString(Decrypt_data)
	return string(sDec)
}

//func main() {
////var wireteString = "测试n"
//	f, _ := os.OpenFile("output1.txt", os.O_APPEND, 0666) //打开文件
//	//n, _ := io.WriteString(f, wireteString)               //写入文件(字符串)
//	//fmt.Printf("写入 %d 个字节n", n)
//	b, err := ioutil.ReadAll(f)
//	if err != nil {
//		panic(err)
//	}
//	data := client.Uncrypt(string(b)) //加密
//	fmt.Printf("%s\n", data)

//	fmt.Printf("%s\n", client.Decrypt(string(data))) //解密

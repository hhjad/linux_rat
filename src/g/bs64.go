package g

import (
	"strings"
	"encoding/base64"
)

//自定义BS64
func Base64_Encode_x(src string) string { //编码
	return string([]byte(coder.EncodeToString([]byte(src))))
}
func Base64_Decode_x(src string) string { //解码
	enbyte, err := coder.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(enbyte)
}



func Path_Encode(www_path string) string { //转换路径
	www_path = strings.Replace(www_path, "?", "#1#", -1)
	www_path = strings.Replace(www_path, "、", "#2#", -1)
	www_path = strings.Replace(www_path, "╲", "#3#", -1)
	www_path = strings.Replace(www_path, "/", "#4#", -1)
	www_path = strings.Replace(www_path, "*", "#5#", -1)
	www_path = strings.Replace(www_path, "“", "#6#", -1)
	www_path = strings.Replace(www_path, "”", "#7#", -1)
	www_path = strings.Replace(www_path, "<", "#8#", -1)
	www_path = strings.Replace(www_path, ">", "#9#", -1)
	www_path = strings.Replace(www_path, "|", "#10#", -1)

	return www_path
}

func Path_Decode(www_path string) string { //解密路径
	www_path = strings.Replace(www_path, "#1#", "?", -1)
	www_path = strings.Replace(www_path, "#2#", "、", -1)
	www_path = strings.Replace(www_path, "#3#", "╲", -1)
	www_path = strings.Replace(www_path, "#4#", "/", -1)
	www_path = strings.Replace(www_path, "#5#", "*", -1)
	www_path = strings.Replace(www_path, "#6#", "“", -1)
	www_path = strings.Replace(www_path, "#7#", "”", -1)
	www_path = strings.Replace(www_path, "#8#", "<", -1)
	www_path = strings.Replace(www_path, "#9#", ">", -1)
	www_path = strings.Replace(www_path, "#10#", "|", -1)

	return www_path
}

func Base64_Encode(data string) string { //编码
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func Base64_Decode(data string) string { //解码
	uDec, err := base64.URLEncoding.DecodeString(data) //base64.StdEncoding.DecodeString(data)
	if err != nil {
		ciphertext := strings.Replace(data, " ", "", -1)
		k, err0 := base64.StdEncoding.DecodeString(ciphertext)
		if err0 != nil {
			return ""
		}
		return string(k)
	}
	return string(uDec)
}
package g

import (
	"fmt"


)

func Panic_Err() {
	if err := recover(); err != nil {
		if Dbug_log == true {
			fmt.Printf("\n=err===%v===", err)
		} else {
			//写入日志

		}
		//fmt.Println(err) //这里的err其实就是panic传入的内容，55
	}
}

func Log_msgbox(msg string) {
	//log.Printf("\n%v---%v", index, msg)
	fmt.Printf("%v   %v\n",Get_time(), msg)
}

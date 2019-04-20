package wss

import (
	"os/exec"
	"strings"
	"fmt"
	"io/ioutil"
	"os"
	"errors"
)

/*
更新主机信息
*/
func UpdateNode(t *Hub) {

	return
}



func CreateUser(args *UserInfo) (result *RstMsg, err error) {

	cmd := exec.Command("/bin/sh", "-c", "/usr/bin/id "+strings.Replace(args.User, "\n", "", -1))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		//return errors.New("Input Name")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		//return errors.New("Input Name")
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		//return errors.New("Input Name")
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		//return errors.New("Input Name")
	}

	if len(bytesErr) != 0 {
		UserAdd(args.User, args.Gid, args.Uid, args.AuthKey)
	}

	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
	}

	if errWait := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", errWait.Error())
	}

	//*result = RstMsg{200, string(bytes)}
	result = &RstMsg{RstCode: 200, Msg:"添加成功", Data:args}
	return result, nil
}


func UserAdd(us string, uid, gid int, key string) {
	fmt.Println("开始添加用户..")
	UserGid := fmt.Sprintf("%d", gid)
	UserId := fmt.Sprintf("%d", uid)
	fmt.Println(UserId)
	fmt.Println(UserGid)
	//var name string = strings.Replace(fmt.Sprintf("useradd -u %s -g %s %s -s /home/users/%s ", us, userId, UserGid, us, us),"\n", "", -1)
	//fmt.Println(name)
	rst, err := ShellCmd("useradd -u " + UserId + " -g " + UserGid + " " + us + " -d /home/users/" + us)
	//rst, err := ShellCmd(name)
	if err != nil {
		fmt.Println("命令执行失败", rst)
	}
	fmt.Printf("用户添加完成: %s\n", us)

	var ssh string = "/home/users/" + us + "/.ssh"

	_, errs := ShellCmd("mkdir " + ssh)
	//rst, err := ShellCmd(name)
	if errs != nil {
		fmt.Println("命令执行失败")
	}

	userFile := ssh + "/authorized_keys"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	fout.WriteString(key)
	fmt.Println("生成key成功")

	_, errown := ShellCmd("chown -R " + us + ":" + UserGid + " " + ssh)
	//rst, err := ShellCmd(name)
	if errown != nil {
		fmt.Println("添加权限组")
	}

	_, errmod7 := ShellCmd("chmod 700 " + ssh)
	//rst, err := ShellCmd(name)
	if errmod7 != nil {
		fmt.Println("添加权限组")
	}

	_, errmod6 := ShellCmd("chmod 600 " + userFile)
	if errmod6 != nil {
		fmt.Println("添加权限组")
	}


}



func ShellCmd(s string) (rst ShellStatus, error error) {
	us_key := exec.Command("/bin/sh", "-c", s)

	stdout, err := us_key.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return ShellStatus{Code: 400, Msg: err.Error()}, errors.New(err.Error())
	}

	stderr, err := us_key.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return ShellStatus{Code: 400, Msg: err.Error()}, errors.New(err.Error())
	}

	if err := us_key.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return ShellStatus{Code: 400, Msg: err.Error()}, errors.New(err.Error())
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return ShellStatus{Code: 400, Msg: err.Error()}, errors.New(err.Error())
	}

	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return ShellStatus{Code: 400, Msg: err.Error()}, errors.New(err.Error())
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return ShellStatus{Code: 400, Msg: err.Error()}, errors.New(err.Error())
	}

	if err := us_key.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return ShellStatus{Code: 400, Msg: err.Error()}, errors.New(err.Error())
	}
	//fmt.Printf("stdout: %s", bytes)
	return ShellStatus{Code: 200, Msg: string(bytes)}, nil
}

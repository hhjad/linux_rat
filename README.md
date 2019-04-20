LINUX集群控制(LINUX反弹式远控)

GET地址 https://github.com/webxscan/linux_rat

zzkey.com BY:QQ:879301117 

应用场景

因为工作需要，需要集群管理树莓派，而这些树莓派并没有外网IP所以无法使用SSH正向链接，所以必须采用反弹式链接。

我之前写过远控，就想到了使用反弹式链接。

使用WEBSOCKET 链接。    

应为时间问题，写的比较简陋还望大家多多见谅。

感谢  苦咖啡（voilet119@163.com）  伙计技术帮助


网站使用的BEEGO


服务端
\src\Client\Client_run.go
var Sx_url = "http://xxxxxxxxx.com/linux_ip.txt" //上线地址   


不多说了  大家自己看代码吧
有问题可以加我BY:QQ:879301117 


项目中保存的有运行效果图


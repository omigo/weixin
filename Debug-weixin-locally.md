如何本地调试微信公众号？
---------------------


微信公众号开发需要配置一个公网 IP 和 80 端口号，我们一般在内网开发完，然后发布到服务器，才能调试，十分不便。

使用 ngrok 可以解决这个问题，本地调试，无须发布到服务器。但国外 ngrok 转发服务器被墙了，好在国内有人搭建了一个，而且速度很快。


ngrok 用法：

1. 下载 ngrok 1.7（注意不是 ngrok 2） 客户端 
    
    链接: https://pan.baidu.com/s/1nvF9bKl 密码: bryg

2. 创建配置文件 ngrok.cfg，内容如下：

```
server_addr: "tunnel.phpor.me:4443"
trust_host_root_certs: false

```


3. 运行 `./ngrok -subdomain omigo -config=ngrok.cfg 3080`，其中 omigo

    是个性子域名，3080 是本地启动服务端口号。启动后提示如下：

```

ngrok                                                                (Ctrl+C to quit)
Tunnel Status                 online  
Version                       1.7/1.7
Forwarding                    http://omigo.tunnel.phpor.me -> 127.0.0.1:3080
Forwarding                    https://omigo.tunnel.phpor.me -> 127.0.0.1:3080  
Web Interface                 127.0.0.1:4040
# Conn                        0
Avg Conn Time                 0.00ms
```

于是直接访问 `http://omigo.tunnel.phpor.me`，请求可以转发到本地 3080 端口

4. 把微信公众号服务器 URL 配置为 `http://omigo.tunnel.phpor.me/weixin` 即可

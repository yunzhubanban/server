# server

《互联网＋校园 云上课堂助手和班级看板软件》服务

## 快速开始

1. 创建数据卷

```bash
$ docker volume create yzbb_data
```

2. 编写配置文件

```toml
... 正在更新中
```

将上述内容命名为 `config.toml` 并保存至数据卷下

> 默认为 `/var/lib/docker/volumes/yzbb_data/_data/config.toml`

3. 启动

<details>
    <summary>Docker Hub</summary>

```bash
$ docker run --name=yunzhubanban-server -p 127.0.0.1:9922:9922 -v yzbb_data:/config -v yzbb_data:/data -v yzbb_data:/.yunzhubanban -d yunzhubanban/server:latest
```

</details>

<details>
    <summary>编译安装</summary>

```bash
$ gf build main.go -a amd64 -s linux
$ docker build -t yunzhubanban/server:unofficial .
$ docker run --name=yunzhubanban-server -p 127.0.0.1:9922:9922 -v yzbb_data:/config -v yzbb_data:/data -v yzbb_data:/.yunzhubanban -d yunzhubanban/server:latestdocker run --name=yunzhubanban-server -p 127.0.0.1:9922:9922 -v yzbb_data:/config -v yzbb_data:/data -v yzbb_data:/.yunzhubanban -d yunzhubanban/server:unofficial
```

有关 `gf` 的相关内容，请参阅：[开发工具 - GoFrame官网](https://goframe.org/pages/viewpage.action?pageId=1114260)

</details>

现在，《互联网＋校园 云上课堂助手和班级看板软件》服务 已安装至你的设备中

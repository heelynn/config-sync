# config-sync

**注意：**`master`分支可能处于开发之中并**非稳定版本**，请通过 tag 下载稳定版本的源代码，或通过[release](https://github.com/heelynn/config-sync/releases)下载已编译的二进制可执行文件。

`config-sync`是一个支持通过`注册中心`、`配置中心`进行配置文件生成的工具服务，在配置生成之后，可以执行指定命令，实现诸如`重启服务`、`重载配置`等操作。

## 安装

### 下载安装包

- 下载最新版本安装包：https://github.com/heelynn/config-sync/releases/latest

## 启动服务
- windows 

进入`bin`目录，双击`startup.cmd`文件


- macOS/Linux

进入`bin`目录，执行`startup.sh`命令

## 配置文件
请参考配置文件示例[英文](https://github.com/heelynn/config-sync/blob/main/doc/application.yaml.example)/[中文](https://github.com/heelynn/config-sync/blob/main/doc/application.yaml.example_zh)
目前支持的配置项：

`config` : 从配置中心同步配置，目前支持`nacos`，`不需要此功能可以不进行配置`
- 从配置中心拉取的配置，写入本地文件文件路径，并监控配置文件是否有变化，有变化则修改配置文件，拉取成功后会执行command配置的命令

`discovery` : 从注册中心同步配置，目前支持`nacos`，`不需要此功能可以不进行配置`
- 注册中心地址，用于定时拉取服务列表，拉取成功后会根据模版文件生成配置文件，并执行command配置的命令
- 模版文件示例：(以nginx的upstream配置为例[upstream.tmpl](https://github.com/heelynn/config-sync/blob/main/doc/upstream.tmpl.example)，自定义其他模版完全遵照以下命名规则取值即可)
  - {{.Name}} 代表服务名
  - {{- range .Instances }} 代表遍历服务列表，下面有子项：
    - {{.Host}} 代表服务地址
    - {{.Port}} 代表服务端口
    - {{.Weight}} 代表服务权重 
- `nginx upstream` 使用示例：
- command配置：可以执行自定义命令，比如重启nginx，执行`restart-nginx.sh`
  - 在nginx中`慎重使用{{.Weight}}`，如果数字不合法配置测试会不通过，比如权重为零，会导致nginx启动失败
  - `restart-nginx.sh`，脚本内容如下：
```shell
#!/bin/bash

# 测试 Nginx 配置文件
nginx -t

# 检查上一个命令的退出状态码
if [ $? -eq 0 ]; then
    echo "Nginx 配置测试成功。正在重新加载 Nginx..."

    # 重新加载 Nginx 配置
    nginx -s reload
else
    # 由于某些原因导致 Nginx 配置测试失败
    # 此处是为了改动配置文件，会导致config-sync下次比对时文件内容发生变化，触发重新生成配置文件以及执行命令
    echo " " > /etc/nginx/conf.d/aaa.conf
    # 输出结果
    echo "Nginx 配置测试失败。"
fi
```



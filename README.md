# j

**注意：**`master`分支可能处于开发之中并**非稳定版本**，请通过 tag 下载稳定版本的源代码，或通过[release](https://github.com/heelynn/conf-sync/releases)下载已编译的二进制可执行文件。

`conf-sync`是一个支持通过`注册中心`、`配置中心`进行配置文件生成的工具服务，在配置生成之后，可以执行指定命令，实现诸如`重启服务`、`重载配置`等操作。

## 安装

### 下载安装包

- 下载最新版本安装包：https://github.com/heelynn/conf-sync/releases/latest

- windows 
进入`bin`目录，双击`startup.cmd`文件

- macOS/Linux

 ```shell
 # 使用shell
 $ mkdir -p $HOME/.j/bin $HOME/.j/version
 # 选择liunx/macos版本下载安装包，下载地址：https://github.com/heelynn/j/releases/latest
 
 # 将j文件解压到$HOME/.j/bin目录下 
 # 给j文件添加可执行权限
 $ chmod +x $HOME/.j/bin/j
 
 # 配置环境变量（适用于 bash、zsh）
 $ export JAVA_HOME="$HOME/.j/java"
 $ export PATH="$HOME/.j/bin:$JAVA_HOME/bin:$PATH"
 ```

## 使用方法

### 下载并安装Java
将下载好的Java安装包`Redhat、openJDk、OracleJDK等`解压到`$HOME/.j/version`目录下。
建议修改为简单文件夹名，如`Java8、Java11、Java17`，方便管理。
 ```shell
 # 列出已安装的Java版本
 $ ls $HOME/.j/version
 Java8
 Java11
 Java17
 ```

### 查看当前Java版本
`使用j命令`查看当前Java版本，与`$HOME/.j/version`目录下Java版本对应
 ```shell
 $ j ls
  java8
  java11
 ```

`使用j命令`切换Java版本，`j ls` 命令查看当前已安装的Java版本
 ```shell
 $ j use java8
 ```
### 验证Java版本
 ```shell
 $ java -version
 ```



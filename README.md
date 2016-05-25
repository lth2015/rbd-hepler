RBD命令辅助工具
-------------------------------------------

#### 起因
-------------------------------------------

在使用Kubernetes和Ceph RBD构建容器云应用时,创建RBD操作一般要经过如下几个步骤:

1. 创建RBD块设备
```bash
rbd create temp -s 1024 -p rbd
```

2. 映射RBD块设备
```bash
rbd map temp
```

3. 查看映射的设备路径
```bash
rbd showmapped | grep temp
/dev/rbd1
```

4. 格式化块设备
```bash
mkfs.ext4 /dev/rbd1
```

5. 取消映射
```bash
rbd unmap /dev/rbd1
```

>每次创建一个RBD块设备都需要执行上述命令,过程繁琐而且容器出错,
>需要编写工具,来快速完成上述操作.

#### 工具安装
-------------------------------------------

下载源码:

```bash
git clone https://github.com/lth2015/rbd-hepler.git
```

编译二进制包:

```bash
cd rbd-helper
export GOPATH=$PWD
export GOBIN=$PWD/bin
cd src/rbdctl
go get "github.com/urfave/cli"
go install
```

成功之后可以直接去rbd-hlper/bin下寻找rbdctl二进制文件

#### 使用命令
-------------------------------------------

##### 创建RBD块并格式化

```bash
./rbdctl create -n RBD块名字 -s 大小(单位兆) -f 文件系统 -p rbd池的名称
```

##### 显示本机已经映射的RBD块设备

```bah
./rbdctl show
```

##### 删除一个RBD块设备

```bash
./rbdctl delete -n RBD块名字
```


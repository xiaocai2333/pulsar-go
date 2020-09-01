# pulsar-go
This is pulsar go client!

## 如何编译安装pulsar C++

### 1. 下载pulsar源码到go/src/github.com/apache目录下
```shell script
    cd ~/go/src/github.com/apache/
    git clone https://github.com/apache/pulsar.git
```

### 2. 安装编译pulsar需要的依赖
```shell script
    apt-get install cmake libssl-dev libcurl4-openssl-dev liblog4cxx-dev \
      libprotobuf-dev protobuf-compiler libboost-all-dev google-mock libgtest-dev libjsoncpp-dev
```
执行这一步时需要注意是否之前安装过protobuf-complie，如果之前安装过，将之前安装的全部删除干净。达到执行
```shell script
    which protoc
```
时，结果为空。

### 3. 编译安装pulsar
进入下载的pulsar项目的cpp目录下，新建build的目录用来编译。
```shell script
    cd ~/go/src/github.com/apache/pulsar/pulsar-client-cpp
    mkdir build && cd build
    cmake .. -DBUILD_TESTS=OFF
    make -j12
    sudo make install
```
这里编译的时候关闭了test，如果有必要，你可以打开它。安装完成之后，就可以在go项目中使用pulsar了。
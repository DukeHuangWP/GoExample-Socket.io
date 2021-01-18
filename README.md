A golang socket.io example(server and clinet) with the go-socket.io and mofadeyunduo.

## Install

### For server
```
cd ./Server
go get github.com/mofadeyunduo/go-socket.io-client
```

### For client
```
cd ./Client
go get github.com/googollee/go-socket.io
```

## File Tree
```
├─Server
│      go.mod → Server 端 go-module 版控
│      go.sum → Server 端 go-module checksum
│      Socket-IO-Server.go →  Server端含(收/廣播)功能介面
│
└─Client
        go.mod → Client 端 go-module 版控
        go.sum → Client 端 go-module checksum
        Socket-IO-Client.go →  Client端含(收/發)功能,Command Line介面
```

## Demo

![Demo](https://github.com/DukeHuangWP/GoExample-Socket.io/blob/master/Demo.gif?raw=true)



# 拉取ens用户地址及使用这些地址注册did流程

## 1. 拉取ens用户地址到本地sql数据库ens.db

golang代码地址：
git@github.com:rockiecn/ens-log.git

使用方法：
先设置要开始拉取的区块数量，以太坊主链从1700万块开始才有ens用户，所以先设置成17000000:
go run main.go setlocal -b=17000000
然后开始拉取ens用户到本地ens.db:
go run main.go update

## 2. 为ens.db的所有用户查询并更新余额

前面拉取的数据没有余额数据，所以需要使用这个脚本为所有用户写入余额数据：
git@github.com:rockiecn/set-balance.git

使用方法：
go run main.go xx
这里的xx是任意数字，表示从第几个记录开始更新。 可以先设置成1， 后面万一程序重启了可以从最后那个序号开始继续更新剩下用户的余额。

## 3. 使用这些地址注册did

注册代码：
git@github.com:rockiecn/createDID.git

使用方法：
go run main.go -eth=dev -sk=246755f26638d9d1d47acca04fdc31dd3731e6c4ea3cdcc439425ee4f0aba880 -start=1 -balance=0
这个-eth表示注册的链，dev为开发链，test为测试链。
-sk是固定地址
-start表示从第几个用户开始注册，万一程序断掉了可以从上一次注册最后那个地址开始继续注册。
-balance表示允许注册的最小余额地址，0表示注册所有地址。

## 备注：

所有脚本要使用ens-log脚本生成的同一个ens.db数据库，所以地址拉取完成后要把ens.db拷贝到set-balance做余额更新，最后把更新好的ens.db再拷贝到createDID脚本目录做注册。

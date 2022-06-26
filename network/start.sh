#!/bin/bash

if [[ `uname` == 'Darwin' ]]; then
    echo "Mac OS"
    export PATH=${PWD}/hyperledger-fabric-darwin-amd64-1.4.12/bin:$PATH
fi
if [[ `uname` == 'Linux' ]]; then
    echo "Linux"
    export PATH=${PWD}/hyperledger-fabric-linux-amd64-1.4.12/bin:$PATH
fi

echo "一、清理环境"
./stop.sh

echo "二、生成证书和秘钥（ MSP 材料），生成结果将保存在 crypto-config 文件夹中"
cryptogen generate --config=./crypto-config.yaml

echo "三、创建排序通道创世区块"
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./config/genesis.block -channelID firstchannel

echo "四、生成通道配置事务'appchannel.tx'"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/appchannel.tx -channelID appchannel

echo "五、为 BENZ 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/BENZAnchor.tx -channelID appchannel -asOrg BENZ

echo "六、为 TESLA 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/TESLAAnchor.tx -channelID appchannel -asOrg TESLA

echo "区块链 ： 启动"
docker-compose up -d
echo "正在等待节点的启动完成，等待10秒"
sleep 10

BENZPeer0Cli="CORE_PEER_ADDRESS=peer0.benz.com:7051 CORE_PEER_LOCALMSPID=BENZMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/benz.com/users/Admin@benz.com/msp"
BENZPeer1Cli="CORE_PEER_ADDRESS=peer1.benz.com:7051 CORE_PEER_LOCALMSPID=BENZMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/benz.com/users/Admin@benz.com/msp"
TESLAPeer0Cli="CORE_PEER_ADDRESS=peer0.tesla.com:7051 CORE_PEER_LOCALMSPID=TESLAMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/tesla.com/users/Admin@tesla.com/msp"
TESLAPeer1Cli="CORE_PEER_ADDRESS=peer1.tesla.com:7051 CORE_PEER_LOCALMSPID=TESLAMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/tesla.com/users/Admin@tesla.com/msp"

echo "七、创建通道"
docker exec cli bash -c "$BENZPeer0Cli peer channel create -o orderer.carunion.com:7050 -c appchannel -f /etc/hyperledger/config/appchannel.tx"

echo "八、将所有节点加入通道"
docker exec cli bash -c "$BENZPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$BENZPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$TESLAPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$TESLAPeer1Cli peer channel join -b appchannel.block"

echo "九、更新锚节点"
docker exec cli bash -c "$BENZPeer0Cli peer channel update -o orderer.carunion.com:7050 -c appchannel -f /etc/hyperledger/config/BENZAnchor.tx"
docker exec cli bash -c "$TESLAPeer0Cli peer channel update -o orderer.carunion.com:7050 -c appchannel -f /etc/hyperledger/config/TESLAAnchor.tx"

# -n 链码名，可以自己随便设置
# -v 版本号
# -p 链码目录，在 /opt/gopath/src/ 目录下
echo "十、安装链码"
docker exec cli bash -c "$BENZPeer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"
docker exec cli bash -c "$TESLAPeer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"

# 只需要其中一个节点实例化
# -n 对应上一步安装链码的名字
# -v 版本号
# -C 是通道，在fabric的世界，一个通道就是一条不同的链
# -c 为传参，传入init参数
echo "十一、实例化链码"
docker exec cli bash -c "$BENZPeer0Cli peer chaincode instantiate -o orderer.carunion.com:7050 -C appchannel -n fabric-realty -l golang -v 1.0.0 -c '{\"Args\":[\"init\"]}' -P \"AND ('BENZMSP.member','TESLAMSP.member')\""

echo "正在等待链码实例化完成，等待5秒"
sleep 5

# 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作
echo "十二、验证链码"
docker exec cli bash -c "$BENZPeer0Cli peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"hello\"]}'"
docker exec cli bash -c "$TESLAPeer0Cli peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"hello\"]}'"
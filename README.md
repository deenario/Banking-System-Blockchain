
# Blockchain Based Election System

The Election system designed using Hyperledger Fabric.

## Setup

```./byfn.sh up -s couchdb```

```./byfn.sh down```

```docker exec cli peer chaincode install -n banking -l golang -p github.com/chaincode/banking -v 1.0```

```export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem```

```docker exec cli peer chaincode instantiate -o orderer.example.com:7050 --cafile $ORDERER_CA -C mychannel -c '{"Args":[]}' -n banking -v 1.0 -P "OR('Org1MSP.member', 'Org2MSP.member')"```

```docker exec cli peer chaincode invoke -o orderer.example.com:7050 --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n banking --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["addUser","_name","_address","example@email.com","_password","type_bank","I like this application"]}'```

```docker exec cli peer chaincode invoke -o orderer.example.com:7050 --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n banking --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["queryUser","example@email.com"]}'```

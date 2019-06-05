
# Blockchain Based Election System

The Election system designed using Hyperledger Fabric.

## Setup

```./byfn.sh up -s couchdb```

```./byfn.sh down```

```docker exec cli peer chaincode install -n banking -l golang -p github.com/chaincode/banking -v 1.0```

```export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem```

```docker exec cli peer chaincode instantiate -o orderer.example.com:7050 --cafile $ORDERER_CA -C mychannel -c '{"Args":[]}' -n banking -v 1.0 -P "OR('Org1MSP.member', 'Org2MSP.member')"```

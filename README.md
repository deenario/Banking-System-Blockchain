
# Blockchain Based Election System

The Election system designed using Hyperledger Fabric. 

## Installations
Ubuntu 16.04 or 18.04.
the first thing you have to do is install the pre-requists for this project. Do that by running the script
```./prereqs-ubuntu.sh```

``` sudo apt install git```

## Installing Hyperledger Fabric
In order to install hyperledger fabric. There are a few steps that needs to be done once.
To install hyperledger fabric

```git clone https://github.com/hyperledger/fabric-samples.git```

``` cd fabric-samples ```
``` ./scripts/bootstrap.sh ``` 

Run the following commands.

## Setup

```./byfn.sh up -s couchdb```

```./byfn.sh down```

```docker exec cli peer chaincode install -n banking -l golang -p github.com/chaincode/banking -v 1.0```

```export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem```

```docker exec cli peer chaincode instantiate -o orderer.example.com:7050 --cafile $ORDERER_CA -C mychannel -c '{"Args":[]}' -n banking -v 1.0 -P "OR('Org1MSP.member', 'Org2MSP.member')"```

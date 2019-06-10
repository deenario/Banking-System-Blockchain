
# Blockchain Based Banking System

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

once it is done. This will install hyperledger fabric in the system. Now go back to Desktop and start working on Hyperledger Fabric.

## Installing the Banking system using blockchain.

Pull the repository for yourself.

``` git clone https://github.com/deenario/Banking-System-Blockchain.git```

``` cd Banking-System-Blockchain/fabric/first-network ```

Run these commands to run the install the blockchain and its chaincode. Run them one by one.

```./byfn.sh up -s couchdb```

```docker exec cli peer chaincode install -n banking -l golang -p github.com/chaincode/banking -v 1.0```

```export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem```

```docker exec cli peer chaincode instantiate -o orderer.example.com:7050 --cafile $ORDERER_CA -C mychannel -c '{"Args":[]}' -n banking -v 1.0 -P "OR('Org1MSP.member', 'Org2MSP.member')"```

Run these 3 commands to create three new users for bank, business and individual.

```docker exec cli peer chaincode invoke -o orderer.example.com:7050 --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n banking --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["addUser","Bank","City","bank@gmail.com","bank", "1000000","bank","National Bank"]}'```

```docker exec cli peer chaincode invoke -o orderer.example.com:7050 --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n banking --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["addUser","Individual","City","individual@gmail.com","individual", "1000","individual","Student"]}'```

```docker exec cli peer chaincode invoke -o orderer.example.com:7050 --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n banking --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["addUser","Business","City","business@gmail.com","business", "1000","business","Business"]}'```


Do not close this terminal and open a new terminal in the Node_api directory.

run these two commands.

``` npm install ```

``` node enrollAdmin.js ```

``` node registerUser.js ```

``` node app.js ```

This will start the API server for the webpage. Keep these two terminals running. 

Go to the SRC folder and open Index.html webpage. Sign up And new user account and Enjoy the application.


## Empty Blockchain and start from scratch.

To empty blockchain. 
Go to fabric/first_network in blockchain and run 

```./byfn.sh down```

then go to node_APIs folder and delete hfkey or something similar to it. 
Once these two steps are done then start from the installing the blockchain step in this document . without git pulling the repo again. 

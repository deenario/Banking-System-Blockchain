
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

Do not close this terminal and open a new terminal in the Node_api directory.

run these two commands.

``` npm install ```

``` node enrollAdmin.js ```

``` node registerUser.js ```

``` node app.js ```

This will start the API server for the webpage. Keep these two terminals running. 

Go to the SRC folder and open Index.html webpage. Sign up And new user account and Enjoy the application.
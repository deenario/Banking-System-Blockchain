'use strict';
/*
* Copyright IBM Corp All Rights Reserved
*
* SPDX-License-Identifier: Apache-2.0
*/
/*
* Chaincode Invoke
*/

var Fabric_Client = require('fabric-client');
const BlockValidator = require('./BlockValidator');
var path = require('path');
var util = require('util');
var os = require('os');
var fs = require('fs');
var blocklistener = false;

module.exports = {

  invokeCreate: async function (request) {

    try {

      var fabric_client = new Fabric_Client();

      // setup the fabric network
      var channel = fabric_client.newChannel('mychannel');
      var peer = fabric_client.newPeer('grpc://localhost:7051');
      channel.addPeer(peer);
      var order = fabric_client.newOrderer('grpc://localhost:7050')
      channel.addOrderer(order);

      // load the base network profile and eventHub
      let event_hub = fabric_client.getEventHub(peer);

      // overlay the client profile over the network profile

      // setup the fabric network - get the channel that was loaded from the network profile
      var tx_id = null;

      //load the user who is going to unteract with the network
      var member_user = null;
      var store_path = path.join(__dirname, 'hfc-key-store');
      console.log('Store path:'+store_path);
      var tx_id = null;
      
      // create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
      let user_from_store = await Fabric_Client.newDefaultKeyValueStore({ path: store_path
      }).then((state_store) => {
        // assign the store to the fabric client
        fabric_client.setStateStore(state_store);
        var crypto_suite = Fabric_Client.newCryptoSuite();
        // use the same location for the state store (where the users' certificate are kept)
        // and the crypto store (where the users' keys are kept)
        var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
        crypto_suite.setCryptoKeyStore(crypto_store);
        fabric_client.setCryptoSuite(crypto_suite);
      
        // get the enrolled user from persistence, this user will sign all requests
        return fabric_client.getUserContext('user1', true);
      });

      if (user_from_store && user_from_store.isEnrolled()) {
        console.log('Successfully loaded user1 from persistence');
      } else {
        throw new Error('Failed to get user1.... run registerUserNetwork.js');
      }

      // get a transaction id object based on the current user assigned to fabric client
      tx_id = fabric_client.newTransactionID();
      console.log("Assigning transaction_id: ", tx_id._transaction_id);

      request = Object.assign(request, { txId: tx_id });
      // send the transaction proposal to the endorsing peers
      let results = await channel.sendTransactionProposal(request);
      console.log("TESTING TEST : " + results);
      var proposalResponses = results[0];
      var proposal = results[1];
      if (proposalResponses && proposalResponses[0].response &&
        proposalResponses[0].response.status === 200) {
        console.log('Transaction proposal was good');

        console.log(util.format(
          'Successfully sent Proposal and received ProposalResponse: Status - %s, message - "%s"',
          proposalResponses[0].response.status, proposalResponses[0].response.message));

        var request = {
          proposalResponses: proposalResponses,
          proposal: proposal
        };
        // set the transaction listener and set a timeout of 30 sec
        // if the transaction did not get committed within the timeout period,
        // report a TIMEOUT status
        var promises = [];
        var sendPromise = channel.sendTransaction(request);
        promises.push(sendPromise); //we want the send transaction first, so that we know where to check status

        // const bl = new BlockValidator(tx_id._transaction_id);
        // if (!blocklistener) {
        //   bl.StartBlockListener(event_hub);
        //   blocklistener = true;
        // }
        return {
          status: 200,
          message: util.format('Transaction ID: %s' , tx_id._transaction_id)
        }

      } else {
        console.log(proposalResponses[0].response.message.toString())
        return {
          status: 500,
          message: proposalResponses[0].response.message.toString()
        }
      }
    }
    catch (err) {
      console.log(err);
      return {
        status: 500,
        message: util.format("%s", err)
      }
    }
  }
}
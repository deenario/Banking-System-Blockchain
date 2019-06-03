'use strict';
/*
* Copyright IBM Corp All Rights Reserved
*
* SPDX-License-Identifier: Apache-2.0
*/
/*
 * Chaincode query
 */

var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');
var fs = require('fs');
var Response;
//make sure we have the profiles we need
var networkConfig = path.join(__dirname, './config/network-profile.json')
var clientConfig = path.join(__dirname, './config/client-profile.json');

module.exports = {

  invokeQuery: async function (request) {

    try {

      //
      var fabric_client = new Fabric_Client();

      // setup the fabric network
      var channel = fabric_client.newChannel('mychannel');
      var peer = fabric_client.newPeer('grpc://localhost:7051');
      channel.addPeer(peer);
      var order = fabric_client.newOrderer('grpc://localhost:7050')
      channel.addOrderer(order);
      var member_user = null;
      var store_path = path.join(__dirname, 'hfc-key-store');
      console.log('Store path:' + store_path);

      // create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
      let user_from_store = await Fabric_Client.newDefaultKeyValueStore({
        path: store_path
      }).then((state_store) => {
        // assign the store to the fabric client
        fabric_client.setStateStore(state_store);
        var crypto_suite = Fabric_Client.newCryptoSuite();
        // use the same location for the state store (where the users' certificate are kept)
        // and the crypto store (where the users' keys are kept)
        var crypto_store = Fabric_Client.newCryptoKeyStore({ path: store_path });
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

      // send the query proposal to the peer
      let query_responses = await channel.queryByChaincode(request);

      // query_responses could have more than one  results if there multiple peers were used as targets
      if (query_responses && query_responses.length == 1) {
        if (query_responses[0] instanceof Error) {
          console.error("error from query = ", query_responses[0]);
          Response = query_responses[0].toString();
          return {
            status: 500,
            message: Response
          }
        } else {
          console.log("Response is ", query_responses[0].toString());
          Response = query_responses[0].toString();
          return {
            status: 200,
            message: Response
          }
        }
      } else {
        console.log("No payloads were returned from query");
        Response = "No payloads were returned from query";
        return {
          status: 500,
          message: Response
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
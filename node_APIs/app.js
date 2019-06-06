'use strict';

//get libraries
const express = require('express');
var queue = require('express-queue');
const bodyParser = require('body-parser');
const request = require('request');
const path = require('path');

//create express web-app
const app = express();
const invoke = require('./invokeNetwork');
const query = require('./queryNetwork');

//declare port
var port = process.env.PORT || 8000;
if (process.env.VCAP_APPLICATION) {
  port = process.env.PORT;
}

// // app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
  extended: true
}));

app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  next();
});

//run app on port
app.listen(port, function () {
  console.log('app running on port: %d', port);
});

//-------------------------------------------------------------
//----------------------  POST API'S    -----------------------
//-------------------------------------------------------------

app.post('/api/adduser', async function (req, res) {

  console.log(req.body.name);
  console.log(req.body.address);
  console.log(req.body.email);
  console.log(req.body.password);
  console.log(req.body.user_type);
  console.log(req.body.biography);

  var request = {
    chaincodeId: 'banking',
    fcn: 'addUser',
    args: [
      req.body.name,
      req.body.address,
      req.body.email,
      req.body.password,
      req.body.user_type,
      req.body.biography
    ]
  };


  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send("The User with email: "+req.body.email+ " is stored in the blockchain with " +response.message);
    else
    res.status(response.status).send(response.message);
  }
});

app.post('/api/addtransaction', async function (req, res) {

  var request = {
    chaincodeId: 'banking',
    fcn: 'addTransaction',
    args: [
      req.body.transaction_ID,
      req.body.to,
      req.body.from,
      req.body.amount,
      req.body.comment
    ]
  };

  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send("The Transaction with ID: "+req.body.transaction_ID+ " is stored in the blockchain with " +response.message);
    else
    res.status(response.status).send(response.message);
  }
});

//-------------------------------------------------------------
//----------------------  GET API'S  --------------------------
//-------------------------------------------------------------

app.get('/api/queryuser', async function (req, res) {

  const request = {
    chaincodeId: 'banking',
    fcn: 'queryUser',
    args: [
      req.query.email,
      req.query.password
    ]
  };
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
    res.status(response.status).send(JSON.parse(response.message));
    else
    res.status(response.status).send(response.message);
  }
});

app.get("/api/test", (req, res, next) => {
  res.json(["Tony","Lisa","Michael","Ginger","Food"]);
 });

app.get('/api/querytransactions', async function (req, res) {

    const request = {
      chaincodeId: 'banking',
      fcn: 'queryTransactions',
      args: [req.query.from]
    };
    let response = await query.invokeQuery(request)
    if (response) {
      if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
      else
      res.status(response.status).send(response.message);
    }
});

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
var _time = "T00:00:00Z";

//declare port
var port = process.env.PORT || 8000;
if (process.env.VCAP_APPLICATION) {
  port = process.env.PORT;
}

app.use(bodyParser.json());

//Using queue middleware
app.use(queue({ activeLimit: 30, queuedLimit: -1 }));

//run app on port
app.listen(port, function () {
  console.log('app running on port: %d', port);
});

//-------------------------------------------------------------
//----------------------  POST API'S    -----------------------
//-------------------------------------------------------------

app.post('/api/addcampaign', async function (req, res) {

  var request = {
    chaincodeId: 'election',
    fcn: 'addCampaign',
    args: [

      req.body.campaign_ID,
      req.body.title,
      req.body.no_of_ballots,
      req.body.timestamp + _time,
      req.body.campaign_startdate + _time,
      req.body.campaign_enddate + _time,
      req.body.campaign_type,
      req.body.county,
      req.body.district

    ]
  };

  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Campaign with ID: "+req.body.campaign_ID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addballot', async function (req, res) {

  var request = {
    chaincodeId: 'election',
    fcn: 'addBallot',
    args: [
      req.body.ballot_ID,
      req.body.campaign_ID,
      req.body.timestamp
    ]
  };

  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Ballot with ID: "+req.body.ballot_ID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addcampaigncandidates', async function (req, res) {

  var request = {
    chaincodeId: 'election',
    fcn: 'addCampaignCandidates',
    args: [
    	req.body.campaign_ID,
	    req.body.user_ID
    ]
  };

  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "A campaign candidiate with ID: "+req.body.ballot_ID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addcampaignresults', async function (req, res) {

  var request = {
    chaincodeId: 'election',
    fcn: 'addCampaignResults',
    args: [
    	req.body.campaign_ID,
	    req.body.user_ID
    ]
  };

  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "A Ballot with ID: "+req.body.ballot_ID+ " has been casted and stored into the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});


//-------------------------------------------------------------
//----------------------  GET API'S  --------------------------
//-------------------------------------------------------------

app.get('/api/querycampaign', async function (req, res) {

  const request = {
    chaincodeId: 'election',
    fcn: 'queryCampaignCandidates',
    args: [req.query.campaign_ID]
  };
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: JSON.parse(response.message) });
    else
    res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryballot', async function (req, res) {

    const request = {
      chaincodeId: 'election',
      fcn: 'queryballot',
      args: [req.query.ballot_ID]
    };
    let response = await query.invokeQuery(request)
    if (response) {
      if(response.status == 200)
      res.status(response.status).send({ message: JSON.parse(response.message) });
      else
      res.status(response.status).send({ message: response.message });
    }
});

app.get('/api/querycampaigncandidates', async function (req, res) {

      const request = {
        chaincodeId: 'election',
        fcn: 'queryCampaignCandidates',
        args: [req.query.campaign_ID]
      };
      let response = await query.invokeQuery(request)
      if (response) {
        if(response.status == 200)
        res.status(response.status).send({ message: JSON.parse(response.message) });
        else
        res.status(response.status).send({ message: response.message });
      }
});

app.get('/api/queryallcampaigncandidates', async function (req, res) {

  const request = {
    chaincodeId: 'election',
    fcn: 'queryAllCampaignCandidates',
    args: []
  };
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: JSON.parse(response.message) });
    else
    res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/querycampaignresults', async function (req, res) {

  const request = {
    chaincodeId: 'election',
    fcn: 'queryCampaignresults',
    args: [req.query.campaign_ID]
  };
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: JSON.parse(response.message) });
    else
    res.status(response.status).send({ message: response.message });
  }
});
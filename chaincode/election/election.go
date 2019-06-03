package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)
//SmartContract is the data structure which represents this contract and on which  various contract lifecycle functions are attached
type SmartContract struct{
}

//Model for Users
type User struct {
	ObjectType			string `json:"Type"`
	User_ID				string `json:"user_ID"`
	Hash				string `json:"hash"`
	Status 				string `json:"status"`
	County				string `json:"county"`
	District			string `json:"district"`
}

type Campaign struct {
	ObjectType			string `json:"Type"`
	Campaign_ID 		string `json:"campaign_ID"`
	Title 				string `json:"title"`
	No_of_ballots		string `json:"no_of_ballots"`
	Timestamp			time.Time `json:"timestamp"`
	Campaign_startdate  time.Time `json:"campaign_startdate"`
	Campaign_enddate	time.Time `json:"campaign_enddate"`
	Campaign_type		string `json:"campaign_type"`
	County				string `json:"county"`
	District			string `json:"district"`
}

type Ballot struct {
	ObjectType 			string `json:"Type"`
	Ballot_ID 			string `json:"ballot_ID"`
	Campaign_ID			string `json:"campaign_ID"`
	Timestamp			time.Time `json:"timestamp"`
}

type CampaignCandidates struct {
	ObjectType			string `json:"Type"`
	Campaign_ID 		string `json:"campaign_ID"`
	User_ID				string `json:"user_ID"`
}

type CampaignResults struct {
	ObjectType			string `json:"Type"`
	Campaign_ID			string `json:"campaign_ID"`
	Candidate_ID		string `json:"candidate_ID"`
	Ballot_ID			string `json:"ballot_ID"`
	Timestamp 			time.Time `json:"timestamp"`
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init Firing!")
	return shim.Success(nil)
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Chaincode Invoke Is Running " + function)

	if function == "addUser" {
		return t.addUser(stub,args)
	}
	if function == "queryUser" {
		return t.queryUser(stub,args)
	}
	if function == "addCampaign" {
		return t.addCampaign(stub,args)
	}
	if function == "queryCampaign" {
		return t.queryCampaign(stub,args)
	}
	if function == "addBallot" {
		return t.addBallot(stub,args)
	}
	if function == "queryBallot" {
		return t.queryBallot(stub,args)
	}
	if function == "addCampaignCandidates"{
		return t.addCampaignCandidates(stub,args)
	}
	if function == "queryCampaignCandidates"{
		return t.queryCampaignCandidates(stub,args)
	}
	if function == "queryAllCampaignCandidates"{
		return t.queryAllCampaignCandidates(stub,args)
	}
	if function == "addCampaignResults" {
		return t.addCampaignResults(stub,args)
	}
	if function == "queryCampaignResults" {
		return t.queryCampaignResults(stub,args)
	}

	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Invoke did not find specified function " + function)
}

func (t *SmartContract) addUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect Number of Aruments. Expecting 5")
	}

	fmt.Println("Adding new User")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	

	user_ID := args[0]
	hash := args[1]
	status := args[2]
	county := args[3]
	district := args[4]


	// ======Check if User Already exists

	userAsBytes, err := stub.GetState(user_ID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if userAsBytes != nil {
		return shim.Error("The Inserted User ID already Exists: " + user_ID)
	}

	// ===== Create User Object and Marshal to JSON

	objectType := "user"
	user := &User{objectType, user_ID, hash, status, county, district}
	userJSONasBytes, err := json.Marshal(user)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save User to State

	err = stub.PutState(user_ID, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved User")
	return shim.Success(nil)
}

func (t *SmartContract) queryUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {

 	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	user_ID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"user\",\"user_ID\":\"%s\"}}", user_ID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addCampaign(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 9 {
		return shim.Error("Incorrect Number of Aruments. Expecting 9")
	}

	fmt.Println("Adding new Campaign")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th Argument Must be a Non-Empty String")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th Argument Must be a Non-Empty String")
	}
	if len(args[8]) <= 0 {
		return shim.Error("9th Argument Must be a Non-Empty String")
	}
	

	campaign_ID := args[0]
	title := args[1]
	no_of_ballots := args[2]
	timestamp, err1 := time.Parse(time.RFC3339, args[3])
	if err1 != nil {
		return shim.Error(err.Error())
	}
	 campaign_startdate, err1 := time.Parse(time.RFC3339, args[4])
	if err1 != nil {
		return shim.Error(err.Error())
	}
	campaign_enddate, err1 := time.Parse(time.RFC3339, args[5])
	if err1 != nil {
		return shim.Error(err.Error())
	}
	campaign_type := args[6]
	county := args[7]
	district := args[8]


	// ======Check if User Already exists

	campaignAsBytes, err := stub.GetState(campaign_ID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if campaignAsBytes != nil {
		return shim.Error("The Inserted campaign ID already Exists: " + campaign_ID)
	}

	// ===== Create campaign Object and Marshal to JSON

	objectType := "campaign"
	campaign := &Campaign{objectType, campaign_ID, title, no_of_ballots, timestamp, campaign_startdate, campaign_enddate, campaign_type, county, district}
	campaignJSONasBytes, err := json.Marshal(campaign)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save campaign to State

	err = stub.PutState(campaign_ID, campaignJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved campaign")
	return shim.Success(nil)
}

func (t *SmartContract) queryCampaign(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
	   return shim.Error("Incorrect number of arguments. Expecting 1")
   }

   campaign_ID := args[0]

   queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"campaign\",\"campaign_ID\":\"%s\"}}", campaign_ID)

   queryResults, err := getQueryResultForQueryString(stub, queryString)
   if err != nil {
	   return shim.Error(err.Error())
   }

   return shim.Success(queryResults)
}

func (t *SmartContract) addBallot(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect Number of Aruments. Expecting 3")
	}

	fmt.Println("Adding new Ballot")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	
	

	ballot_ID := args[0]
	campaign_ID := args[1]
	timestamp, err1 := time.Parse(time.RFC3339, args[2])
	if err1 != nil {
		return shim.Error(err.Error())
	}


	// ======Check if Ballot Already exists

	ballotAsBytes, err := stub.GetState(ballot_ID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if ballotAsBytes != nil {
		return shim.Error("The Inserted Ballot ID already Exists: " + ballot_ID)
	}

	// ===== Create Ballot Object and Marshal to JSON

	objectType := "ballot"
	ballot := &Ballot{objectType, ballot_ID, campaign_ID, timestamp}
	ballotJSONasBytes, err := json.Marshal(ballot)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save Ballot to State

	err = stub.PutState(ballot_ID, ballotJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Ballot")
	return shim.Success(nil)
}

func (t *SmartContract) queryBallot(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
	   return shim.Error("Incorrect number of arguments. Expecting 1")
   }

   ballot_ID := args[0]

   queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"ballot\",\"ballot_ID\":\"%s\"}}", ballot_ID)

   queryResults, err := getQueryResultForQueryString(stub, queryString)
   if err != nil {
	   return shim.Error(err.Error())
   }

   return shim.Success(queryResults)
}

func (t *SmartContract) addCampaignCandidates(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect Number of Aruments. Expecting 2")
	}

	fmt.Println("Adding new Candidiate into a Campaign")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	

	campaign_ID := args[0]
	user_ID := args[1]


	// ===== Create User Object and Marshal to JSON

	objectType := "campaigncandidates"
	campaigncandidates := &CampaignCandidates{objectType, campaign_ID, user_ID}
	campaigncandidatesJSONasBytes, err := json.Marshal(campaigncandidates)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save User to State

	err = stub.PutState(campaign_ID, campaigncandidatesJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Candidate to a Campaign")
	return shim.Success(nil)
}

func (t *SmartContract) queryCampaignCandidates(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
	   return shim.Error("Incorrect number of arguments. Expecting 1")
   }

   campaign_ID := args[0]

   
   queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"campaigncandidates\",\"campaign_ID\":\"%s\"}}",  campaign_ID)

   queryResults, err := getQueryResultForQueryString(stub, queryString)
   if err != nil {
	   return shim.Error(err.Error())
   }

   return shim.Success(queryResults)
}

func (t *SmartContract) queryAllCampaignCandidates(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	startKey := "campaign0"
	endKey := "campaign999999999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func (t *SmartContract) addCampaignResults(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect Number of Arguments. Expecting 4")
	}

	fmt.Println("Adding new Campaign Results")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	
	campaign_ID := args[0]
	candidate_ID := args[1]
	ballot_ID := args[2]
	timestamp, err1 := time.Parse(time.RFC3339, args[3])
	if err1 != nil {
		return shim.Error(err.Error())
	}

	// ===== Create Campaign Results  Object and Marshal to JSON

	objectType := "campaignresults"
	campaignresults := &CampaignResults{objectType, campaign_ID, candidate_ID, ballot_ID, timestamp}
	campaignresultsJSONasBytes, err := json.Marshal(campaignresults)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save Results to State

	err = stub.PutState(campaign_ID, campaignresultsJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Results")
	return shim.Success(nil)
}

func (t *SmartContract) queryCampaignResults(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
	   return shim.Error("Incorrect number of arguments. Expecting 1")
   }

   campaign_ID := args[0]

   queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"campaignresults\",\"campaign_ID\":\"%s\"}}", campaign_ID)

   queryResults, err := getQueryResultForQueryString(stub, queryString)
   if err != nil {
	   return shim.Error(err.Error())
   }

   return shim.Success(queryResults)
}



// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//Main Function starts up the Chaincode
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Smart Contract could not be run. Error Occured: %s", err)
	} else {
		fmt.Println("Smart Contract successfully Initiated")
	}
}

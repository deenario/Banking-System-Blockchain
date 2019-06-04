package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"strings"
	//"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)
//SmartContract is the data structure which represents this contract and on which  various contract lifecycle functions are attached
type SmartContract struct{
}

//Model for Users
type User struct {
	ObjectType			string `json:"Type"`
	Name				string `json:"name"`
	Address 			string `json:"address"`
	Email				string `json:"email"`
	Password			string `json:"password"`
	User_type				string `json:"user_type"`
	Biography			string `json:"biography"`
}

type Transactions struct {
	ObjectType			string `json:"Type"`
	Transaction_ID 		string `json:"transaction_ID"`
	To 					string `json:"to"`
	From				string `json:"from"`
	Amount				string `json:"amount"`
	Comment  			string `json:"comment"`
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
	if function == "addTransaction" {
		return t.addTransaction(stub,args)
	}
	if function == "queryTransactions" {
		return t.queryTransactions(stub,args)
	}
	// if function == "queryBusinessTransactions" {
	// 	return t.queryBusinessTransactions(stub,args)
	// }
	// if function == "queryIndividualTransactions" {
	// 	return t.queryIndividualTransactions(stub,args)
	// }

	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Invoke did not find specified function " + function)
}

func (t *SmartContract) addUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 6 {
		return shim.Error("Incorrect Number of Aruments. Expecting 6")
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
	if len(args[5]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}

	

	name := args[0]
	address := args[1]
	email := args[2]
	password := args[3]
	user_type := args[4]
	biography := args[5]

	// ======Check if User Already exists

	userAsBytes, err := stub.GetState(email)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if userAsBytes != nil {
		return shim.Error("The Inserted User already Exists: " + email)
	}

	// ===== Create User Object and Marshal to JSON

	objectType := "user"
	user := &User{objectType, name, address, email, password, user_type, biography}
	userJSONasBytes, err := json.Marshal(user)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save User to State

	err = stub.PutState(email, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved User")
	return shim.Success(nil)
}

func (t *SmartContract) queryUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {

 	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	email := args[0]
	password := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"user\",\"email\":\"%s\",\"password\":\"%s\"}}", email, password)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addTransaction(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect Number of Aruments. Expecting 5")
	}

	fmt.Println("Adding new Transaction")

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

	transaction_ID := args[0]
	to := args[1]
	from := args[2]
	amount := args[3]
	comment := args[4]

	// ======Check if Transaction Already exists

	transactionsAsBytes, err := stub.GetState(transaction_ID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if transactionsAsBytes != nil {
		return shim.Error("The Inserted transaction already Exists: " + transaction_ID)
	}

	// ===== Create transactions Object and Marshal to JSON

	objectType := "transactions"
	transactions := &Transactions{objectType, transaction_ID, to, from, amount, comment}
	transactionsJSONasBytes, err := json.Marshal(transactions)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save transactions to State

	err = stub.PutState(transaction_ID, transactionsJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved transaction")
	return shim.Success(nil)
}

func (t *SmartContract) queryTransactions(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	from := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"transactions\",\"from\":\"%s\"}}", from)

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

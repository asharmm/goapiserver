package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1254dv",
		Username:  "alex",
	},
	"john": {
		AuthToken: "78dEa",
		Username:  "john",
	},
	"emma": {
		AuthToken: "pqr12",
		Username:  "emma",
	},
	"sarah": {
		AuthToken: "abc45",
		Username:  "sarah",
	},
	"michael": {
		AuthToken: "xyz789",
		Username:  "michael",
	},
	"jane": {
		AuthToken: "def23",
		Username:  "jane",
	},
	

}


var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    1000,
		Username: "alex",
	},
	"john": {
		Coins:    500,
		Username: "john",
	},
	"emma": {
		Coins:    750,
		Username: "emma",
	},
	"sarah": {
		Coins:    300,
		Username: "sarah",
	},
	"michael": {
		Coins:    1200,
		Username: "michael",
	},
	"jane": {
		Coins:    600,
		Username: "jane",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}


func (d *mockDB) GetUserCoins(username string) *CoinDetails{

	time.Sleep(time.Second*1)


	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB ) SetupDatabase() error{
	return nil
}
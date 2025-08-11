package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "Alex",
	},

	"Jake": {
		AuthToken: "455EHB",
		Username:  "Jake",
	},

	"Qoli": {
		AuthToken: "556BNH",
		Username:  "Qoli",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "Alex",
	},
	"Jake": {
		Coins:    200,
		Username: "Jake",
	},
	"Qoli": {
		Coins:    230,
		Username: "Qoli",
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

func (d *mockDB) GetUserCoins(username string) *CoinDetails {

	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

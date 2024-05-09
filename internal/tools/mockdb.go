package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetail = map[string]LoginDetails{
	"will": {
		AuthToken: "123ABC",
		Username:  "will",
	},
	"matt": {
		AuthToken: "456DEF",
		Username:  "matt",
	},
	"brad": {
		AuthToken: "789GHI",
		Username:  "matt",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"will": {
		Coins: 100,
		Username: "will",
	},
	"matt": {
		Coins: 200,
		Username: "matt",
	},
	"brad": {
		Coins: 300,
		Username: "brad",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate db call
	time.Sleep(time.Second * 1)

	var loginDetails = LoginDetails{}
	loginDetails, ok := mockLoginDetail[username]
	if !ok {
		return nil
	}

	return &loginDetails
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate db call
	time.Sleep(time.Second * 1)

	var coinDetails = CoinDetails{}
	coinDetails, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &coinDetails
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

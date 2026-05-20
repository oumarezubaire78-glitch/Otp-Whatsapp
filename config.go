package main

var Config = struct {
	OwnerNumber   string
	BotName       string
	OTPChannelIDs []string
	OTPApiURLs    []string
	Interval      int
}{
	OwnerNumber: "923027665767",
	BotName:     "Ali OTP Monitor",
	OTPChannelIDs: []string{
		"120363406197025409@newsletter",
	},
	OTPApiURLs: []string{
		"http://147.135.212.197/crapi/st/viewstats-SFBTSUpBUzRSlGpkY3KWZESGjHdThXJTQVVmRXh_lHlplndSYWuBRA==
		sms",
		"http://kami-api-production.up.railway.app/api/np?type=sms",
		"http://kami-api1-production.up.railway.app/api/hs?type=sms",
		"http://kami-api1-production.up.railway.app/api/msi?type=sms",
	},
	Interval: 10,
}
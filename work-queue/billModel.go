package main

import (
	"encoding/json"
	"log"
)

type billData struct {
	Msisdn string
	PrincipalAmount string
	Fee string
	Channel string
	Type string
}

func toJson(msisdn string, principalAmount string, fee string, channel string, billerType string) []byte {
	log.Printf("Coverting object payload to json...")
	billData := &billData {
		Msisdn: msisdn,
		PrincipalAmount: principalAmount,
		Fee: fee,
		Channel: channel,
		Type: billerType,
	}

	res, err := json.Marshal(billData)
	FailOnError(err, "Failed to publish a message")

	return res
}

func fromJson(payload []byte) billData {
	log.Printf("Coverting binary payload to object...")

	paidBillData := billData{}
	err := json.Unmarshal(payload, &paidBillData)

	FailOnError(err, "Unable to unmarshall binary raw data from Queue")

	return paidBillData
}
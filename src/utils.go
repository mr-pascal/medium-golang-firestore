package main

import (
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func JsonToCustomer(input map[string]interface{}) (result *Customer) {
	dbByte, _ := json.Marshal(input)

	if err := json.Unmarshal(dbByte, &result); err != nil {
		log.Println(err)
		return nil
	}

	return
}

func CustomerToJson(input *CustomerUpdate) (result map[string]interface{}, err error) {
	b, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := json.Unmarshal(b, &result); err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func CustomerToFirestoreUpdate(input *CustomerUpdate) (result []firestore.Update, err error) {

	var r map[string]interface{}

	if r, err = CustomerToJson(input); err != nil {
		log.Println(err)
		return nil, err
	}

	for k, v := range r {
		result = append(result, firestore.Update{
			Path:  k,
			Value: v,
		})
	}

	return
}

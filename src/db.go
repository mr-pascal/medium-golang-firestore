package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

// Customer collection name
var collection string = "customer"

// Creates a new GCP Firestore client
func CreateClient(ctx context.Context, projectID string) (client *firestore.Client, err error) {
	// Output if the Firestore emulator is being used
	if value := GetEnv("FIRESTORE_EMULATOR_HOST", ""); value != "" {
		log.Printf("Using Firestore Emulator: '%s'", value)
	}

	// Create Firestore client
	client, err = firestore.NewClient(ctx, projectID)
	return
}

// Creates a new customer entity in GCP Firestore
// Additionaly checks if unique customer ID already exists
func CreateCustomer(c *firestore.Client, ctx context.Context, cus *Customer) (err error) {

	// Check if there is already a customer entry existing
	_, err = c.Collection(collection).Doc(cus.CustomerId).Get(ctx)
	if err == nil {
		// -> Something was found, so we break since we don't want to override it
		err = fmt.Errorf("Customer Creation: can't create customer with ID '%s' since it already exists", cus.CustomerId)
		return
	}

	// Create a new customer entry
	_, err = c.Collection(collection).Doc(cus.CustomerId).Set(ctx, cus)
	return
}

// Returns a customer from GCP Firestore
func GetCustomer(c *firestore.Client, ctx context.Context, customerId string) (cus *Customer, err error) {
	doc, err := c.Collection(collection).Doc(customerId).Get(ctx)

	if err != nil {
		if !doc.Exists() {
			// indicates a "Not Found" error
			// Alternatively its possible to check the error code of 'err'
			// -> status.Code(err) == codes.NotFound
			// You might add special handling here
			return nil, err
		}
		return nil, err
	}

	// Converts the JSON from GCP Firestore to our Customer object
	cus = JsonToCustomer(doc.Data())
	return
}

// Updates a customer entry via the Firestore document "Set" method
func UpdateCustomerViaSet(c *firestore.Client, ctx context.Context, customerId string, cus *CustomerUpdate) (err error) {
	res, err := CustomerToJson(cus)
	if err != nil {
		return
	}

	_, err = c.Collection(collection).Doc(customerId).Set(ctx, res, firestore.MergeAll)
	return
}

// Updates a customer entry via the Firestore document "Update" method
func UpdateCustomerViaUpdate(c *firestore.Client, ctx context.Context, customerId string, cus *CustomerUpdate) (err error) {
	res, err := CustomerToFirestoreUpdate(cus)
	if err != nil {
		return
	}
	_, err = c.Collection(collection).Doc(customerId).Update(ctx, res)
	return
}

// Delete a customer document by the provided customerId.
// If the document doesn't exist, it does nothing and returns no error.
func DeleteCustomer(c *firestore.Client, ctx context.Context, customerId string) (err error) {
	_, err = c.Collection(collection).Doc(customerId).Delete(ctx)
	return
}

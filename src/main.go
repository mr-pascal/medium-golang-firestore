package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

var gcpProjectId string = GetEnv("GCP_PROJECT_ID", "my-demo-project")

/////////////////////////////////////
///// Firestore Wrapper Methods /////
/////////////////////////////////////
func createCustomer(c *firestore.Client, ctx context.Context, customer *Customer, customerId string) {
	log.Printf("(ID: '%s'): Create customer ...", customerId)
	if err := CreateCustomer(c, ctx, customer); err != nil {
		log.Println(err)
	} else {
		log.Println("Customer created.")
	}
	log.Println()
}

func getCustomer(c *firestore.Client, ctx context.Context, customerId string) {
	log.Printf("(ID: '%s'): Get customer ...", customerId)
	if retrievedCustomer, err := GetCustomer(c, ctx, customerId); err != nil {
		log.Println(err)
	} else {
		log.Println(retrievedCustomer)
	}
	log.Println()
}

func updateCustomerViaSet(c *firestore.Client, ctx context.Context, customerId string, customerUpdate *CustomerUpdate) {
	log.Printf("(ID: '%s'): Update customer via 'Set' ...", customerId)

	if err := UpdateCustomerViaSet(c, ctx, customerId, customerUpdate); err != nil {
		log.Println(err)
	} else {
		log.Printf("Updated customer.")
	}
	log.Println()
}

func updateCustomerViaUpdate(c *firestore.Client, ctx context.Context, customerId string, customerUpdate *CustomerUpdate) {
	log.Printf("(ID: '%s'): Update customer via 'Update' ...", customerId)

	if err := UpdateCustomerViaUpdate(c, ctx, customerId, customerUpdate); err != nil {
		log.Println(err)
	} else {
		log.Printf("Updated customer.")
	}
	log.Println()
}

func deleteCustomer(c *firestore.Client, ctx context.Context, customerId string) {
	log.Printf("(ID: '%s'): Delete customer ...", customerId)

	if err := DeleteCustomer(c, ctx, customerId); err != nil {
		log.Println(err)
	} else {
		log.Printf("Deleted customer with ID '%s'", customerId)
	}
	log.Println()
}

/////////////////////////////////////
/////////////////////////////////////
/////////////////////////////////////

func main() {
	log.Println("Starting...")

	ctx := context.Background()

	log.Println("Instantiate GCP Firestore client")
	client, err := CreateClient(ctx, gcpProjectId)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close connection after we are done
	defer client.Close()

	validCustomerId := "1"
	invalidCustomerId := "5"
	invalidCustomerId2 := "6"

	newCustomer := Customer{
		CustomerId: validCustomerId,
		Enabled:    true,
		Name:       "Jason",
	}

	b := bool(false)
	customerUpdate := CustomerUpdate{
		Enabled: &b,
	}

	log.Println()
	log.Println("--------")

	/////////

	log.Println("Initial create customer")
	createCustomer(client, ctx, &newCustomer, validCustomerId)

	log.Println("Try to create the same customer again")
	createCustomer(client, ctx, &newCustomer, validCustomerId)

	log.Println("Get created customer")
	getCustomer(client, ctx, validCustomerId)

	log.Println("Update customer")
	updateCustomerViaSet(client, ctx, validCustomerId, &customerUpdate)

	log.Println("Get updated customer")
	getCustomer(client, ctx, validCustomerId)

	log.Println("Delete created customer")
	deleteCustomer(client, ctx, validCustomerId)

	log.Println("--------")
	log.Println()

	/////////

	log.Println("Update not existing customer via 'Set'")
	updateCustomerViaSet(client, ctx, invalidCustomerId2, &customerUpdate)

	log.Println("Get customer that was implicitly created via 'Set'")
	getCustomer(client, ctx, invalidCustomerId2)

	log.Println("Delete implicitly created customer")
	deleteCustomer(client, ctx, invalidCustomerId2)

	log.Println("--------")
	log.Println()

	/////////

	log.Println("Update not existing customer via 'Update'")
	updateCustomerViaUpdate(client, ctx, invalidCustomerId, &customerUpdate)

	log.Println("Delete not existing customer")
	deleteCustomer(client, ctx, invalidCustomerId)

	log.Println("Get not existing customer")
	getCustomer(client, ctx, invalidCustomerId)

	log.Println("--------")
	log.Println()

	/////////

	log.Println("Shutting down...")

}

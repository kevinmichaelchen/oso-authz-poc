package main

import (
	"fmt"
	oso "github.com/osohq/go-oso-cloud"
	"os"
)

func main() {
	apiKey := os.Getenv("OSO_AUTH_KEY")
	if apiKey == "" {
		panic("Missing required environment variable OSO_AUTH_KEY")
	}

	osoClient := oso.NewClient("https://cloud.osohq.com", apiKey)

	actor := oso.Instance{
		Type: "owner",
		ID:   "1",
	}

	resource := oso.Instance{
		Type: "pet",
		ID:   "2",
	}

	sql, err := osoClient.AuthorizeLocal(actor, "walk", resource)
	if err != nil {
		panic(err)
	}

	fmt.Println(sql)
}

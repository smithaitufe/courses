package schema

import (
	"io/ioutil"
	"log"
)

func GetSchema() string {
	schemaFile, err := ioutil.ReadFile("schema/schema.graphql")

	if err != nil {
		log.Fatal(err)
	}

	return string(schemaFile)
}

package main

import (
	"log"
	"os"

	"github.com/vice-registry/vice-api/restapi"
	"github.com/vice-registry/vice-api/restapi/operations"

	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

//
// This is a copy of the file cmd/vice-server/main.go - it is save to edit this file.
//

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewViceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "ViCE Image Registry API"
	parser.LongDescription = "A RESTful API server for the ViCE Image Registry"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

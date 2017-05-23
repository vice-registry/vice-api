# ViCE-API - RESTful Service Endpoint of ViCE Image Registry

This is the RESTful service endpoint of the image registry for the ViCE project.

## Development

The git repository is meant to be used inside a GOPATH environment. 
If go is installed, the following steps will guide to a running development environment:

```
cd $YOUR_VICE-API_WORKSPACE
export GOPATH=$(pwd)
git clone git@omi-gitlab.e-technik.uni-ulm.de:vice/vice-api.git src/omi-gitlab.e-technik.uni-ulm.de/vice/vice-api
(cd src/omi-gitlab.e-technik.uni-ulm.de/vice/vice-api;  go get -u -f ./...)
```

### Compile the swagger API definition
_ Make sure to have the git repository set up in a GOPATH environment._

If you change the API definition in `swagger.yaml`, use the `./swagger-gen` script to (re-)compile the go files.

### Compile vice-api to binary
_ Make sure to have the git repository set up in a GOPATH environment._

```
cd $YOUR_VICE-API_WORKSPACE
export GOPATH=$(pwd)
go build src/omi-gitlab.e-technik.uni-ulm.de/vice/vice-api
```

The last command will create the binary `vice-api`.
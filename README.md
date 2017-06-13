# ViCE-API - RESTful Service of ViCE Registry [![Build Status](https://travis-ci.org/vice-registry/vice-api.svg?branch=master)](https://travis-ci.org/vice-registry/vice-api)

This is the RESTful service endpoint of the image registry for the ViCE project.

## Development

The git repository is meant to be used inside a GOPATH environment. 
If go is installed, the following steps will guide to a running development environment.

```
cd $YOUR_VICE-API_WORKSPACE

export GOPATH=$(pwd)
# clone the src repo
git clone git@omi-gitlab.e-technik.uni-ulm.de:vice/vice-api.git src/github.com/vice-registry/vice-api
# get dependencies
(cd src/github.com/vice-registry/vice-api;  go get -u -f ./...)
```

### Compile the swagger API definition
*Make sure to have the git repository set up in a GOPATH environment.*

The API is specified by [swagger](http://swagger.io/) and compiled into go code with [go-swagger](https://goswagger.io/).
Please note: this is not the official swagger-codegen tool for go!
Make sure to install the go-swagger binary when you want to generate code from the swagger definition.

If you change the API definition in `swagger.yaml`, use the `./swagger-gen` script to (re-)compile the go files.

### Compile vice-api to binary
*Make sure to have the git repository set up in a GOPATH environment.*

```
cd $YOUR_VICE-API_WORKSPACE

export GOPATH=$(pwd)
go install github.com/vice-registry/vice-api
```

The last command will create the binary `bin/vice-api`, which provides the REST API.

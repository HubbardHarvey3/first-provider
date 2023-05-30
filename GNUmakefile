default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m


#build:
#	go build . -o terraform-provider-swapi-provider
#	mv terraform-provider-swapi-provider ~/.terraform.d/plugins/example.com/hharvey/swapi-provider/0.0.1/linux_amd64

build-work:
	go install .
	mv /usr/local/go/bin/swapi-provider /usr/local/go/bin/terraform-provider-swapi-provider

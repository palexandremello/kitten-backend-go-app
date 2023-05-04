test: 
	gotest -v ./app/...

test_coverage: 
	gotest ./app/... -coverprofile coverage.out

show_coverage: 
	go tool cover -html=coverage.out
    
test: 
	gotest -v ./app/...

test_coverage: 
	gotest -v -coverprofile cover.out ./app/...

show_coverage: 
	go tool cover -html=cover.out
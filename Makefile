test: 
	go test -v ./app/...

test_coverage: 
	go test ./app/... -coverprofile coverage.out

show_coverage: 
	go tool cover -html=coverage.out
    

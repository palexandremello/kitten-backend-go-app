test: 
	gotestsum --format testname -- -coverprofile=coverage.out ./app/...

show_coverage: 
	go tool cover -html=coverage.out
    

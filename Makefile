test: 
	gotestsum -- -coverprofile=coverage.out ./app/...

show_coverage: 
	go tool cover -html=coverage.out
    

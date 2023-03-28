# MicroService-SoundTrackDetaction

## About
A university assignment about producing a series of microservices to provide song track detection feature via using the audd.io API
For more infomation, please find the assignment pdf in the repository.

# Installation
## The Go Version
The latest Go version on Linux OS enabled machine. I do not recommend running the microservices in WSL.
## Audd.io API Key
Register an trial account on “https://audd.io” to obtain your API key before running the microservices.
## Inserting the API Key
Open the go file at `addison\search\service\service.go` and modify the code section:

`
const (
	api_url   = "https://api.audd.io/"
	api_token = "" // insert api key here
	
)
`

insert the `API_KEY` string with your own API key.


# Running the Microservices
In a terminal, change directory to the microservice folder:
`cd addison`

then, change the directory to the microservice folder in order:
1. `cd track` to the track microservice folder
2. `cd search` to the search microservice folder
3. `cd cooltown` to the cooltown microservice folder

and run the following command after cd into each the microservice folder:
a. `rm-f go.mod`
b. `go mod init track` <- change the target to `search` or `cooltown` for the other microservices
c. `go mod tidy`
d. `go run main.go`
- repeat the above steps for the other microservices `search` and `cooltown`

# Note
Make sure localhost's port 3000, 3001, 3002 are not occupied by other applications.

## MacOS
When running the search microservice script on MacOS, 
replace the keyword 'base64' with 'gbase64' in the script file.

## WSL
Cannot run the microservices in WSL.

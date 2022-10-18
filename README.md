# Problem Description
Pretend we have a tool that reports information from a machine to our service in
JSON format. We wish to capture this information in a server for easy viewing.

For this project we want to create a RESTful API that accepts the JSON 
payloads <sup>1</sup> and stores it in a datastore<sup>2</sup>. The RESTful API
should  also have a GET request that returns all entries in the store.

The endpoints, project structure, API documentation and library options are up
to you. The project should be unit tested to prove that your project works.

# Project Layout


```bash
├── README.md
├── utils
│   ├── commonUtility.go 
├── cmd
│   ├── main.go       
├── pkg
│   ├── handlers
│   │   └── systemHandler.go
        └── systemHandler_test.go    
│   ├── models
│   │   └── machine.go
│   │   └── stat.go
│   └── services
│       └── storageService.go
        └── storageService_test.go
```
README.md is a detailed description of the project.

cmd contains main packages and having a main.go file

pkg places most of project business logic and locate handler , models and service package.

# Getting Started

## Run
Follow the below step to run the server :

1. Clone the project `git clone https://github.com/Sarikap08/restProject.git`
2. Go to restProject folder
3. Run the command - go run ./cmd
4. To access the API 

    Import the postman `API.postman_collection.json` file for running the API call.

There is two end points : 

GET http://localhost:5051/machines - Fetches all the saved machine information from datastore 

POST http://localhost:5051/machines - Add the machine Information to data store  

# NOTE 

Because of other commitments I didn't integrated any database and using struct to store the machine information





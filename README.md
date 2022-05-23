# Getir project



## General

We’d like you to create a RESTful API with two endpoints.
1. One of them that fetches the data in the provided MongoDB collection and returns the results in the requested format.
2. Second endpoint is to create(POST) and fetch(GET) data from an in-memory database.

## Project info

The project has been developed in Go.
The internal database is simulated with a map[string]string
### Project Structure

**cmd:** contains the file, main.go, that handle the entry point of the service. It's responsible to initialize the structure used by the service, and start the server.

 **handlerMessage:** contains the file used to handle all the errors related to the service logic. Each errors is based on a specific case, and it's coupled with a dedicated error message, http status code and an identifier code.

**server:** contains the file to create the server and handle the client requests. 


### Implementation choices

The project is completely based on the logic of the interfaces that each file of a folder exposes to files of other folders. Each folder, as indicated in the structuring of the project, is responsible for an entity of the service and all its logic, requiring externally managed functionalities through interfaces.
This design choice, together with the use of data structures that in turn contained interfaces as properties, makes it easier to reuse the code and add future changes.

### Automated test

Unit tests were carried out on all functions, and for each one the possible inputs were considered that would allow to cover all the computation branches.
The use of mocks was made in tests where external functions were interrogated by the responsibility of the method under analysis.
All unit tests are present in their respective folders, while a complete function test is present in the integrationTest folder.

### Local start

go run ./cmd

### Docker start

docker build -t <image_name> .

docker run -p 8080:8080 -t <image_name> .
# ExoPlanet Service

This is a microservice for managing different types of exoplanets.

## Features

- Add an Exoplanet: Allows users to add new exoplanets with various properties.
- List Exoplanets: Retrieves a list of all available exoplanets.
- Get Exoplanet by ID: Retrieves information about a specific exoplanet by its unique ID.
- Update Exoplanet: Updates the details of an existing exoplanet.
- Delete Exoplanet: Removes an exoplanet from the catalog.
- Fuel Estimation: Provides an estimation of fuel cost for a trip to any particular exoplanet for a given crew capacity.

## Setup

### Prerequisites

- Go (1.18+)

### Installation

1. Clone the repository: `github.com/Kratos-28/ExoPlanet`
2. Navigate to the project directory: `cd ExoPlanet`
3. Build the project:  `go build`
4. Run the binary: `./ExoPlanet`
5. Alternatively, you can run the service using Docker: `docker build -t exoplanet-service .docker run -p 8080:8080 exoplanet-service`
   
## Usage

- Add an Exoplanet:
  - Endpoint: `POST /exoplanets`
  - Payload: JSON object with exoplanet properties.
- List Exoplanets:
  - Endpoint: `GET /exoplanets`
- Get Exoplanet by ID:
  - Endpoint: `GET /exoplanets/{id}`
- Update Exoplanet:
  - Endpoint: `PUT /exoplanets/{id}`
  - Payload: JSON object with updated exoplanet properties.
- Delete Exoplanet:
  - Endpoint: `DELETE /exoplanets/{id}`
- Fuel Estimation:
  - Endpoint: `GET /exoplanets/{id}/fuel?crew={crew_capacity}`


## Unit Tests

Unit tests are included to ensure the correctness of the service. You can run the tests using the following command:
`go test ./...`


## Sample Output

- Add An Exoplanet
<img width="835" alt="image" src="https://github.com/Kratos-28/ExoPlanet/assets/49074775/1a70124d-8825-4543-ab4a-ba06d33cce2f">

- GetExoPlanetById
<img width="845" alt="image" src="https://github.com/Kratos-28/ExoPlanet/assets/49074775/2f4df6a7-dec5-43ac-bb60-9a17ad41ed67">






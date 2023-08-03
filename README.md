# Admin Service API

This repository contains the documentation for the Admin Service API, which provides endpoints to manage doctors' and patients' data.

## Endpoints

| Endpoint                              | Description                                           |
| ------------------------------------- | ----------------------------------------------------- |
| GET /doctor/{doctorId}                | Get a doctor by ID                                    |
| PUT /doctor/{doctorId}                | Update a doctor by ID                                 |
| DELETE /doctor/{doctorId}             | Delete a doctor by ID                                 |
| POST /doctor                          | Create a new doctor                                   |
| GET /patient/{patientId}              | Get a patient by ID                                   |
| PUT /patient/{patientId}              | Update a patient by ID                                |
| DELETE /patient/{patientId}           | Delete a patient by ID                                |
| POST /patient                         | Create a new patient                                  |

## Description

- **GET /doctor/{doctorId}:** This endpoint allows you to retrieve information about a specific doctor by providing their ID.

- **PUT /doctor/{doctorId}:** Use this endpoint to update the details of a doctor by providing their ID and the new data.

- **DELETE /doctor/{doctorId}:** This endpoint lets you delete a doctor's record based on their ID.

- **POST /doctor:** This endpoint allows you to create a new doctor by providing relevant information in the request body.

- **GET /patient/{patientId}:** This endpoint allows you to retrieve information about a specific patient by providing their ID.

- **PUT /patient/{patientId}:** Use this endpoint to update the details of a patient by providing their ID and the new data.

- **DELETE /patient/{patientId}:** This endpoint lets you delete a patient's record based on their ID.

- **POST /patient:** This endpoint allows you to create a new patient by providing relevant information in the request body.

## Documentation

For more detailed information on each endpoint and their request/response formats, please refer to the [API documentation](https://github.com/Deveimer/admin-service/blob/development/api/openAPI.yaml).

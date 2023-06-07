# README #

This readme still working on progress

# Endpoint for Postman
[link](https://drive.google.com/file/d/1QUdkKN6SdeI_0CU_Dva1WR2_hahZdXwd/view?usp=sharing)

## API Contract
find the api documentation at the following postmant [link](https://app.swaggerhub.com/apis-docs/Lacutee/FitaTest/1.0.0)

## Setup
1. Install Go version 1.20
2. Use GoLand (recommended) or any IDE like Visual Studio Code
4. Download dependencies with command `go mod tidy and go vendor`
5. 

## Run

Use this command to run API app from root directory:

```shell
go run app/main.go
```

## Unit Tests

### Generate Mocks

To generate mock, run:

```shell
go generate ./...
```

### Run Unit Tests

To run unit tests:
```shell
go test ./...
```

---

## Coach Appointment Tech Spec

This README would normally document whatever steps are necessary to get your application up and running.

### Feature Description ###
there are three endpoint that could use for appointment proccess, which are:
1. Create Appointment
2. Approval Appointment
3. Reschedule Appointment

### Acceptance Criteria ###
1. Make an appointment & book the schedule
2. Validate an appointment based on coach availability
3. Coach can decline an appointment request or reschedule it
4. If user decline the rescheduling, then all ended (no need to provide another rescheduling)
5. Unit test

### Architecture and Design ###
this service using onion architecture, there are 5 layers 
from inner to outer which are entity, repository, use case,
controller, and request handler. the usage and responsibility of
each layer are follow:
1. **Entity**: this layer contains the domain model or entities
of the system. These are the core objects that 
represent the business concepts and rules.
2. **Repository**: This layer provides an interface for the 
application to access and manipulate the entities. 
It encapsulates the data access logic and provides
a way to abstract the database implementation details.
3. **Use case** : This layer contains the business logic 
or use cases of the system. It defines the operations 
that can be performed on the entities and orchestrates 
the interactions between the entities and the repository layer.
4. **Controller**: This layer handles the HTTP requests and
responses. It maps the incoming requests to the appropriate 
use case and returns the response to the client.
5. **Request handler**: This layer is responsible for handling 
the incoming HTTP requests and passing them on to 
the controller layer.

### Service State Diagram ###
this diagram will explain the state of appointment through 
th every process at system to achieve the output state.

![state diagram](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/main_state.png)

As explain at state diagram. there are 3 action could be performed
within the system which are create appointment, approval of appointment,
and reschedule appointment. So to cover the business logic of each
action. there i provide the activity diagram for them.

### Create Appointment Activity Diagram ###
The `create appointment` action is the entry point of this system. 
where user create the appointment request, and before the appointment 
create there are several requirment to pass which are:
1. appointment schedule should match the schedule of coach.
2. appointment schedule should not intercept with another user
appointment date with status booked, waiting for coach approval, 
or waiting reschedule approval from user.

Here is the complete diagram for this endpoint.

![create appointment activity](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/create-appointment/activity.png)

### Approval Of Appointment Activity Diagram ###
The `approval of appointment` action is used to approve any appointment
that on state of rescheduling and waiting for coach approval. at this
action, appointment could be rejected and approved. and before reschedule
the appointment, the schedule should fulfill the requirement same as
`create appointment` endpoint.

Here is the complete diagram for this endpoint.

![approval of appointment activity](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/approval-appointment/activity.png)

###  Reschedule Appointment Activity Diagram ###
The `reschedule appointment` action is used by coach to propose a new
schedule to customer. 

Here is the complete diagram for this endpoint.

![reschedule appointment activity](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/reschedule-appointment/activity.png)

###  Data Flow Diagram ###
there are 5 table that dedicated on this services which are `main_users`
`main_roles`, `appointments`, `time_zone_name` and `schedules`. For
complete explaination of database structure you can follow this ddf.


![data flow diagram](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/dfd.finish.png)


![data flow diagram](https://gitlab.com/Nacute/fita-be-test/-/blob/main/_diagrams/dfd.finish.png)

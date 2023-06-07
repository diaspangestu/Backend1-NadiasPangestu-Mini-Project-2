# README #

This readme still working on progress

## Previous Mini Project ##
here is my previous mini project called "Backend SQL Definition", if you want to see the previous project, [CLICK HERE](https://github.com/diaspangestu/BRI-Bootcamp-Internship-Mini-Project1)

## API Contract
find the api documentation at the following postmant [link](https://app.swaggerhub.com/apis-docs/Lacutee/FitaTest/1.0.0)

## Setup
1. Install Go version 1.20
2. Use GoLand (recommended) or any IDE like Visual Studio Code
4. Download dependencies with command `go mod tidy and go vendor`
5. Run MySql server (recommended using xampp or mamp)
6. Create database "bootcampbri_minpro1" or anything you want, you can setup the connection to database on db/gorm.go
7. Create table for database
    1. Table Actors
    ```shell
    CREATE TABLE actors (
   id BIGINT UNSIGNED NOT NULL,
   username VARCHAR(20) NOT NULL,
   password VARCHAR(30) NOT NULL,
   role_id INT UNSIGNED NOT NULL,
   is_verified ENUM('true', 'false'),
   is_active ENUM('true', 'false'),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id)
   ) ENGINE = InnoDB;
    ```
   
    2. Table Customers
   ```shell
    CREATE TABLE customer (
    id BIGINT UNSIGNED NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(50),
    avatar VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
    ) ENGINE = InnoDB;
    ```
   
    3. Table Roles
   ```shell
    id INT UNSIGNED NOT NULL,
    role_name VARCHAR(15),
    PRIMARY KEY (id)
    ) ENGINE = InnoDB;
   ```
   
    4. Table Register Approval
   ```shell
    CREATE TABLE register_approval (
    id BIGINT UNSIGNED NOT NULL,
    admin_id BIGINT UNSIGNED NOT NULL,
    super_admin_id BIGINT UNSIGNED NOT NULL,
    status VARCHAR(25),
    PRIMARY KEY (id)
    ) ENGINE = InnoDB;
    ```
   
    5. Alter Table Actors
   ```shell
    ALTER TABLE actors
    ADD FOREIGN KEY (role_id) REFERENCES role(id);
   ```
   
    6. Alter Table Register Approval
   ```shell
    ALTER TABLE register_approval
    ADD FOREIGN KEY (admin_id) REFERENCES actors(id),
    ADD FOREIGN KEY (super_admin_id) REFERENCES actors(id);
   ```
   
    7. Insert Data Role and Actors (superadmin)
   ```shell
    INSERT INTO role (id, role_name)
    VALUES (1, 'superadmin'),
    (2, 'admin');
    
    INSERT INTO actors (id, username, password, role_id, is_verified, is_active)
    VALUES (1, 'superadmin', 'password123', 1, true, true);
   ```
   
8. Run this command to run API app from root directory:
```shell
go run main.go
```

9. To use the API, please import this collection json to your Postman to test the API
-> [CLICK HERE TO DOWNLOAD COLLECTION](https://galactic-robot-267801.postman.co/workspace/Bootcamp~018ddfa3-b124-42a3-969c-9c65f80dd548/collection/16357233-9d0b805d-0360-404a-a241-b0a90cb1d69e?action=share&creator=16357233) <-

---

## Coach Appointment Tech Spec

This README would normally document whatever steps are necessary to get your application up and running.

### Feature Description ###
there are three endpoint that could use for appointment proccess, which are:

#### Superadmin ####
1. Login
2. Register customer data
3. Delete customer data
4. Read all customer data
5. Read all customer data with parameter (name, email) and pagination
6. Approve/Reject admin registration
7. Active/Deadactive admin
8. Read approval request from admin registration
9. Delete admin data
10. Read all admin data with parameter (username) and pagination

#### Admin ####
1. Login
2. Register admin data
3. Read admin data
4. Update admin data
5. Delete admin data
6. Register customer data
7. Delete customer data
8. Read all customer data
9. Read all admin data
10. Fetch from external API

#### Customer ####
1. Create customer data
2. Read customer data
3. Update customer data
4. Delete customer data

---

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

# HNGx Stage 2 Person Crud API

## Introduction

This is a mini rest api that provides  endpoint to perform crud on person resource

## Table of Contents
1. [Technology stack](#technology-stack)
2. [Application features](#application-features)
4. [Application Setup](#application-setup)
5. [Setup Environment Variable](#setup-environment-variable)
5. [Start Application](#start-application)
6. [Api Endpoints](#api-endpoints)
7. [View UML Diagram](#uml-diagram)
8. [View ER Diagram](#er-diagram)
9. [Run Test](#run-test)
10. [Swagger Documentation Link](#swagger-documentation-link)
6. [Author](#author)
7. [License](#license)

### Technology Stack
- [GO lang](https://www.php.net)
- [GO ORM](https://gorm.io/)
- [GIN Framework](https://gin-gonic.com/docs/)
- [Postgres Database](https://www.mysql.com)
  ##### Testing tools
    - [Testify](https://github.com/stretchr/testify)
    - [SQLite](https://www.sqlite.org/index.html)

### Application Features
* Ability to view list of person
* Ability to get single person using person id or name in path param
* Ability to create new person
* Ability to delete single person using person id or name in path param

### Application Setup
This instruction will get the project working on your local machine for DEVELOPMENT and TESTING purposes.
Make sure you have go install in your machine.
You can confirm go is availabe by running  `go version`
If that is sucessfull, then you have go installed, otherwise go to official documentation [Install Go](https://go.dev/doc/install)

Once you are sure you have installed Go
You need to be sure you have GOPATH setup by running `ech $GOPATH`
If you have that, you are good otherwise visit <br>
For [MAC Users](https://sourabhbajaj.com/mac-setup/Go/) <br>
For [Windows Users](https://www.freecodecamp.org/news/setting-up-go-programming-language-on-windows-f02c8c14e2f/) <br>
For [Linux Users](https://itslinuxfoss.com/set-gopath-environment-variable-ubuntu/)<br>

Once the GOPATH is set up for your device

```
    //Make Application Directory
     mkdir -p $GOPATH/src/github.com/projects/person-crud
    //cd into created directory
    cd $GOPATH/src/github.com/projects/person-crud
    //clone repository
    git clone https://github.com/harmlessprince/person_crud.git .
    // Install Go Packages
    go mod download
```

#### Setup Environment Variable
Default database is sqlite, you can use postgres, mysql.
You can modify the database connection to your choice
```env
APP_NAME=Go Application
APP_ENV=development
APP_DEBUG=true
APP_URL=localhost:3000

APP_PORT=3000

DB_CONNECTION=sqlite
DB_HOST=127.0.0.1
DB_PORT=3307
DB_DATABASE=crud_person
DB_USERNAME=root
DB_PASSWORD=password
```

#### Start Application
  ```
  make run-dev
  ```
or
  ```
    CompileDaemon -command="./person-crud"
  ```

----
#### API Endpoints
Production BASE URL
```
https://taofeeq-hng-stage-two.onrender.com
```

Development BASE URL
```
http://localhost:3000
```

**List Person**
----
Returns json data about all person by in the database

* **URL**

  /api

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   None

* **Data Params**

  None


* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
      ```
          {
            "data": [
              {
                "created_at": "2023-09-11T16:23:57.832758+01:00",
                "deleted_at": "2023-11-11T16:23:57.832758+01:00",
                "id": "1",
                "name": "Mark Essien",
                "updated_at": "2023-10-11T16:23:57.832758+01:00"
              }
            ],
            "message": "string",
            "success": true
          }
      ```

* **Error Response:**

    * **Code:** 400 <br/>
      **Content:**
      ```
      {
        "message": "string",
        "success": false
      }
      ```

* **Sample Call:**

  ```javascript
    var requestOptions = {
      method: 'GET',
      redirect: 'follow'
    };
    
    fetch("https://taofeeq-hng-stage-two.onrender.com/api", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
  ```
----

**Create Person**

----
Returns json data about all person by in the database

* **URL**

  /api

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   None

* **Data Params**

  None


* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
      ```
          {
              "data": {
                "created_at": "2023-09-11T16:23:57.832758+01:00",
                "deleted_at": "2023-11-11T16:23:57.832758+01:00",
                "id": "1",
                "name": "Mark Essien",
                "updated_at": "2023-10-11T16:23:57.832758+01:00"
              },
              "message": "string",
              "success": true
         }
      ```

* **Error Response:**

    * **Code:** 400 Bad Request
      **Content:**
      ```
      {
        "message": "string",
        "success": false
      }
      ```
  OR
* **Code:** 422 Unprocessable Entity
  **Content:**
     ```
     {
         "errors": [
           {
             "condition": "required",
             "error": "Name field is required",
             "key": "name"
           }
         ],
         "message": "string",
         "success": false
     }
     ```

* **Sample Call:**

  ```javascript
    var raw = "{\n    \"name\": \"James T\"\n}";

    var requestOptions = {
      method: 'POST',
      body: raw,
      redirect: 'follow'
    };
    
    fetch("https://taofeeq-hng-stage-two.onrender.com/api", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
  ```

----
**Show Person**
----
Returns json data about a single person by passing either the person id or person name.

* **URL**

  /:user_id

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `user_id=[srring]`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
      ```
          {
            "data": {
                "created_at": "2023-09-11T16:23:57.832758+01:00",
                "deleted_at": "2023-11-11T16:23:57.832758+01:00",
                "id": "1",
                "name": "Mark Essien",
                "updated_at": "2023-10-11T16:23:57.832758+01:00"
            },
            "message": "string",
            "success": true
          }
      ```


* **Error Response:**

    * **Code:** 400 <br/>
      **Content:**
      ```
      {
        "message": "string",
        "success": false
      }
      ```

* **Sample Call:**

  ```javascript
    var requestOptions = {
      method: 'GET',
      redirect: 'follow'
    };
    
    fetch("https://taofeeq-hng-stage-two.onrender.com/api/1", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
  ```
----
**Delete Person**
----
Delete a single person by passing either the person id or person name.

* **URL**

  /:user_id

* **Method:**

  `DELETE`

*  **URL Params**

   **Required:**

   `user_id=[srring]`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
      ```
          {
            "data": null,
            "message": "string",
            "success": true
          }
      ```

* **Error Response:**

    * **Code:** 400 <br/>
      **Content:**
      ```
      {
        "message": "string",
        "success": false
      }
      ```

* **Sample Call:**

  ```javascript
    var requestOptions = {
      method: 'DELETE',
      redirect: 'follow'
    };
    
    fetch("https://taofeeq-hng-stage-two.onrender.com/api/api/1", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
  ```
----
### UML Diagram
![](/uml_crud.png)

### ER Diagram
![](/er_crud.png)


### Run Test

```
make test
```

OR

```
gotest ./tests/... -v
```
### Swagger Documentation Link

```
https://taofeeq-hng-stage-two.onrender.com/docs/index.html
```

### Author
Name: Adewuyi Taofeeq <br>
Email: realolamilekan@gmail.com <br>
LinkenIn:  <a href="#license">Adewuyi Taofeeq Olamikean</a> <br>

### License
ISC

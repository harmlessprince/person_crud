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
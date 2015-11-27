Simple backend for chromecast torrent player.


[Frontend](https://github.com/bmartynov/gochromecast_frontend)

**Add download**
----
* **URL**
  /download/add

* **Method:**
    GET
  
*  **URL Params**

   **Required:**
   
   `uri=[string]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadAdd","result":{"download_id":"<md5 sum from uri>"},"status":0}`
 
 
 
* **Error Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadAdd","message":{"code":5,"message":"download already exists"},"status":1}`

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadAdd","status":1,"message":{"code":4,"message":"cannot create torrent client: <orig error>"}}`

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadAdd","status":1,"message":{"code":5,"message":"download already exists"}}`

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadAdd","status":1,"message":{"code":6,"message":"invalid uri (<uri>)"}}`

**Start Download**
----
* **URL**
  /download/start

* **Method:**
    GET
  
*  **URL Params**

   **Required:**
   
   `id=[string]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadStart","result":{},"status":0}`

* **Error Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadStart","message":{"code":2,"message":"requested download not found (%s)"},"status":1}`

**Stop Download**
----
* **URL**
  /download/stop

* **Method:**
    GET
  
*  **URL Params**

   **Required:**
   
   `id=[string]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadStop","result":{},"status":0}`

* **Error Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadStop","message":{"code":2,"message":"requested download not found (%s)"},"status":1}`
    
**Remove Download**
----
* **URL**
  /download/remove

* **Method:**
    GET
  
*  **URL Params**

   **Required:**
   
   `id=[string]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadRemove","result":{},"status":0}`

* **Error Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadRemove","message":{"code":2,"message":"requested download not found (%s)"},"status":1}`
    
    
**Info all**
----
* **URL**
  /download/info

* **Method:**
    GET
  
*  **URL Params**

   **Required:**
   
   `id=[string]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadInfoAll","result":{},"status":0}`

* **Error Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadInfoAll","message":{"code":2,"message":"requested download not found (%s)"},"status":1}`

**Play**
----
* **URL**
  /download/play

* **Method:**
    GET
  
*  **URL Params**

   **Required:**
   
   `id=[string]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `<content>`

* **Error Response:**

  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadPlay","message":{"code":2,"message":"requested download not found (%s)"},"status":1}`
    
  * **Code:** 200 <br />
    **Content:** `{"method":"DownloadPlay","message":{"code":3,"message":"file not found (%s)"},"status":1}`
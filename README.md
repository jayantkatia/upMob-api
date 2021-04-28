# upcoming_mobiles_api (UpMob API)
<p align="center">UpMob API scraps <a href="https://www.91mobiles.com/upcoming-mobiles-in-india">91mobiles.com</a> to get devices information which are yet to be released in the Indian market and stores it in a postgres db<p>

## API
  <i>Fetch all upcoming devices</i> 

* <b>URL<b>

  <i>/devices</i>

* <b>Method<b>
  
  `GET`
  
* <b>URL Params<b>

    <i>None</i>

* <b>Success Response:<b>

  * <b>Code:<b> 200 <br />
    <b>Content:<b> `[
  {
    "device_name": "Xiaomi Redmi Note 9 Pro 5G",
    "expected": "Expected Launch:May 2021",
    "price": 17990,
    "img_url": "https://www.91-img.com/pictures/141053-v2-xiaomi-redmi-note-9-pro-5g-mobile-phone-medium-1.jpg?tr=q-60",
    "source_url": "https://www.91mobiles.com/xiaomi-redmi-note-9-pro-5g-price-in-india",
    "spec_score": 85
  }, ... ]`
 

## Get Started with the Installation 
1. Required Installations
    1. <a href="https://golang.org/doc/install">Install GoLang in your system</a>
    2. <a href="https://docs.docker.com/engine/install/">Install Docker in your system</a>
    3. <a href="https://github.com/golang-migrate/migrate/tree/master/cmd/migrate">Install Golang Migrate Tool</a> 
    4. Make sure you have ```make``` tool installed.
2. Navigate into the project directory
3. Run
    ```shell
       go mod download
       make postgres
       make createdb
       make migrateup 
    ```
    This sets up and runs your postgres container, creates db in it and migrates database. 
4. Run
    ```shell
        make server
    ```
    This runs main.go and you are good to Go :wink:


## Development and Contributing
Yes, please! Feel free to contribute, raise issues and recommend best practices.
<a href="https://github.com/jayantkatia/upcoming_mobiles_api/blob/main/Makefile"> Makefile</a> is your friend.

A few resources:
- [GoLang Documentation](https://golang.org/doc/)
- [Docker Documentation](https://docs.docker.com/get-started/overview/)

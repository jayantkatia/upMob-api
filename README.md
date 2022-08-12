# upcoming_mobiles_api (UpMob API)
<p align="center">UpMob API scraps <a href="https://www.91mobiles.com/upcoming-mobiles-in-india">91mobiles.com</a> to get devices information which are yet to be released in the Indian market and stores it in a postgres db<p>

## API
  <i>Fetch all upcoming devices</i> 

* <b>URL</b>

  <i>/devices/top100</i>

* <b>Method</b>
  
  `GET`
  
* <b>URL Params</b>

    <i>None</i>

* <b>Success Response:</b>

  * <b>Code:</b> 200 <br />
    <b>Content:</b> 
    <details>
    <summary>See Response</summary>
    <br>
``` 
[
  {
        "device_name": "OnePlus 8T Pro",
        "last_updated": "Updated on: Jul 26, 2021",
        "expected": "August 7, 2021 (Unofficial)",
        "price": 57999,
        "img_url": "https://www.91-img.com/pictures/139350-v1-oneplus-8t-pro-mobile-phone-large-1.jpg?tr=q-60",
        "source_url": "https://www.91mobiles.com/oneplus-8t-pro-price-in-india",
        "spec_score": 0,
        "ram": {
            "String": "8 GB",
            "Valid": true
        },
        "processor": {
            "String": "Qualcomm Snapdragon 865 Plus",
            "Valid": true
        },
        "front_camera": {
            "String": "16 MP",
            "Valid": true
        },
        "rear_camera": {
            "String": "64 MP + 48 MP + 8 MP + 5 MP",
            "Valid": true
        },
        "battery": {
            "String": "4850 mAh",
            "Valid": true
        },
        "display": {
            "String": "6.78 inches",
            "Valid": true
        },
        "operating_system": {
            "String": "Android v10 (Q)",
            "Valid": true
        },
        "custom_ui": {
            "String": "Oxygen OS",
            "Valid": true
        },
        "chipset": {
            "String": "Qualcomm Snapdragon 865 Plus",
            "Valid": true
        },
        "cpu": {
            "String": "Octa core (3.09 GHz, Single core, Kryo 585 + 2.42 GHz, Tri core, Kryo 585 + 1.8 GHz, Quad core, Kryo 585)",
            "Valid": true
        },
        "architecture": {
            "String": "64 bit",
            "Valid": true
        },
        "graphics": {
            "String": "Adreno 650",
            "Valid": true
        },
        "display_type": {
            "String": "Fluid AMOLED",
            "Valid": true
        },
        "screen_size": {
            "String": "6.78 inches (17.22 cm)",
            "Valid": true
        },
        "resolution": {
            "String": "1440 x 3168 pixels",
            "Valid": true
        },
        "pixel_density": {
            "String": "513 ppi",
            "Valid": true
        },
        "touchscreen": {
            "String": "Yes, Capacitive Touchscreen, Multi-touch",
            "Valid": true
        },
        "internal_memory": {
            "String": "256 GB",
            "Valid": true
        },
        "expandable_memory": {
            "String": "No",
            "Valid": true
        },
        "m_camera_setup": {
            "String": "Quad",
            "Valid": true
        },
        "m_resolution": {
            "String": "64 MP Primary Camera48 MP, Wide Angle, Ultra-Wide Angle Camera8 MP Telephoto Camera5 MP Camera",
            "Valid": true
        },
        "m_autofocus": {
            "String": "Yes, Phase Detection autofocus",
            "Valid": true
        },
        "m_ois": {
            "String": "Yes",
            "Valid": true
        },
        "m_sensors": {
            "String": "",
            "Valid": false
        },
        "m_flash": {
            "String": "Yes, LED Flash",
            "Valid": true
        },
        "m_image_resolution": {
            "String": "",
            "Valid": false
        },
        "m_settings": {
            "String": "Exposure compensation, ISO control",
            "Valid": true
        },
        "m_shooting_modes": {
            "String": "Continuos ShootingHigh Dynamic Range mode (HDR)",
            "Valid": true
        },
        "m_camera_features": {
            "String": "Digital ZoomAuto FlashFace detectionTouch to focus",
            "Valid": true
        },
        "m_video_recording": {
            "String": "",
            "Valid": false
        },
        "s_camera_setup": {
            "String": "Single",
            "Valid": true
        },
        "s_resolution": {
            "String": "16 MP Primary Camera",
            "Valid": true
        },
        "s_video_recording": {
            "String": "",
            "Valid": false
        },
        "capacity": {
            "String": "4850 mAh",
            "Valid": true
        },
        "removable_battery": {
            "String": "No",
            "Valid": true
        },
        "wireless_charging": {
            "String": "",
            "Valid": false
        },
        "quick_charging": {
            "String": "Yes, Fast",
            "Valid": true
        },
        "usb": {
            "String": "Yes",
            "Valid": true
        },
        "sim_slots": {
            "String": "Dual SIM, GSM+GSM",
            "Valid": true
        },
        "network_support": {
            "String": "5G supported by device (network not rolled-out in India), 4G (supports Indian bands), 3G, 2G",
            "Valid": true
        },
        "fingerprint_sensor": {
            "String": "Yes",
            "Valid": true
        },
        "other_sensors": {
            "String": "Light sensor, Proximity sensor, Accelerometer, Compass, Gyroscope",
            "Valid": true
        },
        "scrape_timestamp": "2021-07-25T21:33:31.655995Z"
    }, ... ]

```
</details>

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

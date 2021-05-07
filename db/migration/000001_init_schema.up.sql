CREATE TABLE "devices" (
  "device_name" varchar PRIMARY KEY,
  "last_updated" varchar NOT NULL,
  "expected" varchar NOT NULL,
  "price" int NOT NULL,
  "img_url" varchar NOT NULL,
  "source_url" varchar NOT NULL,
  "spec_score" int NOT NULL,

  "ram" varchar,
  "processor" varchar,
  "front_camera" varchar,
  "rear_camera" varchar,
  "battery" varchar,
  "display" varchar,
  "operating_system" varchar,
  
  "custom_ui" varchar, 
  "chipset" varchar,
  "cpu" varchar,
  "architecture" varchar,
  "graphics" varchar,

  "display_type" varchar,
  "screen_size" varchar,
  "resolution" varchar,
  "pixel_density" varchar,
  "touchscreen" varchar,

  "internal_memory" varchar,
  "expandable_memory" varchar,
  
  "m_camera_setup" varchar,
  "m_resolution" varchar,
  "m_autofocus"	varchar,
  "m_ois" varchar,
  "m_sensors"	varchar,
  "m_flash" varchar,
  "m_image_resolution" varchar,
  "m_settings" varchar,
  "m_shooting_modes" varchar,
  "m_camera_features" varchar,
  "m_video_recording" varchar,

  "s_camera_setup" varchar,
  "s_resolution" varchar,
  "s_video_recording" varchar,


  "capacity" varchar,
  "removable_battery" varchar,
  "wireless_charging" varchar,
  "quick_charging" varchar,
  "usb" varchar,

  "sim_slots" varchar,
  "network_support" varchar,

  "fingerprint_sensor" varchar,
  "other_sensors" varchar,
  "scrape_timestamp" timestamp NOT NULL
);


CREATE INDEX ON "devices" ("device_name");
CREATE INDEX ON "devices" ("last_updated");
CREATE INDEX ON "devices" ("scrape_timestamp");


-- name: InsertDevice :one
INSERT INTO devices (
  device_name,
  last_updated,
  expected,
  price,
  img_url,
  source_url,
  spec_score,
  ram,
  processor,
  front_camera,
  rear_camera,
  battery,
  display,
  operating_system,
  custom_ui, 
  chipset,
  cpu,
  architecture,
  graphics,
  display_type,
  screen_size,
  resolution,
  pixel_density,
  touchscreen,
  internal_memory,
  expandable_memory,
  m_camera_setup,
  m_resolution,
  m_autofocus,
  m_ois,
  m_sensors,
  m_flash,
  m_image_resolution,
  m_settings,
  m_shooting_modes,
  m_camera_features,
  m_video_recording,
  s_camera_setup,
  s_resolution,
  s_video_recording,
  capacity,
  removable_battery,
  wireless_charging,
  quick_charging,
  usb,
  sim_slots,
  network_support,
  fingerprint_sensor,
  other_sensors,
  scrape_timestamp
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50
) RETURNING *;

-- name: GetAllDevices :many
SELECT * FROM devices;

-- name: GetLastUpdatedDevice :one
SELECT last_updated FROM devices WHERE device_name = $1;

-- name: GetLastXDevices :many
SELECT * FROM devices ORDER BY scrape_timestamp DESC LIMIT $1; 

-- name: UpdateScrapeTimestamp :exec
UPDATE devices SET scrape_timestamp=$1 WHERE device_name = $1;

-- name: DeleteDevice :exec
DELETE FROM devices WHERE device_name = $1;

-- name: DeleteDevicesXDaysOld :exec
DELETE FROM devices WHERE scrape_timestamp > $1;
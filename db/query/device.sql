-- name: InsertDevice :one
INSERT INTO devices (
    device_name,
    expected,
    price,
    img_url,
    source_url,
    spec_score
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetDevices :many
SELECT * FROM devices;

-- name: DeleteDevices :exec
DELETE FROM devices;
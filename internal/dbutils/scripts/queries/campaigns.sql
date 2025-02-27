-- name: InsertCampaign :one
INSERT INTO campaigns (name, description, key_field, user_id, create_datetime)
VALUES (?, ?, ?, ?, now())
RETURNING *;

-- name: GetCampaignByName :one
SELECT * FROM campaigns WHERE name = ?;

-- name: GetCampaigns :many
SELECT * FROM campaigns WHERE user_id = ?;

-- name: GetCampaignsWithCharacterCount :many
SELECT campaigns.id, campaigns.name, campaigns.description, COUNT(characters.id) AS character_count
FROM campaigns
LEFT JOIN characters ON campaigns.id = characters.campaign_id
WHERE user_id = ?
GROUP BY campaigns.id;

-- name: GetCampaignById :one
SELECT * FROM campaigns WHERE id = ?;
-- name: InsertCampaign :one
INSERT INTO campaigns (name, key, user_id, create_datetime)
VALUES (?, ?, ?, datetime('now', 'localtime'))
RETURNING *;

-- name: GetCampaigns :many
SELECT * FROM campaigns WHERE user_id = ?;

-- name: GetCampaignsWithCharacterCount :many
SELECT campaigns.id, campaigns.name, campaigns.description, COUNT(characters.id) AS character_count
FROM campaigns
LEFT JOIN characters ON campaigns.id = characters.campaign_id
WHERE user_id = ?
GROUP BY campaigns.id;
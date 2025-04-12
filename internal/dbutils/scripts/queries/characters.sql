-- name: InsertCharacter :exec
INSERT INTO characters (name, player_name, xp_bonus, campaign_id, create_datetime, total_xp)
VALUES (?, ?, ?, ?, ?, now());

-- name: GetLatestCharacterByID :one
SELECT * FROM characters WHERE id = LAST_INSERT_ID();

-- name: GetCharactersForCampaign :many
SELECT * FROM characters
WHERE campaign_id = ?
  AND not dead;

-- name: GetDeadCharactersForCampaign :many
SELECT * FROM characters
WHERE campaign_id = ?
  AND dead=true;

-- name: GetCharacterById :one
SELECT * FROM characters
where id = ? LIMIT 1;

-- name: DeleteCharacter :exec
DELETE FROM characters where id = ?;

-- name: UpdateCharacter :exec
UPDATE characters 
SET name = ?, player_name = ?, xp_bonus = ?, total_xp = ?
WHERE id = ?;

-- name: GetAllCharacters :many
SELECT *
FROM characters;


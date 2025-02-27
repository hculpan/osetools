-- name: InsertCharacter :one
INSERT INTO characters (name, player_name, xp_bonus, campaign_id, create_datetime)
VALUES (?, ?, ?, ?, now())
RETURNING *;

-- name: GetCharactersForCampaign :many
SELECT * FROM characters
WHERE campaign_id = ?;

-- name: GetCharacterWithXp :one
SELECT characters.id, characters.name, characters.player_name, characters.campaign_id, characters.xp_bonus, SUM(xp_awards.xp_award_with_bonus) AS total_xp
FROM characters
LEFT JOIN xp_awards ON characters.id = xp_awards.character_id
WHERE characters.id = ?
GROUP BY characters.id LIMIT 1;

-- name: GetCharactersForCampaignWithXp :many
SELECT characters.id, characters.name, characters.player_name, characters.campaign_id, characters.xp_bonus, SUM(xp_awards.xp_award_with_bonus) AS total_xp
FROM characters
LEFT JOIN xp_awards ON characters.id = xp_awards.character_id
WHERE characters.campaign_id = ?
GROUP BY characters.id;


-- name: GetCharacterById :one
SELECT * FROM characters
where id = ? LIMIT 1;

-- name: DeleteCharacter :exec
DELETE FROM characters where id = ?;
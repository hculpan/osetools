-- name: GetXpBonusForCharacter :many
SELECT * FROM xp_bonus_reasons
WHERE character_id = ?;

-- name: InsertXpBonusForCharacter :one
INSERT INTO xp_bonus_reasons (xp_bonus, reason, character_id, create_datetime)
VALUES (?, ?, ?, datetime('now', 'localtime'))
RETURNING *;

-- name: DeleteXpBonusForCharacter :exec
DELETE FROM xp_bonus_reasons
WHERE character_id = ?;
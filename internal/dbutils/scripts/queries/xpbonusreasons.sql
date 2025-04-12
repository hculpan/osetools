-- name: GetXpBonusForCharacter :many
SELECT * FROM xp_bonus_reasons
WHERE character_id = ?;

-- name: InsertXpBonusForCharacter :exec
INSERT INTO xp_bonus_reasons (xp_bonus, reason, character_id, create_datetime)
VALUES (?, ?, ?, now());

-- name: GetLatestXpBonusReasonByID :one
SELECT * FROM xp_bonus_reasons WHERE id = LAST_INSERT_ID();

-- name: DeleteXpBonusForCharacter :exec
DELETE FROM xp_bonus_reasons
WHERE character_id = ?;
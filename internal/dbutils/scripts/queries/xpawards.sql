-- name: GetXpAwardsForCharacter :many
SELECT * FROM xp_awards
WHERE character_id = ?
ORDER BY id desc;

-- name: InsertXpAward :exec
INSERT INTO xp_awards (xp_award, xp_award_with_bonus, reason, character_id, create_datetime)
VALUES (?, ?, ?, ?, now());

-- name: GetLatestXpAwardByID :one
SELECT * FROM xp_awards WHERE id = LAST_INSERT_ID();

-- name: DeleteXpAwardsForCharacter :exec
DELETE FROM xp_awards WHERE character_id = ?;


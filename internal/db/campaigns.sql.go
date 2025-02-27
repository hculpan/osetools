// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: campaigns.sql

package db

import (
	"context"
	"database/sql"
)

const getCampaignById = `-- name: GetCampaignById :one
SELECT id, name, description, key_field, user_id, create_datetime FROM campaigns WHERE id = ?
`

func (q *Queries) GetCampaignById(ctx context.Context, id interface{}) (Campaign, error) {
	row := q.db.QueryRowContext(ctx, getCampaignById, id)
	var i Campaign
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.KeyField,
		&i.UserID,
		&i.CreateDatetime,
	)
	return i, err
}

const getCampaignByName = `-- name: GetCampaignByName :one
SELECT id, name, description, key_field, user_id, create_datetime FROM campaigns WHERE name = ?
`

func (q *Queries) GetCampaignByName(ctx context.Context, name string) (Campaign, error) {
	row := q.db.QueryRowContext(ctx, getCampaignByName, name)
	var i Campaign
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.KeyField,
		&i.UserID,
		&i.CreateDatetime,
	)
	return i, err
}

const getCampaigns = `-- name: GetCampaigns :many
SELECT id, name, description, key_field, user_id, create_datetime FROM campaigns WHERE user_id = ?
`

func (q *Queries) GetCampaigns(ctx context.Context, userID int64) ([]Campaign, error) {
	rows, err := q.db.QueryContext(ctx, getCampaigns, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Campaign
	for rows.Next() {
		var i Campaign
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.KeyField,
			&i.UserID,
			&i.CreateDatetime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCampaignsWithCharacterCount = `-- name: GetCampaignsWithCharacterCount :many
SELECT campaigns.id, campaigns.name, campaigns.description, COUNT(characters.id) AS character_count
FROM campaigns
LEFT JOIN characters ON campaigns.id = characters.campaign_id
WHERE user_id = ?
GROUP BY campaigns.id
`

type GetCampaignsWithCharacterCountRow struct {
	ID             interface{}
	Name           string
	Description    sql.NullString
	CharacterCount int64
}

func (q *Queries) GetCampaignsWithCharacterCount(ctx context.Context, userID int64) ([]GetCampaignsWithCharacterCountRow, error) {
	rows, err := q.db.QueryContext(ctx, getCampaignsWithCharacterCount, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCampaignsWithCharacterCountRow
	for rows.Next() {
		var i GetCampaignsWithCharacterCountRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CharacterCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertCampaign = `-- name: InsertCampaign :one
INSERT INTO campaigns (name, description, key_field, user_id, create_datetime)
VALUES (?, ?, ?, ?, now())
RETURNING id, name, description, key_field, user_id, create_datetime
`

type InsertCampaignParams struct {
	Name        string
	Description sql.NullString
	KeyField    string
	UserID      int64
}

func (q *Queries) InsertCampaign(ctx context.Context, arg InsertCampaignParams) (Campaign, error) {
	row := q.db.QueryRowContext(ctx, insertCampaign,
		arg.Name,
		arg.Description,
		arg.KeyField,
		arg.UserID,
	)
	var i Campaign
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.KeyField,
		&i.UserID,
		&i.CreateDatetime,
	)
	return i, err
}

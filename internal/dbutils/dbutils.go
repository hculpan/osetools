package dbutils

import (
	"context"

	"github.com/hculpan/osetools/internal/db"
)

func UpdateCharacter(ctx context.Context, ch db.Character) error {
	update := db.UpdateCharacterParams{
		ID:         ch.ID,
		PlayerName: ch.PlayerName,
		Name:       ch.Name,
		XpBonus:    ch.XpBonus,
		TotalXp:    ch.TotalXp,
		Dead:       ch.Dead,
	}

	return GetQueries().UpdateCharacter(ctx, update)
}

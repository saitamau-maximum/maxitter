package migrations

import (
	"context"
	"fmt"

	"github.com/saitamau-maximum/maxitter/backend/usecase/model"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		_, err := db.NewCreateTable().Model((*model.User)(nil)).Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		_, err := db.NewDropTable().Model((*model.User)(nil)).IfExists().Exec(ctx)
		return err
	})
}

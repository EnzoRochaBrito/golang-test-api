package adapters

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func getSqlcAdapter(ctx context.Context, connstr, adpter string) {
	conn, err := pgx.Connect(ctx, connstr)
	// pgx.Connect(ctx, connstr)

	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)



}

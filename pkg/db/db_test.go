package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	. "github.com/smartystreets/goconvey/convey"
)

var dbConn = env("DB_CONN", "postgresql://localhost:5432/uteka_test_apisrv?sslmode=disable")

var testDb *DB

func env(v, def string) string {
	if r := os.Getenv(v); r != "" {
		return r
	}

	return def
}

type testdbLogger struct{}

func (d testdbLogger) BeforeQuery(ctx context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d testdbLogger) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func NewTestDb() *DB {
	cfg, err := pg.ParseURL(dbConn)
	if err != nil {
		panic(err)
	}
	dbc := pg.Connect(cfg)
	dbc.AddQueryHook(testdbLogger{})
	return New(dbc)

}

func TestMain(m *testing.M) {
	testDb = NewTestDb()
	runTests := m.Run()
	os.Exit(runTests)
}

func TestDBConnection(t *testing.T) {
	Convey("Test db connection", t, func() {
		err := testDb.Ping(context.Background())
		So(err, ShouldBeNil)
	})
}

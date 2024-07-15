package query


import(
	"context"
	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	// spanneraccessor "github.com/GoogleCloudPlatform/spanner-migration-tool/accessors/spanner"
)

type querySpanner interface {
	// IsSpannerSupoorted checks whether the default value from Source database is supported by Spanner or not. 
	IsSpannerSupported(defaultval string, columntype string) bool
	// firing query to spanner to cast default value based on spanner column type.
	query(db string, defaultval string, ty string) error
}

type querySpannerImpl struct{}

func (q *querySpannerImpl) IsSpannerSupported(defaultval string, columntype string) bool {
	db := "projects/cloud-spanner-intern/instances/testing-instance/databases/functions_test"
	err := q.query(db, defaultval, columntype)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (q *querySpannerImpl) query(db string, defaultval string, ty string) error {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return err
	}
	defer client.Close()

	stmt := spanner.Statement{
		SQL: "SELECT CAST("+defaultval+" AS "+ty+") AS ConvertedDefaultval",
	}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}

	}
}
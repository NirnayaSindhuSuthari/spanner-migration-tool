package spannermetadataaccessor

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	spannermetadataclient "github.com/GoogleCloudPlatform/spanner-migration-tool/accessors/spanner/spannermetadataaccessor/client"
	"github.com/GoogleCloudPlatform/spanner-migration-tool/common/constants"
	"google.golang.org/api/iterator"
)

type SpannerMetadataAccessor interface {
	// IsSpannerSupported checks whether the default value from Source database is supported by Spanner or not.
	IsSpannerSupported(defaultval string, columntype string) bool
	// firing query to spanner to cast default value based on spanner column type.
	query(db string, defaultval string, ty string) error
}

type SpannerMetadataAccessorImpl struct{}

func (spm *SpannerMetadataAccessorImpl) IsSpannerSupportedStatement(SpProjectId string, SpInstanceId string, statement string, columntype string) bool {
	// db := "projects/cloud-spanner-intern/instances/testing-instance/databases/functions_test"
	db := getSpannerUri(SpProjectId, SpInstanceId)
	err := spm.isValidSpannerStatement(db, statement, columntype)
	if err != nil {
		return false
	} else {
		return true
	}
}
func (spm *SpannerMetadataAccessorImpl) isValidSpannerStatement(db string, statement string, ty string) error {
	ctx := context.Background()
	// client, err := spanner.NewClient(ctx, db)
	spmClient, err := spannermetadataclient.GetOrCreateClient(ctx, db)
	if err != nil {
		return err
	}
	// defer spmClient.Close()

	stmt := spanner.Statement{
		SQL: "SELECT CAST(" + statement + " AS " + ty + ") AS ConvertedDefaultval",
	}
	iter := spmClient.Single().Query(ctx, stmt)
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

func getSpannerUri(projectId string, instanceId string) string {
	return fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectId, instanceId, constants.METADATA_DB)
}

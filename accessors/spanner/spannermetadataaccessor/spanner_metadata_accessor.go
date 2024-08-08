// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package spannermetadataaccessor

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	// spannermetadataaccessorclient "github.com/GoogleCloudPlatform/spanner-migration-tool/accessors/clients/spanner/spannermetadataaccessor"
	spannermetadataclient "github.com/GoogleCloudPlatform/spanner-migration-tool/accessors/spanner/spannermetadataaccessor/clients"
	"github.com/GoogleCloudPlatform/spanner-migration-tool/common/constants"
	"google.golang.org/api/iterator"
)

type SpannerMetadataAccessor interface {
	// IsSpannerSupportedStatement checks if the given statement is supported by Spanner.
	IsSpannerSupportedStatement(SpProjectId string, SpInstanceId string, defaultval string, columntype string) bool
	// isValidSpannerStatement queries spanner and checks if statement evaluates to a data corresponding to given type.
	isValidSpannerStatement(db string, defaultval string, datatype string) error
	// isValidSpannerStatement(db string, defaultval string, ty string) error
	// getClient(ctx context.Context, db string) (*spanner.Client, error)
}

type SpannerMetadataAccessorImpl struct{}

func (spm *SpannerMetadataAccessorImpl) IsSpannerSupportedStatement(SpProjectId string, SpInstanceId string, statement string, columntype string) bool {
	db := getSpannerUri(SpProjectId, SpInstanceId)
	if(SpProjectId == "" || SpInstanceId == ""){
		return false
	}
	err := spm.isValidSpannerStatement(db, statement, columntype)
	if err != nil {
		return false
	} else {
		return true
	}
}
// func getClient(ctx context.Context, db string) (*spanner.Client, error) {
// 	// return spannermetadataaccessorclient.GetOrCreateClient(ctx,db)
// 	return spannermetadataclient.GetOrCreateClient(ctx,db)
// }

// type rowIterator interface {
// 	Stop()
// 	Next() (*spanner.Row, error)
// }
// func querySpanner(ctx context.Context, client *spanner.Client, stmt spanner.Statement) (rowIterator) {
// 	return client.Single().Query(ctx, stmt)
// }

func (spm *SpannerMetadataAccessorImpl) isValidSpannerStatement(db string, statement string, datatype string) error {
	ctx := context.Background()
	// spmClient, err := spannermetadataclient.GetOrCreateClient(ctx, db)
	// spmClient, err := getClient(ctx, db)
	spmClient, err := spannermetadataclient.GetOrCreateClient(ctx, db)
	if err != nil {
		return err
	}

	if spmClient == nil {
		return fmt.Errorf("Spanner metadata Client is nil")
	}
	stmt := spanner.Statement{
		SQL: "SELECT CAST(" + statement + " AS " + datatype + ") AS statementValue",
	}
	// iter := querySpanner(ctx, spmClient, stmt)
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
package spannermetadataaccessorclient

import (
	"context"
	"fmt"
	"sync"

	sp "cloud.google.com/go/spanner"
)

var once sync.Once
var spannermetadataClient *sp.Client

var newClient = sp.NewClient

func GetOrCreateClient(ctx context.Context, dbURI string) (*sp.Client, error) {
	var err error
	if spannermetadataClient == nil {
		once.Do(func() {
			spannermetadataClient, err = newClient(ctx, dbURI)
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create spanner metadata database client: %v", err)
		}
		return spannermetadataClient, nil
	}
	return spannermetadataClient, nil
}
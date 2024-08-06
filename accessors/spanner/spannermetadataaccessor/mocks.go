package spannermetadataaccessor

import (
	"context"
	// "testing"

	"cloud.google.com/go/spanner"
	// spannermetadataclient "github.com/GoogleCloudPlatform/spanner-migration-tool/accessors/spanner/spannermetadataaccessor/clients"
	// "google.golang.org/api/iterator"
)

// MockSpannerMetadataAccessorImpl mocks the SpannerMetadataAccessorImpl for testing
type SpannerMetadataAccessorMock struct {
	getClientMock func(ctx context.Context, db string) (*spanner.Client, error) 

}

// NewMockSpannerMetadataAccessorImpl creates a new mock instance
// func NewMockSpannerMetadataAccessorImpl(t *testing.T) *MockSpannerMetadataAccessorImpl {
// 	// ctrl := gomock.NewController(t)
// 	// return &MockSpannerMetadataAccessorImpl{ctrl: ctrl}
// 	return &MockSpannerMetadataAccessorImpl{}
// }
func (m *SpannerMetadataAccessorMock) getClient(ctx context.Context, db string) (*spanner.Client, error) {
		// if m.getClientMock != nil {
			return m.getClientMock(ctx, db)
		// }
		// return m.spmClient, nil // Use pre-defined client if available
}

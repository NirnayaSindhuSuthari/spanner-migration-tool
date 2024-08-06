package spannermetadataaccessor

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/spanner"
	// spannermetadataclient "github.com/GoogleCloudPlatform/spanner-migration-tool/accessors/spanner/spannermetadataaccessor/clients"
	"github.com/stretchr/testify/assert"
	// "google.golang.org/api/iterator"
)

// MockSpannerMetadataAccessorImpl mocks the SpannerMetadataAccessorImpl for testing
type MockSpannerMetadataAccessorImpl struct {
	// ctrl     *gomock.Controller
	mockFn   func(ctx context.Context, db string) (*spanner.Client, error)
	spmClient *spanner.Client // Optional for mocking successful client creation
}

// NewMockSpannerMetadataAccessorImpl creates a new mock instance
func NewMockSpannerMetadataAccessorImpl(t *testing.T) *MockSpannerMetadataAccessorImpl {
	// ctrl := gomock.NewController(t)
	// return &MockSpannerMetadataAccessorImpl{ctrl: ctrl}
	return &MockSpannerMetadataAccessorImpl{}
}
func (spm *MockSpannerMetadataAccessorImpl) getClientMock(ctx context.Context, db string) (*spanner.Client, error) {
	if spm.mockFn != nil {
		return spm.mockFn(ctx, db)
	}
	return spm.spmClient, nil // Use pre-defined client if available
}


// type MockSpannerMetadataAccessor struct {
// 	getClientMock func (ctx context.Context, db string) (*spanner.Client, error)

// 		// ctrl     *gomock.Controller
// 		// mockFn   func(ctx context.Context, db string) (*spanner.Client, error)
// 		// spmClient *spanner.Client // Optional for mocking successful client creation
// }
// func NewMockSpannerMetadataAccessorImpl(t *testing.T) *MockSpannerMetadataAccessorImpl {
// 	return 
// }
// func NewMockSpannerMetadataAccessorImpl(t *testing.T) *MockSpannerMetadataAccessor {
// 		// ctrl := gomock.NewController(t)
// 		return &MockSpannerMetadataAccessor{}
// }
// func (m *MockSpannerMetadataAccessor) getClient(ctx context.Context, db string) (*spanner.Client, error) {
// 	// if m.mockFn != nil {
// 	// 	return m.mockFn(ctx, db)
// 	// }
// 	// return m.spmClient, nil // Use pre-defined client if available
// 	return 	m.getClientMock(ctx, db )
// }

// EXPECTATIONS (Modify as needed)
// func (m *MockSpannerMetadataAccessorImpl) getClientMock(ctx context.Context, db string) (*spanner.Client, error) {
// 	if m.mockFn != nil {
// 		return m.mockFn(ctx, db)
// 	}
// 	return m.spmClient, nil // Use pre-defined client if available
// }

// // TestIsSpannerSupportedStatement tests the IsSpannerSupportedStatement function
// func TestIsSpannerSupportedStatement(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		projectId     string
// 		instanceId    string
// 		statement     string
// 		columntype    string
// 		expected      bool
// 		mockFn        func(ctx context.Context, db string) (*spanner.Client, error)
// 	}{
// 		{
// 			name:          "Valid statement",
// 			projectId:     "my-project",
// 			instanceId:    "my-instance",
// 			statement:     "valid_statement",
// 			columntype:    "STRING",
// 			expected:      true,
// 			mockFn:        func(ctx context.Context, db string) (*spanner.Client, error) { return nil },
// 		},
// 		{
// 			name:          "Invalid statement",
// 			projectId:     "my-project",
// 			instanceId:    "my-instance",
// 			statement:     "invalid_statement",
// 			columntype:    "STRING",
// 			expected:      false,
// 			mockFn:        func(ctx context.Context, db string) (*spanner.Client, error) { return errors.New("invalid statement") },
// 		},
// 		{
// 			name:          "Empty project ID",
// 			projectId:     "",
// 			instanceId:    "my-instance",
// 			statement:     "valid_statement",
// 			columntype:    "STRING",
// 			expected:      false,
// 			mockFn:        nil,
// 		},
// 		{
// 			name:          "Empty instance ID",
// 			projectId:     "my-project",
// 			instanceId:    "",
// 			statement:     "valid_statement",
// 			columntype:    "STRING",
// 			expected:      false,
// 			mockFn:        nil,
// 		},
// 	}
	
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create a mock SpannerMetadataAccessorImpl
// 			mock := NewMockSpannerMetadataAccessorImpl(t)
// 			defer mock.ctrl.Finish()

// 			// Set mock behavior if needed
// 			if tc.mockFn != nil {
// 				mock.EXPECT().getClient(context.Background(), getSpannerUri(tc.projectId, tc.instanceId)).Return(tc.mockFn(context.Background(), getSpannerUri(tc.projectId, tc.instanceId)))
// 			}
// 			// spm.getClient = func(){
// 			// 	return m.mockFn, nil
// 			// }
// 			ctx := context.Background()
// 			db := "projects/project-id/instances/instance-id/databases/spannermigrationtool_metadata"
// 			spm := SpannerMetadataAccessorImpl{}

// 			spm.getClient = m.getClient(ctx,db)
// 			result := spm.IsSpannerSupportedStatement(tc.projectId, tc.instanceId, tc.statement, tc.columntype)
// 			assert.Equal(t, tc.expected, result, tc.name)

// 		})
// 	}
// }


// TestIsValidSpannerStatement tests the isValidSpannerStatement function
func TestIsValidSpannerStatement(t *testing.T) {
	// Define test cases
	tests := []struct {
		name      string
		statement string
		ty        string
		wantErr   error
		// getClientMock    func(ctx context.Context, db string) (*spanner.Client, error)
	}{
		{
			name:      "Valid statement",
			statement: "'John Doe'",
			ty:        "STRING",
			wantErr:   nil,
			// getClientMock: func(ctx context.Context, db string) (*spanner.Client, error) {
			// 	return &spanner.Client{}, nil
			// },
		},
		{
			name:      "Invalid statement",
			statement: "4",
			ty:        "STRING",
			wantErr:   nil,
			// getClientMock: func(ctx context.Context, db string) (*spanner.Client, error) {
			// 	return &spanner.Client{}, nil
			// },
		},
		{
			name:      "Empty statement",
			statement: "",
			ty:        "STRING",
			wantErr:   fmt.Errorf("statement cannot be empty"), // Expected error for empty statement
			// getClientMock: func(ctx context.Context, db string) (*spanner.Client, error) {
			// 	return nil, fmt.Errorf("statement cannot be empty")
			// },
		},
		{
			name:      "Spanner metadata Client is nil",
			statement: "'John Doe'",
			ty:        "STRING",
			wantErr:   fmt.Errorf("Spanner metadata Client is nil"), // Expected error from mock
			// getClientMock: func(ctx context.Context, db string) (*spanner.Client, error) {
			// 	return nil, fmt.Errorf("Spanner metadata Client is nil")
			// },
		},
		{
			name:      "Iterator error",
			statement: "'John Doe'",
			ty:        "STRING",
			wantErr:   fmt.Errorf("some iterator error"), // Simulate iterator error
			// getClientMock: func(ctx context.Context, db string) (*spanner.Client, error) {
			// 	return nil, fmt.Errorf("some iterator error")
			// },
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock SpannerMetadataAccessorImpl (optional)
			// You can uncomment and modify this section if you want to mock the GetOrCreateClient behavior

			// mockCtrl := gomock.NewController(t)
			// defer mockCtrl.Finish()
			db := "projects/project-id/instances/instance-id/databases/spannermigrationtool_metadata"
			// mock := NewMockSpannerMetadataAccessorImpl(t)
			mock := NewMockSpannerMetadataAccessorImpl(t)
			// mock.getClient(context.Background(), db).Return(tc.getClientMock(context.Background(), db))
			// mock.getClient(context.Background(), db).Return(tc.getClientMock(context.Background(), db))

			// Use the actual SpannerMetadataAccessorImpl for this test
			// ctx := context.Background()
			spm := &SpannerMetadataAccessorImpl{}
			// db := "projects/project-id/instances/instance-id/databases/spannermigrationtool_metadata"
			// spmClient := &tc.mockFn
			// spm.getClient := func() {
			// 	return mock.getClientMock() 
			// }
		    spm.getClient = mock.getClientMock
			// getClient := func(ctx context.Context, db string) (*spanner.Client, error) {
			// 	// Simulate successful query execution
			// 	return &spanner.Client{}, nil
			// }
			// spmClient, er := tc.getClientMock(ctx, db)
			// Call the function with test data
			// spm.getClient = func(ctx context.Context, db string) (*spanner.Client, error) {
			// 	// Simulate successful query execution (no need for mocks here)
			// 	return mock.getClientMock(ctx,db)
			// 	// return &spanner.Client{}, nil
			// }
			err := spm.isValidSpannerStatement(db, tc.statement, tc.ty)
			fmt.Println(err)
			assert.Equal(t, tc.wantErr, err != nil, tc.name)
			// Compare the actual error with the expected error
			// if diff := cmp.Diff(tc.wantErr, err, cmp.Comparer(nil)); diff != "" {
			// 	t.Errorf("Error mismatch (-want +got):\n%s", diff)
			// }
		})
	}
}

package education

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersItemAssignmentsItemCategoriesItemRefRequestBuilder provides operations to manage the collection of educationRoot entities.
type UsersItemAssignmentsItemCategoriesItemRefRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// UsersItemAssignmentsItemCategoriesItemRefRequestBuilderDeleteQueryParameters remove an educationCategory from an educationAssignment. Only teachers can perform this operation.
type UsersItemAssignmentsItemCategoriesItemRefRequestBuilderDeleteQueryParameters struct {
	// Delete Uri
	Id *string `uriparametername:"%40id"`
}

// UsersItemAssignmentsItemCategoriesItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemCategoriesItemRefRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *UsersItemAssignmentsItemCategoriesItemRefRequestBuilderDeleteQueryParameters
}

// NewUsersItemAssignmentsItemCategoriesItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemCategoriesItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersItemAssignmentsItemCategoriesItemRefRequestBuilder {
	m := &UsersItemAssignmentsItemCategoriesItemRefRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}/categories/{educationCategory%2Did}/$ref{?%40id*}", pathParameters),
	}
	return m
}

// NewUsersItemAssignmentsItemCategoriesItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemCategoriesItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersItemAssignmentsItemCategoriesItemRefRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewUsersItemAssignmentsItemCategoriesItemRefRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete remove an educationCategory from an educationAssignment. Only teachers can perform this operation.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/educationassignment-remove-category?view=graph-rest-1.0
func (m *UsersItemAssignmentsItemCategoriesItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemCategoriesItemRefRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToDeleteRequestInformation remove an educationCategory from an educationAssignment. Only teachers can perform this operation.
func (m *UsersItemAssignmentsItemCategoriesItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemCategoriesItemRefRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

package education

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ClassesItemAssignmentsItemCategoriesDeltaRequestBuilder provides operations to call the delta method.
type ClassesItemAssignmentsItemCategoriesDeltaRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ClassesItemAssignmentsItemCategoriesDeltaRequestBuilderGetQueryParameters invoke function delta
type ClassesItemAssignmentsItemCategoriesDeltaRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Order items by property values
	Orderby []string `uriparametername:"%24orderby"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// ClassesItemAssignmentsItemCategoriesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemCategoriesDeltaRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ClassesItemAssignmentsItemCategoriesDeltaRequestBuilderGetQueryParameters
}

// NewClassesItemAssignmentsItemCategoriesDeltaRequestBuilderInternal instantiates a new DeltaRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemCategoriesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ClassesItemAssignmentsItemCategoriesDeltaRequestBuilder {
	m := &ClassesItemAssignmentsItemCategoriesDeltaRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/categories/delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
	}
	return m
}

// NewClassesItemAssignmentsItemCategoriesDeltaRequestBuilder instantiates a new DeltaRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemCategoriesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ClassesItemAssignmentsItemCategoriesDeltaRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewClassesItemAssignmentsItemCategoriesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}

// Get invoke function delta
func (m *ClassesItemAssignmentsItemCategoriesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemCategoriesDeltaRequestBuilderGetRequestConfiguration) (ClassesItemAssignmentsItemCategoriesDeltaResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateClassesItemAssignmentsItemCategoriesDeltaResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ClassesItemAssignmentsItemCategoriesDeltaResponseable), nil
}

// ToGetRequestInformation invoke function delta
func (m *ClassesItemAssignmentsItemCategoriesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemCategoriesDeltaRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	requestInfo.Headers.Add("Accept", "application/json")
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

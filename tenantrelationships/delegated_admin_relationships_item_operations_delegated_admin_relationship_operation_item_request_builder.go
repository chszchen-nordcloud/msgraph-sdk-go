package tenantrelationships

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder provides operations to manage the operations property of the microsoft.graph.delegatedAdminRelationship entity.
type DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetQueryParameters read the properties of a delegatedAdminRelationshipOperation object.
type DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetQueryParameters
}

// DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewDelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderInternal instantiates a new DelegatedAdminRelationshipOperationItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder {
	m := &DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}/operations/{delegatedAdminRelationshipOperation%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewDelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder instantiates a new DelegatedAdminRelationshipOperationItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete navigation property operations for tenantRelationships
func (m *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderDeleteRequestConfiguration) error {
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

// Get read the properties of a delegatedAdminRelationshipOperation object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/delegatedadminrelationshipoperation-get?view=graph-rest-1.0
func (m *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipOperationable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipOperationable), nil
}

// Patch update the navigation property operations in tenantRelationships
func (m *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipOperationable, requestConfiguration *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderPatchRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipOperationable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipOperationable), nil
}

// ToDeleteRequestInformation delete navigation property operations for tenantRelationships
func (m *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToGetRequestInformation read the properties of a delegatedAdminRelationshipOperation object.
func (m *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the navigation property operations in tenantRelationships
func (m *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipOperationable, requestConfiguration *DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
	requestInfo.Headers.Add("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

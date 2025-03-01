package tenantrelationships

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder provides operations to manage the accessAssignments property of the microsoft.graph.delegatedAdminRelationship entity.
type DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetQueryParameters read the properties of a delegatedAdminAccessAssignment object.
type DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetQueryParameters
}

// DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewDelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderInternal instantiates a new DelegatedAdminAccessAssignmentItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder {
	m := &DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}/accessAssignments/{delegatedAdminAccessAssignment%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewDelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder instantiates a new DelegatedAdminAccessAssignmentItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete a delegatedAdminAccessAssignment object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/delegatedadminaccessassignment-delete?view=graph-rest-1.0
func (m *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderDeleteRequestConfiguration) error {
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

// Get read the properties of a delegatedAdminAccessAssignment object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/delegatedadminaccessassignment-get?view=graph-rest-1.0
func (m *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminAccessAssignmentable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminAccessAssignmentable), nil
}

// Patch update the properties of a delegatedAdminAccessAssignment object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/delegatedadminaccessassignment-update?view=graph-rest-1.0
func (m *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminAccessAssignmentable, requestConfiguration *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderPatchRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminAccessAssignmentable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminAccessAssignmentable), nil
}

// ToDeleteRequestInformation delete a delegatedAdminAccessAssignment object.
func (m *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToGetRequestInformation read the properties of a delegatedAdminAccessAssignment object.
func (m *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the properties of a delegatedAdminAccessAssignment object.
func (m *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminAccessAssignmentable, requestConfiguration *DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

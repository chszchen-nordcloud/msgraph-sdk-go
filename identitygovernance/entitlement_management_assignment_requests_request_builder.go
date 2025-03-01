package identitygovernance

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementAssignmentRequestsRequestBuilder provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
type EntitlementManagementAssignmentRequestsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EntitlementManagementAssignmentRequestsRequestBuilderGetQueryParameters in Azure AD entitlement management, retrieve a list of accessPackageAssignmentRequest objects.  The resulting list includes all the assignment requests, current and well as expired, that the caller has access to read, across all catalogs and access packages.
type EntitlementManagementAssignmentRequestsRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
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

// EntitlementManagementAssignmentRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAssignmentRequestsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *EntitlementManagementAssignmentRequestsRequestBuilderGetQueryParameters
}

// EntitlementManagementAssignmentRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAssignmentRequestsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByAccessPackageAssignmentRequestId provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
func (m *EntitlementManagementAssignmentRequestsRequestBuilder) ByAccessPackageAssignmentRequestId(accessPackageAssignmentRequestId string) *EntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if accessPackageAssignmentRequestId != "" {
		urlTplParams["accessPackageAssignmentRequest%2Did"] = accessPackageAssignmentRequestId
	}
	return NewEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewEntitlementManagementAssignmentRequestsRequestBuilderInternal instantiates a new AssignmentRequestsRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EntitlementManagementAssignmentRequestsRequestBuilder {
	m := &EntitlementManagementAssignmentRequestsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
	}
	return m
}

// NewEntitlementManagementAssignmentRequestsRequestBuilder instantiates a new AssignmentRequestsRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EntitlementManagementAssignmentRequestsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEntitlementManagementAssignmentRequestsRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *EntitlementManagementAssignmentRequestsRequestBuilder) Count() *EntitlementManagementAssignmentRequestsCountRequestBuilder {
	return NewEntitlementManagementAssignmentRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
func (m *EntitlementManagementAssignmentRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string) *EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnRequestBuilder {
	return NewEntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}

// Get in Azure AD entitlement management, retrieve a list of accessPackageAssignmentRequest objects.  The resulting list includes all the assignment requests, current and well as expired, that the caller has access to read, across all catalogs and access packages.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/entitlementmanagement-list-assignmentrequests?view=graph-rest-1.0
func (m *EntitlementManagementAssignmentRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentRequestsRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentRequestCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestCollectionResponseable), nil
}

// Post in Azure AD Entitlement Management, create a new accessPackageAssignmentRequest object.  This operation is used to assign a user to an access package, update the assignment, or to remove an access package assignment.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/entitlementmanagement-post-assignmentrequests?view=graph-rest-1.0
func (m *EntitlementManagementAssignmentRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable, requestConfiguration *EntitlementManagementAssignmentRequestsRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable), nil
}

// ToGetRequestInformation in Azure AD entitlement management, retrieve a list of accessPackageAssignmentRequest objects.  The resulting list includes all the assignment requests, current and well as expired, that the caller has access to read, across all catalogs and access packages.
func (m *EntitlementManagementAssignmentRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentRequestsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation in Azure AD Entitlement Management, create a new accessPackageAssignmentRequest object.  This operation is used to assign a user to an access package, update the assignment, or to remove an access package assignment.
func (m *EntitlementManagementAssignmentRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable, requestConfiguration *EntitlementManagementAssignmentRequestsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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

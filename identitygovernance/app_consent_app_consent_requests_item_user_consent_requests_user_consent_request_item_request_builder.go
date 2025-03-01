package identitygovernance

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder provides operations to manage the userConsentRequests property of the microsoft.graph.appConsentRequest entity.
type AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderGetQueryParameters read the properties and relationships of a userConsentRequest object.
type AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderGetQueryParameters
}

// AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// Approval provides operations to manage the approval property of the microsoft.graph.userConsentRequest entity.
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder) Approval() *AppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilder {
	return NewAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewAppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderInternal instantiates a new UserConsentRequestItemRequestBuilder and sets the default values.
func NewAppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder {
	m := &AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/appConsent/appConsentRequests/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewAppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder instantiates a new UserConsentRequestItemRequestBuilder and sets the default values.
func NewAppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete navigation property userConsentRequests for identityGovernance
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderDeleteRequestConfiguration) error {
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

// Get read the properties and relationships of a userConsentRequest object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/userconsentrequest-get?view=graph-rest-1.0
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserConsentRequestFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable), nil
}

// Patch update the navigation property userConsentRequests in identityGovernance
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderPatchRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserConsentRequestFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable), nil
}

// ToDeleteRequestInformation delete navigation property userConsentRequests for identityGovernance
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToGetRequestInformation read the properties and relationships of a userConsentRequest object.
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the navigation property userConsentRequests in identityGovernance
func (m *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, requestConfiguration *AppConsentAppConsentRequestsItemUserConsentRequestsUserConsentRequestItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

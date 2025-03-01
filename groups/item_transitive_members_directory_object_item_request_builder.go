package groups

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTransitiveMembersDirectoryObjectItemRequestBuilder provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
type ItemTransitiveMembersDirectoryObjectItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemTransitiveMembersDirectoryObjectItemRequestBuilderGetQueryParameters the direct and transitive members of a group. Nullable.
type ItemTransitiveMembersDirectoryObjectItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemTransitiveMembersDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitiveMembersDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemTransitiveMembersDirectoryObjectItemRequestBuilderGetQueryParameters
}

// NewItemTransitiveMembersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemTransitiveMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTransitiveMembersDirectoryObjectItemRequestBuilder {
	m := &ItemTransitiveMembersDirectoryObjectItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/transitiveMembers/{directoryObject%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemTransitiveMembersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemTransitiveMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTransitiveMembersDirectoryObjectItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemTransitiveMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the direct and transitive members of a group. Nullable.
func (m *ItemTransitiveMembersDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitiveMembersDirectoryObjectItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}

// GraphApplication casts the previous resource to application.
func (m *ItemTransitiveMembersDirectoryObjectItemRequestBuilder) GraphApplication() *ItemTransitiveMembersItemGraphApplicationRequestBuilder {
	return NewItemTransitiveMembersItemGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// GraphDevice casts the previous resource to device.
func (m *ItemTransitiveMembersDirectoryObjectItemRequestBuilder) GraphDevice() *ItemTransitiveMembersItemGraphDeviceRequestBuilder {
	return NewItemTransitiveMembersItemGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// GraphGroup casts the previous resource to group.
func (m *ItemTransitiveMembersDirectoryObjectItemRequestBuilder) GraphGroup() *ItemTransitiveMembersItemGraphGroupRequestBuilder {
	return NewItemTransitiveMembersItemGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// GraphOrgContact casts the previous resource to orgContact.
func (m *ItemTransitiveMembersDirectoryObjectItemRequestBuilder) GraphOrgContact() *ItemTransitiveMembersItemGraphOrgContactRequestBuilder {
	return NewItemTransitiveMembersItemGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// GraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemTransitiveMembersDirectoryObjectItemRequestBuilder) GraphServicePrincipal() *ItemTransitiveMembersItemGraphServicePrincipalRequestBuilder {
	return NewItemTransitiveMembersItemGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// GraphUser casts the previous resource to user.
func (m *ItemTransitiveMembersDirectoryObjectItemRequestBuilder) GraphUser() *ItemTransitiveMembersItemGraphUserRequestBuilder {
	return NewItemTransitiveMembersItemGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation the direct and transitive members of a group. Nullable.
func (m *ItemTransitiveMembersDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitiveMembersDirectoryObjectItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

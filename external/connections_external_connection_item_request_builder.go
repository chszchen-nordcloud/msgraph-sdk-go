package external

import (
	"context"
	i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc "github.com/chszchen-nordcloud/msgraph-sdk-go/models/externalconnectors"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ConnectionsExternalConnectionItemRequestBuilder provides operations to manage the connections property of the microsoft.graph.externalConnectors.external entity.
type ConnectionsExternalConnectionItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ConnectionsExternalConnectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsExternalConnectionItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ConnectionsExternalConnectionItemRequestBuilderGetQueryParameters read the properties and relationships of an externalConnection object.
type ConnectionsExternalConnectionItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ConnectionsExternalConnectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsExternalConnectionItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ConnectionsExternalConnectionItemRequestBuilderGetQueryParameters
}

// ConnectionsExternalConnectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsExternalConnectionItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewConnectionsExternalConnectionItemRequestBuilderInternal instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewConnectionsExternalConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ConnectionsExternalConnectionItemRequestBuilder {
	m := &ConnectionsExternalConnectionItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/connections/{externalConnection%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewConnectionsExternalConnectionItemRequestBuilder instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewConnectionsExternalConnectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ConnectionsExternalConnectionItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewConnectionsExternalConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes an externalConnection object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/externalconnectors-externalconnection-delete?view=graph-rest-1.0
func (m *ConnectionsExternalConnectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConnectionsExternalConnectionItemRequestBuilderDeleteRequestConfiguration) error {
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

// Get read the properties and relationships of an externalConnection object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/externalconnectors-externalconnection-get?view=graph-rest-1.0
func (m *ConnectionsExternalConnectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectionsExternalConnectionItemRequestBuilderGetRequestConfiguration) (i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.CreateExternalConnectionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable), nil
}

// Groups provides operations to manage the groups property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ConnectionsExternalConnectionItemRequestBuilder) Groups() *ConnectionsItemGroupsRequestBuilder {
	return NewConnectionsItemGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Items provides operations to manage the items property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ConnectionsExternalConnectionItemRequestBuilder) Items() *ConnectionsItemItemsRequestBuilder {
	return NewConnectionsItemItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Operations provides operations to manage the operations property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ConnectionsExternalConnectionItemRequestBuilder) Operations() *ConnectionsItemOperationsRequestBuilder {
	return NewConnectionsItemOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Patch update the properties of an externalConnection object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/externalconnectors-externalconnection-update?view=graph-rest-1.0
func (m *ConnectionsExternalConnectionItemRequestBuilder) Patch(ctx context.Context, body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable, requestConfiguration *ConnectionsExternalConnectionItemRequestBuilderPatchRequestConfiguration) (i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.CreateExternalConnectionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable), nil
}

// Schema provides operations to manage the schema property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ConnectionsExternalConnectionItemRequestBuilder) Schema() *ConnectionsItemSchemaRequestBuilder {
	return NewConnectionsItemSchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation deletes an externalConnection object.
func (m *ConnectionsExternalConnectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConnectionsExternalConnectionItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToGetRequestInformation read the properties and relationships of an externalConnection object.
func (m *ConnectionsExternalConnectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectionsExternalConnectionItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the properties of an externalConnection object.
func (m *ConnectionsExternalConnectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalConnectionable, requestConfiguration *ConnectionsExternalConnectionItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

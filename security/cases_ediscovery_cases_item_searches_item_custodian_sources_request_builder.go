package security

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/chszchen-nordcloud/msgraph-sdk-go/models/security"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder provides operations to manage the custodianSources property of the microsoft.graph.security.ediscoverySearch entity.
type CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderGetQueryParameters get the list of custodial data sources associated with an eDiscovery search.
type CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderGetQueryParameters struct {
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

// CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderGetQueryParameters
}

// ByDataSourceId provides operations to manage the custodianSources property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder) ByDataSourceId(dataSourceId string) *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesDataSourceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if dataSourceId != "" {
		urlTplParams["dataSource%2Did"] = dataSourceId
	}
	return NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesDataSourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderInternal instantiates a new CustodianSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder {
	m := &CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/custodianSources{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
	}
	return m
}

// NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder instantiates a new CustodianSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder) Count() *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesCountRequestBuilder {
	return NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get the list of custodial data sources associated with an eDiscovery search.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverysearch-list-custodiansources?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderGetRequestConfiguration) (idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateDataSourceCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceCollectionResponseable), nil
}

// ToGetRequestInformation get the list of custodial data sources associated with an eDiscovery search.
func (m *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

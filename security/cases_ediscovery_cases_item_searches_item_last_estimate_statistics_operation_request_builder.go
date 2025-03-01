package security

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/chszchen-nordcloud/msgraph-sdk-go/models/security"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder provides operations to manage the lastEstimateStatisticsOperation property of the microsoft.graph.security.ediscoverySearch entity.
type CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderGetQueryParameters get the last  ediscoveryEstimateOperation objects and their properties.
type CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderGetQueryParameters
}

// NewCasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderInternal instantiates a new LastEstimateStatisticsOperationRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder {
	m := &CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/lastEstimateStatisticsOperation{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewCasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder instantiates a new LastEstimateStatisticsOperationRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the last  ediscoveryEstimateOperation objects and their properties.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverysearch-list-lastestimatestatisticsoperation?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration) (idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryEstimateOperationable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryEstimateOperationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryEstimateOperationable), nil
}

// ToGetRequestInformation get the last  ediscoveryEstimateOperation objects and their properties.
func (m *CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

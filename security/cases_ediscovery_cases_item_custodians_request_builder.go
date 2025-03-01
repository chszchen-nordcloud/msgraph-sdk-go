package security

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/chszchen-nordcloud/msgraph-sdk-go/models/security"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CasesEdiscoveryCasesItemCustodiansRequestBuilder provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoveryCasesItemCustodiansRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CasesEdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters get a list of the custodian objects and their properties.
type CasesEdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters struct {
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

// CasesEdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CasesEdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters
}

// CasesEdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByEdiscoveryCustodianId provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) ByEdiscoveryCustodianId(ediscoveryCustodianId string) *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if ediscoveryCustodianId != "" {
		urlTplParams["ediscoveryCustodian%2Did"] = ediscoveryCustodianId
	}
	return NewCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal instantiates a new CustodiansRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemCustodiansRequestBuilder {
	m := &CasesEdiscoveryCasesItemCustodiansRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
	}
	return m
}

// NewCasesEdiscoveryCasesItemCustodiansRequestBuilder instantiates a new CustodiansRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemCustodiansRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) Count() *CasesEdiscoveryCasesItemCustodiansCountRequestBuilder {
	return NewCasesEdiscoveryCasesItemCustodiansCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get a list of the custodian objects and their properties.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverycase-list-custodians?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration) (idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCustodianCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianCollectionResponseable), nil
}

// MicrosoftGraphSecurityApplyHold provides operations to call the applyHold method.
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) MicrosoftGraphSecurityApplyHold() *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder {
	return NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// MicrosoftGraphSecurityRemoveHold provides operations to call the removeHold method.
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) MicrosoftGraphSecurityRemoveHold() *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityRemoveHoldRequestBuilder {
	return NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post create a new ediscoveryCustodian object.After the custodian object is created, you will need to create the custodian's userSource to reference their mailbox and OneDrive for Business site.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverycase-post-custodians?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) Post(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, requestConfiguration *CasesEdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration) (idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCustodianFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable), nil
}

// ToGetRequestInformation get a list of the custodian objects and their properties.
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation create a new ediscoveryCustodian object.After the custodian object is created, you will need to create the custodian's userSource to reference their mailbox and OneDrive for Business site.
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) ToPostRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, requestConfiguration *CasesEdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

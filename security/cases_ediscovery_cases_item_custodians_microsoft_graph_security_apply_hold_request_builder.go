package security

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder provides operations to call the applyHold method.
type CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderInternal instantiates a new MicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder {
	m := &CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/microsoft.graph.security.applyHold", pathParameters),
	}
	return m
}

// NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder instantiates a new MicrosoftGraphSecurityApplyHoldRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(urlParams, requestAdapter)
}

// Post invoke action applyHold
func (m *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldApplyHoldPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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

// ToPostRequestInformation invoke action applyHold
func (m *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldApplyHoldPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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

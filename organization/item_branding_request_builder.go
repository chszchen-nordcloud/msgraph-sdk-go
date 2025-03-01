package organization

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemBrandingRequestBuilder provides operations to manage the branding property of the microsoft.graph.organization entity.
type ItemBrandingRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemBrandingRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemBrandingRequestBuilderGetQueryParameters retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
type ItemBrandingRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemBrandingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemBrandingRequestBuilderGetQueryParameters
}

// ItemBrandingRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// BackgroundImage provides operations to manage the media for the organization entity.
func (m *ItemBrandingRequestBuilder) BackgroundImage() *ItemBrandingBackgroundImageRequestBuilder {
	return NewItemBrandingBackgroundImageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// BannerLogo provides operations to manage the media for the organization entity.
func (m *ItemBrandingRequestBuilder) BannerLogo() *ItemBrandingBannerLogoRequestBuilder {
	return NewItemBrandingBannerLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemBrandingRequestBuilderInternal instantiates a new BrandingRequestBuilder and sets the default values.
func NewItemBrandingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemBrandingRequestBuilder {
	m := &ItemBrandingRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/branding{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemBrandingRequestBuilder instantiates a new BrandingRequestBuilder and sets the default values.
func NewItemBrandingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemBrandingRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemBrandingRequestBuilderInternal(urlParams, requestAdapter)
}

// CustomCSS provides operations to manage the media for the organization entity.
func (m *ItemBrandingRequestBuilder) CustomCSS() *ItemBrandingCustomCSSRequestBuilder {
	return NewItemBrandingCustomCSSRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Delete delete the default organizational branding object. To delete the organizationalBranding object, all images (Stream types) must first be removed from the object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/organizationalbranding-delete?view=graph-rest-1.0
func (m *ItemBrandingRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemBrandingRequestBuilderDeleteRequestConfiguration) error {
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

// Favicon provides operations to manage the media for the organization entity.
func (m *ItemBrandingRequestBuilder) Favicon() *ItemBrandingFaviconRequestBuilder {
	return NewItemBrandingFaviconRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/organizationalbranding-get?view=graph-rest-1.0
func (m *ItemBrandingRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBrandingRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationalBrandingFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable), nil
}

// HeaderLogo provides operations to manage the media for the organization entity.
func (m *ItemBrandingRequestBuilder) HeaderLogo() *ItemBrandingHeaderLogoRequestBuilder {
	return NewItemBrandingHeaderLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Localizations provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
func (m *ItemBrandingRequestBuilder) Localizations() *ItemBrandingLocalizationsRequestBuilder {
	return NewItemBrandingLocalizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Patch update the properties of the default branding object specified by the organizationalBranding resource.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/organizationalbranding-update?view=graph-rest-1.0
func (m *ItemBrandingRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable, requestConfiguration *ItemBrandingRequestBuilderPatchRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationalBrandingFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable), nil
}

// SquareLogo provides operations to manage the media for the organization entity.
func (m *ItemBrandingRequestBuilder) SquareLogo() *ItemBrandingSquareLogoRequestBuilder {
	return NewItemBrandingSquareLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// SquareLogoDark provides operations to manage the media for the organization entity.
func (m *ItemBrandingRequestBuilder) SquareLogoDark() *ItemBrandingSquareLogoDarkRequestBuilder {
	return NewItemBrandingSquareLogoDarkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation delete the default organizational branding object. To delete the organizationalBranding object, all images (Stream types) must first be removed from the object.
func (m *ItemBrandingRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToGetRequestInformation retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
func (m *ItemBrandingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the properties of the default branding object specified by the organizationalBranding resource.
func (m *ItemBrandingRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationalBrandingable, requestConfiguration *ItemBrandingRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

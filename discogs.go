package discogs

import (
	"net/http"

	"github.com/irlndts/go-apirequest"
)

const (
	discogsAPI = "https://api.discogs.com/"
)

// Client is a Discogs client for making Discogs API requests.
type Client struct {
	api      *apirequest.API
	currency string
	//Release *ReleaseService
	Master *MasterService
	Artist *ArtistService
	Label  *LabelService
	Search *SearchService
}

// NewClient returns a new Client.
func NewClient(useragent, token string) *Client {
	base := apirequest.New().Client(&http.Client{}).Base(discogsAPI).Add("User-Agent", useragent).Add("Authorization", "Discogs token="+token)
	return &Client{
		api:      base,
		currency: "USD",

		Artist: newArtistService(base.New()),
		Label:  newLabelService(base.New()),
		Master: newMasterService(base.New()),
		//Release: newReleaseService(base.New()),
		Search: newSearchService(base.New()),
	}
}

// Token sets tokens, it's required for some queries like search
func (c *Client) Token(token string) *Client {
	c.api.Set("Authorization", "Discogs token="+token)
	return c
}

// UserAgent sets specified user agent
// Discogs requires it
func (c *Client) UserAgent(useragent string) *Client {
	c.api.Set("User-Agent", useragent)
	return c
}

// SetCurrency determines currency for marketplace data.
// Defaults to the authenticated users currency. Must be one of the following:
// USD GBP EUR CAD AUD JPY CHF MXN BRL NZD SEK ZAR
func (c *Client) Currency(currency string) error {
	c.currency = currency
	return nil
}

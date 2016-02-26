package metacpan

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("metacpan", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{
}

// BuildURI generates a search URL for MetaCPAN.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://metacpan.org/search?q=%s", url.QueryEscape(q))
}
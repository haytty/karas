package page

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
)

type Page struct {
	URL string
	DOM map[string]string
	mu  sync.Mutex
}

func NewPage() *Page {
	dom := make(map[string]string)
	return &Page{
		DOM: dom,
	}
}

func (p *Page) String() string {
	return fmt.Sprintf(
		color.GreenString("URL: %s, %v"),
		p.URL,
		p.DOM,
	)
}

func (p *Page) Set(key, value string) {
	p.mu.Lock()
	p.DOM[key] = value
	p.mu.Unlock()
}

func (base *Page) Match(p *Page) bool {
	// url check
	if !base.isSameURL(p) {
		return false
	}

	// dom check
	if !base.isSameDOM(p) {
		return false
	}

	return true
}

func (base *Page) isSameURL(p *Page) bool {
	return base.URL == p.URL
}

func (base *Page) isSameDOM(p *Page) bool {
	for baseKey, baseValue := range base.DOM {
		pv, ok := p.DOM[baseKey]
		if !ok {
			return false
		}
		if pv != baseValue {
			return false
		}
	}
	return true
}

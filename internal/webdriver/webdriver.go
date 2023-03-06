package webdriver

import (
	"fmt"

	"github.com/tebeka/selenium/chrome"

	"github.com/tebeka/selenium"
)

type WebDriver interface {
	NewWebDriver() (*selenium.Service, selenium.WebDriver, error)
}

type Selenium struct {
	selenium     string
	chrome       string
	chromeDriver string
	port         int
}

func (s *Selenium) NewWebDriver() (*selenium.Service, selenium.WebDriver, error) {
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(s.chromeDriver),
	}

	service, err := selenium.NewSeleniumService(s.selenium, s.port, opts...)
	if err != nil {
		return nil, nil, err
	}
	caps := selenium.Capabilities{"browserName": "chrome"}

	chrCaps := chrome.Capabilities{
		Path: s.chrome,
		Args: []string{
			// This flag is needed to test against Chrome binaries that are not the
			// default installation. The sandbox requires a setuid binary.
			"--no-sandbox",
		},
		W3C: true,
	}

	chrCaps.Args = append(chrCaps.Args, "--headless")
	caps.AddChrome(chrCaps)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", s.port))
	return service, wd, err
}

func NewSelenium(selenium, chrome, chromeDriver string, port int) *Selenium {
	return &Selenium{
		selenium:     selenium,
		chrome:       chrome,
		chromeDriver: chromeDriver,
		port:         port,
	}
}

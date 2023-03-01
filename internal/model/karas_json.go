package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/tebeka/selenium/chrome"

	"github.com/tebeka/selenium"

	"github.com/haytty/karas/internal/model/action"
)

type KarasJSON struct {
	Karas struct {
		Url     string          `json:"url"`
		Output  string          `json:"output"`
		Actions []action.Action `json:"actions"`
	} `json:"karas"`
}

func NewKarasJSON() *KarasJSON {
	return &KarasJSON{}
}

func (s *KarasJSON) Do() error {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath     = "drivers/selenium-server.jar"
		chromeDriverPath = "drivers/chromedriver"
		geckoDriverPath  = "drivers/geckodriver"
		firefoxPath      = "drivers/firefox/firefox"
		port             = 8080
	)

	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		//selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.ChromeDriver(chromeDriverPath),
		//selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	//selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		fmt.Println("NewSeleniumService")
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}

	chrCaps := chrome.Capabilities{
		Path: "./drivers/chrome-linux/chrome",
		Args: []string{
			// This flag is needed to test against Chrome binaries that are not the
			// default installation. The sandbox requires a setuid binary.
			"--no-sandbox",
		},
		W3C: true,
	}
	chrCaps.Args = append(chrCaps.Args, "--headless")
	caps.AddChrome(chrCaps)

	//f := firefox.Capabilities{
	//	Binary: firefoxPath,
	//}
	//caps.AddFirefox(f)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	//wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		fmt.Println("NewRemote")
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys(`
		package main
		import "fmt"

		func main() {
			fmt.Println("Hello WebDriver!\n")
		}
	`)

	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// please check
	time.Sleep(10 * time.Second)

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "span.stdout")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s\n", strings.Replace(output, "\n\n", "\n", -1))

	// Example Output:
	// Hello WebDriver!
	//
	// Program exited.

	return nil
}

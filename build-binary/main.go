package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {

	for true {

		var err error

		// create context
		ctxt, cancel := context.WithCancel(context.Background())
		defer cancel()

		// create chrome instance
		//c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
		c, err := chromedp.New(ctxt)

		if err != nil {
			log.Fatal(err)
		}

		// run task list
		var site, res string
		err = c.Run(ctxt, googleSearch("dog", "Home", &site, &res))
		if err != nil {
			log.Fatal(err)
		}

		// shutdown chrome
		err = c.Shutdown(ctxt)
		if err != nil {
			log.Fatal(err)
		}

		// wait for chrome to finish
		err = c.Wait()
		if err != nil {
			log.Fatal(err)
		}

	}
}

func googleSearch(q, text string, site, res *string) chromedp.Tasks {
	sel := fmt.Sprintf(`//a[text()[contains(., '%s')]]`, text)
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.google.com`),
		chromedp.WaitVisible(`#hplogo`, chromedp.ByID),
		chromedp.SendKeys(`#lst-ib`, q+"\n", chromedp.ByID),
		chromedp.WaitVisible(`#res`, chromedp.ByID),
		chromedp.Text(sel, res),
		chromedp.Click(sel),
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			_, err := fmt.Println("result:", *res)
			return err
		}),
	}
}

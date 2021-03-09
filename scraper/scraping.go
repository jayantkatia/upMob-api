package scraper

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	db "github.com/jayantkatia/upcoming_mobiles_api/db/sqlc"
)

const host = "https://www.91mobiles.com"

func remDupDevices(arrDevices db.Devices) db.Devices {
	mapDevices := make(map[string]bool)
	var result db.Devices

	for i, entry := range arrDevices {
		if _, ok := mapDevices[arrDevices[i].DeviceName]; !ok {
			mapDevices[arrDevices[i].DeviceName] = true
			result = append(result, entry)
		}
	}
	return result
}

func StartScraping(store *db.Store) {

	arrDevices, err := Scrap("/upcoming-mobiles-in-india")
	if arrDevices == nil {
		return
	}
	arrDevices = remDupDevices(arrDevices)

	if len(arrDevices) < 20 {
		log.Println("SCRAP:: No db Tx will occur since entries below 20")
		return
	}
	if err != nil {
		log.Println("SCRAP:: Overall error: ", err)
	}
	fmt.Println("Scraping done successfully")
	err = store.UpdateDBTx(context.Background(), arrDevices)
	if err != nil {
		log.Println("SCRAP:: Transaction error: ", err)
	}
	fmt.Println("Tx done successfully")
}

/* Scraping errors are subjected to change in css selectors by the host maintainers.*/

func Scrap(path string) (db.Devices, error) {

	// opts := append(chromedp.DefaultExecAllocatorOptions[:],
	// 	chromedp.DisableGPU,
	// 	chromedp.Flag("headless", false),
	// )

	// ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	// defer cancel()

	// ctx, cancel = chromedp.NewContext(ctx)
	// defer cancel()

	// Headless mode
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var containerDivs, name, date, price, image, specScore []*cdp.Node
	var arrDevices db.Devices

	err := chromedp.Run(ctx, chromedp.Navigate(host+path))
	if err != nil {
		log.Println("SCRAP:: Failed Navigation to host:", err)
		return nil, err
	}

	// If user wants strictly 100 results, then may change condition where to
	// len(arrDevices)<=100
	for page := 1; page <= 5; page++ {

		err = chromedp.Run(ctx, chromedp.Tasks{
			chromedp.ScrollIntoView(`#footer_gad`, chromedp.ByID),
			chromedp.WaitVisible(fmt.Sprintf("div[data-row=\"%d\"]", page*20), chromedp.ByQuery),

			chromedp.Nodes(".finder_snipet_wrap ", &containerDivs, chromedp.ByQueryAll), //a[data-type=\"name\"]
			chromedp.ActionFunc(func(context.Context) error {
				log.Printf(">>>> %d Page Scraped", page)
				return nil
			}),
		})
		if err != nil {
			log.Printf("SCRAP:: Failed surfing %d", page)
			log.Println(err)
			return arrDevices, err
		}

		for node := 0; node < len(containerDivs); node++ {

			if err = chromedp.Run(ctx, chromedp.Tasks{
				chromedp.Nodes("a[data-type=\"name\"]", &name, chromedp.ByQuery, chromedp.FromNode(containerDivs[node]), chromedp.AtLeast(0)),
				chromedp.Nodes("div.pro_list_date", &date, chromedp.ByQuery, chromedp.FromNode(containerDivs[node]), chromedp.AtLeast(0)),
				chromedp.Nodes("span.price_float", &price, chromedp.ByQuery, chromedp.FromNode(containerDivs[node]), chromedp.AtLeast(0)),
				chromedp.Nodes("div[data-type=\"Knofindewmore\"] > .rating_box_new_list", &specScore, chromedp.ByQuery, chromedp.FromNode(containerDivs[node]), chromedp.AtLeast(0)),
				chromedp.Nodes("img[data-type=\"image\"]", &image, chromedp.ByQuery, chromedp.FromNode(containerDivs[node]), chromedp.AtLeast(0)),
			}); err != nil {
				log.Printf("SCRAP:: Failed at %d page %d node", page, node)
				log.Println(err)
				continue
			}
			var device db.Device

			if len(name) > 0 {
				device.DeviceName = name[0].Children[0].NodeValue

				href, exists := name[0].Attribute("href")
				device.SourceUrl = host
				if exists {
					device.SourceUrl += href
				}
			} else {
				continue
			}

			if len(date) > 0 {
				device.Expected = date[0].Children[0].NodeValue
			} else {
				device.Expected = "-"
			}

			if len(specScore) > 0 {
				specString := strings.Trim(specScore[0].Children[0].NodeValue, " %")
				spec, err := strconv.ParseInt(specString, 10, 32)
				if err != nil {
					device.SpecScore = 0
				} else {
					device.SpecScore = int32(spec)
				}
			} else {
				device.SpecScore = 0
			}

			if len(image) > 0 {
				src, exists := image[0].Attribute("src")
				if exists {
					device.ImgUrl = "https:" + strings.Split(src, "%")[0]
				}
			} else {
				continue
			}

			if len(price) > 0 {
				var text string
				if err = chromedp.Run(ctx,
					chromedp.Text([]cdp.NodeID{price[0].NodeID}, &text, chromedp.ByNodeID),
				); err != nil {
					continue
				}
				commaSeparated := strings.Split(text, ".")[1]
				arrayStrings := strings.Split(commaSeparated, ",")
				priceString := strings.Join(arrayStrings, "")
				if priceInt, err := strconv.ParseInt(priceString, 10, 64); err != nil {
					continue
				} else {
					device.Price = priceInt
				}

			} else {
				continue
			}

			arrDevices = append(arrDevices, device)
		}
		err := chromedp.Run(ctx, chromedp.Click(`div.listing-btns4 > span.list-bttnn`, chromedp.NodeVisible))
		if err != nil {
			log.Printf("SCRAP:: Failed to go to %d on click", page+1)
			log.Println(err)
			//returning since no need to add duplicates by remaining on the same page
			return arrDevices, err
		}
	}
	return arrDevices, nil
}

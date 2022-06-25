package scraper

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	db "github.com/jayantkatia/upcoming_mobiles_api/db/sqlc"
)

type Response struct {
	ResponseData string `json:"response"`
}

var links []string
var host = "https://www.91mobiles.com"

func Scraper(queries *db.Queries) {
	// creating an http client with timeout
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	// scraping links from list pages
	scrapeLinks(getLink(0, 1), client)
	for i := 1; i < 6; i++ {
		scrapeLinks(getLink(1, i), client)
	}

	// Links visited and data scraped
	fmt.Println(len(links))
	for i := len(links) - 1; i >= 0; i-- {
		visitPage(links[i], client, queries)
	}
}

func getLink(show_next int, page int) string {
	return "https://www.91mobiles.com/template/category_finder/finder_ajax.php?show_next=" + strconv.Itoa(show_next) + "&excludeId=&hash=&search=&hidFrmSubFlag=1&page=" + strconv.Itoa(page) + "&category=mobile&unique_sort=ga_views&gaCategory=Upcoming+Mobiles+Price+List+in+India-filter&requestType=1&showPagination=1&listType=list&listType_v3=list&listType_v1=list&listType_v2=list&listType_v4=list&listType_v5=list&listType_v6=list&page_type=upcoming&finderRuleUrl=&selMobSort=ga_views&hdnCategory=mobile&user_search=&url_feat_rule=upcoming-mobiles-in-india&buygaCat=upcoming-mob&amount=0%3B200000&sCatName=mobile&price_range_apply=1&tr_fl%5B%5D=mob_market_status_filter.marketstatus_filter%3Aupcoming&tr_fl%5B%5D=mob_market_status_filter.marketstatus_filter%3Arumoured"
}

func scrapeLinks(link string, client *http.Client) {
	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Printf("ERROR :: while making request for a list page")
		return
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 UpMob-WebScrapingBot/scrapes webpages to fetch upcoming device launches(www.github.com/jayantkatia/upmob-api)")
	request.Header.Set("x-requested-with", "XMLHttpRequest")
	response, err := client.Do(request)
	if err != nil {
		log.Printf("ERROR :: while getting response from list page")
		return
	}
	defer response.Body.Close()

	var responseData Response
	s, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("ERROR :: reading list page response")
		return
	}
	json.Unmarshal(s, &responseData)

	document, err := goquery.NewDocumentFromReader(strings.NewReader(responseData.ResponseData))
	if err != nil {
		log.Printf("ERROR :: creating goquery document from list page")
		return
	}

	document.Find("div.finder_snipet_wrap").Each(func(i int, s *goquery.Selection) {
		name := s.Find("a[data-type=\"name\"]")
		href, exists := name.Attr("href")
		if exists {
			links = append(links, host+href)
		}
	})
}

func visitPage(link string, client *http.Client, queries *db.Queries) {
	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Printf("ERROR :: while making request for details page")
		return
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 UpMob-WebScrapingBot/scrapes webpages to fetch upcoming device launches(www.github.com/jayantkatia/upmob-api)")
	response, err := client.Do(request)
	if err != nil {
		log.Printf("ERROR :: while getting response for details page")
		return
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("ERROR :: while making goquery document for details page")
		return
	}

	last_updated := document.Find(".label_hd").First()
	name := document.Find("h1.h1_pro_head")

	lu, err := queries.GetLastUpdatedDevice(context.Background(), name.Text())
	isRecordExists := false
	if err == nil && lu == last_updated.Text() {
		queries.UpdateScrapeTimestamp(context.Background(), time.Now())
		log.Printf("Updating ScrapeTimeStamp, Already existing record")
		return
	} else if err == nil {
		log.Printf("Already existing record")
		isRecordExists = true
	}

	price := document.Find("span.big_prc[itemprop=\"price\"]").Text()
	arrayStrings := strings.Split(price, ",")
	priceString := strings.Join(arrayStrings, "")
	priceInt, err := strconv.ParseInt(strings.TrimSpace(priceString), 10, 32)
	if err != nil {
		log.Printf("ERROR:: while converting price string to int in details page")
		return
	}

	spec_score := document.Find("div.top_box>div>div").Text()
	spec_score_string := strings.Trim(spec_score, "%")
	spec_score_int, err := strconv.Atoi(spec_score_string)
	if err != nil {
		spec_score_int = 0
	}

	img := document.Find("img.overview_lrg_pic_img")
	img_src, _ := img.Attr("src")

	device := make(map[string]string)
	document.Find("div.spec_box").Each(func(i int, spec_box *goquery.Selection) {
		spec_box.Find("tr").Each(func(i int, spec *goquery.Selection) {
			spec_title := spec.Find(".spec_ttle").Text()
			spec_des := spec.Find(".spec_des").Contents().Not("span").Text()
			if len(device[spec_title]) == 0 {
				device[spec_title] = strings.TrimSpace(spec_des)
			} else if len(device["m "+spec_title]) == 0 {
				device["m "+spec_title] = strings.TrimSpace(spec_des)
			} else {
				device["t "+spec_title] = strings.TrimSpace(spec_des)
			}
		})
	})

	if isRecordExists {
		err := queries.DeleteDevice(context.Background(), name.Text())
		if err != nil {
			log.Printf("ERROR:: while deleting previous record")
			return
		}
	}

	_, err = queries.InsertDevice(context.Background(), db.InsertDeviceParams{
		DeviceName:        name.Text(),
		LastUpdated:       last_updated.Text(),
		Expected:          device["Launch Date"],
		Price:             int32(priceInt),
		ImgUrl:            "https:" + img_src,
		SourceUrl:         link,
		SpecScore:         int32(spec_score_int),
		Ram:               convertString(device["RAM"]),
		Processor:         convertString(device["Processor"]),
		FrontCamera:       convertString(device["Front Camera"]),
		RearCamera:        convertString(device["Rear Camera"]),
		Battery:           convertString(device["Battery"]),
		Display:           convertString(device["Display"]),
		OperatingSystem:   convertString(device["Operating System"]),
		CustomUi:          convertString(device["Custom UI"]),
		Chipset:           convertString(device["Chipset"]),
		Cpu:               convertString(device["CPU"]),
		Architecture:      convertString(device["Architecture"]),
		Graphics:          convertString(device["Graphics"]),
		DisplayType:       convertString(device["Display Type"]),
		ScreenSize:        convertString(device["Screen Size"]),
		Resolution:        convertString(device["Resolution"]),
		PixelDensity:      convertString(device["Pixel Density"]),
		Touchscreen:       convertString(device["Touch Screen"]),
		InternalMemory:    convertString(device["Internal Memory"]),
		ExpandableMemory:  convertString(device["Expandable Memory"]),
		MCameraSetup:      convertString(device["Camera Setup"]),
		MResolution:       convertString(device["m Resolution"]),
		MAutofocus:        convertString(device["Autofocus"]),
		MOis:              convertString(device["OIS"]),
		MSensors:          convertString(device["Sensor"]),
		MFlash:            convertString(device["Flash"]),
		MImageResolution:  convertString(device["Image Flash"]),
		MSettings:         convertString(device["Settings"]),
		MShootingModes:    convertString(device["Shooting Modes"]),
		MCameraFeatures:   convertString(device["Camera Features"]),
		MVideoRecording:   convertString(device["Video Recording"]),
		SCameraSetup:      convertString(device["m Camera Setup"]),
		SResolution:       convertString(device["t Resolution"]),
		SVideoRecording:   convertString(device["m Video Recording"]),
		Capacity:          convertString(device["Capacity"]),
		RemovableBattery:  convertString(device["Removable"]),
		WirelessCharging:  convertString(device["Wireless Charging"]),
		QuickCharging:     convertString(device["Quick Charging"]),
		Usb:               convertString(device["USB Type-C"]),
		SimSlots:          convertString(device["SIM Slot(s)"]),
		NetworkSupport:    convertString(device["Network Support"]),
		FingerprintSensor: convertString(device["Fingerprint Sensor"]),
		OtherSensors:      convertString(device["Other Sensors"]),
		ScrapeTimestamp:   time.Now(),
	})
	if err != nil {
		log.Printf("ERROR:: inserting entry to database")
		return
	}
}
func convertString(value string) sql.NullString {
	if value != "" && len(value) > 0 {
		return sql.NullString{
			String: value,
			Valid:  true,
		}
	}
	return sql.NullString{
		Valid: false,
	}
}

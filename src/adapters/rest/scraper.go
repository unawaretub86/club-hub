package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/likexian/whois"

	"github.com/unawaretub86/club-hub/src/core/domain"
	"github.com/unawaretub86/club-hub/src/core/ports"
)

type Scrapper struct {
	scrapperPort ports.ScrapperPort
}

func NewScrapper(scrapperPort ports.ScrapperPort) *Scrapper {
	return &Scrapper{
		scrapperPort,
	}
}

func (scrapper *Scrapper) ScrapCompanyData(franchises []domain.Franchise) ([]domain.FranchiseScrapData, error) {
	var wg sync.WaitGroup
	franchiseResultChan := make(chan domain.FranchiseScrapData, len(franchises))
	domainResultChan := make(chan domain.DomainScrapData)

	for _, franchise := range franchises {
		wg.Add(1)
		go func(franchise domain.Franchise) {
			defer wg.Done()

			url := addHTTPScheme(*franchise.URL)

			franchiseInfo, err := getFranchiseInfo(url)
			if err != nil {
				log.Println("Error getting data from Franchise:", err)
				return
			}

			franchiseResultChan <- franchiseInfo
		}(franchise)

		wg.Add(1)
		go func(franchise domain.Franchise) {
			defer wg.Done()

			url := addHTTPScheme(*franchise.URL)

			domainInfo, err := getDomainInfo(url)
			if err != nil {
				log.Println("Error getting data from Domain:", err)
				return
			}

			domainResultChan <- domainInfo
		}(franchise)
	}

	go func() {
		wg.Wait()
		close(franchiseResultChan)
		close(domainResultChan)
	}()

	franchiseInfoList := []domain.FranchiseScrapData{}
	for info := range franchiseResultChan {
		info.DomainScrapData = <-domainResultChan
		franchiseInfoList = append(franchiseInfoList, info)
	}

	return franchiseInfoList, nil
}

func getFranchiseInfo(franchiseURL string) (domain.FranchiseScrapData, error) {
	info := &domain.FranchiseScrapData{}

	info.ImageURL = getImageURL(franchiseURL)

	hopCount, servers, status, protocol, err := getDataFromSSL(franchiseURL)
	if err != nil {
		return domain.FranchiseScrapData{}, err
	}

	info.HopCount = hopCount
	info.Servers = servers
	info.Status = *status
	info.CommunicationType = *protocol

	return *info, nil
}

func getImageURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error getting web:", err)
		return ""
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error analyzing web:", err)
		return ""
	}

	imageURL, exists := doc.Find("meta[property='og:image']").Attr("content")
	if !exists {
		log.Println("property 'og:image' not found")
		return ""
	}

	return imageURL
}

func getDataFromSSL(url string) (int, []string, *string, *string, error) {
	resp, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + url)
	if err != nil {
		return 0, nil, nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, nil, nil, err
	}

	var sslLabsResponse domain.SSLDomainScrapData
	err = json.Unmarshal(body, &sslLabsResponse)
	if err != nil {
		return 0, nil, nil, nil, err
	}

	var hopCount int

	var servers []string
	for _, endpoint := range sslLabsResponse.Servers {
		servers = append(servers, endpoint.ServerName)
		hopCount = len(servers)
	}

	status := sslLabsResponse.Status
	protocol := sslLabsResponse.Protocol

	return hopCount, servers, &status, &protocol, nil
}

func getDomainInfo(url string) (domain.DomainScrapData, error) {

	franchiseURLNoProtocol := removeHTTP(url)

	whoisResult, err := whois.Whois(franchiseURLNoProtocol)
	if err != nil {
		return domain.DomainScrapData{}, err
	}

	createdDate, err := parseValue(whoisResult, "Creation Date:")
	if err != nil {
		return domain.DomainScrapData{}, err
	}

	expiresDate, err := parseValue(whoisResult, "Expiration Date:")
	if err != nil {
		return domain.DomainScrapData{}, err
	}

	registrant, err := parseValue(whoisResult, "Registrant Name:")
	if err != nil {
		return domain.DomainScrapData{}, err
	}

	contactEmail, err := parseValue(whoisResult, "Tech Email:")
	if err != nil {
		return domain.DomainScrapData{}, err
	}

	return domain.DomainScrapData{
		CreatedAt:    createdDate,
		ExpiresAt:    expiresDate,
		Registrant:   registrant,
		ContactEmail: contactEmail,
	}, nil
}

func addHTTPScheme(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	return url
}

func removeHTTP(url string) string {
	re := regexp.MustCompile(`^https://www\.`)
	return re.ReplaceAllString(url, "")
}

func parseValue(whoisResult, key string) (string, error) {
	index := strings.Index(whoisResult, key)
	if index == -1 {
		return "", fmt.Errorf("clave no encontrada: %s", key)
	}

	start := index + len(key)
	end := strings.Index(whoisResult[start:], "\n") + start

	return strings.TrimSpace(whoisResult[start:end]), nil
}

package rest

import (
	"fmt"
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
				log.Println("Error obteniendo información de la franquicia:", err)
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
				log.Println("Error obteniendo información para la franquicia:", err)
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

func addHTTPScheme(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	return url
}

func removeHTTPWWW(url string) string {
	re := regexp.MustCompile(`^https://www\.`)
	return re.ReplaceAllString(url, "")
}

func getFranchiseInfo(franchiseURL string) (domain.FranchiseScrapData, error) {
	info := &domain.FranchiseScrapData{}

	info.ImageURL = getImageURL(franchiseURL)

	info.IsWebsiteUp = isWebsiteUp(franchiseURL)

	info.CommunicationType = getCommunicationType(franchiseURL)

	hopCount, _, err := getHopCountAndServers(franchiseURL)
	if err != nil {
		return domain.FranchiseScrapData{}, err
	}

	info.HopCount = hopCount

	return *info, nil
}

func getImageURL(url string) string {
	// Implementa el web scraping aquí para obtener la URL de la imagen.
	// Puedes usar goquery u otra biblioteca de scraping de tu elección.

	// Ejemplo de implementación con goquery:
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error al obtener la página web:", err)
		return ""
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error al analizar la página web:", err)
		return ""
	}

	imageURL, exists := doc.Find("meta[property='og:image']").Attr("content")
	if !exists {
		log.Println("No se encontró la etiqueta meta con propiedad 'og:image'")
		return ""
	}

	return imageURL
}

func isWebsiteUp(url string) bool {
	// Implementa la verificación del estado del sitio web aquí.
	// Puedes usar la librería net/http para enviar una solicitud HTTP y verificar el código de respuesta.
	resp, err := http.Head(url)
	if err != nil {
		log.Println("Error al enviar la solicitud HEAD:", err)
		return false
	}

	return resp.StatusCode == http.StatusOK
}

func getCommunicationType(url string) string {
	if strings.HasPrefix(url, "https://") {
		return "HTTPS"
	}

	return "HTTP"
}

func getHopCountAndServers(url string) (int, []string, error) {
	resp, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + url)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	// ... procesar la respuesta de la API de SSL Labs y extraer la información necesaria.

	return 3, []string{"Server1", "Server2", "Server3"}, nil
}

func getDomainInfo(url string) (domain.DomainScrapData, error) {

	franchiseURLNoProtocol := removeHTTPWWW(url)

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

func parseValue(whoisResult, key string) (string, error) {
	index := strings.Index(whoisResult, key)
	if index == -1 {
		return "", fmt.Errorf("clave no encontrada: %s", key)
	}

	start := index + len(key)
	end := strings.Index(whoisResult[start:], "\n") + start

	return strings.TrimSpace(whoisResult[start:end]), nil
}

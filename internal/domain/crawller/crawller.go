package crawller

import (
	"log"
	"regexp"
	"strings"

	"github.com/itelitte/ofsistorage-work/internal/domain"
	"github.com/itelitte/ofsistorage-work/internal/utils"

	"github.com/gocolly/colly"
)

// pega os dados da pagina https://docs.fcdo.gov.uk/docs/UK-Sanctions-List.html

func GetNamesEntries() []domain.SanctionEntry {
	// Crie um novo coletor
	c := colly.NewCollector()

	// Defina o seletor CSS para os elementos <li>
	liSelector := "li.SpacedOut"

	// Crie uma fatia para armazenar os textos dos elementos <li>
	var liTexts []string

	// Configure o manipulador para processar os elementos <li>
	c.OnHTML(liSelector, func(e *colly.HTMLElement) {
		// Extraia o texto do elemento <li> e adicione à fatia
		liTexts = append(liTexts, e.Text)
	})

	// Visite a página da web e colete os dados
	err := c.Visit("https://ofsistorage.blob.core.windows.net/publishlive/2022format/ConList.html")
	if err != nil {
		log.Fatal(err)
	}

	var entries []domain.SanctionEntry

	// Imprima os textos dos elementos <li>
	for _, text := range liTexts {
		//fmt.Println(text)
		if strings.Contains(text, "Name 6:") {
			e := extractSanctionEntry(text)
			if len(e.Names) != 0 {
				entries = append(entries, e)
			}

		}
	}

	return entries
}

func GetNamesOrganization() []domain.OrganizationEntry {
	// Crie um novo coletor
	c := colly.NewCollector()

	// Defina o seletor CSS para os elementos <li>
	liSelector := "li.SpacedOut"

	// Crie uma fatia para armazenar os textos dos elementos <li>
	var liTexts []string

	// Configure o manipulador para processar os elementos <li>
	c.OnHTML(liSelector, func(e *colly.HTMLElement) {
		// Extraia o texto do elemento <li> e adicione à fatia
		liTexts = append(liTexts, e.Text)
	})

	// Visite a página da web e colete os dados
	err := c.Visit("https://ofsistorage.blob.core.windows.net/publishlive/2022format/ConList.html")
	if err != nil {
		log.Fatal(err)
	}

	var organization []domain.OrganizationEntry
	// Imprima os textos dos elementos <li>
	for _, text := range liTexts {
		//fmt.Println(text)
		if strings.Contains(text, "Organisation Name:") {
			organization = append(organization, extractOrganizationEntry(text))
		}
	}

	return organization
}

func extractSanctionEntry(input string) domain.SanctionEntry {

	reTexto := regexp.MustCompile(`\x{00A0}`)
	data := reTexto.ReplaceAllString(input, " ")
	entry := domain.SanctionEntry{}
	// Use expressões regulares para extrair as informações necessárias
	reName := regexp.MustCompile(`Name 6:(.+?)(?:Name|Title|DOB)`)
	reNameNonLatin := regexp.MustCompile(`Name \(non-Latin script\): (.+?)(?:DOB|Title)`)
	reDOB := regexp.MustCompile(`DOB: (.+)POB`)
	rePOB := regexp.MustCompile(`POB: (.+)Good`)
	reAKA := regexp.MustCompile(`a\.k\.a:\s*(.+?)\s*(?:Other|Nationality|National|Passport)`)
	reOtherInformation := regexp.MustCompile(`Other Information: (.+)Listed`)
	reListedOn := regexp.MustCompile(`Listed on: (\S+)`)
	reSanctionsListDate := regexp.MustCompile(`UK Sanctions List Date Designated: (\S+)`)
	reLastUpdated := regexp.MustCompile(`Last Updated: (\S+)`)
	reGroupID := regexp.MustCompile(`Group ID: (\S+)`)

	// Use FindStringSubmatch para encontrar a primeira correspondência de cada padrão
	if matches := reName.FindStringSubmatch(data); len(matches) > 1 {
		entry.Names = utils.ExtractValues(matches[1])
	}

	if matches := reAKA.FindStringSubmatch(data); len(matches) > 1 {
		entry.AKA = matches[1]
		entry.AKA = strings.ReplaceAll(entry.AKA, "Low quality a.k.a:", ",")
	}

	if matches := reNameNonLatin.FindStringSubmatch(data); len(matches) > 1 {
		entry.NameNonLatin = matches[1]
	}

	if matches := reDOB.FindStringSubmatch(data); len(matches) > 1 {
		entry.DOB = matches[1]
	}

	if matches := rePOB.FindStringSubmatch(data); len(matches) > 1 {
		entry.POB = matches[1]
	}

	if matches := reOtherInformation.FindStringSubmatch(data); len(matches) > 1 {
		entry.OtherInformation = matches[1]
	}

	if matches := reListedOn.FindStringSubmatch(data); len(matches) > 1 {
		entry.ListedOn = matches[1]
	}

	if matches := reSanctionsListDate.FindStringSubmatch(data); len(matches) > 1 {
		entry.SanctionsListDate = matches[1]
	}

	if matches := reLastUpdated.FindStringSubmatch(data); len(matches) > 1 {
		entry.LastUpdated = matches[1]
	}

	if matches := reGroupID.FindStringSubmatch(data); len(matches) > 1 {
		entry.GroupID = matches[1]
	}

	return entry
}

func extractOrganizationEntry(data string) domain.OrganizationEntry {
	entry := domain.OrganizationEntry{}

	// Use expressões regulares para extrair as informações necessárias
	reOrganizationName := regexp.MustCompile(`Organisation Name:\s+(.+?)\s*(?:a\.k\.a|\(non-Latin script\)|Address:|Other Information:|NameNonLatin|$)`)
	reNameNonLatin := regexp.MustCompile(`Name \(non-Latin script\):\s+(.+)`)
	reAKA := regexp.MustCompile(`a\.k\.a:\s+(.+)`)
	reAddress := regexp.MustCompile(`Address:\s+(.+)`)
	reOtherInformation := regexp.MustCompile(`Other Information:\s+(.+)`)
	reListedOn := regexp.MustCompile(`Listed on:\s+(\S+)`)
	reSanctionsListDate := regexp.MustCompile(`UK Sanctions List Date Designated:\s+(\S+)`)
	reLastUpdated := regexp.MustCompile(`Last Updated:\s+(\S+)`)
	reGroupID := regexp.MustCompile(`Group ID:\s+(\S+)`)

	if matches := reOrganizationName.FindStringSubmatch(data); len(matches) > 1 {
		entry.OrganizationName = reomoveWordName(matches[1])
	}

	// Extrair a.k.a. (se houver)
	if matches := reAKA.FindAllStringSubmatch(data, -1); len(matches) > 0 {
		for _, m := range matches {
			entry.AKA = append(entry.AKA, m[1])
		}
	}

	// Extrair nome (não em script latino) (se houver)
	if matches := reNameNonLatin.FindStringSubmatch(data); len(matches) > 1 {
		entry.NameNonLatin = matches[1]
	}

	// Extrair o endereço (se houver)
	if matches := reAddress.FindAllStringSubmatch(data, -1); len(matches) > 0 {
		for _, m := range matches {
			entry.Address = append(entry.Address, m[1])
		}
	}

	if matches := reOtherInformation.FindStringSubmatch(data); len(matches) > 1 {
		entry.OtherInformation = matches[1]
	}

	if matches := reListedOn.FindStringSubmatch(data); len(matches) > 1 {
		entry.ListedOn = matches[1]
	}

	if matches := reSanctionsListDate.FindStringSubmatch(data); len(matches) > 1 {
		entry.SanctionsListDate = matches[1]
	}

	if matches := reLastUpdated.FindStringSubmatch(data); len(matches) > 1 {
		entry.LastUpdated = matches[1]
	}

	if matches := reGroupID.FindStringSubmatch(data); len(matches) > 1 {
		entry.GroupID = matches[1]
	}

	return entry
}

// imput OJE PARVAZ MADO NAFAR COMPANYName
// output OJE PARVAZ MADO NAFAR COMPANY
func reomoveWordName(data string) string {
	reName := regexp.MustCompile(`Name`)
	return reName.ReplaceAllString(data, "")
}

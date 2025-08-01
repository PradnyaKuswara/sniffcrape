package services

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"github.com/PradnyaKuswara/sniffcrape/internal/repositories"
	"github.com/gocolly/colly/v2"
)

type ScrapeResultService struct {
	ScrapeResultRepositoryInterface repositories.ScrapeResultRepositoryInterface
}

func NewScrapeResultService(repo repositories.ScrapeResultRepositoryInterface) *ScrapeResultService {
	return &ScrapeResultService{
		ScrapeResultRepositoryInterface: repo,
	}
}

func (s *ScrapeResultService) GetAllScrapeResults() ([]models.ScrapeResult, error) {
	scrapeResults, err := s.ScrapeResultRepositoryInterface.GetAllScrapeResults()
	if err != nil {
		return nil, err
	}
	return scrapeResults, nil
}

func (s *ScrapeResultService) GetScrapeResultByID(id string) (models.ScrapeResult, error) {
	scrapeResult, err := s.ScrapeResultRepositoryInterface.GetScrapeResultByID(id)
	if err != nil {
		return models.ScrapeResult{}, err
	}
	return scrapeResult, nil
}

func (s *ScrapeResultService) CreateScrapeResult(req models.ScrapeResultRequest) (models.ScrapeResult, error) {
	data := models.ScrapeResult{
		Title:       req.Title,
		Url:         req.Url,
		Description: req.Description,
	}

	result, err := s.ScrapeResultRepositoryInterface.CreateScrapeResult(data)
	if err != nil {
		return models.ScrapeResult{}, err
	}

	return result, nil
}

func (s *ScrapeResultService) CollyScrap(url string) (*models.ScrapeResultResponse, error) {
	c := colly.NewCollector()

	var (
		title, description, keywords, canonical string
		ogTitle, ogDesc, ogImage, ogUrl         string
		twTitle, twDesc, twImage, twCard        string
		robots, viewport, charset, author       string

		h1s, h2s, h3s []string
		imageUrls     []string
		links         []string
		favicons      []string
		scripts       []string
	)

	// Title
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	// Meta Tags
	c.OnHTML(`meta[name="description"]`, func(e *colly.HTMLElement) {
		description = e.Attr("content")
	})
	c.OnHTML(`meta[name="keywords"]`, func(e *colly.HTMLElement) {
		keywords = e.Attr("content")
	})
	c.OnHTML(`meta[name="robots"]`, func(e *colly.HTMLElement) {
		robots = e.Attr("content")
	})
	c.OnHTML(`meta[name="viewport"]`, func(e *colly.HTMLElement) {
		viewport = e.Attr("content")
	})
	c.OnHTML(`meta[charset]`, func(e *colly.HTMLElement) {
		charset = e.Attr("charset")
	})
	c.OnHTML(`meta[name="author"]`, func(e *colly.HTMLElement) {
		author = e.Attr("content")
	})

	// Canonical
	c.OnHTML(`link[rel="canonical"]`, func(e *colly.HTMLElement) {
		canonical = e.Attr("href")
	})

	// Open Graph Tags
	c.OnHTML(`meta[property="og:title"]`, func(e *colly.HTMLElement) {
		ogTitle = e.Attr("content")
	})
	c.OnHTML(`meta[property="og:description"]`, func(e *colly.HTMLElement) {
		ogDesc = e.Attr("content")
	})
	c.OnHTML(`meta[property="og:image"]`, func(e *colly.HTMLElement) {
		ogImage = e.Attr("content")
	})
	c.OnHTML(`meta[property="og:url"]`, func(e *colly.HTMLElement) {
		ogUrl = e.Attr("content")
	})

	// Twitter Card Tags
	c.OnHTML(`meta[name="twitter:title"]`, func(e *colly.HTMLElement) {
		twTitle = e.Attr("content")
	})
	c.OnHTML(`meta[name="twitter:description"]`, func(e *colly.HTMLElement) {
		twDesc = e.Attr("content")
	})
	c.OnHTML(`meta[name="twitter:image"]`, func(e *colly.HTMLElement) {
		twImage = e.Attr("content")
	})
	c.OnHTML(`meta[name="twitter:card"]`, func(e *colly.HTMLElement) {
		twCard = e.Attr("content")
	})

	// Headings
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		h1s = append(h1s, e.Text)
	})
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		h2s = append(h2s, e.Text)
	})
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		h3s = append(h3s, e.Text)
	})

	// Images
	c.OnHTML("img", func(e *colly.HTMLElement) {
		src := e.Request.AbsoluteURL(e.Attr("src"))
		if src != "" {
			imageUrls = append(imageUrls, src)
		}
	})

	// Links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Request.AbsoluteURL(e.Attr("href"))
		if href != "" {
			links = append(links, href)
		}
	})

	// Favicons
	c.OnHTML(`link[rel="icon"], link[rel="shortcut icon"], link[rel="apple-touch-icon"]`, func(e *colly.HTMLElement) {
		icon := e.Request.AbsoluteURL(e.Attr("href"))
		if icon != "" {
			favicons = append(favicons, icon)
		}
	})

	// External scripts
	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		src := e.Request.AbsoluteURL(e.Attr("src"))
		if src != "" {
			scripts = append(scripts, src)
		}
	})

	// Start scraping
	if err := c.Visit(url); err != nil {
		return nil, err
	}

	return &models.ScrapeResultResponse{
		Title:       title,
		Url:         url,
		Description: description,
		Keywords:    keywords,
		Canonical:   canonical,
		Robots:      robots,
		Viewport:    viewport,
		Charset:     charset,
		Author:      author,

		OgTitle: ogTitle,
		OgDesc:  ogDesc,
		OgImage: ogImage,
		OgUrl:   ogUrl,

		TwTitle: twTitle,
		TwDesc:  twDesc,
		TwImage: twImage,
		TwCard:  twCard,

		H1:     h1s,
		H2:     h2s,
		H3:     h3s,
		Images: imageUrls,
		Links:  links,

		Favicons: favicons,
		Scripts:  scripts,
	}, nil
}

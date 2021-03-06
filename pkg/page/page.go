package page

import (
	"encoding/json"
	"fmt"
	"gocms/pkg/files"
	"io/ioutil"
	"time"
)

// Status type
type Status int

// Type type
type Type int

// Draft or Published status constants
const (
	Draft Status = iota
	Published
)

// Post or Page type constants
const (
	Page Type = iota
	Post
)

// Info type
type Info struct {
	ID          int
	Parent      int
	Depth       int
	Title       string
	Description string
	Content     string
	Route       string
	Author      string
	Status      Status
	Category    int
	Type        Type
	Template    string
	PubDate     time.Time
	UpdateDate  time.Time
}

// GetByID by ID a page or post
func GetByID(domain string, id int, getContent bool) Info {
	return find(domain, id, "-", getContent)
}

// GetByRoute a page or post
func GetByRoute(domain string, route string, getContent bool) Info {
	return find(domain, -1, route, getContent)
}

var pages []Info

func find(domain string, id int, route string, getContent bool) Info {
	var page Info
	if pages == nil {
		LoadData(domain)
	}
	for _, p := range pages {
		if p.ID == id || p.Route == route {
			if getContent {
				p.Content = LoadContent(domain, p.ID)
			}
			page = p
			break
		}
	}
	return page
}

// LoadContent return the contents of the page content file
func LoadContent(domain string, id int) string {
	content, _ := ioutil.ReadFile(fmt.Sprintf("%s%s/pages/%d.html", files.Root, domain, id))
	return string(content)
}

// SaveContent saves page content to a file
func SaveContent(domain string, id int, content string) {
	err := ioutil.WriteFile(fmt.Sprintf("%s%s/pages/%d.html", files.Root, domain, id), []byte(content), 0660)
	if err != nil {
		panic(err)
	}
}

const pagesFile = "/data/pages.json"

// LoadData loads the data from the pages data file
func LoadData(domain string) {
	data, err := ioutil.ReadFile(files.Root + domain + pagesFile)
	if err != nil {
		pages = []Info{Info{}}
		SaveData(domain)
	} else {
		err = json.Unmarshal(data, &pages)
		if err != nil {
			panic(err)
		}
	}
}

// SaveData saves the pages data to a file
func SaveData(domain string) {
	b, _ := json.Marshal(pages)
	err := ioutil.WriteFile(files.Root+domain+pagesFile, b, 0660)
	if err != nil {
		panic(err)
	}
}

// Save a page or post
func Save(domain string, info Info, saveContent bool) int {
	if saveContent {
		SaveContent(domain, info.ID, info.Content)
	}
	return 1
}

// GetPages returns a slice of pages data
func GetPages(parent int, status Status) []Info {
	list := []Info{}
	for _, p := range pages {
		if p.Parent == parent && p.Status == status && p.Template == "post" || p.Template == "page" {
			list = append(list, p)
		}
	}
	return list
}

// GetAllPages return all the page data
func GetAllPages() []Info {
	return pages
}

// GetCategoryPages returns a slice of category pages in a category
func GetCategoryPages(id int, status Status) []Info {
	list := []Info{}
	for _, p := range pages {
		if p.Category == id && p.Status == status && p.Template == "category" {
			list = append(list, p)
		}
	}
	return list
}

// Add new page data
func Add(newPage Info) {
	pages = append(pages, newPage)
}

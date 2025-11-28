package readarr_test

import (
	"net/http"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golift.io/starr"
	"golift.io/starr/readarr"
	"golift.io/starr/starrtest"
)

var testSearchJSON = `{
		"foreignId": "3634639",
		"book": {
			"title": "Dune",
			"authorTitle": "herbert, frank Dune",
			"seriesTitle": "",
			"disambiguation": "",
			"overview": "Set on the desert planet Arrakis, Dune is the story of the boy Paul Atreides, heir to a noble family tasked with ruling an inhospitable world where the only thing of value is the “spice” melange, a drug capable of extending life and enhancing consciousness. Coveted across the known universe, melange is a prize worth killing for...\r\n\r\nWhen House Atreides is betrayed, the destruction of Paul’s family will set the boy on a journey toward a destiny greater than he could ever have imagined. And as he evolves into the mysterious man known as Muad’Dib, he will bring to fruition humankind’s most ancient and unattainable dream.",
			"authorId": 53,
			"foreignBookId": "3634639",
			"foreignEditionId": "44767458",
			"titleSlug": "3634639",
			"monitored": true,
			"anyEditionOk": true,
			"ratings": {
				"votes": 784771,
				"value": 4.33,
				"popularity": 3398058.43
			},
			"releaseDate": "1965-06-01T07:00:00Z",
			"pageCount": 658,
			"genres": [
				"Science Fiction",
				"Fiction",
				"Fantasy",
				"Classics",
				"Audiobook",
				"Science Fiction Fantasy",
				"Novels",
				"Space Opera",
				"Adventure",
				"Book Club"
			],
			"author": {
				"authorMetadataId": 53,
				"status": "continuing",
				"ended": false,
				"authorName": "Frank Herbert",
				"authorNameLastFirst": "Herbert, Frank",
				"foreignAuthorId": "58",
				"titleSlug": "58",
				"overview": "Franklin Patrick Herbert Jr. was an American science fiction author best known for the 1965 novel Dune and its five sequels. Though he became famous for his novels, he also wrote short stories and worked as a newspaper journalist, photographer, book reviewer, ecological consultant, and lecturer.The Dune saga, set in the distant future, and taking place over millennia, explores complex themes, such as the long-term survival of the human species, human evolution, planetary science and ecology, and the intersection of religion, politics, economics and power in a future where humanity has long since developed interstellar travel and settled many thousands of worlds. Dune is the best-selling science fiction novel of all time, and the entire series is considered to be among the classics of the genre.",
				"links": [
					{
						"url": "https://www.goodreads.com/author/show/58.Frank_Herbert",
						"name": "Goodreads"
					}
				],
				"images": [
					{
						"url": "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/authors/1591018335i/58._UX200_CR0,0,200,200_.jpg",
						"coverType": "poster",
						"extension": ".jpg"
					}
				],
				"path": "/books/Frank Herbert",
				"qualityProfileId": 1,
				"metadataProfileId": 2,
				"monitored": true,
				"monitorNewItems": "none",
				"folder": "Frank Herbert",
				"genres": [],
				"cleanName": "frankherbert",
				"sortName": "frank herbert",
				"sortNameLastFirst": "herbert, frank",
				"tags": [],
				"added": "2025-11-21T08:58:56Z",
				"ratings": {
					"votes": 1446269,
					"value": 4.154275,
					"popularity": 6008199.149975
				},
				"statistics": {
					"bookFileCount": 0,
					"bookCount": 0,
					"availableBookCount": 0,
					"totalBookCount": 0,
					"sizeOnDisk": 0,
					"percentOfBooks": 0
				},
				"id": 53
			},
			"images": [
				{
					"url": "/MediaCover/Books/187/cover.jpg?lastWrite=636992634250000000",
					"coverType": "cover",
					"extension": ".jpg",
					"remoteUrl": "https://m.media-amazon.com/images/S/compressed.photo.goodreads.com/books/1555447414i/44767458.jpg"
				}
			],
			"links": [
				{
					"url": "https://www.goodreads.com/work/3634639-dune",
					"name": "Goodreads Editions"
				},
				{
					"url": "https://www.goodreads.com/book/show/44767458-dune",
					"name": "Goodreads Book"
				}
			],
			"added": "2025-11-21T08:58:56Z",
			"remoteCover": "https://m.media-amazon.com/images/S/compressed.photo.goodreads.com/books/1555447414i/44767458.jpg",
			"editions": [
				{
					"bookId": 187,
					"foreignEditionId": "44767458",
					"titleSlug": "44767458",
					"isbn13": "9780593099322",
					"asin": "059309932X",
					"title": "Dune",
					"language": "eng",
					"overview": "Set on the desert planet Arrakis, Dune is the story of the boy Paul Atreides, heir to a noble family tasked with ruling an inhospitable world where the only thing of value is the “spice” melange, a drug capable of extending life and enhancing consciousness. Coveted across the known universe, melange is a prize worth killing for...\r\n\r\nWhen House Atreides is betrayed, the destruction of Paul’s family will set the boy on a journey toward a destiny greater than he could ever have imagined. And as he evolves into the mysterious man known as Muad’Dib, he will bring to fruition humankind’s most ancient and unattainable dream.",
					"format": "Hardcover",
					"isEbook": false,
					"disambiguation": "",
					"publisher": "Ace",
					"pageCount": 658,
					"releaseDate": "2019-10-01T07:00:00Z",
					"images": [
						{
							"url": "/MediaCover/Books/187/cover.jpg?lastWrite=636992634250000000",
							"coverType": "cover",
							"extension": ".jpg",
							"remoteUrl": "https://m.media-amazon.com/images/S/compressed.photo.goodreads.com/books/1555447414i/44767458.jpg"
						}
					],
					"links": [
						{
							"url": "https://www.goodreads.com/book/show/44767458-dune",
							"name": "Goodreads Book"
						}
					],
					"ratings": {
						"votes": 784771,
						"value": 4.33,
						"popularity": 3398058.43
					},
					"monitored": true,
					"manualAdd": false,
					"grabbed": false,
					"id": 187
				}
			],
			"grabbed": false,
			"id": 187
		},
		"id": 2
}`

var testEditionImage = starr.Image{
	CoverType: "cover",
	URL:       "/MediaCover/Books/187/cover.jpg?lastWrite=636992634250000000",
	RemoteURL: "https://m.media-amazon.com/images/S/compressed.photo.goodreads.com/books/1555447414i/44767458.jpg",
	Extension: ".jpg",
}

var testBookImage = testEditionImage

var testAuthorImage = starr.Image{
	CoverType: "poster",
	URL:       "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/authors/1591018335i/58._UX200_CR0,0,200,200_.jpg",
	Extension: ".jpg",
}

var testAuthorLink = starr.Link{
	URL:  "https://www.goodreads.com/author/show/58.Frank_Herbert",
	Name: "Goodreads",
}

var testEditionLink = starr.Link{
	URL:  "https://www.goodreads.com/book/show/44767458-dune",
	Name: "Goodreads Book",
}

var testBookLink1 = starr.Link{
	URL:  "https://www.goodreads.com/work/3634639-dune",
	Name: "Goodreads Editions",
}

var testBookLink2 = testEditionLink

var testBookRating = starr.Ratings{
	Votes:      784771,
	Value:      4.33,
	Popularity: 3398058.43,
}

var testAuthorRating = starr.Ratings{
	Votes:      1446269,
	Value:      4.154275,
	Popularity: 6008199.149975,
}

var testEditionRating = testBookRating

var testEdition = readarr.Edition{
	ID:               187,
	BookID:           187,
	ForeignEditionID: "44767458",
	TitleSlug:        "44767458",
	Isbn13:           "9780593099322",
	Asin:             "059309932X",
	Title:            "Dune",
	Overview:         "Set on the desert planet Arrakis, Dune is the story of the boy Paul Atreides, heir to a noble family tasked with ruling an inhospitable world where the only thing of value is the “spice” melange, a drug capable of extending life and enhancing consciousness. Coveted across the known universe, melange is a prize worth killing for...\r\n\r\nWhen House Atreides is betrayed, the destruction of Paul’s family will set the boy on a journey toward a destiny greater than he could ever have imagined. And as he evolves into the mysterious man known as Muad’Dib, he will bring to fruition humankind’s most ancient and unattainable dream.",
	Format:           "Hardcover",
	Publisher:        "Ace",
	PageCount:        658,
	ReleaseDate:      time.Date(2019, 10, 1, 7, 0, 0, 0, time.UTC),
	Images:           []*starr.Image{&testEditionImage},
	Links:            []*starr.Link{&testEditionLink},
	Ratings:          &testEditionRating,
	Monitored:        true,
	ManualAdd:        false,
	IsEbook:          false,
}

var testStatistics = readarr.Statistics{
	BookCount:          0,
	BookFileCount:      0,
	TotalBookCount:     0,
	SizeOnDisk:         0,
	PercentOfBooks:     0,
	AvailableBookCount: 0,
}

var testAuthor = readarr.Author{
	ID:                  53,
	Status:              "continuing",
	AuthorName:          "Frank Herbert",
	ForeignAuthorID:     "58",
	TitleSlug:           "58",
	Overview:            "Franklin Patrick Herbert Jr. was an American science fiction author best known for the 1965 novel Dune and its five sequels. Though he became famous for his novels, he also wrote short stories and worked as a newspaper journalist, photographer, book reviewer, ecological consultant, and lecturer.The Dune saga, set in the distant future, and taking place over millennia, explores complex themes, such as the long-term survival of the human species, human evolution, planetary science and ecology, and the intersection of religion, politics, economics and power in a future where humanity has long since developed interstellar travel and settled many thousands of worlds. Dune is the best-selling science fiction novel of all time, and the entire series is considered to be among the classics of the genre.",
	Links:               []*starr.Link{&testAuthorLink},
	Images:              []*starr.Image{&testAuthorImage},
	Path:                "/books/Frank Herbert",
	QualityProfileID:    1,
	MetadataProfileID:   2,
	Genres:              []string{},
	CleanName:           "frankherbert",
	SortName:            "frank herbert",
	Tags:                []int{},
	Added:               time.Date(2025, 11, 21, 8, 58, 56, 0, time.UTC),
	Ratings:             &testAuthorRating,
	Statistics:          &testStatistics,
	LastBook:            nil,
	NextBook:            nil,
	Ended:               false,
	Monitored:           true,
	AuthorMetadataID:    53,
	AuthorNameLastFirst: "Herbert, Frank",
	MonitorNewItems:     "none",
	SortNameLastFirst:   "herbert, frank",
}

func TestSearch(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "search?term=Dune"),
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[` + testSearchJSON + `]`,
			WithRequest:    "Dune",
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse: []*readarr.SearchResult{{
				ForeignID: "3634639",
				ID:        2,
				Book: &readarr.Book{
					Added:         time.Date(2025, 11, 21, 8, 58, 56, 0, time.UTC),
					AnyEditionOk:  true,
					AuthorID:      53,
					AuthorTitle:   "herbert, frank Dune",
					Editions:      []*readarr.Edition{&testEdition},
					ForeignBookID: "3634639",
					Genres: []string{
						"Science Fiction",
						"Fiction",
						"Fantasy",
						"Classics",
						"Audiobook",
						"Science Fiction Fantasy",
						"Novels",
						"Space Opera",
						"Adventure",
						"Book Club",
					},
					ID:          187,
					Images:      []*starr.Image{&testBookImage},
					Links:       []*starr.Link{&testBookLink1, &testBookLink2},
					Monitored:   true,
					Grabbed:     false,
					Overview:    "Set on the desert planet Arrakis, Dune is the story of the boy Paul Atreides, heir to a noble family tasked with ruling an inhospitable world where the only thing of value is the “spice” melange, a drug capable of extending life and enhancing consciousness. Coveted across the known universe, melange is a prize worth killing for...\r\n\r\nWhen House Atreides is betrayed, the destruction of Paul’s family will set the boy on a journey toward a destiny greater than he could ever have imagined. And as he evolves into the mysterious man known as Muad’Dib, he will bring to fruition humankind’s most ancient and unattainable dream.",
					PageCount:   658,
					Ratings:     &testBookRating,
					RemoteCover: "https://m.media-amazon.com/images/S/compressed.photo.goodreads.com/books/1555447414i/44767458.jpg",
					ReleaseDate: time.Date(1965, 6, 1, 7, 0, 0, 0, time.UTC),
					SeriesTitle: "",
					Title:       "Dune",
					TitleSlug:   "3634639",
					Author:      &testAuthor,
				},
				Author: nil,
			}},
		},

		{
			Name:           "noresults",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "search?term=somestringthatisntabooktitle"),
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[]`,
			WithRequest:    "somestringthatisntabooktitle",
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*readarr.SearchResult{},
		},

		{
			Name:           "emptyterm",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "search?term="),
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[]`,
			WithRequest:    "",
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*readarr.SearchResult(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.Search(test.WithRequest.(string))
			require.ErrorIs(t, err, test.WithError, "wrong error")
			assert.EqualValues(t, test.WithResponse, output, "response mismatch")
		})
	}
}

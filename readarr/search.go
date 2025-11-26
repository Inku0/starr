package readarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpSearch = APIver + "/search"

// SearchResult is the struct returned from the /api/v1/search endpoint.
// ID in this context means the index of the search result, not the book's ID.
type SearchResult struct {
	ForeignID string  `json:"foreignId,omitempty"`
	Author    *Author `json:"author,omitempty"`
	Book      *Book   `json:"book,omitempty"`
	ID        int     `json:"id,omitempty"`
}

// Search returns a slice of pointers to SearchResult.
func (r *Readarr) Search(term string) ([]*SearchResult, error) {
	return r.SearchContext(context.Background(), term)
}

// SearchContext returns a slice of pointers to SearchResult.
func (r *Readarr) SearchContext(ctx context.Context, term string) ([]*SearchResult, error) {
	if term == "" {
		return nil, fmt.Errorf("search term can't be empty")
	}

	params := make(url.Values)
	params.Set("term", term)

	var output []*SearchResult

	req := starr.Request{URI: bpSearch, Query: params}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

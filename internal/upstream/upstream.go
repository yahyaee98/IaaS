package upstream

import (
	"iaas/internal/data"
)

// Upstream is meant to be a bridge between external libraries from third parties(in this case "itunes" and "google_books"
// in the "third_party" package). So, an Upstream will translate a third party's response to the struct that is meaningful
// to us.
type Upstream interface {
	Search(search string) ([]*data.Item, error)
}

package v2

import (
	"crypto"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Source is a single entry map of URI to hash or Git Reference
type Source map[string]string

// NewSource builds a new Source from a URI
func NewSource(URI string) (src Source, err error) {
	src = make(Source)
	var in io.ReadCloser
	if strings.HasPrefix(URI, "git|") {
		// Git sources
		pieces := strings.Split(URI, ":")
		if len(pieces) < 3 {
			return "", fmt.Errorf("no hash in git URI, should resemble 'git|http://path/to/repo:commit hash")
		}
		src[URI] = pieces[len(pieces)-1]
		return
	} else if strings.HasPrefix(URI, "file://") {
		// Local File sources
		in, err = os.Open(strings.TrimPrefix(URI, "file://"))
		if err != nil {
			return
		}
		defer in.Close()
	} else if strings.HasPrefix(URI, "http") {
		// HTTP Sources
		var r *http.Response
		r, err = http.Get(URI)
		if err != nil {
			return
		}
		in = r.Body
		defer r.Body.Close()
	} else {
		err = fmt.Errorf("unsupported source type")
		return
	}
	// All hashed are SHA256 hashes
	hash := crypto.SHA256.New()
	_, err = io.Copy(hash, in)
	if err != nil {
		return
	}
	src[URI] = fmt.Sprintf("%x", hash.Sum(nil))
	return
}

// UpdateSources gets the hashes for one or more URI sources
func UpdateSources(URIs []string) (srcs []Source, err error) {
	// for each URI
	var src Source
	for _, URI := range URIs {
		// - URI : HASH
		// Get Hash for URI
		src, err = NewSource(URI)
		if err != nil {
			return
		}
		src[URI] = h
		// Add source to all sources
		srcs = append(srcs, src)
	}
	return
}

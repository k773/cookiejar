package cookiejar

import (
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

// ExtractCookies : j must be of type *(github.com/k773/cookiejar).Jar
func ExtractCookies(j http.CookieJar, domainIncludes string) (m map[string]string, e error) {
	if c, s := j.(*Jar); s {
		m = map[string]string{}
		var m2 = map[string]Entry{}

		for domain, cookieName2Cookie := range c.Entries {
			if domainIncludes == "" || strings.Contains(domain, domainIncludes) {
				for k, v := range cookieName2Cookie {
					var prev Entry
					var alreadyWritten bool
					if prev, alreadyWritten = m2[k]; alreadyWritten {
						if v.Creation.After(prev.Creation) {
							alreadyWritten = false
						}
					}

					if !alreadyWritten {
						m2[k] = v
						m[v.Name] = v.Value
					}
				}
			}

		}
	} else {
		e = errors.New("ExtractCookies: unsupported cookiejar")
	}
	return
}

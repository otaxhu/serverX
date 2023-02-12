package serverX

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Metadata interface{}

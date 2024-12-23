package browser

import (
	"go.k6.io/k6/js/modules/k6/browser/common"
)

// syncMapRequest is like mapRequest but returns synchronous functions.
func syncMapRequest(vu moduleVU, r *common.Request) mapping {
	maps := mapping{
		"allHeaders": r.AllHeaders,
		"frame":      func() mapping { return syncMapFrame(vu, r.Frame()) },
		"headerValue": func(name string) any {
			v, ok := r.HeaderValue(name)
			if !ok {
				return nil
			}

			return v
		},
		"headers":             r.Headers,
		"headersArray":        r.HeadersArray,
		"isNavigationRequest": r.IsNavigationRequest,
		"method":              r.Method,
		"postData":            r.PostData,
		"postDataBuffer":      r.PostDataBuffer,
		"resourceType":        r.ResourceType,
		"response":            func() mapping { return syncMapResponse(vu, r.Response()) },
		"size":                r.Size,
		"timing":              r.Timing,
		"url":                 r.URL,
	}

	return maps
}

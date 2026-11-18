// Copyright 2023 The narqo/go-badge Authors. All rights reserved.
// SPDX-License-Identifier: MIT.

package badges

// cspell:words Verdana

// Style represents the visual style of the badge.
type Style string

const (
	// StyleFlat is the default style with rounded corners and a slight gradient.
	StyleFlat Style = "flat"
	// StyleFlatSquare has square corners and uses solid colors only.
	StyleFlatSquare Style = "flat-square"
)

// ParseStyle returns the style matching the given name or StyleFlat if the name is unknown.
func ParseStyle(name string) Style {
	switch Style(name) {
	case StyleFlatSquare:
		return StyleFlatSquare
	default:
		return StyleFlat
	}
}

var flatTemplate = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Bounds.Dx}}" height="20"><linearGradient id="smooth" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="round"><rect width="{{.Bounds.Dx}}" height="20" rx="3" fill="#fff"/></mask><g mask="url(#round)"><rect width="{{.Bounds.SubjectDx}}" height="20" fill="#555"/><rect x="{{.Bounds.SubjectDx}}" width="{{.Bounds.StatusDx}}" height="20" fill="{{or .Color "#4c1" | html}}"/><rect width="{{.Bounds.Dx}}" height="20" fill="url(#smooth)"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="{{.Bounds.SubjectX}}" y="15" fill="#010101" fill-opacity=".3">{{.Subject | html}}</text><text x="{{.Bounds.SubjectX}}" y="14">{{.Subject | html}}</text><text x="{{.Bounds.StatusX}}" y="15" fill="#010101" fill-opacity=".3">{{.Status | html}}</text><text x="{{.Bounds.StatusX}}" y="14">{{.Status | html}}</text></g></svg>`

var flatSquareTemplate = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Bounds.Dx}}" height="20"><g shape-rendering="crispEdges"><rect width="{{.Bounds.SubjectDx}}" height="20" fill="#555"/><rect x="{{.Bounds.SubjectDx}}" width="{{.Bounds.StatusDx}}" height="20" fill="{{or .Color "#4c1" | html}}"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="{{.Bounds.SubjectX}}" y="14">{{.Subject | html}}</text><text x="{{.Bounds.StatusX}}" y="14">{{.Status | html}}</text></g></svg>`

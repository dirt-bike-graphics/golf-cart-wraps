// Package cartmodels provides open reference data for 34 golf cart models
// and body configurations across Club Car, E-Z-GO, Yamaha, Evolution, ICON,
// Bintelli, and Denago EV.
//
// Published by [Golf Cart Wraps] to support wrap sizing research and body
// style comparisons. Preview patterns on these models in 3D with the
// [Golf Cart Wrap Designer].
//
// [Golf Cart Wraps]: https://golfcart-wraps.com
// [Golf Cart Wrap Designer]: https://golfcart-wraps.com/design?utm_source=pkg_go_dev&utm_medium=referral&utm_campaign=go_module
package cartmodels

import (
    _ "embed"
    "encoding/json"
    "strings"
)

//go:embed models.json
var rawJSON []byte

// Model describes a catalog body configuration, without dimensional fitment claims.
type Model struct {
    ID string `json:"id"`
    Brand string `json:"brand"`
    Model string `json:"model"`
    ReferenceYear string `json:"reference_year"`
    BodyConfiguration string `json:"body_configuration"`
}

// All returns an independent copy of all 34 reference configurations.
func All() []Model {
    var c struct { Models []Model `json:"models"` }
    if err := json.Unmarshal(rawJSON, &c); err != nil { panic(err) }
    return c.Models
}

// ByBrand returns models for a brand using case insensitive comparison.
func ByBrand(brand string) []Model {
    out := []Model{}
    for _, m := range All() {
        if strings.EqualFold(m.Brand, strings.TrimSpace(brand)) { out = append(out, m) }
    }
    return out
}

// ByID returns the matching model and whether it exists.
func ByID(id string) (Model, bool) {
    for _, m := range All() {
        if m.ID == id { return m, true }
    }
    return Model{}, false
}

// Brands returns the seven brands in catalog order.
func Brands() []string {
    out := []string{}
    seen := map[string]bool{}
    for _, m := range All() {
        if !seen[m.Brand] { out = append(out, m.Brand); seen[m.Brand] = true }
    }
    return out
}

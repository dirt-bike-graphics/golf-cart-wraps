package cartmodels

import "testing"

func TestCatalog(t *testing.T) {
    if len(All()) != 34 { t.Fatal("expected 34 models") }
    if len(Brands()) != 7 { t.Fatal("expected 7 brands") }
    seen := map[string]bool{}
    for _, m := range All() {
        if seen[m.ID] || m.ID == "" { t.Fatal("invalid or duplicate ID") }
        seen[m.ID] = true
        got, ok := ByID(m.ID)
        if !ok || got != m { t.Fatal("ID lookup differs from catalog") }
    }
}

func TestQueries(t *testing.T) {
    if len(ByBrand(" club car ")) != 7 { t.Fatal("expected 7 Club Car configurations") }
    for _, m := range ByBrand("e-z-go") { if m.Brand != "E-Z-GO" { t.Fatal("incorrect brand") } }
    if len(ByBrand("unknown")) != 0 { t.Fatal("unknown brand matched") }
    if _, ok := ByID("unknown"); ok { t.Fatal("unknown ID matched") }
    m, ok := ByID("denago-rover-xl")
    if !ok || m.Brand != "Denago EV" || m.Model != "Rover XL" { t.Fatal("incorrect Rover XL") }
    copy := All(); copy[0].Brand = "changed"
    if All()[0].Brand == "changed" { t.Fatal("catalog was mutated") }
}

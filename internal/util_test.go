package internal

import "testing"

func TestClean(t *testing.T) {
	input := `---
aliases:
  - Fremdinformation
---
Die [[Externes Rechnungswesen|Finanzbuchhaltung]] hat zur Fremdinformation die [[Finanzbuchhaltung Informationsfunktion|Informationsfunktion]], außenstehende Personen, z.B. Kredit gewährende Banken und Lieferanten Informationen über die wirtschaftliche Lage des Unternehmens zu unterrichten und dient somit insbesondere dem [[Gläubigerschutz]].`
	expect := `Die Finanzbuchhaltung hat zur Fremdinformation die Informationsfunktion, außenstehende Personen, z.B. Kredit gewährende Banken und Lieferanten Informationen über die wirtschaftliche Lage des Unternehmens zu unterrichten und dient somit insbesondere dem Gläubigerschutz.`
	got := Clean(input)
	if expect != got {
		t.Errorf("expected:\n===\n%v\n===\n But got\n===\n%v", expect, got)
	}
}

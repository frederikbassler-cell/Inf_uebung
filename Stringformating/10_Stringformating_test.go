package stringformating

import (
	"testing"
)

// Hinweis: Diese Tests prüfen exakte Strings (inkl. Leerzeichen).
// Stelle sicher, dass deine Funktionen genau das geforderte Format zurückgeben.

func TestFormatPerson(t *testing.T) {
	got := FormatPerson("Max", 21)
	want := "Name: Max, Alter: 21"
	if got != want {
		t.Fatalf("FormatPerson: expected %q, got %q", want, got)
	}
}

func TestFormatPrice(t *testing.T) {
	tests := []struct {
		price float64
		want  string
	}{
		{12.5, "Preis: 12.50 €"},
		{0, "Preis: 0.00 €"},
		{3.14159, "Preis: 3.14 €"},
		{99.999, "Preis: 100.00 €"}, // Rundung prüfen
	}

	for _, tc := range tests {
		got := FormatPrice(tc.price)
		if got != tc.want {
			t.Fatalf("FormatPrice(%v): expected %q, got %q", tc.price, tc.want, got)
		}
	}
}

func TestFormatOrder(t *testing.T) {
	got := FormatOrder("Apfel", 3, 0.5)
	want := "Produkt: Apfel | Menge: 3 | Einzelpreis: 0.50 €"
	if got != want {
		t.Fatalf("FormatOrder: expected %q, got %q", want, got)
	}

	// Weiterer Testfall: andere Werte
	got = FormatOrder("Banane", 10, 1.2)
	want = "Produkt: Banane | Menge: 10 | Einzelpreis: 1.20 €"
	if got != want {
		t.Fatalf("FormatOrder: expected %q, got %q", want, got)
	}
}

func TestDescribeValue(t *testing.T) {
	got := DescribeValue(42)
	want := "Wert: 42 | Typ: int"
	if got != want {
		t.Fatalf("DescribeValue(int): expected %q, got %q", want, got)
	}

	got = DescribeValue("Hi")
	want = "Wert: Hi | Typ: string"
	if got != want {
		t.Fatalf("DescribeValue(string): expected %q, got %q", want, got)
	}

	// bool test
	got = DescribeValue(true)
	want = "Wert: true | Typ: bool"
	if got != want {
		t.Fatalf("DescribeValue(bool): expected %q, got %q", want, got)
	}
}

func TestUserText(t *testing.T) {
	got := UserText("Max", 21)
	want := "Der Benutzer Max ist 21 Jahre alt."
	if got != want {
		t.Fatalf("UserText: expected %q, got %q", want, got)
	}

	got = UserText("Lisa", 0)
	want = "Der Benutzer Lisa ist 0 Jahre alt."
	if got != want {
		t.Fatalf("UserText: expected %q, got %q", want, got)
	}
}

func TestFormatScore(t *testing.T) {
	// Name linksbündig in Breite 10 + Leerzeichen + Score
	got := FormatScore("Max", 120)
	want := "Max        120" // "Max" + 7 Spaces = 10 Zeichen, dann Space, dann 120
	if got != want {
		t.Fatalf("FormatScore: expected %q, got %q", want, got)
	}

	got = FormatScore("Alexander", 5)
	want = "Alexander  5" // "Alexander" (9 Zeichen) + 1 Space, dann Space, dann 5
	if got != want {
		t.Fatalf("FormatScore: expected %q, got %q", want, got)
	}
}

package entities

import "testing"

func TestMarshalTppMessage(t *testing.T) {
	text := "test"
	m := NewTppMessage(text, CategoryERROR, MessageCodeCERTIFICATEBLOCKED, "")
	b, err := m.Marshal()
	if err != nil {
		t.Errorf("failed to marshal object - %s", err)
	}
	m2 := TppMessage{}
	err = m2.Unmarshal(b)
	if err != nil {
		t.Errorf("failed to unmarshal object -  %s", err)
	}
	if *m2.Text != text {
		t.Errorf("error unmarshaling object - want %s got %s", text, *m2.Text)
	}
}

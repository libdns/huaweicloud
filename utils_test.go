package huaweicloud

import (
	"testing"
)

func TestPrepareRecordValueTXTWithQuotes(t *testing.T) {
	result := prepareRecordValue("TXT", `"example"`)
	expected := []string{`"example"`}
	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPrepareRecordValueTXTWithoutStartingQuote(t *testing.T) {
	result := prepareRecordValue("TXT", `example"`)
	expected := []string{`"example"`}
	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPrepareRecordValueTXTWithoutEndingQuote(t *testing.T) {
	result := prepareRecordValue("TXT", `"example`)
	expected := []string{`"example"`}
	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPrepareRecordValueTXTWithoutQuotes(t *testing.T) {
	result := prepareRecordValue("TXT", `example`)
	expected := []string{`"example"`}
	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPrepareRecordValueNonTXT(t *testing.T) {
	result := prepareRecordValue("A", `example`)
	expected := []string{`example`}
	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

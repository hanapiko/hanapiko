package main

import (
	"go-reloaded/myfunctions"
	"reflect"
	"testing"
)

func TestCapitalize(t *testing.T) {
	sample := []string{"it", "(cap)", "was", "the", "best", "(cap,", "2)", "times"}
	expected := []string{"It", "was", "The", "Best", "times"}
	result := myfunctions.Capitalize(sample)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected: %v, Got: %v", expected, result)
	}

}

func TestHex(t *testing.T) {
	sample := []string{"1E", "(hex)", "files", "were", "added"}
	expected := []string{"30", "files", "were", "added"}
	result := myfunctions.Hex(sample)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected: %v, Got: %v", expected, result)
	}

}

func TestBin(t *testing.T) {
	sample := []string{"It", "has", "been", "10", "(bin)", "years"}
	expected := []string{"It", "has", "been", "2", "years"}
	result := myfunctions.Bin(sample)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed. Expected: %v, got: %v", expected, result)
	}
}

func TestVowel(t *testing.T) {
	sample := []string{"it", "was", "a", "untold", "story"}
	expected := []string{"it", "was", "an", "untold", "story"}
	result := myfunctions.Vowel(sample)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("failed. expected: %v, got: %v", expected, result)
	}
}

func TestPunctuation(t *testing.T) {
	sample := []string{"Punctuation", "tests", "are", "...", "kinda", "boring", ",don't", "you", "think", "!?"}
	expected := []string{"Punctuation", "tests", "are...", "kinda", "boring,", "don't", "you", "think!?"}
	result := myfunctions.Punctuations(sample)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("failed. expected %v, found %v", expected, result)
	}
}

func TestApostrophe(t *testing.T) {
	sample := []string{"Punctuation", "is", "'", "awesome", "'"}
	expected := []string{"Punctuation", "is", "'awesome'"}
	result := myfunctions.Apostrophe(sample)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("failed. expected %v, found %v", expected, result)
	}
}

package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("search with known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("search with unknown word", func(t *testing.T){
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add with success", func(t *testing.T){
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("add existing word", func(t *testing.T){
		commonKey := "test"
		value := "this is just a test"
		dictionary := Dictionary{commonKey: value}
		err := dictionary.Add(commonKey, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, commonKey, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update dictionary with existing word", func(t *testing.T){
		key := "test"
		value := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{key: value}

		err := dictionary.Update(key, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newDefinition)
	})

	t.Run("new word", func(t *testing.T){
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete key", func(t *testing.T){
		key := "test"
		dictionary := Dictionary{key: "test definition"}
		dictionary.Delete(key)

		_, err := dictionary.Search(key)
		assertError(t, err, ErrorNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
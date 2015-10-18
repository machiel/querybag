package querybag

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestNewMapper(t *testing.T) {
	m, err := New("queries")

	if err != nil {
		t.Errorf("Failed TestNewMapper because of error %s", err)
	}

	expected := &Bag{
		"retrieve_comments": `SELECT *
FROM comments
WHERE post_id = ?
`,
		"retrieve_deleted_posts": `SELECT *
FROM posts
WHERE deleted_at IS NOT NULL
`,
		"retrieve_users": `SELECT *
FROM users
WHERE active = 1
`,
	}

	if !reflect.DeepEqual(m, expected) {
		t.Error("The generated map didn't match the expected result.")
	}

	m, err = New("bogus-dir")
	if err == nil {
		t.Error("Expected unexistent directory to have failed")
	}

	dir, _ := ioutil.TempDir("", "querybag")
	defer os.RemoveAll(dir)
	ioutil.WriteFile(dir+"/test.sql", nil, 0)

	m, err = New(dir)
	if err == nil {
		t.Error("Expected unreadable file to have failed")
	}
}

func TestBag_Get(t *testing.T) {
	m, _ := New("queries")
	expected := `SELECT *
FROM comments
WHERE post_id = ?
`
	result := m.Get("retrieve_comments")

	if expected != result {
		t.Errorf("Expected query to be equal to: %q\ngot: %q", expected, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected getting missing key to panic, it did not")
		}
	}()

	m.Get("bogus_query")
}

func TestIsSQL(t *testing.T) {

	results := map[string]bool{
		"hi":        false,
		"hello.txt": false,
		"test.sql":  true,
	}

	for fileName, expected := range results {
		if isSQL(fileName) != expected {
			t.Errorf("Expected '%s' to be considered SQL (isSQL should've returned %t)", fileName, expected)
		}
	}

}

func TestSanitizeName(t *testing.T) {

	results := map[string]string{
		"hello.sql": "hello",
		"hello.txt": "hello.txt",
	}

	for fileName, expected := range results {
		if sanitizeName(fileName) != expected {
			t.Errorf("Expected '%s' to be rewritten to '%s'", fileName, expected)
		}
	}

}

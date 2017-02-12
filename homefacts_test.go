package homefacts

import (
	"encoding/json"
	Ω "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)

const (
	testAddress = "22280 S 209th Way, Queen Creek, AZ 85142"
	filePath    = "test/test.json"
)

func TestAPIWithProxy(t *testing.T) {
	Ω.RegisterTestingT(t)

	expectedJSON, err := ioutil.ReadFile(filePath)

	if err != nil {
		t.Fatal(err.Error())
	}

	c := New()

	resp, err := c.GetFacts(&Request{
		Address: testAddress,
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	data, err := json.Marshal(resp)

	if err != nil {
		t.Fatal(err.Error())
	}

	Ω.Ω(data).Should(Ω.MatchJSON(expectedJSON), "JSON Mismatch")
}

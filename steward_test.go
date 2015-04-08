package steward

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gorilla/mux"
	. "github.com/onsi/gomega"
)

func TestAuthserver(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Status Endpoint", func() {
		g.It("should return a 200", func() {
			r := mux.NewRouter()
			r.Path(endpoint).HandlerFunc(Status)
			server := httptest.NewServer(r)
			defer server.Close()
			req, err := http.NewRequest("GET", server.URL+endpoint, nil)
			Expect(err).NotTo(HaveOccurred())

			client := http.Client{}
			res, err := client.Do(req)
			Expect(err).NotTo(HaveOccurred())
			defer res.Body.Close()

			Expect(res.StatusCode).To(Equal(200))
			bodyBytes, err := ioutil.ReadAll(res.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(bodyBytes).To(ContainSubstring("null"))
		})
	})
}

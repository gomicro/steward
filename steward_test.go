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
		var r *mux.Router
		var server *httptest.Server
		var wasHit bool
		var client http.Client

		g.BeforeEach(func() {
			r = mux.NewRouter()
			wasHit = false
			r.Path(endpoint).HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				wasHit = true
				Status(w, req)
			})
			server = httptest.NewServer(r)
			client = http.Client{}
		})

		g.AfterEach(func() {
			server.Close()
		})

		g.It("should return a StatusOK for a healthy state", func() {
			req, err := http.NewRequest("GET", server.URL+endpoint, nil)
			Expect(err).NotTo(HaveOccurred())

			res, err := client.Do(req)
			Expect(err).NotTo(HaveOccurred())
			defer res.Body.Close()

			Expect(res.StatusCode).To(Equal(http.StatusOK))
			Expect(wasHit).To(BeTrue())
			bodyBytes, err := ioutil.ReadAll(res.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(bodyBytes).To(ContainSubstring("null"))
		})

		g.It("should return a custom status response payload", func() {
			StatusResponse(&Response{Application: "Foo", Version: "1.0.0", BuildTime: "now"})
			req, err := http.NewRequest("GET", server.URL+endpoint, nil)
			Expect(err).NotTo(HaveOccurred())

			res, err := client.Do(req)
			Expect(err).NotTo(HaveOccurred())
			defer res.Body.Close()

			Expect(res.StatusCode).To(Equal(http.StatusOK))
			Expect(wasHit).To(BeTrue())
			bodyBytes, err := ioutil.ReadAll(res.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(bodyBytes).To(ContainSubstring(`"app": "Foo",`))
			Expect(bodyBytes).To(ContainSubstring(`"version": "1.0.0",`))
			Expect(bodyBytes).To(ContainSubstring(`"buildTime": "now"`))
		})
	})
}

type Response struct {
	Application string `json:"app"`
	Version     string `json:"version"`
	BuildTime   string `json:"buildTime"`
}

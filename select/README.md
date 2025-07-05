# select

當遇到多個 goroutine 需要溝通時，可以透過 channel 來傳遞資料
而當遇到多個資料需要傳遞給一個 goroutine 時

這時可以透過 select 語法來最有效的等候多個 channel 的傳遞資訊

## code

```golang
func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayServer(5 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Errorf("error should be nit")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := makeDelayServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayServer(timeout time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(timeout)
			w.WriteHeader(http.StatusOK)
		},
	))
}
```
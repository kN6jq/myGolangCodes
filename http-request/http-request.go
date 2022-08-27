package http_request

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:9090", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	_, err = fmt.Fprintf(os.Stdout, "%s\n", dump)
	if err != nil {
		return
	}
}

package platforms

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	/*"strings"
	"mime"
	"mime/multipart"
	"io"
	"io/ioutil"
	"fmt"*/
)

func EmailBot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//json.NewDecoder(r.Body).Decode(update)
	//r.MultipartForm.Value
	/*mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(msg.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			slurp, err := ioutil.ReadAll(p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Part %q: %q\n", p.Header.Get("Foo"), slurp)
		}
	}
	*/
}

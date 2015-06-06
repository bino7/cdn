package cdn

import (
	//"io"
	"net/http"
	//"time"

	"github.com/go-martini/martini"
	"os"
	"io"
)

const FORMAT = "Mon, 2 Jan 2006 15:04:05 GMT"

func get(w http.ResponseWriter, req *http.Request, vars martini.Params) {
	name, fileName := vars["name"], vars["file"]
	filePath := conf.Root+name+"/"+fileName
	if fileExist(filePath) {
		file, err := os.Open(filePath)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		/*// check to crop/resize
		cr, isCrop := req.Form["crop"]
		scr, isSCrop := req.Form["scrop"]
		rsz, isResize := req.Form["resize"]

		isIn := isImage(file)

		if isCrop && isIn && cr != nil {
			parsed, _ := parseParams(cr[0])
			if parsed != nil {
				crop(w, file, parsed)
				return
			}
		} else if isSCrop && isIn && scr != nil {
			parsed, _ := parseParams(scr[0])
			if parsed != nil {
				smartCrop(w, file, parsed)
				return
			}

		} else if isResize && isIn && rsz != nil {
			parsed, _ := parseParams(rsz[0])
			if parsed != nil {
				resize(w, file, parsed)
				return
			}
		} else {*/
			io.Copy(w, file)
			w.WriteHeader(http.StatusOK)
		//}
		return
	}
		w.WriteHeader(http.StatusNotFound)

}

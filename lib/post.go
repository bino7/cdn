package cdn

import (
	"io"
	"net/http"
	"fmt"
	"os"
	"github.com/go-martini/martini"
	"strings"
)

func post(w http.ResponseWriter, r *http.Request,vars martini.Params) {

	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}


	defer file.Close()

	root:=conf.Root
	//remove any directory names in the filename
	//START: work around IE sending full filepath and manually get filename
	itemHead := header.Header["Content-Disposition"][0]
	lookfor := "filename=\""
	fileIndex := strings.Index(itemHead, lookfor)

	if fileIndex < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}

	filename := itemHead[fileIndex+len(lookfor):]
	filename = filename[:strings.Index(filename, "\"")]

	slashIndex := strings.LastIndex(filename, "\\")
	if slashIndex > 0 {
		filename = filename[slashIndex+1:]
	}

	slashIndex = strings.LastIndex(filename, "/")
	if slashIndex > 0 {
		filename = filename[slashIndex+1:]
	}
	//END: work around IE sending full filepath
	name:=vars["name"]
	dir:=root+name+"/"
	_, err = os.Stat(dir)
	if fileExist(dir)==false{
		os.MkdirAll(dir,0777)
	}
	out, err := os.Create(dir+filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)
}

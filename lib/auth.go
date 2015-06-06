package cdn
import (
    "net/http"
    "github.com/go-martini/martini"
    "fmt"
)

func auth(w http.ResponseWriter, req *http.Request, vars martini.Params){
    fmt.Println("authorize")
    ///w.WriteHeader(http.StatusUnauthorized)
}
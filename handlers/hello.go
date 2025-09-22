package handlers
import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct{
	l *log.Logger
}
func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter,r *http.Request){
	h.l.Println("Hello microservice")//log output on server
	dt,err:=io.ReadAll(r.Body)
	if err!=nil{
		http.Error(rw,"Oops",http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oops"))
		return
	}
	// h.l.Printf("Data %s\n",dt)//logging data to server
	fmt.Fprintf(rw,"Hello genius boy %s",dt)
}
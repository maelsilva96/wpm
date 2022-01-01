package basic

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/maelsilva96/wpm"
)

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	wpm.LoadModelBind()
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		wpm.SendError(err)
	}
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			wpm.SendError(err)
		}
		wpm.SendData(200, map[string]string{}, body)
	}
	str := fmt.Sprintf("Falha ao tentar recuperar a informação da API (%s)", resp.Status)
	wpm.SendError(fmt.Errorf(str))
}
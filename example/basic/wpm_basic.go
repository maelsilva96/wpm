package basic

import (
	"fmt"
	"github.com/maelsilva96/wpm"
	"io/ioutil"
	"net/http"
)

func main() {
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
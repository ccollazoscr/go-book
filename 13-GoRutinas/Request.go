package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"
)

func main() {
	SincornizacionMedianteMutex()
}

const (
	token1      = ""
	token2      = ""
	companykey1 = ""
	companykey2 = ""
)

func SincornizacionMedianteMutex() {
	fmt.Printf("\n***** SincornizacionMedianteMutex *****\n")
	//Obtiene el número de procesadores del equipo
	//Recover se utiliza para recuperarse de un panic que se haya lanzado
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recuperandose del pánico: ", p)
		}
	}()
	tareasParalelas := runtime.GOMAXPROCS(0)
	fmt.Println("tareasParalelas:", tareasParalelas)

	cont := 0
	for i := 0; i < 20; i++ {
		wg := sync.WaitGroup{}
		wg.Add(tareasParalelas)
		for t := 0; t < tareasParalelas; t++ {
			cont++
			if t%2 == 0 {
				go Request(&wg, cont, companykey1, token1)
			} else {
				go Request(&wg, cont, companykey2, token2)
			}
		}
		wg.Wait()
	}

}

func Request(wg *sync.WaitGroup, i int, companykey string, token string) {
	defer wg.Done()
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://example.com", nil)

	if err != nil {
		fmt.Println("Error creando Request", i)
		return
	}

	req.Header.Add("Autorization", token)
	resp, errRequest := client.Do(req)
	defer resp.Body.Close()

	if errRequest != nil {
		fmt.Println("Error realizando Request", i)
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err == nil {
		fmt.Println("Resutlado: ", string(body) == companykey)
	}
}

// func Request2(wg *sync.WaitGroup, i int) {
// 	defer wg.Done()
// 	fmt.Println("Request2-", i)
// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var seed int64

var crypt []string

var encrypt = make(map[string]string)
var decrypt = make(map[string]string)

// ServeRefresh .
func ServeRefresh(w http.ResponseWriter, r *http.Request) {
	seed = time.Now().UnixNano()
	rand.Seed(seed)

	for k, v := range rand.Perm(len(crypt)) {
		encrypt[crypt[k]] = crypt[v]
		decrypt[crypt[v]] = crypt[k]
	}
}

func doCrypt(s string, m map[string]string) (r string) {
	for _, v := range s {
		if vv, ok := m[string(v)]; ok {
			r += vv
		} else {
			r += string(v)
		}
	}
	return
}

// ServeTranslate .
func ServeTranslate(w http.ResponseWriter, r *http.Request) {
	in := r.FormValue("in")
	out := r.FormValue("out")

	if out != "" {
		in = doCrypt(out, decrypt)
	} else if in != "" {
		out = doCrypt(in, encrypt)
	}

	fmt.Fprintf(w, `<html>
    <head>
        <title>Just a joke</title>
    </head>
    <body>
        <p>Time: %s</p>
        <form method="POST">
            <textarea name="in" style="width: 80%%; height: 40%%;">%s</textarea>
            <br /><br />
            <textarea name="out" style="width: 80%%; height: 40%%;">%s</textarea>
            <br /><br />
            <input type="submit" value="" style="width: 40%%; height: 10%%; background-color: gray;">
        </form>
    </body>
</html>`, time.Unix(seed/1000000000, 0).Format("2006-01-02 15:04:05"), in, out)
}

func main() {
	f, err := os.Open("3500.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		crypt = append(crypt, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	ServeRefresh(nil, nil)

	http.HandleFunc("/refresh", ServeRefresh)
	http.HandleFunc("/translate", ServeTranslate)

	log.Fatal(http.ListenAndServe(":4444", nil))
}

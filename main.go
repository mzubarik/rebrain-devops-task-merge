package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

<<<<<<< HEAD
func sendResponse(w http.ResponseWriter, r *http.Request) {
  keys, ok := r.URL.Query()["a"]
  if !ok || len(keys[0]) < 1 {
      log.Println("Url Param 'a' is missing")
      return
  }
=======
func handler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["a"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'a' is missing")
		return
	}
>>>>>>> develop

	a, err := strconv.Atoi(keys[0])

	if err != nil {
		return
	}

	keys, ok = r.URL.Query()["b"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'b' is missing")
		return
	}
	b, err := strconv.Atoi(keys[0])
	if err != nil {
		log.Println("Bad number!")
		return
	}

	fmt.Fprintf(w, "%d + %d = %d", a, b, sum(a, b))
}

func main() {
<<<<<<< HEAD
    http.HandleFunc("/", sendResponse)
    log.Fatal(http.ListenAndServe(":7000", nil))
=======
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":7000", nil))
>>>>>>> develop
}

<<<<<<< HEAD
func sum (a, b int) int {
<<<<<<< HEAD
  t := b + a
  return t
=======
  return a + b
=======
func sum(a, b int) int {
	t := a + b
	return t
>>>>>>> origin/feature
>>>>>>> develop
}

package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strings"

	"github.com/Paypi/paypi-go"
)

// ToYodaSpeak cuts a string into two peices and flips them round.
// Yoda actually uses object-subject-verb word ordering insead of
// English's subject-verb-object, however instead of whipping out
// NLP or training a TF model, just flipping the text works quite well.
func ToYodaSpeak(s string) string {
	split := strings.Split(strings.ToLower(s), " ")

	// Find the center point of the text
	splitPoint := int(math.Ceil(float64(len(split)) / float64(2)))

	start := split[0:splitPoint]
	end := split[splitPoint:]

	// Put the end of the sentence at the start of the string.
	flipped := strings.Join(end, " ") + " " + strings.Join(start, " ")

	// Capitalise first character
	return strings.ToUpper(string([]rune(end[0])[0])) + flipped[1:]
}

func handleYodaRequest(w http.ResponseWriter, r *http.Request) {
	// fetch auth Header and check subscriberToken
	auth := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(auth) != 2 {
		http.Error(w, "User token not given", http.StatusUnauthorized)
		return
	}
	user, err := paypi.Authenticate(auth[1])

	if err != nil {
		http.Error(w, "User token is unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	text := r.URL.Query().Get("text")

	// Charge the user once processing is complete
	_, err = user.MakeCharge(paypi.MakeChargeInput{
		ChargeIdentifier: "<STATIC_CHARGE_IDENTIFIER>",
	})
	if err != nil {
		http.Error(w, "Unable to make charge", http.StatusServiceUnavailable)
		return
	}

	info := struct {
		Text string `json:"text"`
	}{
		Text: ToYodaSpeak(text),
	}
	json.NewEncoder(w).Encode(info)
}

func main() {
	// Setup PayPI
	paypi.Key = "<API_SECRET_KEY>"

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleYodaRequest)

	s := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}

	// Begin listening for requests.
	log.Printf("Yodafying text on %s", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		log.Println("Server an error has", err)
	}
}

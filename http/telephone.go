package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// getTelephone
// @summary get a telephone information
// @description returns one telephone information by telephone number
// @router /telephones/{number} [get]
// @tags telephone
// @produce json
// @param number path string true "telephone number" minlength(11)
// @success 200 {string} telephone-json
func getTelephone(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")

	telephone, err := telephoneRepository.GetTelephoneByNumber(number)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get a record: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}

	tj, err := json.Marshal(telephone)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute json.Marshal() on getTelephone(): %s", err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

	w.Write([]byte(tj))
}

// listTelephone
// @summary list telephone informations
// @description returns all telephone informations
// @router /telephones [get]
// @tags telephone
// @produce json
// @success 200 {string} telephone-json
func listTelephone(w http.ResponseWriter, r *http.Request) {
	telephone, err := telephoneRepository.ListTelephone()

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to select table: %s\n", err.Error())
	}

	tj, err := json.Marshal(telephone)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute json.Marshal() on listTelephone(): %s", err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

	w.Write([]byte(tj))
}

// postTelephone
// @summary post a telephone information
// @description register one telephone information by number
// @router /telephones [post]
// @tags telephone
// @produce json
// @success 200 {string} telephone-json
func postTelephone(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")

	reqBody := map[string]string{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute json decode on postTelephone: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}

	defer r.Body.Close()

	ownerId, err := strconv.Atoi(reqBody["owner_id"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to covert owner_id in request body to integer on postTelephone: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}
	iccId, err := strconv.Atoi(reqBody["icc_id"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to covert icc_id in request body to integer on postTelephone: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}

	if err = telephoneRepository.PostTelephone(ownerId, iccId, number); err != nil {
		fmt.Fprintf(os.Stderr, "failed to insert a record: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}
}

func putTelephone(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")

	reqBody := map[string]string{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute json decode on postTelephone: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}

	ownerId, err := strconv.Atoi(reqBody["owner_id"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to covert owner_id in request body to integer on putTelephone: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}
	iccId, err := strconv.Atoi(reqBody["icc_id"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to covert icc_id in request body to integer on putTelephone: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}

	defer r.Body.Close()

	_, err = telephoneRepository.GetTelephoneByNumber(number)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get a record: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}

	if err := telephoneRepository.PutTelephoneByNumber(number, ownerId, iccId); err != nil {
		fmt.Fprintf(os.Stderr, "failed to update a record: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}
}

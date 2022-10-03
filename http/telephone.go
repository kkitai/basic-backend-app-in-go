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

	telephone, err := conn.GetTelephoneByNumber(number)

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
	telephone, err := conn.ListTelephone()

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

	err = conn.PostTelephone(ownerId, iccId, number)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to insert a record: %s\n", err)
		http.Error(w, http.StatusText(400), 400)
	}
}

//func putTelephone(w http.ResponseWriter, r *http.Request) {
//	number := chi.URLParam(r, "number")
//
//	reqBody := map[string]string{}
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&reqBody); err != nil {
//		fmt.Fprintf(os.Stderr, "failed to execute json decode on postTelephone: %s\n", err)
//		http.Error(w, http.StatusText(400), 400)
//	}
//
//	if _, err := isTelephoneExists(number); err != nil {
//		fmt.Fprintf(os.Stderr, "failed to get a telephone by given number_code(%s): %s\n", number, err)
//		http.Error(w, http.StatusText(400), 400)
//	}
//
//	defer r.Body.Close()
//
//	if _, err := conn.PutTelephone(reqBody["name"], reqBody["owner_id"], number); err != nil {
//		fmt.Fprintf(os.Stderr, "failed to update the record: %s\n", err)
//	}
//}
//
//func isTelephoneExists(number string) (bool, error) {
//	rows, err := conn.GetTelephone(number)
//	if err != nil {
//		return false, fmt.Errorf("failed to get a record: %w", err)
//	}
//
//	var vs []Telephone
//	for rows.Next() {
//		var v Telephone
//		rows.Scan(&v.ID, &v.Name, &v.NACCSCode, &v.OwnerID)
//		vs = append(vs, v)
//	}
//
//	if len(vs) != 1 {
//		return false, fmt.Errorf("telephone length was not 1 (actual: %d)", len(vs))
//	}
//
//	return true, nil
//}

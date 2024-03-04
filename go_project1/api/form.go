package api

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/schema"
	"github.com/rs/zerolog/log"

	"net/http"
	db "sqlc"
	util "util"
	web "web"
)

// "github.com/go-chi/render"

// "bytes"
// "time"

// "github.com/gorilla/schema"

// func (server *Server) formHandler(w http.ResponseWriter, r *http.Request) error {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
// 		return err
// 	}

// 	msg := &db.SubmitFormTxParams{}

// 	decoder := schema.NewDecoder()
// 	decodeErr := decoder.Decode(msg, r.PostForm)
// 	if decodeErr != nil {
// 		log.Error().Err(decodeErr).Msg("Error mapping parsed form data to struct")
// 		http.Error(w, "Error processing form data", http.StatusInternalServerError)
// 		return err
// 	}

// 	if !msg.ValidateForm() {
// 		// log.Info().Msg(msg)
// 		web.Render(w, "lienhe", msg)
// 		return err
// 	}

// 	fmt.Printf("\n formHandler dbHandler %v\n", server.DbHandler)

// 	if _, err := server.DbHandler.SubmitFormTx(r.Context(), *msg); err != nil {
// 		log.Error().Err(err)
// 		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
// 		return err
// 	}

// 	http.Redirect(w, r, "/thankyou", http.StatusSeeOther)
// 	return util.WriteJSON(w, http.StatusOK, msg)

// }

func (server *Server) formHandler(w http.ResponseWriter, r *http.Request) {
	// apiURL := getBaseURL(r)
	// apiURL := "/api/v1/form"
	msg, err := parseForm(r)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	fmt.Printf("\n formHandler msg %v\n", msg)
	if !msg.ValidateForm() {
		// log.Info().Msg(msg)
		web.Render(w, "lienhe", msg)
		return
	}
	fmt.Printf("\n formHandler msg àtẻ validate %v\n", msg)

	if _, err := server.DbHandler.SubmitFormTx(r.Context(), *msg); err != nil {
		log.Error().Err(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}
	// http.Redirect(w, r, "/thankyou", http.StatusSeeOther)

	util.WriteJSON(w, http.StatusOK, msg)

	// payload, err := json.MarshalIndent(msg, "", "  ")
	// fmt.Printf("\n formHandler payload %v\n", string(payload))
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error encoding JSON payload")
	// 	http.Error(w, "Error processing form data", http.StatusInternalServerError)
	// 	return
	// }

	// // Make a POST request to the specified URL for each test case
	// post, err := APIRequest(apiURL, payload)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error making API request")
	// 	http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Printf("\n formHandler msg post %v\n", post)
	// fmt.Printf("\n formHandler msg post.StatusCode %v\n", post.StatusCode)
	// fmt.Printf("\n formHandler msg post.Body %v\n", post.Body)

	// resp, err := http.Get(apiURL)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error making GET request")
	// }
	// defer resp.Body.Close()

	// // Read the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error reading response body")
	// }

	// // Parse the JSON data
	// var data map[string]interface{}
	// if err := json.Unmarshal(body, &data); err != nil {
	// 	log.Error().Err(err).Msg("Error parsing JSON data")
	// }

	// // Access the JSON data
	// fmt.Printf("Received JSON data: %+v\n", data)

	// http.Redirect(w, r, "/thankyou", http.StatusSeeOther)

	// render.JSON(w, r, payload)
	// util.WriteJSON(w, post.StatusCode, msg)
}

// func getBaseURL(r *http.Request) string {
// 	base := r.Header.Get("X-Forwarded-Prefix")
// 	if base == "" {
// 		base = "http://localhost:1414"
// 	}
// 	apiURL := fmt.Sprintf("%s/api/v1/form", base)
// 	fmt.Printf("\n formHandler apiURL %v\n", apiURL)

// 	return apiURL
// }

func parseForm(r *http.Request) (*db.SubmitFormTxParams, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	msg := &db.SubmitFormTxParams{}
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(msg, r.PostForm)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return msg, nil
}

// func sendAPIRequests(w http.ResponseWriter, r *http.Request, apiURL string, msg db.SubmitFormTxParams) {
// 	var resp *http.Response
// 	fmt.Printf("\n sendAPIRequests msg %v\n", msg)
// 	payload, err := json.Marshal(msg)
// 	fmt.Printf("\n sendAPIRequests payload %v\n", payload)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error encoding JSON payload")
// 		http.Error(w, "Error processing form data", http.StatusInternalServerError)
// 		return
// 	}

// 	_, err = APIRequest(apiURL, payload)
// 	fmt.Printf("\n sendAPIRequests resp %v\n", resp)
// 	fmt.Printf("\n sendAPIRequests resp.StatusCode %v\n", resp.StatusCode)

// 	if err != nil {
// 		log.Error().Err(err).Msg("Error making API request")
// 		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
// 		return
// 	}
// }

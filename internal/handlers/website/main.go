package website

import (
	
)
func (conn *Handlers) Homepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/html")
	err := templates.Homepage.Execute(w, nil)
	if err != nil {
		go conn.Logger.Println(err)
	}
}
package rest

import "net/http"

func (a *UserApp) GetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/register", a.HandleRegister)
	mux.HandleFunc("/user/login", a.HandleLogin)
	// mux.HandleFunc("/user/logout")

	// mux.HandleFunc("/user/ping")

	return mux
}

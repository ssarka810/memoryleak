package api

import "net/http"

func (server *Server)RouteApi(){
	router := server.router
	//to use middlewire for all paths
	// router.Use()
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome home!!"))
	})

	router.HandleFunc("/v1/api/dynamic/struct/decoder",server.DynamicStructDecoderApiHandler).Methods("GET")
	router.HandleFunc("/v1/api/dynamic/input",server.DynamicInputrApiHandler)
	router.HandleFunc("/v1/api/test/memory/leakage",server.TestMemoryLeakageHandler)
}
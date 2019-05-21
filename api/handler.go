package api

//
// import (
// 	"context"
// 	"net/http"
// )
//
// func Config() http.HandlerFunc {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		resp := configer(req.Context())
// 		Success(http.StatusOK, resp, rw)
// 	})
// }
//
// func configer(ctx context.Context) (response configResponse) {
// 	var profiles []mapElement
// 	for v, n := range Profiles {
// 		profiles = append(profiles, mapElement{Name: n, Value: v})
// 	}
//
// 	response = configResponse{
// 		Profiles: profiles,
// 	}
// 	return
// }

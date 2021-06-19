package main

import (
	g "RecursosTuristicos/knn"
	m "RecursosTuristicos/reader"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var recursos []m.Recurso
// func printParam(aea string){
// 	if aea == "" {
// 		fmt.Print("aea")
// 	}else{fmt.Print(aea)}
// }
func Listar(res http.ResponseWriter, req *http.Request) {
	region := req.FormValue("region")
	provincia := req.FormValue("provincia")
	distrito := req.FormValue("distrito")
	categoria := req.FormValue("categoria")
	tipo := req.FormValue("tipo")
	subtipo := req.FormValue("subtipo")
	var aux []m.Recurso
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	for _, recurso := range recursos {
		if !strings.EqualFold(recurso.REGION, region) && region != "" {
			continue
		}
		if !strings.EqualFold(recurso.PROVINCIA, provincia) && provincia != "" {
			continue
		}
		if !strings.EqualFold(recurso.DISTRITO, distrito) && distrito != "" {
			continue
		}
		if !strings.EqualFold(recurso.CATEGORIA, categoria) && categoria != "" {
			continue
		}
		if !strings.EqualFold(recurso.TIPO, tipo) && tipo != "" {
			continue
		}
		if !strings.EqualFold(recurso.SUBTIPO, subtipo) && subtipo != "" {
			continue
		}
		aux = append(aux,recurso)
	}
	//fmt.Print(aux)
	jsonBytes, _ := json.MarshalIndent(aux, "", " ")
		io.WriteString(res, string(jsonBytes))
}
func ListarFiltrado(res http.ResponseWriter, req *http.Request) {
	readK := req.FormValue("k")
	K, _ := strconv.Atoi(readK)
	fmt.Print(K)
	res.Header().Set("Content-Type", "application/json")
	var recursoz []g.RecursoPredict

	for i :=0; i< len(recursos); i++ {
		recursoz = append(recursoz,g.Kn(recursos,K,recursos[i]))
		//fmt.Print(i,": siguiente recurso | ")
	}
	jsonBytes, _ := json.MarshalIndent(recursoz, "", " ")
    io.WriteString(res, string(jsonBytes))
}
func serveFiles(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    p := r.URL.Path
    if p == "/" {
        p = "./static/index.html"
    }
    http.ServeFile(w, r, p)
}
// func getRecursoByRegionName(res http.ResponseWriter, req *http.Request) {
//     nombre := req.FormValue("nombre")
//     //ivalDni, _ := strconv.Atoi(valDni)
//     //definir el tipo de contenido a devolver
//     res.Header().Set("Content-Type", "application/json")
//     //busqueda
//     for _, recurso := range recursos {
//         if recurso.REGION == nombre {
//             //renderizar
//             jsonBytes, _ := json.MarshalIndent(recurso, "", " ")
//             io.WriteString(res, string(jsonBytes))
//         }
//     }
// }

func handleRequest() {
	http.HandleFunc("/listar", Listar)
	http.HandleFunc("/listarF", ListarFiltrado)
	http.HandleFunc("/", serveFiles)
	//http.HandleFunc("/region",getRecursoByRegionName)
	log.Fatal(http.ListenAndServe(":9000", nil))
}


func main() {
	recursos = m.CargarCanales()
	handleRequest()
}
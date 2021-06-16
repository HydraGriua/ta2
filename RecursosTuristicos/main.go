package main

import (
	m "RecursosTuristicos/reader"
	g "RecursosTuristicos/sorter"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
)
type Ubicacion struct {
	x float64
	y float64
}
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
		jsonBytes, _ := json.MarshalIndent(recurso, "", " ")
		io.WriteString(res, string(jsonBytes))
	}
}
func ListarFiltrado(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	aea := Ubicacion{
		x: -76.17,
		y: -6.60,
	}
	jsonBytes, _ := json.MarshalIndent(kn(recursos,5,aea), "", " ")
	io.WriteString(res, string(jsonBytes))
}
func kn(ar []m.Recurso, k int, x Ubicacion) []m.RecursoD {
	var arraux, arraux2 []m.RecursoD
	for i := 0; i < len(ar); i++{
		distance := math.Sqrt(math.Pow(ar[i].LONGITUD - x.x,2) + math.Pow(ar[i].LATITUD - x.y,2));
		recurso := m.RecursoD {
			RECURSO: ar[i],
			DIST: distance,
		}
		arraux = append(arraux,recurso)
	}
	dist := func(p1, p2 *m.RecursoD) bool {
		 		return p1.DIST < p2.DIST
	}
	// decreasingDist := func(p1, p2 *m.RecursoD) bool {
	// 	return dist(p2, p1)
	// }
	g.By(dist).Sort(arraux)
	//TODO: agregar criterios de prediccion de categorias X
	for i:=0; i< k; i++ {
		arraux2 = append(arraux2,arraux[i])
	}
	return arraux2
}
func serveFiles(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    p := r.URL.Path
    if p == "/listarI" {
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
	http.HandleFunc("/listarI", serveFiles)
	//http.HandleFunc("/region",getRecursoByRegionName)
	log.Fatal(http.ListenAndServe(":9000", nil))
}


func main() {
	recursos = m.CargarCanales()
	handleRequest()
}
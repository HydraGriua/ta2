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
var recursos []m.Recurso
func printParam(aea string){
	if aea == "" {
		fmt.Print("aea")
	}else{fmt.Print(aea)}
}
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
func kn(ar []m.Recurso, k int, v m.Recurso) []m.RecursoD {
	var arraux, arraux2 []m.RecursoD
	for i := 0; i < len(ar); i++{
		distance := math.Sqrt(math.Pow(ar[i].LONGITUD - v.LONGITUD,2) + math.Pow(ar[i].LATITUD - v.LATITUD,2));
		recurso := m.RecursoD {
			RECURSO: ar[i],
			DIST: int(distance),
		}
		arraux = append(arraux,recurso)
	}
	dist := func(p1, p2 *m.RecursoD) bool {
		 		return p1.DIST < p2.DIST
	}
	decreasingDist := func(p1, p2 *m.RecursoD) bool {
		return dist(p2, p1)
	}
	g.By(decreasingDist).Sort(arraux)
	arraux2 = append(arraux2,arraux[0],arraux[1],arraux[2])
	return arraux2
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
	//http.HandleFunc("/region",getRecursoByRegionName)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	recursos = m.CargarCanales()
	handleRequest()
}
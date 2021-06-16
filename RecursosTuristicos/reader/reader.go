package reader

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Recurso struct {
	REGION string
	PROVINCIA string
	DISTRITO string
	CODIGO string
	NOMBRE string
	CATEGORIA string
	TIPO string
	SUBTIPO string
	URL string
	LATITUD float64
	LONGITUD float64
}
type RecursoD struct {
	RECURSO Recurso
	DIST float64
}
func Errcheck(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
func menor(a int,b int) int{
	if a<=b{return a}else{return b}
}

func CargarCanales() []Recurso{
	records, err := ReadCSVFromUrl("https://raw.githubusercontent.com/HydraGriua/ta2-PrograConcurrenteDistribuida/main/Inventario.csv")
	//records, err := readData("Inventario.csv")
	Errcheck(err)
	chns := make([]chan Recurso, 5)
	lim := len(records) / 4
	it := 0
	for i :=0; i<len(records); i += lim{
		cnk := records[i:menor(i+lim,len(records))]
		chns[it] = make(chan Recurso)
		go Cargar(cnk, chns[it])
		it++
	}
	var recursos []Recurso
	for _, channel := range chns {
		for recurso := range channel {
			recursos = append(recursos, recurso)
		}
	}
	return recursos
}
func Cargar(records [][]string, c chan Recurso ) {
	var latitud, longitud float64
	//var recursos []Recurso
	for _, record := range records {
		if record[9] != "" {
			aux, err := strconv.ParseFloat(strings.TrimSpace(record[9]), 64)
			Errcheck(err)
			latitud = aux
		} else {
			continue
		}
		if record[10] != "" {
			aux, err := strconv.ParseFloat(strings.TrimSpace(record[10]), 64)
			Errcheck(err)
			longitud = aux
		} else {
			continue
		}
		recurso := Recurso{
			REGION: record[0],
			PROVINCIA:record[1],
			DISTRITO: record[2],
			CODIGO: record[3],
			NOMBRE: record[4],
			CATEGORIA: record[5],
			TIPO: record[6],
			SUBTIPO: record[7],
			URL: record[8],
			LATITUD: latitud,
			LONGITUD: longitud,
		}
		//ecursos = append(recursos, recurso)
		c <- recurso
	}
	//return recursos
	close(c)

}
func ReadData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
func ReadCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return [][]string{}, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	if _, err := reader.Read(); err != nil {
		return [][]string{}, err
	}
	data, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return data, nil
}
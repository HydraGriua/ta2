package knn

import (
	m "RecursosTuristicos/reader"
	g "RecursosTuristicos/sorter"
	//"fmt"
	"math"
)

type RecursoPredict struct {
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

	RegionP string
	ProvinciaP string
	DistritoP string
	CategoriaP string
	TipoP string
	SubtipoP string
}


func Pred(ar []m.RecursoD, k int, x m.Recurso) RecursoPredict {
	regiones,provincias,distritos,categorias,tipos,subtipos := make(map[string]int),make(map[string]int),make(map[string]int),make(map[string]int),make(map[string]int),make(map[string]int)
	for i:= 0; i< k; i++{
		if _, ok := regiones[ar[i].RECURSO.REGION]; !ok{
			regiones[ar[i].RECURSO.REGION] = 1
		}else{
			regiones[ar[i].RECURSO.REGION] += 1
		}
		if _, ok := provincias[ar[i].RECURSO.PROVINCIA]; !ok{
			provincias[ar[i].RECURSO.PROVINCIA] = 1
		}else{
			provincias[ar[i].RECURSO.PROVINCIA] += 1
		}
		if _, ok := distritos[ar[i].RECURSO.DISTRITO]; !ok{
			distritos[ar[i].RECURSO.DISTRITO] = 1
		}else{
			distritos[ar[i].RECURSO.DISTRITO] += 1
		}
		if _, ok := categorias[ar[i].RECURSO.CATEGORIA]; !ok{
			categorias[ar[i].RECURSO.CATEGORIA] = 1
		}else{
			categorias[ar[i].RECURSO.CATEGORIA] += 1
		}
		if _, ok := tipos[ar[i].RECURSO.TIPO]; !ok{
			tipos[ar[i].RECURSO.TIPO] = 1
		}else{
			tipos[ar[i].RECURSO.TIPO] += 1
		}
		if _, ok := subtipos[ar[i].RECURSO.SUBTIPO]; !ok{
			subtipos[ar[i].RECURSO.SUBTIPO] = 1
		}else{
			subtipos[ar[i].RECURSO.SUBTIPO] += 1
		}
	}
	//fmt.Print(regiones,provincias,distritos,categorias,tipos,subtipos)
	regionesC,provinciasC,distritosC,categoriasC,tiposC,subtiposC := 0,0,0,0,0,0
	regionP,provinciaP,distritoP,categoriaP,tipoP,subtipoP := "","","","","",""
	for key, value := range regiones {
		if value  > regionesC{
			regionP = key
		}
	}
	for key, value := range provincias {
		if value  > provinciasC{
			provinciaP = key
		}
	}
	for key, value := range distritos {
		if value  > distritosC{
			distritoP = key
		}
	}
	for key, value := range categorias {
		if value  > categoriasC{
			categoriaP = key
		}
	}
	for key, value := range tipos {
		if value  > tiposC{
			tipoP = key
		}
	}
	for key, value := range subtipos {
		if value  > subtiposC{
			subtipoP = key
		}
	}
	//fmt.Print(x, "<-prediccion")
	rec := RecursoPredict{
		REGION: x.REGION,
		PROVINCIA: x.PROVINCIA,
		DISTRITO: x.DISTRITO,
		CODIGO: x.CODIGO,
		NOMBRE: x.NOMBRE,
		CATEGORIA: x.CATEGORIA,
		TIPO: x.TIPO,
		SUBTIPO: x.SUBTIPO,
		URL: x.URL,
		LATITUD: x.LATITUD,
		LONGITUD: x.LONGITUD,
		RegionP: regionP,
		ProvinciaP: provinciaP,
		DistritoP: distritoP,
		CategoriaP: categoriaP,
		TipoP: tipoP,
		SubtipoP: subtipoP,
	}
	//fmt.Print(rec, "<-recurso predicho")
	return rec;
}
func Dist(y m.Recurso, x m.Recurso, c chan m.RecursoD){
	distance := math.Sqrt(math.Pow(y.LONGITUD - x.LONGITUD,2) + math.Pow(y.LATITUD - x.LATITUD,2));
	recurso := m.RecursoD {
		RECURSO: y,
		DIST: distance,
	}
	c<- recurso
}

func Kn(ar []m.Recurso, k int, x m.Recurso) RecursoPredict{
	//var arraux []m.RecursoD
	chn := make(chan m.RecursoD)
	//bools:= make(chan int, 4)
	//lim := len(ar) / 4
	it := 0
	//fmt.Print(len(ar), "<-arreglo")
	for i :=0; i<len(ar); i ++{
		//cnk := ar[i:m.Menor(i+lim,len(ar))]
		//chns[it] = make(chan m.RecursoD)
		//bools[it] = make(chan int)
		if ar[i].LONGITUD != x.LONGITUD && ar[i].LATITUD != x.LATITUD{
			it++
			go Dist(ar[i],x, chn)
		}
		//fmt.Print(it, " | ")
	}
	//fmt.Print(it, "<-iterador")
	//fmt.Print("antes de canales")
	//fmt.Print(len(chn))
	var recursos []m.RecursoD
	for recurso := range chn {
		recursos = append(recursos, recurso)
		it--
		//fmt.Print(it, " | ")
		if it == 0{
			close(chn)
		}
		//fmt.Print(recurso, it)
	}
	//fmt.Print("se cerro")
	//fmt.Print("siguiente canal")
	dist := func(p1, p2 *m.RecursoD) bool {
		return p1.DIST < p2.DIST
	}
	g.By(dist).Sort(recursos)
	//fmt.Print(x, "<- recurso a predecir")
	return Pred(recursos,k, x)
}
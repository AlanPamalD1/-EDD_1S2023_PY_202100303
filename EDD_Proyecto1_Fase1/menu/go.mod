module menu

go 1.20

//require estudiante v0.0.0-00010101000000-000000000000
//require listadobleenlazada v0.0.0-00010101000000-000000000000

require (
	estudiante v0.0.0-00010101000000-000000000000
	listadobleenlazada v0.0.0-00010101000000-000000000000
)

require (
	cola v0.0.0-00010101000000-000000000000 // indirect
	pilabitacora v0.0.0-00010101000000-000000000000 // indirect
	pilaregistro v0.0.0-00010101000000-000000000000 // indirect
)

replace estudiante => ../estudiante

replace listadobleenlazada => ../listadobleenlazada

replace cola => ../cola

replace pilabitacora => ../pilabitacora

replace pilaregistro => ../pilaregistro

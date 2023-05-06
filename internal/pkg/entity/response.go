package entity

type Local struct {
	Fecha                      string `json:"fecha" xml:"fecha"`
	LocalId                    string `json:"local_id" xml:"local_id"`
	LocalNombre                string `json:"local_nombre" xml:"local_nombre"`
	ComunaNombre               string `json:"comuna_nombre" xml:"comuna_nombre"`
	LocalidadNombre            string `json:"localidad_nombre" xml:"localidad_nombre"`
	LocalDireccion             string `json:"local_direccion" xml:"local_direccion"`
	FuncionamientoHoraApertura string `json:"funcionamiento_hora_apertura" xml:"funcionamiento_hora_apertura"`
	FuncionamientoHoraCierre   string `json:"funcionamiento_hora_cierre" xml:"funcionamiento_hora_cierre"`
	LocalTelefono              string `json:"local_telefono" xml:"local_telefono"`
	LocalLat                   string `json:"local_lat" xml:"local_lat"`
	LocalLng                   string `json:"local_lng" xml:"local_lng"`
	FuncionamientoDia          string `json:"funcionamiento_dia" xml:"funcionamiento_dia"`
	FkRegion                   string `json:"fk_region" xml:"fk_region"`
	FkComuna                   string `json:"fk_comuna" xml:"fk_comuna"`
	FkLocalidad                string `json:"fk_localidad" xml:"fk_localidad"`
}

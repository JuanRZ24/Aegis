package monitor


type ResourceMonitor interface {
	Name()				string //Nombre del recurso(CPU,RAM,etc.)
	Collect()			Data   //Devuelve los datos estructurados
	Format(data Data)	string //Convierte datos en string bonito para consola
}

type Data struct {
    Metrics map[string]string
}


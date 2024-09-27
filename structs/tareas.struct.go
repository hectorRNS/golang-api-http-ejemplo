package structs

type Tareas struct {
	Id    int    `json:"id"`
	Tarea string `json:"tarea"`
}

type TareaId struct {
	TareaId int `json:"tarea-id" xml:"tarea-id" form:"tarea-id"`
}

type NuevaTarea struct {
	Tarea string `json:"tarea" xml:"tarea" form:"tarea"`
}

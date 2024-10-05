package structs

type Tareas struct {
	Id    int    `json:"id" validate:"required,min=1,max=999999999"`
	Tarea string `json:"tarea"  validate:"required,texto,min=1,max=60"`
}

type TareaId struct {
	TareaId int `json:"tarea-id" xml:"tarea-id" form:"tarea-id" validate:"required,min=1,max=999999999"`
}

type NuevaTarea struct {
	Tarea string `json:"tarea" xml:"tarea" form:"tarea" validate:"required,texto,min=1,max=60"`
}

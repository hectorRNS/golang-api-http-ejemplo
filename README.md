# API de ejemplo
Este proyecto fue creado cómo ejemplo de demostración de conocimiento de uso del lenguaje Golang.

## Desarrollo
- log.Println(variables) (para debuggear)
- log.Info(variable) (para debuggear)

### Correr codigo
- go run main.go
- air (correr live reload ubuntu)

### Probar API con curl en terminal
- Inicio / curl http://localhost:3000 && echo
- Tarea Listado / curl http://localhost:3000/api/tareas && echo
- Tarea Listado Por ID / curl http://localhost:3000/api/tareas/(tareaId) && echo
- Crear Tarea / curl -H "Content-type: application/x-www-form-urlencoded" \
     -d "tarea=foo" \
     -X POST \
     http://localhost:3000/api/tareas && echo
- Borrar Tarea / curl -H "Content-type: application/x-www-form-urlencoded" \
     -d "tarea-id=4" \
     -X DELETE \
     http://localhost:3000/api/tareas && echo
- Actualizar Tarea / curl -H "Content-type: application/x-www-form-urlencoded" \
     -d "id=3" \
     -d "tarea=nuevo dato" \
     -X PUT \
     http://localhost:3000/api/tareas && echo

### Crear un alias para air (live reloader)
- alias air='$(go env GOPATH)/bin/air'

---

## Produción

### Copilar y ejecutar ubuntu
- go build main.go
- ./main

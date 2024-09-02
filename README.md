# API de ejemplo
Este proyecto fue creado cómo ejemplo de demostración de conocimiento de uso del lenguaje Golang.

## Desarrollo

### Correr codigo
- go run main.go
- air (correr live reload ubuntu)

### Probar API con curl en terminal
- Inicio / curl http://localhost:3000 && echo
- Tarea Listado / curl http://localhost:3000/api/tareas && echo
- Tarea Listado Por ID / curl http://localhost:3000/api/tareas/(tareaId) && echo

### crear un alias para air (live reloader)
- alias air='$(go env GOPATH)/bin/air'

---

## Produción

### Copilar y ejecutar ubuntu
- go build main.go
- ./main

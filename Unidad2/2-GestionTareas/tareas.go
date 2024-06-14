package main

import "fmt"

type Tarea struct {
	Nombre      string
	Descripcion string
	Responsable string
	Estado      string
}

type ListaTareas struct {
	Tareas []Tarea
}

func main() {
	lista := ListaTareas{}
	lista.crearTarea("Limpieza", "Limpiar el patio", "Juan")
	lista.crearTarea("Compras", "Compra del pan diario", "Rocío")
	lista.crearTarea("Limpieza", "Limpiar la sala", "Juan")
	lista.crearTarea("Ayudas varias", "Ayudar a la abuela", "Rocío")

	lista.updateEstadoTarea(&lista.Tareas[0], "En Progreso")
	lista.updateEstadoTarea(&lista.Tareas[2], "Completada")

	lista.printTareasPendientes()
}

// Permite agregar tareas a la lista
func (lp *ListaTareas) crearTarea(nombre string, descripcion string, responsable string) {
	nuevaTarea := Tarea{Nombre: nombre, Descripcion: descripcion, Responsable: responsable, Estado: "Pendiente"}
	lp.Tareas = append(lp.Tareas, nuevaTarea)
}

func (lt *ListaTareas) updateEstadoTarea(tareaAbuscar *Tarea, nuevoestado string) {
	estadoAnterior := tareaAbuscar.Estado
	for i, tarea := range lt.Tareas {
		if tarea == *tareaAbuscar {
			lt.Tareas[i].Estado = nuevoestado
			return
		}
	}
	fmt.Printf("Estado anterior de la tarea: %s\n", estadoAnterior)
	fmt.Printf("   Nuevo estado de la tarea: %s\n", nuevoestado)
}

//Imprime el listado de productos
func (lt *ListaTareas) printTareasPendientes() {
	fmt.Println("LISTADO DE TAREAS PENDIENTES")
	fmt.Println("============================")
	for _, tarea := range lt.Tareas {
		fmt.Printf("     Nombre: %s\n", tarea.Nombre)
		fmt.Printf("Descripcion: %s\n", tarea.Descripcion)
		fmt.Printf("Responsable: %s\n", tarea.Responsable)
		fmt.Printf("     Estado: %s\n", tarea.Estado)
		fmt.Println("--------------------------------")
	}
}

package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad int
}

type ListaProductos struct {
	Productos []Producto
}

func main() {
	lista := ListaProductos{}
	lista.addProductos(Producto{Nombre: "Monitor LG 19\"", Precio: 78520.15, Cantidad: 25})
	lista.addProductos(Producto{Nombre: "Teclado Sentey", Precio: 25300.78, Cantidad: 50})
	lista.addProductos(Producto{Nombre: "Parlantes Edifier 5.0", Precio: 689500, Cantidad: 4})
	lista.addProductos(Producto{Nombre: "Mother ASUS Bx45-DL", Precio: 180320.45, Cantidad: 10})

	lista.updateCantidadProducto("Teclado Sentey", 33)

	lista.eliminarProducto("Mother ASUS Bx45-DL")

	lista.printProductos()
}

// Permite agregar productos a la lista
func (lp *ListaProductos) addProductos(p Producto) {
	lp.Productos = append(lp.Productos, p)
}

//Actualiza la cantidad de un determinado producto que coincida con el nombre por parámetro
func (lp *ListaProductos) updateCantidadProducto(nombreProducto string, nuevaCantidad int) {
	for i := 0; i < len(lp.Productos); i++ {
		if lp.Productos[i].Nombre == nombreProducto {
			lp.Productos[i].Cantidad = nuevaCantidad
			break
		}
	}
}

func (lp *ListaProductos) eliminarProducto(nombreProducto string) {
	for i, producto := range lp.Productos {
		if producto.Nombre == nombreProducto {
			//Elimino el producto de la lista
			lp.Productos = append(lp.Productos[:i], lp.Productos[i+1:]...)
			fmt.Printf("Producto borrado: %s\n", nombreProducto)
			return
		}
	}
	fmt.Printf("No se encontró el producto: %s\n", nombreProducto)
}

//Imprime el listado de productos
func (lp *ListaProductos) printProductos() {
	fmt.Println("LISTADO DE PRODUCTOS")
	fmt.Println("=====================")
	for _, producto := range lp.Productos {

		fmt.Printf("  Nombre: %s\n", producto.Nombre)
		fmt.Printf("  Precio: %.2f\n", producto.Precio)
		fmt.Printf("Cantidad: %d\n", producto.Cantidad)
		fmt.Println("--------------------------------")
	}
}

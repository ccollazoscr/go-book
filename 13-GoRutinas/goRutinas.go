package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	FirstGoRutina()
	CondicionDeCarrera()
	SincornizacionMedianteMutex()
	SincornizacionMedianteAtomix()
}

func SincornizacionMedianteAtomix() {
	fmt.Printf("\n***** SincornizacionMedianteAtomix *****\n")
	//Obtiene el número de procesadores del equipo
	//Recover se utiliza para recuperarse de un panic que se haya lanzado
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recuperandose del pánico: ", p)
		}
	}()
	tareasParalelas := runtime.GOMAXPROCS(0)
	fmt.Println("tareasParalelas:", tareasParalelas)
	v := []int{0, 1, 3, 1, 0, 7, 8, 9, 3, 3, 0, 2}

	wg := sync.WaitGroup{}
	wg.Add(tareasParalelas)

	totalSuma := int64(0)
	for t := 0; t < tareasParalelas; t++ {
		s := t
		go func() {
			defer wg.Done()
			inicio := s * len(v) / tareasParalelas
			fin := (s + 1) * len(v) / tareasParalelas
			suma := Suma(v[inicio:fin])

			//Bloquea a todas las rutinas para que unicamente una pueda actualizar el valor de totalSuma
			atomic.AddInt64(&totalSuma, int64(suma))

		}()
	}
	wg.Wait()
	//El resultado en algun momento podria dar diferente de 37 y estaría mal.
	//Esto sucede porque dos procesos acceden a la misma variable en el mismo instante y su valor queda desactualizado
	if totalSuma != 37 {
		panic(fmt.Sprint("totalSuma: ", totalSuma))
	} else {
		fmt.Println("totalSuma:", totalSuma)
	}
}

func SincornizacionMedianteMutex() {
	fmt.Printf("\n***** SincornizacionMedianteMutex *****\n")
	//Obtiene el número de procesadores del equipo
	//Recover se utiliza para recuperarse de un panic que se haya lanzado
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recuperandose del pánico: ", p)
		}
	}()
	tareasParalelas := runtime.GOMAXPROCS(0)
	fmt.Println("tareasParalelas:", tareasParalelas)
	v := []int{0, 1, 3, 1, 0, 7, 8, 9, 3, 3, 0, 2}

	mt := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(tareasParalelas)

	totalSuma := 0
	for t := 0; t < tareasParalelas; t++ {
		s := t
		go func() {
			defer wg.Done()
			inicio := s * len(v) / tareasParalelas
			fin := (s + 1) * len(v) / tareasParalelas
			suma := Suma(v[inicio:fin])

			//Bloquea a todas las rutinas para que unicamente una pueda actualizar el valor de totalSuman
			mt.Lock()
			totalSuma += suma
			mt.Unlock()
		}()
	}
	wg.Wait()
	//El resultado en algun momento podria dar diferente de 37 y estaría mal.
	//Esto sucede porque dos procesos acceden a la misma variable en el mismo instante y su valor queda desactualizado
	if totalSuma != 37 {
		panic(fmt.Sprint("totalSuma: ", totalSuma))
	} else {
		fmt.Println("totalSuma:", totalSuma)
	}
}

func CondicionDeCarrera() {
	fmt.Printf("\n***** CondicionDeCarrera *****\n")
	//Obtiene el número de procesadores del equipo
	//Recover se utiliza para recuperarse de un panic que se haya lanzado
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recuperandose del pánico: ", p)
		}
	}()
	tareasParalelas := runtime.GOMAXPROCS(0)
	fmt.Println("tareasParalelas:", tareasParalelas)
	v := []int{0, 1, 3, 1, 0, 7, 8, 9, 3, 3, 0, 2}

	wg := sync.WaitGroup{}
	wg.Add(tareasParalelas)

	totalSuma := 0
	for t := 0; t < tareasParalelas; t++ {
		s := t
		go func() {
			defer wg.Done()
			inicio := s * len(v) / tareasParalelas
			fin := (s + 1) * len(v) / tareasParalelas
			suma := Suma(v[inicio:fin])
			totalSuma += suma
		}()
	}
	wg.Wait()
	//Haciendo uso de mutex el resultado siempre va a dar 37 ya que no van a existir condiciones de carrera en la aplicación
	if totalSuma != 37 {
		panic(fmt.Sprint("totalSuma: ", totalSuma))
	} else {
		fmt.Println("totalSuma:", totalSuma)
	}
}

func FirstGoRutina() {
	fmt.Printf("\n***** FirstGoRutina *****\n")
	const numTareas = 3
	//Objeto utilizado para sincronizar varias goRutinas
	wg := sync.WaitGroup{}
	wg.Add(numTareas)

	for i := 0; i < numTareas; i++ {
		numTarea := i

		//Para invocar una GoRutina se debe anteponer la palabra go
		go func() {
			//El Done decrementa en 1 el contador que en un principío se inicioalizo en 3-
			defer wg.Done()

			fmt.Println("Ejecutando tarea", numTarea)
		}()
	}

	wg.Wait()
	fmt.Println("Completadas todas las tareas. Finanlizando")
}

func Suma(porcion []int) int {
	total := 0
	for _, n := range porcion {
		total += n
	}
	return total
}

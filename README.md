# EDD GoDrive | Proyecto 1

<div align="center">

 <a href="https://go.dev/" target="_blank"><img src="https://img.shields.io/badge/-Go-blue?style=for-the-badge&logo=go&logoColor=white"/></a> <a href="https://es.wikipedia.org/wiki/JSON" target="_blank"><img src="https://img.shields.io/badge/-JSON-gray?style=for-the-badge&logo=json&logoColor=white"/></a>

</div>

#### Estructura de Datos

> 馃懁 *Creado por **`Alan Rodrigo Pamal De Le贸n`***

Aplicaci贸n manejada por consola para el manejo de usuarios de la Facultad de Ingenier铆a de la Universidad de San Carlos de Guatemala, su representaci贸n gr谩fica por medio de grafos y reportes JSON.


### Funci贸namiento

Al ingresar el programa se muestra un men煤 para iniciar sesi贸n o cerrar el programa.

<div align="center">
    <a href="" target="_blank"><img src="https://i.imgur.com/2xCvb54.png" style="width:20rem"></a>
</div>

Si se ingresa con los datos de sesi贸n del administrador se cuenta con varias opciones, entre las cuales estan:

1. Ver estudiantes pendientes: estudiante que esten registrados o cargados masivamente en la lista de espera.
    + Mostrara una lista donde indica cuantos estudiantes hay en la cola de pendintes, dando las opciones de aceptar estudiante, rechazar o terminar el proceso de aceptaci贸n.
    <div align="center">
        <a href="" target="_blank"><img src="https://i.imgur.com/FCaG4Qz.png" style="width:20rem"></a>
    </div>
2. Ver estudiante del sistema: estudiantes que fueron aceptados de la lista de pendientes y fueron ingresados al sistema.
3. Registrar nuevo estudiante: Cargar a la lista de estudiantes pendientes manualmente.
4. Carga masiva de estudiantes: Cargar masivamente estudiantes con un archivo en formato `csv`.
5. Reportes: Mostrara un men煤 donde se puede generar una gr谩fica de las listas, colas pilas del sistema.
    + En este menu puedes generar:
        1. Estudiantes en espera: Cola de espera.
        2. Estudiantes registrados: Registrados en sistema.
        3. Reporte adminsitrador: Logs de lo que realiza el administrador con los estudiantes en espera.
        4. JSON: Reporte en formato `json` de los estudiantes en sistema.
        <div align="center">
            <a href="" target="_blank"><img src="https://i.imgur.com/XUPgE3Y.png" style="width:20rem"></a>
        </div>
6. Cerrar sesi贸n: Cerra sesi贸n y regresar al menu principal.




### Estructuras

Se utiliz贸 de 5 estructuras, 2 pilas, 1 lista doble enzada, 1 cola y un objeto de tipo Estudiante.

* Estudiante
    
    Estructura que guarda los datos de los `estudiantes` a registrar en el sistema, contando con la estructura:

    ```go
    type Estudiante struct {
        nombre   string
        carnet   string
        password string
    }
    ```

    Contando con las funciones:

    + Constrctor `New()`.
    + Funci贸n `Print()` que imprime los datos del estudiante.
    + Getters `GetNombre()`, `GerCarnet()` y `GetPassword()` para obtener los datos y poder utilizarlos en otras clases.


* Cola

    La cola con base a una lista simple cuenta con su propia estructura nodo, que tiene los atributos de apuntadores `next` y un `value` que guarda un tipo de clase Estudiante, esta es utilizada para guardar los estudiantes que estan pendientes de aceptar para su ingreso al sistema.

    ```go
    type Nodo struct {
        next  *Nodo
        value *est.Estudiante
    }
    ```

    A su vez la estructura de `Cola` que guarda los nodos `cabeza` y `cola`.

    ```go
    type Cola struct {
        cabeza *Nodo
        cola   *Nodo
    }
    ```

    Las funciones de la estructura Cola son:

    + Constructor `New()`.
    + Funci贸n `Add()` para a帽adir el nodo al final de la cola.
    + Funci贸n `Size()` que retorna el tama帽o de la lista.
    + Funci贸n `ExistStudent()` que verifica la existencia de un estudiante con ese numero de carne en la lista.
    + Funci贸n `Print()` que imprime la lista.
    + Las funciones `Dequeue()` que sirven para sacar un nodo del inicio de la cola.
    + Las funciones `ejecutar()`, `crearArchivoDot()` y `esribirArchivoDot()` los cuales son utilizados por la funci贸n `Graficar( nombre )` el cual genera una gr谩fica con `graphviz` de la lista, generando un archivo .dot y un .jpg para poder ver la gr谩fica.

    <div align="center">
        <a href="" target="_blank"><img src="https://i.imgur.com/GD7J8b8.png" style="width:25rem"></a>
    </div>


* Lista doble enlazada

    La lista doble enlazada es utilizada para guardar los estudiantes que sean aceptados y sean registrados en el sistema, cuenta con su propia estructura nodo, que tiene los atributos de apuntadores `next` y `before`, un `value` que guarda el estudiante  y una `bitacora` que guarda una pila.

    ```go
    type Nodo struct {
        next     *Nodo
        before   *Nodo
        value    *est.Estudiante
        bitacora *plbtc.Pila
    }
    ```

    A su vez la estructura de `ListaDoble` que guarda los nodos `cabeza` y `cola`.

    ```go
    type ListaDoble struct {
        cabeza *Nodo
        cola   *Nodo
    }
    ```

    Las funciones de la estructura Lista Doble son:

    + Constructor `New()`.
    + Funci贸n `AddFirst()` para a帽adir al inicio de la lista.
    + Funci贸n `AddInPos()` para a帽adir el nodo en la posici贸n indicada.
    + Funci贸n `Size()` que retorna el tama帽o de la lista.
    + Funci贸n `ExistStudent()` que verifica la existencia de un estudiante con ese numero de carne en la lista.
    + Funci贸n `GetStudentByCarnet()` que devuelve el `nodo.value` (estudiante).
    + Funci贸n `GetStackByCarnet()` que devuelve la bitacora (pila).
    + Funci贸n `Print()` que imprime la lista.
    + Las funciones `PopFirst()`, `PopLast()` y `PopPos(posicion)` que sirven para sacar un nodo de la lista.
    + Funci贸n `SortByCarnet()` que ordena la lista en orden ascendente segun el carnet del value (estudiante) en el nodo.
    + Las funciones `ejecutar()`, `crearArchivoDot()` y `esribirArchivoDot()` los cuales son utilizados por la funci贸n `Graficar( nombre )` el cual genera una gr谩fica con `graphviz` de la lista, generando un archivo .dot y un .jpg para poder ver la gr谩fica.

    <div align="center">
        <a href="" target="_blank"><img src="https://i.imgur.com/Prup6Oy.png" style="width:25rem"></a>
    </div>

* Pilas

    Las pilas son generadas para guardar tanto los registros del sistema, la primera pila (bitacora) guarda la informaci贸n de a que hora y fecha inicia sesi贸n un estudiante ya registrado. La segunda pila (registro) sirve para guardar los logs del administrador, guardando cuando acepta o rechaza une estudiante de la lista de pendientes en el sistema.

    Ambas pilas cuentan con la mismsa estructura `Pila` y sus mismas funciones, pero cambiando su `Nodo`, los cuales varian en sus atributos, los cuales serian:

    Para la pila (bitacora) de inicios de sesion se tiene un nodo con:

    ```go
    type Nodo struct {
        next  *Nodo
        fecha string
        hora  string
    }
    ```

    Para la segunda pila (registro) se tiene un nodo con:

     ```go
    type Nodo struct {
        next       *Nodo
        value      *est.Estudiante
        fecha      string
        hora       string
        aceptacion string
    }
    ```

    Las funciones de la estructuras pila constan de:

    + Constructor `New()`.
    + Funci贸n `Push()` para a帽adir el nodo al inicio de la pila.
    + Funci贸n `Size()` que retorna el tama帽o de la lista.
    + Funci贸n `Exist()` que verifica la existencia de un estudiante con ese numero de carne en la pila.
    + Funci贸n `Print()` que imprime la cola.
    + Las funciones `Pop()` que sirven para sacar un nodo de la punta de la pila.
    + Las funciones `ejecutar()`, `crearArchivoDot()` y `esribirArchivoDot()` los cuales son utilizados por la funci贸n `Graficar()` el cual genera una gr谩fica con `graphviz` de la lista, generando un archivo .dot y un .jpg para poder ver la gr谩fica.


    Pila bit谩cora:

    <div align="center">
        <a href="" target="_blank"><img src="https://i.imgur.com/SG5hlPv.png" style="width:10rem"></a>
    </div>

    Pila registro:
    
    <div align="center">
        <a href="" target="_blank"><img src="https://i.imgur.com/qPYHPNW.png" style="width:10rem"></a>
    </div>

## 📁 Config
En esta se encuentra el documento 📄const.go, en el cual se encuentran las constantes utilizadas a lo largo del proyecto, con el objetivo de facilitar el acceso y modificaciones a lo largo del proyecto por medio de acciones mínimas, sin la necesidad de revisitar cada apartado donde se encuentre el valor.
## 📁 Models
En esta se encuentre el documento 📄types.go, en el cual se encuentra la definición de las estructuras de las entidades de la puerta, estacionamiento y vehículo.
## 📁 Services
Comprende a las entidades y sus correspondientes acciones, separadas en unidades de código, que consisten en una única acción, con el objetivo de realizar la separación de responsabilidades en la mejor medida posible.
## 📁 Vehicle
En esta se encuentran las actividades que puede ejercer un vehículo.
- `📄 entrance_service.go`: Lógica de acceso al estacionamiento, controlando la disponibilidad de cajones, la cola de espera, y el estado y sentido de la puerta. También actualiza el estado de la puerta y gestiona la asignación de casillas usando variables condicionales.
- `📄 freeSpace_service.go`: Libera el espacio una vez el tiempo de estacionamiento ha finalizado.
- `📄 exit_service.go`: Lógica de salida del estacionamiento, gestionando el flujo de salida en base a la cola de vehículos y el estado de la puerta.
## 📁 Door
- `📄 door_service.go`: Cambia el estado de la puerta, notificando y ajustando el color según su estado (Entrando, Saliendo o Libre).
## 📄parking_service.go
Se encarga de definir la estructura del estacionamiento, englobando la instancia del Estacionamiento utilizando la estructura definida en los modelos y configura los estados iniciales de sus atributos, así como la inicialización de las variables condicionales y finalmente se asocian los servicios de la puerta y del vehículo.
## 📁 Simulator
En esta se encuentra el documento 📄parking_simulator.go el cual se encarga de gestionar el inicio de la simulación como una go routine (lanzada por el main), dentro de la cual, basado en las constantes establecidas, lanza go routines correspondientes a cada vehículo en intervalos de dos segundos, posteriormente, se determina el tiempo en el que cada vehículo estará estacionado, haciendo uso de los servicios para efectuar la simulación.
## 📁 ui
- `📄 scene.go`: Contiene el método `Setup()`, que inicializa la interfaz gráfica con la configuración de `view.go` y actualiza la información presentada, como ocupación y desocupación de casillas.
- `📄 view.go`: Configuración de la interfaz gráfica, creando leyendas, contenedores y organizando los elementos.
- `📄 main.go`: Crea la aplicación de Fyne, inicializa los servicios del estacionamiento, simulador e interfaz, realiza el setup de la interfaz, lanza la go routine principal (como se menciona en `📄parking_simulator.go`) y despliega la interfaz de usuario.

## Mi experiencia
Las herramientas utilizadas aportan significativamente al avance académico, además de impactar significativamente en el abanico de conocimientos que en el futuro impactaran en el desarrollo laboral, añadiendo competencias especialmente destacables gracias al temprano surgimiento de Go. 
La metodología implementada se encuentra abierta a áreas de mejora significativas, ya que, erróneamente, se comienza el desarrollo sin una arquitectura sólida, lo que añade desorganización y conflictos en la modulación y separación de responsabilidades, por lo que se destaca la importancia de trabajar bajo una arquitectura desde etapas tempranas de desarrollo de código, de esta manera, el flujo de trabajo es considerablemente más efectivo, rápido e intuitivo, posponer la modulación causa conflictos y atrasa a gran medida la finalización de una aplicación.
De igual manera, se tiene como área de mejora la identificación de responsabilidades únicas, durante el desarrollo, se asumieron ciertas acciones como un conjunto de otras que, una vez terminado el proyecto o para correcciones finales, unificar responsabilidades, intervienen con la manipulación de la misma, ya que su alto acoplamiento dan paso a que pequeñas modificaciones presenten un impacto más alto del esperado a otros módulos del código.

## ¿Como usarlo?
- Compilar: `go run main.go`
- Instalar dependencias: `go mod tidy`

## ¿Como se ve?
![image](https://github.com/user-attachments/assets/e2d344ca-ea48-43eb-a6af-094ba537e973)

## Más información sobre GO
![223238 Monica](https://github.com/user-attachments/assets/efc4e9c4-4a82-4bb8-ab7b-c5f06ccf825d)


  

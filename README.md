📁 Config
En esta se encuentra el documento 📄const.go, en el cual se encuentran las constantes utilizadas a lo largo del proyecto, con el objetivo de facilitar el acceso y modificaciones a lo largo del proyecto por medio de acciones mínimas, sin la necesidad de revisitar cada apartado donde se encuentre el valor.
📁 Models
En esta se encuentre el documento 📄types.go, en el cual se encuentra la definición de las estructuras de las entidades de la puerta, estacionamiento y vehículo.
📁 Services
Comprende a las entidades y sus correspondientes acciones, separadas en unidades de código, que consisten en una única acción, con el objetivo de realizar la separación de responsabilidades en la mejor medida posible.
📁 Vehicle
En esta se encuentran las actividades que puede ejercer un vehículo.
  📄entrance_service.go
Lógica de acceso al estacionamiento y aparcamiento, donde se toman en consideración la disponibilidad de cajones, cola de espera para entrada al estacionamiento y posición en la misma, estado y sentido de la puerta. De igual manera, se encarga de actualizar el estado de la puerta y de la lógica de asignación de una casilla. Se hace utilización de variables de tipos de condición para controlar el acceso al estacionamiento.
  📄freeSpace_service.go
Esta se encarga de liberar el espacio una vez el tiempo de estacionamiento haya finalizado.
  📄exit_service.go
Lógica de salida del estacionamiento, consiste en admitir los carros que han liberado su espacio de estacionamiento en una cola, y bajo las condiciones adecuadas, gestiona el flujo de la salida del estacionamiento, tomando en consideración el estado de la puerta.
📁 Door
  📄door_service.go
Engloba el cambio de estado de la puerta, se encarga de notificar y cambiar el color de la puerta dependiendo de su estado (Entrando, Saliendo o Libre)
📄parking_service.go
Se encarga de definir la estructura del estacionamiento, englobando la instancia del Estacionamiento utilizando la estructura definida en los modelos y configura los estados iniciales de sus atributos, así como la inicialización de las variables condicionales y finalmente se asocian los servicios de la puerta y del vehículo.
📁 Simulator
En esta se encuentra el documento 📄parking_simulator.go el cual se encarga de gestionar el inicio de la simulación como una go routine (lanzada por el main), dentro de la cual, basado en las constantes establecidas, lanza go routines correspondientes a cada vehículo en intervalos de dos segundos, posteriormente, se determina el tiempo en el que cada vehículo estará estacionado, haciendo uso de los servicios para efectuar la simulación.
📁 ui
	📄scene.go
Contiene el método Setup(), encargada de inicializar la interfaz gráfica utilizando la configuración establecida en 📄view.go, posteriormente se encarga de la actualización de la información presentada en la interfaz, como lo es la ocupación y desocupación de casillas.
  📄view.go
Configuración de la interfaz gráfica, creando las leyendas, contenedores y la disposición general de los elementos en la interfaz.
  📄main.go
Se crea la aplicación de Fyne, se inicializan los servicios del estacionamiento, simulador e interfaz, se realiza el Setup de la interfaz, se lanza la go routine principal (mencionada previamente en 📄parking_simulator.go) y finalmente se despliega la interfaz de usuario.

Fabian Jorquera 201573513-1
Rodrigo Andres Perez Jamett 201873619-8
Daniel Magaña Magaña 201873504-3

Instrucciones:
Se asume que todos los inputs de comandos estarán bien escritos y sin faltas ortográficas.
Corra los archivos en este orden: Broker, Fulcrum 1, 2 y 3, Informantes y Leia. Esto para asegurar que los servidores esten abiertos antes que los clientes intenten unirse.
Dentro de la carpeta Fulcrum hay 3 carpetas, "0", "1" y "2", el servidor que se inició escribiendo 0 guardará sus registros planetarios y logs ahí, el que se inició con 1, en "1" y lo mismo para el 2
Tener en cuenta que cuando a un servidor le llega un comando de parte de un informante y se encuentra una inconsistencia, primero se resuelve, se hace el merge con los otros servidores y despues se ejecuta el comando, por lo que el cambio (hecho por ese último comando) no se verá reflejado en los otros servidores inmediatamente
Puede que Leia no lea la informacion mas actualizada, pero se asegura monotonic reads, por lo que nunca leerá informacion mas antigua que la que ya ha leido
PD: Hicimos todo lo pedido en la tarea y lo probamos de pie a cabeza :D

*Fulcrum:
- Asegurarse de abrir los 3 servidores dentro de un periodo de maximo 2 minutos, dado que como estos cada 2 minutos hacen un merge, si falta un servidor saltará un error
- Al ejecutar un servidor, se pedirá ingresar 0, 1 o 2 en funcion de qué servidor se quiere iniciar(1, 2 o 3 respectivamente), cada uno de los 3 servidores se deberá iniciar con un numero diferente

Makefile:
    clean:   Limpia todos los posibles archivos txt de testeos pasados existentes.
    FULCRUM: clean + Iniciar servidor Fulcrum.
    BROKER:  clean + Iniciar servidor Broker.
    INFO:    clean + Iniciar cliente informante.
    LEIA:    clean + Inicia cliente de Leia.

Debera correr en el siguiente orden los comandos, en caso contrario fallará:*

*Alternativamente, puede usar 7 consolas distintas y así ver los prints que pusimos para facilitar ver los datos que maneja cada servidor.

Abrir máquina 120:
    estando en dentro de carpeta StarWars: 
        make BROKER

Abrir máquina 117:
    estando en dentro de carpeta StarWars: 
        make FULCRUM     //Se le preguntara que fulcrum abrir, ahi debe indicar 0

Abrir máquina 118:
    estando en dentro de carpeta StarWars:
        make FULCRUM      //Se le preguntara que fulcrum abrir, ahi debe indicar 1

Abrir máquina 119:
    estando en dentro de carpeta StarWars: 
        make FULCRUM      //Se le preguntara que fulcrum abrir, ahi debe indicar 2

Volver a la máquina 117:
    Apretrar CTRL+Z y luego escribir bg, así el proceso seguira ejecutandose por detras y podra iniciar Leia.

    make LEIA   //Iniciara a la cliente Leia

Volver a la máquina 118:
    Apretrar CTRL+Z y luego escribir bg, así el proceso seguira ejecutandose por detras y podra iniciar a la Informante 1 (Ahsoka).

    make INFO   //Iniciara al informante  

Volver a la máquina 119:
    Apretrar CTRL+Z y luego escribir bg, así el proceso seguira ejecutandose por detras y podra iniciar al Informante 2 (Almirante Thrawn).

    make INFO     //Iniciara al informante

SI QUIERE VOLVER A CORRER TODO DENUEVO RECUERDE MATAR LOS PROCESOS QUE ESTAN CORRIENDO EN EL BACKGROUND, PARA ELLO VAYA A LAS CONSOLAS Y HAGA LO SIGUIENTE:

	jobs		//Para verificar que hay procesos corriendo en background

	kill %1 	//Para matar el proceso corriendo en background

UNA VEZ HAGA ESTO EN TODAS LAS CONSOLAS CON PROCESOS CORRIENDO EN BACKGROUND PRODRA COMENZAR A ESCRIBIR TODOS LOS COMANDOS DE MAKE DENUEVO (OJO QUE ESTOS BORRARAN LOS REGISTROS CREADOS ANTERIORMENTE ANTES DE COMENZAR)


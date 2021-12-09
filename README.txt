Rodrigo Andres Perez Jamett
Daniel Magaña Magaña 201873504-3

Instrucciones:
Se asume que todos los inputs de comandos estarán bien escritos y sin faltas ortográficas.
Corra los archivos en este orden: Broker, Fulcrum 1, 2 y 3, Informantes y Leia. Esto para asegurar que los servidores esten abiertos antes que los clientes intenten unirse.

*Fulcrum:
- Asegurarse de abrir los 3 servidores dentro de un periodo de maximo 2 minutos, dado que como estos cada 2 minutos hacen un merge, si falta un servidor saltará un error
- Al ejecutar un servidor, se pedirá ingresar 0, 1 o 2 en funcion de qué servidor se quiere iniciar(1, 2 o 3 respectivamente), cada uno de los 3 servidores se deberá iniciar con un numero diferente

Makefile:
    clean:   Limpia todos los posibles archivos txt de testeos pasados existentes.
    FULCRUM: clean + Iniciar servidor Fulcrum.
    BROKER:  clean + Iniciar servidor Broker.
    INFO:    clean + Iniciar cliente informante.
    LEIA:    clean + Inicia cliente de Leia.


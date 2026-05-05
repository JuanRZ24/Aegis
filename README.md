# Aegis - Monitor de Servidor CLI

**Aegis** es una herramienta de monitoreo de sistema ligera escrita en Go. Proporciona información en tiempo real sobre los recursos críticos del sistema directamente en la terminal.

## Características

- **CPU**: Monitoreo de carga y uso.
- **Memoria**: Visualización del consumo de RAM.
- **Disco**: Estado de almacenamiento y particiones.s
- **Procesos**: Listado de los 5 procesos que más recursos consumen.
- **Temperatura**: Sensores térmicos del hardware (si están disponibles).
- **Interfaz Fluida**: Limpieza automática de pantalla cada 2 segundos para una visualización dinámica.

## Tecnologías Utilizadas

- **Go**: Lenguaje base para alta eficiencia.
- **Cobra**: Framework para interfaces de línea de comandos.

## Requisitos

- [Go](https://go.dev/dl/) 1.18 o superior instalado.

## Instalación y Uso

1.  Clona el repositorio:
    ```bash
    git clone https://github.com/JuanRZ24/aegis.git
    cd aegis
    ```

2.  Ejecuta la aplicación directamente:
    ```bash
    go run main.go
    ```

3.  O compila el binario:
    ```bash
    go build -o aegis
    ./aegis
    ```

## Licencia

Este proyecto se distribuye bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.

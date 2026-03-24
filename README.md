# Webp Crusher

Descripción
- Herramienta de línea de comandos en Go para convertir imágenes (.jpg, .png) a Webp y guardar las salidas en una carpeta `webp` dentro del directorio de entrada.

Requisitos
- Go 1.18 o superior.

Instalación
1. Clona el repositorio o descarga los archivos al directorio de trabajo.
2. En el directorio raíz del proyecto, compila:

Windows (executable):

```powershell
go build -o WebpCrusher.exe
```

Linux/macOS (binary name will be `WebpCrusher`):

```bash
go build -o WebpCrusher
```

Uso
- Ejecuta el binario indicando el directorio que contiene las imágenes con la bandera `-p`.

```powershell
.\WebpCrusher -p "C:\ruta\a\imagenes"
```

Comportamiento
- El programa recorre recursivamente el directorio indicado y convierte archivos con extensión `.jpg` y `.png`.
- Los archivos WebP resultantes se escriben en `webp` dentro del directorio proporcionado (por ejemplo `C:\ruta\a\imagenes\webp`).
- Si la carpeta `webp` no existe, se crea automáticamente.

Parámetros
- `-p`: Ruta al directorio raíz que contiene las imágenes de entrada. Obligatorio.

Ejemplo rápido
1. Crear y poblar una carpeta de prueba:

```powershell
mkdir C:\temp\imgs
copy C:\ruta\a\algunas\imagenes\*.jpg C:\temp\imgs\
```
2. Ejecutar la conversión:

```powershell
.\WebpCrusher -p C:\temp\imgs
```
3. Revisar los archivos convertidos en `C:\temp\imgs\webp`.

Pruebas
- El repositorio incluye un test que compila el binario desde el proyecto y verifica que la ayuda muestre la descripción del parámetro `-p`.
- Ejecutar tests localmente:

```powershell
go test ./...
```

Contribuciones
- Para contribuciones, abra un pull request con cambios claros y pruebas cuando corresponda.
- Respete la licencia del proyecto.

Licencia
- Este proyecto está bajo Creative Commons Attribution-NonCommercial 4.0 International (CC BY-NC 4.0). No se permite el uso comercial (venta) del código sin permiso explícito del autor. Consulte el archivo `LICENSE` para más detalles.

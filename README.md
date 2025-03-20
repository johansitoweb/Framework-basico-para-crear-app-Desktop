# Framework-basico-para-crear-app-Desktop

# Mi Framework Minimalista Go

Un framework minimalista para crear aplicaciones de escritorio con Go y SDL2.

## Descripción

Este framework proporciona una base para crear aplicaciones de escritorio simples con Go, utilizando SDL2 para el renderizado. Permite crear ventanas, botones y etiquetas con estilos básicos, y ofrece un layout vertical simple (`VBox`).

## Características

* **Ventanas:** Creación y gestión de ventanas.
* **Botones:** Botones con texto y funcionalidad de clic.
* **Etiquetas:** Mostrar texto estático.
* **Estilos:** Estilos básicos para personalizar la apariencia de los componentes.
* **Layout Vertical:** Organización de componentes en una columna.

## Requisitos

* Go (versión 1.16 o superior)
* SDL2 (bibliotecas de desarrollo)

## Instalación

1.  Asegúrate de tener Go y SDL2 instalados.
2.  Clona este repositorio:

    ```bash
    git clone <github.com/johansitoweb/Framework-basico-para-crear-app-Desktop>
    ```

3.  Navega al directorio del proyecto:

    ```bash
    cd <Framework-basico-para-crear-app-Desktop>
    ```

4.  Descarga las dependencias:

    ```bash
    go get [github.com/veandco/go-sdl2/sdl](https://www.google.com/search?q=https://github.com/veandco/go-sdl2/sdl)
    ```

## Uso

1.  Crea un archivo `main.go` en tu proyecto.
2.  Importa el paquete del framework.
3.  Utiliza las funciones `NewWindow`, `CreateButton`, `CreateLabel` y `DrawVBox` para crear tu interfaz de usuario.
4.  Ejecuta la aplicación:

    ```bash
    go run main.go
    ```

## Ejemplo

```go
    package main

    import (
            "fmt"
            "os"

            "[github.com/veandco/go-sdl2/sdl](https://www.google.com/search?q=https://github.com/veandco/go-sdl2/sdl)"
    )

    // ... (Código del framework) ...

    func main() {
            if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
                    fmt.Fprintf(os.Stderr, "Error al inicializar SDL: %s\n", err)
                    return
            }
            defer sdl.Quit()

            style := Style{
                    TextColor: sdl.Color{R: 255, G: 255, B: 255, A: 255},
                    BackColor: sdl.Color{R: 0, G: 0, B: 150, A: 255},
            }

            window, err := NewWindow("Mi Aplicación", 800, 600, style)
            if err != nil {
                    fmt.Fprintf(os.Stderr, "Error al crear la ventana: %s\n", err)
                    return
            }
            defer window.sdlWindow.Destroy()
            defer window.sdlRenderer.Destroy()

            button := &Button{
                    rect: sdl.Rect{X: 100, Y: 100, W: 150, H: 60},
                    text: "Clic aquí",
                    style: Style{
                            TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 255},
                            BackColor: sdl.Color{R: 200, G: 200, B: 200, A: 255},
                    },
                    onClick: func() {
                            fmt.Println("Botón clicado")
                    },
            }

            label := &Label{
                    rect:  sdl.Rect{X: 100, Y: 200, W: 200, H: 30},
                    text:  "¡Hola, mundo!",
                    style: style,
            }

            vbox := &VBox{
                    children: []interface{}{button, label},
                    rect:     sdl.Rect{X: 0, Y: 0, W: 800, H: 600},
            }

            running := true
            for running {
                    for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
                            switch event.(type) {
                            case *sdl.QuitEvent:
                                    running = false
                                    break
                            }
                    }

                    window.sdlRenderer.SetDrawColor(window.style.BackColor.R, window.style.BackColor.G, window.style.BackColor.B, window.style.BackColor.A)
                    window.sdlRenderer.Clear()

                    window.DrawVBox(vbox)

                    window.sdlRenderer.Present()
                    sdl.Delay(16)
            }
    }

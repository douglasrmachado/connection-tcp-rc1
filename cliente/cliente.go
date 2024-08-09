package main

import "fmt"
import "net"
import "os"

func main() {
    conn, err := net.Dial("tcp", "localhost:5000")
    if err != nil {
        fmt.Println("Erro ao conectar ao servidor:", err)
        os.Exit(1)
    }
    defer conn.Close()

    var message string
    fmt.Print("Digite sua mensagem: ")
    fmt.Scanln(&message)

    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Erro ao enviar mensagem:", err)
        return
    }

    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Erro ao receber resposta:", err)
        return
    }
    fmt.Println("Resposta do servidor:", string(buffer[:n]))
}

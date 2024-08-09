package main

import "fmt"
import "net"

func main() {
    listener, err := net.Listen("tcp", ":5000")
    if err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Servidor rodando...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Erro ao aceitar a conexão:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Erro ao ler dados:", err)
        return
    }
    fmt.Println("Mensagem recebida:", string(buffer[:n]))

    response := []byte("Mensagem recebida!")
    conn.Write(response)
}

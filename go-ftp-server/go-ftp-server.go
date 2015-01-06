package main

import 
(
  "net"
  "fmt"
  "strings"
  "strconv"
)

func main() {
  ln, err := net.Listen("tcp", ":1337")
  if err != nil {
    fmt.Println("Error starting listener")
    return
  }

  fmt.Println("Listening on port 1337")

  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Error opening connection with cleint")
    }

    go handleClient(conn)
  }
}

// Handle the new client
func handleClient(conn net.Conn) {
  fmt.Printf("New client at %v\net", conn.RemoteAddr())

  _, err := conn.Write([]byte("220 Welcome to this \"FTP Server\"\n"))
  if err != nil {
    fmt.Println("Error starting FTP session with client")
  }

  buff := make([]byte, 256)
  //clientAddr := ""

  for {
    readCount, err := conn.Read(buff)
    if err != nil {
      fmt.Println("Error reading from client")
      conn.Close()
    }

    if readCount > 0 {

      readString := strings.TrimSpace(string(buff[:readCount-1]))
      fmt.Printf("Received: %v\n", readString)
      cmdString := strings.Split(readString, " ")[0]
      fmt.Printf("Command: %v\n", cmdString)
      switch cmdString {
        case "USER":
          writeResponse(userCommand(), conn)
        case "SYST":
          writeResponse(systCommand(), conn)
        case "FEAT":
          writeResponse(notImplementedCommand(), conn)
        case "PWD":
          writeResponse(pwdCommand(), conn)
        case "TYPE":
          writeResponse(typeCommand(readString), conn)
        case "PASV":
          writeResponse(pasvCommand(), conn)
        case "EPSV":
          writeResponse(notImplementedCommand(), conn)
        case "PORT":
          resp, _ := portCommand(readString) 
          writeResponse(resp, conn)
        case "LIST":
          writeResponse(listCommand(), conn)
        default:
          fmt.Printf("Unknown command: %v", cmdString)
      }
    }
  }
}

func writeResponse(response string, conn net.Conn) {
  fmt.Printf("Sending: %v\n", response)
  conn.Write([]byte(response))
}

func userCommand() string{
  return "230 I don't particularly feel like writing auth code\n"
}

func systCommand() string{
  return "215 UNIX Type: L8\n"
}

func notImplementedCommand() string{
  return "502 Don't really feel like doing that\n"
}

func pwdCommand() string{
  return "257 /\n"
}

func typeCommand(cmd string) string{
  return fmt.Sprintf("200 Type set to %v\n", strings.Split(cmd, " ")[1])
}

func pasvCommand() string{
  return "227 Entering passive mode\n"
}

func portCommand(cmd string) (string, string) {
  dataString := strings.Split(cmd, " ")[1]
  dataSplit := strings.Split(dataString, ",")
  ip := fmt.Sprintf("%v:%v:%v:%v", dataSplit[0], dataSplit[1], dataSplit[2], dataSplit[3])
  portPartOne, _ := strconv.Atoi(dataSplit[4])
  portPartTwo, _ := strconv.Atoi(dataSplit[5])
  port := (portPartOne * 256) + portPartTwo

  return "200 PORT command successful\n", fmt.Sprintf("%v:%v", ip, port)
}

func listCommand() string{
  return "150 Here comes the directory listing.\n-rw-r--r--    1 ftp      ftp           178 Apr 25  2014 README\n226 Directory send OK.\n"
}
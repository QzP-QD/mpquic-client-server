package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"os"
	//"strconv"
	//"strings"

	quic "github.com/lucas-clemente/quic-go"
)

// CHUNK size to read
const CHUNK = 1024 * 10
const checkgit = "checkgit"

const flagpth string = "/home/mininet/peekaboo/flag1"

func main() {
	addr := "10.0.5.2:8000"
	server(addr)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func server(addr string) {
	//infofile, err1 := os.Open("/home/mininet/peekaboo/pathinfo1.txt")
	//if err1 != nil {
	//	panic(any(err1))
	//}
	//defer infofile.Close()
	//reader := bufio.NewReader(infofile)
	//for {
	//	str, err2 := reader.ReadString('\n')
	//	if err2 == io.EOF{
	//		break
	//	}
	//	stringSlice := strings.Split(str, " ")
	//
	//	ip := stringSlice[0]
	//	ip = strings.Replace(ip, "\n", "", -1)
	//
	//	loss, _ := strconv.ParseFloat(stringSlice[1], 64)
	//
	//	bthstr := strings.Replace(stringSlice[2], "\n", "", -1)
	//	bth, _ := strconv.ParseFloat(bthstr, 64)
	//}

	// Configure multipath
	quicConfig := &quic.Config{
		Flagpth:       flagpth,
		Missiontype:   "back",
		SchedulerName: "peek",
		CreatePaths:   true,
	}

	listener, err := quic.ListenAddr(addr, generateTLSConfig(), quicConfig)
	check(err)

	// Listen forever
	for {
		sess, err := listener.Accept()
		check(err)
		fmt.Println("Accepted connection")
		go handleClient(sess)
	}
}

func handleClient(sess quic.Session) {
	stream, err := sess.AcceptStream()
	check(err)
	defer stream.Close()

	cmd := readMessage(stream)
	if cmd != "SETUP" {
		return
	}
	fmt.Println("Received SETUP request...")
	sendMessage("OK", stream)
	filename := readMessage(stream)
	fmt.Println("Filename:", filename)
	f, err := os.Open("/home/mininet/Downloads/SampleVideos/" + filename)
	check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	_, err = io.Copy(stream, r)
	if err != nil {
		fmt.Println("Client disconnected...")
		fmt.Println(err)
	}
	fmt.Println("Exited...")
}

func sendMessage(msg string, stream quic.Stream) {
	l := uint32(len(msg))
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, l)
	stream.Write(data)
	stream.Write([]byte(msg))
}

func readMessage(stream quic.Stream) string {
	data := make([]byte, 4)
	stream.Read(data)
	l := binary.LittleEndian.Uint32(data)
	data = make([]byte, l)
	stream.Read(data)
	return string(data)
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
}

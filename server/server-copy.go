package main
//
//import (
//	"bufio"
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/tls"
//	"crypto/x509"
//	"encoding/binary"
//	"encoding/csv"
//	"encoding/pem"
//	"fmt"
//	"io"
//	"math/big"
//	"os"
//	"strconv"
//	"strings"
//
//	quic "github.com/lucas-clemente/quic-go"
//)
//
//// CHUNK size to read
//const CHUNK = 1024 * 10
//
//const flagpth string = "/home/mininet/peekaboo/flag1"
//
//func main() {
//	addr := "10.0.5.2:8000"
//	file, err := os.Open("./matrix.txt")
//	check(err)
//	defer file.Close()
//
//	var matrix [3] float64
//	reader := bufio.NewReader(file)
//	for {
//		str, err2 := reader.ReadString('\n')
//		if err2 == io.EOF {
//			break
//		}
//		stringSlice := strings.Split(str, ",")
//		for i:= 0; i < 3; i++{
//			tmp := strings.Replace(stringSlice[i], "\r", "", -1)
//			tmp = strings.Replace(tmp, "\n", "", -1)
//			tmp = strings.Replace(tmp, " ", "", -1)
//			tmpnum, err3 := strconv.ParseFloat(tmp, 64)
//			check(err3)
//			matrix[i] = tmpnum
//		}
//		fmt.Println(matrix)
//
//		server(addr, matrix)
//	}
//}
//
//func check(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//
//func server(addr string, matrix [3]float64) {
//	// Configure multipath
//	quicConfig := &quic.Config{
//		// TestAHP
//		Flagpth: flagpth,
//		AHPs: matrix,
//		Missiontype: "audio",
//		SchedulerName: "peek",
//		CreatePaths: true,
//	}
//
//	listener, err := quic.ListenAddr(addr, generateTLSConfig(), quicConfig)
//	check(err)
//
//	file, _ := os.OpenFile("/home/mininet/peekaboo/result/newnewres/ahp_testing.csv", os.O_WRONLY|os.O_APPEND,0600)
//	defer file.Close()
//
//	writer := csv.NewWriter(file)
//	matrixstr := "[ "+strconv.FormatFloat(matrix[0], 'f',5,64)+", "+strconv.FormatFloat(matrix[1], 'f',5,64)+", "+
//		strconv.FormatFloat(matrix[2], 'f',5,64)+" ] "
//	writer.Write([]string{matrixstr})
//	writer.Flush()
//
//	for index := 0; index < 3; index ++{
//		sess, err1 := listener.Accept()
//		check(err1)
//		fmt.Println("Accepted connection")
//		handleClient(sess)
//	}
//	listener.Close()
//}
//
//func handleClient(sess quic.Session) {
//	stream, err := sess.AcceptStream()
//	check(err)
//	defer stream.Close()
//
//	cmd := readMessage(stream)
//	if cmd != "SETUP" {
//		return
//	}
//	fmt.Println("Received SETUP request...")
//	sendMessage("OK", stream)
//	filename := readMessage(stream)
//	fmt.Println("Filename:", filename)
//	f, err := os.Open("/home/mininet/Downloads/SampleVideos/" + filename)
//	check(err)
//	defer f.Close()
//
//	r := bufio.NewReader(f)
//	_, err = io.Copy(stream, r)
//	if err != nil {
//		fmt.Println("Client disconnected...")
//		fmt.Println(err)
//	}
//
//	fmt.Println("Exited...")
//
//	return
//}
//
//func sendMessage(msg string, stream quic.Stream) {
//	l := uint32(len(msg))
//	data := make([]byte, 4)
//	binary.LittleEndian.PutUint32(data, l)
//	stream.Write(data)
//	stream.Write([]byte(msg))
//}
//
//func readMessage(stream quic.Stream) string {
//	data := make([]byte, 4)
//	stream.Read(data)
//	l := binary.LittleEndian.Uint32(data)
//	data = make([]byte, l)
//	stream.Read(data)
//	return string(data)
//}
//
//// Setup a bare-bones TLS config for the server
//func generateTLSConfig() *tls.Config {
//	key, err := rsa.GenerateKey(rand.Reader, 1024)
//	if err != nil {
//		panic(err)
//	}
//	template := x509.Certificate{SerialNumber: big.NewInt(1)}
//	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
//	if err != nil {
//		panic(err)
//	}
//	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
//	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
//
//	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
//	if err != nil {
//		panic(err)
//	}
//	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
//}

package main
//
//import (
//	//"bufio"
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/tls"
//	"crypto/x509"
//	"encoding/binary"
//	"encoding/pem"
//	"fmt"
//	"io"
//
//	//"io"
//	//"os"
//
//	//"os"
//	//"time"
//
//	//"io"
//	"math/big"
//	//"os"
//	"os/exec"
//
//	quic "github.com/lucas-clemente/quic-go"
//)
//
//// CHUNK size to read
//const CHUNK = 1024 * 10
//
//const flagpth string = "/home/mininet/peekaboo/flag1"
//
//const outputpth = "/home/mininet/LToutput/output.mp4"
//
//func main() {
//	addr := "10.0.5.2:8000"
//	server(addr)
//}
//
//func check(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//
//func server(addr string) {
//	// Configure multipath
//	quicConfig := &quic.Config{
//		Flagpth: flagpth,
//		Missiontype: "audio",
//		SchedulerName: "peek",
//		CreatePaths: true,
//	}
//
//	listener, err := quic.ListenAddr(addr, generateTLSConfig(), quicConfig)
//	check(err)
//
//	// Listen forever
//	for {
//		sess, err := listener.Accept()
//		check(err)
//		fmt.Println("Accepted connection")
//		go handleClient(sess)
//	}
//}
//
//func handleClient(sess quic.Session) {
//	stream, err := sess.AcceptStream()
//	check(err)
//	defer stream.Close()
//
//	mess := readMessage(stream)
//	if mess != "SETUP" {
//		return
//	}
//	fmt.Println("Received SETUP request...")
//	sendMessage("OK", stream)
//	filename := readMessage(stream)
//	fmt.Println("Filename:", filename)
//
//	filepth := "/home/mininet/Downloads/SampleVideos/" + filename
//	args := []string{"/home/mininet/MPQUIC-video-streaming/encode.py", filepth}
//	fmt.Println(filepth)
//
//	// 命令行的方式运行ecnode.py
//	cmd := exec.Command("python3", args...)
//	//执行cmd并获取stdout
//	stdout, outerr := cmd.StdoutPipe()
//	check(outerr)
//	err = cmd.Start()
//	check(err)
//
//	//r := bufio.NewReader(stdout)
//	//scanner := bufio.NewScanner(r)
//	//for scanner.Scan(){
//	//	_, tmperr := os.Stat(outputpth)
//	//	if tmperr == nil{
//	//		break
//	//	}
//	//
//	//	curmes := scanner.Text()
//	//	sendMessage(curmes, stream)
//	//}
//
//	messg, err11 := io.Copy(stream, stdout)
//	//fmt.Println("Sending...")
//	if err11 != nil {
//		fmt.Println("Client disconnected...")
//		fmt.Println(err)
//	}
//	fmt.Println(messg)
//
//	//for scanner.Scan() {
//	//	//TODO:小trick，后面要改下收发程序
//	//	_, tmperr := os.Stat(outputpth)
//	//	if tmperr == nil{
//	//		break
//	//	}
//	//	messg, err11 := io.Copy(stream, r)
//	//	//fmt.Println("Sending...")
//	//	if err11 != nil {
//	//		fmt.Println("Client disconnected...")
//	//		fmt.Println(err)
//	//	}
//	//	fmt.Println(messg)
//	//}
//
//	fmt.Println("Exited...")
//}
//
//func sendMessage(msg string, stream quic.Stream) {
//	l := uint32(len(msg))
//	fmt.Println(l)
//	data := make([]byte, 4)
//	binary.LittleEndian.PutUint32(data, l)
//	stream.Write(data)
//	stream.Write([]byte(msg))
//}
//
//func sendMessagebytes(msg []byte, stream quic.Stream) {
//	l := uint32(len(msg))
//	data := make([]byte, 4)
//	binary.LittleEndian.PutUint32(data, l)
//	stream.Write(data)
//	stream.Write(msg)
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

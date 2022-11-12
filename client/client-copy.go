package main
//
//import (
//	//"bufio"
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/tls"
//	"crypto/x509"
//	"encoding/binary"
//	"encoding/csv"
//
//	//"encoding/csv"
//	"strconv"
//	//"strings"
//
//	//"encoding/csv"
//	"encoding/pem"
//	"fmt"
//	"io"
//	"math/big"
//	"os"
//	"os/exec"
//	//"strconv"
//	"time"
//
//	"github.com/lucas-clemente/quic-go"
//
//	//"encoding/json"
//)
//
//// CHUNK size to read
//const CHUNK = 1024 * 10
//
//const flagpth string = "/home/mininet/peekaboo/flag2"
//
//func main() {
//	addr := "10.0.5.2:8000"
//	filename := "300M.mp4"
//
//	//client(addr, filename, "peek")
//	scheculerNames := [1]string{"peek"}
//	//scheculerNames := [4]string{"peek", "rtt","random","ecf","rr","blest","ahp"，"dqnAgent"，"lowband‘}
//	//file, _ := os.OpenFile("/home/mininet/peekaboo/result/newnewres/ahp_testing.csv", os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
//	//defer file.Close()
//	//
//	//writer := csv.NewWriter(file)
//
//	for _, scheculername := range scheculerNames{
//		//writer.Write([]string{scheculername})
//		//writer.Flush()
//		fmt.Println("Testing "+scheculername)
//		for {
//			os.Remove("video.ts")
//			client(addr, filename, scheculername)
//			//writer.Write([]string{strconv.FormatFloat(tasktime, 'f', 5, 64)})
//			//writer.Flush()
//			time.Sleep(time.Duration(1)*time.Second)
//		}
//	}
//}
//
//func check(err error) {
//	if err != nil {
//		panic(any(err))
//	}
//}
//
////video, audio, file
//func client(addr string, filename string, scheculerName string) {
//	// setup multipath configuration
//	quicConfig := &quic.Config{
//		Flagpth: flagpth,
//		Missiontype: "audio",
//		SchedulerName: scheculerName,
//		CreatePaths: true,
//	}
//	// connect to server
//	session, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, quicConfig)
//	check(err)
//	stream, err := session.OpenStreamSync()
//	defer stream.Close()
//
//	// initiate SETUP
//	sendMessage("SETUP", stream)
//
//	// send filename
//	sendMessage(filename, stream)
//
//	// get reponse
//	msg := readMessage(stream)
//	if msg != "OK" {
//		return
//	}
//
//	//recivefilepth := "video.ts"
//	// start ffmpeg
//	// ffmpeg := exec.Command("ffplay", "-f", "mp4", "-i", "pipe:")
//	ffmpeg := exec.Command("ffmpeg", "-i", "pipe:", "-c", "copy", "video.ts")
//	inpipe, err := ffmpeg.StdinPipe()
//	check(err)
//	starttime := time.Now()
//	err = ffmpeg.Start()
//
//	// write
//	_, err = io.Copy(inpipe, stream)
//	if err != nil {
//		fmt.Println("Stream closed...")
//		fmt.Println(err)
//	}
//	fmt.Println("Exited...")
//	durationtime := time.Since(starttime)
//	file, _ := os.OpenFile("/home/mininet/peekaboo/result/newnewres/ahp_testing.csv", os.O_WRONLY|os.O_APPEND,0600)
//	defer file.Close()
//
//	writer := csv.NewWriter(file)
//	fmt.Println("Cost Time ", durationtime.Minutes())
//	writer.Write([]string{strconv.FormatFloat(durationtime.Minutes(), 'f', 5, 64)})
//	writer.Flush()
//
//	return
//}
//
//func sendMessage(msg string, stream quic.Stream) {
//	// utility for sending control messages
//	l := uint32(len(msg))
//	data := make([]byte, 4)
//	binary.LittleEndian.PutUint32(data, l)
//	stream.Write(data)
//	stream.Write([]byte(msg))
//}
//
//func readMessage(stream quic.Stream) string {
//	// utility for receiving control messages
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

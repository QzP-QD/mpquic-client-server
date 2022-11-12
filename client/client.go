package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/csv"
	"strconv"
	"strings"

	//"encoding/csv"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"os"
	"os/exec"
	//"strconv"
	"time"

	"github.com/lucas-clemente/quic-go"

	//"encoding/json"
)

// CHUNK size to read
const CHUNK = 1024 * 10

const flagpth string = "/home/mininet/peekaboo/flag2"

func main() {
	addr := "10.0.5.2:8000"
	filename := "300M.mp4"

	//client(addr, filename, "peek")
	scheculerNames := [1]string{"peek"}
	//scheculerNames := [4]string{"peek", "rtt","random","ecf","rr","blest","ahp"，"dqnAgent"，"lowband‘}
	file, _ := os.OpenFile("/home/mininet/peekaboo/result/newnewnewres/low_result_multi_"+scheculerNames[0] + ".csv", os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, scheculername := range scheculerNames{
		writer.Write([]string{scheculername})
		writer.Flush()
		fmt.Println("Testing "+scheculername)
		for i := 0; i < 1 ; i++ {
			fmt.Print(strconv.FormatInt(int64(i),10) + ">>time>>")
			os.Remove("video.ts")
			//TODO:记录下
			//go recordprocess(int64(i), scheculername)
			tasktime := client(addr, filename, scheculername)
			fmt.Print("Cost Time(min):  ")
			fmt.Println(tasktime)
			writer.Write([]string{strconv.FormatFloat(tasktime, 'f', 5, 64)})
			writer.Flush()
		}
	}
}

func recordprocess(index int64, scheculername string){
	var nameslice []string
	indexes := [11]string{"0", "1","2", "3","4","5","6","7","8","9","10"}
	nameslice = append(nameslice, "./testresult/recordprocess")
	nameslice = append(nameslice, indexes[index])
	nameslice = append(nameslice, scheculername)
	nameslice = append(nameslice, ".txt")
	recordfilename := strings.Join(nameslice, "_")

	recordfile, err1:= os.OpenFile(recordfilename, os.O_WRONLY|os.O_CREATE,0600)
	if err1 != nil {
		panic(any(err1))
	}
	defer recordfile.Close()

	recordwriter := bufio.NewWriter(recordfile)
	starttime := time.Now()
	flag := false
	for {
		fi, err := os.Stat("video.ts")
		if err == nil {
			flag = true
			curtime := time.Since(starttime).Seconds()
			filesize := fi.Size()

			var slice []string
			slice = append(slice, strconv.FormatInt(int64(curtime),10))
			slice = append(slice, strconv.FormatInt(filesize,10))
			outputstr := strings.Join(slice, ",")
			outputstr += "\n"		// 换行符

			recordwriter.WriteString(outputstr)
			recordwriter.Flush()
		}else{
			if flag {
				break
			}
			starttime = time.Now()
			continue
		}
		time.Sleep(time.Duration(3)*time.Second)
	}
	fmt.Println(recordfilename + "finish!")
}

func check(err error) {
	if err != nil {
		panic(any(err))
	}
}

//video, audio, file
func client(addr string, filename string, scheculerName string) float64 {
	// setup multipath configuration
	quicConfig := &quic.Config{
		Flagpth: flagpth,
		Missiontype: "back",
		SchedulerName: scheculerName,
		CreatePaths: true,
	}
	// connect to server
	session, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, quicConfig)
	check(err)
	stream, err := session.OpenStreamSync()
	defer stream.Close()

	// initiate SETUP
	sendMessage("SETUP", stream)

	// send filename
	sendMessage(filename, stream)

	// get reponse
	msg := readMessage(stream)
	if msg != "OK" {
		return  -1
	}

	//recivefilepth := "video.ts"
	// start ffmpeg
	// ffmpeg := exec.Command("ffplay", "-f", "mp4", "-i", "pipe:")
	ffmpeg := exec.Command("ffmpeg", "-i", "pipe:", "-c", "copy", "video.ts")
	inpipe, err := ffmpeg.StdinPipe()
	check(err)
	starttime := time.Now()
	err = ffmpeg.Start()

	// write
	_, err = io.Copy(inpipe, stream)
	if err != nil {
		fmt.Println("Stream closed...")
		fmt.Println(err)
	}
	fmt.Println("Exited...")
	durationtime := time.Since(starttime)

	return durationtime.Minutes()
}

func sendMessage(msg string, stream quic.Stream) {
	// utility for sending control messages
	l := uint32(len(msg))
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, l)
	stream.Write(data)
	stream.Write([]byte(msg))
}

func readMessage(stream quic.Stream) string {
	// utility for receiving control messages
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

package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: parselog src_file")
		os.Exit(-1)
	}

	data, _ := ioutil.ReadFile(os.Args[1])
	buffer := bytes.NewBuffer(data)
	for buffer.Len() > 0 {
		// 消息体
		dataLen := readInt32(buffer)
		dataBytes := readBytes(buffer, dataLen)
		unzipedDataBytes := unzip(dataBytes)
		// 参数
		paramLen := readInt32(buffer)
		paramBytes := readBytes(buffer, paramLen)
		// 接收时间
		timestap := readInt64(buffer)

		fmt.Printf("%s\t%s\t%d\n", base64.StdEncoding.EncodeToString(unzipedDataBytes), string(paramBytes), timestap)
	}
}

func readInt32(buffer *bytes.Buffer) int32 {
	var i int32
	binary.Read(buffer, binary.BigEndian, &i)
	return i
}

func readInt64(buffer *bytes.Buffer) int64 {
	var i int64
	binary.Read(buffer, binary.BigEndian, &i)
	return i
}

func readBytes(buffer *bytes.Buffer, len int32) []byte {
	temp := make([]byte, len)
	buffer.Read(temp)
	return temp
}

func unzip(zipData []byte) []byte {
	z, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil
	}

	buffer := bytes.NewBufferString("")
	for _, zf := range z.File {
		f, err := zf.Open()
		if err != nil {
			continue
		}
		io.Copy(buffer, f)
	}

	return buffer.Bytes()
}

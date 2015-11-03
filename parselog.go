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
	r := bytes.NewReader(data)
	for r.Len() > 0 {
		// 消息体
		dataLen := readInt32(r)
		dataBytes := readBytes(r, dataLen)
		unzipedDataBytes := unzip(dataBytes)
		// 参数
		paramLen := readInt32(r)
		paramBytes := readBytes(r, paramLen)
		// 接收时间
		timestap := readInt64(r)

		fmt.Printf("%s\t%s\t%d\n", base64.StdEncoding.EncodeToString(unzipedDataBytes), string(paramBytes), timestap)
	}
}

func readInt32(r *bytes.Reader) int32 {
	var i int32
	binary.Read(r, binary.BigEndian, &i)
	return i
}

func readInt64(r *bytes.Reader) int64 {
	var i int64
	binary.Read(r, binary.BigEndian, &i)
	return i
}

func readBytes(r *bytes.Reader, len int32) []byte {
	temp := make([]byte, len)
	r.Read(temp)
	return temp
}

func unzip(zipData []byte) []byte {
	zr, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil
	}

	buffer := bytes.NewBufferString("")
	for _, zf := range zr.File {
		f, err := zf.Open()
		if err != nil {
			continue
		}
		io.Copy(buffer, f)
	}

	return buffer.Bytes()
}

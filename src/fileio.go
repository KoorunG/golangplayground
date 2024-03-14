package main

import (
	"io"
	"os"
)

/**
 * 파일 읽고 쓰기
 */

func main() {
	// 1. 파일 열기
	fi, err := os.Open("out/test1.txt")
	// 에러처리
	if err != nil {
		panic(err)
	}
	// 작업이 수행된 뒤 Close() 호출
	defer fi.Close()

	// 2. 파일 생성하기
	fo, err := os.Create("out/test2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	// make()로 byte buffer 생성
	buff := make([]byte, 1024)

	for {
		// Read
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		// 파일의 끝일 경우 루프 종료
		if cnt == 0 {
			break
		}
		// 쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}

package main

import (
    "fmt"
    "github.com/lazywei/go-opencv/opencv"
)

func main() {
    // 이미지 파일 읽기
    img := opencv.LoadImage("test.jpg")

    // 윈도우 생성 및 이미지 출력
    win := opencv.NewWindow("Image")
    win.ShowImage(img)

    // 키 이벤트 대기
    opencv.WaitKey(0)

    // 메모리 해제
    img.Release()
    win.Destroy()
}
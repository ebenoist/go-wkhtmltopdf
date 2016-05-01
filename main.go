package main

// Cribbed from https://github.com/hailocab/wkhtmltopdf-go/blob/master/wkhtmltopdf/pdf_c_api.go

//#cgo CFLAGS: -I/usr/local/include
//#cgo LDFLAGS: -L/usr/local/lib -lwkhtmltox -Wall -ansi -pedantic -ggdb
//#include <stdbool.h>
//#include <wkhtmltox/pdf.h>
import "C"

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"unsafe"
)

const WKHTMLTOPDF_CMD = "wkhtmltopdf"

func main() {
	var html string
	var cbinding bool

	flag.BoolVar(&cbinding, "cbinding", false, "Perform with c binding")
	flag.StringVar(&html, "html", "<b>Hello, World</b>", "HTML to convert to PDF")
	flag.Parse()

	if cbinding {
		convertWithBindings(html)
	} else {
		HTMLToPDF(html)
	}
}

func convertWithBindings(html string) {
	fmt.Printf("with C")
	C.wkhtmltopdf_init(C.false)
	globalSettings := C.wkhtmltopdf_create_global_settings()
	objectSettings := C.wkhtmltopdf_create_object_settings()
	converter := C.wkhtmltopdf_create_converter(globalSettings)
	C.wkhtmltopdf_add_object(converter, objectSettings, C.CString(html))
	C.wkhtmltopdf_convert(converter)

	emptyString := C.CString("")
	outputBuffer := (**C.uchar)(unsafe.Pointer(&emptyString))
	length := C.wkhtmltopdf_get_output(converter, outputBuffer)

	pdf := C.GoStringN((*C.char)(unsafe.Pointer(*outputBuffer)), C.int(length))

	fmt.Printf("%s", pdf)
}

func HTMLToPDF(html string) {
	fmt.Printf("with shell")
	cmd := exec.Command(WKHTMLTOPDF_CMD, "-", "-")

	var out bytes.Buffer
	cmd.Stdout = &out

	writer, _ := cmd.StdinPipe()

	// Write to the buffer asynchronously so that it does not hang on large payloads
	go func() {
		writer.Write([]byte(html))
		writer.Close()
	}()

	cmd.Run()

	fmt.Printf("%s", out.Bytes())
}

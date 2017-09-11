/*

Golang h264,aac decoder/encoder libav wrapper

	d, err = codec.NewAACEncoder()
	data, err = d.Encode(samples)

	d, err = codec.NewAACDecoder(aaccfg)
	samples, err = d.Decode(data)

	var img *image.YCbCr
	d, err = codec.NewH264Encoder(640, 480)
	img, err = d.Encode(img)

	d, err = codec.NewH264Decoder(pps)
	img, err = d.Decode(nal)
*/
package codec

import (
	"reflect"
	"unsafe"

	/*
		#cgo LDFLAGS: -Wdeprecated-declarations -lavformat -lavutil -lavcodec
		#cgo CFLAGS: -Wdeprecated-declarations
		#include <libavutil/avutil.h>
		#include <libavformat/avformat.h>

		static void libav_init() {
			av_register_all();
			av_log_set_level(AV_LOG_FATAL);
		}
	*/
	"C"
)
import "log"

func init() {
	C.libav_init()
	log.Println("avutil_version", uint(C.avutil_version()))
	log.Println("avformat_version", uint(C.avformat_version()))
	log.Println("avcodec_version", uint(C.avcodec_version()))

}

func fromCPtr(buf unsafe.Pointer, size int) (ret []uint8) {
	hdr := (*reflect.SliceHeader)((unsafe.Pointer(&ret)))
	hdr.Cap = size
	hdr.Len = size
	hdr.Data = uintptr(buf)
	return
}

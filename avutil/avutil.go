// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <stdlib.h>
//#include <libavutil/channel_layout.h>
//#include <libavutil/samplefmt.h>
//#include <libavutil/audio_fifo.h>
//#include <libavutil/mathematics.h>
//#include <libavutil/time.h>
//#include <libavutil/imgutils.h>
import "C"
import (
	"unsafe"
)

type (
	Options       C.struct_AVOptions
	AvTree        C.struct_AVTree
	Rational      C.struct_AVRational
	MediaType     C.enum_AVMediaType
	AvPictureType C.enum_AVPictureType
	PixelFormat   C.enum_AVPixelFormat
	SampleFormat  C.enum_AVSampleFormat
	File          C.FILE
	AudioFifo     C.struct_AVAudioFifo
	AVRounding    C.enum_AVRounding
)

const (
	AV_SAMPLE_FMT_U8   = int(C.AV_SAMPLE_FMT_U8)
	AV_SAMPLE_FMT_S16  = int(C.AV_SAMPLE_FMT_S16)
	AV_SAMPLE_FMT_S32  = int(C.AV_SAMPLE_FMT_S32)
	AV_SAMPLE_FMT_FLT  = int(C.AV_SAMPLE_FMT_FLT)
	AV_SAMPLE_FMT_DBL  = int(C.AV_SAMPLE_FMT_DBL)
	AV_SAMPLE_FMT_U8P  = int(C.AV_SAMPLE_FMT_U8P)
	AV_SAMPLE_FMT_S16P = int(C.AV_SAMPLE_FMT_S16P)
	AV_SAMPLE_FMT_S32P = int(C.AV_SAMPLE_FMT_S32P)
	AV_SAMPLE_FMT_FLTP = int(C.AV_SAMPLE_FMT_FLTP)
	AV_SAMPLE_FMT_DBLP = int(C.AV_SAMPLE_FMT_DBLP)
	AV_SAMPLE_FMT_S64  = int(C.AV_SAMPLE_FMT_S64)
	AV_SAMPLE_FMT_S64P = int(C.AV_SAMPLE_FMT_S64P)
	AV_SAMPLE_FMT_NB   = int(C.AV_SAMPLE_FMT_NB)
	AV_ROUND_ZERO      = int(C.AV_ROUND_ZERO)
	AV_ROUND_INF       = int(C.AV_ROUND_INF)
	AV_ROUND_DOWN      = int(C.AV_ROUND_DOWN)
	AV_ROUND_UP        = int(C.AV_ROUND_UP)
	AV_TIME_BASE       = int(C.AV_TIME_BASE)
)

//Return the LIBAvUTIL_VERSION_INT constant.
func AvutilVersion() uint {
	return uint(C.avutil_version())
}

//Return the libavutil build-time configuration.
func AvutilConfiguration() string {
	return C.GoString(C.avutil_configuration())
}

//Return the libavutil license.
func AvutilLicense() string {
	return C.GoString(C.avutil_license())
}

//Return a string describing the media_type enum, NULL if media_type is unknown.
func AvGetMediaTypeString(mt MediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

//Return a single letter to describe the given picture type pict_type.
func AvGetPictureTypeChar(pt AvPictureType) string {
	return string(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

//Return x default pointer in case p is NULL.
func AvXIfNull(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

//Compute the length of an integer list.
func AvIntListLengthForSize(e uint, l int, t uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(e), unsafe.Pointer(&l), (C.uint64_t)(t)))
}

//Open a file using a UTF-8 filename.
func AvFopenUtf8(p, m string) *File {
	f := C.av_fopen_utf8(C.CString(p), C.CString(m))
	return (*File)(f)
}

//Return the fractional representation of the internal time base.
func AvGetTimeBaseQ() Rational {
	return (Rational)(C.av_get_time_base_q())
}

func AvGetChannelLayoutNbChannels(channelLayout uint64) int {
	return (int)(C.av_get_channel_layout_nb_channels((C.uint64_t)(channelLayout)))
}

func AvGetDefaultChannelLayout(nbChannels int) int {
	return (int)(C.av_get_default_channel_layout((C.int)(nbChannels)))
}

func AvAudioFifoAlloc(sampleFmt SampleFormat, channels int, nbSamples int) *AudioFifo {
	return (*AudioFifo)((C.av_audio_fifo_alloc((C.enum_AVSampleFormat)(sampleFmt), (C.int)(channels), (C.int)(nbSamples))))
}

func AvAudioFifoSize(af *AudioFifo) int {
	return int(C.av_audio_fifo_size((*C.struct_AVAudioFifo)(af)))
}

func AvSamplesAlloc(audioData **uint8, linesize *int, nbChannels int, nbSamples int, sampleFmt SampleFormat, align int) int {
	return int(C.av_samples_alloc((**C.uint8_t)(unsafe.Pointer(audioData)), (*C.int)(unsafe.Pointer(linesize)), C.int(nbChannels), C.int(nbSamples), (C.enum_AVSampleFormat)(sampleFmt), C.int(align)))
}

func AvAudioFifoReAlloc(af *AudioFifo, nbSamples int) int {
	return (int)(C.av_audio_fifo_realloc((*C.struct_AVAudioFifo)(af), (C.int)(nbSamples)))
}

func AvAudioFifoWrite(af *AudioFifo, data *unsafe.Pointer, nbSamples int) int {
	return (int)(C.av_audio_fifo_write((*C.struct_AVAudioFifo)(af), data, (C.int)(nbSamples)))
}

func AvAudioFifoRead(af *AudioFifo, data *unsafe.Pointer, nbSamples int) int {
	return (int)(C.av_audio_fifo_read((*C.struct_AVAudioFifo)(af), data, (C.int)(nbSamples)))
}

func AvOptSetBin(obj unsafe.Pointer, name string, value *uint8, len int, searchFlags int) int {
	return int(C.av_opt_set_bin(obj, C.CString(name), (*C.uint8_t)(value), C.int(len), C.int(searchFlags)))
}

func AvRescaleRnd(a int64, b int64, c int64, rnd AVRounding) int64 {
	return int64(C.av_rescale_rnd(C.int64_t(a), C.int64_t(b), C.int64_t(c), C.enum_AVRounding(rnd)))
}

func AvUsleep(sleep int) int {
	return int(C.av_usleep(C.unsigned(sleep)))
}

func AvImageAlloc(frame *Frame, width int, height int, pixFmt int, align int) int {
	return int(C.av_image_alloc((**C.uint8_t)(unsafe.Pointer(frame.AvFrameGetData())), (*C.int)(unsafe.Pointer(frame.AvFrameGetLineSize())), C.int(width), C.int(height), (C.enum_AVPixelFormat)(pixFmt), C.int(align)))
}

func AvGetSampleFmtName(sampleFmt SampleFormat) string {
	return C.GoString(C.av_get_sample_fmt_name((C.enum_AVSampleFormat)(sampleFmt)))
}

func AvRescaleQ(a int64, bq Rational, cq Rational) int64 {
	return int64(C.av_rescale_q(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq)))
}

func (r *Rational) Set(num, den int) {
	r.num, r.den = C.int(num), C.int(den)
}

func AvGetTime() int64 {
	return (int64)(C.av_gettime())
}

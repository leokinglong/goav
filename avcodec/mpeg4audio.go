// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavcodec/mpeg4audio.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"unsafe"
)

type (
	MPEG4AudioConfig C.struct_MPEG4AudioConfig
)

func AvprivMpeg4AudioGetConfig2(c *MPEG4AudioConfig, buf *uint8, size int, syncExtension int, logctx unsafe.Pointer) int {
	return int(C.avpriv_mpeg4audio_get_config2((*C.struct_MPEG4AudioConfig)(c), (*C.uint8_t)(buf), C.int(size), C.int(syncExtension), logctx))
}

func (mpeg4 *MPEG4AudioConfig)  AvprivMpeg4AudioGetSampleRate() int {
	return *((*int)(unsafe.Pointer(&mpeg4.sample_rate)))
}

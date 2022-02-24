// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avutil

/*
	#cgo pkg-config: libavutil
	#include <stdlib.h>
    #include <libavformat/avformat.h>
    #include <libavcodec/avcodec.h>
    #include <libavutil/avutil.h>
    #include <libavutil/opt.h>
    #include <libavdevice/avdevice.h>
*/
import "C"

type (
	Context                       C.struct_AVCodecContext
)

func AvOptSet(ctxt *Context ,name string,val string, searchFlags int) int{
	return int(C.av_opt_set(((*C.struct_AVCodecContext)(ctxt)).priv_data,  C.CString(name), C.CString(val), C.int(searchFlags)))
}

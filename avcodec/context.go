// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

//func (ctxt *Context) AvCodecGetPktTimebase() Rational {
//	return Rational(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctxt)))
//}

func (ctxt *Context) AvCodecGetTimebase() Rational {
	return Rational(((*C.struct_AVCodecContext)(ctxt)).time_base)
}

func (ctxt *Context) AvCodecGetTimebaseNum() int {
	return (int)(*((*int)(unsafe.Pointer(&ctxt.time_base.num))))
}

func (ctxt *Context) AvCodecGetTimebaseDen() int {
	return (int)(*((*int)(unsafe.Pointer(&ctxt.time_base.den))))
}

func (ctxt *Context) AvCodecGetSamAspRatioNum() int {
	return (int)(*((*int)(unsafe.Pointer(&ctxt.sample_aspect_ratio.num))))
}

func (ctxt *Context) AvCodecGetSamAspRatioDen() int {
	return (int)(*((*int)(unsafe.Pointer(&ctxt.sample_aspect_ratio.den))))
}

// AvCodecGetPktTimebase2 returns the timebase rational number as numerator and denominator
//func (ctxt *Context) AvCodecGetPktTimebase2() Rational {
//	return ctxt.AvCodecGetPktTimebase()
//}

//func (ctxt *Context) AvCodecSetPktTimebase(r Rational) {
//	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctxt), (C.struct_AVRational)(r))
//}

//func (ctxt *Context) AvCodecGetCodecDescriptor() *Descriptor {
//	return (*Descriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctxt)))
//}

//func (ctxt *Context) AvCodecSetCodecDescriptor(d *Descriptor) {
//	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecDescriptor)(d))
//}

//func (ctxt *Context) AvCodecGetLowres() int {
//	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctxt)))
//}

//func (ctxt *Context) AvCodecSetLowres(i int) {
//	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctxt), C.int(i))
//}

//func (ctxt *Context) AvCodecGetSeekPreroll() int {
//	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctxt)))
//}

//func (ctxt *Context) AvCodecSetSeekPreroll(i int) {
//	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctxt), C.int(i))
//}

//func (ctxt *Context) AvCodecGetChromaIntraMatrix() *uint16 {
//	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt)))
//}

//func (ctxt *Context) AvCodecSetChromaIntraMatrix(t *uint16) {
//	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt), (*C.uint16_t)(t))
//}

//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctxt *Context) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(&ctxt)))
}

//Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
//func (ctxt *Context) AvcodecGetContextDefaults3(c *Codec) int {
//	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c)))
//}

//Copy the settings of the source Context into the destination Context.
//func (ctxt *Context) AvcodecCopyContext(ctxt2 *Context) int {
//	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecContext)(ctxt2)))
//}

//Initialize the Context to use the given Codec
func (ctxt *Context) AvcodecOpen2(c *Codec, d **Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given Context and free all the data associated with it (but not the Context itself).
func (ctxt *Context) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctxt)))
}

//The default callback for Context.get_buffer2().
func (s *Context) AvcodecDefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(s), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctxt *Context) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctxt *Context) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode the audio frame of size avpkt->size from avpkt->data into frame.
//func (ctxt *Context) AvcodecDecodeAudio4(f *Frame, g *int, a *Packet) int {
//	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
//}

//Decode the video frame of size avpkt->size from avpkt->data into picture.
//func (ctxt *Context) AvcodecDecodeVideo2(p *Frame, g *int, a *Packet) int {
//	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
//}

//Decode a subtitle message.
func (ctxt *Context) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Encode a frame of audio.
//func (ctxt *Context) AvcodecEncodeAudio2(p *Packet, f *Frame, gp *int) int {
//	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
//}

//Encode a frame of video
//func (ctxt *Context) AvcodecEncodeVideo2(p *Packet, f *Frame, gp *int) int {
//	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
//}

func (ctxt *Context) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctxt), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (ctxt *Context) AvcodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctxt), (*C.enum_AVPixelFormat)(f)))
}

//Reset the internal decoder state / flush internal buffers.
func (ctxt *Context) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctxt))
}

//Return audio frame duration.
func (ctxt *Context) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctxt), C.int(f)))
}

func (ctxt *Context) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctxt)))
}

//Parse a packet.
//func (ctxt *Context) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
//	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
//}

//func (ctxt *Context) AvParserChange(ctxtp *ParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
//	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
//}

func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

//func (p *Parser) AvParserNext() *Parser {
//	return (*Parser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
//}

//func (p *Parser) AvRegisterCodecParser() {
//	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
//}

func (ctxt *Context) SetTimebase(num1 int, den1 int) {
	ctxt.time_base.num = C.int(num1)
	ctxt.time_base.den = C.int(den1)
}

func (ctxt *Context) SetTimebase2(rational Rational) {
	ctxt.time_base = C.struct_AVRational(rational)
}

func (ctxt *Context) SetPktTimebase(rational Rational) {
	ctxt.pkt_timebase = C.struct_AVRational(rational)
}

func (ctxt *Context) SetStrictStdCompliance(com int) {
	ctxt.strict_std_compliance = C.int(com)
}

func (ctxt *Context) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int, bitrate int64) {
	ctxt.width = C.int(width)
	ctxt.height = C.int(height)
	ctxt.bit_rate = C.long(bitrate)
	ctxt.gop_size = C.int(gopSize)
	if hasBframes {
		ctxt.has_b_frames = 1
	} else {
		ctxt.has_b_frames = 0
		ctxt.max_b_frames = 0
	}
	// ctxt.extradata = nil
	// ctxt.extradata_size = 0
	// ctxt.channels = 0
	ctxt.pix_fmt = int32(pxlFmt)
	//C.av_opt_set(ctxt.priv_data, "preset", "ultrafast", 0)
}

func (ctxt *Context) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	ctxt.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10, 7000000)
}

func (ctxt *Context) AvcodecSendPacket(packet *Packet) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(packet)))
}

func (context *Context) AvcodecReceivePacket(dest *Packet) int {
	return int(C.avcodec_receive_packet((*C.struct_AVCodecContext)(context), (*C.struct_AVPacket)(dest)))
	//fmt.Printf("[%p] AvcodecReceivePacket result %v\n",context,ret)
}

func (ctxt *Context) AvcodecReceiveFrame(frame *Frame) int {
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(frame)))
}

func (ctxt *Context) AvcodecSendFrame(frame *Frame) int {
	return int(C.avcodec_send_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(frame)))
}

func (ctxt *Context) SetFrameRate(frameRate Rational) {
	ctxt.framerate = (C.struct_AVRational)(frameRate)
}

func (ctxt *Context) GetFrameRate() Rational {
	return Rational(ctxt.framerate)
}

func (ctxt *Context) GetFrameRateInt() int {
	if ctxt.framerate.den == 0 {
		return 25
	}
	return (int)(ctxt.framerate.num / ctxt.framerate.den)
}

func (ctxt *Context) SetFrameRate2(num int, den int) {
	ctxt.framerate.num = (C.int)(num)
	ctxt.framerate.den = (C.int)(den)
}

func (ctxt *Context) SetSampleAspectRatio(sampleAspectRatio Rational) {
	ctxt.sample_aspect_ratio = (C.struct_AVRational)(sampleAspectRatio)
}

func (ctxt *Context) SetSampleRate(sampleRate int) {
	ctxt.sample_rate = (C.int)(sampleRate)
}

func (ctxt *Context) SetChannelLayout(channelLayout uint64) {
	ctxt.channel_layout = (C.uint64_t)(channelLayout)
}

func (ctxt *Context) SetChannels(channels int) {
	ctxt.channels = (C.int)(channels)
}

func AvInvQ(q Rational) Rational {
	return (Rational)(C.av_inv_q((C.struct_AVRational)(q)))
}

func (ctxt *Context) SetBitRate(bitrate int64) {
	ctxt.bit_rate = C.long(bitrate)
}

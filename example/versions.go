package main

import (
	"log"

	"github.com/LeoKingLong/goav/avdevice"
	"github.com/LeoKingLong/goav/avfilter"
	"github.com/LeoKingLong/goav/avutil"
	"github.com/LeoKingLong/goav/swresample"
	"github.com/LeoKingLong/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}

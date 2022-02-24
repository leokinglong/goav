#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <libavformat/avformat.h>
#include <libavutil/avutil.h>

static int CheckInterrupt(void *ctx)
{
    int *stop  = (int*)ctx;
    
    //av_log(NULL, AV_LOG_WARNING,"stop:%d\n",*stop);
    return (*stop == 1) ? 1 : 0;
}

void avformat_set_callback(AVFormatContext *context, int *param)
{
    context->interrupt_callback.callback = CheckInterrupt;
    context->interrupt_callback.opaque = (void*)param;
}

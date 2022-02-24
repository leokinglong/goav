#ifndef AVFORMAT_CALLBACK_H
#define AVFORMAT_CALLBACK_H

struct AVFormatContext;
void avformat_set_callback(AVFormatContext *context, int *param);
#endif
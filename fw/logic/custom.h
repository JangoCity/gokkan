#include <stdint.h>

#ifndef __CUSTOM_H
#define __CUSTOM_H


void QueueToUsb(void const *argument);

void CommandQueueProcess(void const *argument);

void println(const char *str);

void receiveBytesFromCDC(uint8_t *Buf, uint32_t *Len);

void receiveBytesFromCAN(void);

void customInit(void);

#endif


#include "custom.h"
#include "stm32f1xx_hal.h"
#include "messages.pb-c.h"
#include "FreeRTOS.h"

#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <queue.h>
#include <cmsis_os.h>
#include <stm32f1xx_hal_can.h>
#include <usbd_cdc_if.h>

#define INCOMING_BUFFER_LENGTH 512


extern CAN_HandleTypeDef hcan1;
extern UART_HandleTypeDef huart3;

extern osMessageQId CAN_TO_SERIALHandle;
extern osMessageQId COMMAND_QUEUEHandle;


uint8_t serial_RX_Data[30];

uint8_t serialLocked = 0;
uint8_t canInited = 0;
uint8_t CR = '\n';

CanTxMsgTypeDef TxMessage;
CanRxMsgTypeDef RxMessage;


void println(const char *str) {
  HAL_UART_Transmit(&huart3, (uint8_t *) str, strlen(str), 10);
  HAL_UART_Transmit(&huart3, &CR, 1, 10);
}


void flushBuffers() {
  println("FlushBuffer");
  memset(&serial_RX_Data, 0x00, sizeof(serial_RX_Data));
}


void customInit() {

  CAN_FilterConfTypeDef f0FilterConfig;
  f0FilterConfig.FilterNumber = 0;
  f0FilterConfig.FilterMode = CAN_FILTERMODE_IDMASK;
  f0FilterConfig.FilterScale = CAN_FILTERSCALE_32BIT;
  f0FilterConfig.FilterIdHigh = 0x0000;
  f0FilterConfig.FilterIdLow = 0x0000;
  f0FilterConfig.FilterMaskIdHigh = 0x0000;
  f0FilterConfig.FilterMaskIdLow = 0x0000;
  f0FilterConfig.FilterFIFOAssignment = 0;
  f0FilterConfig.FilterActivation = ENABLE;
  f0FilterConfig.BankNumber = 13;

  CAN_FilterConfTypeDef f1FilterConfig;
  f1FilterConfig.FilterNumber = 0;
  f1FilterConfig.FilterMode = CAN_FILTERMODE_IDMASK;
  f1FilterConfig.FilterScale = CAN_FILTERSCALE_32BIT;
  f1FilterConfig.FilterIdHigh = 0x0000;
  f1FilterConfig.FilterIdLow = 0x0000;
  f1FilterConfig.FilterMaskIdHigh = 0x0000;
  f1FilterConfig.FilterMaskIdLow = 0x0000;
  f1FilterConfig.FilterFIFOAssignment = 1;
  f1FilterConfig.FilterActivation = ENABLE;
  f1FilterConfig.BankNumber = 14;

  HAL_CAN_ConfigFilter(&hcan1, &f0FilterConfig);
  HAL_CAN_ConfigFilter(&hcan1, &f1FilterConfig);


  hcan1.pTxMsg = &TxMessage;
  hcan1.pRxMsg = &RxMessage;


  flushBuffers();

  HAL_UART_Receive_IT(&huart3, serial_RX_Data, 30);
  HAL_CAN_Receive_IT(&hcan1, CAN_FIFO1);


}


void QueueToUsb(void const *argument) {
  FromDevice fromDevice;
  for (;;) {
    uint8_t *buf;                     // Buffer to store serialized data
    uint16_t len;
    if (xQueuePeek(CAN_TO_SERIALHandle, &fromDevice, 10)) {
      len = from_device__get_packed_size(&fromDevice);
      buf = malloc(len);
      from_device__pack(&fromDevice, buf);
      CDC_Transmit_FS(buf, len);
    }
    osDelay(1);
  }
}

void receiveBytesFromCAN() {
  FromDevice *fromDevice;
  uint8_t *data;
  data = hcan1.pRxMsg->Data;

  fromDevice->type = FROM_DEVICE__MESSAGE_TYPE__GET_FRAME;
  fromDevice->frame->id = hcan1.pRxMsg->StdId;
  fromDevice->frame->eid = hcan1.pRxMsg->ExtId;
  fromDevice->frame->ide = hcan1.pRxMsg->IDE;
  fromDevice->frame->rtr = hcan1.pRxMsg->RTR;
  fromDevice->frame->dlc = hcan1.pRxMsg->DLC;
  fromDevice->frame->data =
      (data[7] << 56) +
      (data[6] << 48) +
      (data[5] << 40) +
      (data[4] << 32) +
      (data[3] << 24) +
      (data[2] << 16) +
      (data[1] << 8) +
      data[0];
  xQueueSendFromISR(CAN_TO_SERIALHandle, &fromDevice, NULL);
}

void receiveBytesFromCDC(uint8_t *Buf, uint32_t *Len) {
  ToDevice *toDevice;
  toDevice = to_device__unpack(NULL, *Len, Buf);
  if (toDevice != NULL) {
    xQueueSendFromISR(COMMAND_QUEUEHandle, toDevice, NULL);
  }
}

void deinitCan() {
  if (canInited == 1) {
    canInited = 0;
    HAL_CAN_DeInit(&hcan1);
//    HAL_GPIO_WritePin(CAN_CTRL_GPIO_Port, CAN_CTRL_Pin, GPIO_PIN_RESET);
  }
}

void initCan() {
  if (canInited == 0) {
    canInited = 1;

//    HAL_GPIO_WritePin(CAN_CTRL_GPIO_Port, CAN_CTRL_Pin, GPIO_PIN_SET);
  }
}

void sendCanFrame(ToDevice *toDevice) {
  if (HAL_CAN_GetState(&hcan1) != HAL_CAN_STATE_BUSY_TX_RX ||
      HAL_CAN_GetState(&hcan1) != HAL_CAN_STATE_BUSY_TX) {
    hcan1.pTxMsg->StdId = toDevice->frame->id;
    hcan1.pTxMsg->ExtId = toDevice->frame->eid;
    hcan1.pTxMsg->IDE = toDevice->frame->ide;
    hcan1.pTxMsg->RTR = toDevice->frame->rtr;
    hcan1.pTxMsg->DLC = toDevice->frame->dlc;
    hcan1.pTxMsg->Data[0] = toDevice->frame->data;
    hcan1.pTxMsg->Data[1] = toDevice->frame->data >> 8;
    hcan1.pTxMsg->Data[2] = toDevice->frame->data >> 16;
    hcan1.pTxMsg->Data[3] = toDevice->frame->data >> 24;
    hcan1.pTxMsg->Data[4] = toDevice->frame->data >> 32;
    hcan1.pTxMsg->Data[5] = toDevice->frame->data >> 40;
    hcan1.pTxMsg->Data[6] = toDevice->frame->data >> 48;
    hcan1.pTxMsg->Data[7] = toDevice->frame->data >> 56;
    HAL_CAN_Transmit_IT(&hcan1);
  } else {
    println("Skipping can frame");
  }
}


void setBaudRateInRegisters(uint32_t Prescaler, uint32_t BS1, uint32_t BS2) {
  hcan1.Init.Prescaler = Prescaler;
  hcan1.Init.BS1 = BS1;
  hcan1.Init.BS2 = BS2;
}

void setBaudRate(ToDevice *toDevice) {
  BaudRate__Rate rate = toDevice->baudrate->rate;

  switch (rate) {
    case BAUD_RATE__RATE__K20:
      //pr = 100;bs1=15;bs2=2
      setBaudRateInRegisters(100, CAN_BS1_15TQ, CAN_BS2_2TQ);
    case BAUD_RATE__RATE__K50:
      //45;13;2
      setBaudRateInRegisters(45, CAN_BS1_13TQ, CAN_BS2_2TQ);
    case BAUD_RATE__RATE__K83:
      //24;15;2
      setBaudRateInRegisters(24, CAN_BS1_15TQ, CAN_BS2_2TQ);
    case BAUD_RATE__RATE__K100:
      //20;15;2
      setBaudRateInRegisters(20, CAN_BS1_15TQ, CAN_BS2_2TQ);
    case BAUD_RATE__RATE__K125:
      //18;13;2
      setBaudRateInRegisters(18, CAN_BS1_13TQ, CAN_BS2_2TQ);
    case BAUD_RATE__RATE__K250:
      //9;13;2
      setBaudRateInRegisters(9, CAN_BS1_13TQ, CAN_BS2_2TQ);
    case BAUD_RATE__RATE__K500:
      //4;15;2
      setBaudRateInRegisters(4, CAN_BS1_15TQ, CAN_BS2_2TQ);
    case BAUD_RATE__RATE__K800:
      //3;12;2
      setBaudRateInRegisters(3, CAN_BS1_12TQ, CAN_BS2_2TQ);
    case BAUD_RATE__RATE__K1000:
      //2;15;2
      setBaudRateInRegisters(2, CAN_BS1_15TQ, CAN_BS2_2TQ);
  }

}


void CommandQueueProcess(void const *argument) {
  ToDevice toDevice;
  for (;;) {
    if (xQueuePeek(COMMAND_QUEUEHandle, &toDevice, 10)) {
      if (toDevice.has_type) {
        if (toDevice.type == TO_DEVICE__MESSAGE_TYPE__SEND_FRAME) {
          if (canInited == 1) {
            sendCanFrame(&toDevice);
          }
          println("Sending can frame");
        } else if (toDevice.type == TO_DEVICE__MESSAGE_TYPE__SET_BAUDRATE) {
          if (canInited == 0) {
            setBaudRate(&toDevice);
          }
          println("Setting baudrate");
        } else if (toDevice.type == TO_DEVICE__MESSAGE_TYPE__INIT) {
          initCan();
          println("Initting can");
        } else if (toDevice.type == TO_DEVICE__MESSAGE_TYPE__DEINIT) {
          deinitCan();
          println("Deinitting can");
        } else if (toDevice.type == TO_DEVICE__MESSAGE_TYPE__GET_STATUS) {
          println("getting status");
        } else if (toDevice.type == TO_DEVICE__MESSAGE_TYPE__SET_FILTER) {
          println("Setting filter");
        }
      }
    }
    osDelay(1);
  }
}

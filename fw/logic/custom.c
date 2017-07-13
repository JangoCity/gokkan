
#include "custom.h"
#include "stm32f1xx_hal.h"


#include <stdlib.h>
#include <string.h>


extern CAN_HandleTypeDef hcan1;
extern UART_HandleTypeDef huart3;


uint8_t serial_RX_Data[30];

uint8_t canLocked = 0;
uint8_t serialLocked = 0;


static CanTxMsgTypeDef TxMessage;
static CanRxMsgTypeDef RxMessage;


void flushBuffers() {
  memset(&serial_RX_Data, 0x00, sizeof(serial_RX_Data));
}


void init() {

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


void QueryToCAN(void const *argument) {

  /* USER CODE BEGIN 5 */
//    CANMessage canMsg;
//
//
//    /* Infinite loop */
//    for (;;) {
//
//        if (canLocked == 0 && xQueuePeek(SERIAL_TO_CANHandle, &canMsg, 10)) {
//            uint8_t* msg = (uint8_t *) hcan.pTxMsg;
//            uint8_t* data = canMsg.dataPointer;
//            for (int i = 0; i < sizeof(CanTxMsgTypeDef); i++) {
//              msg[i] = data[i];
//            }
//
//            HAL_StatusTypeDef r = HAL_CAN_Transmit_IT(&hcan);
//            if (r == HAL_OK) {
//                canLocked = 1;
//            }
//        }
//
//        osDelay(1);
//    }
//  /* USER CODE END 5 */
}

void QueryToSerial(void const *argument) {
//  /* USER CODE BEGIN QToSerialF */
//  SMessage sMessage;
//    for (;;) {
//        if (serialLocked == 0 && xQueuePeek(CAN_TO_SERIALHandle, &sMessage, 10)) {
//            pt_serialMessage = malloc(rxMessageSize + 2);
//            memcpy(pt_serialMessage + 1, sMessage.dataPointer, rxMessageSize);
//            pt_serialMessage[rxMessageSize + 1] = '\n';
//            pt_serialMessage[0] = '!';
//            free(sMessage.dataPointer);
//            HAL_StatusTypeDef r = HAL_UART_Transmit_IT(&huart3, pt_serialMessage, rxMessageSize + 2);
//            if (r == HAL_OK) {
//                serialLocked = 1;
//            }
//        }
//      osDelay(1);
//    }
}



#include "custom.h"
#include "stm32f1xx_hal.h"
#include "cmsis_os.h"

#include <string.h>
#include <stdlib.h>

CAN_HandleTypeDef hcan;
UART_HandleTypeDef huart3;


osThreadId QToCANHandle;
osThreadId QToSerialHandle;
osMessageQId CAN_TO_SERIALHandle;
osMessageQId SERIAL_TO_CANHandle;


uint8_t serial_RX_Data[30];

uint8_t* pt_serialMessage;
uint8_t* pt_CANMessage;

uint8_t canLocked = 0;
uint8_t serialLocked = 0;

//skip fmi and fifo number;
uint8_t rxMessageSize = sizeof(CanRxMsgTypeDef) - 8;

static CanTxMsgTypeDef TxMessage;
static CanRxMsgTypeDef RxMessage;



void SystemClock_Config(void);
static void MX_GPIO_Init(void);
static void MX_USART3_UART_Init(void);
static void MX_CAN_Init(void);
static void MX_CRC_Init(void);
void QToCANF(void const * argument);
void QToSerialF(void const * argument);
static void MX_NVIC_Init(void);


void flushBuffers() {
    memset(&serial_RX_Data, 0x00, sizeof(serial_RX_Data));
}


int main(void)
{

  HAL_Init();

  SystemClock_Config();

  MX_GPIO_Init();
  MX_USART3_UART_Init();
  MX_CAN_Init();
  MX_CRC_Init();

  MX_NVIC_Init();


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

    HAL_CAN_ConfigFilter(&hcan, &f0FilterConfig);
    HAL_CAN_ConfigFilter(&hcan, &f1FilterConfig);


    hcan.pTxMsg = &TxMessage;
    hcan.pRxMsg = &RxMessage;


    flushBuffers();

    HAL_UART_Receive_IT(&huart3, serial_RX_Data, 30);
    HAL_CAN_Receive_IT(&hcan, CAN_FIFO1);

  osThreadDef(QToCAN, QToCANF, osPriorityNormal, 0, 128);
  QToCANHandle = osThreadCreate(osThread(QToCAN), NULL);

  osThreadDef(QToSerial, QToSerialF, osPriorityNormal, 0, 128);
  QToSerialHandle = osThreadCreate(osThread(QToSerial), NULL);

  osMessageQDef(CAN_TO_SERIAL, 16, CANMessage);
  CAN_TO_SERIALHandle = osMessageCreate(osMessageQ(CAN_TO_SERIAL), NULL);

  osMessageQDef(SERIAL_TO_CAN, 16, SMessage);
  SERIAL_TO_CANHandle = osMessageCreate(osMessageQ(SERIAL_TO_CAN), NULL);

  osKernelStart();


    while (1) {


    }

}



void QToCANF(void const * argument)
{

  /* USER CODE BEGIN 5 */
    CANMessage canMsg;


    /* Infinite loop */
    for (;;) {

        if (canLocked == 0 && xQueuePeek(SERIAL_TO_CANHandle, &canMsg, 10)) {
            uint8_t* msg = (uint8_t *) hcan.pTxMsg;
            uint8_t* data = canMsg.dataPointer;
            for (int i = 0; i < sizeof(CanTxMsgTypeDef); i++) {
              msg[i] = data[i];
            }

            HAL_StatusTypeDef r = HAL_CAN_Transmit_IT(&hcan);
            if (r == HAL_OK) {
                canLocked = 1;
            }
        }

        osDelay(1);
    }
  /* USER CODE END 5 */
}

/* QToSerialF function */
void QToSerialF(void const * argument)
{
  /* USER CODE BEGIN QToSerialF */
  SMessage sMessage;
    for (;;) {
        if (serialLocked == 0 && xQueuePeek(CAN_TO_SERIALHandle, &sMessage, 10)) {
            pt_serialMessage = malloc(rxMessageSize + 2);
            memcpy(pt_serialMessage + 1, sMessage.dataPointer, rxMessageSize);
            pt_serialMessage[rxMessageSize + 1] = '\n';
            pt_serialMessage[0] = '!';
            free(sMessage.dataPointer);
            HAL_StatusTypeDef r = HAL_UART_Transmit_IT(&huart3, pt_serialMessage, rxMessageSize + 2);
            if (r == HAL_OK) {
                serialLocked = 1;
            }
        }
      osDelay(1);
    }
}


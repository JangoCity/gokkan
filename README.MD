GOKKAN(極寒)
=======================

Gokkan - just a japanese word for mid winter.

----------


WTF
---

Some dirty and messy code to interact from go-lang to CAN bus.

 - fw - Some CubeMX generated sources for Keil. Configured for stm32f103rb board (cheap shit, no CAN and USB simultaneously) to proxy CANFrames to USART and back. FreeRTOS, STM32 and so other shit code has it own licenses and copyrights, Blah-blah-blah. Probably you could regenerate it for some free toolchain with CubeMX
 - protocol - implementation of my howegrown protocol callen PotatoBus (just because)
 - data - right now simply CANFrame marshaller\unmarshaller
 - driver - serila connectivity utils
_______________________

Not implemented
---------------
 - CAN bus error handling
 - USART bus error handling
 - FreeRTOS error handling
 - General errors and strange IRQs handling
 - CANFilters
 - ~~Serial through coroutines~~
 - Move to nice and more suitable hardware
 - Some UDS (ISO 14229-1:2013) routines
 - Sniff, dump and simple DoS examples 

//go:build forge

package machine

const (
	T0        = PA0  // Thermistor 0
	T1        = PA1  // Thermistor 1
	T2        = PA2  // Thermistor 2
	HOTEND    = PA3  // Hotend
	BED       = PA4  // Printer bed
	FAN3      = PA5  // Fan 3
	FAN2      = PA6  // Fan 2
	FAN1      = PA7  // Fan 1
	TX        = PA9  // UART Tx
	RX        = PA10 // UART Rx
	DM        = PA11 // USB D-
	DP        = PA12 // USB D+
	SWDIO_DBG = PA13 // Debug SWDIO
	SWCLK_DBG = PA14 // Debug SWCLK

	EN_M2   = PB0 // Enable Motor 2
	DIAG_M2 = PB1 // Diagnose Motor 2
	EN_M4   = PB7 // Enable Motor 4
	DIAG_M4 = PB8 // Diagnose Motor 4
	STEP_M4 = PB9 // Step Motor 4

	EN_M1       = PC0  // Enable Motor 1
	STEP_M1     = PC1  // Step Motor 1
	DIR_M1      = PC2  // Direction Motor 1
	DIAG_M1     = PC3  // Diagnose Motor 1
	DIR_M2      = PC4  // Direction Motor 2
	STEP_M2     = PC5  // Step Motor 2
	EN_M3       = PC10 // Enable Motor 3
	DIAG_M3     = PC11 // Diagnose Motor 3
	STEP_M3     = PC12 // Step Motor 3
	NEOPIXEL_IN = PC15 // Neopixel In

	SCRN_SDA = PD0 // Screen Serial Data In
	SCRN_SCL = PD1 // Screen Clock
	EN_BTN   = PD2 // Encoder Button
	EN_REV   = PD3 // Encoder Reverse
	EN_FOR   = PD4 // Encoder Forward
	DIR_M3   = PD5 // Direction Motor 3

	DIR_M4 = PE0 // Direction Motor 4
	Z_END  = PE1 // Bed Leveling
	ABL    = PE2 // Bed Leveling
)

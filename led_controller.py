import board
import neopixel
import led

DATA_PIN = board.D18
NUMBER_OF_LEDS = 102

class LedContraller():
	led = None

	def __init__(self):
		self.led = led.LedStripe(DATA_PIN, NUMBER_OF_LEDS)

	def Fill_white(self):
		self.led.Fill(led.COLOR_WHITE, is_slowly=True)

	def Clear(self):
		self.led.Clear(is_slowly=True)

import time

import board
import neopixel

DATA_PIN = board.D18
NUMBER_OF_LEDS = 102
ORDER = neopixel.GRB

BRIGHTNESS = 0.5

NO_COLOR = (0, 0, 0)
COLOR_WHITE = (255, 255, 255)

NUMBER_OF_SLOWLY_STEPS = 255
SLOWLY_STEPS_TIME = 1.5

class LedStripe():
	controller = None
	pixels_state = []
	number_of_leds = 0

	def __init__(self, data_pin, number_of_leds):
		self.controller = neopixel.NeoPixel(data_pin, number_of_leds, pixel_order=ORDER, auto_write=False, brightness=BRIGHTNESS)
		self.number_of_leds = number_of_leds
		self.pixels_state = [NO_COLOR] * self.number_of_leds

		self.controller.show()

	def __del__(self):
		self.controller.deinit()

	def Fill(self, color, is_slowly):
		new_pixels_state = [color] * self.number_of_leds

		self._show(new_pixels_state, is_slowly)

	def Clear(self, is_slowly):
		new_pixels_state = [NO_COLOR] * self.number_of_leds

		self._show(new_pixels_state, is_slowly)

	def _show(self, new_pixels_state, is_slowly=False):
		if is_slowly is True:
			for slow_step_number in range(NUMBER_OF_SLOWLY_STEPS):

				for i, pixel_state in enumerate(self.pixels_state):
					deltaR = int(abs(self.pixels_state[i][0] - new_pixels_state[i][0]) / NUMBER_OF_SLOWLY_STEPS)
					deltaG = int(abs(self.pixels_state[i][1] - new_pixels_state[i][1]) / NUMBER_OF_SLOWLY_STEPS)
					deltaB = int(abs(self.pixels_state[i][2] - new_pixels_state[i][2]) / NUMBER_OF_SLOWLY_STEPS)

					R = pixel_state[0] + deltaR * (slow_step_number + 1) if new_pixels_state[i][0] > pixel_state[0] else pixel_state[0] - deltaR * (slow_step_number + 1)
					G = pixel_state[1] + deltaG * (slow_step_number + 1) if new_pixels_state[i][1] > pixel_state[1] else pixel_state[1] - deltaG * (slow_step_number + 1)
					B = pixel_state[2] + deltaB * (slow_step_number + 1) if new_pixels_state[i][2] > pixel_state[2] else pixel_state[2] - deltaB * (slow_step_number + 1)

					self.controller[i] = (R, G, B)

				self.controller.show()
				time.sleep(SLOWLY_STEPS_TIME / NUMBER_OF_SLOWLY_STEPS)

			self.pixels_state = new_pixels_state
			return

		for i, pixel in enumerate(self.pixels_state):
			self.controller[i] = pixel

		self.pixels_state = new_pixels_state
		self.controller.show()


	def _slow_change(self):
		pass

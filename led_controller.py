import led


class LedContraller():
	_led = None
	_isSwichedOn = False

	def __init__(self):
		self._led = led.LedStripe()

	def Switch(self):
		if self._isSwichedOn is True:
			self.Clear()
		else:
			self.Fill_white()

	def Fill_white(self, is_slowly=False):
		self._led.Fill(led.COLOR_WHITE, is_slowly)
		self._isSwichedOn = not self._isSwichedOn

	def Clear(self, is_slowly=False):
		self._led.Clear(is_slowly)
		self._isSwichedOn = not self._isSwichedOn

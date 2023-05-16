import led


class LedContraller():
	_led = None
	_isSwichedOn = None

	def __init__(self):
		self._led = led.LedStripe()
		self.Clear()

	def Switch(self):
		if self._isSwichedOn:
			self.Fill_white()
		else:
			self.Clear()

	def Fill_white(self, is_slowly=False):
		self._led.Fill(led.COLOR_WHITE, is_slowly)
		self._isSwichedOn = True

	def Clear(self, is_slowly=False):
		self._led.Clear(is_slowly)
		self._isSwichedOn = False

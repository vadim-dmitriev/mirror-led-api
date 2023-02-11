from concurrent import futures
import logging

import grpc
from grpc_reflection.v1alpha import reflection

import led_api_pb2
import led_api_pb2_grpc
from led_controller import LedContraller

port = '8090'


class LedService(led_api_pb2_grpc.LedService):
	led = None

	def __init__(self) -> None:
		super().__init__()
		self.led = LedContraller()

	def LightLED(self, request, context):
		logging.warning("LightLED calls")

		if request.mode == led_api_pb2.LightLEDRequest.NO_MODE:
			logging.warning("NO_MODE")
			self.led.Clear()

		if request.mode == led_api_pb2.LightLEDRequest.KEY_WORD:
			logging.warning("KEY_WORD")
			self.led.Fill_white()

		return led_api_pb2.LightLEDResponse(success=True)


def serve():
	server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
	led_api_pb2_grpc.add_LedServiceServicer_to_server(LedService(), server)
	SERVICE_NAMES = (
		led_api_pb2.DESCRIPTOR.services_by_name['LedService'].full_name,
		reflection.SERVICE_NAME,
	)

	reflection.enable_server_reflection(SERVICE_NAMES, server)

	server.add_insecure_port('0.0.0.0:' + port)
	server.start()

	print("Server started, listening on " + port)

	server.wait_for_termination()


if __name__ == '__main__':
	logging.basicConfig()
	serve()

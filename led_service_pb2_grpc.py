# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import led_service_pb2 as led__service__pb2


class LedServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.LightLED = channel.unary_unary(
                '/led_service.LedService/LightLED',
                request_serializer=led__service__pb2.LightLEDRequest.SerializeToString,
                response_deserializer=led__service__pb2.LightLEDResponse.FromString,
                )
        self.SwitchLED = channel.unary_unary(
                '/led_service.LedService/SwitchLED',
                request_serializer=led__service__pb2.Empty.SerializeToString,
                response_deserializer=led__service__pb2.Empty.FromString,
                )


class LedServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def LightLED(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SwitchLED(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_LedServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'LightLED': grpc.unary_unary_rpc_method_handler(
                    servicer.LightLED,
                    request_deserializer=led__service__pb2.LightLEDRequest.FromString,
                    response_serializer=led__service__pb2.LightLEDResponse.SerializeToString,
            ),
            'SwitchLED': grpc.unary_unary_rpc_method_handler(
                    servicer.SwitchLED,
                    request_deserializer=led__service__pb2.Empty.FromString,
                    response_serializer=led__service__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'led_service.LedService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class LedService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def LightLED(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/led_service.LedService/LightLED',
            led__service__pb2.LightLEDRequest.SerializeToString,
            led__service__pb2.LightLEDResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SwitchLED(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/led_service.LedService/SwitchLED',
            led__service__pb2.Empty.SerializeToString,
            led__service__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

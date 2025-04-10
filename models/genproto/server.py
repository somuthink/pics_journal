from models.model_chat import mistral_chat
import com_pb2_grpc, com_pb2
import grpc
from concurrent import futures


class Communication(com_pb2_grpc.CommunicationServicer):
    def GetAReply(self, request, context):
        return com_pb2.ModuleReply(answer=mistral_chat(request))


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    com_pb2_grpc.add_Communication_to_server(Communication(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()

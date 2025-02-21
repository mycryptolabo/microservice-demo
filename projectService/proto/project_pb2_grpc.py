# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . 0


class ProjectSvcStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.createProject = channel.unary_unary(
                '/demo_project.ProjectSvc/createProject',
                request_serializer=project__pb2.CreateProjectRequest.SerializeToString,
                response_deserializer=project__pb2.ProjectResponse.FromString,
                )
        self.createTag = channel.unary_unary(
                '/demo_project.ProjectSvc/createTag',
                request_serializer=project__pb2.CreateTagRequest.SerializeToString,
                response_deserializer=project__pb2.TagResponse.FromString,
                )
        self.getProject = channel.unary_unary(
                '/demo_project.ProjectSvc/getProject',
                request_serializer=project__pb2.GetProjectRequest.SerializeToString,
                response_deserializer=project__pb2.ProjectResponse.FromString,
                )


class ProjectSvcServicer(object):
    """Missing associated documentation comment in .proto file."""

    def createProject(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def createTag(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def getProject(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProjectSvcServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'createProject': grpc.unary_unary_rpc_method_handler(
                    servicer.createProject,
                    request_deserializer=project__pb2.CreateProjectRequest.FromString,
                    response_serializer=project__pb2.ProjectResponse.SerializeToString,
            ),
            'createTag': grpc.unary_unary_rpc_method_handler(
                    servicer.createTag,
                    request_deserializer=project__pb2.CreateTagRequest.FromString,
                    response_serializer=project__pb2.TagResponse.SerializeToString,
            ),
            'getProject': grpc.unary_unary_rpc_method_handler(
                    servicer.getProject,
                    request_deserializer=project__pb2.GetProjectRequest.FromString,
                    response_serializer=project__pb2.ProjectResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'demo_project.ProjectSvc', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ProjectSvc(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def createProject(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/demo_project.ProjectSvc/createProject',
            project__pb2.CreateProjectRequest.SerializeToString,
            project__pb2.ProjectResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def createTag(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/demo_project.ProjectSvc/createTag',
            project__pb2.CreateTagRequest.SerializeToString,
            project__pb2.TagResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def getProject(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/demo_project.ProjectSvc/getProject',
            project__pb2.GetProjectRequest.SerializeToString,
            project__pb2.ProjectResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: user/user.proto for package 'demo_user'

require 'grpc'
require 'user/user_pb'

module DemoUser
  module UserSvc
    class Service

      include ::GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'demo_user.UserSvc'

      rpc :register, ::DemoUser::RegisterRequest, ::DemoUser::UserResponse
      rpc :login, ::DemoUser::LoginRequest, ::DemoUser::UserResponse
      rpc :verify, ::DemoUser::VerifyRequest, ::DemoUser::VerifyResponse
      rpc :getUser, ::DemoUser::GetUserRequest, ::DemoUser::VerifyResponse
    end

    Stub = Service.rpc_stub_class
  end
end

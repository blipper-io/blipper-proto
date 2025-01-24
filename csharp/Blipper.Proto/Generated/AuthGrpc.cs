// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: auth/auth.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Blipper.Proto.AuthService {
  public static partial class AuthService
  {
    static readonly string __ServiceName = "blipper.auth.v1.AuthService";

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static void __Helper_SerializeMessage(global::Google.Protobuf.IMessage message, grpc::SerializationContext context)
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (message is global::Google.Protobuf.IBufferMessage)
      {
        context.SetPayloadLength(message.CalculateSize());
        global::Google.Protobuf.MessageExtensions.WriteTo(message, context.GetBufferWriter());
        context.Complete();
        return;
      }
      #endif
      context.Complete(global::Google.Protobuf.MessageExtensions.ToByteArray(message));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static class __Helper_MessageCache<T>
    {
      public static readonly bool IsBufferMessage = global::System.Reflection.IntrospectionExtensions.GetTypeInfo(typeof(global::Google.Protobuf.IBufferMessage)).IsAssignableFrom(typeof(T));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static T __Helper_DeserializeMessage<T>(grpc::DeserializationContext context, global::Google.Protobuf.MessageParser<T> parser) where T : global::Google.Protobuf.IMessage<T>
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (__Helper_MessageCache<T>.IsBufferMessage)
      {
        return parser.ParseFrom(context.PayloadAsReadOnlySequence());
      }
      #endif
      return parser.ParseFrom(context.PayloadAsNewBuffer());
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.RegisterRequest> __Marshaller_blipper_auth_v1_RegisterRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.RegisterRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.RegisterResponse> __Marshaller_blipper_auth_v1_RegisterResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.RegisterResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.LoginRequest> __Marshaller_blipper_auth_v1_LoginRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.LoginRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.LoginResponse> __Marshaller_blipper_auth_v1_LoginResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.LoginResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.LogoutRequest> __Marshaller_blipper_auth_v1_LogoutRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.LogoutRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Google.Protobuf.WellKnownTypes.Empty> __Marshaller_google_protobuf_Empty = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Google.Protobuf.WellKnownTypes.Empty.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.RefreshTokenRequest> __Marshaller_blipper_auth_v1_RefreshTokenRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.RefreshTokenRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.RefreshTokenResponse> __Marshaller_blipper_auth_v1_RefreshTokenResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.RefreshTokenResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.ValidateTokenRequest> __Marshaller_blipper_auth_v1_ValidateTokenRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.ValidateTokenRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Blipper.Proto.AuthService.ValidateTokenResponse> __Marshaller_blipper_auth_v1_ValidateTokenResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Blipper.Proto.AuthService.ValidateTokenResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Blipper.Proto.AuthService.RegisterRequest, global::Blipper.Proto.AuthService.RegisterResponse> __Method_Register = new grpc::Method<global::Blipper.Proto.AuthService.RegisterRequest, global::Blipper.Proto.AuthService.RegisterResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Register",
        __Marshaller_blipper_auth_v1_RegisterRequest,
        __Marshaller_blipper_auth_v1_RegisterResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Blipper.Proto.AuthService.LoginRequest, global::Blipper.Proto.AuthService.LoginResponse> __Method_Login = new grpc::Method<global::Blipper.Proto.AuthService.LoginRequest, global::Blipper.Proto.AuthService.LoginResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Login",
        __Marshaller_blipper_auth_v1_LoginRequest,
        __Marshaller_blipper_auth_v1_LoginResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Blipper.Proto.AuthService.LogoutRequest, global::Google.Protobuf.WellKnownTypes.Empty> __Method_Logout = new grpc::Method<global::Blipper.Proto.AuthService.LogoutRequest, global::Google.Protobuf.WellKnownTypes.Empty>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Logout",
        __Marshaller_blipper_auth_v1_LogoutRequest,
        __Marshaller_google_protobuf_Empty);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Blipper.Proto.AuthService.RefreshTokenRequest, global::Blipper.Proto.AuthService.RefreshTokenResponse> __Method_RefreshToken = new grpc::Method<global::Blipper.Proto.AuthService.RefreshTokenRequest, global::Blipper.Proto.AuthService.RefreshTokenResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "RefreshToken",
        __Marshaller_blipper_auth_v1_RefreshTokenRequest,
        __Marshaller_blipper_auth_v1_RefreshTokenResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Blipper.Proto.AuthService.ValidateTokenRequest, global::Blipper.Proto.AuthService.ValidateTokenResponse> __Method_ValidateToken = new grpc::Method<global::Blipper.Proto.AuthService.ValidateTokenRequest, global::Blipper.Proto.AuthService.ValidateTokenResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "ValidateToken",
        __Marshaller_blipper_auth_v1_ValidateTokenRequest,
        __Marshaller_blipper_auth_v1_ValidateTokenResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Blipper.Proto.AuthService.AuthReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of AuthService</summary>
    [grpc::BindServiceMethod(typeof(AuthService), "BindService")]
    public abstract partial class AuthServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Blipper.Proto.AuthService.RegisterResponse> Register(global::Blipper.Proto.AuthService.RegisterRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Blipper.Proto.AuthService.LoginResponse> Login(global::Blipper.Proto.AuthService.LoginRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Google.Protobuf.WellKnownTypes.Empty> Logout(global::Blipper.Proto.AuthService.LogoutRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Blipper.Proto.AuthService.RefreshTokenResponse> RefreshToken(global::Blipper.Proto.AuthService.RefreshTokenRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Blipper.Proto.AuthService.ValidateTokenResponse> ValidateToken(global::Blipper.Proto.AuthService.ValidateTokenRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for AuthService</summary>
    public partial class AuthServiceClient : grpc::ClientBase<AuthServiceClient>
    {
      /// <summary>Creates a new client for AuthService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public AuthServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for AuthService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public AuthServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected AuthServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected AuthServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Blipper.Proto.AuthService.RegisterResponse Register(global::Blipper.Proto.AuthService.RegisterRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Register(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Blipper.Proto.AuthService.RegisterResponse Register(global::Blipper.Proto.AuthService.RegisterRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Register, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Blipper.Proto.AuthService.RegisterResponse> RegisterAsync(global::Blipper.Proto.AuthService.RegisterRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RegisterAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Blipper.Proto.AuthService.RegisterResponse> RegisterAsync(global::Blipper.Proto.AuthService.RegisterRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Register, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Blipper.Proto.AuthService.LoginResponse Login(global::Blipper.Proto.AuthService.LoginRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Login(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Blipper.Proto.AuthService.LoginResponse Login(global::Blipper.Proto.AuthService.LoginRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Login, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Blipper.Proto.AuthService.LoginResponse> LoginAsync(global::Blipper.Proto.AuthService.LoginRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return LoginAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Blipper.Proto.AuthService.LoginResponse> LoginAsync(global::Blipper.Proto.AuthService.LoginRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Login, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Google.Protobuf.WellKnownTypes.Empty Logout(global::Blipper.Proto.AuthService.LogoutRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Logout(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Google.Protobuf.WellKnownTypes.Empty Logout(global::Blipper.Proto.AuthService.LogoutRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Logout, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Google.Protobuf.WellKnownTypes.Empty> LogoutAsync(global::Blipper.Proto.AuthService.LogoutRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return LogoutAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Google.Protobuf.WellKnownTypes.Empty> LogoutAsync(global::Blipper.Proto.AuthService.LogoutRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Logout, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Blipper.Proto.AuthService.RefreshTokenResponse RefreshToken(global::Blipper.Proto.AuthService.RefreshTokenRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RefreshToken(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Blipper.Proto.AuthService.RefreshTokenResponse RefreshToken(global::Blipper.Proto.AuthService.RefreshTokenRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_RefreshToken, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Blipper.Proto.AuthService.RefreshTokenResponse> RefreshTokenAsync(global::Blipper.Proto.AuthService.RefreshTokenRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RefreshTokenAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Blipper.Proto.AuthService.RefreshTokenResponse> RefreshTokenAsync(global::Blipper.Proto.AuthService.RefreshTokenRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_RefreshToken, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Blipper.Proto.AuthService.ValidateTokenResponse ValidateToken(global::Blipper.Proto.AuthService.ValidateTokenRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return ValidateToken(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Blipper.Proto.AuthService.ValidateTokenResponse ValidateToken(global::Blipper.Proto.AuthService.ValidateTokenRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_ValidateToken, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Blipper.Proto.AuthService.ValidateTokenResponse> ValidateTokenAsync(global::Blipper.Proto.AuthService.ValidateTokenRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return ValidateTokenAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Blipper.Proto.AuthService.ValidateTokenResponse> ValidateTokenAsync(global::Blipper.Proto.AuthService.ValidateTokenRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_ValidateToken, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override AuthServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new AuthServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(AuthServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_Register, serviceImpl.Register)
          .AddMethod(__Method_Login, serviceImpl.Login)
          .AddMethod(__Method_Logout, serviceImpl.Logout)
          .AddMethod(__Method_RefreshToken, serviceImpl.RefreshToken)
          .AddMethod(__Method_ValidateToken, serviceImpl.ValidateToken).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, AuthServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_Register, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Blipper.Proto.AuthService.RegisterRequest, global::Blipper.Proto.AuthService.RegisterResponse>(serviceImpl.Register));
      serviceBinder.AddMethod(__Method_Login, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Blipper.Proto.AuthService.LoginRequest, global::Blipper.Proto.AuthService.LoginResponse>(serviceImpl.Login));
      serviceBinder.AddMethod(__Method_Logout, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Blipper.Proto.AuthService.LogoutRequest, global::Google.Protobuf.WellKnownTypes.Empty>(serviceImpl.Logout));
      serviceBinder.AddMethod(__Method_RefreshToken, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Blipper.Proto.AuthService.RefreshTokenRequest, global::Blipper.Proto.AuthService.RefreshTokenResponse>(serviceImpl.RefreshToken));
      serviceBinder.AddMethod(__Method_ValidateToken, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Blipper.Proto.AuthService.ValidateTokenRequest, global::Blipper.Proto.AuthService.ValidateTokenResponse>(serviceImpl.ValidateToken));
    }

  }
}
#endregion

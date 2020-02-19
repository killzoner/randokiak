/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';


/**
 *  The request message containing the user's name.
 */
export interface HelloRequest {
  name: string;
}

/**
 *  The response message containing the greetings
 */
export interface HelloReply {
  message: string;
}

const baseHelloRequest: object = {
  name: "",
};

const baseHelloReply: object = {
  message: "",
};

/**
 *  The greeting service definition.
 */
export interface Greeter {

  /**
   *  Sends a greeting
   */
  SayHello(request: HelloRequest): Promise<HelloReply>;

}

export class GreeterClientImpl implements Greeter {

  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }

  SayHello(request: HelloRequest): Promise<HelloReply> {
    const data = HelloRequest.encode(request).finish();
    const promise = this.rpc.request("protocols.Greeter", "SayHello", data);
    return promise.then(data => HelloReply.decode(new Reader(data)));
  }

}

interface Rpc {

  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;

}

export const HelloRequest = {
  encode(message: HelloRequest, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.name);
    return writer;
  },
  decode(reader: Reader, length?: number): HelloRequest {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseHelloRequest) as HelloRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): HelloRequest {
    const message = Object.create(baseHelloRequest) as HelloRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<HelloRequest>): HelloRequest {
    const message = Object.create(baseHelloRequest) as HelloRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
  toJSON(message: HelloRequest): unknown {
    const obj: any = {};
    obj.name = message.name || "";
    return obj;
  },
};

export const HelloReply = {
  encode(message: HelloReply, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.message);
    return writer;
  },
  decode(reader: Reader, length?: number): HelloReply {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseHelloReply) as HelloReply;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): HelloReply {
    const message = Object.create(baseHelloReply) as HelloReply;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<HelloReply>): HelloReply {
    const message = Object.create(baseHelloReply) as HelloReply;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
  toJSON(message: HelloReply): unknown {
    const obj: any = {};
    obj.message = message.message || "";
    return obj;
  },
};

type DeepPartial<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T[P] extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T[P] extends Date | Function | Uint8Array | undefined
  ? T[P]
  : T[P] extends infer U | undefined
  ? DeepPartial<U>
  : T[P] extends object
  ? DeepPartial<T[P]>
  : T[P]
};
/* eslint-disable */
import { Profile } from './randomprofile';
import { Reader, Writer } from 'protobufjs/minimal';


/**
 *  The request message containing the user's name.
 */
export interface ProfileRequest {
}

/**
 *  The response message containing the greetings
 */
export interface ProfileReply {
  profile: Profile | undefined;
}

const baseProfileRequest: object = {
};

const baseProfileReply: object = {
  profile: undefined,
};

export interface Randomizer {

  GetProfile(request: ProfileRequest): Promise<ProfileReply>;

}

export class RandomizerClientImpl implements Randomizer {

  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }

  GetProfile(request: ProfileRequest): Promise<ProfileReply> {
    const data = ProfileRequest.encode(request).finish();
    const promise = this.rpc.request("protocols.Randomizer", "GetProfile", data);
    return promise.then(data => ProfileReply.decode(new Reader(data)));
  }

}

interface Rpc {

  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;

}

export const ProfileRequest = {
  encode(message: ProfileRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },
  decode(reader: Reader, length?: number): ProfileRequest {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseProfileRequest) as ProfileRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): ProfileRequest {
    const message = Object.create(baseProfileRequest) as ProfileRequest;
    return message;
  },
  fromPartial(object: DeepPartial<ProfileRequest>): ProfileRequest {
    const message = Object.create(baseProfileRequest) as ProfileRequest;
    return message;
  },
  toJSON(message: ProfileRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const ProfileReply = {
  encode(message: ProfileReply, writer: Writer = Writer.create()): Writer {
    if (message.profile !== undefined && message.profile !== undefined) {
      Profile.encode(message.profile, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },
  decode(reader: Reader, length?: number): ProfileReply {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseProfileReply) as ProfileReply;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.profile = Profile.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): ProfileReply {
    const message = Object.create(baseProfileReply) as ProfileReply;
    if (object.profile !== undefined && object.profile !== null) {
      message.profile = Profile.fromJSON(object.profile);
    } else {
      message.profile = undefined;
    }
    return message;
  },
  fromPartial(object: DeepPartial<ProfileReply>): ProfileReply {
    const message = Object.create(baseProfileReply) as ProfileReply;
    if (object.profile !== undefined && object.profile !== null) {
      message.profile = Profile.fromPartial(object.profile);
    } else {
      message.profile = undefined;
    }
    return message;
  },
  toJSON(message: ProfileReply): unknown {
    const obj: any = {};
    obj.profile = message.profile ? Profile.toJSON(message.profile) : undefined;
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
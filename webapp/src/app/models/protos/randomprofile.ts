/* eslint-disable */
import * as Long from 'long';
import { Writer, Reader } from 'protobufjs/minimal';


export interface ID {
  name: string;
}

export interface Location {
  street: string;
  city: string;
  state: string;
  postcode: number;
}

export interface Login {
  username: string;
  password: string;
  salt: string;
  md5: string;
  sha1: string;
  sha256: string;
}

export interface Name {
  first: string;
  last: string;
  title: string;
}

export interface Picture {
  large: string;
  medium: string;
  thumbnail: string;
}

export interface Profile {
  gender: string;
  name: ProfileName | undefined;
  location: ProfileLocation | undefined;
  email: string;
  login: ProfileLogin | undefined;
  dob: string;
  registered: string;
  phone: string;
  cell: string;
  id: ProfileId | undefined;
  picture: ProfilePicture | undefined;
  nat: string;
}

export interface ProfileId {
  name: string;
}

export interface ProfileLocation {
  street: string;
  city: string;
  state: string;
  postcode: number;
}

export interface ProfileLogin {
  username: string;
  password: string;
  salt: string;
  md5: string;
  sha1: string;
  sha256: string;
}

export interface ProfileName {
  first: string;
  last: string;
  title: string;
}

export interface ProfilePicture {
  large: string;
  medium: string;
  thumbnail: string;
}

const baseID: object = {
  name: "",
};

const baseLocation: object = {
  street: "",
  city: "",
  state: "",
  postcode: 0,
};

const baseLogin: object = {
  username: "",
  password: "",
  salt: "",
  md5: "",
  sha1: "",
  sha256: "",
};

const baseName: object = {
  first: "",
  last: "",
  title: "",
};

const basePicture: object = {
  large: "",
  medium: "",
  thumbnail: "",
};

const baseProfile: object = {
  gender: "",
  name: undefined,
  location: undefined,
  email: "",
  login: undefined,
  dob: "",
  registered: "",
  phone: "",
  cell: "",
  id: undefined,
  picture: undefined,
  nat: "",
};

const baseProfileId: object = {
  name: "",
};

const baseProfileLocation: object = {
  street: "",
  city: "",
  state: "",
  postcode: 0,
};

const baseProfileLogin: object = {
  username: "",
  password: "",
  salt: "",
  md5: "",
  sha1: "",
  sha256: "",
};

const baseProfileName: object = {
  first: "",
  last: "",
  title: "",
};

const baseProfilePicture: object = {
  large: "",
  medium: "",
  thumbnail: "",
};

function longToNumber(long: Long) {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

export const ID = {
  encode(message: ID, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.name);
    return writer;
  },
  decode(reader: Reader, length?: number): ID {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseID) as ID;
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
  fromJSON(object: any): ID {
    const message = Object.create(baseID) as ID;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<ID>): ID {
    const message = Object.create(baseID) as ID;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
  toJSON(message: ID): unknown {
    const obj: any = {};
    obj.name = message.name || "";
    return obj;
  },
};

export const Location = {
  encode(message: Location, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.street);
    writer.uint32(18).string(message.city);
    writer.uint32(26).string(message.state);
    writer.uint32(32).int64(message.postcode);
    return writer;
  },
  decode(reader: Reader, length?: number): Location {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseLocation) as Location;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.street = reader.string();
          break;
        case 2:
          message.city = reader.string();
          break;
        case 3:
          message.state = reader.string();
          break;
        case 4:
          message.postcode = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Location {
    const message = Object.create(baseLocation) as Location;
    if (object.street !== undefined && object.street !== null) {
      message.street = String(object.street);
    } else {
      message.street = "";
    }
    if (object.city !== undefined && object.city !== null) {
      message.city = String(object.city);
    } else {
      message.city = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
    if (object.postcode !== undefined && object.postcode !== null) {
      message.postcode = Number(object.postcode);
    } else {
      message.postcode = 0;
    }
    return message;
  },
  fromPartial(object: DeepPartial<Location>): Location {
    const message = Object.create(baseLocation) as Location;
    if (object.street !== undefined && object.street !== null) {
      message.street = object.street;
    } else {
      message.street = "";
    }
    if (object.city !== undefined && object.city !== null) {
      message.city = object.city;
    } else {
      message.city = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
    if (object.postcode !== undefined && object.postcode !== null) {
      message.postcode = object.postcode;
    } else {
      message.postcode = 0;
    }
    return message;
  },
  toJSON(message: Location): unknown {
    const obj: any = {};
    obj.street = message.street || "";
    obj.city = message.city || "";
    obj.state = message.state || "";
    obj.postcode = message.postcode || 0;
    return obj;
  },
};

export const Login = {
  encode(message: Login, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.username);
    writer.uint32(18).string(message.password);
    writer.uint32(26).string(message.salt);
    writer.uint32(34).string(message.md5);
    writer.uint32(42).string(message.sha1);
    writer.uint32(50).string(message.sha256);
    return writer;
  },
  decode(reader: Reader, length?: number): Login {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseLogin) as Login;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        case 3:
          message.salt = reader.string();
          break;
        case 4:
          message.md5 = reader.string();
          break;
        case 5:
          message.sha1 = reader.string();
          break;
        case 6:
          message.sha256 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Login {
    const message = Object.create(baseLogin) as Login;
    if (object.username !== undefined && object.username !== null) {
      message.username = String(object.username);
    } else {
      message.username = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = String(object.password);
    } else {
      message.password = "";
    }
    if (object.salt !== undefined && object.salt !== null) {
      message.salt = String(object.salt);
    } else {
      message.salt = "";
    }
    if (object.md5 !== undefined && object.md5 !== null) {
      message.md5 = String(object.md5);
    } else {
      message.md5 = "";
    }
    if (object.sha1 !== undefined && object.sha1 !== null) {
      message.sha1 = String(object.sha1);
    } else {
      message.sha1 = "";
    }
    if (object.sha256 !== undefined && object.sha256 !== null) {
      message.sha256 = String(object.sha256);
    } else {
      message.sha256 = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<Login>): Login {
    const message = Object.create(baseLogin) as Login;
    if (object.username !== undefined && object.username !== null) {
      message.username = object.username;
    } else {
      message.username = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = object.password;
    } else {
      message.password = "";
    }
    if (object.salt !== undefined && object.salt !== null) {
      message.salt = object.salt;
    } else {
      message.salt = "";
    }
    if (object.md5 !== undefined && object.md5 !== null) {
      message.md5 = object.md5;
    } else {
      message.md5 = "";
    }
    if (object.sha1 !== undefined && object.sha1 !== null) {
      message.sha1 = object.sha1;
    } else {
      message.sha1 = "";
    }
    if (object.sha256 !== undefined && object.sha256 !== null) {
      message.sha256 = object.sha256;
    } else {
      message.sha256 = "";
    }
    return message;
  },
  toJSON(message: Login): unknown {
    const obj: any = {};
    obj.username = message.username || "";
    obj.password = message.password || "";
    obj.salt = message.salt || "";
    obj.md5 = message.md5 || "";
    obj.sha1 = message.sha1 || "";
    obj.sha256 = message.sha256 || "";
    return obj;
  },
};

export const Name = {
  encode(message: Name, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.first);
    writer.uint32(18).string(message.last);
    writer.uint32(26).string(message.title);
    return writer;
  },
  decode(reader: Reader, length?: number): Name {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseName) as Name;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.first = reader.string();
          break;
        case 2:
          message.last = reader.string();
          break;
        case 3:
          message.title = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Name {
    const message = Object.create(baseName) as Name;
    if (object.first !== undefined && object.first !== null) {
      message.first = String(object.first);
    } else {
      message.first = "";
    }
    if (object.last !== undefined && object.last !== null) {
      message.last = String(object.last);
    } else {
      message.last = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<Name>): Name {
    const message = Object.create(baseName) as Name;
    if (object.first !== undefined && object.first !== null) {
      message.first = object.first;
    } else {
      message.first = "";
    }
    if (object.last !== undefined && object.last !== null) {
      message.last = object.last;
    } else {
      message.last = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    return message;
  },
  toJSON(message: Name): unknown {
    const obj: any = {};
    obj.first = message.first || "";
    obj.last = message.last || "";
    obj.title = message.title || "";
    return obj;
  },
};

export const Picture = {
  encode(message: Picture, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.large);
    writer.uint32(18).string(message.medium);
    writer.uint32(26).string(message.thumbnail);
    return writer;
  },
  decode(reader: Reader, length?: number): Picture {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(basePicture) as Picture;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.large = reader.string();
          break;
        case 2:
          message.medium = reader.string();
          break;
        case 3:
          message.thumbnail = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Picture {
    const message = Object.create(basePicture) as Picture;
    if (object.large !== undefined && object.large !== null) {
      message.large = String(object.large);
    } else {
      message.large = "";
    }
    if (object.medium !== undefined && object.medium !== null) {
      message.medium = String(object.medium);
    } else {
      message.medium = "";
    }
    if (object.thumbnail !== undefined && object.thumbnail !== null) {
      message.thumbnail = String(object.thumbnail);
    } else {
      message.thumbnail = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<Picture>): Picture {
    const message = Object.create(basePicture) as Picture;
    if (object.large !== undefined && object.large !== null) {
      message.large = object.large;
    } else {
      message.large = "";
    }
    if (object.medium !== undefined && object.medium !== null) {
      message.medium = object.medium;
    } else {
      message.medium = "";
    }
    if (object.thumbnail !== undefined && object.thumbnail !== null) {
      message.thumbnail = object.thumbnail;
    } else {
      message.thumbnail = "";
    }
    return message;
  },
  toJSON(message: Picture): unknown {
    const obj: any = {};
    obj.large = message.large || "";
    obj.medium = message.medium || "";
    obj.thumbnail = message.thumbnail || "";
    return obj;
  },
};

export const Profile = {
  encode(message: Profile, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.gender);
    if (message.name !== undefined && message.name !== undefined) {
      ProfileName.encode(message.name, writer.uint32(18).fork()).ldelim();
    }
    if (message.location !== undefined && message.location !== undefined) {
      ProfileLocation.encode(message.location, writer.uint32(26).fork()).ldelim();
    }
    writer.uint32(34).string(message.email);
    if (message.login !== undefined && message.login !== undefined) {
      ProfileLogin.encode(message.login, writer.uint32(42).fork()).ldelim();
    }
    writer.uint32(50).string(message.dob);
    writer.uint32(58).string(message.registered);
    writer.uint32(66).string(message.phone);
    writer.uint32(74).string(message.cell);
    if (message.id !== undefined && message.id !== undefined) {
      ProfileId.encode(message.id, writer.uint32(82).fork()).ldelim();
    }
    if (message.picture !== undefined && message.picture !== undefined) {
      ProfilePicture.encode(message.picture, writer.uint32(90).fork()).ldelim();
    }
    writer.uint32(98).string(message.nat);
    return writer;
  },
  decode(reader: Reader, length?: number): Profile {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseProfile) as Profile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gender = reader.string();
          break;
        case 2:
          message.name = ProfileName.decode(reader, reader.uint32());
          break;
        case 3:
          message.location = ProfileLocation.decode(reader, reader.uint32());
          break;
        case 4:
          message.email = reader.string();
          break;
        case 5:
          message.login = ProfileLogin.decode(reader, reader.uint32());
          break;
        case 6:
          message.dob = reader.string();
          break;
        case 7:
          message.registered = reader.string();
          break;
        case 8:
          message.phone = reader.string();
          break;
        case 9:
          message.cell = reader.string();
          break;
        case 10:
          message.id = ProfileId.decode(reader, reader.uint32());
          break;
        case 11:
          message.picture = ProfilePicture.decode(reader, reader.uint32());
          break;
        case 12:
          message.nat = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Profile {
    const message = Object.create(baseProfile) as Profile;
    if (object.gender !== undefined && object.gender !== null) {
      message.gender = String(object.gender);
    } else {
      message.gender = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = ProfileName.fromJSON(object.name);
    } else {
      message.name = undefined;
    }
    if (object.location !== undefined && object.location !== null) {
      message.location = ProfileLocation.fromJSON(object.location);
    } else {
      message.location = undefined;
    }
    if (object.email !== undefined && object.email !== null) {
      message.email = String(object.email);
    } else {
      message.email = "";
    }
    if (object.login !== undefined && object.login !== null) {
      message.login = ProfileLogin.fromJSON(object.login);
    } else {
      message.login = undefined;
    }
    if (object.dob !== undefined && object.dob !== null) {
      message.dob = String(object.dob);
    } else {
      message.dob = "";
    }
    if (object.registered !== undefined && object.registered !== null) {
      message.registered = String(object.registered);
    } else {
      message.registered = "";
    }
    if (object.phone !== undefined && object.phone !== null) {
      message.phone = String(object.phone);
    } else {
      message.phone = "";
    }
    if (object.cell !== undefined && object.cell !== null) {
      message.cell = String(object.cell);
    } else {
      message.cell = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = ProfileId.fromJSON(object.id);
    } else {
      message.id = undefined;
    }
    if (object.picture !== undefined && object.picture !== null) {
      message.picture = ProfilePicture.fromJSON(object.picture);
    } else {
      message.picture = undefined;
    }
    if (object.nat !== undefined && object.nat !== null) {
      message.nat = String(object.nat);
    } else {
      message.nat = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<Profile>): Profile {
    const message = Object.create(baseProfile) as Profile;
    if (object.gender !== undefined && object.gender !== null) {
      message.gender = object.gender;
    } else {
      message.gender = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = ProfileName.fromPartial(object.name);
    } else {
      message.name = undefined;
    }
    if (object.location !== undefined && object.location !== null) {
      message.location = ProfileLocation.fromPartial(object.location);
    } else {
      message.location = undefined;
    }
    if (object.email !== undefined && object.email !== null) {
      message.email = object.email;
    } else {
      message.email = "";
    }
    if (object.login !== undefined && object.login !== null) {
      message.login = ProfileLogin.fromPartial(object.login);
    } else {
      message.login = undefined;
    }
    if (object.dob !== undefined && object.dob !== null) {
      message.dob = object.dob;
    } else {
      message.dob = "";
    }
    if (object.registered !== undefined && object.registered !== null) {
      message.registered = object.registered;
    } else {
      message.registered = "";
    }
    if (object.phone !== undefined && object.phone !== null) {
      message.phone = object.phone;
    } else {
      message.phone = "";
    }
    if (object.cell !== undefined && object.cell !== null) {
      message.cell = object.cell;
    } else {
      message.cell = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = ProfileId.fromPartial(object.id);
    } else {
      message.id = undefined;
    }
    if (object.picture !== undefined && object.picture !== null) {
      message.picture = ProfilePicture.fromPartial(object.picture);
    } else {
      message.picture = undefined;
    }
    if (object.nat !== undefined && object.nat !== null) {
      message.nat = object.nat;
    } else {
      message.nat = "";
    }
    return message;
  },
  toJSON(message: Profile): unknown {
    const obj: any = {};
    obj.gender = message.gender || "";
    obj.name = message.name ? ProfileName.toJSON(message.name) : undefined;
    obj.location = message.location ? ProfileLocation.toJSON(message.location) : undefined;
    obj.email = message.email || "";
    obj.login = message.login ? ProfileLogin.toJSON(message.login) : undefined;
    obj.dob = message.dob || "";
    obj.registered = message.registered || "";
    obj.phone = message.phone || "";
    obj.cell = message.cell || "";
    obj.id = message.id ? ProfileId.toJSON(message.id) : undefined;
    obj.picture = message.picture ? ProfilePicture.toJSON(message.picture) : undefined;
    obj.nat = message.nat || "";
    return obj;
  },
};

export const ProfileId = {
  encode(message: ProfileId, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.name);
    return writer;
  },
  decode(reader: Reader, length?: number): ProfileId {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseProfileId) as ProfileId;
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
  fromJSON(object: any): ProfileId {
    const message = Object.create(baseProfileId) as ProfileId;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<ProfileId>): ProfileId {
    const message = Object.create(baseProfileId) as ProfileId;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
  toJSON(message: ProfileId): unknown {
    const obj: any = {};
    obj.name = message.name || "";
    return obj;
  },
};

export const ProfileLocation = {
  encode(message: ProfileLocation, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.street);
    writer.uint32(18).string(message.city);
    writer.uint32(26).string(message.state);
    writer.uint32(32).int64(message.postcode);
    return writer;
  },
  decode(reader: Reader, length?: number): ProfileLocation {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseProfileLocation) as ProfileLocation;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.street = reader.string();
          break;
        case 2:
          message.city = reader.string();
          break;
        case 3:
          message.state = reader.string();
          break;
        case 4:
          message.postcode = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): ProfileLocation {
    const message = Object.create(baseProfileLocation) as ProfileLocation;
    if (object.street !== undefined && object.street !== null) {
      message.street = String(object.street);
    } else {
      message.street = "";
    }
    if (object.city !== undefined && object.city !== null) {
      message.city = String(object.city);
    } else {
      message.city = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
    if (object.postcode !== undefined && object.postcode !== null) {
      message.postcode = Number(object.postcode);
    } else {
      message.postcode = 0;
    }
    return message;
  },
  fromPartial(object: DeepPartial<ProfileLocation>): ProfileLocation {
    const message = Object.create(baseProfileLocation) as ProfileLocation;
    if (object.street !== undefined && object.street !== null) {
      message.street = object.street;
    } else {
      message.street = "";
    }
    if (object.city !== undefined && object.city !== null) {
      message.city = object.city;
    } else {
      message.city = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
    if (object.postcode !== undefined && object.postcode !== null) {
      message.postcode = object.postcode;
    } else {
      message.postcode = 0;
    }
    return message;
  },
  toJSON(message: ProfileLocation): unknown {
    const obj: any = {};
    obj.street = message.street || "";
    obj.city = message.city || "";
    obj.state = message.state || "";
    obj.postcode = message.postcode || 0;
    return obj;
  },
};

export const ProfileLogin = {
  encode(message: ProfileLogin, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.username);
    writer.uint32(18).string(message.password);
    writer.uint32(26).string(message.salt);
    writer.uint32(34).string(message.md5);
    writer.uint32(42).string(message.sha1);
    writer.uint32(50).string(message.sha256);
    return writer;
  },
  decode(reader: Reader, length?: number): ProfileLogin {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseProfileLogin) as ProfileLogin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        case 3:
          message.salt = reader.string();
          break;
        case 4:
          message.md5 = reader.string();
          break;
        case 5:
          message.sha1 = reader.string();
          break;
        case 6:
          message.sha256 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): ProfileLogin {
    const message = Object.create(baseProfileLogin) as ProfileLogin;
    if (object.username !== undefined && object.username !== null) {
      message.username = String(object.username);
    } else {
      message.username = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = String(object.password);
    } else {
      message.password = "";
    }
    if (object.salt !== undefined && object.salt !== null) {
      message.salt = String(object.salt);
    } else {
      message.salt = "";
    }
    if (object.md5 !== undefined && object.md5 !== null) {
      message.md5 = String(object.md5);
    } else {
      message.md5 = "";
    }
    if (object.sha1 !== undefined && object.sha1 !== null) {
      message.sha1 = String(object.sha1);
    } else {
      message.sha1 = "";
    }
    if (object.sha256 !== undefined && object.sha256 !== null) {
      message.sha256 = String(object.sha256);
    } else {
      message.sha256 = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<ProfileLogin>): ProfileLogin {
    const message = Object.create(baseProfileLogin) as ProfileLogin;
    if (object.username !== undefined && object.username !== null) {
      message.username = object.username;
    } else {
      message.username = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = object.password;
    } else {
      message.password = "";
    }
    if (object.salt !== undefined && object.salt !== null) {
      message.salt = object.salt;
    } else {
      message.salt = "";
    }
    if (object.md5 !== undefined && object.md5 !== null) {
      message.md5 = object.md5;
    } else {
      message.md5 = "";
    }
    if (object.sha1 !== undefined && object.sha1 !== null) {
      message.sha1 = object.sha1;
    } else {
      message.sha1 = "";
    }
    if (object.sha256 !== undefined && object.sha256 !== null) {
      message.sha256 = object.sha256;
    } else {
      message.sha256 = "";
    }
    return message;
  },
  toJSON(message: ProfileLogin): unknown {
    const obj: any = {};
    obj.username = message.username || "";
    obj.password = message.password || "";
    obj.salt = message.salt || "";
    obj.md5 = message.md5 || "";
    obj.sha1 = message.sha1 || "";
    obj.sha256 = message.sha256 || "";
    return obj;
  },
};

export const ProfileName = {
  encode(message: ProfileName, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.first);
    writer.uint32(18).string(message.last);
    writer.uint32(26).string(message.title);
    return writer;
  },
  decode(reader: Reader, length?: number): ProfileName {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseProfileName) as ProfileName;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.first = reader.string();
          break;
        case 2:
          message.last = reader.string();
          break;
        case 3:
          message.title = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): ProfileName {
    const message = Object.create(baseProfileName) as ProfileName;
    if (object.first !== undefined && object.first !== null) {
      message.first = String(object.first);
    } else {
      message.first = "";
    }
    if (object.last !== undefined && object.last !== null) {
      message.last = String(object.last);
    } else {
      message.last = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<ProfileName>): ProfileName {
    const message = Object.create(baseProfileName) as ProfileName;
    if (object.first !== undefined && object.first !== null) {
      message.first = object.first;
    } else {
      message.first = "";
    }
    if (object.last !== undefined && object.last !== null) {
      message.last = object.last;
    } else {
      message.last = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    return message;
  },
  toJSON(message: ProfileName): unknown {
    const obj: any = {};
    obj.first = message.first || "";
    obj.last = message.last || "";
    obj.title = message.title || "";
    return obj;
  },
};

export const ProfilePicture = {
  encode(message: ProfilePicture, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.large);
    writer.uint32(18).string(message.medium);
    writer.uint32(26).string(message.thumbnail);
    return writer;
  },
  decode(reader: Reader, length?: number): ProfilePicture {
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = Object.create(baseProfilePicture) as ProfilePicture;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.large = reader.string();
          break;
        case 2:
          message.medium = reader.string();
          break;
        case 3:
          message.thumbnail = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): ProfilePicture {
    const message = Object.create(baseProfilePicture) as ProfilePicture;
    if (object.large !== undefined && object.large !== null) {
      message.large = String(object.large);
    } else {
      message.large = "";
    }
    if (object.medium !== undefined && object.medium !== null) {
      message.medium = String(object.medium);
    } else {
      message.medium = "";
    }
    if (object.thumbnail !== undefined && object.thumbnail !== null) {
      message.thumbnail = String(object.thumbnail);
    } else {
      message.thumbnail = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<ProfilePicture>): ProfilePicture {
    const message = Object.create(baseProfilePicture) as ProfilePicture;
    if (object.large !== undefined && object.large !== null) {
      message.large = object.large;
    } else {
      message.large = "";
    }
    if (object.medium !== undefined && object.medium !== null) {
      message.medium = object.medium;
    } else {
      message.medium = "";
    }
    if (object.thumbnail !== undefined && object.thumbnail !== null) {
      message.thumbnail = object.thumbnail;
    } else {
      message.thumbnail = "";
    }
    return message;
  },
  toJSON(message: ProfilePicture): unknown {
    const obj: any = {};
    obj.large = message.large || "";
    obj.medium = message.medium || "";
    obj.thumbnail = message.thumbnail || "";
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
/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "healthchain.healthchain";

export interface MsgCreateRecord {
  creator: string;
  patientId: number;
  sys: number;
  dia: number;
}

export interface MsgCreateRecordResponse {
  id: number;
}

export interface MsgUpdateRecord {
  creator: string;
  id: number;
  patientId: number;
  sys: number;
  dia: number;
}

export interface MsgUpdateRecordResponse {
}

export interface MsgDeleteRecord {
  creator: string;
  id: number;
}

export interface MsgDeleteRecordResponse {
}

function createBaseMsgCreateRecord(): MsgCreateRecord {
  return { creator: "", patientId: 0, sys: 0, dia: 0 };
}

export const MsgCreateRecord = {
  encode(message: MsgCreateRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.patientId !== 0) {
      writer.uint32(16).uint64(message.patientId);
    }
    if (message.sys !== 0) {
      writer.uint32(24).uint64(message.sys);
    }
    if (message.dia !== 0) {
      writer.uint32(32).uint64(message.dia);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.patientId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.sys = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.dia = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateRecord {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      patientId: isSet(object.patientId) ? Number(object.patientId) : 0,
      sys: isSet(object.sys) ? Number(object.sys) : 0,
      dia: isSet(object.dia) ? Number(object.dia) : 0,
    };
  },

  toJSON(message: MsgCreateRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.patientId !== undefined && (obj.patientId = Math.round(message.patientId));
    message.sys !== undefined && (obj.sys = Math.round(message.sys));
    message.dia !== undefined && (obj.dia = Math.round(message.dia));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateRecord>, I>>(object: I): MsgCreateRecord {
    const message = createBaseMsgCreateRecord();
    message.creator = object.creator ?? "";
    message.patientId = object.patientId ?? 0;
    message.sys = object.sys ?? 0;
    message.dia = object.dia ?? 0;
    return message;
  },
};

function createBaseMsgCreateRecordResponse(): MsgCreateRecordResponse {
  return { id: 0 };
}

export const MsgCreateRecordResponse = {
  encode(message: MsgCreateRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateRecordResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreateRecordResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateRecordResponse>, I>>(object: I): MsgCreateRecordResponse {
    const message = createBaseMsgCreateRecordResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgUpdateRecord(): MsgUpdateRecord {
  return { creator: "", id: 0, patientId: 0, sys: 0, dia: 0 };
}

export const MsgUpdateRecord = {
  encode(message: MsgUpdateRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.patientId !== 0) {
      writer.uint32(24).uint64(message.patientId);
    }
    if (message.sys !== 0) {
      writer.uint32(32).uint64(message.sys);
    }
    if (message.dia !== 0) {
      writer.uint32(40).uint64(message.dia);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.patientId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.sys = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.dia = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateRecord {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      patientId: isSet(object.patientId) ? Number(object.patientId) : 0,
      sys: isSet(object.sys) ? Number(object.sys) : 0,
      dia: isSet(object.dia) ? Number(object.dia) : 0,
    };
  },

  toJSON(message: MsgUpdateRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.patientId !== undefined && (obj.patientId = Math.round(message.patientId));
    message.sys !== undefined && (obj.sys = Math.round(message.sys));
    message.dia !== undefined && (obj.dia = Math.round(message.dia));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateRecord>, I>>(object: I): MsgUpdateRecord {
    const message = createBaseMsgUpdateRecord();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.patientId = object.patientId ?? 0;
    message.sys = object.sys ?? 0;
    message.dia = object.dia ?? 0;
    return message;
  },
};

function createBaseMsgUpdateRecordResponse(): MsgUpdateRecordResponse {
  return {};
}

export const MsgUpdateRecordResponse = {
  encode(_: MsgUpdateRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateRecordResponse();
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

  fromJSON(_: any): MsgUpdateRecordResponse {
    return {};
  },

  toJSON(_: MsgUpdateRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateRecordResponse>, I>>(_: I): MsgUpdateRecordResponse {
    const message = createBaseMsgUpdateRecordResponse();
    return message;
  },
};

function createBaseMsgDeleteRecord(): MsgDeleteRecord {
  return { creator: "", id: 0 };
}

export const MsgDeleteRecord = {
  encode(message: MsgDeleteRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteRecord {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgDeleteRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteRecord>, I>>(object: I): MsgDeleteRecord {
    const message = createBaseMsgDeleteRecord();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgDeleteRecordResponse(): MsgDeleteRecordResponse {
  return {};
}

export const MsgDeleteRecordResponse = {
  encode(_: MsgDeleteRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteRecordResponse();
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

  fromJSON(_: any): MsgDeleteRecordResponse {
    return {};
  },

  toJSON(_: MsgDeleteRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteRecordResponse>, I>>(_: I): MsgDeleteRecordResponse {
    const message = createBaseMsgDeleteRecordResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateRecord(request: MsgCreateRecord): Promise<MsgCreateRecordResponse>;
  UpdateRecord(request: MsgUpdateRecord): Promise<MsgUpdateRecordResponse>;
  DeleteRecord(request: MsgDeleteRecord): Promise<MsgDeleteRecordResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateRecord = this.CreateRecord.bind(this);
    this.UpdateRecord = this.UpdateRecord.bind(this);
    this.DeleteRecord = this.DeleteRecord.bind(this);
  }
  CreateRecord(request: MsgCreateRecord): Promise<MsgCreateRecordResponse> {
    const data = MsgCreateRecord.encode(request).finish();
    const promise = this.rpc.request("healthchain.healthchain.Msg", "CreateRecord", data);
    return promise.then((data) => MsgCreateRecordResponse.decode(new _m0.Reader(data)));
  }

  UpdateRecord(request: MsgUpdateRecord): Promise<MsgUpdateRecordResponse> {
    const data = MsgUpdateRecord.encode(request).finish();
    const promise = this.rpc.request("healthchain.healthchain.Msg", "UpdateRecord", data);
    return promise.then((data) => MsgUpdateRecordResponse.decode(new _m0.Reader(data)));
  }

  DeleteRecord(request: MsgDeleteRecord): Promise<MsgDeleteRecordResponse> {
    const data = MsgDeleteRecord.encode(request).finish();
    const promise = this.rpc.request("healthchain.healthchain.Msg", "DeleteRecord", data);
    return promise.then((data) => MsgDeleteRecordResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

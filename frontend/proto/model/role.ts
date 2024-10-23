// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: model/role.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Model } from "./common";

export const protobufPackage = "proto.model";

export interface Role {
  model: Model | undefined;
  name: string;
  page: Page | undefined;
}

export interface PageOperation {
  create: boolean;
  read: boolean;
  update: boolean;
  delete: boolean;
}

export interface Page {
  role: PageOperation | undefined;
  setting: PageOperation | undefined;
  account: PageOperation | undefined;
}

function createBaseRole(): Role {
  return { model: undefined, name: "", page: undefined };
}

export const Role: MessageFns<Role> = {
  encode(message: Role, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.model !== undefined) {
      Model.encode(message.model, writer.uint32(10).fork()).join();
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.page !== undefined) {
      Page.encode(message.page, writer.uint32(26).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Role {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRole();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.model = Model.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.page = Page.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Role {
    return {
      model: isSet(object.model) ? Model.fromJSON(object.model) : undefined,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      page: isSet(object.page) ? Page.fromJSON(object.page) : undefined,
    };
  },

  toJSON(message: Role): unknown {
    const obj: any = {};
    if (message.model !== undefined) {
      obj.model = Model.toJSON(message.model);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.page !== undefined) {
      obj.page = Page.toJSON(message.page);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Role>, I>>(base?: I): Role {
    return Role.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Role>, I>>(object: I): Role {
    const message = createBaseRole();
    message.model = (object.model !== undefined && object.model !== null) ? Model.fromPartial(object.model) : undefined;
    message.name = object.name ?? "";
    message.page = (object.page !== undefined && object.page !== null) ? Page.fromPartial(object.page) : undefined;
    return message;
  },
};

function createBasePageOperation(): PageOperation {
  return { create: false, read: false, update: false, delete: false };
}

export const PageOperation: MessageFns<PageOperation> = {
  encode(message: PageOperation, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.create !== false) {
      writer.uint32(8).bool(message.create);
    }
    if (message.read !== false) {
      writer.uint32(16).bool(message.read);
    }
    if (message.update !== false) {
      writer.uint32(24).bool(message.update);
    }
    if (message.delete !== false) {
      writer.uint32(32).bool(message.delete);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): PageOperation {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePageOperation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.create = reader.bool();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.read = reader.bool();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.update = reader.bool();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.delete = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PageOperation {
    return {
      create: isSet(object.create) ? globalThis.Boolean(object.create) : false,
      read: isSet(object.read) ? globalThis.Boolean(object.read) : false,
      update: isSet(object.update) ? globalThis.Boolean(object.update) : false,
      delete: isSet(object.delete) ? globalThis.Boolean(object.delete) : false,
    };
  },

  toJSON(message: PageOperation): unknown {
    const obj: any = {};
    if (message.create !== false) {
      obj.create = message.create;
    }
    if (message.read !== false) {
      obj.read = message.read;
    }
    if (message.update !== false) {
      obj.update = message.update;
    }
    if (message.delete !== false) {
      obj.delete = message.delete;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PageOperation>, I>>(base?: I): PageOperation {
    return PageOperation.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<PageOperation>, I>>(object: I): PageOperation {
    const message = createBasePageOperation();
    message.create = object.create ?? false;
    message.read = object.read ?? false;
    message.update = object.update ?? false;
    message.delete = object.delete ?? false;
    return message;
  },
};

function createBasePage(): Page {
  return { role: undefined, setting: undefined, account: undefined };
}

export const Page: MessageFns<Page> = {
  encode(message: Page, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.role !== undefined) {
      PageOperation.encode(message.role, writer.uint32(10).fork()).join();
    }
    if (message.setting !== undefined) {
      PageOperation.encode(message.setting, writer.uint32(18).fork()).join();
    }
    if (message.account !== undefined) {
      PageOperation.encode(message.account, writer.uint32(26).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Page {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.role = PageOperation.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.setting = PageOperation.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.account = PageOperation.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Page {
    return {
      role: isSet(object.role) ? PageOperation.fromJSON(object.role) : undefined,
      setting: isSet(object.setting) ? PageOperation.fromJSON(object.setting) : undefined,
      account: isSet(object.account) ? PageOperation.fromJSON(object.account) : undefined,
    };
  },

  toJSON(message: Page): unknown {
    const obj: any = {};
    if (message.role !== undefined) {
      obj.role = PageOperation.toJSON(message.role);
    }
    if (message.setting !== undefined) {
      obj.setting = PageOperation.toJSON(message.setting);
    }
    if (message.account !== undefined) {
      obj.account = PageOperation.toJSON(message.account);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Page>, I>>(base?: I): Page {
    return Page.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Page>, I>>(object: I): Page {
    const message = createBasePage();
    message.role = (object.role !== undefined && object.role !== null)
      ? PageOperation.fromPartial(object.role)
      : undefined;
    message.setting = (object.setting !== undefined && object.setting !== null)
      ? PageOperation.fromPartial(object.setting)
      : undefined;
    message.account = (object.account !== undefined && object.account !== null)
      ? PageOperation.fromPartial(object.account)
      : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}

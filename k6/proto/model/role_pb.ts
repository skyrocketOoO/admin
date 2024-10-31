// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file model/role.proto (package proto.model, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Model } from "./common_pb.js";

/**
 * @generated from message proto.model.Role
 */
export class Role extends Message<Role> {
  /**
   * @generated from field: proto.model.Model model = 1;
   */
  model?: Model;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: proto.model.Page page = 3;
   */
  page?: Page;

  constructor(data?: PartialMessage<Role>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.model.Role";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "model", kind: "message", T: Model },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "page", kind: "message", T: Page },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Role {
    return new Role().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Role {
    return new Role().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Role {
    return new Role().fromJsonString(jsonString, options);
  }

  static equals(a: Role | PlainMessage<Role> | undefined, b: Role | PlainMessage<Role> | undefined): boolean {
    return proto3.util.equals(Role, a, b);
  }
}

/**
 * @generated from message proto.model.PageOperation
 */
export class PageOperation extends Message<PageOperation> {
  /**
   * @generated from field: bool create = 1;
   */
  create = false;

  /**
   * @generated from field: bool read = 2;
   */
  read = false;

  /**
   * @generated from field: bool update = 3;
   */
  update = false;

  /**
   * @generated from field: bool delete = 4;
   */
  delete = false;

  constructor(data?: PartialMessage<PageOperation>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.model.PageOperation";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "create", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "read", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "update", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "delete", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PageOperation {
    return new PageOperation().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PageOperation {
    return new PageOperation().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PageOperation {
    return new PageOperation().fromJsonString(jsonString, options);
  }

  static equals(a: PageOperation | PlainMessage<PageOperation> | undefined, b: PageOperation | PlainMessage<PageOperation> | undefined): boolean {
    return proto3.util.equals(PageOperation, a, b);
  }
}

/**
 * @generated from message proto.model.Page
 */
export class Page extends Message<Page> {
  /**
   * @generated from field: proto.model.PageOperation role = 1;
   */
  role?: PageOperation;

  /**
   * @generated from field: proto.model.PageOperation setting = 2;
   */
  setting?: PageOperation;

  /**
   * @generated from field: proto.model.PageOperation account = 3;
   */
  account?: PageOperation;

  constructor(data?: PartialMessage<Page>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.model.Page";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "role", kind: "message", T: PageOperation },
    { no: 2, name: "setting", kind: "message", T: PageOperation },
    { no: 3, name: "account", kind: "message", T: PageOperation },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Page {
    return new Page().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Page {
    return new Page().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Page {
    return new Page().fromJsonString(jsonString, options);
  }

  static equals(a: Page | PlainMessage<Page> | undefined, b: Page | PlainMessage<Page> | undefined): boolean {
    return proto3.util.equals(Page, a, b);
  }
}


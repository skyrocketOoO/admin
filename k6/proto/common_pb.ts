// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file common.proto (package proto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum proto.Concator
 */
export enum Concator {
  /**
   * @generated from enum value: AND = 0;
   */
  AND = 0,

  /**
   * @generated from enum value: OR = 1;
   */
  OR = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(Concator)
proto3.util.setEnumType(Concator, "proto.Concator", [
  { no: 0, name: "AND" },
  { no: 1, name: "OR" },
]);

/**
 * @generated from message proto.Empty
 */
export class Empty extends Message<Empty> {
  constructor(data?: PartialMessage<Empty>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.Empty";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Empty {
    return new Empty().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Empty {
    return new Empty().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Empty {
    return new Empty().fromJsonString(jsonString, options);
  }

  static equals(a: Empty | PlainMessage<Empty> | undefined, b: Empty | PlainMessage<Empty> | undefined): boolean {
    return proto3.util.equals(Empty, a, b);
  }
}

/**
 * @generated from message proto.ListOption
 */
export class ListOption extends Message<ListOption> {
  /**
   * @generated from field: optional proto.Pager pager = 1;
   */
  pager?: Pager;

  /**
   * @generated from field: repeated proto.Sorter sorters = 2;
   */
  sorters: Sorter[] = [];

  /**
   * Root of filter expressions.
   *
   * @generated from field: optional proto.ConditionGroup conditionGroup = 3;
   */
  conditionGroup?: ConditionGroup;

  constructor(data?: PartialMessage<ListOption>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.ListOption";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pager", kind: "message", T: Pager, opt: true },
    { no: 2, name: "sorters", kind: "message", T: Sorter, repeated: true },
    { no: 3, name: "conditionGroup", kind: "message", T: ConditionGroup, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListOption {
    return new ListOption().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListOption {
    return new ListOption().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListOption {
    return new ListOption().fromJsonString(jsonString, options);
  }

  static equals(a: ListOption | PlainMessage<ListOption> | undefined, b: ListOption | PlainMessage<ListOption> | undefined): boolean {
    return proto3.util.equals(ListOption, a, b);
  }
}

/**
 * @generated from message proto.Pager
 */
export class Pager extends Message<Pager> {
  /**
   * @generated from field: int32 number = 1;
   */
  number = 0;

  /**
   * @generated from field: int32 size = 2;
   */
  size = 0;

  constructor(data?: PartialMessage<Pager>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.Pager";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "number", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pager {
    return new Pager().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pager {
    return new Pager().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pager {
    return new Pager().fromJsonString(jsonString, options);
  }

  static equals(a: Pager | PlainMessage<Pager> | undefined, b: Pager | PlainMessage<Pager> | undefined): boolean {
    return proto3.util.equals(Pager, a, b);
  }
}

/**
 * @generated from message proto.Sorter
 */
export class Sorter extends Message<Sorter> {
  /**
   * @generated from field: bool ascending = 1;
   */
  ascending = false;

  /**
   * @generated from field: string field = 2;
   */
  field = "";

  constructor(data?: PartialMessage<Sorter>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.Sorter";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ascending", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "field", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Sorter {
    return new Sorter().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Sorter {
    return new Sorter().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Sorter {
    return new Sorter().fromJsonString(jsonString, options);
  }

  static equals(a: Sorter | PlainMessage<Sorter> | undefined, b: Sorter | PlainMessage<Sorter> | undefined): boolean {
    return proto3.util.equals(Sorter, a, b);
  }
}

/**
 * @generated from message proto.Condition
 */
export class Condition extends Message<Condition> {
  /**
   * @generated from field: optional string field = 1;
   */
  field?: string;

  /**
   * same as sql operator
   *
   * @generated from field: string operator = 2;
   */
  operator = "";

  /**
   * @generated from field: string value = 3;
   */
  value = "";

  constructor(data?: PartialMessage<Condition>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.Condition";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "field", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "operator", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Condition {
    return new Condition().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Condition {
    return new Condition().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Condition {
    return new Condition().fromJsonString(jsonString, options);
  }

  static equals(a: Condition | PlainMessage<Condition> | undefined, b: Condition | PlainMessage<Condition> | undefined): boolean {
    return proto3.util.equals(Condition, a, b);
  }
}

/**
 * Group of filters with a logical operator.
 *
 * @generated from message proto.ConditionGroup
 */
export class ConditionGroup extends Message<ConditionGroup> {
  /**
   * Logical operator (AND/OR) to combine filters.
   *
   * @generated from field: proto.Concator concator = 1;
   */
  concator = Concator.AND;

  /**
   * List of individual filters.
   *
   * @generated from field: repeated proto.Condition conditions = 2;
   */
  conditions: Condition[] = [];

  /**
   * Nested filter groups for recursive logic.
   *
   * @generated from field: repeated proto.ConditionGroup nestedGroups = 3;
   */
  nestedGroups: ConditionGroup[] = [];

  constructor(data?: PartialMessage<ConditionGroup>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.ConditionGroup";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "concator", kind: "enum", T: proto3.getEnumType(Concator) },
    { no: 2, name: "conditions", kind: "message", T: Condition, repeated: true },
    { no: 3, name: "nestedGroups", kind: "message", T: ConditionGroup, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConditionGroup {
    return new ConditionGroup().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConditionGroup {
    return new ConditionGroup().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConditionGroup {
    return new ConditionGroup().fromJsonString(jsonString, options);
  }

  static equals(a: ConditionGroup | PlainMessage<ConditionGroup> | undefined, b: ConditionGroup | PlainMessage<ConditionGroup> | undefined): boolean {
    return proto3.util.equals(ConditionGroup, a, b);
  }
}


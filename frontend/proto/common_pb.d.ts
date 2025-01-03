// @generated by protoc-gen-es v1.10.0
// @generated from file common.proto (package proto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum proto.Concator
 */
export declare enum Concator {
  /**
   * @generated from enum value: AND = 0;
   */
  AND = 0,

  /**
   * @generated from enum value: OR = 1;
   */
  OR = 1,
}

/**
 * @generated from message proto.Empty
 */
export declare class Empty extends Message<Empty> {
  constructor(data?: PartialMessage<Empty>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.Empty";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Empty;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Empty;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Empty;

  static equals(a: Empty | PlainMessage<Empty> | undefined, b: Empty | PlainMessage<Empty> | undefined): boolean;
}

/**
 * @generated from message proto.ListOption
 */
export declare class ListOption extends Message<ListOption> {
  /**
   * @generated from field: optional proto.Pager pager = 1;
   */
  pager?: Pager;

  /**
   * @generated from field: repeated proto.Sorter sorters = 2;
   */
  sorters: Sorter[];

  /**
   * Root of filter expressions.
   *
   * @generated from field: optional proto.ConditionGroup conditionGroup = 3;
   */
  conditionGroup?: ConditionGroup;

  constructor(data?: PartialMessage<ListOption>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.ListOption";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListOption;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListOption;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListOption;

  static equals(a: ListOption | PlainMessage<ListOption> | undefined, b: ListOption | PlainMessage<ListOption> | undefined): boolean;
}

/**
 * @generated from message proto.Pager
 */
export declare class Pager extends Message<Pager> {
  /**
   * @generated from field: int32 number = 1;
   */
  number: number;

  /**
   * @generated from field: int32 size = 2;
   */
  size: number;

  constructor(data?: PartialMessage<Pager>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.Pager";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pager;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pager;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pager;

  static equals(a: Pager | PlainMessage<Pager> | undefined, b: Pager | PlainMessage<Pager> | undefined): boolean;
}

/**
 * @generated from message proto.Sorter
 */
export declare class Sorter extends Message<Sorter> {
  /**
   * @generated from field: bool ascending = 1;
   */
  ascending: boolean;

  /**
   * @generated from field: string field = 2;
   */
  field: string;

  constructor(data?: PartialMessage<Sorter>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.Sorter";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Sorter;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Sorter;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Sorter;

  static equals(a: Sorter | PlainMessage<Sorter> | undefined, b: Sorter | PlainMessage<Sorter> | undefined): boolean;
}

/**
 * NOT SUPPORTED: "BETWEEN" operator.
 *
 * @generated from message proto.Condition
 */
export declare class Condition extends Message<Condition> {
  /**
   * @generated from field: string field = 1;
   */
  field: string;

  /**
   * same as sql operator
   *
   * @generated from field: string operator = 2;
   */
  operator: string;

  /**
   * @generated from field: string value = 3;
   */
  value: string;

  constructor(data?: PartialMessage<Condition>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.Condition";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Condition;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Condition;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Condition;

  static equals(a: Condition | PlainMessage<Condition> | undefined, b: Condition | PlainMessage<Condition> | undefined): boolean;
}

/**
 * Group of filters with a logical operator.
 *
 * @generated from message proto.ConditionGroup
 */
export declare class ConditionGroup extends Message<ConditionGroup> {
  /**
   * Logical operator (AND/OR) to combine filters.
   *
   * @generated from field: proto.Concator concator = 1;
   */
  concator: Concator;

  /**
   * List of individual filters.
   *
   * @generated from field: repeated proto.Condition conditions = 2;
   */
  conditions: Condition[];

  /**
   * Nested filter groups for recursive logic.
   *
   * @generated from field: repeated proto.ConditionGroup conditionGroups = 3;
   */
  conditionGroups: ConditionGroup[];

  constructor(data?: PartialMessage<ConditionGroup>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.ConditionGroup";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConditionGroup;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConditionGroup;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConditionGroup;

  static equals(a: ConditionGroup | PlainMessage<ConditionGroup> | undefined, b: ConditionGroup | PlainMessage<ConditionGroup> | undefined): boolean;
}


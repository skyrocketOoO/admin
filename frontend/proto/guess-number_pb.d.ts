// @generated by protoc-gen-es v1.10.0
// @generated from file guess-number.proto (package proto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message proto.GuessNumberReq
 */
export declare class GuessNumberReq extends Message<GuessNumberReq> {
  /**
   * @generated from field: int32 lowerBound = 1;
   */
  lowerBound: number;

  /**
   * @generated from field: int32 upperBound = 2;
   */
  upperBound: number;

  /**
   * @generated from field: repeated proto.Proportion proportions = 3;
   */
  proportions: Proportion[];

  constructor(data?: PartialMessage<GuessNumberReq>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.GuessNumberReq";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GuessNumberReq;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GuessNumberReq;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GuessNumberReq;

  static equals(a: GuessNumberReq | PlainMessage<GuessNumberReq> | undefined, b: GuessNumberReq | PlainMessage<GuessNumberReq> | undefined): boolean;
}

/**
 * @generated from message proto.Proportion
 */
export declare class Proportion extends Message<Proportion> {
  /**
   * @generated from field: repeated float values = 1;
   */
  values: number[];

  constructor(data?: PartialMessage<Proportion>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.Proportion";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Proportion;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Proportion;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Proportion;

  static equals(a: Proportion | PlainMessage<Proportion> | undefined, b: Proportion | PlainMessage<Proportion> | undefined): boolean;
}

/**
 * @generated from message proto.GuessNubmerResp
 */
export declare class GuessNubmerResp extends Message<GuessNubmerResp> {
  /**
   * @generated from field: repeated int32 possibleResults = 1;
   */
  possibleResults: number[];

  constructor(data?: PartialMessage<GuessNubmerResp>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proto.GuessNubmerResp";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GuessNubmerResp;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GuessNubmerResp;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GuessNubmerResp;

  static equals(a: GuessNubmerResp | PlainMessage<GuessNubmerResp> | undefined, b: GuessNubmerResp | PlainMessage<GuessNubmerResp> | undefined): boolean;
}


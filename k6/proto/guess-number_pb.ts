// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file guess-number.proto (package proto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message proto.GuessNumberReq
 */
export class GuessNumberReq extends Message<GuessNumberReq> {
  /**
   * @generated from field: int32 lowerBound = 1;
   */
  lowerBound = 0;

  /**
   * @generated from field: int32 upperBound = 2;
   */
  upperBound = 0;

  /**
   * @generated from field: repeated proto.Proportion proportions = 3;
   */
  proportions: Proportion[] = [];

  constructor(data?: PartialMessage<GuessNumberReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.GuessNumberReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "lowerBound", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "upperBound", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "proportions", kind: "message", T: Proportion, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GuessNumberReq {
    return new GuessNumberReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GuessNumberReq {
    return new GuessNumberReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GuessNumberReq {
    return new GuessNumberReq().fromJsonString(jsonString, options);
  }

  static equals(a: GuessNumberReq | PlainMessage<GuessNumberReq> | undefined, b: GuessNumberReq | PlainMessage<GuessNumberReq> | undefined): boolean {
    return proto3.util.equals(GuessNumberReq, a, b);
  }
}

/**
 * @generated from message proto.Proportion
 */
export class Proportion extends Message<Proportion> {
  /**
   * @generated from field: repeated float values = 1;
   */
  values: number[] = [];

  constructor(data?: PartialMessage<Proportion>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.Proportion";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "values", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Proportion {
    return new Proportion().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Proportion {
    return new Proportion().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Proportion {
    return new Proportion().fromJsonString(jsonString, options);
  }

  static equals(a: Proportion | PlainMessage<Proportion> | undefined, b: Proportion | PlainMessage<Proportion> | undefined): boolean {
    return proto3.util.equals(Proportion, a, b);
  }
}

/**
 * @generated from message proto.GuessNubmerResp
 */
export class GuessNubmerResp extends Message<GuessNubmerResp> {
  /**
   * @generated from field: repeated int32 possibleResults = 1;
   */
  possibleResults: number[] = [];

  constructor(data?: PartialMessage<GuessNubmerResp>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.GuessNubmerResp";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "possibleResults", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GuessNubmerResp {
    return new GuessNubmerResp().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GuessNubmerResp {
    return new GuessNubmerResp().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GuessNubmerResp {
    return new GuessNubmerResp().fromJsonString(jsonString, options);
  }

  static equals(a: GuessNubmerResp | PlainMessage<GuessNubmerResp> | undefined, b: GuessNubmerResp | PlainMessage<GuessNubmerResp> | undefined): boolean {
    return proto3.util.equals(GuessNubmerResp, a, b);
  }
}


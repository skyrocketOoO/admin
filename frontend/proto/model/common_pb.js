// @generated by protoc-gen-es v1.10.0
// @generated from file model/common.proto (package proto.model, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message proto.model.Model
 */
export const Model = /*@__PURE__*/ proto3.makeMessageType(
  "proto.model.Model",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "created_at", kind: "message", T: Timestamp },
    { no: 3, name: "updated_at", kind: "message", T: Timestamp },
    { no: 4, name: "deleted_at", kind: "message", T: Timestamp },
  ],
);

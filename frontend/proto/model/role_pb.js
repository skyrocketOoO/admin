// @generated by protoc-gen-es v1.10.0
// @generated from file model/role.proto (package proto.model, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Model } from "./common_pb.js";

/**
 * @generated from message proto.model.Role
 */
export const Role = /*@__PURE__*/ proto3.makeMessageType(
  "proto.model.Role",
  () => [
    { no: 1, name: "model", kind: "message", T: Model },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "page", kind: "message", T: Page },
  ],
);

/**
 * @generated from message proto.model.PageOperation
 */
export const PageOperation = /*@__PURE__*/ proto3.makeMessageType(
  "proto.model.PageOperation",
  () => [
    { no: 1, name: "create", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "read", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "update", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "delete", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message proto.model.Page
 */
export const Page = /*@__PURE__*/ proto3.makeMessageType(
  "proto.model.Page",
  () => [
    { no: 1, name: "role", kind: "message", T: PageOperation },
    { no: 2, name: "setting", kind: "message", T: PageOperation },
    { no: 3, name: "account", kind: "message", T: PageOperation },
  ],
);


// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts,import_extension=ts"
// @generated from file main.proto (package proto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { BindRoleReq, CreateAccountReq, CreateRoleReq, DeleteAccountReq, DeleteRoleReq, GetRoleAuthReq, GetRoleAuthResp, ListAccountReq, ListAccountResp, ListRoleReq, ListRoleResp, LoginReq, LoginResp, LogoutReq, UnBindRoleReq, UpdateAccountReq, UpdateRoleReq } from "./main_pbts";
import { MethodKind } from "@bufbuild/protobuf";
import { Empty } from "./common_pbts";

/**
 * @generated from service proto.Main
 */
export const Main = {
  typeName: "proto.Main",
  methods: {
    /**
     * @generated from rpc proto.Main.Login
     */
    login: {
      name: "Login",
      I: LoginReq,
      O: LoginResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.Logout
     */
    logout: {
      name: "Logout",
      I: LogoutReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.CreateAccount
     */
    createAccount: {
      name: "CreateAccount",
      I: CreateAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.ListAccount
     */
    listAccount: {
      name: "ListAccount",
      I: ListAccountReq,
      O: ListAccountResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.UpdateAccount
     */
    updateAccount: {
      name: "UpdateAccount",
      I: UpdateAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.DeleteAccount
     */
    deleteAccount: {
      name: "DeleteAccount",
      I: DeleteAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.DeactiveAccount
     */
    deactiveAccount: {
      name: "DeactiveAccount",
      I: DeleteAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.ActiveAccount
     */
    activeAccount: {
      name: "ActiveAccount",
      I: DeleteAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.CreateRole
     */
    createRole: {
      name: "CreateRole",
      I: CreateRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.ListRole
     */
    listRole: {
      name: "ListRole",
      I: ListRoleReq,
      O: ListRoleResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.GetRoleAuth
     */
    getRoleAuth: {
      name: "GetRoleAuth",
      I: GetRoleAuthReq,
      O: GetRoleAuthResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.UpdateRole
     */
    updateRole: {
      name: "UpdateRole",
      I: UpdateRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.DeleteRole
     */
    deleteRole: {
      name: "DeleteRole",
      I: DeleteRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.BindRole
     */
    bindRole: {
      name: "BindRole",
      I: BindRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.Main.UnBindRole
     */
    unBindRole: {
      name: "UnBindRole",
      I: UnBindRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;


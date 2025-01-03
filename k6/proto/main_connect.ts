// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts,import_extension=ts"
// @generated from file main.proto (package proto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { BindRoleReq, CreateAccountReq, CreateRoleReq, DeleteAccountReq, DeleteRoleReq, GetRoleAuthReq, GetRoleAuthResp, ListAccountReq, ListAccountResp, ListRoleReq, ListRoleResp, LoginReq, LoginResp, LogoutReq, RasSqlReq, RasSqlResp, UnBindRoleReq, UpdateAccountReq, UpdateRoleReq } from "./main_pbts";
import { MethodKind } from "@bufbuild/protobuf";
import { Empty } from "./common_pbts";

/**
 * @generated from service proto.MainService
 */
export const MainService = {
  typeName: "proto.MainService",
  methods: {
    /**
     * @generated from rpc proto.MainService.RawSql
     */
    rawSql: {
      name: "RawSql",
      I: RasSqlReq,
      O: RasSqlResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.Login
     */
    login: {
      name: "Login",
      I: LoginReq,
      O: LoginResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.Logout
     */
    logout: {
      name: "Logout",
      I: LogoutReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.CreateAccount
     */
    createAccount: {
      name: "CreateAccount",
      I: CreateAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.ListAccount
     */
    listAccount: {
      name: "ListAccount",
      I: ListAccountReq,
      O: ListAccountResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.UpdateAccount
     */
    updateAccount: {
      name: "UpdateAccount",
      I: UpdateAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.DeleteAccount
     */
    deleteAccount: {
      name: "DeleteAccount",
      I: DeleteAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.DeactiveAccount
     */
    deactiveAccount: {
      name: "DeactiveAccount",
      I: DeleteAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.ActiveAccount
     */
    activeAccount: {
      name: "ActiveAccount",
      I: DeleteAccountReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.CreateRole
     */
    createRole: {
      name: "CreateRole",
      I: CreateRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.ListRole
     */
    listRole: {
      name: "ListRole",
      I: ListRoleReq,
      O: ListRoleResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.GetRoleAuth
     */
    getRoleAuth: {
      name: "GetRoleAuth",
      I: GetRoleAuthReq,
      O: GetRoleAuthResp,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.UpdateRole
     */
    updateRole: {
      name: "UpdateRole",
      I: UpdateRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.DeleteRole
     */
    deleteRole: {
      name: "DeleteRole",
      I: DeleteRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.BindRole
     */
    bindRole: {
      name: "BindRole",
      I: BindRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.MainService.UnBindRole
     */
    unBindRole: {
      name: "UnBindRole",
      I: UnBindRoleReq,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;


# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [common.proto](#common-proto)
    - [Empty](#proto-Empty)
    - [ListOption](#proto-ListOption)
    - [Pager](#proto-Pager)
    - [Query](#proto-Query)
    - [Sorter](#proto-Sorter)
  
- [main.proto](#main-proto)
    - [AccountData](#proto-AccountData)
    - [BindRoleRequest](#proto-BindRoleRequest)
    - [CreateAccountRequest](#proto-CreateAccountRequest)
    - [CreateRoleRequest](#proto-CreateRoleRequest)
    - [DeleteAccountRequest](#proto-DeleteAccountRequest)
    - [DeleteRoleRequest](#proto-DeleteRoleRequest)
    - [GetRoleAuthRequest](#proto-GetRoleAuthRequest)
    - [GetRoleAuthResponse](#proto-GetRoleAuthResponse)
    - [ListAccountRequest](#proto-ListAccountRequest)
    - [ListAccountResponse](#proto-ListAccountResponse)
    - [ListRoleRequest](#proto-ListRoleRequest)
    - [ListRoleResponse](#proto-ListRoleResponse)
    - [LoginRequest](#proto-LoginRequest)
    - [LoginResponse](#proto-LoginResponse)
    - [LogoutRequest](#proto-LogoutRequest)
    - [Page](#proto-Page)
    - [PageOperation](#proto-PageOperation)
    - [RoleData](#proto-RoleData)
    - [UnBindRoleRequest](#proto-UnBindRoleRequest)
    - [UpdateAccountRequest](#proto-UpdateAccountRequest)
    - [UpdateRoleRequest](#proto-UpdateRoleRequest)
  
    - [Main](#proto-Main)
  
- [Scalar Value Types](#scalar-value-types)



<a name="common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common.proto



<a name="proto-Empty"></a>

### Empty







<a name="proto-ListOption"></a>

### ListOption



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Pager | [Pager](#proto-Pager) | optional |  |
| Sorter | [Sorter](#proto-Sorter) | optional |  |
| Query | [Query](#proto-Query) | repeated |  |






<a name="proto-Pager"></a>

### Pager



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Number | [int32](#int32) |  |  |
| Size | [int32](#int32) |  |  |






<a name="proto-Query"></a>

### Query



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Fuzzy | [bool](#bool) |  |  |
| Field | [string](#string) |  |  |
| Value | [string](#string) |  |  |






<a name="proto-Sorter"></a>

### Sorter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Asc | [bool](#bool) |  |  |
| Field | [string](#string) |  |  |





 

 

 

 



<a name="main-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## main.proto



<a name="proto-AccountData"></a>

### AccountData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserName | [string](#string) |  |  |
| DisplayName | [string](#string) |  |  |
| Email | [string](#string) |  |  |
| State | [int32](#int32) |  |  |






<a name="proto-BindRoleRequest"></a>

### BindRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="proto-CreateAccountRequest"></a>

### CreateAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserName | [string](#string) |  |  |
| Password | [string](#string) |  |  |
| DisplayName | [string](#string) |  |  |






<a name="proto-CreateRoleRequest"></a>

### CreateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| Page | [Page](#proto-Page) |  |  |






<a name="proto-DeleteAccountRequest"></a>

### DeleteAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Page | [Page](#proto-Page) |  |  |






<a name="proto-DeleteRoleRequest"></a>

### DeleteRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="proto-GetRoleAuthRequest"></a>

### GetRoleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) | optional | if ID is &#34;&#34;, return auth based on session |






<a name="proto-GetRoleAuthResponse"></a>

### GetRoleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Page | [Page](#proto-Page) |  |  |






<a name="proto-ListAccountRequest"></a>

### ListAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Option | [ListOption](#proto-ListOption) | optional |  |






<a name="proto-ListAccountResponse"></a>

### ListAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| List | [AccountData](#proto-AccountData) | repeated |  |
| Total | [int64](#int64) |  |  |






<a name="proto-ListRoleRequest"></a>

### ListRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Option | [ListOption](#proto-ListOption) | optional |  |






<a name="proto-ListRoleResponse"></a>

### ListRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| List | [RoleData](#proto-RoleData) | repeated |  |
| Total | [int64](#int64) |  |  |






<a name="proto-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserName | [string](#string) |  |  |
| Password | [string](#string) |  |  |






<a name="proto-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SessionID | [string](#string) |  |  |
| Role | [model.Role](#proto-model-Role) |  |  |






<a name="proto-LogoutRequest"></a>

### LogoutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SessionID | [string](#string) |  |  |






<a name="proto-Page"></a>

### Page



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountList | [PageOperation](#proto-PageOperation) |  |  |
| Role | [PageOperation](#proto-PageOperation) |  |  |
| Setting | [PageOperation](#proto-PageOperation) |  |  |






<a name="proto-PageOperation"></a>

### PageOperation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Create | [bool](#bool) |  |  |
| Read | [bool](#bool) |  |  |
| Update | [bool](#bool) |  |  |
| Delete | [bool](#bool) |  |  |






<a name="proto-RoleData"></a>

### RoleData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="proto-UnBindRoleRequest"></a>

### UnBindRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="proto-UpdateAccountRequest"></a>

### UpdateAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| DisplayName | [string](#string) |  |  |
| State | [int32](#int32) |  | 0: no change, 1: active, 2: inactive |






<a name="proto-UpdateRoleRequest"></a>

### UpdateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Page | [Page](#proto-Page) |  |  |





 

 

 


<a name="proto-Main"></a>

### Main


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#proto-LoginRequest) | [LoginResponse](#proto-LoginResponse) |  |
| Logout | [LogoutRequest](#proto-LogoutRequest) | [Empty](#proto-Empty) |  |
| CreateAccount | [CreateAccountRequest](#proto-CreateAccountRequest) | [Empty](#proto-Empty) |  |
| ListAccount | [ListAccountRequest](#proto-ListAccountRequest) | [ListAccountResponse](#proto-ListAccountResponse) |  |
| UpdateAccount | [UpdateAccountRequest](#proto-UpdateAccountRequest) | [Empty](#proto-Empty) |  |
| DeleteAccount | [DeleteAccountRequest](#proto-DeleteAccountRequest) | [Empty](#proto-Empty) |  |
| DeactiveAccount | [DeleteAccountRequest](#proto-DeleteAccountRequest) | [Empty](#proto-Empty) |  |
| ActiveAccount | [DeleteAccountRequest](#proto-DeleteAccountRequest) | [Empty](#proto-Empty) |  |
| CreateRole | [CreateRoleRequest](#proto-CreateRoleRequest) | [Empty](#proto-Empty) |  |
| ListRole | [ListRoleRequest](#proto-ListRoleRequest) | [ListRoleResponse](#proto-ListRoleResponse) |  |
| GetRoleAuth | [GetRoleAuthRequest](#proto-GetRoleAuthRequest) | [GetRoleAuthResponse](#proto-GetRoleAuthResponse) |  |
| UpdateRole | [UpdateRoleRequest](#proto-UpdateRoleRequest) | [Empty](#proto-Empty) |  |
| DeleteRole | [DeleteRoleRequest](#proto-DeleteRoleRequest) | [Empty](#proto-Empty) |  |
| BindRole | [BindRoleRequest](#proto-BindRoleRequest) | [Empty](#proto-Empty) |  |
| UnBindRole | [UnBindRoleRequest](#proto-UnBindRoleRequest) | [Empty](#proto-Empty) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |


# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [internal/proto/word.proto](#internal_proto_word-proto)
    - [CreateWordRequest](#word-CreateWordRequest)
    - [DeleteRequest](#word-DeleteRequest)
    - [GetWordRequest](#word-GetWordRequest)
    - [GetWordResponse](#word-GetWordResponse)
    - [ListWordRequest](#word-ListWordRequest)
    - [ListWordResponse](#word-ListWordResponse)
    - [UpdateRequest](#word-UpdateRequest)
    - [Word](#word-Word)
  
    - [WordService](#word-WordService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="internal_proto_word-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## internal/proto/word.proto



<a name="word-CreateWordRequest"></a>

### CreateWordRequest
単語を作成する際のリクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [string](#string) |  |  |
| word | [string](#string) |  |  |
| meaning | [string](#string) |  |  |






<a name="word-DeleteRequest"></a>

### DeleteRequest
指定したIDの単語を削除する際のリクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="word-GetWordRequest"></a>

### GetWordRequest
単語を取得する際のリクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| userID | [string](#string) |  |  |






<a name="word-GetWordResponse"></a>

### GetWordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| word | [Word](#word-Word) |  |  |






<a name="word-ListWordRequest"></a>

### ListWordRequest
単語の一覧を取得する際のリクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [string](#string) |  |  |






<a name="word-ListWordResponse"></a>

### ListWordResponse
指定したユーザIDの単語一覧を返す


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| words | [Word](#word-Word) | repeated |  |






<a name="word-UpdateRequest"></a>

### UpdateRequest
指定したIDの単語を更新する際のリクエスト


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| meaning | [string](#string) |  |  |






<a name="word-Word"></a>

### Word
単語の型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| userID | [string](#string) |  |  |
| word | [string](#string) |  |  |
| meaning | [string](#string) |  |  |





 

 

 


<a name="word-WordService"></a>

### WordService
単語を管理するサービス

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateWord | [CreateWordRequest](#word-CreateWordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 特定の文字列をユーザの単語帳として登録する。 |
| ListWord | [ListWordRequest](#word-ListWordRequest) | [ListWordResponse](#word-ListWordResponse) | 指定したユーザIDの単語一覧を表示する。 |
| GetWord | [GetWordRequest](#word-GetWordRequest) | [GetWordResponse](#word-GetWordResponse) | 指定したIDの単語を取得する。 |
| UpdateWord | [UpdateRequest](#word-UpdateRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 指定したIDの単語の説明文を変更する。 |
| DeleteWord | [DeleteRequest](#word-DeleteRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 指定したIDの単語を削除する。 |

 



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


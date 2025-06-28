# vkGetShaderBinaryDataEXT(3) Manual Page

## Name

vkGetShaderBinaryDataEXT - Get the binary shader code from a shader object



## [](#_c_specification)C Specification

Binary shader code **can** be retrieved from a shader object using the command:

```c++
// Provided by VK_EXT_shader_object
VkResult vkGetShaderBinaryDataEXT(
    VkDevice                                    device,
    VkShaderEXT                                 shader,
    size_t*                                     pDataSize,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that shader object was created from.
- `shader` is the shader object to retrieve binary shader code from.
- `pDataSize` is a pointer to a `size_t` value related to the size of the binary shader code, as described below.
- `pData` is either `NULL` or a pointer to a buffer.

## [](#_description)Description

If `pData` is `NULL`, then the size of the binary shader code of the shader object, in bytes, is returned in `pDataSize`. Otherwise, `pDataSize` **must** point to a variable set by the application to the size of the buffer, in bytes, pointed to by `pData`, and on return the variable is overwritten with the amount of data actually written to `pData`. If `pDataSize` is less than the size of the binary shader code, nothing is written to `pData`, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`.

Note

The behavior of this command when `pDataSize` is too small differs from how some other getter-type commands work in Vulkan. Because shader binary data is only usable in its entirety, it would never be useful for the implementation to return partial data. Because of this, nothing is written to `pData` unless `pDataSize` is large enough to fit the data in its entirety.

This behavior is not consistent with the behavior described in [Opaque Binary Data Results](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-binaryresults), for historical reasons.

If the amount of data available is larger than the passed `pDataSize`, the query returns a `VK_INCOMPLETE` success status instead of a `VK_ERROR_NOT_ENOUGH_SPACE_KHR` error status.

Binary shader code retrieved using `vkGetShaderBinaryDataEXT` **can** be passed to a subsequent call to [vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateShadersEXT.html) on a compatible physical device by specifying `VK_SHADER_CODE_TYPE_BINARY_EXT` in the `codeType` member of `VkShaderCreateInfoEXT`.

The shader code returned by repeated calls to this function with the same `VkShaderEXT` is guaranteed to be invariant for the lifetime of the `VkShaderEXT` object.

Valid Usage

- [](#VUID-vkGetShaderBinaryDataEXT-None-08461)VUID-vkGetShaderBinaryDataEXT-None-08461  
  The [`shaderObject`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderObject) feature **must** be enabled
- [](#VUID-vkGetShaderBinaryDataEXT-None-08499)VUID-vkGetShaderBinaryDataEXT-None-08499  
  If `pData` is not `NULL`, it **must** be aligned to `16` bytes

Valid Usage (Implicit)

- [](#VUID-vkGetShaderBinaryDataEXT-device-parameter)VUID-vkGetShaderBinaryDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetShaderBinaryDataEXT-shader-parameter)VUID-vkGetShaderBinaryDataEXT-shader-parameter  
  `shader` **must** be a valid [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) handle
- [](#VUID-vkGetShaderBinaryDataEXT-pDataSize-parameter)VUID-vkGetShaderBinaryDataEXT-pDataSize-parameter  
  `pDataSize` **must** be a valid pointer to a `size_t` value
- [](#VUID-vkGetShaderBinaryDataEXT-pData-parameter)VUID-vkGetShaderBinaryDataEXT-pData-parameter  
  If the value referenced by `pDataSize` is not `0`, and `pData` is not `NULL`, `pData` **must** be a valid pointer to an array of `pDataSize` bytes
- [](#VUID-vkGetShaderBinaryDataEXT-shader-parent)VUID-vkGetShaderBinaryDataEXT-shader-parent  
  `shader` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetShaderBinaryDataEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# vkGetMemoryMetalHandlePropertiesEXT(3) Manual Page

## Name

vkGetMemoryMetalHandlePropertiesEXT - Get Properties of External Memory Metal Handles



## [](#_c_specification)C Specification

Metal memory handles compatible with Vulkan **may** also be created by non-Vulkan APIs using methods beyond the scope of this specification. To determine the correct parameters to use when importing such handles, call:

```c++
// Provided by VK_EXT_external_memory_metal
VkResult vkGetMemoryMetalHandlePropertiesEXT(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    const void*                                 pHandle,
    VkMemoryMetalHandlePropertiesEXT*           pMemoryMetalHandleProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that will be importing `pHandle`.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of the handle `pHandle`.
- `pHandle` is the handle which will be imported.
- `pMemoryMetalHandleProperties` is a pointer to a [VkMemoryMetalHandlePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMetalHandlePropertiesEXT.html) structure in which properties of `pHandle` are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetMemoryMetalHandlePropertiesEXT-handle-10416)VUID-vkGetMemoryMetalHandlePropertiesEXT-handle-10416  
  `pHandle` **must** point to a valid id&lt;MTLBuffer&gt;, id&lt;MTLTexture&gt; or id&lt;MTLDevice&gt;
- [](#VUID-vkGetMemoryMetalHandlePropertiesEXT-handleType-10417)VUID-vkGetMemoryMetalHandlePropertiesEXT-handleType-10417  
  `handleType` **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLBUFFER_BIT_EXT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLTEXTURE_BIT_EXT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLHEAP_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryMetalHandlePropertiesEXT-device-parameter)VUID-vkGetMemoryMetalHandlePropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryMetalHandlePropertiesEXT-handleType-parameter)VUID-vkGetMemoryMetalHandlePropertiesEXT-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value
- [](#VUID-vkGetMemoryMetalHandlePropertiesEXT-pHandle-parameter)VUID-vkGetMemoryMetalHandlePropertiesEXT-pHandle-parameter  
  `pHandle` **must** be a pointer value
- [](#VUID-vkGetMemoryMetalHandlePropertiesEXT-pMemoryMetalHandleProperties-parameter)VUID-vkGetMemoryMetalHandlePropertiesEXT-pMemoryMetalHandleProperties-parameter  
  `pMemoryMetalHandleProperties` **must** be a valid pointer to a [VkMemoryMetalHandlePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMetalHandlePropertiesEXT.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_external\_memory\_metal](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_external_memory_metal.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkMemoryMetalHandlePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMetalHandlePropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryMetalHandlePropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
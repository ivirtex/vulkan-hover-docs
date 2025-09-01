# vkGetMemoryHostPointerPropertiesEXT(3) Manual Page

## Name

vkGetMemoryHostPointerPropertiesEXT - Get properties of external memory host pointer



## [](#_c_specification)C Specification

To determine the correct parameters to use when importing host pointers, call:

```c++
// Provided by VK_EXT_external_memory_host
VkResult vkGetMemoryHostPointerPropertiesEXT(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    const void*                                 pHostPointer,
    VkMemoryHostPointerPropertiesEXT*           pMemoryHostPointerProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that will be importing `pHostPointer`.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of the handle `pHostPointer`.
- `pHostPointer` is the host pointer to import from.
- `pMemoryHostPointerProperties` is a pointer to a [VkMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHostPointerPropertiesEXT.html) structure in which the host pointer properties are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01752)VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01752  
  `handleType` **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`
- [](#VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-01753)VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-01753  
  `pHostPointer` **must** be a pointer aligned to an integer multiple of `VkPhysicalDeviceExternalMemoryHostPropertiesEXT`::`minImportedHostPointerAlignment`
- [](#VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01754)VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01754  
  If `handleType` is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT`, `pHostPointer` **must** be a pointer to host memory
- [](#VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01755)VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01755  
  If `handleType` is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`, `pHostPointer` **must** be a pointer to host mapped foreign memory

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryHostPointerPropertiesEXT-device-parameter)VUID-vkGetMemoryHostPointerPropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-parameter)VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value
- [](#VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-parameter)VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-parameter  
  `pHostPointer` **must** be a pointer value
- [](#VUID-vkGetMemoryHostPointerPropertiesEXT-pMemoryHostPointerProperties-parameter)VUID-vkGetMemoryHostPointerPropertiesEXT-pMemoryHostPointerProperties-parameter  
  `pMemoryHostPointerProperties` **must** be a valid pointer to a [VkMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHostPointerPropertiesEXT.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_external\_memory\_host](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_external_memory_host.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHostPointerPropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryHostPointerPropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
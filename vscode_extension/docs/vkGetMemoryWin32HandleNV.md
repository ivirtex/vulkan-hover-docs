# vkGetMemoryWin32HandleNV(3) Manual Page

## Name

vkGetMemoryWin32HandleNV - Retrieve Win32 handle to a device memory object



## [](#_c_specification)C Specification

To retrieve the handle corresponding to a device memory object created with [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfoNV.html)::`handleTypes` set to include `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_NV`, call:

```c++
// Provided by VK_NV_external_memory_win32
VkResult vkGetMemoryWin32HandleNV(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    VkExternalMemoryHandleTypeFlagsNV           handleType,
    HANDLE*                                     pHandle);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object.
- `handleType` is a bitmask of [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) containing a single bit specifying the type of handle requested.
- `pHandle` is a pointer to a Windows `HANDLE` in which the handle is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetMemoryWin32HandleNV-handleType-01326)VUID-vkGetMemoryWin32HandleNV-handleType-01326  
  `handleType` **must** be a flag specified in [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfoNV.html)::`handleTypes` when allocating `memory`

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryWin32HandleNV-device-parameter)VUID-vkGetMemoryWin32HandleNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryWin32HandleNV-memory-parameter)VUID-vkGetMemoryWin32HandleNV-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-vkGetMemoryWin32HandleNV-handleType-parameter)VUID-vkGetMemoryWin32HandleNV-handleType-parameter  
  `handleType` **must** be a valid combination of [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) values
- [](#VUID-vkGetMemoryWin32HandleNV-handleType-requiredbitmask)VUID-vkGetMemoryWin32HandleNV-handleType-requiredbitmask  
  `handleType` **must** not be `0`
- [](#VUID-vkGetMemoryWin32HandleNV-pHandle-parameter)VUID-vkGetMemoryWin32HandleNV-pHandle-parameter  
  `pHandle` **must** be a valid pointer to a `HANDLE` value
- [](#VUID-vkGetMemoryWin32HandleNV-memory-parent)VUID-vkGetMemoryWin32HandleNV-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_NV\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_win32.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryWin32HandleNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
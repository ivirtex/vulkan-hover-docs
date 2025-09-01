# VkMemoryGetWin32HandleInfoKHR(3) Manual Page

## Name

VkMemoryGetWin32HandleInfoKHR - Structure describing a Win32 handle memory export operation



## [](#_c_specification)C Specification

The `VkMemoryGetWin32HandleInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_memory_win32
typedef struct VkMemoryGetWin32HandleInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetWin32HandleInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` is the memory object from which the handle will be exported.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of handle requested.

## [](#_description)Description

The properties of the handle returned depend on the value of `handleType`. See [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) for a description of the properties of the defined external memory handle types.

Valid Usage

- [](#VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00662)VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00662  
  `handleType` **must** have been included in [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes` when `memory` was created
- [](#VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00663)VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00663  
  If `handleType` is defined as an NT handle, [vkGetMemoryWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryWin32HandleKHR.html) **must** be called no more than once for each valid unique combination of `memory` and `handleType`
- [](#VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00664)VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00664  
  `handleType` **must** be defined as an NT handle or a global share handle

Valid Usage (Implicit)

- [](#VUID-VkMemoryGetWin32HandleInfoKHR-sType-sType)VUID-VkMemoryGetWin32HandleInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR`
- [](#VUID-VkMemoryGetWin32HandleInfoKHR-pNext-pNext)VUID-VkMemoryGetWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryGetWin32HandleInfoKHR-memory-parameter)VUID-VkMemoryGetWin32HandleInfoKHR-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkMemoryGetWin32HandleInfoKHR-handleType-parameter)VUID-VkMemoryGetWin32HandleInfoKHR-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_win32.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryWin32HandleKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryGetWin32HandleInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
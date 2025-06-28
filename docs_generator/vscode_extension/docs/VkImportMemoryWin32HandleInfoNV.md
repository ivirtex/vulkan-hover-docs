# VkImportMemoryWin32HandleInfoNV(3) Manual Page

## Name

VkImportMemoryWin32HandleInfoNV - Import Win32 memory created on the same physical device



## [](#_c_specification)C Specification

To import memory created on the same physical device but outside of the current Vulkan instance, add a [VkImportMemoryWin32HandleInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryWin32HandleInfoNV.html) structure to the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure, specifying a handle to and the type of the memory.

The `VkImportMemoryWin32HandleInfoNV` structure is defined as:

```c++
// Provided by VK_NV_external_memory_win32
typedef struct VkImportMemoryWin32HandleInfoNV {
    VkStructureType                      sType;
    const void*                          pNext;
    VkExternalMemoryHandleTypeFlagsNV    handleType;
    HANDLE                               handle;
} VkImportMemoryWin32HandleInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleType` is `0` or a [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) value specifying the type of memory handle in `handle`.
- `handle` is a Windows `HANDLE` referring to the memory.

## [](#_description)Description

If `handleType` is `0`, this structure is ignored by consumers of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure it is chained from.

Valid Usage

- [](#VUID-VkImportMemoryWin32HandleInfoNV-handleType-01327)VUID-VkImportMemoryWin32HandleInfoNV-handleType-01327  
  `handleType` **must** not have more than one bit set
- [](#VUID-VkImportMemoryWin32HandleInfoNV-handle-01328)VUID-VkImportMemoryWin32HandleInfoNV-handle-01328  
  `handle` **must** be a valid handle to memory, obtained as specified by `handleType`

Valid Usage (Implicit)

- [](#VUID-VkImportMemoryWin32HandleInfoNV-sType-sType)VUID-VkImportMemoryWin32HandleInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV`
- [](#VUID-VkImportMemoryWin32HandleInfoNV-handleType-parameter)VUID-VkImportMemoryWin32HandleInfoNV-handleType-parameter  
  `handleType` **must** be a valid combination of [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) values

## [](#_see_also)See Also

[VK\_NV\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_win32.html), [VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMemoryWin32HandleInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
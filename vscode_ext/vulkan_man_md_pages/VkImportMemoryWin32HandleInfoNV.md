# VkImportMemoryWin32HandleInfoNV(3) Manual Page

## Name

VkImportMemoryWin32HandleInfoNV - Import Win32 memory created on the
same physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To import memory created on the same physical device but outside of the
current Vulkan instance, add a
[VkImportMemoryWin32HandleInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoNV.html)
structure to the `pNext` chain of the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure, specifying
a handle to and the type of the memory.

The `VkImportMemoryWin32HandleInfoNV` structure is defined as:

``` c
// Provided by VK_NV_external_memory_win32
typedef struct VkImportMemoryWin32HandleInfoNV {
    VkStructureType                      sType;
    const void*                          pNext;
    VkExternalMemoryHandleTypeFlagsNV    handleType;
    HANDLE                               handle;
} VkImportMemoryWin32HandleInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleType` is `0` or a
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  value specifying the type of memory handle in `handle`.

- `handle` is a Windows `HANDLE` referring to the memory.

## <a href="#_description" class="anchor"></a>Description

If `handleType` is `0`, this structure is ignored by consumers of the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure it is
chained from.

Valid Usage

- <a href="#VUID-VkImportMemoryWin32HandleInfoNV-handleType-01327"
  id="VUID-VkImportMemoryWin32HandleInfoNV-handleType-01327"></a>
  VUID-VkImportMemoryWin32HandleInfoNV-handleType-01327  
  `handleType` **must** not have more than one bit set

- <a href="#VUID-VkImportMemoryWin32HandleInfoNV-handle-01328"
  id="VUID-VkImportMemoryWin32HandleInfoNV-handle-01328"></a>
  VUID-VkImportMemoryWin32HandleInfoNV-handle-01328  
  `handle` **must** be a valid handle to memory, obtained as specified
  by `handleType`

Valid Usage (Implicit)

- <a href="#VUID-VkImportMemoryWin32HandleInfoNV-sType-sType"
  id="VUID-VkImportMemoryWin32HandleInfoNV-sType-sType"></a>
  VUID-VkImportMemoryWin32HandleInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV`

- <a href="#VUID-VkImportMemoryWin32HandleInfoNV-handleType-parameter"
  id="VUID-VkImportMemoryWin32HandleInfoNV-handleType-parameter"></a>
  VUID-VkImportMemoryWin32HandleInfoNV-handleType-parameter  
  `handleType` **must** be a valid combination of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_win32.html),
[VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMemoryWin32HandleInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

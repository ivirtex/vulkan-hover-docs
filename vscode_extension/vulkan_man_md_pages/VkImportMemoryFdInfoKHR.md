# VkImportMemoryFdInfoKHR(3) Manual Page

## Name

VkImportMemoryFdInfoKHR - Import memory created on the same physical
device from a file descriptor



## <a href="#_c_specification" class="anchor"></a>C Specification

To import memory from a POSIX file descriptor handle, add a
[VkImportMemoryFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryFdInfoKHR.html) structure to the
`pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html)
structure. The `VkImportMemoryFdInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_memory_fd
typedef struct VkImportMemoryFdInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
    int                                   fd;
} VkImportMemoryFdInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the handle type of `fd`.

- `fd` is the external handle to import.

## <a href="#_description" class="anchor"></a>Description

Importing memory from a file descriptor transfers ownership of the file
descriptor from the application to the Vulkan implementation. The
application **must** not perform any operations on the file descriptor
after a successful import. The imported memory object holds a reference
to its payload.

Applications **can** import the same payload into multiple instances of
Vulkan, into the same instance from which it was exported, and multiple
times into a given Vulkan instance. In all cases, each import operation
**must** create a distinct `VkDeviceMemory` object.

Valid Usage

- <a href="#VUID-VkImportMemoryFdInfoKHR-handleType-00667"
  id="VUID-VkImportMemoryFdInfoKHR-handleType-00667"></a>
  VUID-VkImportMemoryFdInfoKHR-handleType-00667  
  If `handleType` is not `0`, it **must** be supported for import, as
  reported by
  [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html)
  or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)

- <a href="#VUID-VkImportMemoryFdInfoKHR-fd-00668"
  id="VUID-VkImportMemoryFdInfoKHR-fd-00668"></a>
  VUID-VkImportMemoryFdInfoKHR-fd-00668  
  The memory from which `fd` was exported **must** have been created on
  the same underlying physical device as `device`

- <a href="#VUID-VkImportMemoryFdInfoKHR-handleType-00669"
  id="VUID-VkImportMemoryFdInfoKHR-handleType-00669"></a>
  VUID-VkImportMemoryFdInfoKHR-handleType-00669  
  If `handleType` is not `0`, it **must** be
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT` or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT`

- <a href="#VUID-VkImportMemoryFdInfoKHR-handleType-00670"
  id="VUID-VkImportMemoryFdInfoKHR-handleType-00670"></a>
  VUID-VkImportMemoryFdInfoKHR-handleType-00670  
  If `handleType` is not `0`, `fd` **must** be a valid handle of the
  type specified by `handleType`

- <a href="#VUID-VkImportMemoryFdInfoKHR-fd-01746"
  id="VUID-VkImportMemoryFdInfoKHR-fd-01746"></a>
  VUID-VkImportMemoryFdInfoKHR-fd-01746  
  The memory represented by `fd` **must** have been created from a
  physical device and driver that is compatible with `device` and
  `handleType`, as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-memory-handle-types-compatibility"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-memory-handle-types-compatibility</a>

- <a href="#VUID-VkImportMemoryFdInfoKHR-fd-01520"
  id="VUID-VkImportMemoryFdInfoKHR-fd-01520"></a>
  VUID-VkImportMemoryFdInfoKHR-fd-01520  
  `fd` **must** obey any requirements listed for `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-memory-handle-types-compatibility"
  target="_blank" rel="noopener">external memory handle types
  compatibility</a>

Valid Usage (Implicit)

- <a href="#VUID-VkImportMemoryFdInfoKHR-sType-sType"
  id="VUID-VkImportMemoryFdInfoKHR-sType-sType"></a>
  VUID-VkImportMemoryFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR`

- <a href="#VUID-VkImportMemoryFdInfoKHR-handleType-parameter"
  id="VUID-VkImportMemoryFdInfoKHR-handleType-parameter"></a>
  VUID-VkImportMemoryFdInfoKHR-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_fd.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMemoryFdInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

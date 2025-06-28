# VkImportMemoryFdInfoKHR(3) Manual Page

## Name

VkImportMemoryFdInfoKHR - Import memory created on the same physical device from a file descriptor



## [](#_c_specification)C Specification

To import memory from a POSIX file descriptor handle, add a [VkImportMemoryFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryFdInfoKHR.html) structure to the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure. The `VkImportMemoryFdInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_memory_fd
typedef struct VkImportMemoryFdInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
    int                                   fd;
} VkImportMemoryFdInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the handle type of `fd`.
- `fd` is the external handle to import.

## [](#_description)Description

Importing memory from a file descriptor transfers ownership of the file descriptor from the application to the Vulkan implementation. The application **must** not perform any operations on the file descriptor after a successful import. The imported memory object holds a reference to its payload.

Applications **can** import the same payload into multiple instances of Vulkan, into the same instance from which it was exported, and multiple times into a given Vulkan instance. In all cases, each import operation **must** create a distinct `VkDeviceMemory` object.

Valid Usage

- [](#VUID-VkImportMemoryFdInfoKHR-handleType-09862)VUID-VkImportMemoryFdInfoKHR-handleType-09862  
  If `handleType` is not `0`, it **must** be supported for import, as reported by [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html), [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatProperties.html) or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html)
- [](#VUID-VkImportMemoryFdInfoKHR-fd-00668)VUID-VkImportMemoryFdInfoKHR-fd-00668  
  The memory from which `fd` was exported **must** have been created on the same underlying physical device as `device`
- [](#VUID-VkImportMemoryFdInfoKHR-handleType-00669)VUID-VkImportMemoryFdInfoKHR-handleType-00669  
  If `handleType` is not `0`, it **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT`
- [](#VUID-VkImportMemoryFdInfoKHR-handleType-00670)VUID-VkImportMemoryFdInfoKHR-handleType-00670  
  If `handleType` is not `0`, `fd` **must** be a valid handle of the type specified by `handleType`
- [](#VUID-VkImportMemoryFdInfoKHR-fd-01746)VUID-VkImportMemoryFdInfoKHR-fd-01746  
  The memory represented by `fd` **must** have been created from a physical device and driver that is compatible with `device` and `handleType`, as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-memory-handle-types-compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-memory-handle-types-compatibility)
- [](#VUID-VkImportMemoryFdInfoKHR-fd-01520)VUID-VkImportMemoryFdInfoKHR-fd-01520  
  `fd` **must** obey any requirements listed for `handleType` in [external memory handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-memory-handle-types-compatibility)

Valid Usage (Implicit)

- [](#VUID-VkImportMemoryFdInfoKHR-sType-sType)VUID-VkImportMemoryFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR`
- [](#VUID-VkImportMemoryFdInfoKHR-handleType-parameter)VUID-VkImportMemoryFdInfoKHR-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_external\_memory\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_fd.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMemoryFdInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
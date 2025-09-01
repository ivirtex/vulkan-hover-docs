# VkImportMemoryMetalHandleInfoEXT(3) Manual Page

## Name

VkImportMemoryMetalHandleInfoEXT - Import Metal memory created on the same physical device



## [](#_c_specification)C Specification

To import memory from a Metal handle, add a [VkImportMemoryMetalHandleInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryMetalHandleInfoEXT.html) structure to the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure.

The `VkImportMemoryMetalHandleInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_external_memory_metal
typedef struct VkImportMemoryMetalHandleInfoEXT {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
    void*                                 handle;
} VkImportMemoryMetalHandleInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of `handle` or `name`.
- `handle` is `NULL` or the external handle to import.

## [](#_description)Description

Importing memory object payloads from Metal handles shares the ownership of the handle to the Vulkan implementation.

Applications **can** import the same payload into multiple instances of Vulkan, into the same instance from which it was exported, and multiple times into a given Vulkan instance. In all cases, each import operation **must** create a distinct `VkDeviceMemory` object.

Valid Usage

- [](#VUID-VkImportMemoryMetalHandleInfoEXT-handleType-10408)VUID-VkImportMemoryMetalHandleInfoEXT-handleType-10408  
  If `handleType` is not `0`, it **must** be supported for import, as reported by [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatProperties.html) or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html)
- [](#VUID-VkImportMemoryMetalHandleInfoEXT-handle-10409)VUID-VkImportMemoryMetalHandleInfoEXT-handle-10409  
  The memory from which `handle` was exported **must** have been created on the same underlying physical device as `device`
- [](#VUID-VkImportMemoryMetalHandleInfoEXT-handleType-10410)VUID-VkImportMemoryMetalHandleInfoEXT-handleType-10410  
  If `handleType` is not `0`, it **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLBUFFER_BIT_EXT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLTEXTURE_BIT_EXT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLHEAP_BIT_EXT`
- [](#VUID-VkImportMemoryMetalHandleInfoEXT-handleType-10411)VUID-VkImportMemoryMetalHandleInfoEXT-handleType-10411  
  If `handleType` is not `0` , `handle` **must** be a valid non-NULL handle of the type specified by `handleType`
- [](#VUID-VkImportMemoryMetalHandleInfoEXT-handle-10412)VUID-VkImportMemoryMetalHandleInfoEXT-handle-10412  
  `handle` **must** obey any requirements listed for `handleType` in [external memory handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-memory-handle-types-compatibility)

Valid Usage (Implicit)

- [](#VUID-VkImportMemoryMetalHandleInfoEXT-sType-sType)VUID-VkImportMemoryMetalHandleInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_MEMORY_METAL_HANDLE_INFO_EXT`
- [](#VUID-VkImportMemoryMetalHandleInfoEXT-handleType-parameter)VUID-VkImportMemoryMetalHandleInfoEXT-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_EXT\_external\_memory\_metal](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_external_memory_metal.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMemoryMetalHandleInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
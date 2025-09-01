# VkMemoryGetMetalHandleInfoEXT(3) Manual Page

## Name

VkMemoryGetMetalHandleInfoEXT - Structure describing a Metal handle memory export operation



## [](#_c_specification)C Specification

The `VkMemoryGetMetalHandleInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_external_memory_metal
typedef struct VkMemoryGetMetalHandleInfoEXT {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetMetalHandleInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` is the memory object from which the handle will be exported.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of handle requested.

## [](#_description)Description

The properties of the handle returned depend on the value of `handleType`. See [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) for a description of the properties of the defined external memory handle types.

Valid Usage

- [](#VUID-VkMemoryGetMetalHandleInfoEXT-memory-10413)VUID-VkMemoryGetMetalHandleInfoEXT-memory-10413  
  `memory` **must** have been created with a valid [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)
- [](#VUID-VkMemoryGetMetalHandleInfoEXT-handleType-10414)VUID-VkMemoryGetMetalHandleInfoEXT-handleType-10414  
  `handleType` **must** have been included in [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes` when `memory` was created
- [](#VUID-VkMemoryGetMetalHandleInfoEXT-handleType-10415)VUID-VkMemoryGetMetalHandleInfoEXT-handleType-10415  
  `handleType` **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLBUFFER_BIT_EXT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLTEXTURE_BIT_EXT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLHEAP_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkMemoryGetMetalHandleInfoEXT-sType-sType)VUID-VkMemoryGetMetalHandleInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_GET_METAL_HANDLE_INFO_EXT`
- [](#VUID-VkMemoryGetMetalHandleInfoEXT-pNext-pNext)VUID-VkMemoryGetMetalHandleInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryGetMetalHandleInfoEXT-memory-parameter)VUID-VkMemoryGetMetalHandleInfoEXT-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkMemoryGetMetalHandleInfoEXT-handleType-parameter)VUID-VkMemoryGetMetalHandleInfoEXT-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_EXT\_external\_memory\_metal](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_external_memory_metal.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryMetalHandleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryMetalHandleEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryGetMetalHandleInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
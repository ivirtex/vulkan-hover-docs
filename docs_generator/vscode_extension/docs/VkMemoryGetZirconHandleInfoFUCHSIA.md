# VkMemoryGetZirconHandleInfoFUCHSIA(3) Manual Page

## Name

VkMemoryGetZirconHandleInfoFUCHSIA - Structure specifying export parameters for Zircon handle to device memory



## [](#_c_specification)C Specification

`VkMemoryGetZirconHandleInfoFUCHSIA` is defined as:

```c++
// Provided by VK_FUCHSIA_external_memory
typedef struct VkMemoryGetZirconHandleInfoFUCHSIA {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetZirconHandleInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) being exported.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of the handle pointed to by [vkGetMemoryZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryZirconHandleFUCHSIA.html)::`pZirconHandle`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04775)VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04775  
  `handleType` **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`
- [](#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04776)VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04776  
  `handleType` **must** have been included in the `handleTypes` field of the `VkExportMemoryAllocateInfo` structure when the external memory was allocated

Valid Usage (Implicit)

- [](#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-sType-sType)VUID-VkMemoryGetZirconHandleInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_GET_ZIRCON_HANDLE_INFO_FUCHSIA`
- [](#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-pNext-pNext)VUID-VkMemoryGetZirconHandleInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-memory-parameter)VUID-VkMemoryGetZirconHandleInfoFUCHSIA-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-parameter)VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_FUCHSIA\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_memory.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryZirconHandleFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryGetZirconHandleInfoFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
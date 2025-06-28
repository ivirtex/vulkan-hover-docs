# VkSemaphoreCreateInfo(3) Manual Page

## Name

VkSemaphoreCreateInfo - Structure specifying parameters of a newly created semaphore



## [](#_c_specification)C Specification

The `VkSemaphoreCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSemaphoreCreateInfo {
    VkStructureType           sType;
    const void*               pNext;
    VkSemaphoreCreateFlags    flags;
} VkSemaphoreCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSemaphoreCreateInfo-pNext-06789)VUID-VkSemaphoreCreateInfo-pNext-06789  
  If the `pNext` chain includes a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure, its `exportObjectType` member **must** be `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkSemaphoreCreateInfo-sType-sType)VUID-VkSemaphoreCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO`
- [](#VUID-VkSemaphoreCreateInfo-pNext-pNext)VUID-VkSemaphoreCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html), [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfo.html), [VkExportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreWin32HandleInfoKHR.html), [VkImportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMetalSharedEventInfoEXT.html), [VkQueryLowLatencySupportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryLowLatencySupportNV.html), or [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreTypeCreateInfo.html)
- [](#VUID-VkSemaphoreCreateInfo-sType-unique)VUID-VkSemaphoreCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique, with the exception of structures of type [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html)
- [](#VUID-VkSemaphoreCreateInfo-flags-zerobitmask)VUID-VkSemaphoreCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkSemaphoreCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSemaphore.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
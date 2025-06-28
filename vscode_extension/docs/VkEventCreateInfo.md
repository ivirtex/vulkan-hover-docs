# VkEventCreateInfo(3) Manual Page

## Name

VkEventCreateInfo - Structure specifying parameters of a newly created event



## [](#_c_specification)C Specification

The `VkEventCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkEventCreateInfo {
    VkStructureType       sType;
    const void*           pNext;
    VkEventCreateFlags    flags;
} VkEventCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkEventCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateFlagBits.html) defining additional creation parameters.

## [](#_description)Description

Valid Usage

- [](#VUID-VkEventCreateInfo-pNext-06790)VUID-VkEventCreateInfo-pNext-06790  
  If the `pNext` chain includes a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure, its `exportObjectType` member **must** be `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkEventCreateInfo-sType-sType)VUID-VkEventCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EVENT_CREATE_INFO`
- [](#VUID-VkEventCreateInfo-pNext-pNext)VUID-VkEventCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) or [VkImportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMetalSharedEventInfoEXT.html)
- [](#VUID-VkEventCreateInfo-sType-unique)VUID-VkEventCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique, with the exception of structures of type [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html)
- [](#VUID-VkEventCreateInfo-flags-parameter)VUID-VkEventCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkEventCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkEventCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateEvent.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkEventCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
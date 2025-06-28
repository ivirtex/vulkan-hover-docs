# VkFenceCreateInfo(3) Manual Page

## Name

VkFenceCreateInfo - Structure specifying parameters of a newly created fence



## [](#_c_specification)C Specification

The `VkFenceCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkFenceCreateInfo {
    VkStructureType       sType;
    const void*           pNext;
    VkFenceCreateFlags    flags;
} VkFenceCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkFenceCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateFlagBits.html) specifying the initial state and behavior of the fence.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkFenceCreateInfo-sType-sType)VUID-VkFenceCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FENCE_CREATE_INFO`
- [](#VUID-VkFenceCreateInfo-pNext-pNext)VUID-VkFenceCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportFenceCreateInfo.html) or [VkExportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportFenceWin32HandleInfoKHR.html)
- [](#VUID-VkFenceCreateInfo-sType-unique)VUID-VkFenceCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkFenceCreateInfo-flags-parameter)VUID-VkFenceCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkFenceCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFenceCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateFence](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateFence.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFenceCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
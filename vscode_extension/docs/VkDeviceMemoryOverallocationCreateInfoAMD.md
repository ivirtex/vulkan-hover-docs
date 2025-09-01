# VkDeviceMemoryOverallocationCreateInfoAMD(3) Manual Page

## Name

VkDeviceMemoryOverallocationCreateInfoAMD - Specify memory overallocation behavior for a Vulkan device



## [](#_c_specification)C Specification

To specify whether device memory allocation is allowed beyond the size reported by [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html), add a [VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html) structure to the `pNext` chain of the [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) structure. If this structure is not specified, it is as if the `VK_MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD` value is used.

```c++
// Provided by VK_AMD_memory_overallocation_behavior
typedef struct VkDeviceMemoryOverallocationCreateInfoAMD {
    VkStructureType                      sType;
    const void*                          pNext;
    VkMemoryOverallocationBehaviorAMD    overallocationBehavior;
} VkDeviceMemoryOverallocationCreateInfoAMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `overallocationBehavior` is the desired overallocation behavior.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDeviceMemoryOverallocationCreateInfoAMD-sType-sType)VUID-VkDeviceMemoryOverallocationCreateInfoAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD`
- [](#VUID-VkDeviceMemoryOverallocationCreateInfoAMD-overallocationBehavior-parameter)VUID-VkDeviceMemoryOverallocationCreateInfoAMD-overallocationBehavior-parameter  
  `overallocationBehavior` **must** be a valid [VkMemoryOverallocationBehaviorAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryOverallocationBehaviorAMD.html) value

## [](#_see_also)See Also

[VK\_AMD\_memory\_overallocation\_behavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_memory_overallocation_behavior.html), [VkMemoryOverallocationBehaviorAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryOverallocationBehaviorAMD.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceMemoryOverallocationCreateInfoAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
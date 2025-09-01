# VkMemoryOverallocationBehaviorAMD(3) Manual Page

## Name

VkMemoryOverallocationBehaviorAMD - Specify memory overallocation behavior



## [](#_c_specification)C Specification

Possible values for [VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html)::`overallocationBehavior` include:

```c++
// Provided by VK_AMD_memory_overallocation_behavior
typedef enum VkMemoryOverallocationBehaviorAMD {
    VK_MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD = 0,
    VK_MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD = 1,
    VK_MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD = 2,
} VkMemoryOverallocationBehaviorAMD;
```

## [](#_description)Description

- `VK_MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD` lets the implementation decide if overallocation is allowed.
- `VK_MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD` specifies overallocation is allowed if platform permits.
- `VK_MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD` specifies the application is not allowed to allocate device memory beyond the heap sizes reported by [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html). Allocations that are not explicitly made by the application within the scope of the Vulkan instance are not accounted for.

## [](#_see_also)See Also

[VK\_AMD\_memory\_overallocation\_behavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_memory_overallocation_behavior.html), [VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryOverallocationBehaviorAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
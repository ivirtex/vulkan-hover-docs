# VkMemoryOverallocationBehaviorAMD(3) Manual Page

## Name

VkMemoryOverallocationBehaviorAMD - Specify memory overallocation
behavior



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values for
[VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html)::`overallocationBehavior`
include:

``` c
// Provided by VK_AMD_memory_overallocation_behavior
typedef enum VkMemoryOverallocationBehaviorAMD {
    VK_MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD = 0,
    VK_MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD = 1,
    VK_MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD = 2,
} VkMemoryOverallocationBehaviorAMD;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD` lets the
  implementation decide if overallocation is allowed.

- `VK_MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD` specifies
  overallocation is allowed if platform permits.

- `VK_MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD` specifies the
  application is not allowed to allocate device memory beyond the heap
  sizes reported by
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html).
  Allocations that are not explicitly made by the application within the
  scope of the Vulkan instance are not accounted for.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_memory_overallocation_behavior](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_memory_overallocation_behavior.html),
[VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryOverallocationBehaviorAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

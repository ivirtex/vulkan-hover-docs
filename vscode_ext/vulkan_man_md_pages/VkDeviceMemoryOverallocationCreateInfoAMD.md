# VkDeviceMemoryOverallocationCreateInfoAMD(3) Manual Page

## Name

VkDeviceMemoryOverallocationCreateInfoAMD - Specify memory
overallocation behavior for a Vulkan device



## <a href="#_c_specification" class="anchor"></a>C Specification

To specify whether device memory allocation is allowed beyond the size
reported by
[VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html),
add a
[VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html)
structure to the `pNext` chain of the
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) structure. If this
structure is not specified, it is as if the
`VK_MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD` value is used.

``` c
// Provided by VK_AMD_memory_overallocation_behavior
typedef struct VkDeviceMemoryOverallocationCreateInfoAMD {
    VkStructureType                      sType;
    const void*                          pNext;
    VkMemoryOverallocationBehaviorAMD    overallocationBehavior;
} VkDeviceMemoryOverallocationCreateInfoAMD;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `overallocationBehavior` is the desired overallocation behavior.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceMemoryOverallocationCreateInfoAMD-sType-sType"
  id="VUID-VkDeviceMemoryOverallocationCreateInfoAMD-sType-sType"></a>
  VUID-VkDeviceMemoryOverallocationCreateInfoAMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD`

- <a
  href="#VUID-VkDeviceMemoryOverallocationCreateInfoAMD-overallocationBehavior-parameter"
  id="VUID-VkDeviceMemoryOverallocationCreateInfoAMD-overallocationBehavior-parameter"></a>
  VUID-VkDeviceMemoryOverallocationCreateInfoAMD-overallocationBehavior-parameter  
  `overallocationBehavior` **must** be a valid
  [VkMemoryOverallocationBehaviorAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOverallocationBehaviorAMD.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_memory_overallocation_behavior](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_memory_overallocation_behavior.html),
[VkMemoryOverallocationBehaviorAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOverallocationBehaviorAMD.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceMemoryOverallocationCreateInfoAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# VkPerformanceOverrideInfoINTEL(3) Manual Page

## Name

VkPerformanceOverrideInfoINTEL - Performance override information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceOverrideInfoINTEL` structure is defined as:

``` c
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceOverrideInfoINTEL {
    VkStructureType                   sType;
    const void*                       pNext;
    VkPerformanceOverrideTypeINTEL    type;
    VkBool32                          enable;
    uint64_t                          parameter;
} VkPerformanceOverrideInfoINTEL;
```

## <a href="#_members" class="anchor"></a>Members

- `type` is the particular
  [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideTypeINTEL.html)
  to set.

- `enable` defines whether the override is enabled.

- `parameter` is a potential required parameter for the override.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPerformanceOverrideInfoINTEL-sType-sType"
  id="VUID-VkPerformanceOverrideInfoINTEL-sType-sType"></a>
  VUID-VkPerformanceOverrideInfoINTEL-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL`

- <a href="#VUID-VkPerformanceOverrideInfoINTEL-pNext-pNext"
  id="VUID-VkPerformanceOverrideInfoINTEL-pNext-pNext"></a>
  VUID-VkPerformanceOverrideInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPerformanceOverrideInfoINTEL-type-parameter"
  id="VUID-VkPerformanceOverrideInfoINTEL-type-parameter"></a>
  VUID-VkPerformanceOverrideInfoINTEL-type-parameter  
  `type` **must** be a valid
  [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideTypeINTEL.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideTypeINTEL.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdSetPerformanceOverrideINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPerformanceOverrideINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceOverrideInfoINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

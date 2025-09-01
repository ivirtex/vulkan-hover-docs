# VkPerformanceOverrideInfoINTEL(3) Manual Page

## Name

VkPerformanceOverrideInfoINTEL - Performance override information



## [](#_c_specification)C Specification

The `VkPerformanceOverrideInfoINTEL` structure is defined as:

```c++
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceOverrideInfoINTEL {
    VkStructureType                   sType;
    const void*                       pNext;
    VkPerformanceOverrideTypeINTEL    type;
    VkBool32                          enable;
    uint64_t                          parameter;
} VkPerformanceOverrideInfoINTEL;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` is the particular [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideTypeINTEL.html) to set.
- `enable` defines whether the override is enabled.
- `parameter` is a potential required parameter for the override.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPerformanceOverrideInfoINTEL-sType-sType)VUID-VkPerformanceOverrideInfoINTEL-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL`
- [](#VUID-VkPerformanceOverrideInfoINTEL-pNext-pNext)VUID-VkPerformanceOverrideInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPerformanceOverrideInfoINTEL-type-parameter)VUID-VkPerformanceOverrideInfoINTEL-type-parameter  
  `type` **must** be a valid [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideTypeINTEL.html) value

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideTypeINTEL.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdSetPerformanceOverrideINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPerformanceOverrideINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceOverrideInfoINTEL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
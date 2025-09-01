# VkPerformanceConfigurationAcquireInfoINTEL(3) Manual Page

## Name

VkPerformanceConfigurationAcquireInfoINTEL - Acquire a configuration to capture performance data



## [](#_c_specification)C Specification

The `VkPerformanceConfigurationAcquireInfoINTEL` structure is defined as:

```c++
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceConfigurationAcquireInfoINTEL {
    VkStructureType                        sType;
    const void*                            pNext;
    VkPerformanceConfigurationTypeINTEL    type;
} VkPerformanceConfigurationAcquireInfoINTEL;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` is one of the [VkPerformanceConfigurationTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationTypeINTEL.html) type of performance configuration that will be acquired.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPerformanceConfigurationAcquireInfoINTEL-sType-sType)VUID-VkPerformanceConfigurationAcquireInfoINTEL-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL`
- [](#VUID-VkPerformanceConfigurationAcquireInfoINTEL-pNext-pNext)VUID-VkPerformanceConfigurationAcquireInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPerformanceConfigurationAcquireInfoINTEL-type-parameter)VUID-VkPerformanceConfigurationAcquireInfoINTEL-type-parameter  
  `type` **must** be a valid [VkPerformanceConfigurationTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationTypeINTEL.html) value

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkPerformanceConfigurationTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationTypeINTEL.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkAcquirePerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquirePerformanceConfigurationINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceConfigurationAcquireInfoINTEL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkPerformanceValueINTEL(3) Manual Page

## Name

VkPerformanceValueINTEL - Container for value and types of parameters that can be queried



## [](#_c_specification)C Specification

The `VkPerformanceValueINTEL` structure is defined as:

```c++
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceValueINTEL {
    VkPerformanceValueTypeINTEL    type;
    VkPerformanceValueDataINTEL    data;
} VkPerformanceValueINTEL;
```

## [](#_members)Members

- `type` is a [VkPerformanceValueTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueTypeINTEL.html) value specifying the type of the returned data.
- `data` is a [VkPerformanceValueDataINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueDataINTEL.html) union specifying the value of the returned data.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkPerformanceValueDataINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueDataINTEL.html), [VkPerformanceValueTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueTypeINTEL.html), [vkGetPerformanceParameterINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPerformanceParameterINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceValueINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
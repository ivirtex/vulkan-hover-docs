# VkPerformanceValueTypeINTEL(3) Manual Page

## Name

VkPerformanceValueTypeINTEL - Type of the parameters that can be queried



## [](#_c_specification)C Specification

Possible values of [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueINTEL.html)::`type`, specifying the type of the data returned in [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueINTEL.html)::`data`, are:

- `VK_PERFORMANCE_VALUE_TYPE_UINT32_INTEL` specifies that unsigned 32-bit integer data is returned in `data.value32`.
- `VK_PERFORMANCE_VALUE_TYPE_UINT64_INTEL` specifies that unsigned 64-bit integer data is returned in `data.value64`.
- `VK_PERFORMANCE_VALUE_TYPE_FLOAT_INTEL` specifies that floating-point data is returned in `data.valueFloat`.
- `VK_PERFORMANCE_VALUE_TYPE_BOOL_INTEL` specifies that [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) data is returned in `data.valueBool`.
- `VK_PERFORMANCE_VALUE_TYPE_STRING_INTEL` specifies that a pointer to a null-terminated UTF-8 string is returned in `data.valueString`. The pointer is valid for the lifetime of the `device` parameter passed to [vkGetPerformanceParameterINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPerformanceParameterINTEL.html).

```c++
// Provided by VK_INTEL_performance_query
typedef enum VkPerformanceValueTypeINTEL {
    VK_PERFORMANCE_VALUE_TYPE_UINT32_INTEL = 0,
    VK_PERFORMANCE_VALUE_TYPE_UINT64_INTEL = 1,
    VK_PERFORMANCE_VALUE_TYPE_FLOAT_INTEL = 2,
    VK_PERFORMANCE_VALUE_TYPE_BOOL_INTEL = 3,
    VK_PERFORMANCE_VALUE_TYPE_STRING_INTEL = 4,
} VkPerformanceValueTypeINTEL;
```

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceValueTypeINTEL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
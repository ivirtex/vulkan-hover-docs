# VkPerformanceOverrideTypeINTEL(3) Manual Page

## Name

VkPerformanceOverrideTypeINTEL - Performance override type



## [](#_c_specification)C Specification

Possible values of [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideInfoINTEL.html)::`type`, specifying performance override types, are:

```c++
// Provided by VK_INTEL_performance_query
typedef enum VkPerformanceOverrideTypeINTEL {
    VK_PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL = 0,
    VK_PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL = 1,
} VkPerformanceOverrideTypeINTEL;
```

## [](#_description)Description

- `VK_PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL` turns all rendering operations into noop.
- `VK_PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL` stalls the stream of commands until all previously emitted commands have completed and all caches been flushed and invalidated.

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideInfoINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceOverrideTypeINTEL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
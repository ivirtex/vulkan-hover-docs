# VkPerformanceOverrideTypeINTEL(3) Manual Page

## Name

VkPerformanceOverrideTypeINTEL - Performance override type



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideInfoINTEL.html)::`type`,
specifying performance override types, are:

``` c
// Provided by VK_INTEL_performance_query
typedef enum VkPerformanceOverrideTypeINTEL {
    VK_PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL = 0,
    VK_PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL = 1,
} VkPerformanceOverrideTypeINTEL;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL` turns all rendering
  operations into noop.

- `VK_PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL` stalls the
  stream of commands until all previously emitted commands have
  completed and all caches been flushed and invalidated.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideInfoINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceOverrideTypeINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

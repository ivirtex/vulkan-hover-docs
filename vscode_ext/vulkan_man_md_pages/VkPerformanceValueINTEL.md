# VkPerformanceValueINTEL(3) Manual Page

## Name

VkPerformanceValueINTEL - Container for value and types of parameters
that can be queried



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceValueINTEL` structure is defined as:

``` c
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceValueINTEL {
    VkPerformanceValueTypeINTEL    type;
    VkPerformanceValueDataINTEL    data;
} VkPerformanceValueINTEL;
```

## <a href="#_members" class="anchor"></a>Members

- `type` is a
  [VkPerformanceValueTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueTypeINTEL.html) value
  specifying the type of the returned data.

- `data` is a
  [VkPerformanceValueDataINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueDataINTEL.html) union
  specifying the value of the returned data.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkPerformanceValueDataINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueDataINTEL.html),
[VkPerformanceValueTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueTypeINTEL.html),
[vkGetPerformanceParameterINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPerformanceParameterINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceValueINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

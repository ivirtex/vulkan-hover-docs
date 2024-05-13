# VkPerformanceValueTypeINTEL(3) Manual Page

## Name

VkPerformanceValueTypeINTEL - Type of the parameters that can be queried



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueINTEL.html)::`type`,
specifying the type of the data returned in
[VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueINTEL.html)::`data`, are:

- `VK_PERFORMANCE_VALUE_TYPE_UINT32_INTEL` specifies that unsigned
  32-bit integer data is returned in `data.value32`.

- `VK_PERFORMANCE_VALUE_TYPE_UINT64_INTEL` specifies that unsigned
  64-bit integer data is returned in `data.value64`.

- `VK_PERFORMANCE_VALUE_TYPE_FLOAT_INTEL` specifies that floating-point
  data is returned in `data.valueFloat`.

- `VK_PERFORMANCE_VALUE_TYPE_BOOL_INTEL` specifies that
  [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) data is returned in `data.valueBool`.

- `VK_PERFORMANCE_VALUE_TYPE_STRING_INTEL` specifies that a pointer to a
  null-terminated UTF-8 string is returned in `data.valueString`. The
  pointer is valid for the lifetime of the `device` parameter passed to
  [vkGetPerformanceParameterINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPerformanceParameterINTEL.html).

``` c
// Provided by VK_INTEL_performance_query
typedef enum VkPerformanceValueTypeINTEL {
    VK_PERFORMANCE_VALUE_TYPE_UINT32_INTEL = 0,
    VK_PERFORMANCE_VALUE_TYPE_UINT64_INTEL = 1,
    VK_PERFORMANCE_VALUE_TYPE_FLOAT_INTEL = 2,
    VK_PERFORMANCE_VALUE_TYPE_BOOL_INTEL = 3,
    VK_PERFORMANCE_VALUE_TYPE_STRING_INTEL = 4,
} VkPerformanceValueTypeINTEL;
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceValueTypeINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

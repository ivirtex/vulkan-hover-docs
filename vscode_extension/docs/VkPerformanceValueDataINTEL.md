# VkPerformanceValueDataINTEL(3) Manual Page

## Name

VkPerformanceValueDataINTEL - Values returned for the parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceValueDataINTEL` union is defined as:

``` c
// Provided by VK_INTEL_performance_query
typedef union VkPerformanceValueDataINTEL {
    uint32_t       value32;
    uint64_t       value64;
    float          valueFloat;
    VkBool32       valueBool;
    const char*    valueString;
} VkPerformanceValueDataINTEL;
```

## <a href="#_members" class="anchor"></a>Members

- `data.value32` represents 32-bit integer data.

- `data.value64` represents 64-bit integer data.

- `data.valueFloat` represents floating-point data.

- `data.valueBool` represents [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) data.

- `data.valueString` represents a pointer to a null-terminated UTF-8
  string.

## <a href="#_description" class="anchor"></a>Description

The correct member of the union is determined by the associated
[VkPerformanceValueTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueTypeINTEL.html) value.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceValueDataINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# VkPerformanceValueDataINTEL(3) Manual Page

## Name

VkPerformanceValueDataINTEL - Values returned for the parameters



## [](#_c_specification)C Specification

The `VkPerformanceValueDataINTEL` union is defined as:

```c++
// Provided by VK_INTEL_performance_query
typedef union VkPerformanceValueDataINTEL {
    uint32_t       value32;
    uint64_t       value64;
    float          valueFloat;
    VkBool32       valueBool;
    const char*    valueString;
} VkPerformanceValueDataINTEL;
```

## [](#_members)Members

- `value32` represents 32-bit integer data.
- `value64` represents 64-bit integer data.
- `valueFloat` represents floating-point data.
- `valueBool` represents [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) data.
- `valueString` represents a pointer to a null-terminated UTF-8 string.

## [](#_description)Description

The correct member of the union is determined by the associated [VkPerformanceValueTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueTypeINTEL.html) value.

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceValueDataINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
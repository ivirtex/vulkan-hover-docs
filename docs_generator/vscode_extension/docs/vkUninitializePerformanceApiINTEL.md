# vkUninitializePerformanceApiINTEL(3) Manual Page

## Name

vkUninitializePerformanceApiINTEL - Uninitialize a device for performance queries



## [](#_c_specification)C Specification

Once performance query operations have completed, uninitialize the device for performance queries with the call:

```c++
// Provided by VK_INTEL_performance_query
void vkUninitializePerformanceApiINTEL(
    VkDevice                                    device);
```

## [](#_parameters)Parameters

- `device` is the logical device used for the queries.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkUninitializePerformanceApiINTEL-device-parameter)VUID-vkUninitializePerformanceApiINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkUninitializePerformanceApiINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
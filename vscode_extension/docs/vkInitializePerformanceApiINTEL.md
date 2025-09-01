# vkInitializePerformanceApiINTEL(3) Manual Page

## Name

vkInitializePerformanceApiINTEL - Initialize a device for performance queries



## [](#_c_specification)C Specification

Prior to creating a performance query pool, initialize the device for performance queries with the call:

```c++
// Provided by VK_INTEL_performance_query
VkResult vkInitializePerformanceApiINTEL(
    VkDevice                                    device,
    const VkInitializePerformanceApiInfoINTEL*  pInitializeInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device used for the queries.
- `pInitializeInfo` is a pointer to a [VkInitializePerformanceApiInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInitializePerformanceApiInfoINTEL.html) structure specifying initialization parameters.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkInitializePerformanceApiINTEL-device-parameter)VUID-vkInitializePerformanceApiINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkInitializePerformanceApiINTEL-pInitializeInfo-parameter)VUID-vkInitializePerformanceApiINTEL-pInitializeInfo-parameter  
  `pInitializeInfo` **must** be a valid pointer to a valid [VkInitializePerformanceApiInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInitializePerformanceApiInfoINTEL.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkInitializePerformanceApiInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInitializePerformanceApiInfoINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkInitializePerformanceApiINTEL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
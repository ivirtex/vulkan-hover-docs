# vkGetPerformanceParameterINTEL(3) Manual Page

## Name

vkGetPerformanceParameterINTEL - Query performance capabilities of the device



## [](#_c_specification)C Specification

Some performance query features of a device can be discovered with the call:

```c++
// Provided by VK_INTEL_performance_query
VkResult vkGetPerformanceParameterINTEL(
    VkDevice                                    device,
    VkPerformanceParameterTypeINTEL             parameter,
    VkPerformanceValueINTEL*                    pValue);
```

## [](#_parameters)Parameters

- `device` is the logical device to query.
- `parameter` is the parameter to query.
- `pValue` is a pointer to a [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueINTEL.html) structure in which the type and value of the parameter are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPerformanceParameterINTEL-device-parameter)VUID-vkGetPerformanceParameterINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPerformanceParameterINTEL-parameter-parameter)VUID-vkGetPerformanceParameterINTEL-parameter-parameter  
  `parameter` **must** be a valid [VkPerformanceParameterTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceParameterTypeINTEL.html) value
- [](#VUID-vkGetPerformanceParameterINTEL-pValue-parameter)VUID-vkGetPerformanceParameterINTEL-pValue-parameter  
  `pValue` **must** be a valid pointer to a [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueINTEL.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPerformanceParameterTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceParameterTypeINTEL.html), [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPerformanceParameterINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
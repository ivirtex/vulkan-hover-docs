# vkAcquirePerformanceConfigurationINTEL(3) Manual Page

## Name

vkAcquirePerformanceConfigurationINTEL - Acquire the performance query capability



## [](#_c_specification)C Specification

To acquire a device performance configuration, call:

```c++
// Provided by VK_INTEL_performance_query
VkResult vkAcquirePerformanceConfigurationINTEL(
    VkDevice                                    device,
    const VkPerformanceConfigurationAcquireInfoINTEL* pAcquireInfo,
    VkPerformanceConfigurationINTEL*            pConfiguration);
```

## [](#_parameters)Parameters

- `device` is the logical device that the performance query commands will be submitted to.
- `pAcquireInfo` is a pointer to a [VkPerformanceConfigurationAcquireInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html) structure, specifying the performance configuration to acquire.
- `pConfiguration` is a pointer to a `VkPerformanceConfigurationINTEL` handle in which the resulting configuration object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkAcquirePerformanceConfigurationINTEL-device-parameter)VUID-vkAcquirePerformanceConfigurationINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAcquirePerformanceConfigurationINTEL-pAcquireInfo-parameter)VUID-vkAcquirePerformanceConfigurationINTEL-pAcquireInfo-parameter  
  `pAcquireInfo` **must** be a valid pointer to a valid [VkPerformanceConfigurationAcquireInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html) structure
- [](#VUID-vkAcquirePerformanceConfigurationINTEL-pConfiguration-parameter)VUID-vkAcquirePerformanceConfigurationINTEL-pConfiguration-parameter  
  `pConfiguration` **must** be a valid pointer to a [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationINTEL.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPerformanceConfigurationAcquireInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html), [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAcquirePerformanceConfigurationINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
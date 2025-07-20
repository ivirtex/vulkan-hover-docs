# vkReleasePerformanceConfigurationINTEL(3) Manual Page

## Name

vkReleasePerformanceConfigurationINTEL - Release a configuration to capture performance data



## [](#_c_specification)C Specification

To release a device performance configuration, call:

```c++
// Provided by VK_INTEL_performance_query
VkResult vkReleasePerformanceConfigurationINTEL(
    VkDevice                                    device,
    VkPerformanceConfigurationINTEL             configuration);
```

## [](#_parameters)Parameters

- `device` is the device associated to the configuration object to release.
- `configuration` is the configuration object to release.

## [](#_description)Description

Valid Usage

- [](#VUID-vkReleasePerformanceConfigurationINTEL-configuration-02737)VUID-vkReleasePerformanceConfigurationINTEL-configuration-02737  
  `configuration` **must** not be released before all command buffers submitted while the configuration was set are in [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)

Valid Usage (Implicit)

- [](#VUID-vkReleasePerformanceConfigurationINTEL-device-parameter)VUID-vkReleasePerformanceConfigurationINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkReleasePerformanceConfigurationINTEL-configuration-parameter)VUID-vkReleasePerformanceConfigurationINTEL-configuration-parameter  
  If `configuration` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `configuration` **must** be a valid [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationINTEL.html) handle
- [](#VUID-vkReleasePerformanceConfigurationINTEL-configuration-parent)VUID-vkReleasePerformanceConfigurationINTEL-configuration-parent  
  If `configuration` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `configuration` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkReleasePerformanceConfigurationINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
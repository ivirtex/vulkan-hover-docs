# vkReleaseProfilingLockKHR(3) Manual Page

## Name

vkReleaseProfilingLockKHR - Releases the profiling lock



## [](#_c_specification)C Specification

To release the profiling lock, call:

```c++
// Provided by VK_KHR_performance_query
void vkReleaseProfilingLockKHR(
    VkDevice                                    device);
```

## [](#_parameters)Parameters

- `device` is the logical device to cease profiling on.

## [](#_description)Description

Valid Usage

- [](#VUID-vkReleaseProfilingLockKHR-device-03235)VUID-vkReleaseProfilingLockKHR-device-03235  
  The profiling lock of `device` **must** have been held via a previous successful call to [vkAcquireProfilingLockKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireProfilingLockKHR.html)

Valid Usage (Implicit)

- [](#VUID-vkReleaseProfilingLockKHR-device-parameter)VUID-vkReleaseProfilingLockKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkReleaseProfilingLockKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
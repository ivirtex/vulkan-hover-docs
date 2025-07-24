# vkAcquireProfilingLockKHR(3) Manual Page

## Name

vkAcquireProfilingLockKHR - Acquires the profiling lock



## [](#_c_specification)C Specification

To record and submit a command buffer containing a performance query pool the profiling lock **must** be held. The profiling lock **must** be acquired prior to any call to [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBeginCommandBuffer.html) that will be using a performance query pool. The profiling lock **must** be held while any command buffer containing a performance query pool is in the *recording*, *executable*, or *pending state*. To acquire the profiling lock, call:

```c++
// Provided by VK_KHR_performance_query
VkResult vkAcquireProfilingLockKHR(
    VkDevice                                    device,
    const VkAcquireProfilingLockInfoKHR*        pInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device to profile.
- `pInfo` is a pointer to a [VkAcquireProfilingLockInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireProfilingLockInfoKHR.html) structure containing information about how the profiling is to be acquired.

## [](#_description)Description

Implementations **may** allow multiple actors to hold the profiling lock concurrently.

Valid Usage (Implicit)

- [](#VUID-vkAcquireProfilingLockKHR-device-parameter)VUID-vkAcquireProfilingLockKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAcquireProfilingLockKHR-pInfo-parameter)VUID-vkAcquireProfilingLockKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkAcquireProfilingLockInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireProfilingLockInfoKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`
- `VK_TIMEOUT`

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkAcquireProfilingLockInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireProfilingLockInfoKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAcquireProfilingLockKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
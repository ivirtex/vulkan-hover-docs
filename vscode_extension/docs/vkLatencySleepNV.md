# vkLatencySleepNV(3) Manual Page

## Name

vkLatencySleepNV - Trigger low latency mode Sleep



## [](#_c_specification)C Specification

To provide the synchronization primitive used to delay host CPU work for lower latency rendering, call:

```c++
// Provided by VK_NV_low_latency2
VkResult vkLatencySleepNV(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkLatencySleepInfoNV*                 pSleepInfo);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to delay associated CPU work based on [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySubmissionPresentIdNV.html) submissions.
- `pSleepInfo` is a pointer to a [VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepInfoNV.html) structure specifying the parameters of the latency sleep.

## [](#_description)Description

`vkLatencySleepNV` returns immediately. Applications **should** use [vkWaitSemaphores](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitSemaphores.html) with `pSleepInfo->signalSemaphore` to delay host CPU work. CPU work refers to application work done before presenting which includes but is not limited to: input sampling, simulation, command buffer recording, command buffer submission, and present submission. Applications **should** call this function before input sampling, and exactly once between presents.

Valid Usage (Implicit)

- [](#VUID-vkLatencySleepNV-device-parameter)VUID-vkLatencySleepNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkLatencySleepNV-swapchain-parameter)VUID-vkLatencySleepNV-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkLatencySleepNV-pSleepInfo-parameter)VUID-vkLatencySleepNV-pSleepInfo-parameter  
  `pSleepInfo` **must** be a valid pointer to a valid [VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepInfoNV.html) structure
- [](#VUID-vkLatencySleepNV-swapchain-parent)VUID-vkLatencySleepNV-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepInfoNV.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkLatencySleepNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
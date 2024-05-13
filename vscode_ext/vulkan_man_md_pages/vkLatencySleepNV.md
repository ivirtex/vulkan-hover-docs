# vkLatencySleepNV(3) Manual Page

## Name

vkLatencySleepNV - Trigger low latency mode Sleep



## <a href="#_c_specification" class="anchor"></a>C Specification

To provide the synchronization primitive used to delay host CPU work for
lower latency rendering, call:

``` c
// Provided by VK_NV_low_latency2
VkResult vkLatencySleepNV(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkLatencySleepInfoNV*                 pSleepInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to delay associated CPU work based on
  [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySubmissionPresentIdNV.html)
  submissions.

- `pSleepInfo` is a pointer to a
  [VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepInfoNV.html) structure specifying
  the parameters of the latency sleep.

## <a href="#_description" class="anchor"></a>Description

`vkLatencySleepNV` returns immediately. Applications **should** use
[vkWaitSemaphores](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphores.html) with
`pSleepInfo->signalSemaphore` to delay host CPU work. CPU work refers to
application work done before presenting which includes but is not
limited to: input sampling, simulation, command buffer recording,
command buffer submission, and present submission. It is recommended to
call this function before input sampling. When using this function, it
**should** be called exactly once between presents.

Valid Usage (Implicit)

- <a href="#VUID-vkLatencySleepNV-device-parameter"
  id="VUID-vkLatencySleepNV-device-parameter"></a>
  VUID-vkLatencySleepNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkLatencySleepNV-swapchain-parameter"
  id="VUID-vkLatencySleepNV-swapchain-parameter"></a>
  VUID-vkLatencySleepNV-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkLatencySleepNV-pSleepInfo-parameter"
  id="VUID-vkLatencySleepNV-pSleepInfo-parameter"></a>
  VUID-vkLatencySleepNV-pSleepInfo-parameter  
  `pSleepInfo` **must** be a valid pointer to a valid
  [VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepInfoNV.html) structure

- <a href="#VUID-vkLatencySleepNV-swapchain-parent"
  id="VUID-vkLatencySleepNV-swapchain-parent"></a>
  VUID-vkLatencySleepNV-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

This command does not return any failure codes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepInfoNV.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkLatencySleepNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

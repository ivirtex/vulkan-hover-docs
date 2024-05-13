# vkAcquireProfilingLockKHR(3) Manual Page

## Name

vkAcquireProfilingLockKHR - Acquires the profiling lock



## <a href="#_c_specification" class="anchor"></a>C Specification

To record and submit a command buffer containing a performance query
pool the profiling lock **must** be held. The profiling lock **must** be
acquired prior to any call to
[vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html) that will be using a
performance query pool. The profiling lock **must** be held while any
command buffer containing a performance query pool is in the
*recording*, *executable*, or *pending state*. To acquire the profiling
lock, call:

``` c
// Provided by VK_KHR_performance_query
VkResult vkAcquireProfilingLockKHR(
    VkDevice                                    device,
    const VkAcquireProfilingLockInfoKHR*        pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device to profile.

- `pInfo` is a pointer to a `VkAcquireProfilingLockInfoKHR` structure
  containing information about how the profiling is to be acquired.

## <a href="#_description" class="anchor"></a>Description

Implementations **may** allow multiple actors to hold the profiling lock
concurrently.

Valid Usage (Implicit)

- <a href="#VUID-vkAcquireProfilingLockKHR-device-parameter"
  id="VUID-vkAcquireProfilingLockKHR-device-parameter"></a>
  VUID-vkAcquireProfilingLockKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkAcquireProfilingLockKHR-pInfo-parameter"
  id="VUID-vkAcquireProfilingLockKHR-pInfo-parameter"></a>
  VUID-vkAcquireProfilingLockKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkAcquireProfilingLockInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireProfilingLockInfoKHR.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_TIMEOUT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkAcquireProfilingLockInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireProfilingLockInfoKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAcquireProfilingLockKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

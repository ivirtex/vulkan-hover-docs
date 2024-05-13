# vkReleaseProfilingLockKHR(3) Manual Page

## Name

vkReleaseProfilingLockKHR - Releases the profiling lock



## <a href="#_c_specification" class="anchor"></a>C Specification

To release the profiling lock, call:

``` c
// Provided by VK_KHR_performance_query
void vkReleaseProfilingLockKHR(
    VkDevice                                    device);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device to cease profiling on.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkReleaseProfilingLockKHR-device-03235"
  id="VUID-vkReleaseProfilingLockKHR-device-03235"></a>
  VUID-vkReleaseProfilingLockKHR-device-03235  
  The profiling lock of `device` **must** have been held via a previous
  successful call to
  [vkAcquireProfilingLockKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireProfilingLockKHR.html)

Valid Usage (Implicit)

- <a href="#VUID-vkReleaseProfilingLockKHR-device-parameter"
  id="VUID-vkReleaseProfilingLockKHR-device-parameter"></a>
  VUID-vkReleaseProfilingLockKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkReleaseProfilingLockKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

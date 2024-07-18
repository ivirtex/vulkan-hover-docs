# VkFence(3) Manual Page

## Name

VkFence - Opaque handle to a fence object



## <a href="#_c_specification" class="anchor"></a>C Specification

Fences are a synchronization primitive that **can** be used to insert a
dependency from a queue to the host. Fences have two states - signaled
and unsignaled. A fence **can** be signaled as part of the execution of
a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-submission"
target="_blank" rel="noopener">queue submission</a> command. Fences
**can** be unsignaled on the host with
[vkResetFences](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetFences.html). Fences **can** be waited on by the
host with the [vkWaitForFences](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitForFences.html) command, and the
current state **can** be queried with
[vkGetFenceStatus](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFenceStatus.html).

The internal data of a fence **may** include a reference to any
resources and pending work associated with signal or unsignal operations
performed on that fence object, collectively referred to as the fence’s
*payload*. Mechanisms to import and export that internal data to and
from fences are provided
<a href="VkExportFenceCreateInfo.html" target="_blank"
rel="noopener">below</a>. These mechanisms indirectly enable
applications to share fence state between two or more fences and other
synchronization primitives across process and API boundaries.

Fences are represented by `VkFence` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkFence)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_DEFINE_NON_DISPATCHABLE_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html),
[VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetFdInfoKHR.html),
[VkFenceGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetWin32HandleInfoKHR.html),
[VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceFdInfoKHR.html),
[VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html),
[VkSwapchainPresentFenceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentFenceInfoEXT.html),
[vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html),
[vkCreateFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateFence.html),
[vkDestroyFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyFence.html),
[vkGetFenceStatus](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFenceStatus.html),
[vkQueueBindSparse](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueBindSparse.html),
[vkQueueSubmit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html),
[vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit2.html),
[vkQueueSubmit2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit2KHR.html),
[vkRegisterDeviceEventEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkRegisterDeviceEventEXT.html),
[vkRegisterDisplayEventEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkRegisterDisplayEventEXT.html),
[vkResetFences](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetFences.html),
[vkWaitForFences](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitForFences.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFence"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

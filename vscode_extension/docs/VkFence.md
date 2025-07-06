# VkFence(3) Manual Page

## Name

VkFence - Opaque handle to a fence object



## [](#_c_specification)C Specification

Fences are a synchronization primitive that **can** be used to insert a dependency from a queue to the host. Fences have two states - signaled and unsignaled. A fence **can** be signaled as part of the execution of a [queue submission](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission) command. Fences **can** be unsignaled on the host with [vkResetFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetFences.html). Fences **can** be waited on by the host with the [vkWaitForFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForFences.html) command, and the current state **can** be queried with [vkGetFenceStatus](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFenceStatus.html).

The internal data of a fence **may** include a reference to any resources and pending work associated with signal or unsignal operations performed on that fence object, collectively referred to as the fence’s *payload*. Mechanisms to import and export that internal data to and from fences are provided [below](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportFenceCreateInfo.html). These mechanisms indirectly enable applications to share fence state between two or more fences and other synchronization primitives across process and API boundaries.

Fences are represented by `VkFence` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkFence)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html), [VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceGetFdInfoKHR.html), [VkFenceGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceGetWin32HandleInfoKHR.html), [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceFdInfoKHR.html), [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceWin32HandleInfoKHR.html), [VkSwapchainPresentFenceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentFenceInfoKHR.html), [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html), [vkCreateFence](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateFence.html), [vkDestroyFence](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyFence.html), [vkGetFenceStatus](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFenceStatus.html), [vkQueueBindSparse](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueBindSparse.html), [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html), [vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2.html), [vkQueueSubmit2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2KHR.html), [vkRegisterDeviceEventEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkRegisterDeviceEventEXT.html), [vkRegisterDisplayEventEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkRegisterDisplayEventEXT.html), [vkResetFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetFences.html), [vkWaitForFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForFences.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFence)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkSemaphore(3) Manual Page

## Name

VkSemaphore - Opaque handle to a semaphore object



## [](#_c_specification)C Specification

Semaphores are a synchronization primitive that **can** be used to insert a dependency between queue operations or between a queue operation and the host. [Binary semaphores](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary) have two states - signaled and unsignaled. [Timeline semaphores](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary) have a strictly increasing 64-bit unsigned integer payload and are signaled with respect to a particular reference value. A semaphore **can** be signaled after execution of a queue operation is completed, and a queue operation **can** wait for a semaphore to become signaled before it begins execution. A timeline semaphore **can** additionally be signaled from the host with the [vkSignalSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSignalSemaphore.html) command and waited on from the host with the [vkWaitSemaphores](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitSemaphores.html) command.

The internal data of a semaphore **may** include a reference to any resources and pending work associated with signal or unsignal operations performed on that semaphore object, collectively referred to as the semaphore’s *payload*. Mechanisms to import and export that internal data to and from semaphores are provided [below](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfo.html). These mechanisms indirectly enable applications to share semaphore state between two or more semaphores and other synchronization primitives across process and API boundaries.

Semaphores are represented by `VkSemaphore` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSemaphore)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html), [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html), [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalSharedEventInfoEXT.html), [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreFdInfoKHR.html), [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreWin32HandleInfoKHR.html), [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html), [VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepInfoNV.html), [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html), [VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetFdInfoKHR.html), [VkSemaphoreGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetWin32HandleInfoKHR.html), [VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html), [VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSignalInfo.html), [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html), [VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitInfo.html), [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html), [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html), [vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSemaphore.html), [vkDestroySemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySemaphore.html), [vkGetSemaphoreCounterValue](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreCounterValue.html), [vkGetSemaphoreCounterValue](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreCounterValue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphore).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
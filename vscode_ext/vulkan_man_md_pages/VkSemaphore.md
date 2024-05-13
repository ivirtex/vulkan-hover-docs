# VkSemaphore(3) Manual Page

## Name

VkSemaphore - Opaque handle to a semaphore object



## <a href="#_c_specification" class="anchor"></a>C Specification

Semaphores are a synchronization primitive that **can** be used to
insert a dependency between queue operations or between a queue
operation and the host. <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
target="_blank" rel="noopener">Binary semaphores</a> have two states -
signaled and unsignaled. <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
target="_blank" rel="noopener">Timeline semaphores</a> have a strictly
increasing 64-bit unsigned integer payload and are signaled with respect
to a particular reference value. A semaphore **can** be signaled after
execution of a queue operation is completed, and a queue operation
**can** wait for a semaphore to become signaled before it begins
execution. A timeline semaphore **can** additionally be signaled from
the host with the [vkSignalSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSignalSemaphore.html) command
and waited on from the host with the
[vkWaitSemaphores](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphores.html) command.

The internal data of a semaphore **may** include a reference to any
resources and pending work associated with signal or unsignal operations
performed on that semaphore object, collectively referred to as the
semaphore’s *payload*. Mechanisms to import and export that internal
data to and from semaphores are provided
<a href="VkExportSemaphoreCreateInfo.html" target="_blank"
rel="noopener">below</a>. These mechanisms indirectly enable
applications to share semaphore state between two or more semaphores and
other synchronization primitives across process and API boundaries.

Semaphores are represented by `VkSemaphore` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSemaphore)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html),
[VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html),
[VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalSharedEventInfoEXT.html),
[VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreFdInfoKHR.html),
[VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html),
[VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html),
[VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepInfoNV.html),
[VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html),
[VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetFdInfoKHR.html),
[VkSemaphoreGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetWin32HandleInfoKHR.html),
[VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html),
[VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSignalInfo.html),
[VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html),
[VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitInfo.html),
[VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html),
[vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html),
[vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSemaphore.html),
[vkDestroySemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySemaphore.html),
[vkGetSemaphoreCounterValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreCounterValue.html),
[vkGetSemaphoreCounterValueKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreCounterValueKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphore"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

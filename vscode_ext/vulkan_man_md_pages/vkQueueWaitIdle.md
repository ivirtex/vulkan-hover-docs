# vkQueueWaitIdle(3) Manual Page

## Name

vkQueueWaitIdle - Wait for a queue to become idle



## <a href="#_c_specification" class="anchor"></a>C Specification

To wait on the host for the completion of outstanding queue operations
for a given queue, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkQueueWaitIdle(
    VkQueue                                     queue);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the queue on which to wait.

## <a href="#_description" class="anchor"></a>Description

`vkQueueWaitIdle` is equivalent to having submitted a valid fence to
every previously executed <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-submission"
target="_blank" rel="noopener">queue submission command</a> that accepts
a fence, then waiting for all of those fences to signal using
[vkWaitForFences](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitForFences.html) with an infinite timeout and
`waitAll` set to `VK_TRUE`.

Valid Usage (Implicit)

- <a href="#VUID-vkQueueWaitIdle-queue-parameter"
  id="VUID-vkQueueWaitIdle-queue-parameter"></a>
  VUID-vkQueueWaitIdle-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

Host Synchronization

- Host access to `queue` **must** be externally synchronized

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|------------------------------------------------|--------------------------------------------|-------------------------------------------------|-------------------------------------------|------------------------------------------------------------|
| \-                                             | \-                                         | \-                                              | Any                                       | \-                                                         |

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueueWaitIdle"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

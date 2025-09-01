# vkQueueEndDebugUtilsLabelEXT(3) Manual Page

## Name

vkQueueEndDebugUtilsLabelEXT - Close a queue debug label region



## [](#_c_specification)C Specification

A queue debug label region is closed by calling:

```c++
// Provided by VK_EXT_debug_utils
void vkQueueEndDebugUtilsLabelEXT(
    VkQueue                                     queue);
```

## [](#_parameters)Parameters

- `queue` is the queue in which a debug label region should be closed.

## [](#_description)Description

The calls to [vkQueueBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueBeginDebugUtilsLabelEXT.html) and [vkQueueEndDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueEndDebugUtilsLabelEXT.html) **must** be matched and balanced.

Valid Usage

- [](#VUID-vkQueueEndDebugUtilsLabelEXT-None-01911)VUID-vkQueueEndDebugUtilsLabelEXT-None-01911  
  There **must** be an outstanding `vkQueueBeginDebugUtilsLabelEXT` command prior to the `vkQueueEndDebugUtilsLabelEXT` on the queue

Valid Usage (Implicit)

- [](#VUID-vkQueueEndDebugUtilsLabelEXT-queue-parameter)VUID-vkQueueEndDebugUtilsLabelEXT-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle

Host Synchronization

- Host access to `queue` **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

\-

\-

\-

Any

\-

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkQueueEndDebugUtilsLabelEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# vkQueueBeginDebugUtilsLabelEXT(3) Manual Page

## Name

vkQueueBeginDebugUtilsLabelEXT - Open a queue debug label region



## [](#_c_specification)C Specification

A queue debug label region is opened by calling:

```c++
// Provided by VK_EXT_debug_utils
void vkQueueBeginDebugUtilsLabelEXT(
    VkQueue                                     queue,
    const VkDebugUtilsLabelEXT*                 pLabelInfo);
```

## [](#_parameters)Parameters

- `queue` is the queue in which to start a debug label region.
- `pLabelInfo` is a pointer to a [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsLabelEXT.html) structure specifying parameters of the label region to open.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkQueueBeginDebugUtilsLabelEXT-queue-parameter)VUID-vkQueueBeginDebugUtilsLabelEXT-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle
- [](#VUID-vkQueueBeginDebugUtilsLabelEXT-pLabelInfo-parameter)VUID-vkQueueBeginDebugUtilsLabelEXT-pLabelInfo-parameter  
  `pLabelInfo` **must** be a valid pointer to a valid [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsLabelEXT.html) structure

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

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsLabelEXT.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkQueueBeginDebugUtilsLabelEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
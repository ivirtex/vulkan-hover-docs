# vkQueueEndDebugUtilsLabelEXT(3) Manual Page

## Name

vkQueueEndDebugUtilsLabelEXT - Close a queue debug label region



## <a href="#_c_specification" class="anchor"></a>C Specification

A queue debug label region is closed by calling:

``` c
// Provided by VK_EXT_debug_utils
void vkQueueEndDebugUtilsLabelEXT(
    VkQueue                                     queue);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the queue in which a debug label region should be closed.

## <a href="#_description" class="anchor"></a>Description

The calls to
[vkQueueBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueBeginDebugUtilsLabelEXT.html)
and [vkQueueEndDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueEndDebugUtilsLabelEXT.html)
**must** be matched and balanced.

Valid Usage

- <a href="#VUID-vkQueueEndDebugUtilsLabelEXT-None-01911"
  id="VUID-vkQueueEndDebugUtilsLabelEXT-None-01911"></a>
  VUID-vkQueueEndDebugUtilsLabelEXT-None-01911  
  There **must** be an outstanding `vkQueueBeginDebugUtilsLabelEXT`
  command prior to the `vkQueueEndDebugUtilsLabelEXT` on the queue

Valid Usage (Implicit)

- <a href="#VUID-vkQueueEndDebugUtilsLabelEXT-queue-parameter"
  id="VUID-vkQueueEndDebugUtilsLabelEXT-queue-parameter"></a>
  VUID-vkQueueEndDebugUtilsLabelEXT-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|----|----|----|----|----|
| \- | \- | \- | Any | \- |

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html), [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueueEndDebugUtilsLabelEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

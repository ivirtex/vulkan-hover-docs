# vkQueueBeginDebugUtilsLabelEXT(3) Manual Page

## Name

vkQueueBeginDebugUtilsLabelEXT - Open a queue debug label region



## <a href="#_c_specification" class="anchor"></a>C Specification

A queue debug label region is opened by calling:

``` c
// Provided by VK_EXT_debug_utils
void vkQueueBeginDebugUtilsLabelEXT(
    VkQueue                                     queue,
    const VkDebugUtilsLabelEXT*                 pLabelInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the queue in which to start a debug label region.

- `pLabelInfo` is a pointer to a
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) structure specifying
  parameters of the label region to open.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkQueueBeginDebugUtilsLabelEXT-queue-parameter"
  id="VUID-vkQueueBeginDebugUtilsLabelEXT-queue-parameter"></a>
  VUID-vkQueueBeginDebugUtilsLabelEXT-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a href="#VUID-vkQueueBeginDebugUtilsLabelEXT-pLabelInfo-parameter"
  id="VUID-vkQueueBeginDebugUtilsLabelEXT-pLabelInfo-parameter"></a>
  VUID-vkQueueBeginDebugUtilsLabelEXT-pLabelInfo-parameter  
  `pLabelInfo` **must** be a valid pointer to a valid
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) structure

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|----|----|----|----|----|
| \- | \- | \- | Any | \- |

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html),
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueueBeginDebugUtilsLabelEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

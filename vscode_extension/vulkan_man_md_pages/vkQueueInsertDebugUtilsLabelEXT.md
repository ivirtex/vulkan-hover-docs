# vkQueueInsertDebugUtilsLabelEXT(3) Manual Page

## Name

vkQueueInsertDebugUtilsLabelEXT - Insert a label into a queue



## <a href="#_c_specification" class="anchor"></a>C Specification

A single label can be inserted into a queue by calling:

``` c
// Provided by VK_EXT_debug_utils
void vkQueueInsertDebugUtilsLabelEXT(
    VkQueue                                     queue,
    const VkDebugUtilsLabelEXT*                 pLabelInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the queue into which a debug label will be inserted.

- `pLabelInfo` is a pointer to a
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) structure specifying
  parameters of the label to insert.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkQueueInsertDebugUtilsLabelEXT-queue-parameter"
  id="VUID-vkQueueInsertDebugUtilsLabelEXT-queue-parameter"></a>
  VUID-vkQueueInsertDebugUtilsLabelEXT-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a href="#VUID-vkQueueInsertDebugUtilsLabelEXT-pLabelInfo-parameter"
  id="VUID-vkQueueInsertDebugUtilsLabelEXT-pLabelInfo-parameter"></a>
  VUID-vkQueueInsertDebugUtilsLabelEXT-pLabelInfo-parameter  
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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueueInsertDebugUtilsLabelEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

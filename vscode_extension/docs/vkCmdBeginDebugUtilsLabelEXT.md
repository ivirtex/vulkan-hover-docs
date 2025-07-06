# vkCmdBeginDebugUtilsLabelEXT(3) Manual Page

## Name

vkCmdBeginDebugUtilsLabelEXT - Open a command buffer debug label region



## [](#_c_specification)C Specification

A command buffer debug label region can be opened by calling:

```c++
// Provided by VK_EXT_debug_utils
void vkCmdBeginDebugUtilsLabelEXT(
    VkCommandBuffer                             commandBuffer,
    const VkDebugUtilsLabelEXT*                 pLabelInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `pLabelInfo` is a pointer to a [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsLabelEXT.html) structure specifying parameters of the label region to open.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginDebugUtilsLabelEXT-commandBuffer-parameter)VUID-vkCmdBeginDebugUtilsLabelEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginDebugUtilsLabelEXT-pLabelInfo-parameter)VUID-vkCmdBeginDebugUtilsLabelEXT-pLabelInfo-parameter  
  `pLabelInfo` **must** be a valid pointer to a valid [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsLabelEXT.html) structure
- [](#VUID-vkCmdBeginDebugUtilsLabelEXT-commandBuffer-recording)VUID-vkCmdBeginDebugUtilsLabelEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginDebugUtilsLabelEXT-commandBuffer-cmdpool)VUID-vkCmdBeginDebugUtilsLabelEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, compute, decode, encode, or optical flow operations

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Both

Transfer  
Graphics  
Compute  
Decode  
Encode  
Opticalflow

Action  
State

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsLabelEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginDebugUtilsLabelEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
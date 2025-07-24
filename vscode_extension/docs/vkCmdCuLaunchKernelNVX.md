# vkCmdCuLaunchKernelNVX(3) Manual Page

## Name

vkCmdCuLaunchKernelNVX - Stub description of vkCmdCuLaunchKernelNVX



## [](#_c_specification)C Specification

There is currently no specification language written for this command. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_NVX_binary_import
void vkCmdCuLaunchKernelNVX(
    VkCommandBuffer                             commandBuffer,
    const VkCuLaunchInfoNVX*                    pLaunchInfo);
```

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCmdCuLaunchKernelNVX-commandBuffer-parameter)VUID-vkCmdCuLaunchKernelNVX-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCuLaunchKernelNVX-pLaunchInfo-parameter)VUID-vkCmdCuLaunchKernelNVX-pLaunchInfo-parameter  
  `pLaunchInfo` **must** be a valid pointer to a valid [VkCuLaunchInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuLaunchInfoNVX.html) structure
- [](#VUID-vkCmdCuLaunchKernelNVX-commandBuffer-recording)VUID-vkCmdCuLaunchKernelNVX-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCuLaunchKernelNVX-commandBuffer-cmdpool)VUID-vkCmdCuLaunchKernelNVX-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdCuLaunchKernelNVX-videocoding)VUID-vkCmdCuLaunchKernelNVX-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics  
Compute

Action

Conditional Rendering

vkCmdCuLaunchKernelNVX is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NVX\_binary\_import](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_binary_import.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCuLaunchInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuLaunchInfoNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCuLaunchKernelNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
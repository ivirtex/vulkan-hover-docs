# vkCmdResolveImage2(3) Manual Page

## Name

vkCmdResolveImage2 - Resolve regions of an image



## [](#_c_specification)C Specification

To resolve a multisample image to a non-multisample image, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdResolveImage2(
    VkCommandBuffer                             commandBuffer,
    const VkResolveImageInfo2*                  pResolveImageInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_copy_commands2
void vkCmdResolveImage2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkResolveImageInfo2*                  pResolveImageInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pResolveImageInfo` is a pointer to a [VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveImageInfo2.html) structure describing the resolve parameters.

## [](#_description)Description

This command is functionally identical to [vkCmdResolveImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResolveImage.html), but includes extensible sub-structures that include `sType` and `pNext` parameters, allowing them to be more easily extended.

Valid Usage

- [](#VUID-vkCmdResolveImage2-commandBuffer-01837)VUID-vkCmdResolveImage2-commandBuffer-01837  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `srcImage` **must** not be a protected image
- [](#VUID-vkCmdResolveImage2-commandBuffer-01838)VUID-vkCmdResolveImage2-commandBuffer-01838  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be a protected image
- [](#VUID-vkCmdResolveImage2-commandBuffer-01839)VUID-vkCmdResolveImage2-commandBuffer-01839  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be an unprotected image

Valid Usage (Implicit)

- [](#VUID-vkCmdResolveImage2-commandBuffer-parameter)VUID-vkCmdResolveImage2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdResolveImage2-pResolveImageInfo-parameter)VUID-vkCmdResolveImage2-pResolveImageInfo-parameter  
  `pResolveImageInfo` **must** be a valid pointer to a valid [VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveImageInfo2.html) structure
- [](#VUID-vkCmdResolveImage2-commandBuffer-recording)VUID-vkCmdResolveImage2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdResolveImage2-commandBuffer-cmdpool)VUID-vkCmdResolveImage2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdResolveImage2-renderpass)VUID-vkCmdResolveImage2-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdResolveImage2-videocoding)VUID-vkCmdResolveImage2-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Graphics

Action

Conditional Rendering

vkCmdResolveImage2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveImageInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdResolveImage2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
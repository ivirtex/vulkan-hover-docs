# vkCmdEndVideoCodingKHR(3) Manual Page

## Name

vkCmdEndVideoCodingKHR - End video coding scope



## [](#_c_specification)C Specification

To end a video coding scope, call:

```c++
// Provided by VK_KHR_video_queue
void vkCmdEndVideoCodingKHR(
    VkCommandBuffer                             commandBuffer,
    const VkVideoEndCodingInfoKHR*              pEndCodingInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pEndCodingInfo` is a pointer to a [VkVideoEndCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEndCodingInfoKHR.html) structure specifying the parameters for ending the video coding scope.

## [](#_description)Description

After ending a video coding scope, the video session object, the optional video session parameters object, and all [reference picture resources](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#bound-reference-picture-resources) previously bound by the corresponding [vkCmdBeginVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginVideoCodingKHR.html) command are *unbound*.

Valid Usage

- [](#VUID-vkCmdEndVideoCodingKHR-None-07251)VUID-vkCmdEndVideoCodingKHR-None-07251  
  There **must** be no [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-active) queries

Valid Usage (Implicit)

- [](#VUID-vkCmdEndVideoCodingKHR-commandBuffer-parameter)VUID-vkCmdEndVideoCodingKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndVideoCodingKHR-pEndCodingInfo-parameter)VUID-vkCmdEndVideoCodingKHR-pEndCodingInfo-parameter  
  `pEndCodingInfo` **must** be a valid pointer to a valid [VkVideoEndCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEndCodingInfoKHR.html) structure
- [](#VUID-vkCmdEndVideoCodingKHR-commandBuffer-recording)VUID-vkCmdEndVideoCodingKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndVideoCodingKHR-commandBuffer-cmdpool)VUID-vkCmdEndVideoCodingKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR, or VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR operations
- [](#VUID-vkCmdEndVideoCodingKHR-renderpass)VUID-vkCmdEndVideoCodingKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdEndVideoCodingKHR-videocoding)VUID-vkCmdEndVideoCodingKHR-videocoding  
  This command **must** only be called inside of a video coding scope
- [](#VUID-vkCmdEndVideoCodingKHR-bufferlevel)VUID-vkCmdEndVideoCodingKHR-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary

Outside

Inside

VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR  
VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR

Action  
State

Conditional Rendering

vkCmdEndVideoCodingKHR is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkVideoEndCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEndCodingInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndVideoCodingKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
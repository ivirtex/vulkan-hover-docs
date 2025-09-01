# vkCmdControlVideoCodingKHR(3) Manual Page

## Name

vkCmdControlVideoCodingKHR - Control video coding parameters



## [](#_c_specification)C Specification

To apply dynamic controls to the bound video session object, call:

```c++
// Provided by VK_KHR_video_queue
void vkCmdControlVideoCodingKHR(
    VkCommandBuffer                             commandBuffer,
    const VkVideoCodingControlInfoKHR*          pCodingControlInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pCodingControlInfo` is a pointer to a [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html) structure specifying the control parameters.

## [](#_description)Description

The control parameters provided in this call are applied to the video session at the time the command executes on the device and are in effect until a subsequent call to this command with the same video session bound changes the corresponding control parameters.

A newly created video session **must** be reset before performing video coding operations using it by including `VK_VIDEO_CODING_CONTROL_RESET_BIT_KHR` in `pCodingControlInfo->flags`. The reset operation also returns all DPB slots of the video session to the [inactive state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states). Correspondingly, any DPB slot index associated with the [bound reference picture resources](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#bound-reference-picture-resources) is removed.

For encode sessions, the reset operation returns [rate control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control) configuration to implementation default settings and sets the [video encode quality level](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quality-level) to zero.

After video coding operations are performed using a video session, the reset operation **can** be used to return the video session to the same *initial* state as after the reset of a newly created video session. This **can** be used, for example, when different video sequences are needed to be processed with the same video session object.

If `pCodingControlInfo->flags` includes `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, then the command replaces the [rate control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control) configuration maintained by the video session with the configuration specified in the [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlInfoKHR.html) structure included in the `pCodingControlInfo->pNext` chain.

If `pCodingControlInfo->flags` includes `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR`, then the command changes the current [video encode quality level](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quality-level) to the value specified in the `qualityLevel` member of the [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html) structure included in the `pCodingControlInfo->pNext` chain.

Valid Usage

- [](#VUID-vkCmdControlVideoCodingKHR-flags-07017)VUID-vkCmdControlVideoCodingKHR-flags-07017  
  If `pCodingControlInfo->flags` does not include `VK_VIDEO_CODING_CONTROL_RESET_BIT_KHR`, then the bound video session **must** not be in [uninitialized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-uninitialized) state at the time the command is executed on the device
- [](#VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-08243)VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-08243  
  If the bound video session was not created with an encode operation, then `pCodingControlInfo->flags` **must** not include `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR` or `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-vkCmdControlVideoCodingKHR-commandBuffer-parameter)VUID-vkCmdControlVideoCodingKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-parameter)VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-parameter  
  `pCodingControlInfo` **must** be a valid pointer to a valid [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html) structure
- [](#VUID-vkCmdControlVideoCodingKHR-commandBuffer-recording)VUID-vkCmdControlVideoCodingKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdControlVideoCodingKHR-commandBuffer-cmdpool)VUID-vkCmdControlVideoCodingKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support decode, or encode operations
- [](#VUID-vkCmdControlVideoCodingKHR-renderpass)VUID-vkCmdControlVideoCodingKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdControlVideoCodingKHR-videocoding)VUID-vkCmdControlVideoCodingKHR-videocoding  
  This command **must** only be called inside of a video coding scope
- [](#VUID-vkCmdControlVideoCodingKHR-bufferlevel)VUID-vkCmdControlVideoCodingKHR-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary

Outside

Inside

Decode  
Encode

Action

Conditional Rendering

vkCmdControlVideoCodingKHR is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdControlVideoCodingKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
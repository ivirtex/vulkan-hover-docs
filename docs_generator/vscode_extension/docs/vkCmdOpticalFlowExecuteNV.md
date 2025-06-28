# vkCmdOpticalFlowExecuteNV(3) Manual Page

## Name

vkCmdOpticalFlowExecuteNV - Calculate optical flow vectors



## [](#_c_specification)C Specification

Default direction of flow estimation is forward which calculates the optical flow from input frame to reference frame. Optionally backward flow estimation can be additionally calculated. An output flow vector (Vx, Vy) means that current pixel (x, y) of input frame can be found at location (x+Vx, y+Vy) in reference frame. A backward flow vector (Vx, Vy) means that current pixel (x, y) of reference frame can be found at location (x+Vx, y+Vy) in input frame.

To calculate optical flow vectors from two input frames, call:

```c++
// Provided by VK_NV_optical_flow
void vkCmdOpticalFlowExecuteNV(
    VkCommandBuffer                             commandBuffer,
    VkOpticalFlowSessionNV                      session,
    const VkOpticalFlowExecuteInfoNV*           pExecuteInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `session` is the optical flow session object on which this command is operating.
- `pExecuteInfo` Info is a pointer to a [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteInfoNV.html).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-parameter)VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdOpticalFlowExecuteNV-session-parameter)VUID-vkCmdOpticalFlowExecuteNV-session-parameter  
  `session` **must** be a valid [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html) handle
- [](#VUID-vkCmdOpticalFlowExecuteNV-pExecuteInfo-parameter)VUID-vkCmdOpticalFlowExecuteNV-pExecuteInfo-parameter  
  `pExecuteInfo` **must** be a valid pointer to a valid [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteInfoNV.html) structure
- [](#VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-recording)VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-cmdpool)VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support optical flow operations
- [](#VUID-vkCmdOpticalFlowExecuteNV-renderpass)VUID-vkCmdOpticalFlowExecuteNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdOpticalFlowExecuteNV-videocoding)VUID-vkCmdOpticalFlowExecuteNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdOpticalFlowExecuteNV-commonparent)VUID-vkCmdOpticalFlowExecuteNV-commonparent  
  Both of `commandBuffer`, and `session` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Opticalflow

Action

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteInfoNV.html), [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdOpticalFlowExecuteNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
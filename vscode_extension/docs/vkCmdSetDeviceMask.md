# vkCmdSetDeviceMask(3) Manual Page

## Name

vkCmdSetDeviceMask - Modify device mask of a command buffer



## [](#_c_specification)C Specification

To update the current device mask of a command buffer, call:

```c++
// Provided by VK_VERSION_1_1
void vkCmdSetDeviceMask(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    deviceMask);
```

or the equivalent command

```c++
// Provided by VK_KHR_device_group
void vkCmdSetDeviceMaskKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    deviceMask);
```

## [](#_parameters)Parameters

- `commandBuffer` is command buffer whose current device mask is modified.
- `deviceMask` is the new value of the current device mask.

## [](#_description)Description

`deviceMask` is used to filter out subsequent commands from executing on all physical devices whose bit indices are not set in the mask, except commands beginning a render pass instance, commands transitioning to the next subpass in the render pass instance, and commands ending a render pass instance, which always execute on the set of physical devices whose bit indices are included in the `deviceMask` member of the [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html) structure passed to the command beginning the corresponding render pass instance.

Valid Usage

- [](#VUID-vkCmdSetDeviceMask-deviceMask-00108)VUID-vkCmdSetDeviceMask-deviceMask-00108  
  `deviceMask` **must** be a valid device mask value
- [](#VUID-vkCmdSetDeviceMask-deviceMask-00109)VUID-vkCmdSetDeviceMask-deviceMask-00109  
  `deviceMask` **must** not be zero
- [](#VUID-vkCmdSetDeviceMask-deviceMask-00110)VUID-vkCmdSetDeviceMask-deviceMask-00110  
  `deviceMask` **must** not include any set bits that were not in the [VkDeviceGroupCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupCommandBufferBeginInfo.html)::`deviceMask` value when the command buffer began recording
- [](#VUID-vkCmdSetDeviceMask-deviceMask-00111)VUID-vkCmdSetDeviceMask-deviceMask-00111  
  If `vkCmdSetDeviceMask` is called inside a render pass instance, `deviceMask` **must** not include any set bits that were not in the [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html)::`deviceMask` value when the render pass instance began recording

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDeviceMask-commandBuffer-parameter)VUID-vkCmdSetDeviceMask-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDeviceMask-commandBuffer-recording)VUID-vkCmdSetDeviceMask-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDeviceMask-commandBuffer-cmdpool)VUID-vkCmdSetDeviceMask-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, or VK\_QUEUE\_TRANSFER\_BIT operations

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Both

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_TRANSFER\_BIT

State

Conditional Rendering

vkCmdSetDeviceMask is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDeviceMask).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
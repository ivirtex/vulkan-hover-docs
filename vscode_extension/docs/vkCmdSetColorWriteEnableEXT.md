# vkCmdSetColorWriteEnableEXT(3) Manual Page

## Name

vkCmdSetColorWriteEnableEXT - Enable or disable writes to a color attachment dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically enable or disable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) writes to a color attachment, call:

```c++
// Provided by VK_EXT_color_write_enable
void                                    vkCmdSetColorWriteEnableEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    attachmentCount,
    const VkBool32*                             pColorWriteEnables);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `attachmentCount` is the number of [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) elements in `pColorWriteEnables`.
- `pColorWriteEnables` is a pointer to an array of per target attachment boolean values specifying whether color writes are enabled for the given attachment.

## [](#_description)Description

This command sets the color write enables for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineColorWriteCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorWriteCreateInfoEXT.html)::`pColorWriteEnables` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetColorWriteEnableEXT-None-04803)VUID-vkCmdSetColorWriteEnableEXT-None-04803  
  The [`colorWriteEnable`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-colorWriteEnable) feature **must** be enabled
- [](#VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-06656)VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-06656  
  `attachmentCount` **must** be less than or equal to the `maxColorAttachments` member of `VkPhysicalDeviceLimits`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-parameter)VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetColorWriteEnableEXT-pColorWriteEnables-parameter)VUID-vkCmdSetColorWriteEnableEXT-pColorWriteEnables-parameter  
  `pColorWriteEnables` **must** be a valid pointer to an array of `attachmentCount` [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) values
- [](#VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-recording)VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-cmdpool)VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetColorWriteEnableEXT-videocoding)VUID-vkCmdSetColorWriteEnableEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-arraylength)VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-arraylength  
  `attachmentCount` **must** be greater than `0`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics

State

Conditional Rendering

vkCmdSetColorWriteEnableEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_color\_write\_enable](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_color_write_enable.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetColorWriteEnableEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
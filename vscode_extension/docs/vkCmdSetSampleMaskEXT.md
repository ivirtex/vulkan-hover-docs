# vkCmdSetSampleMaskEXT(3) Manual Page

## Name

vkCmdSetSampleMaskEXT - Specify the sample mask dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the sample mask, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetSampleMaskEXT(
    VkCommandBuffer                             commandBuffer,
    VkSampleCountFlagBits                       samples,
    const VkSampleMask*                         pSampleMask);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `samples` specifies the number of sample bits in the `pSampleMask`.
- `pSampleMask` is a pointer to an array of [VkSampleMask](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleMask.html) values, where the array size is based on the `samples` parameter.

## [](#_description)Description

This command sets the sample mask for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html)::`pSampleMask` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetSampleMaskEXT-None-09423)VUID-vkCmdSetSampleMaskEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3SampleMask`](#features-extendedDynamicState3SampleMask) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetSampleMaskEXT-commandBuffer-parameter)VUID-vkCmdSetSampleMaskEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetSampleMaskEXT-samples-parameter)VUID-vkCmdSetSampleMaskEXT-samples-parameter  
  `samples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value
- [](#VUID-vkCmdSetSampleMaskEXT-pSampleMask-parameter)VUID-vkCmdSetSampleMaskEXT-pSampleMask-parameter  
  `pSampleMask` **must** be a valid pointer to an array of ⌈32samples​⌉ [VkSampleMask](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleMask.html) values
- [](#VUID-vkCmdSetSampleMaskEXT-commandBuffer-recording)VUID-vkCmdSetSampleMaskEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetSampleMaskEXT-commandBuffer-cmdpool)VUID-vkCmdSetSampleMaskEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetSampleMaskEXT-videocoding)VUID-vkCmdSetSampleMaskEXT-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdSetSampleMaskEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkSampleMask](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleMask.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetSampleMaskEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
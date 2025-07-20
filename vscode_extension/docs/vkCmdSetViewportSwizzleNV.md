# vkCmdSetViewportSwizzleNV(3) Manual Page

## Name

vkCmdSetViewportSwizzleNV - Specify the viewport swizzle state dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the viewport swizzle state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_viewport_swizzle, VK_EXT_shader_object with VK_NV_viewport_swizzle
void vkCmdSetViewportSwizzleNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstViewport,
    uint32_t                                    viewportCount,
    const VkViewportSwizzleNV*                  pViewportSwizzles);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstViewport` is the index of the first viewport whose parameters are updated by the command.
- `viewportCount` is the number of viewports whose parameters are updated by the command.
- `pViewportSwizzles` is a pointer to an array of [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html) structures specifying viewport swizzles.

## [](#_description)Description

This command sets the viewport swizzle state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)::`viewportCount`, and [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)::`pViewportSwizzles` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetViewportSwizzleNV-None-09423)VUID-vkCmdSetViewportSwizzleNV-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ViewportSwizzle`](#features-extendedDynamicState3ViewportSwizzle) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetViewportSwizzleNV-commandBuffer-parameter)VUID-vkCmdSetViewportSwizzleNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetViewportSwizzleNV-pViewportSwizzles-parameter)VUID-vkCmdSetViewportSwizzleNV-pViewportSwizzles-parameter  
  `pViewportSwizzles` **must** be a valid pointer to an array of `viewportCount` valid [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html) structures
- [](#VUID-vkCmdSetViewportSwizzleNV-commandBuffer-recording)VUID-vkCmdSetViewportSwizzleNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetViewportSwizzleNV-commandBuffer-cmdpool)VUID-vkCmdSetViewportSwizzleNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetViewportSwizzleNV-videocoding)VUID-vkCmdSetViewportSwizzleNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetViewportSwizzleNV-viewportCount-arraylength)VUID-vkCmdSetViewportSwizzleNV-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

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

vkCmdSetViewportSwizzleNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_NV\_viewport\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_viewport_swizzle.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetViewportSwizzleNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
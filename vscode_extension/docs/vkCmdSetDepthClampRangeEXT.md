# vkCmdSetDepthClampRangeEXT(3) Manual Page

## Name

vkCmdSetDepthClampRangeEXT - Set the viewport depth clamp range dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the viewport depth clamp range parameters, call:

```c++
// Provided by VK_EXT_depth_clamp_control, VK_EXT_depth_clamp_control with VK_EXT_shader_object
void vkCmdSetDepthClampRangeEXT(
    VkCommandBuffer                             commandBuffer,
    VkDepthClampModeEXT                         depthClampMode,
    const VkDepthClampRangeEXT*                 pDepthClampRange);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `depthClampMode` determines how the clamp range is determined for each viewport.
- `pDepthClampRange` sets the depth clamp range for all viewports if `depthClampMode` is `VK_DEPTH_CLAMP_MODE_USER_DEFINED_RANGE_EXT`.

## [](#_description)Description

This command sets the viewport depth clamp range for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_DEPTH_CLAMP_RANGE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportDepthClampControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClampControlCreateInfoEXT.html)::`depthClampMode` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetDepthClampRangeEXT-pDepthClampRange-09647)VUID-vkCmdSetDepthClampRangeEXT-pDepthClampRange-09647  
  If `depthClampMode` is `VK_DEPTH_CLAMP_MODE_USER_DEFINED_RANGE_EXT`, then `pDepthClampRange` must be a valid pointer to a valid `VkDepthClampRangeEXT` structure

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDepthClampRangeEXT-commandBuffer-parameter)VUID-vkCmdSetDepthClampRangeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDepthClampRangeEXT-depthClampMode-parameter)VUID-vkCmdSetDepthClampRangeEXT-depthClampMode-parameter  
  `depthClampMode` **must** be a valid [VkDepthClampModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampModeEXT.html) value
- [](#VUID-vkCmdSetDepthClampRangeEXT-pDepthClampRange-parameter)VUID-vkCmdSetDepthClampRangeEXT-pDepthClampRange-parameter  
  If `pDepthClampRange` is not `NULL`, `pDepthClampRange` **must** be a valid pointer to a valid [VkDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampRangeEXT.html) structure
- [](#VUID-vkCmdSetDepthClampRangeEXT-commandBuffer-recording)VUID-vkCmdSetDepthClampRangeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDepthClampRangeEXT-commandBuffer-cmdpool)VUID-vkCmdSetDepthClampRangeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetDepthClampRangeEXT-videocoding)VUID-vkCmdSetDepthClampRangeEXT-videocoding  
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

vkCmdSetDepthClampRangeEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_depth\_clamp\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clamp_control.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDepthClampModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampModeEXT.html), [VkDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampRangeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDepthClampRangeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
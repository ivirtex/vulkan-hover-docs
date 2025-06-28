# vkCmdSetPatchControlPointsEXT(3) Manual Page

## Name

vkCmdSetPatchControlPointsEXT - Specify the number of control points per patch dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the number of control points per patch, call:

```c++
// Provided by VK_EXT_extended_dynamic_state2, VK_EXT_shader_object
void vkCmdSetPatchControlPointsEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    patchControlPoints);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `patchControlPoints` specifies the number of control points per patch.

## [](#_description)Description

This command sets the number of control points per patch for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationStateCreateInfo.html)::`patchControlPoints` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetPatchControlPointsEXT-None-09422)VUID-vkCmdSetPatchControlPointsEXT-None-09422  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState2PatchControlPoints`](#features-extendedDynamicState2PatchControlPoints) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetPatchControlPointsEXT-patchControlPoints-04874)VUID-vkCmdSetPatchControlPointsEXT-patchControlPoints-04874  
  `patchControlPoints` **must** be greater than zero and less than or equal to `VkPhysicalDeviceLimits`::`maxTessellationPatchSize`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetPatchControlPointsEXT-commandBuffer-parameter)VUID-vkCmdSetPatchControlPointsEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetPatchControlPointsEXT-commandBuffer-recording)VUID-vkCmdSetPatchControlPointsEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetPatchControlPointsEXT-commandBuffer-cmdpool)VUID-vkCmdSetPatchControlPointsEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetPatchControlPointsEXT-videocoding)VUID-vkCmdSetPatchControlPointsEXT-videocoding  
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

Graphics

State

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state2.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetPatchControlPointsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
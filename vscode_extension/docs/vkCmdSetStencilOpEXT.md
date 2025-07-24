# vkCmdSetStencilOp(3) Manual Page

## Name

vkCmdSetStencilOp - Set stencil operation dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the stencil operation, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdSetStencilOp(
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    VkStencilOp                                 failOp,
    VkStencilOp                                 passOp,
    VkStencilOp                                 depthFailOp,
    VkCompareOp                                 compareOp);
```

or the equivalent command

```c++
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdSetStencilOpEXT(
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    VkStencilOp                                 failOp,
    VkStencilOp                                 passOp,
    VkStencilOp                                 depthFailOp,
    VkCompareOp                                 compareOp);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `faceMask` is a bitmask of [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlagBits.html) specifying the set of stencil state for which to update the stencil operation.
- `failOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value specifying the action performed on samples that fail the stencil test.
- `passOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value specifying the action performed on samples that pass both the depth and stencil tests.
- `depthFailOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value specifying the action performed on samples that pass the stencil test and fail the depth test.
- `compareOp` is a [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html) value specifying the comparison operator used in the stencil test.

## [](#_description)Description

This command sets the stencil operation for subsequent drawing commands when when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_STENCIL_OP` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the corresponding `VkPipelineDepthStencilStateCreateInfo`::`failOp`, `passOp`, `depthFailOp`, and `compareOp` values used to create the currently active pipeline, for both front and back faces.

Valid Usage

- [](#VUID-vkCmdSetStencilOp-None-08971)VUID-vkCmdSetStencilOp-None-08971  
  At least one of the following **must** be true:
  
  - the [`extendedDynamicState`](#features-extendedDynamicState) feature is enabled
  - the [`shaderObject`](#features-shaderObject) feature is enabled
  - the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) parent of `commandBuffer` is greater than or equal to Version 1.3

Valid Usage (Implicit)

- [](#VUID-vkCmdSetStencilOp-commandBuffer-parameter)VUID-vkCmdSetStencilOp-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetStencilOp-faceMask-parameter)VUID-vkCmdSetStencilOp-faceMask-parameter  
  `faceMask` **must** be a valid combination of [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlagBits.html) values
- [](#VUID-vkCmdSetStencilOp-faceMask-requiredbitmask)VUID-vkCmdSetStencilOp-faceMask-requiredbitmask  
  `faceMask` **must** not be `0`
- [](#VUID-vkCmdSetStencilOp-failOp-parameter)VUID-vkCmdSetStencilOp-failOp-parameter  
  `failOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value
- [](#VUID-vkCmdSetStencilOp-passOp-parameter)VUID-vkCmdSetStencilOp-passOp-parameter  
  `passOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value
- [](#VUID-vkCmdSetStencilOp-depthFailOp-parameter)VUID-vkCmdSetStencilOp-depthFailOp-parameter  
  `depthFailOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value
- [](#VUID-vkCmdSetStencilOp-compareOp-parameter)VUID-vkCmdSetStencilOp-compareOp-parameter  
  `compareOp` **must** be a valid [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html) value
- [](#VUID-vkCmdSetStencilOp-commandBuffer-recording)VUID-vkCmdSetStencilOp-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetStencilOp-commandBuffer-cmdpool)VUID-vkCmdSetStencilOp-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetStencilOp-videocoding)VUID-vkCmdSetStencilOp-videocoding  
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

Conditional Rendering

vkCmdSetStencilOp is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html), [VkStencilFaceFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlags.html), [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetStencilOp)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
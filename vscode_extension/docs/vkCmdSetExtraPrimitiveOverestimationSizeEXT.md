# vkCmdSetExtraPrimitiveOverestimationSizeEXT(3) Manual Page

## Name

vkCmdSetExtraPrimitiveOverestimationSizeEXT - Specify the conservative rasterization extra primitive overestimation size dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `extraPrimitiveOverestimationSize`, call:

```c++
// Provided by VK_EXT_conservative_rasterization with VK_EXT_extended_dynamic_state3, VK_EXT_conservative_rasterization with VK_EXT_shader_object
void vkCmdSetExtraPrimitiveOverestimationSizeEXT(
    VkCommandBuffer                             commandBuffer,
    float                                       extraPrimitiveOverestimationSize);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `extraPrimitiveOverestimationSize` specifies the `extraPrimitiveOverestimationSize`.

## [](#_description)Description

This command sets the `extraPrimitiveOverestimationSize` for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)::`extraPrimitiveOverestimationSize` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-None-09423)VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ExtraPrimitiveOverestimationSize`](#features-extendedDynamicState3ExtraPrimitiveOverestimationSize) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-extraPrimitiveOverestimationSize-07428)VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-extraPrimitiveOverestimationSize-07428  
  `extraPrimitiveOverestimationSize` **must** be in the range of `0.0` to `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`maxExtraPrimitiveOverestimationSize` inclusive

Valid Usage (Implicit)

- [](#VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-commandBuffer-parameter)VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-commandBuffer-recording)VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-commandBuffer-cmdpool)VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-videocoding)VUID-vkCmdSetExtraPrimitiveOverestimationSizeEXT-videocoding  
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

vkCmdSetExtraPrimitiveOverestimationSizeEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_conservative\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conservative_rasterization.html), [VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetExtraPrimitiveOverestimationSizeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
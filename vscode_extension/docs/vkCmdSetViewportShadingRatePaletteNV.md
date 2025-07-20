# vkCmdSetViewportShadingRatePaletteNV(3) Manual Page

## Name

vkCmdSetViewportShadingRatePaletteNV - Set shading rate image palettes dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the per-viewport shading rate image palettes, call:

```c++
// Provided by VK_NV_shading_rate_image
void vkCmdSetViewportShadingRatePaletteNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstViewport,
    uint32_t                                    viewportCount,
    const VkShadingRatePaletteNV*               pShadingRatePalettes);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstViewport` is the index of the first viewport whose shading rate palette is updated by the command.
- `viewportCount` is the number of viewports whose shading rate palettes are updated by the command.
- `pShadingRatePalettes` is a pointer to an array of [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShadingRatePaletteNV.html) structures defining the palette for each viewport.

## [](#_description)Description

This command sets the per-viewport shading rate image palettes for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html)::`pShadingRatePalettes` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-None-02064)VUID-vkCmdSetViewportShadingRatePaletteNV-None-02064  
  The [`shadingRateImage`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shadingRateImage) feature **must** be enabled
- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02067)VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02067  
  The sum of `firstViewport` and `viewportCount` **must** be between `1` and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive
- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02068)VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02068  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `firstViewport` **must** be `0`
- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-02069)VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-02069  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `viewportCount` **must** be `1`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-parameter)VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-pShadingRatePalettes-parameter)VUID-vkCmdSetViewportShadingRatePaletteNV-pShadingRatePalettes-parameter  
  `pShadingRatePalettes` **must** be a valid pointer to an array of `viewportCount` valid [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShadingRatePaletteNV.html) structures
- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-recording)VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-cmdpool)VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-videocoding)VUID-vkCmdSetViewportShadingRatePaletteNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-arraylength)VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-arraylength  
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

vkCmdSetViewportShadingRatePaletteNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShadingRatePaletteNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetViewportShadingRatePaletteNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
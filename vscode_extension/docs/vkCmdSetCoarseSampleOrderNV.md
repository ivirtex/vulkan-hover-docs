# vkCmdSetCoarseSampleOrderNV(3) Manual Page

## Name

vkCmdSetCoarseSampleOrderNV - Set order of coverage samples for coarse fragments dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the order of coverage samples in fragments larger than one pixel, call:

```c++
// Provided by VK_NV_shading_rate_image
void vkCmdSetCoarseSampleOrderNV(
    VkCommandBuffer                             commandBuffer,
    VkCoarseSampleOrderTypeNV                   sampleOrderType,
    uint32_t                                    customSampleOrderCount,
    const VkCoarseSampleOrderCustomNV*          pCustomSampleOrders);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `sampleOrderType` specifies the mechanism used to order coverage samples in fragments larger than one pixel.
- `customSampleOrderCount` specifies the number of custom sample orderings to use when ordering coverage samples.
- `pCustomSampleOrders` is a pointer to an array of [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderCustomNV.html) structures, each structure specifying the coverage sample order for a single combination of fragment area and coverage sample count.

## [](#_description)Description

If `sampleOrderType` is `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV`, the coverage sample order used for any combination of fragment area and coverage sample count not enumerated in `pCustomSampleOrders` will be identical to that used for `VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV`.

This command sets the order of coverage samples for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportCoarseSampleOrderStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html) values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-02081)VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-02081  
  If `sampleOrderType` is not `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV`, `customSamplerOrderCount` **must** be `0`
- [](#VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-02235)VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-02235  
  The array `pCustomSampleOrders` **must** not contain two structures with matching values for both the `shadingRate` and `sampleCount` members

Valid Usage (Implicit)

- [](#VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-parameter)VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-parameter)VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-parameter  
  `sampleOrderType` **must** be a valid [VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderTypeNV.html) value
- [](#VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-parameter)VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-parameter  
  If `customSampleOrderCount` is not `0`, `pCustomSampleOrders` **must** be a valid pointer to an array of `customSampleOrderCount` valid [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderCustomNV.html) structures
- [](#VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-recording)VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-cmdpool)VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetCoarseSampleOrderNV-videocoding)VUID-vkCmdSetCoarseSampleOrderNV-videocoding  
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

vkCmdSetCoarseSampleOrderNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderCustomNV.html), [VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderTypeNV.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetCoarseSampleOrderNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
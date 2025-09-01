# vkCmdBindShadingRateImageNV(3) Manual Page

## Name

vkCmdBindShadingRateImageNV - Bind a shading rate image on a command buffer



## [](#_c_specification)C Specification

When shading rate image usage is enabled in the bound pipeline, the pipeline uses a shading rate image specified by the command:

```c++
// Provided by VK_NV_shading_rate_image
void vkCmdBindShadingRateImageNV(
    VkCommandBuffer                             commandBuffer,
    VkImageView                                 imageView,
    VkImageLayout                               imageLayout);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `imageView` is an image view handle specifying the shading rate image. `imageView` **may** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), which is equivalent to specifying a view of an image filled with zero values.
- `imageLayout` is the layout that the image subresources accessible from `imageView` will be in when the shading rate image is accessed.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdBindShadingRateImageNV-None-02058)VUID-vkCmdBindShadingRateImageNV-None-02058  
  The [`shadingRateImage`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shadingRateImage) feature **must** be enabled
- [](#VUID-vkCmdBindShadingRateImageNV-imageView-02059)VUID-vkCmdBindShadingRateImageNV-imageView-02059  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle of type `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`
- [](#VUID-vkCmdBindShadingRateImageNV-imageView-02060)VUID-vkCmdBindShadingRateImageNV-imageView-02060  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** have a format of `VK_FORMAT_R8_UINT`
- [](#VUID-vkCmdBindShadingRateImageNV-imageView-02061)VUID-vkCmdBindShadingRateImageNV-imageView-02061  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** have been created with a `usage` value including `VK_IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV`
- [](#VUID-vkCmdBindShadingRateImageNV-imageView-02062)VUID-vkCmdBindShadingRateImageNV-imageView-02062  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** match the actual [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) of each subresource accessible from `imageView` at the time the subresource is accessed
- [](#VUID-vkCmdBindShadingRateImageNV-imageLayout-02063)VUID-vkCmdBindShadingRateImageNV-imageLayout-02063  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** be `VK_IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV` or `VK_IMAGE_LAYOUT_GENERAL`

Valid Usage (Implicit)

- [](#VUID-vkCmdBindShadingRateImageNV-commandBuffer-parameter)VUID-vkCmdBindShadingRateImageNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindShadingRateImageNV-imageView-parameter)VUID-vkCmdBindShadingRateImageNV-imageView-parameter  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-vkCmdBindShadingRateImageNV-imageLayout-parameter)VUID-vkCmdBindShadingRateImageNV-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdBindShadingRateImageNV-commandBuffer-recording)VUID-vkCmdBindShadingRateImageNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindShadingRateImageNV-commandBuffer-cmdpool)VUID-vkCmdBindShadingRateImageNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBindShadingRateImageNV-videocoding)VUID-vkCmdBindShadingRateImageNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindShadingRateImageNV-commonparent)VUID-vkCmdBindShadingRateImageNV-commonparent  
  Both of `commandBuffer`, and `imageView` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

vkCmdBindShadingRateImageNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindShadingRateImageNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
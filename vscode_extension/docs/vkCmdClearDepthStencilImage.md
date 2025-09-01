# vkCmdClearDepthStencilImage(3) Manual Page

## Name

vkCmdClearDepthStencilImage - Fill regions of a combined depth/stencil image



## [](#_c_specification)C Specification

To clear one or more subranges of a depth/stencil image, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdClearDepthStencilImage(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     image,
    VkImageLayout                               imageLayout,
    const VkClearDepthStencilValue*             pDepthStencil,
    uint32_t                                    rangeCount,
    const VkImageSubresourceRange*              pRanges);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `image` is the image to be cleared.
- `imageLayout` specifies the current layout of the image subresource ranges to be cleared, and **must** be `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`.
- `pDepthStencil` is a pointer to a [VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearDepthStencilValue.html) structure containing the values that the depth and stencil image subresource ranges will be cleared to (see [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears-values](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears-values) below).
- `rangeCount` is the number of image subresource range structures in `pRanges`.
- `pRanges` is a pointer to an array of [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structures describing a range of mipmap levels, array layers, and aspects to be cleared, as described in [Image Views](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views).

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdClearDepthStencilImage-image-01994)VUID-vkCmdClearDepthStencilImage-image-01994  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-format-features) of `image` **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`
- [](#VUID-vkCmdClearDepthStencilImage-pRanges-02658)VUID-vkCmdClearDepthStencilImage-pRanges-02658  
  If the `aspect` member of any element of `pRanges` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, and `image` was created with [separate stencil usage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html), `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage` used to create `image`
- [](#VUID-vkCmdClearDepthStencilImage-pRanges-02659)VUID-vkCmdClearDepthStencilImage-pRanges-02659  
  If the `aspect` member of any element of `pRanges` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, and `image` was not created with [separate stencil usage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html), `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` used to create `image`
- [](#VUID-vkCmdClearDepthStencilImage-pRanges-02660)VUID-vkCmdClearDepthStencilImage-pRanges-02660  
  If the `aspect` member of any element of `pRanges` includes `VK_IMAGE_ASPECT_DEPTH_BIT`, `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` used to create `image`
- [](#VUID-vkCmdClearDepthStencilImage-image-00010)VUID-vkCmdClearDepthStencilImage-image-00010  
  If `image` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdClearDepthStencilImage-imageLayout-00011)VUID-vkCmdClearDepthStencilImage-imageLayout-00011  
  `imageLayout` **must** specify the layout of the image subresource ranges of `image` specified in `pRanges` at the time this command is executed on a `VkDevice`
- [](#VUID-vkCmdClearDepthStencilImage-imageLayout-00012)VUID-vkCmdClearDepthStencilImage-imageLayout-00012  
  `imageLayout` **must** be either of `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-vkCmdClearDepthStencilImage-aspectMask-02824)VUID-vkCmdClearDepthStencilImage-aspectMask-02824  
  The [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`aspectMask` member of each element of the `pRanges` array **must** not include bits other than `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkCmdClearDepthStencilImage-image-02825)VUID-vkCmdClearDepthStencilImage-image-02825  
  If the `image`’s format does not have a stencil component, then the [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`aspectMask` member of each element of the `pRanges` array **must** not include the `VK_IMAGE_ASPECT_STENCIL_BIT` bit
- [](#VUID-vkCmdClearDepthStencilImage-image-02826)VUID-vkCmdClearDepthStencilImage-image-02826  
  If the `image`’s format does not have a depth component, then the [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`aspectMask` member of each element of the `pRanges` array **must** not include the `VK_IMAGE_ASPECT_DEPTH_BIT` bit
- [](#VUID-vkCmdClearDepthStencilImage-baseMipLevel-01474)VUID-vkCmdClearDepthStencilImage-baseMipLevel-01474  
  The [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseMipLevel` members of the elements of the `pRanges` array **must** each be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-vkCmdClearDepthStencilImage-pRanges-01694)VUID-vkCmdClearDepthStencilImage-pRanges-01694  
  For each [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) element of `pRanges`, if the `levelCount` member is not `VK_REMAINING_MIP_LEVELS`, then `baseMipLevel` + `levelCount` **must** be less than or equal to the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-vkCmdClearDepthStencilImage-baseArrayLayer-01476)VUID-vkCmdClearDepthStencilImage-baseArrayLayer-01476  
  The [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseArrayLayer` members of the elements of the `pRanges` array **must** each be less than the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-vkCmdClearDepthStencilImage-pRanges-01695)VUID-vkCmdClearDepthStencilImage-pRanges-01695  
  For each [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) element of `pRanges`, if the `layerCount` member is not `VK_REMAINING_ARRAY_LAYERS`, then `baseArrayLayer` + `layerCount` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-vkCmdClearDepthStencilImage-image-00014)VUID-vkCmdClearDepthStencilImage-image-00014  
  `image` **must** have a depth/stencil format
- [](#VUID-vkCmdClearDepthStencilImage-commandBuffer-01807)VUID-vkCmdClearDepthStencilImage-commandBuffer-01807  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, `image` **must** not be a protected image
- [](#VUID-vkCmdClearDepthStencilImage-commandBuffer-01808)VUID-vkCmdClearDepthStencilImage-commandBuffer-01808  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, `image` **must** not be an unprotected image

Valid Usage (Implicit)

- [](#VUID-vkCmdClearDepthStencilImage-commandBuffer-parameter)VUID-vkCmdClearDepthStencilImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdClearDepthStencilImage-image-parameter)VUID-vkCmdClearDepthStencilImage-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkCmdClearDepthStencilImage-imageLayout-parameter)VUID-vkCmdClearDepthStencilImage-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdClearDepthStencilImage-pDepthStencil-parameter)VUID-vkCmdClearDepthStencilImage-pDepthStencil-parameter  
  `pDepthStencil` **must** be a valid pointer to a valid [VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearDepthStencilValue.html) structure
- [](#VUID-vkCmdClearDepthStencilImage-pRanges-parameter)VUID-vkCmdClearDepthStencilImage-pRanges-parameter  
  `pRanges` **must** be a valid pointer to an array of `rangeCount` valid [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structures
- [](#VUID-vkCmdClearDepthStencilImage-commandBuffer-recording)VUID-vkCmdClearDepthStencilImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdClearDepthStencilImage-commandBuffer-cmdpool)VUID-vkCmdClearDepthStencilImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdClearDepthStencilImage-renderpass)VUID-vkCmdClearDepthStencilImage-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdClearDepthStencilImage-videocoding)VUID-vkCmdClearDepthStencilImage-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdClearDepthStencilImage-rangeCount-arraylength)VUID-vkCmdClearDepthStencilImage-rangeCount-arraylength  
  `rangeCount` **must** be greater than `0`
- [](#VUID-vkCmdClearDepthStencilImage-commonparent)VUID-vkCmdClearDepthStencilImage-commonparent  
  Both of `commandBuffer`, and `image` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Graphics

Action

Conditional Rendering

vkCmdClearDepthStencilImage is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearDepthStencilValue.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdClearDepthStencilImage).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
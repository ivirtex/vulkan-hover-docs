# vkCmdClearColorImage(3) Manual Page

## Name

vkCmdClearColorImage - Clear regions of a color image



## [](#_c_specification)C Specification

To clear one or more subranges of a color image, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdClearColorImage(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     image,
    VkImageLayout                               imageLayout,
    const VkClearColorValue*                    pColor,
    uint32_t                                    rangeCount,
    const VkImageSubresourceRange*              pRanges);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `image` is the image to be cleared.
- `imageLayout` specifies the current layout of the image subresource ranges to be cleared, and **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`.
- `pColor` is a pointer to a [VkClearColorValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearColorValue.html) structure containing the values that the image subresource ranges will be cleared to (see [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears-values](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears-values) below).
- `rangeCount` is the number of image subresource range structures in `pRanges`.
- `pRanges` is a pointer to an array of [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structures describing a range of mipmap levels, array layers, and aspects to be cleared, as described in [Image Views](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views).

## [](#_description)Description

Each specified range in `pRanges` is cleared to the value specified by `pColor`.

Valid Usage

- [](#VUID-vkCmdClearColorImage-image-01993)VUID-vkCmdClearColorImage-image-01993  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-format-features) of `image` **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`
- [](#VUID-vkCmdClearColorImage-image-00002)VUID-vkCmdClearColorImage-image-00002  
  `image` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdClearColorImage-image-01545)VUID-vkCmdClearColorImage-image-01545  
  `image` **must** not use any of the [formats that require a sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion)
- [](#VUID-vkCmdClearColorImage-image-00003)VUID-vkCmdClearColorImage-image-00003  
  If `image` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdClearColorImage-imageLayout-00004)VUID-vkCmdClearColorImage-imageLayout-00004  
  `imageLayout` **must** specify the layout of the image subresource ranges of `image` specified in `pRanges` at the time this command is executed on a `VkDevice`
- [](#VUID-vkCmdClearColorImage-imageLayout-01394)VUID-vkCmdClearColorImage-imageLayout-01394  
  `imageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-vkCmdClearColorImage-aspectMask-02498)VUID-vkCmdClearColorImage-aspectMask-02498  
  The [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`aspectMask` members of the elements of the `pRanges` array **must** each only include `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-vkCmdClearColorImage-baseMipLevel-01470)VUID-vkCmdClearColorImage-baseMipLevel-01470  
  The [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseMipLevel` members of the elements of the `pRanges` array **must** each be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-vkCmdClearColorImage-pRanges-01692)VUID-vkCmdClearColorImage-pRanges-01692  
  For each [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) element of `pRanges`, if the `levelCount` member is not `VK_REMAINING_MIP_LEVELS`, then `baseMipLevel` + `levelCount` **must** be less than or equal to the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-vkCmdClearColorImage-baseArrayLayer-01472)VUID-vkCmdClearColorImage-baseArrayLayer-01472  
  The [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseArrayLayer` members of the elements of the `pRanges` array **must** each be less than the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-vkCmdClearColorImage-pRanges-01693)VUID-vkCmdClearColorImage-pRanges-01693  
  For each [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) element of `pRanges`, if the `layerCount` member is not `VK_REMAINING_ARRAY_LAYERS`, then `baseArrayLayer` + `layerCount` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-vkCmdClearColorImage-image-00007)VUID-vkCmdClearColorImage-image-00007  
  `image` **must** not have a compressed or depth/stencil format
- [](#VUID-vkCmdClearColorImage-pColor-04961)VUID-vkCmdClearColorImage-pColor-04961  
  `pColor` **must** be a valid pointer to a [VkClearColorValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearColorValue.html) union
- [](#VUID-vkCmdClearColorImage-commandBuffer-01805)VUID-vkCmdClearColorImage-commandBuffer-01805  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, `image` **must** not be a protected image
- [](#VUID-vkCmdClearColorImage-commandBuffer-01806)VUID-vkCmdClearColorImage-commandBuffer-01806  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, **must** not be an unprotected image
- [](#VUID-vkCmdClearColorImage-image-09678)VUID-vkCmdClearColorImage-image-09678  
  If `image`’s format has components other than R and G, it **must** not have a 64-bit component width

Valid Usage (Implicit)

- [](#VUID-vkCmdClearColorImage-commandBuffer-parameter)VUID-vkCmdClearColorImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdClearColorImage-image-parameter)VUID-vkCmdClearColorImage-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkCmdClearColorImage-imageLayout-parameter)VUID-vkCmdClearColorImage-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdClearColorImage-pRanges-parameter)VUID-vkCmdClearColorImage-pRanges-parameter  
  `pRanges` **must** be a valid pointer to an array of `rangeCount` valid [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structures
- [](#VUID-vkCmdClearColorImage-commandBuffer-recording)VUID-vkCmdClearColorImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdClearColorImage-commandBuffer-cmdpool)VUID-vkCmdClearColorImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, or VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdClearColorImage-renderpass)VUID-vkCmdClearColorImage-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdClearColorImage-videocoding)VUID-vkCmdClearColorImage-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdClearColorImage-rangeCount-arraylength)VUID-vkCmdClearColorImage-rangeCount-arraylength  
  `rangeCount` **must** be greater than `0`
- [](#VUID-vkCmdClearColorImage-commonparent)VUID-vkCmdClearColorImage-commonparent  
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

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT

Action

Conditional Rendering

vkCmdClearColorImage is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkClearColorValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearColorValue.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdClearColorImage).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
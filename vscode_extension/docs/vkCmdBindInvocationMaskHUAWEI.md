# vkCmdBindInvocationMaskHUAWEI(3) Manual Page

## Name

vkCmdBindInvocationMaskHUAWEI - Bind an invocation mask image on a command buffer



## [](#_c_specification)C Specification

When invocation mask image usage is enabled in the bound ray tracing pipeline, the pipeline uses an invocation mask image specified by the command:

```c++
// Provided by VK_HUAWEI_invocation_mask
void vkCmdBindInvocationMaskHUAWEI(
    VkCommandBuffer                             commandBuffer,
    VkImageView                                 imageView,
    VkImageLayout                               imageLayout);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded
- `imageView` is an image view handle specifying the invocation mask image `imageView` **may** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), which is equivalent to specifying a view of an image filled with ones value.
- `imageLayout` is the layout that the image subresources accessible from `imageView` will be in when the invocation mask image is accessed

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdBindInvocationMaskHUAWEI-None-04976)VUID-vkCmdBindInvocationMaskHUAWEI-None-04976  
  The [`invocationMask`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-invocationMask) feature **must** be enabled
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04977)VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04977  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle of type `VK_IMAGE_VIEW_TYPE_2D`
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04978)VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04978  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** have a format of `VK_FORMAT_R8_UINT`
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04979)VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04979  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** have been created with `VK_IMAGE_USAGE_INVOCATION_MASK_BIT_HUAWEI` set
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04980)VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04980  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** be `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-width-04981)VUID-vkCmdBindInvocationMaskHUAWEI-width-04981  
  Thread mask image resolution **must** match the `width` and `height` in [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html)
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-None-04982)VUID-vkCmdBindInvocationMaskHUAWEI-None-04982  
  Each element in the invocation mask image **must** have the value `0` or `1`. The value 1 means the invocation is active
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-depth-04983)VUID-vkCmdBindInvocationMaskHUAWEI-depth-04983  
  `depth` in [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html) **must** be 1

Valid Usage (Implicit)

- [](#VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-parameter)VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-parameter)VUID-vkCmdBindInvocationMaskHUAWEI-imageView-parameter  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-imageLayout-parameter)VUID-vkCmdBindInvocationMaskHUAWEI-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-recording)VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-cmdpool)VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-renderpass)VUID-vkCmdBindInvocationMaskHUAWEI-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-videocoding)VUID-vkCmdBindInvocationMaskHUAWEI-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindInvocationMaskHUAWEI-commonparent)VUID-vkCmdBindInvocationMaskHUAWEI-commonparent  
  Both of `commandBuffer`, and `imageView` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Compute

State

## [](#_see_also)See Also

[VK\_HUAWEI\_invocation\_mask](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HUAWEI_invocation_mask.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindInvocationMaskHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
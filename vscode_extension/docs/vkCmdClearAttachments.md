# vkCmdClearAttachments(3) Manual Page

## Name

vkCmdClearAttachments - Clear regions within bound framebuffer attachments



## [](#_c_specification)C Specification

To clear one or more regions of color and depth/stencil attachments inside a render pass instance, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdClearAttachments(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    attachmentCount,
    const VkClearAttachment*                    pAttachments,
    uint32_t                                    rectCount,
    const VkClearRect*                          pRects);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `attachmentCount` is the number of entries in the `pAttachments` array.
- `pAttachments` is a pointer to an array of [VkClearAttachment](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearAttachment.html) structures defining the attachments to clear and the clear values to use.
- `rectCount` is the number of entries in the `pRects` array.
- `pRects` is a pointer to an array of [VkClearRect](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearRect.html) structures defining regions to clear for every attachment in `pAttachments`.

## [](#_description)Description

If the render pass has a [fragment density map attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-fragmentdensitymapattachment), clears follow the [operations of fragment density maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragmentdensitymapops) as if each clear region was a primitive which generates fragments. The clear color is applied to all pixels inside each fragment’s area regardless if the pixels lie outside of the clear region. Clears **may** have a different set of supported fragment areas than draws.

Unlike other [clear commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears), [vkCmdClearAttachments](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearAttachments.html) is not a transfer command. It performs its operations in [rasterization order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-order). For color attachments, the operations are executed as color attachment writes, by the `VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT` stage. For depth/stencil attachments, the operations are executed as [depth writes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth) and [stencil writes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-stencil) by the `VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT` and `VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT` stages.

`vkCmdClearAttachments` is not affected by the bound pipeline state.

Note

It is generally preferable to clear attachments by using the `VK_ATTACHMENT_LOAD_OP_CLEAR` load operation at the start of rendering, as it is more efficient on some implementations.

If any attachment’s `aspectMask` to be cleared is not backed by an image view, the clear has no effect on that aspect.

If an attachment being cleared refers to an image view created with an `aspectMask` equal to one of `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT` or `VK_IMAGE_ASPECT_PLANE_2_BIT`, it is considered to be `VK_IMAGE_ASPECT_COLOR_BIT` for purposes of this command, and **must** be cleared with the `VK_IMAGE_ASPECT_COLOR_BIT` aspect as specified by [image view creation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#image-views-plane-promotion).

Valid Usage

- [](#VUID-vkCmdClearAttachments-aspectMask-07884)VUID-vkCmdClearAttachments-aspectMask-07884  
  If the current render pass instance does not use dynamic rendering, and the `aspectMask` member of any element of `pAttachments` contains `VK_IMAGE_ASPECT_DEPTH_BIT`, the current subpass instance’s depth-stencil attachment **must** be either `VK_ATTACHMENT_UNUSED` or the attachment `format` **must** contain a depth component
- [](#VUID-vkCmdClearAttachments-aspectMask-07885)VUID-vkCmdClearAttachments-aspectMask-07885  
  If the current render pass instance does not use dynamic rendering, and the `aspectMask` member of any element of `pAttachments` contains `VK_IMAGE_ASPECT_STENCIL_BIT`, the current subpass instance’s depth-stencil attachment **must** be either `VK_ATTACHMENT_UNUSED` or the attachment `format` **must** contain a stencil component
- [](#VUID-vkCmdClearAttachments-aspectMask-07271)VUID-vkCmdClearAttachments-aspectMask-07271  
  If the `aspectMask` member of any element of `pAttachments` contains `VK_IMAGE_ASPECT_COLOR_BIT`, the `colorAttachment` **must** be a valid color attachment index in the current render pass instance
- [](#VUID-vkCmdClearAttachments-rect-02682)VUID-vkCmdClearAttachments-rect-02682  
  The `rect` member of each element of `pRects` **must** have an `extent.width` greater than `0`
- [](#VUID-vkCmdClearAttachments-rect-02683)VUID-vkCmdClearAttachments-rect-02683  
  The `rect` member of each element of `pRects` **must** have an `extent.height` greater than `0`
- [](#VUID-vkCmdClearAttachments-pRects-00016)VUID-vkCmdClearAttachments-pRects-00016  
  The rectangular region specified by each element of `pRects` **must** be contained within the render area of the current render pass instance
- [](#VUID-vkCmdClearAttachments-pRects-06937)VUID-vkCmdClearAttachments-pRects-06937  
  The layers specified by each element of `pRects` **must** be contained within every attachment that `pAttachments` refers to, i.e. for each element of `pRects`, [VkClearRect](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearRect.html)::`baseArrayLayer` + [VkClearRect](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearRect.html)::`layerCount` **must** be less than or equal to the number of layers rendered to in the current render pass instance
- [](#VUID-vkCmdClearAttachments-layerCount-01934)VUID-vkCmdClearAttachments-layerCount-01934  
  The `layerCount` member of each element of `pRects` **must** not be `0`
- [](#VUID-vkCmdClearAttachments-commandBuffer-02504)VUID-vkCmdClearAttachments-commandBuffer-02504  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, each attachment to be cleared **must** not be a protected image
- [](#VUID-vkCmdClearAttachments-commandBuffer-02505)VUID-vkCmdClearAttachments-commandBuffer-02505  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, each attachment to be cleared **must** not be an unprotected image
- [](#VUID-vkCmdClearAttachments-baseArrayLayer-00018)VUID-vkCmdClearAttachments-baseArrayLayer-00018  
  If the render pass instance this is recorded in uses multiview, then `baseArrayLayer` **must** be zero and `layerCount` **must** be one
- [](#VUID-vkCmdClearAttachments-colorAttachment-09503)VUID-vkCmdClearAttachments-colorAttachment-09503  
  The `colorAttachment` member of each element of `pAttachments` **must** not identify a color attachment that is currently mapped to `VK_ATTACHMENT_UNUSED` in `commandBuffer` via [VkRenderingAttachmentLocationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentLocationInfo.html)
- [](#VUID-vkCmdClearAttachments-aspectMask-09298)VUID-vkCmdClearAttachments-aspectMask-09298  
  If the subpass this is recorded in performs an external format resolve, the `aspectMask` member of any element of `pAttachments` **must** not include `VK_IMAGE_ASPECT_PLANE_i_BIT` for any index *i*
- [](#VUID-vkCmdClearAttachments-None-09679)VUID-vkCmdClearAttachments-None-09679  
  If the attachment format has components other than R and G, it **must** not have a 64-bit component width
- [](#VUID-vkCmdClearAttachments-None-10616)VUID-vkCmdClearAttachments-None-10616  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdClearAttachments-commandBuffer-parameter)VUID-vkCmdClearAttachments-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdClearAttachments-pAttachments-parameter)VUID-vkCmdClearAttachments-pAttachments-parameter  
  `pAttachments` **must** be a valid pointer to an array of `attachmentCount` valid [VkClearAttachment](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearAttachment.html) structures
- [](#VUID-vkCmdClearAttachments-pRects-parameter)VUID-vkCmdClearAttachments-pRects-parameter  
  `pRects` **must** be a valid pointer to an array of `rectCount` [VkClearRect](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearRect.html) structures
- [](#VUID-vkCmdClearAttachments-commandBuffer-recording)VUID-vkCmdClearAttachments-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdClearAttachments-commandBuffer-cmdpool)VUID-vkCmdClearAttachments-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdClearAttachments-renderpass)VUID-vkCmdClearAttachments-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdClearAttachments-videocoding)VUID-vkCmdClearAttachments-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdClearAttachments-attachmentCount-arraylength)VUID-vkCmdClearAttachments-attachmentCount-arraylength  
  `attachmentCount` **must** be greater than `0`
- [](#VUID-vkCmdClearAttachments-rectCount-arraylength)VUID-vkCmdClearAttachments-rectCount-arraylength  
  `rectCount` **must** be greater than `0`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Inside

Outside

Graphics

Action

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkClearAttachment](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearAttachment.html), [VkClearRect](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearRect.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdClearAttachments)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
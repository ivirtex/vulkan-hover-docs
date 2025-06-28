# vkCmdBeginRenderPass(3) Manual Page

## Name

vkCmdBeginRenderPass - Begin a new render pass



## [](#_c_specification)C Specification

To begin a render pass instance, call:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkCmdBeginRenderPass(
    VkCommandBuffer                             commandBuffer,
    const VkRenderPassBeginInfo*                pRenderPassBegin,
    VkSubpassContents                           contents);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pRenderPassBegin` is a pointer to a [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html) structure specifying the render pass to begin an instance of, and the framebuffer the instance uses.
- `contents` is a [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html) value specifying how the commands in the first subpass will be provided.

## [](#_description)Description

After beginning a render pass instance, the command buffer is ready to record the commands for the first subpass of that render pass.

Valid Usage

- [](#VUID-vkCmdBeginRenderPass-initialLayout-00895)VUID-vkCmdBeginRenderPass-initialLayout-00895  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass-initialLayout-01758)VUID-vkCmdBeginRenderPass-initialLayout-01758  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass-initialLayout-02842)VUID-vkCmdBeginRenderPass-initialLayout-02842  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass-stencilInitialLayout-02843)VUID-vkCmdBeginRenderPass-stencilInitialLayout-02843  
  If any of the `stencilInitialLayout` or `stencilFinalLayout` member of the `VkAttachmentDescriptionStencilLayout` structures or the `stencilLayout` member of the `VkAttachmentReferenceStencilLayout` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass-initialLayout-00897)VUID-vkCmdBeginRenderPass-initialLayout-00897  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_SAMPLED_BIT` or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass-initialLayout-00898)VUID-vkCmdBeginRenderPass-initialLayout-00898  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_TRANSFER_SRC_BIT`
- [](#VUID-vkCmdBeginRenderPass-initialLayout-00899)VUID-vkCmdBeginRenderPass-initialLayout-00899  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_TRANSFER_DST_BIT`
- [](#VUID-vkCmdBeginRenderPass-initialLayout-00900)VUID-vkCmdBeginRenderPass-initialLayout-00900  
  If the `initialLayout` member of any of the `VkAttachmentDescription` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is not `VK_IMAGE_LAYOUT_UNDEFINED`, then each such `initialLayout` **must** be equal to the current layout of the corresponding attachment image subresource of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin`
- [](#VUID-vkCmdBeginRenderPass-srcStageMask-06451)VUID-vkCmdBeginRenderPass-srcStageMask-06451  
  The `srcStageMask` members of any element of the `pDependencies` member of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) used to create `renderPass` **must** be supported by the capabilities of the queue family identified by the `queueFamilyIndex` member of the [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) used to create the command pool which `commandBuffer` was allocated from
- [](#VUID-vkCmdBeginRenderPass-dstStageMask-06452)VUID-vkCmdBeginRenderPass-dstStageMask-06452  
  The `dstStageMask` members of any element of the `pDependencies` member of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) used to create `renderPass` **must** be supported by the capabilities of the queue family identified by the `queueFamilyIndex` member of the [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) used to create the command pool which `commandBuffer` was allocated from
- [](#VUID-vkCmdBeginRenderPass-framebuffer-02532)VUID-vkCmdBeginRenderPass-framebuffer-02532  
  For any attachment in `framebuffer` that is used by `renderPass` and is bound to memory locations that are also bound to another attachment used by `renderPass`, and if at least one of those uses causes either attachment to be written to, both attachments **must** have had the `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` set
- [](#VUID-vkCmdBeginRenderPass-framebuffer-09045)VUID-vkCmdBeginRenderPass-framebuffer-09045  
  If any attachments specified in `framebuffer` are used by `renderPass` and are bound to overlapping memory locations, there **must** be only one that is used as a color attachment, depth/stencil, or resolve attachment in any subpass
- [](#VUID-vkCmdBeginRenderPass-initialLayout-07000)VUID-vkCmdBeginRenderPass-initialLayout-07000  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including either the `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` and either the `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT` usage bits
- [](#VUID-vkCmdBeginRenderPass-initialLayout-07001)VUID-vkCmdBeginRenderPass-initialLayout-07001  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value the `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` usage bit
- [](#VUID-vkCmdBeginRenderPass-initialLayout-09537)VUID-vkCmdBeginRenderPass-initialLayout-09537  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including either `VK_IMAGE_USAGE_STORAGE_BIT`, or both `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` and either of `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass-contents-09640)VUID-vkCmdBeginRenderPass-contents-09640  
  If `contents` is `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_KHR`, then at least one of the following features **must** be enabled:
  
  - [`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7)
  - [`nestedCommandBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nestedCommandBuffer)

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginRenderPass-commandBuffer-parameter)VUID-vkCmdBeginRenderPass-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginRenderPass-pRenderPassBegin-parameter)VUID-vkCmdBeginRenderPass-pRenderPassBegin-parameter  
  `pRenderPassBegin` **must** be a valid pointer to a valid [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html) structure
- [](#VUID-vkCmdBeginRenderPass-contents-parameter)VUID-vkCmdBeginRenderPass-contents-parameter  
  `contents` **must** be a valid [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html) value
- [](#VUID-vkCmdBeginRenderPass-commandBuffer-recording)VUID-vkCmdBeginRenderPass-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginRenderPass-commandBuffer-cmdpool)VUID-vkCmdBeginRenderPass-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginRenderPass-renderpass)VUID-vkCmdBeginRenderPass-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBeginRenderPass-videocoding)VUID-vkCmdBeginRenderPass-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBeginRenderPass-bufferlevel)VUID-vkCmdBeginRenderPass-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary

Outside

Outside

Graphics

Action  
State  
Synchronization

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html), [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginRenderPass)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
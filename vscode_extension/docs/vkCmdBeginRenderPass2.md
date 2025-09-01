# vkCmdBeginRenderPass2(3) Manual Page

## Name

vkCmdBeginRenderPass2 - Begin a new render pass



## [](#_c_specification)C Specification

Alternatively to begin a render pass, call:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
void vkCmdBeginRenderPass2(
    VkCommandBuffer                             commandBuffer,
    const VkRenderPassBeginInfo*                pRenderPassBegin,
    const VkSubpassBeginInfo*                   pSubpassBeginInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_create_renderpass2
void vkCmdBeginRenderPass2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkRenderPassBeginInfo*                pRenderPassBegin,
    const VkSubpassBeginInfo*                   pSubpassBeginInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pRenderPassBegin` is a pointer to a [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html) structure specifying the render pass to begin an instance of, and the framebuffer the instance uses. After recording this command, the render pass and framebuffer **may** be accessed at any point that `commandBuffer` is in the recording or pending state until it is reset.
- `pSubpassBeginInfo` is a pointer to a [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassBeginInfo.html) structure containing information about the subpass which is about to begin rendering.

## [](#_description)Description

After beginning a render pass instance, the command buffer is ready to record the commands for the first subpass of that render pass.

Valid Usage

- [](#VUID-vkCmdBeginRenderPass2-framebuffer-02779)VUID-vkCmdBeginRenderPass2-framebuffer-02779  
  Both the `framebuffer` and `renderPass` members of `pRenderPassBegin` **must** have been created on the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) that `commandBuffer` was allocated on
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-03094)VUID-vkCmdBeginRenderPass2-initialLayout-03094  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-03096)VUID-vkCmdBeginRenderPass2-initialLayout-03096  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-02844)VUID-vkCmdBeginRenderPass2-initialLayout-02844  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass2-stencilInitialLayout-02845)VUID-vkCmdBeginRenderPass2-stencilInitialLayout-02845  
  If any of the `stencilInitialLayout` or `stencilFinalLayout` member of the `VkAttachmentDescriptionStencilLayout` structures or the `stencilLayout` member of the `VkAttachmentReferenceStencilLayout` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-03097)VUID-vkCmdBeginRenderPass2-initialLayout-03097  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_SAMPLED_BIT` or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-03098)VUID-vkCmdBeginRenderPass2-initialLayout-03098  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_TRANSFER_SRC_BIT`
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-03099)VUID-vkCmdBeginRenderPass2-initialLayout-03099  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_TRANSFER_DST_BIT`
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-03100)VUID-vkCmdBeginRenderPass2-initialLayout-03100  
  If the `initialLayout` member of any of the `VkAttachmentDescription` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is not `VK_IMAGE_LAYOUT_UNDEFINED`, then each such `initialLayout` **must** be equal to the current layout of the corresponding attachment image subresource of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin`
- [](#VUID-vkCmdBeginRenderPass2-srcStageMask-06453)VUID-vkCmdBeginRenderPass2-srcStageMask-06453  
  The `srcStageMask` members of any element of the `pDependencies` member of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) used to create `renderPass` **must** be supported by the capabilities of the queue family identified by the `queueFamilyIndex` member of the [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) used to create the command pool which `commandBuffer` was allocated from
- [](#VUID-vkCmdBeginRenderPass2-dstStageMask-06454)VUID-vkCmdBeginRenderPass2-dstStageMask-06454  
  The `dstStageMask` members of any element of the `pDependencies` member of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) used to create `renderPass` **must** be supported by the capabilities of the queue family identified by the `queueFamilyIndex` member of the [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) used to create the command pool which `commandBuffer` was allocated from
- [](#VUID-vkCmdBeginRenderPass2-framebuffer-02533)VUID-vkCmdBeginRenderPass2-framebuffer-02533  
  For any attachment in `framebuffer` that is used by `renderPass` and is bound to memory locations that are also bound to another attachment used by `renderPass`, and if at least one of those uses causes either attachment to be written to, both attachments **must** have had the `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` set
- [](#VUID-vkCmdBeginRenderPass2-framebuffer-09046)VUID-vkCmdBeginRenderPass2-framebuffer-09046  
  If any attachments specified in `framebuffer` are used by `renderPass` and are bound to overlapping memory locations, there **must** be only one that is used as a color attachment, depth/stencil, or resolve attachment in any subpass
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-07002)VUID-vkCmdBeginRenderPass2-initialLayout-07002  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including either the `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` and either the `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT` usage bits
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-07003)VUID-vkCmdBeginRenderPass2-initialLayout-07003  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value the `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` usage bit
- [](#VUID-vkCmdBeginRenderPass2-initialLayout-09538)VUID-vkCmdBeginRenderPass2-initialLayout-09538  
  If any of the `initialLayout` or `finalLayout` member of the `VkAttachmentDescription` structures or the `layout` member of the `VkAttachmentReference` structures specified when creating the render pass specified in the `renderPass` member of `pRenderPassBegin` is `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ` then the corresponding attachment image view of the framebuffer specified in the `framebuffer` member of `pRenderPassBegin` **must** have been created with a `usage` value including either `VK_IMAGE_USAGE_STORAGE_BIT`, or both `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` and either of `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-vkCmdBeginRenderPass2-flags-10652)VUID-vkCmdBeginRenderPass2-flags-10652  
  If `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` was included in the [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags` used to create the `renderPass`, `commandBuffer` **must** not have been created with `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginRenderPass2-commandBuffer-parameter)VUID-vkCmdBeginRenderPass2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginRenderPass2-pRenderPassBegin-parameter)VUID-vkCmdBeginRenderPass2-pRenderPassBegin-parameter  
  `pRenderPassBegin` **must** be a valid pointer to a valid [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html) structure
- [](#VUID-vkCmdBeginRenderPass2-pSubpassBeginInfo-parameter)VUID-vkCmdBeginRenderPass2-pSubpassBeginInfo-parameter  
  `pSubpassBeginInfo` **must** be a valid pointer to a valid [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassBeginInfo.html) structure
- [](#VUID-vkCmdBeginRenderPass2-commandBuffer-recording)VUID-vkCmdBeginRenderPass2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginRenderPass2-commandBuffer-cmdpool)VUID-vkCmdBeginRenderPass2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginRenderPass2-renderpass)VUID-vkCmdBeginRenderPass2-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBeginRenderPass2-videocoding)VUID-vkCmdBeginRenderPass2-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBeginRenderPass2-bufferlevel)VUID-vkCmdBeginRenderPass2-bufferlevel  
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

Conditional Rendering

vkCmdBeginRenderPass2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html), [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassBeginInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginRenderPass2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# vkCmdBeginRenderPass2(3) Manual Page

## Name

vkCmdBeginRenderPass2 - Begin a new render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively to begin a render pass, call:

``` c
// Provided by VK_VERSION_1_2
void vkCmdBeginRenderPass2(
    VkCommandBuffer                             commandBuffer,
    const VkRenderPassBeginInfo*                pRenderPassBegin,
    const VkSubpassBeginInfo*                   pSubpassBeginInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_create_renderpass2
void vkCmdBeginRenderPass2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkRenderPassBeginInfo*                pRenderPassBegin,
    const VkSubpassBeginInfo*                   pSubpassBeginInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pRenderPassBegin` is a pointer to a
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) structure
  specifying the render pass to begin an instance of, and the
  framebuffer the instance uses.

- `pSubpassBeginInfo` is a pointer to a
  [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfo.html) structure containing
  information about the subpass which is about to begin rendering.

## <a href="#_description" class="anchor"></a>Description

After beginning a render pass instance, the command buffer is ready to
record the commands for the first subpass of that render pass.

Valid Usage

- <a href="#VUID-vkCmdBeginRenderPass2-framebuffer-02779"
  id="VUID-vkCmdBeginRenderPass2-framebuffer-02779"></a>
  VUID-vkCmdBeginRenderPass2-framebuffer-02779  
  Both the `framebuffer` and `renderPass` members of `pRenderPassBegin`
  **must** have been created on the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) that
  `commandBuffer` was allocated on

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-03094"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-03094"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-03094  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-03096"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-03096"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-03096  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL` then the
  corresponding attachment image view of the framebuffer specified in
  the `framebuffer` member of `pRenderPassBegin` **must** have been
  created with a `usage` value including
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-02844"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-02844"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-02844  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-vkCmdBeginRenderPass2-stencilInitialLayout-02845"
  id="VUID-vkCmdBeginRenderPass2-stencilInitialLayout-02845"></a>
  VUID-vkCmdBeginRenderPass2-stencilInitialLayout-02845  
  If any of the `stencilInitialLayout` or `stencilFinalLayout` member of
  the `VkAttachmentDescriptionStencilLayout` structures or the
  `stencilLayout` member of the `VkAttachmentReferenceStencilLayout`
  structures specified when creating the render pass specified in the
  `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-03097"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-03097"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-03097  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including `VK_IMAGE_USAGE_SAMPLED_BIT` or
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-03098"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-03098"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-03098  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including `VK_IMAGE_USAGE_TRANSFER_SRC_BIT`

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-03099"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-03099"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-03099  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including `VK_IMAGE_USAGE_TRANSFER_DST_BIT`

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-03100"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-03100"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-03100  
  If the `initialLayout` member of any of the `VkAttachmentDescription`
  structures specified when creating the render pass specified in the
  `renderPass` member of `pRenderPassBegin` is not
  `VK_IMAGE_LAYOUT_UNDEFINED`, then each such `initialLayout` **must**
  be equal to the current layout of the corresponding attachment image
  subresource of the framebuffer specified in the `framebuffer` member
  of `pRenderPassBegin`

- <a href="#VUID-vkCmdBeginRenderPass2-srcStageMask-06453"
  id="VUID-vkCmdBeginRenderPass2-srcStageMask-06453"></a>
  VUID-vkCmdBeginRenderPass2-srcStageMask-06453  
  The `srcStageMask` members of any element of the `pDependencies`
  member of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) used
  to create `renderPass` **must** be supported by the capabilities of
  the queue family identified by the `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) used to create
  the command pool which `commandBuffer` was allocated from

- <a href="#VUID-vkCmdBeginRenderPass2-dstStageMask-06454"
  id="VUID-vkCmdBeginRenderPass2-dstStageMask-06454"></a>
  VUID-vkCmdBeginRenderPass2-dstStageMask-06454  
  The `dstStageMask` members of any element of the `pDependencies`
  member of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) used
  to create `renderPass` **must** be supported by the capabilities of
  the queue family identified by the `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) used to create
  the command pool which `commandBuffer` was allocated from

- <a href="#VUID-vkCmdBeginRenderPass2-framebuffer-02533"
  id="VUID-vkCmdBeginRenderPass2-framebuffer-02533"></a>
  VUID-vkCmdBeginRenderPass2-framebuffer-02533  
  For any attachment in `framebuffer` that is used by `renderPass` and
  is bound to memory locations that are also bound to another attachment
  used by `renderPass`, and if at least one of those uses causes either
  attachment to be written to, both attachments **must** have had the
  `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` set

- <a href="#VUID-vkCmdBeginRenderPass2-framebuffer-09046"
  id="VUID-vkCmdBeginRenderPass2-framebuffer-09046"></a>
  VUID-vkCmdBeginRenderPass2-framebuffer-09046  
  If any attachments specified in `framebuffer` are used by `renderPass`
  and are bound to overlapping memory locations, there **must** be only
  one that is used as a color attachment, depth/stencil, or resolve
  attachment in any subpass

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-07002"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-07002"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-07002  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then the
  corresponding attachment image view of the framebuffer specified in
  the `framebuffer` member of `pRenderPassBegin` **must** have been
  created with a `usage` value including either the
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` and either the
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT`
  usage bits

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-07003"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-07003"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-07003  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then the
  corresponding attachment image view of the framebuffer specified in
  the `framebuffer` member of `pRenderPassBegin` **must** have been
  created with a `usage` value the
  `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` usage bit

- <a href="#VUID-vkCmdBeginRenderPass2-initialLayout-09538"
  id="VUID-vkCmdBeginRenderPass2-initialLayout-09538"></a>
  VUID-vkCmdBeginRenderPass2-initialLayout-09538  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including either `VK_IMAGE_USAGE_STORAGE_BIT`, or
  both `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` and either of
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBeginRenderPass2-commandBuffer-parameter"
  id="VUID-vkCmdBeginRenderPass2-commandBuffer-parameter"></a>
  VUID-vkCmdBeginRenderPass2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBeginRenderPass2-pRenderPassBegin-parameter"
  id="VUID-vkCmdBeginRenderPass2-pRenderPassBegin-parameter"></a>
  VUID-vkCmdBeginRenderPass2-pRenderPassBegin-parameter  
  `pRenderPassBegin` **must** be a valid pointer to a valid
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) structure

- <a href="#VUID-vkCmdBeginRenderPass2-pSubpassBeginInfo-parameter"
  id="VUID-vkCmdBeginRenderPass2-pSubpassBeginInfo-parameter"></a>
  VUID-vkCmdBeginRenderPass2-pSubpassBeginInfo-parameter  
  `pSubpassBeginInfo` **must** be a valid pointer to a valid
  [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfo.html) structure

- <a href="#VUID-vkCmdBeginRenderPass2-commandBuffer-recording"
  id="VUID-vkCmdBeginRenderPass2-commandBuffer-recording"></a>
  VUID-vkCmdBeginRenderPass2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBeginRenderPass2-commandBuffer-cmdpool"
  id="VUID-vkCmdBeginRenderPass2-commandBuffer-cmdpool"></a>
  VUID-vkCmdBeginRenderPass2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBeginRenderPass2-renderpass"
  id="VUID-vkCmdBeginRenderPass2-renderpass"></a>
  VUID-vkCmdBeginRenderPass2-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBeginRenderPass2-videocoding"
  id="VUID-vkCmdBeginRenderPass2-videocoding"></a>
  VUID-vkCmdBeginRenderPass2-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBeginRenderPass2-bufferlevel"
  id="VUID-vkCmdBeginRenderPass2-bufferlevel"></a>
  VUID-vkCmdBeginRenderPass2-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p>Primary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State<br />
Synchronization</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html),
[VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBeginRenderPass2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

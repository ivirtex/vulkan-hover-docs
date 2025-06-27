# vkCmdBeginRenderPass(3) Manual Page

## Name

vkCmdBeginRenderPass - Begin a new render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin a render pass instance, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdBeginRenderPass(
    VkCommandBuffer                             commandBuffer,
    const VkRenderPassBeginInfo*                pRenderPassBegin,
    VkSubpassContents                           contents);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pRenderPassBegin` is a pointer to a
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) structure
  specifying the render pass to begin an instance of, and the
  framebuffer the instance uses.

- `contents` is a [VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html) value
  specifying how the commands in the first subpass will be provided.

## <a href="#_description" class="anchor"></a>Description

After beginning a render pass instance, the command buffer is ready to
record the commands for the first subpass of that render pass.

Valid Usage

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-00895"
  id="VUID-vkCmdBeginRenderPass-initialLayout-00895"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-00895  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-01758"
  id="VUID-vkCmdBeginRenderPass-initialLayout-01758"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-01758  
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

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-02842"
  id="VUID-vkCmdBeginRenderPass-initialLayout-02842"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-02842  
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

- <a href="#VUID-vkCmdBeginRenderPass-stencilInitialLayout-02843"
  id="VUID-vkCmdBeginRenderPass-stencilInitialLayout-02843"></a>
  VUID-vkCmdBeginRenderPass-stencilInitialLayout-02843  
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

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-00897"
  id="VUID-vkCmdBeginRenderPass-initialLayout-00897"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-00897  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including `VK_IMAGE_USAGE_SAMPLED_BIT` or
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-00898"
  id="VUID-vkCmdBeginRenderPass-initialLayout-00898"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-00898  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including `VK_IMAGE_USAGE_TRANSFER_SRC_BIT`

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-00899"
  id="VUID-vkCmdBeginRenderPass-initialLayout-00899"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-00899  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` then the corresponding
  attachment image view of the framebuffer specified in the
  `framebuffer` member of `pRenderPassBegin` **must** have been created
  with a `usage` value including `VK_IMAGE_USAGE_TRANSFER_DST_BIT`

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-00900"
  id="VUID-vkCmdBeginRenderPass-initialLayout-00900"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-00900  
  If the `initialLayout` member of any of the `VkAttachmentDescription`
  structures specified when creating the render pass specified in the
  `renderPass` member of `pRenderPassBegin` is not
  `VK_IMAGE_LAYOUT_UNDEFINED`, then each such `initialLayout` **must**
  be equal to the current layout of the corresponding attachment image
  subresource of the framebuffer specified in the `framebuffer` member
  of `pRenderPassBegin`

- <a href="#VUID-vkCmdBeginRenderPass-srcStageMask-06451"
  id="VUID-vkCmdBeginRenderPass-srcStageMask-06451"></a>
  VUID-vkCmdBeginRenderPass-srcStageMask-06451  
  The `srcStageMask` members of any element of the `pDependencies`
  member of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) used
  to create `renderPass` **must** be supported by the capabilities of
  the queue family identified by the `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) used to create
  the command pool which `commandBuffer` was allocated from

- <a href="#VUID-vkCmdBeginRenderPass-dstStageMask-06452"
  id="VUID-vkCmdBeginRenderPass-dstStageMask-06452"></a>
  VUID-vkCmdBeginRenderPass-dstStageMask-06452  
  The `dstStageMask` members of any element of the `pDependencies`
  member of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) used
  to create `renderPass` **must** be supported by the capabilities of
  the queue family identified by the `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) used to create
  the command pool which `commandBuffer` was allocated from

- <a href="#VUID-vkCmdBeginRenderPass-framebuffer-02532"
  id="VUID-vkCmdBeginRenderPass-framebuffer-02532"></a>
  VUID-vkCmdBeginRenderPass-framebuffer-02532  
  For any attachment in `framebuffer` that is used by `renderPass` and
  is bound to memory locations that are also bound to another attachment
  used by `renderPass`, and if at least one of those uses causes either
  attachment to be written to, both attachments **must** have had the
  `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` set

- <a href="#VUID-vkCmdBeginRenderPass-framebuffer-09045"
  id="VUID-vkCmdBeginRenderPass-framebuffer-09045"></a>
  VUID-vkCmdBeginRenderPass-framebuffer-09045  
  If any attachments specified in `framebuffer` are used by `renderPass`
  and are bound to overlapping memory locations, there **must** be only
  one that is used as a color attachment, depth/stencil, or resolve
  attachment in any subpass

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-07000"
  id="VUID-vkCmdBeginRenderPass-initialLayout-07000"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-07000  
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

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-07001"
  id="VUID-vkCmdBeginRenderPass-initialLayout-07001"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-07001  
  If any of the `initialLayout` or `finalLayout` member of the
  `VkAttachmentDescription` structures or the `layout` member of the
  `VkAttachmentReference` structures specified when creating the render
  pass specified in the `renderPass` member of `pRenderPassBegin` is
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then the
  corresponding attachment image view of the framebuffer specified in
  the `framebuffer` member of `pRenderPassBegin` **must** have been
  created with a `usage` value the
  `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` usage bit

- <a href="#VUID-vkCmdBeginRenderPass-initialLayout-09537"
  id="VUID-vkCmdBeginRenderPass-initialLayout-09537"></a>
  VUID-vkCmdBeginRenderPass-initialLayout-09537  
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

- <a href="#VUID-vkCmdBeginRenderPass-contents-09640"
  id="VUID-vkCmdBeginRenderPass-contents-09640"></a>
  VUID-vkCmdBeginRenderPass-contents-09640  
  If `contents` is
  `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_KHR`, then
  at least one of the following features **must** be enabled:

  - <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
    target="_blank" rel="noopener"><code>maintenance7</code></a>

  - <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
    target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBeginRenderPass-commandBuffer-parameter"
  id="VUID-vkCmdBeginRenderPass-commandBuffer-parameter"></a>
  VUID-vkCmdBeginRenderPass-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBeginRenderPass-pRenderPassBegin-parameter"
  id="VUID-vkCmdBeginRenderPass-pRenderPassBegin-parameter"></a>
  VUID-vkCmdBeginRenderPass-pRenderPassBegin-parameter  
  `pRenderPassBegin` **must** be a valid pointer to a valid
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) structure

- <a href="#VUID-vkCmdBeginRenderPass-contents-parameter"
  id="VUID-vkCmdBeginRenderPass-contents-parameter"></a>
  VUID-vkCmdBeginRenderPass-contents-parameter  
  `contents` **must** be a valid
  [VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html) value

- <a href="#VUID-vkCmdBeginRenderPass-commandBuffer-recording"
  id="VUID-vkCmdBeginRenderPass-commandBuffer-recording"></a>
  VUID-vkCmdBeginRenderPass-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBeginRenderPass-commandBuffer-cmdpool"
  id="VUID-vkCmdBeginRenderPass-commandBuffer-cmdpool"></a>
  VUID-vkCmdBeginRenderPass-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBeginRenderPass-renderpass"
  id="VUID-vkCmdBeginRenderPass-renderpass"></a>
  VUID-vkCmdBeginRenderPass-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBeginRenderPass-videocoding"
  id="VUID-vkCmdBeginRenderPass-videocoding"></a>
  VUID-vkCmdBeginRenderPass-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBeginRenderPass-bufferlevel"
  id="VUID-vkCmdBeginRenderPass-bufferlevel"></a>
  VUID-vkCmdBeginRenderPass-bufferlevel  
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

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html),
[VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBeginRenderPass"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

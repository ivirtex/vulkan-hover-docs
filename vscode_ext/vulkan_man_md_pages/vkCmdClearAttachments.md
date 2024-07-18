# vkCmdClearAttachments(3) Manual Page

## Name

vkCmdClearAttachments - Clear regions within bound framebuffer
attachments



## <a href="#_c_specification" class="anchor"></a>C Specification

To clear one or more regions of color and depth/stencil attachments
inside a render pass instance, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdClearAttachments(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    attachmentCount,
    const VkClearAttachment*                    pAttachments,
    uint32_t                                    rectCount,
    const VkClearRect*                          pRects);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `attachmentCount` is the number of entries in the `pAttachments`
  array.

- `pAttachments` is a pointer to an array of
  [VkClearAttachment](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearAttachment.html) structures defining the
  attachments to clear and the clear values to use.

- `rectCount` is the number of entries in the `pRects` array.

- `pRects` is a pointer to an array of [VkClearRect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearRect.html)
  structures defining regions within each selected attachment to clear.

## <a href="#_description" class="anchor"></a>Description

If the render pass has a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-fragmentdensitymapattachment"
target="_blank" rel="noopener">fragment density map attachment</a>,
clears follow the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragmentdensitymapops"
target="_blank" rel="noopener">operations of fragment density maps</a>
as if each clear region was a primitive which generates fragments. The
clear color is applied to all pixels inside each fragment’s area
regardless if the pixels lie outside of the clear region. Clears **may**
have a different set of supported fragment areas than draws.

Unlike other <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears"
target="_blank" rel="noopener">clear commands</a>,
[vkCmdClearAttachments](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearAttachments.html) is not a transfer
command. It performs its operations in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-order"
target="_blank" rel="noopener">rasterization order</a>. For color
attachments, the operations are executed as color attachment writes, by
the `VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT` stage. For
depth/stencil attachments, the operations are executed as <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth"
target="_blank" rel="noopener">depth writes</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-stencil"
target="_blank" rel="noopener">stencil writes</a> by the
`VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT` and
`VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT` stages.

`vkCmdClearAttachments` is not affected by the bound pipeline state.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is generally preferable to clear attachments by using the
<code>VK_ATTACHMENT_LOAD_OP_CLEAR</code> load operation at the start of
rendering, as it is more efficient on some implementations.</p></td>
</tr>
</tbody>
</table>

If any attachment’s `aspectMask` to be cleared is not backed by an image
view, the clear has no effect on that aspect.

If an attachment being cleared refers to an image view created with an
`aspectMask` equal to one of `VK_IMAGE_ASPECT_PLANE_0_BIT`,
`VK_IMAGE_ASPECT_PLANE_1_BIT` or `VK_IMAGE_ASPECT_PLANE_2_BIT`, it is
considered to be `VK_IMAGE_ASPECT_COLOR_BIT` for purposes of this
command, and **must** be cleared with the `VK_IMAGE_ASPECT_COLOR_BIT`
aspect as specified by <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#image-views-plane-promotion"
target="_blank" rel="noopener">image view creation</a>.

Valid Usage

- <a href="#VUID-vkCmdClearAttachments-aspectMask-07884"
  id="VUID-vkCmdClearAttachments-aspectMask-07884"></a>
  VUID-vkCmdClearAttachments-aspectMask-07884  
  If the current render pass instance does not use dynamic rendering,
  and the `aspectMask` member of any element of `pAttachments` contains
  `VK_IMAGE_ASPECT_DEPTH_BIT`, the current subpass instance’s
  depth-stencil attachment **must** be either `VK_ATTACHMENT_UNUSED` or
  the attachment `format` **must** contain a depth component

- <a href="#VUID-vkCmdClearAttachments-aspectMask-07885"
  id="VUID-vkCmdClearAttachments-aspectMask-07885"></a>
  VUID-vkCmdClearAttachments-aspectMask-07885  
  If the current render pass instance does not use dynamic rendering,
  and the `aspectMask` member of any element of `pAttachments` contains
  `VK_IMAGE_ASPECT_STENCIL_BIT`, the current subpass instance’s
  depth-stencil attachment **must** be either `VK_ATTACHMENT_UNUSED` or
  the attachment `format` **must** contain a stencil component

- <a href="#VUID-vkCmdClearAttachments-aspectMask-07271"
  id="VUID-vkCmdClearAttachments-aspectMask-07271"></a>
  VUID-vkCmdClearAttachments-aspectMask-07271  
  If the `aspectMask` member of any element of `pAttachments` contains
  `VK_IMAGE_ASPECT_COLOR_BIT`, the `colorAttachment` **must** be a valid
  color attachment index in the current render pass instance

- <a href="#VUID-vkCmdClearAttachments-rect-02682"
  id="VUID-vkCmdClearAttachments-rect-02682"></a>
  VUID-vkCmdClearAttachments-rect-02682  
  The `rect` member of each element of `pRects` **must** have an
  `extent.width` greater than `0`

- <a href="#VUID-vkCmdClearAttachments-rect-02683"
  id="VUID-vkCmdClearAttachments-rect-02683"></a>
  VUID-vkCmdClearAttachments-rect-02683  
  The `rect` member of each element of `pRects` **must** have an
  `extent.height` greater than `0`

- <a href="#VUID-vkCmdClearAttachments-pRects-00016"
  id="VUID-vkCmdClearAttachments-pRects-00016"></a>
  VUID-vkCmdClearAttachments-pRects-00016  
  The rectangular region specified by each element of `pRects` **must**
  be contained within the render area of the current render pass
  instance

- <a href="#VUID-vkCmdClearAttachments-pRects-06937"
  id="VUID-vkCmdClearAttachments-pRects-06937"></a>
  VUID-vkCmdClearAttachments-pRects-06937  
  The layers specified by each element of `pRects` **must** be contained
  within every attachment that `pAttachments` refers to, i.e. for each
  element of `pRects`,
  [VkClearRect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearRect.html)::`baseArrayLayer` +
  [VkClearRect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearRect.html)::`layerCount` **must** be less than or
  equal to the number of layers rendered to in the current render pass
  instance

- <a href="#VUID-vkCmdClearAttachments-layerCount-01934"
  id="VUID-vkCmdClearAttachments-layerCount-01934"></a>
  VUID-vkCmdClearAttachments-layerCount-01934  
  The `layerCount` member of each element of `pRects` **must** not be
  `0`

- <a href="#VUID-vkCmdClearAttachments-commandBuffer-02504"
  id="VUID-vkCmdClearAttachments-commandBuffer-02504"></a>
  VUID-vkCmdClearAttachments-commandBuffer-02504  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, each attachment to be cleared **must** not be a
  protected image

- <a href="#VUID-vkCmdClearAttachments-commandBuffer-02505"
  id="VUID-vkCmdClearAttachments-commandBuffer-02505"></a>
  VUID-vkCmdClearAttachments-commandBuffer-02505  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, each attachment to be cleared **must** not be an
  unprotected image

- <a href="#VUID-vkCmdClearAttachments-baseArrayLayer-00018"
  id="VUID-vkCmdClearAttachments-baseArrayLayer-00018"></a>
  VUID-vkCmdClearAttachments-baseArrayLayer-00018  
  If the render pass instance this is recorded in uses multiview, then
  `baseArrayLayer` **must** be zero and `layerCount` **must** be one

- <a href="#VUID-vkCmdClearAttachments-colorAttachment-09503"
  id="VUID-vkCmdClearAttachments-colorAttachment-09503"></a>
  VUID-vkCmdClearAttachments-colorAttachment-09503  
  The `colorAttachment` member of each element of `pAttachments`
  **must** not identify a color attachment that is currently mapped to
  `VK_ATTACHMENT_UNUSED` in `commandBuffer` via
  [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentLocationInfoKHR.html)

- <a href="#VUID-vkCmdClearAttachments-aspectMask-09298"
  id="VUID-vkCmdClearAttachments-aspectMask-09298"></a>
  VUID-vkCmdClearAttachments-aspectMask-09298  
  If the subpass this is recorded in performs an external format
  resolve, the `aspectMask` member of any element of `pAttachments`
  **must** not include `VK_IMAGE_ASPECT_PLANE`*`_i_`*`BIT` for any index
  *i*

- <a href="#VUID-vkCmdClearAttachments-None-09679"
  id="VUID-vkCmdClearAttachments-None-09679"></a>
  VUID-vkCmdClearAttachments-None-09679  
  If the attachment format has components other than R and G, it
  **must** not have a 64-bit component width

Valid Usage (Implicit)

- <a href="#VUID-vkCmdClearAttachments-commandBuffer-parameter"
  id="VUID-vkCmdClearAttachments-commandBuffer-parameter"></a>
  VUID-vkCmdClearAttachments-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdClearAttachments-pAttachments-parameter"
  id="VUID-vkCmdClearAttachments-pAttachments-parameter"></a>
  VUID-vkCmdClearAttachments-pAttachments-parameter  
  `pAttachments` **must** be a valid pointer to an array of
  `attachmentCount` valid [VkClearAttachment](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearAttachment.html)
  structures

- <a href="#VUID-vkCmdClearAttachments-pRects-parameter"
  id="VUID-vkCmdClearAttachments-pRects-parameter"></a>
  VUID-vkCmdClearAttachments-pRects-parameter  
  `pRects` **must** be a valid pointer to an array of `rectCount`
  [VkClearRect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearRect.html) structures

- <a href="#VUID-vkCmdClearAttachments-commandBuffer-recording"
  id="VUID-vkCmdClearAttachments-commandBuffer-recording"></a>
  VUID-vkCmdClearAttachments-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdClearAttachments-commandBuffer-cmdpool"
  id="VUID-vkCmdClearAttachments-commandBuffer-cmdpool"></a>
  VUID-vkCmdClearAttachments-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdClearAttachments-renderpass"
  id="VUID-vkCmdClearAttachments-renderpass"></a>
  VUID-vkCmdClearAttachments-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdClearAttachments-videocoding"
  id="VUID-vkCmdClearAttachments-videocoding"></a>
  VUID-vkCmdClearAttachments-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdClearAttachments-attachmentCount-arraylength"
  id="VUID-vkCmdClearAttachments-attachmentCount-arraylength"></a>
  VUID-vkCmdClearAttachments-attachmentCount-arraylength  
  `attachmentCount` **must** be greater than `0`

- <a href="#VUID-vkCmdClearAttachments-rectCount-arraylength"
  id="VUID-vkCmdClearAttachments-rectCount-arraylength"></a>
  VUID-vkCmdClearAttachments-rectCount-arraylength  
  `rectCount` **must** be greater than `0`

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
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Inside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkClearAttachment](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearAttachment.html),
[VkClearRect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearRect.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdClearAttachments"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

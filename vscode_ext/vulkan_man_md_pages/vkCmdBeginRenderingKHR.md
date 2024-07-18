# vkCmdBeginRendering(3) Manual Page

## Name

vkCmdBeginRendering - Begin a dynamic render pass instance



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin a render pass instance, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdBeginRendering(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingInfo*                      pRenderingInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_dynamic_rendering
void vkCmdBeginRenderingKHR(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingInfo*                      pRenderingInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pRenderingInfo` is a pointer to a
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) structure specifying details
  of the render pass instance to begin.

## <a href="#_description" class="anchor"></a>Description

After beginning a render pass instance, the command buffer is ready to
record <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing"
target="_blank" rel="noopener">draw commands</a>.

If `pRenderingInfo->flags` includes `VK_RENDERING_RESUMING_BIT` then
this render pass is resumed from a render pass instance that has been
suspended earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

Valid Usage

- <a href="#VUID-vkCmdBeginRendering-dynamicRendering-06446"
  id="VUID-vkCmdBeginRendering-dynamicRendering-06446"></a>
  VUID-vkCmdBeginRendering-dynamicRendering-06446  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRendering"
  target="_blank" rel="noopener"><code>dynamicRendering</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdBeginRendering-commandBuffer-06068"
  id="VUID-vkCmdBeginRendering-commandBuffer-06068"></a>
  VUID-vkCmdBeginRendering-commandBuffer-06068  
  If `commandBuffer` is a secondary command buffer, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
  target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>
  feature is not enabled, `pRenderingInfo->flags` **must** not include
  `VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT`

- <a href="#VUID-vkCmdBeginRendering-pRenderingInfo-09588"
  id="VUID-vkCmdBeginRendering-pRenderingInfo-09588"></a>
  VUID-vkCmdBeginRendering-pRenderingInfo-09588  
  If `pRenderingInfo->pDepthAttachment` is not `NULL` and
  `pRenderingInfo->pDepthAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pRenderingInfo->pDepthAttachment->imageView` **must** be in the
  layout specified by `pRenderingInfo->pDepthAttachment->imageLayout`

- <a href="#VUID-vkCmdBeginRendering-pRenderingInfo-09589"
  id="VUID-vkCmdBeginRendering-pRenderingInfo-09589"></a>
  VUID-vkCmdBeginRendering-pRenderingInfo-09589  
  If `pRenderingInfo->pDepthAttachment` is not `NULL`,
  `pRenderingInfo->pDepthAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pRenderingInfo->pDepthAttachment->imageResolveMode` is not
  `VK_RESOLVE_MODE_NONE`, and
  `pRenderingInfo->pDepthAttachment->resolveImageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pRenderingInfo->pDepthAttachment->resolveImageView` **must** be in
  the layout specified by
  `pRenderingInfo->pDepthAttachment->resolveImageLayout`

- <a href="#VUID-vkCmdBeginRendering-pRenderingInfo-09590"
  id="VUID-vkCmdBeginRendering-pRenderingInfo-09590"></a>
  VUID-vkCmdBeginRendering-pRenderingInfo-09590  
  If `pRenderingInfo->pStencilAttachment` is not `NULL` and
  `pRenderingInfo->pStencilAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pRenderingInfo->pStencilAttachment->imageView` **must** be in the
  layout specified by `pRenderingInfo->pStencilAttachment->imageLayout`

- <a href="#VUID-vkCmdBeginRendering-pRenderingInfo-09591"
  id="VUID-vkCmdBeginRendering-pRenderingInfo-09591"></a>
  VUID-vkCmdBeginRendering-pRenderingInfo-09591  
  If `pRenderingInfo->pStencilAttachment` is not `NULL`,
  `pRenderingInfo->pStencilAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pRenderingInfo->pStencilAttachment->imageResolveMode` is not
  `VK_RESOLVE_MODE_NONE`, and
  `pRenderingInfo->pStencilAttachment->resolveImageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pRenderingInfo->pStencilAttachment->resolveImageView` **must** be in
  the layout specified by
  `pRenderingInfo->pStencilAttachment->resolveImageLayout`

- <a href="#VUID-vkCmdBeginRendering-pRenderingInfo-09592"
  id="VUID-vkCmdBeginRendering-pRenderingInfo-09592"></a>
  VUID-vkCmdBeginRendering-pRenderingInfo-09592  
  For any element of `pRenderingInfo->pColorAttachments`, if `imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), that image view **must**
  be in the layout specified by `imageLayout`

- <a href="#VUID-vkCmdBeginRendering-pRenderingInfo-09593"
  id="VUID-vkCmdBeginRendering-pRenderingInfo-09593"></a>
  VUID-vkCmdBeginRendering-pRenderingInfo-09593  
  For any element of `pRenderingInfo->pColorAttachments`, if either
  `imageResolveMode` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, or `imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not
  `VK_RESOLVE_MODE_NONE`, and `resolveImageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `resolveImageView` **must** be
  in the layout specified by `resolveImageLayout`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBeginRendering-commandBuffer-parameter"
  id="VUID-vkCmdBeginRendering-commandBuffer-parameter"></a>
  VUID-vkCmdBeginRendering-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBeginRendering-pRenderingInfo-parameter"
  id="VUID-vkCmdBeginRendering-pRenderingInfo-parameter"></a>
  VUID-vkCmdBeginRendering-pRenderingInfo-parameter  
  `pRenderingInfo` **must** be a valid pointer to a valid
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) structure

- <a href="#VUID-vkCmdBeginRendering-commandBuffer-recording"
  id="VUID-vkCmdBeginRendering-commandBuffer-recording"></a>
  VUID-vkCmdBeginRendering-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBeginRendering-commandBuffer-cmdpool"
  id="VUID-vkCmdBeginRendering-commandBuffer-cmdpool"></a>
  VUID-vkCmdBeginRendering-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBeginRendering-renderpass"
  id="VUID-vkCmdBeginRendering-renderpass"></a>
  VUID-vkCmdBeginRendering-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBeginRendering-videocoding"
  id="VUID-vkCmdBeginRendering-videocoding"></a>
  VUID-vkCmdBeginRendering-videocoding  
  This command **must** only be called outside of a video coding scope

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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBeginRendering"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# vkCmdPipelineBarrier2(3) Manual Page

## Name

vkCmdPipelineBarrier2 - Insert a memory dependency



## <a href="#_c_specification" class="anchor"></a>C Specification

To record a pipeline barrier, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdPipelineBarrier2(
    VkCommandBuffer                             commandBuffer,
    const VkDependencyInfo*                     pDependencyInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_synchronization2
void vkCmdPipelineBarrier2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkDependencyInfo*                     pDependencyInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `pDependencyInfo` is a pointer to a
  [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html) structure defining the
  scopes of this operation.

## <a href="#_description" class="anchor"></a>Description

When [vkCmdPipelineBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier2.html) is submitted to
a queue, it defines memory dependencies between commands that were
submitted to the same queue before it, and those submitted to the same
queue after it.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> of each memory
dependency defined by `pDependencyInfo` are applied to operations that
occurred earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> of each memory
dependency defined by `pDependencyInfo` are applied to operations that
occurred later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

If `vkCmdPipelineBarrier2` is recorded within a render pass instance,
the synchronization scopes are limited to a subset of operations within
the same subpass or render pass instance.

Valid Usage

- <a href="#VUID-vkCmdPipelineBarrier2-None-07889"
  id="VUID-vkCmdPipelineBarrier2-None-07889"></a>
  VUID-vkCmdPipelineBarrier2-None-07889  
  If `vkCmdPipelineBarrier2` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, the render pass
  **must** have been created with at least one subpass dependency that
  expresses a dependency from the current subpass to itself, does not
  include `VK_DEPENDENCY_BY_REGION_BIT` if this command does not, does
  not include `VK_DEPENDENCY_VIEW_LOCAL_BIT` if this command does not,
  and has [synchronization scopes](#synchronization-dependencies-scopes)
  and [access scopes](#synchronization-dependencies-access-scopes) that
  are all supersets of the scopes defined in this command

- <a href="#VUID-vkCmdPipelineBarrier2-bufferMemoryBarrierCount-01178"
  id="VUID-vkCmdPipelineBarrier2-bufferMemoryBarrierCount-01178"></a>
  VUID-vkCmdPipelineBarrier2-bufferMemoryBarrierCount-01178  
  If `vkCmdPipelineBarrier2` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, it **must** not
  include any buffer memory barriers

- <a href="#VUID-vkCmdPipelineBarrier2-image-04073"
  id="VUID-vkCmdPipelineBarrier2-image-04073"></a>
  VUID-vkCmdPipelineBarrier2-image-04073  
  If `vkCmdPipelineBarrier2` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, the `image` member
  of any image memory barrier included in this command **must** be an
  attachment used in the current subpass both as an input attachment,
  and as either a color, color resolve, or depth/stencil attachment

- <a href="#VUID-vkCmdPipelineBarrier2-image-09373"
  id="VUID-vkCmdPipelineBarrier2-image-09373"></a>
  VUID-vkCmdPipelineBarrier2-image-09373  
  If `vkCmdPipelineBarrier2` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, and the `image`
  member of any image memory barrier is a color resolve attachment, the
  corresponding color attachment **must** be `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-vkCmdPipelineBarrier2-image-09374"
  id="VUID-vkCmdPipelineBarrier2-image-09374"></a>
  VUID-vkCmdPipelineBarrier2-image-09374  
  If `vkCmdPipelineBarrier2` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, and the `image`
  member of any image memory barrier is a color resolve attachment, it
  **must** have been created with a non-zero
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value

- <a href="#VUID-vkCmdPipelineBarrier2-oldLayout-01181"
  id="VUID-vkCmdPipelineBarrier2-oldLayout-01181"></a>
  VUID-vkCmdPipelineBarrier2-oldLayout-01181  
  If `vkCmdPipelineBarrier2` is called within a render pass instance,
  the `oldLayout` and `newLayout` members of any image memory barrier
  included in this command **must** be equal

- <a href="#VUID-vkCmdPipelineBarrier2-srcQueueFamilyIndex-01182"
  id="VUID-vkCmdPipelineBarrier2-srcQueueFamilyIndex-01182"></a>
  VUID-vkCmdPipelineBarrier2-srcQueueFamilyIndex-01182  
  If `vkCmdPipelineBarrier2` is called within a render pass instance,
  the `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members of any
  memory barrier included in this command **must** be equal

- <a href="#VUID-vkCmdPipelineBarrier2-None-07890"
  id="VUID-vkCmdPipelineBarrier2-None-07890"></a>
  VUID-vkCmdPipelineBarrier2-None-07890  
  If `vkCmdPipelineBarrier2` is called within a render pass instance,
  and the source stage masks of any memory barriers include
  [framebuffer-space stages](#synchronization-framebuffer-regions),
  destination stage masks of all memory barriers **must** only include
  [framebuffer-space stages](#synchronization-framebuffer-regions)

- <a href="#VUID-vkCmdPipelineBarrier2-dependencyFlags-07891"
  id="VUID-vkCmdPipelineBarrier2-dependencyFlags-07891"></a>
  VUID-vkCmdPipelineBarrier2-dependencyFlags-07891  
  If `vkCmdPipelineBarrier2` is called within a render pass instance,
  and the source stage masks of any memory barriers include
  [framebuffer-space stages](#synchronization-framebuffer-regions), then
  `dependencyFlags` **must** include `VK_DEPENDENCY_BY_REGION_BIT`

- <a href="#VUID-vkCmdPipelineBarrier2-None-07892"
  id="VUID-vkCmdPipelineBarrier2-None-07892"></a>
  VUID-vkCmdPipelineBarrier2-None-07892  
  If `vkCmdPipelineBarrier2` is called within a render pass instance,
  the source and destination stage masks of any memory barriers **must**
  only include graphics pipeline stages

- <a href="#VUID-vkCmdPipelineBarrier2-dependencyFlags-01186"
  id="VUID-vkCmdPipelineBarrier2-dependencyFlags-01186"></a>
  VUID-vkCmdPipelineBarrier2-dependencyFlags-01186  
  If `vkCmdPipelineBarrier2` is called outside of a render pass
  instance, the dependency flags **must** not include
  `VK_DEPENDENCY_VIEW_LOCAL_BIT`

- <a href="#VUID-vkCmdPipelineBarrier2-None-07893"
  id="VUID-vkCmdPipelineBarrier2-None-07893"></a>
  VUID-vkCmdPipelineBarrier2-None-07893  
  If `vkCmdPipelineBarrier2` is called inside a render pass instance,
  and there is more than one view in the current subpass, dependency
  flags **must** include `VK_DEPENDENCY_VIEW_LOCAL_BIT`

- <a href="#VUID-vkCmdPipelineBarrier2-None-09553"
  id="VUID-vkCmdPipelineBarrier2-None-09553"></a>
  VUID-vkCmdPipelineBarrier2-None-09553  
  If none of the
  [`shaderTileImageColorReadAccess`](#features-shaderTileImageColorReadAccess),
  [`shaderTileImageStencilReadAccess`](#features-shaderTileImageStencilReadAccess),
  or
  [`shaderTileImageDepthReadAccess`](#features-shaderTileImageDepthReadAccess)
  features are enabled, and the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `vkCmdPipelineBarrier2` **must** not be called
  within a render pass instance started with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)

- <a href="#VUID-vkCmdPipelineBarrier2-None-09554"
  id="VUID-vkCmdPipelineBarrier2-None-09554"></a>
  VUID-vkCmdPipelineBarrier2-None-09554  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, and `vkCmdPipelineBarrier2` is called within a
  render pass instance started with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there **must** be no
  buffer or image memory barriers specified by this command

- <a href="#VUID-vkCmdPipelineBarrier2-None-09586"
  id="VUID-vkCmdPipelineBarrier2-None-09586"></a>
  VUID-vkCmdPipelineBarrier2-None-09586  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, and `vkCmdPipelineBarrier2` is called within a
  render pass instance started with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), memory barriers
  specified by this command **must** only include
  `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`,
  `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`,
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, or
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT` in their access masks

- <a href="#VUID-vkCmdPipelineBarrier2-image-09555"
  id="VUID-vkCmdPipelineBarrier2-image-09555"></a>
  VUID-vkCmdPipelineBarrier2-image-09555  
  If `vkCmdPipelineBarrier2` is called within a render pass instance
  started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), and the
  `image` member of any image memory barrier is used as an attachment in
  the current render pass instance, it **must** be in the
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR` or
  `VK_IMAGE_LAYOUT_GENERAL` layout

- <a href="#VUID-vkCmdPipelineBarrier2-srcStageMask-09556"
  id="VUID-vkCmdPipelineBarrier2-srcStageMask-09556"></a>
  VUID-vkCmdPipelineBarrier2-srcStageMask-09556  
  If `vkCmdPipelineBarrier2` is called within a render pass instance
  started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), this
  command **must** only specify [framebuffer-space
  stages](#synchronization-framebuffer-regions) in `srcStageMask` and
  `dstStageMask`

- <a href="#VUID-vkCmdPipelineBarrier2-synchronization2-03848"
  id="VUID-vkCmdPipelineBarrier2-synchronization2-03848"></a>
  VUID-vkCmdPipelineBarrier2-synchronization2-03848  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdPipelineBarrier2-srcStageMask-03849"
  id="VUID-vkCmdPipelineBarrier2-srcStageMask-03849"></a>
  VUID-vkCmdPipelineBarrier2-srcStageMask-03849  
  The `srcStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfo` **must** only include pipeline stages valid for the
  queue family that was used to create the command pool that
  `commandBuffer` was allocated from

- <a href="#VUID-vkCmdPipelineBarrier2-dstStageMask-03850"
  id="VUID-vkCmdPipelineBarrier2-dstStageMask-03850"></a>
  VUID-vkCmdPipelineBarrier2-dstStageMask-03850  
  The `dstStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfo` **must** only include pipeline stages valid for the
  queue family that was used to create the command pool that
  `commandBuffer` was allocated from

Valid Usage (Implicit)

- <a href="#VUID-vkCmdPipelineBarrier2-commandBuffer-parameter"
  id="VUID-vkCmdPipelineBarrier2-commandBuffer-parameter"></a>
  VUID-vkCmdPipelineBarrier2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdPipelineBarrier2-pDependencyInfo-parameter"
  id="VUID-vkCmdPipelineBarrier2-pDependencyInfo-parameter"></a>
  VUID-vkCmdPipelineBarrier2-pDependencyInfo-parameter  
  `pDependencyInfo` **must** be a valid pointer to a valid
  [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html) structure

- <a href="#VUID-vkCmdPipelineBarrier2-commandBuffer-recording"
  id="VUID-vkCmdPipelineBarrier2-commandBuffer-recording"></a>
  VUID-vkCmdPipelineBarrier2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdPipelineBarrier2-commandBuffer-cmdpool"
  id="VUID-vkCmdPipelineBarrier2-commandBuffer-cmdpool"></a>
  VUID-vkCmdPipelineBarrier2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, compute, decode, or encode operations

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
<tr class="header">
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
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute<br />
Decode<br />
Encode</p></td>
<td
class="tableblock halign-left valign-top"><p>Synchronization</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdPipelineBarrier2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

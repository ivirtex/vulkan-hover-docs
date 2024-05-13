# vkCmdClearDepthStencilImage(3) Manual Page

## Name

vkCmdClearDepthStencilImage - Fill regions of a combined depth/stencil
image



## <a href="#_c_specification" class="anchor"></a>C Specification

To clear one or more subranges of a depth/stencil image, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdClearDepthStencilImage(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     image,
    VkImageLayout                               imageLayout,
    const VkClearDepthStencilValue*             pDepthStencil,
    uint32_t                                    rangeCount,
    const VkImageSubresourceRange*              pRanges);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `image` is the image to be cleared.

- `imageLayout` specifies the current layout of the image subresource
  ranges to be cleared, and **must** be `VK_IMAGE_LAYOUT_GENERAL` or
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`.

- `pDepthStencil` is a pointer to a
  [VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearDepthStencilValue.html) structure
  containing the values that the depth and stencil image subresource
  ranges will be cleared to (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears-values"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears-values</a>
  below).

- `rangeCount` is the number of image subresource range structures in
  `pRanges`.

- `pRanges` is a pointer to an array of
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html) structures
  describing a range of mipmap levels, array layers, and aspects to be
  cleared, as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views"
  target="_blank" rel="noopener">Image Views</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCmdClearDepthStencilImage-image-01994"
  id="VUID-vkCmdClearDepthStencilImage-image-01994"></a>
  VUID-vkCmdClearDepthStencilImage-image-01994  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-format-features"
  target="_blank" rel="noopener">format features</a> of `image` **must**
  contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

- <a href="#VUID-vkCmdClearDepthStencilImage-pRanges-02658"
  id="VUID-vkCmdClearDepthStencilImage-pRanges-02658"></a>
  VUID-vkCmdClearDepthStencilImage-pRanges-02658  
  If the `aspect` member of any element of `pRanges` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, and `image` was created with
  <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html" target="_blank"
  rel="noopener">separate stencil usage</a>,
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  used to create `image`

- <a href="#VUID-vkCmdClearDepthStencilImage-pRanges-02659"
  id="VUID-vkCmdClearDepthStencilImage-pRanges-02659"></a>
  VUID-vkCmdClearDepthStencilImage-pRanges-02659  
  If the `aspect` member of any element of `pRanges` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, and `image` was not created with
  <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html" target="_blank"
  rel="noopener">separate stencil usage</a>,
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` used to create
  `image`

- <a href="#VUID-vkCmdClearDepthStencilImage-pRanges-02660"
  id="VUID-vkCmdClearDepthStencilImage-pRanges-02660"></a>
  VUID-vkCmdClearDepthStencilImage-pRanges-02660  
  If the `aspect` member of any element of `pRanges` includes
  `VK_IMAGE_ASPECT_DEPTH_BIT`, `VK_IMAGE_USAGE_TRANSFER_DST_BIT`
  **must** have been included in the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` used to create
  `image`

- <a href="#VUID-vkCmdClearDepthStencilImage-image-00010"
  id="VUID-vkCmdClearDepthStencilImage-image-00010"></a>
  VUID-vkCmdClearDepthStencilImage-image-00010  
  If `image` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdClearDepthStencilImage-imageLayout-00011"
  id="VUID-vkCmdClearDepthStencilImage-imageLayout-00011"></a>
  VUID-vkCmdClearDepthStencilImage-imageLayout-00011  
  `imageLayout` **must** specify the layout of the image subresource
  ranges of `image` specified in `pRanges` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdClearDepthStencilImage-imageLayout-00012"
  id="VUID-vkCmdClearDepthStencilImage-imageLayout-00012"></a>
  VUID-vkCmdClearDepthStencilImage-imageLayout-00012  
  `imageLayout` **must** be either of
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdClearDepthStencilImage-aspectMask-02824"
  id="VUID-vkCmdClearDepthStencilImage-aspectMask-02824"></a>
  VUID-vkCmdClearDepthStencilImage-aspectMask-02824  
  The
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`aspectMask`
  member of each element of the `pRanges` array **must** not include
  bits other than `VK_IMAGE_ASPECT_DEPTH_BIT` or
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-vkCmdClearDepthStencilImage-image-02825"
  id="VUID-vkCmdClearDepthStencilImage-image-02825"></a>
  VUID-vkCmdClearDepthStencilImage-image-02825  
  If the `image`’s format does not have a stencil component, then the
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`aspectMask`
  member of each element of the `pRanges` array **must** not include the
  `VK_IMAGE_ASPECT_STENCIL_BIT` bit

- <a href="#VUID-vkCmdClearDepthStencilImage-image-02826"
  id="VUID-vkCmdClearDepthStencilImage-image-02826"></a>
  VUID-vkCmdClearDepthStencilImage-image-02826  
  If the `image`’s format does not have a depth component, then the
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`aspectMask`
  member of each element of the `pRanges` array **must** not include the
  `VK_IMAGE_ASPECT_DEPTH_BIT` bit

- <a href="#VUID-vkCmdClearDepthStencilImage-baseMipLevel-01474"
  id="VUID-vkCmdClearDepthStencilImage-baseMipLevel-01474"></a>
  VUID-vkCmdClearDepthStencilImage-baseMipLevel-01474  
  The
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseMipLevel`
  members of the elements of the `pRanges` array **must** each be less
  than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-vkCmdClearDepthStencilImage-pRanges-01694"
  id="VUID-vkCmdClearDepthStencilImage-pRanges-01694"></a>
  VUID-vkCmdClearDepthStencilImage-pRanges-01694  
  For each [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)
  element of `pRanges`, if the `levelCount` member is not
  `VK_REMAINING_MIP_LEVELS`, then `baseMipLevel` + `levelCount` **must**
  be less than or equal to the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-vkCmdClearDepthStencilImage-baseArrayLayer-01476"
  id="VUID-vkCmdClearDepthStencilImage-baseArrayLayer-01476"></a>
  VUID-vkCmdClearDepthStencilImage-baseArrayLayer-01476  
  The
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseArrayLayer`
  members of the elements of the `pRanges` array **must** each be less
  than the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-vkCmdClearDepthStencilImage-pRanges-01695"
  id="VUID-vkCmdClearDepthStencilImage-pRanges-01695"></a>
  VUID-vkCmdClearDepthStencilImage-pRanges-01695  
  For each [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)
  element of `pRanges`, if the `layerCount` member is not
  `VK_REMAINING_ARRAY_LAYERS`, then `baseArrayLayer` + `layerCount`
  **must** be less than or equal to the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-vkCmdClearDepthStencilImage-image-00014"
  id="VUID-vkCmdClearDepthStencilImage-image-00014"></a>
  VUID-vkCmdClearDepthStencilImage-image-00014  
  `image` **must** have a depth/stencil format

- <a href="#VUID-vkCmdClearDepthStencilImage-commandBuffer-01807"
  id="VUID-vkCmdClearDepthStencilImage-commandBuffer-01807"></a>
  VUID-vkCmdClearDepthStencilImage-commandBuffer-01807  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, `image` **must** not be a protected image

- <a href="#VUID-vkCmdClearDepthStencilImage-commandBuffer-01808"
  id="VUID-vkCmdClearDepthStencilImage-commandBuffer-01808"></a>
  VUID-vkCmdClearDepthStencilImage-commandBuffer-01808  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, `image` **must** not be an unprotected image

Valid Usage (Implicit)

- <a href="#VUID-vkCmdClearDepthStencilImage-commandBuffer-parameter"
  id="VUID-vkCmdClearDepthStencilImage-commandBuffer-parameter"></a>
  VUID-vkCmdClearDepthStencilImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdClearDepthStencilImage-image-parameter"
  id="VUID-vkCmdClearDepthStencilImage-image-parameter"></a>
  VUID-vkCmdClearDepthStencilImage-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdClearDepthStencilImage-imageLayout-parameter"
  id="VUID-vkCmdClearDepthStencilImage-imageLayout-parameter"></a>
  VUID-vkCmdClearDepthStencilImage-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-vkCmdClearDepthStencilImage-pDepthStencil-parameter"
  id="VUID-vkCmdClearDepthStencilImage-pDepthStencil-parameter"></a>
  VUID-vkCmdClearDepthStencilImage-pDepthStencil-parameter  
  `pDepthStencil` **must** be a valid pointer to a valid
  [VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearDepthStencilValue.html) structure

- <a href="#VUID-vkCmdClearDepthStencilImage-pRanges-parameter"
  id="VUID-vkCmdClearDepthStencilImage-pRanges-parameter"></a>
  VUID-vkCmdClearDepthStencilImage-pRanges-parameter  
  `pRanges` **must** be a valid pointer to an array of `rangeCount`
  valid [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)
  structures

- <a href="#VUID-vkCmdClearDepthStencilImage-commandBuffer-recording"
  id="VUID-vkCmdClearDepthStencilImage-commandBuffer-recording"></a>
  VUID-vkCmdClearDepthStencilImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdClearDepthStencilImage-commandBuffer-cmdpool"
  id="VUID-vkCmdClearDepthStencilImage-commandBuffer-cmdpool"></a>
  VUID-vkCmdClearDepthStencilImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdClearDepthStencilImage-renderpass"
  id="VUID-vkCmdClearDepthStencilImage-renderpass"></a>
  VUID-vkCmdClearDepthStencilImage-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdClearDepthStencilImage-videocoding"
  id="VUID-vkCmdClearDepthStencilImage-videocoding"></a>
  VUID-vkCmdClearDepthStencilImage-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdClearDepthStencilImage-rangeCount-arraylength"
  id="VUID-vkCmdClearDepthStencilImage-rangeCount-arraylength"></a>
  VUID-vkCmdClearDepthStencilImage-rangeCount-arraylength  
  `rangeCount` **must** be greater than `0`

- <a href="#VUID-vkCmdClearDepthStencilImage-commonparent"
  id="VUID-vkCmdClearDepthStencilImage-commonparent"></a>
  VUID-vkCmdClearDepthStencilImage-commonparent  
  Both of `commandBuffer`, and `image` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearDepthStencilValue.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdClearDepthStencilImage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

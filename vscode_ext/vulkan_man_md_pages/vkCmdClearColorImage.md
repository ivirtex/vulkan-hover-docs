# vkCmdClearColorImage(3) Manual Page

## Name

vkCmdClearColorImage - Clear regions of a color image



## <a href="#_c_specification" class="anchor"></a>C Specification

To clear one or more subranges of a color image, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdClearColorImage(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     image,
    VkImageLayout                               imageLayout,
    const VkClearColorValue*                    pColor,
    uint32_t                                    rangeCount,
    const VkImageSubresourceRange*              pRanges);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `image` is the image to be cleared.

- `imageLayout` specifies the current layout of the image subresource
  ranges to be cleared, and **must** be
  `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_GENERAL` or
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`.

- `pColor` is a pointer to a [VkClearColorValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearColorValue.html)
  structure containing the values that the image subresource ranges will
  be cleared to (see <a
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

Each specified range in `pRanges` is cleared to the value specified by
`pColor`.

Valid Usage

- <a href="#VUID-vkCmdClearColorImage-image-01993"
  id="VUID-vkCmdClearColorImage-image-01993"></a>
  VUID-vkCmdClearColorImage-image-01993  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-format-features"
  target="_blank" rel="noopener">format features</a> of `image` **must**
  contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

- <a href="#VUID-vkCmdClearColorImage-image-00002"
  id="VUID-vkCmdClearColorImage-image-00002"></a>
  VUID-vkCmdClearColorImage-image-00002  
  `image` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdClearColorImage-image-01545"
  id="VUID-vkCmdClearColorImage-image-01545"></a>
  VUID-vkCmdClearColorImage-image-01545  
  `image` **must** not use any of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">formats that require a sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion</a>

- <a href="#VUID-vkCmdClearColorImage-image-00003"
  id="VUID-vkCmdClearColorImage-image-00003"></a>
  VUID-vkCmdClearColorImage-image-00003  
  If `image` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdClearColorImage-imageLayout-00004"
  id="VUID-vkCmdClearColorImage-imageLayout-00004"></a>
  VUID-vkCmdClearColorImage-imageLayout-00004  
  `imageLayout` **must** specify the layout of the image subresource
  ranges of `image` specified in `pRanges` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdClearColorImage-imageLayout-01394"
  id="VUID-vkCmdClearColorImage-imageLayout-01394"></a>
  VUID-vkCmdClearColorImage-imageLayout-01394  
  `imageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdClearColorImage-aspectMask-02498"
  id="VUID-vkCmdClearColorImage-aspectMask-02498"></a>
  VUID-vkCmdClearColorImage-aspectMask-02498  
  The
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`aspectMask`
  members of the elements of the `pRanges` array **must** each only
  include `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-vkCmdClearColorImage-baseMipLevel-01470"
  id="VUID-vkCmdClearColorImage-baseMipLevel-01470"></a>
  VUID-vkCmdClearColorImage-baseMipLevel-01470  
  The
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseMipLevel`
  members of the elements of the `pRanges` array **must** each be less
  than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-vkCmdClearColorImage-pRanges-01692"
  id="VUID-vkCmdClearColorImage-pRanges-01692"></a>
  VUID-vkCmdClearColorImage-pRanges-01692  
  For each [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)
  element of `pRanges`, if the `levelCount` member is not
  `VK_REMAINING_MIP_LEVELS`, then `baseMipLevel` + `levelCount` **must**
  be less than or equal to the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-vkCmdClearColorImage-baseArrayLayer-01472"
  id="VUID-vkCmdClearColorImage-baseArrayLayer-01472"></a>
  VUID-vkCmdClearColorImage-baseArrayLayer-01472  
  The
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseArrayLayer`
  members of the elements of the `pRanges` array **must** each be less
  than the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-vkCmdClearColorImage-pRanges-01693"
  id="VUID-vkCmdClearColorImage-pRanges-01693"></a>
  VUID-vkCmdClearColorImage-pRanges-01693  
  For each [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)
  element of `pRanges`, if the `layerCount` member is not
  `VK_REMAINING_ARRAY_LAYERS`, then `baseArrayLayer` + `layerCount`
  **must** be less than or equal to the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-vkCmdClearColorImage-image-00007"
  id="VUID-vkCmdClearColorImage-image-00007"></a>
  VUID-vkCmdClearColorImage-image-00007  
  `image` **must** not have a compressed or depth/stencil format

- <a href="#VUID-vkCmdClearColorImage-pColor-04961"
  id="VUID-vkCmdClearColorImage-pColor-04961"></a>
  VUID-vkCmdClearColorImage-pColor-04961  
  `pColor` **must** be a valid pointer to a
  [VkClearColorValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearColorValue.html) union

- <a href="#VUID-vkCmdClearColorImage-commandBuffer-01805"
  id="VUID-vkCmdClearColorImage-commandBuffer-01805"></a>
  VUID-vkCmdClearColorImage-commandBuffer-01805  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, `image` **must** not be a protected image

- <a href="#VUID-vkCmdClearColorImage-commandBuffer-01806"
  id="VUID-vkCmdClearColorImage-commandBuffer-01806"></a>
  VUID-vkCmdClearColorImage-commandBuffer-01806  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, **must** not be an unprotected image

Valid Usage (Implicit)

- <a href="#VUID-vkCmdClearColorImage-commandBuffer-parameter"
  id="VUID-vkCmdClearColorImage-commandBuffer-parameter"></a>
  VUID-vkCmdClearColorImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdClearColorImage-image-parameter"
  id="VUID-vkCmdClearColorImage-image-parameter"></a>
  VUID-vkCmdClearColorImage-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdClearColorImage-imageLayout-parameter"
  id="VUID-vkCmdClearColorImage-imageLayout-parameter"></a>
  VUID-vkCmdClearColorImage-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-vkCmdClearColorImage-pRanges-parameter"
  id="VUID-vkCmdClearColorImage-pRanges-parameter"></a>
  VUID-vkCmdClearColorImage-pRanges-parameter  
  `pRanges` **must** be a valid pointer to an array of `rangeCount`
  valid [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)
  structures

- <a href="#VUID-vkCmdClearColorImage-commandBuffer-recording"
  id="VUID-vkCmdClearColorImage-commandBuffer-recording"></a>
  VUID-vkCmdClearColorImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdClearColorImage-commandBuffer-cmdpool"
  id="VUID-vkCmdClearColorImage-commandBuffer-cmdpool"></a>
  VUID-vkCmdClearColorImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdClearColorImage-renderpass"
  id="VUID-vkCmdClearColorImage-renderpass"></a>
  VUID-vkCmdClearColorImage-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdClearColorImage-videocoding"
  id="VUID-vkCmdClearColorImage-videocoding"></a>
  VUID-vkCmdClearColorImage-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdClearColorImage-rangeCount-arraylength"
  id="VUID-vkCmdClearColorImage-rangeCount-arraylength"></a>
  VUID-vkCmdClearColorImage-rangeCount-arraylength  
  `rangeCount` **must** be greater than `0`

- <a href="#VUID-vkCmdClearColorImage-commonparent"
  id="VUID-vkCmdClearColorImage-commonparent"></a>
  VUID-vkCmdClearColorImage-commonparent  
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkClearColorValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearColorValue.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdClearColorImage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# vkCmdResolveImage(3) Manual Page

## Name

vkCmdResolveImage - Resolve regions of an image



## <a href="#_c_specification" class="anchor"></a>C Specification

To resolve a multisample color image to a non-multisample color image,
call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdResolveImage(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     srcImage,
    VkImageLayout                               srcImageLayout,
    VkImage                                     dstImage,
    VkImageLayout                               dstImageLayout,
    uint32_t                                    regionCount,
    const VkImageResolve*                       pRegions);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `srcImage` is the source image.

- `srcImageLayout` is the layout of the source image subresources for
  the resolve.

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the resolve.

- `regionCount` is the number of regions to resolve.

- `pRegions` is a pointer to an array of
  [VkImageResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve.html) structures specifying the
  regions to resolve.

## <a href="#_description" class="anchor"></a>Description

During the resolve the samples corresponding to each pixel location in
the source are converted to a single sample before being written to the
destination. If the source formats are floating-point or normalized
types, the sample values for each pixel are resolved in an
implementation-dependent manner. If the source formats are integer
types, a single sample’s value is selected for each pixel.

`srcOffset` and `dstOffset` select the initial `x`, `y`, and `z` offsets
in texels of the sub-regions of the source and destination image data.
`extent` is the size in texels of the source image to resolve in
`width`, `height` and `depth`. Each element of `pRegions` **must** be a
region that is contained within its corresponding image.

Resolves are done layer by layer starting with `baseArrayLayer` member
of `srcSubresource` for the source and `dstSubresource` for the
destination. `layerCount` layers are resolved to the destination image.

Valid Usage

- <a href="#VUID-vkCmdResolveImage-commandBuffer-01837"
  id="VUID-vkCmdResolveImage-commandBuffer-01837"></a>
  VUID-vkCmdResolveImage-commandBuffer-01837  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcImage` **must** not be a protected image

- <a href="#VUID-vkCmdResolveImage-commandBuffer-01838"
  id="VUID-vkCmdResolveImage-commandBuffer-01838"></a>
  VUID-vkCmdResolveImage-commandBuffer-01838  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be a protected image

- <a href="#VUID-vkCmdResolveImage-commandBuffer-01839"
  id="VUID-vkCmdResolveImage-commandBuffer-01839"></a>
  VUID-vkCmdResolveImage-commandBuffer-01839  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be an unprotected image

<!-- -->

- <a href="#VUID-vkCmdResolveImage-pRegions-00255"
  id="VUID-vkCmdResolveImage-pRegions-00255"></a>
  VUID-vkCmdResolveImage-pRegions-00255  
  The union of all source regions, and the union of all destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-vkCmdResolveImage-srcImage-00256"
  id="VUID-vkCmdResolveImage-srcImage-00256"></a>
  VUID-vkCmdResolveImage-srcImage-00256  
  If `srcImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdResolveImage-srcImage-00257"
  id="VUID-vkCmdResolveImage-srcImage-00257"></a>
  VUID-vkCmdResolveImage-srcImage-00257  
  `srcImage` **must** have a sample count equal to any valid sample
  count value other than `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-vkCmdResolveImage-dstImage-00258"
  id="VUID-vkCmdResolveImage-dstImage-00258"></a>
  VUID-vkCmdResolveImage-dstImage-00258  
  If `dstImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdResolveImage-dstImage-00259"
  id="VUID-vkCmdResolveImage-dstImage-00259"></a>
  VUID-vkCmdResolveImage-dstImage-00259  
  `dstImage` **must** have a sample count equal to
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-vkCmdResolveImage-srcImageLayout-00260"
  id="VUID-vkCmdResolveImage-srcImageLayout-00260"></a>
  VUID-vkCmdResolveImage-srcImageLayout-00260  
  `srcImageLayout` **must** specify the layout of the image subresources
  of `srcImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdResolveImage-srcImageLayout-01400"
  id="VUID-vkCmdResolveImage-srcImageLayout-01400"></a>
  VUID-vkCmdResolveImage-srcImageLayout-01400  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdResolveImage-dstImageLayout-00262"
  id="VUID-vkCmdResolveImage-dstImageLayout-00262"></a>
  VUID-vkCmdResolveImage-dstImageLayout-00262  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdResolveImage-dstImageLayout-01401"
  id="VUID-vkCmdResolveImage-dstImageLayout-01401"></a>
  VUID-vkCmdResolveImage-dstImageLayout-01401  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdResolveImage-dstImage-02003"
  id="VUID-vkCmdResolveImage-dstImage-02003"></a>
  VUID-vkCmdResolveImage-dstImage-02003  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-vkCmdResolveImage-linearColorAttachment-06519"
  id="VUID-vkCmdResolveImage-linearColorAttachment-06519"></a>
  VUID-vkCmdResolveImage-linearColorAttachment-06519  
  If the [`linearColorAttachment`](#features-linearColorAttachment)
  feature is enabled and the image is created with
  `VK_IMAGE_TILING_LINEAR`, the [format
  features](#resources-image-format-features) of `dstImage` **must**
  contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

- <a href="#VUID-vkCmdResolveImage-srcImage-01386"
  id="VUID-vkCmdResolveImage-srcImage-01386"></a>
  VUID-vkCmdResolveImage-srcImage-01386  
  `srcImage` and `dstImage` **must** have been created with the same
  image format

- <a href="#VUID-vkCmdResolveImage-srcSubresource-01709"
  id="VUID-vkCmdResolveImage-srcSubresource-01709"></a>
  VUID-vkCmdResolveImage-srcSubresource-01709  
  The `srcSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-vkCmdResolveImage-dstSubresource-01710"
  id="VUID-vkCmdResolveImage-dstSubresource-01710"></a>
  VUID-vkCmdResolveImage-dstSubresource-01710  
  The `dstSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-vkCmdResolveImage-srcSubresource-01711"
  id="VUID-vkCmdResolveImage-srcSubresource-01711"></a>
  VUID-vkCmdResolveImage-srcSubresource-01711  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-vkCmdResolveImage-dstSubresource-01712"
  id="VUID-vkCmdResolveImage-dstSubresource-01712"></a>
  VUID-vkCmdResolveImage-dstSubresource-01712  
  If `dstSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-vkCmdResolveImage-dstImage-02546"
  id="VUID-vkCmdResolveImage-dstImage-02546"></a>
  VUID-vkCmdResolveImage-dstImage-02546  
  `dstImage` and `srcImage` **must** not have been created with `flags`
  containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- <a href="#VUID-vkCmdResolveImage-srcImage-04446"
  id="VUID-vkCmdResolveImage-srcImage-04446"></a>
  VUID-vkCmdResolveImage-srcImage-04446  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `srcSubresource.layerCount` **must** be `1`

- <a href="#VUID-vkCmdResolveImage-srcImage-04447"
  id="VUID-vkCmdResolveImage-srcImage-04447"></a>
  VUID-vkCmdResolveImage-srcImage-04447  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `dstSubresource.baseArrayLayer` **must** be `0` and
  `dstSubresource.layerCount` **must** be `1`

- <a href="#VUID-vkCmdResolveImage-srcOffset-00269"
  id="VUID-vkCmdResolveImage-srcOffset-00269"></a>
  VUID-vkCmdResolveImage-srcOffset-00269  
  For each element of `pRegions`, `srcOffset.x` and (`extent.width` +
  `srcOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-vkCmdResolveImage-srcOffset-00270"
  id="VUID-vkCmdResolveImage-srcOffset-00270"></a>
  VUID-vkCmdResolveImage-srcOffset-00270  
  For each element of `pRegions`, `srcOffset.y` and (`extent.height` +
  `srcOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-vkCmdResolveImage-srcImage-00271"
  id="VUID-vkCmdResolveImage-srcImage-00271"></a>
  VUID-vkCmdResolveImage-srcImage-00271  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-vkCmdResolveImage-srcOffset-00272"
  id="VUID-vkCmdResolveImage-srcOffset-00272"></a>
  VUID-vkCmdResolveImage-srcOffset-00272  
  For each element of `pRegions`, `srcOffset.z` and (`extent.depth` +
  `srcOffset.z`) **must** both be greater than or equal to `0` and less
  than or equal to the depth of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-vkCmdResolveImage-srcImage-00273"
  id="VUID-vkCmdResolveImage-srcImage-00273"></a>
  VUID-vkCmdResolveImage-srcImage-00273  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `srcOffset.z` **must** be `0` and
  `extent.depth` **must** be `1`

- <a href="#VUID-vkCmdResolveImage-dstOffset-00274"
  id="VUID-vkCmdResolveImage-dstOffset-00274"></a>
  VUID-vkCmdResolveImage-dstOffset-00274  
  For each element of `pRegions`, `dstOffset.x` and (`extent.width` +
  `dstOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-vkCmdResolveImage-dstOffset-00275"
  id="VUID-vkCmdResolveImage-dstOffset-00275"></a>
  VUID-vkCmdResolveImage-dstOffset-00275  
  For each element of `pRegions`, `dstOffset.y` and (`extent.height` +
  `dstOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-vkCmdResolveImage-dstImage-00276"
  id="VUID-vkCmdResolveImage-dstImage-00276"></a>
  VUID-vkCmdResolveImage-dstImage-00276  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-vkCmdResolveImage-dstOffset-00277"
  id="VUID-vkCmdResolveImage-dstOffset-00277"></a>
  VUID-vkCmdResolveImage-dstOffset-00277  
  For each element of `pRegions`, `dstOffset.z` and (`extent.depth` +
  `dstOffset.z`) **must** both be greater than or equal to `0` and less
  than or equal to the depth of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-vkCmdResolveImage-dstImage-00278"
  id="VUID-vkCmdResolveImage-dstImage-00278"></a>
  VUID-vkCmdResolveImage-dstImage-00278  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `dstOffset.z` **must** be `0` and
  `extent.depth` **must** be `1`

- <a href="#VUID-vkCmdResolveImage-srcImage-06762"
  id="VUID-vkCmdResolveImage-srcImage-06762"></a>
  VUID-vkCmdResolveImage-srcImage-06762  
  `srcImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-vkCmdResolveImage-srcImage-06763"
  id="VUID-vkCmdResolveImage-srcImage-06763"></a>
  VUID-vkCmdResolveImage-srcImage-06763  
  The [format features](#resources-image-format-features) of `srcImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`

- <a href="#VUID-vkCmdResolveImage-dstImage-06764"
  id="VUID-vkCmdResolveImage-dstImage-06764"></a>
  VUID-vkCmdResolveImage-dstImage-06764  
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdResolveImage-dstImage-06765"
  id="VUID-vkCmdResolveImage-dstImage-06765"></a>
  VUID-vkCmdResolveImage-dstImage-06765  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdResolveImage-commandBuffer-parameter"
  id="VUID-vkCmdResolveImage-commandBuffer-parameter"></a>
  VUID-vkCmdResolveImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdResolveImage-srcImage-parameter"
  id="VUID-vkCmdResolveImage-srcImage-parameter"></a>
  VUID-vkCmdResolveImage-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdResolveImage-srcImageLayout-parameter"
  id="VUID-vkCmdResolveImage-srcImageLayout-parameter"></a>
  VUID-vkCmdResolveImage-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkCmdResolveImage-dstImage-parameter"
  id="VUID-vkCmdResolveImage-dstImage-parameter"></a>
  VUID-vkCmdResolveImage-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdResolveImage-dstImageLayout-parameter"
  id="VUID-vkCmdResolveImage-dstImageLayout-parameter"></a>
  VUID-vkCmdResolveImage-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkCmdResolveImage-pRegions-parameter"
  id="VUID-vkCmdResolveImage-pRegions-parameter"></a>
  VUID-vkCmdResolveImage-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkImageResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve.html) structures

- <a href="#VUID-vkCmdResolveImage-commandBuffer-recording"
  id="VUID-vkCmdResolveImage-commandBuffer-recording"></a>
  VUID-vkCmdResolveImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdResolveImage-commandBuffer-cmdpool"
  id="VUID-vkCmdResolveImage-commandBuffer-cmdpool"></a>
  VUID-vkCmdResolveImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdResolveImage-renderpass"
  id="VUID-vkCmdResolveImage-renderpass"></a>
  VUID-vkCmdResolveImage-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdResolveImage-videocoding"
  id="VUID-vkCmdResolveImage-videocoding"></a>
  VUID-vkCmdResolveImage-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdResolveImage-regionCount-arraylength"
  id="VUID-vkCmdResolveImage-regionCount-arraylength"></a>
  VUID-vkCmdResolveImage-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-vkCmdResolveImage-commonparent"
  id="VUID-vkCmdResolveImage-commonparent"></a>
  VUID-vkCmdResolveImage-commonparent  
  Each of `commandBuffer`, `dstImage`, and `srcImage` **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdResolveImage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

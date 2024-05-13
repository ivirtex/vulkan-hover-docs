# vkCmdCopyImageToBuffer(3) Manual Page

## Name

vkCmdCopyImageToBuffer - Copy image data into a buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data from an image object to a buffer object, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdCopyImageToBuffer(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     srcImage,
    VkImageLayout                               srcImageLayout,
    VkBuffer                                    dstBuffer,
    uint32_t                                    regionCount,
    const VkBufferImageCopy*                    pRegions);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `srcImage` is the source image.

- `srcImageLayout` is the layout of the source image subresources for
  the copy.

- `dstBuffer` is the destination buffer.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html) structures specifying the
  regions to copy.

## <a href="#_description" class="anchor"></a>Description

Each source region specified by `pRegions` is copied from the source
image to the destination region of the destination buffer according to
the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-buffers-images-addressing"
target="_blank" rel="noopener">addressing calculations</a> for each
resource. If any of the specified regions in `srcImage` overlaps in
memory with any of the specified regions in `dstBuffer`, values read
from those overlapping regions are undefined.

Copy regions for the image **must** be aligned to a multiple of the
texel block extent in each dimension, except at the edges of the image,
where region extents **must** match the edge of the image.

Valid Usage

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07966"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07966"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-vkCmdCopyImageToBuffer-imageSubresource-07967"
  id="VUID-vkCmdCopyImageToBuffer-imageSubresource-07967"></a>
  VUID-vkCmdCopyImageToBuffer-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-vkCmdCopyImageToBuffer-imageSubresource-07968"
  id="VUID-vkCmdCopyImageToBuffer-imageSubresource-07968"></a>
  VUID-vkCmdCopyImageToBuffer-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of
  each element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07969"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07969"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-vkCmdCopyImageToBuffer-imageSubresource-07970"
  id="VUID-vkCmdCopyImageToBuffer-imageSubresource-07970"></a>
  VUID-vkCmdCopyImageToBuffer-imageSubresource-07970  
  The image region specified by each element of `pRegions` **must** be
  contained within the specified `imageSubresource` of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-imageSubresource-07971"
  id="VUID-vkCmdCopyImageToBuffer-imageSubresource-07971"></a>
  VUID-vkCmdCopyImageToBuffer-imageSubresource-07971  
  For each element of `pRegions`, `imageOffset.x` and
  (`imageExtent.width` + `imageOffset.x`) **must** both be greater than
  or equal to `0` and less than or equal to the width of the specified
  `imageSubresource` of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-imageSubresource-07972"
  id="VUID-vkCmdCopyImageToBuffer-imageSubresource-07972"></a>
  VUID-vkCmdCopyImageToBuffer-imageSubresource-07972  
  For each element of `pRegions`, `imageOffset.y` and
  (`imageExtent.height` + `imageOffset.y`) **must** both be greater than
  or equal to `0` and less than or equal to the height of the specified
  `imageSubresource` of `srcImage`

<!-- -->

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07973"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07973"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07973  
  `srcImage` **must** have a sample count equal to
  `VK_SAMPLE_COUNT_1_BIT`

<!-- -->

- <a href="#VUID-vkCmdCopyImageToBuffer-commandBuffer-01831"
  id="VUID-vkCmdCopyImageToBuffer-commandBuffer-01831"></a>
  VUID-vkCmdCopyImageToBuffer-commandBuffer-01831  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcImage` **must** not be a protected image

- <a href="#VUID-vkCmdCopyImageToBuffer-commandBuffer-01832"
  id="VUID-vkCmdCopyImageToBuffer-commandBuffer-01832"></a>
  VUID-vkCmdCopyImageToBuffer-commandBuffer-01832  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdCopyImageToBuffer-commandBuffer-01833"
  id="VUID-vkCmdCopyImageToBuffer-commandBuffer-01833"></a>
  VUID-vkCmdCopyImageToBuffer-commandBuffer-01833  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstBuffer` **must** not be an unprotected buffer

- <a href="#VUID-vkCmdCopyImageToBuffer-commandBuffer-07746"
  id="VUID-vkCmdCopyImageToBuffer-commandBuffer-07746"></a>
  VUID-vkCmdCopyImageToBuffer-commandBuffer-07746  
  If the queue family used to create the
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) which `commandBuffer` was
  allocated from does not support `VK_QUEUE_GRAPHICS_BIT` or
  `VK_QUEUE_COMPUTE_BIT`, the `bufferOffset` member of any element of
  `pRegions` **must** be a multiple of `4`

- <a href="#VUID-vkCmdCopyImageToBuffer-imageOffset-07747"
  id="VUID-vkCmdCopyImageToBuffer-imageOffset-07747"></a>
  VUID-vkCmdCopyImageToBuffer-imageOffset-07747  
  The `imageOffset` and `imageExtent` members of each element of
  `pRegions` **must** respect the image transfer granularity
  requirements of `commandBuffer`’s command pool’s queue family, as
  described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

<!-- -->

- <a href="#VUID-vkCmdCopyImageToBuffer-pRegions-00183"
  id="VUID-vkCmdCopyImageToBuffer-pRegions-00183"></a>
  VUID-vkCmdCopyImageToBuffer-pRegions-00183  
  `dstBuffer` **must** be large enough to contain all buffer locations
  that are accessed according to [Buffer and Image
  Addressing](#copies-buffers-images-addressing), for each element of
  `pRegions`

- <a href="#VUID-vkCmdCopyImageToBuffer-pRegions-00184"
  id="VUID-vkCmdCopyImageToBuffer-pRegions-00184"></a>
  VUID-vkCmdCopyImageToBuffer-pRegions-00184  
  The union of all source regions, and the union of all destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-00186"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-00186"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-00186  
  `srcImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-01998"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-01998"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-01998  
  The [format features](#resources-image-format-features) of `srcImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`

- <a href="#VUID-vkCmdCopyImageToBuffer-dstBuffer-00191"
  id="VUID-vkCmdCopyImageToBuffer-dstBuffer-00191"></a>
  VUID-vkCmdCopyImageToBuffer-dstBuffer-00191  
  `dstBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdCopyImageToBuffer-dstBuffer-00192"
  id="VUID-vkCmdCopyImageToBuffer-dstBuffer-00192"></a>
  VUID-vkCmdCopyImageToBuffer-dstBuffer-00192  
  If `dstBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImageLayout-00189"
  id="VUID-vkCmdCopyImageToBuffer-srcImageLayout-00189"></a>
  VUID-vkCmdCopyImageToBuffer-srcImageLayout-00189  
  `srcImageLayout` **must** specify the layout of the image subresources
  of `srcImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImageLayout-01397"
  id="VUID-vkCmdCopyImageToBuffer-srcImageLayout-01397"></a>
  VUID-vkCmdCopyImageToBuffer-srcImageLayout-01397  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

<!-- -->

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07979"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07979"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07979  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height`
  **must** be `1`

- <a href="#VUID-vkCmdCopyImageToBuffer-imageOffset-09104"
  id="VUID-vkCmdCopyImageToBuffer-imageOffset-09104"></a>
  VUID-vkCmdCopyImageToBuffer-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and
  (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than
  or equal to `0` and less than or equal to the depth of the specified
  `imageSubresource` of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07980"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07980"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07980  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `imageOffset.z` **must** be `0`
  and `imageExtent.depth` **must** be `1`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07274"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07274"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07274  
  For each element of `pRegions`, `imageOffset.x` **must** be a multiple
  of the [texel block extent width](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07275"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07275"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07275  
  For each element of `pRegions`, `imageOffset.y` **must** be a multiple
  of the [texel block extent height](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07276"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07276"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-00207"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-00207"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-00207  
  For each element of `pRegions`, if the sum of `imageOffset.x` and
  `extent.width` does not equal the width of the subresource specified
  by `srcSubresource`, `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-00208"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-00208"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-00208  
  For each element of `pRegions`, if the sum of `imageOffset.y` and
  `extent.height` does not equal the height of the subresource specified
  by `srcSubresource`, `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-00209"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-00209"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-imageSubresource-09105"
  id="VUID-vkCmdCopyImageToBuffer-imageSubresource-09105"></a>
  VUID-vkCmdCopyImageToBuffer-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must**
  specify aspects present in `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07981"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07981"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07981  
  If `srcImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `imageSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07983"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07983"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07983  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, for each element of
  `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and
  `imageSubresource.layerCount` **must** be `1`

<!-- -->

- <a href="#VUID-vkCmdCopyImageToBuffer-bufferRowLength-09106"
  id="VUID-vkCmdCopyImageToBuffer-bufferRowLength-09106"></a>
  VUID-vkCmdCopyImageToBuffer-bufferRowLength-09106  
  For each element of `pRegions`, `bufferRowLength` **must** be a
  multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-bufferImageHeight-09107"
  id="VUID-vkCmdCopyImageToBuffer-bufferImageHeight-09107"></a>
  VUID-vkCmdCopyImageToBuffer-bufferImageHeight-09107  
  For each element of `pRegions`, `bufferImageHeight` **must** be a
  multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImageToBuffer-bufferRowLength-09108"
  id="VUID-vkCmdCopyImageToBuffer-bufferRowLength-09108"></a>
  VUID-vkCmdCopyImageToBuffer-bufferRowLength-09108  
  For each element of `pRegions`, `bufferRowLength` divided by the
  [texel block extent width](#formats-compatibility-classes) and then
  multiplied by the texel block size of `srcImage` **must** be less than
  or equal to 2<sup>31</sup>-1

<!-- -->

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07975"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07975"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07975  
  If `srcImage` does not have either a depth/stencil format or a
  [multi-planar format](#formats-requiring-sampler-ycbcr-conversion),
  then for each element of `pRegions`, `bufferOffset` **must** be a
  multiple of the [texel block size](#formats-compatibility-classes)

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07976"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07976"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07976  
  If `srcImage` has a [multi-planar
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `bufferOffset` **must** be a multiple of the
  element size of the compatible format for the format and the
  `aspectMask` of the `imageSubresource` as defined in
  [\[formats-compatible-planes\]](#formats-compatible-planes)

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-07978"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-07978"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-07978  
  If `srcImage` has a depth/stencil format, the `bufferOffset` member of
  any element of `pRegions` **must** be a multiple of `4`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyImageToBuffer-commandBuffer-parameter"
  id="VUID-vkCmdCopyImageToBuffer-commandBuffer-parameter"></a>
  VUID-vkCmdCopyImageToBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImage-parameter"
  id="VUID-vkCmdCopyImageToBuffer-srcImage-parameter"></a>
  VUID-vkCmdCopyImageToBuffer-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdCopyImageToBuffer-srcImageLayout-parameter"
  id="VUID-vkCmdCopyImageToBuffer-srcImageLayout-parameter"></a>
  VUID-vkCmdCopyImageToBuffer-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkCmdCopyImageToBuffer-dstBuffer-parameter"
  id="VUID-vkCmdCopyImageToBuffer-dstBuffer-parameter"></a>
  VUID-vkCmdCopyImageToBuffer-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdCopyImageToBuffer-pRegions-parameter"
  id="VUID-vkCmdCopyImageToBuffer-pRegions-parameter"></a>
  VUID-vkCmdCopyImageToBuffer-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html) structures

- <a href="#VUID-vkCmdCopyImageToBuffer-commandBuffer-recording"
  id="VUID-vkCmdCopyImageToBuffer-commandBuffer-recording"></a>
  VUID-vkCmdCopyImageToBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyImageToBuffer-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyImageToBuffer-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyImageToBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyImageToBuffer-renderpass"
  id="VUID-vkCmdCopyImageToBuffer-renderpass"></a>
  VUID-vkCmdCopyImageToBuffer-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyImageToBuffer-videocoding"
  id="VUID-vkCmdCopyImageToBuffer-videocoding"></a>
  VUID-vkCmdCopyImageToBuffer-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdCopyImageToBuffer-regionCount-arraylength"
  id="VUID-vkCmdCopyImageToBuffer-regionCount-arraylength"></a>
  VUID-vkCmdCopyImageToBuffer-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-vkCmdCopyImageToBuffer-commonparent"
  id="VUID-vkCmdCopyImageToBuffer-commonparent"></a>
  VUID-vkCmdCopyImageToBuffer-commonparent  
  Each of `commandBuffer`, `dstBuffer`, and `srcImage` **must** have
  been created, allocated, or retrieved from the same
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
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyImageToBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

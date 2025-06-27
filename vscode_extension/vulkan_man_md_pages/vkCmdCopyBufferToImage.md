# vkCmdCopyBufferToImage(3) Manual Page

## Name

vkCmdCopyBufferToImage - Copy data from a buffer into an image



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data from a buffer object to an image object, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdCopyBufferToImage(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    srcBuffer,
    VkImage                                     dstImage,
    VkImageLayout                               dstImageLayout,
    uint32_t                                    regionCount,
    const VkBufferImageCopy*                    pRegions);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `srcBuffer` is the source buffer.

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the copy.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html) structures specifying the
  regions to copy.

## <a href="#_description" class="anchor"></a>Description

Each source region specified by `pRegions` is copied from the source
buffer to the destination region of the destination image according to
the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-buffers-images-addressing"
target="_blank" rel="noopener">addressing calculations</a> for each
resource. If any of the specified regions in `srcBuffer` overlaps in
memory with any of the specified regions in `dstImage`, values read from
those overlapping regions are undefined. If any region accesses a depth
aspect in `dstImage` and the
[`VK_EXT_depth_range_unrestricted`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_range_unrestricted.html)
extension is not enabled, values copied from `srcBuffer` outside of the
range \[0,1\] will be written as undefined values to the destination
image.

Copy regions for the image **must** be aligned to a multiple of the
texel block extent in each dimension, except at the edges of the image,
where region extents **must** match the edge of the image.

Valid Usage

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07966"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07966"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-vkCmdCopyBufferToImage-imageSubresource-07967"
  id="VUID-vkCmdCopyBufferToImage-imageSubresource-07967"></a>
  VUID-vkCmdCopyBufferToImage-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-vkCmdCopyBufferToImage-imageSubresource-07968"
  id="VUID-vkCmdCopyBufferToImage-imageSubresource-07968"></a>
  VUID-vkCmdCopyBufferToImage-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of
  each element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07969"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07969"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-vkCmdCopyBufferToImage-imageSubresource-07970"
  id="VUID-vkCmdCopyBufferToImage-imageSubresource-07970"></a>
  VUID-vkCmdCopyBufferToImage-imageSubresource-07970  
  The image region specified by each element of `pRegions` **must** be
  contained within the specified `imageSubresource` of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageSubresource-07971"
  id="VUID-vkCmdCopyBufferToImage-imageSubresource-07971"></a>
  VUID-vkCmdCopyBufferToImage-imageSubresource-07971  
  For each element of `pRegions`, `imageOffset.x` and
  (`imageExtent.width` + `imageOffset.x`) **must** both be greater than
  or equal to `0` and less than or equal to the width of the specified
  `imageSubresource` of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageSubresource-07972"
  id="VUID-vkCmdCopyBufferToImage-imageSubresource-07972"></a>
  VUID-vkCmdCopyBufferToImage-imageSubresource-07972  
  For each element of `pRegions`, `imageOffset.y` and
  (`imageExtent.height` + `imageOffset.y`) **must** both be greater than
  or equal to `0` and less than or equal to the height of the specified
  `imageSubresource` of `dstImage`

<!-- -->

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07973"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07973"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07973  
  `dstImage` **must** have a sample count equal to
  `VK_SAMPLE_COUNT_1_BIT`

<!-- -->

- <a href="#VUID-vkCmdCopyBufferToImage-commandBuffer-01828"
  id="VUID-vkCmdCopyBufferToImage-commandBuffer-01828"></a>
  VUID-vkCmdCopyBufferToImage-commandBuffer-01828  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdCopyBufferToImage-commandBuffer-01829"
  id="VUID-vkCmdCopyBufferToImage-commandBuffer-01829"></a>
  VUID-vkCmdCopyBufferToImage-commandBuffer-01829  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be a protected image

- <a href="#VUID-vkCmdCopyBufferToImage-commandBuffer-01830"
  id="VUID-vkCmdCopyBufferToImage-commandBuffer-01830"></a>
  VUID-vkCmdCopyBufferToImage-commandBuffer-01830  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be an unprotected image

- <a href="#VUID-vkCmdCopyBufferToImage-commandBuffer-07737"
  id="VUID-vkCmdCopyBufferToImage-commandBuffer-07737"></a>
  VUID-vkCmdCopyBufferToImage-commandBuffer-07737  
  If the queue family used to create the
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) which `commandBuffer` was
  allocated from does not support `VK_QUEUE_GRAPHICS_BIT` or
  `VK_QUEUE_COMPUTE_BIT`, the `bufferOffset` member of any element of
  `pRegions` **must** be a multiple of `4`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-07738"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-07738"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-07738  
  The `imageOffset` and `imageExtent` members of each element of
  `pRegions` **must** respect the image transfer granularity
  requirements of `commandBuffer`’s command pool’s queue family, as
  described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

- <a href="#VUID-vkCmdCopyBufferToImage-commandBuffer-07739"
  id="VUID-vkCmdCopyBufferToImage-commandBuffer-07739"></a>
  VUID-vkCmdCopyBufferToImage-commandBuffer-07739  
  If the queue family used to create the
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) which `commandBuffer` was
  allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each
  element of `pRegions`, the `aspectMask` member of `imageSubresource`
  **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or
  `VK_IMAGE_ASPECT_STENCIL_BIT`

<!-- -->

- <a href="#VUID-vkCmdCopyBufferToImage-pRegions-00171"
  id="VUID-vkCmdCopyBufferToImage-pRegions-00171"></a>
  VUID-vkCmdCopyBufferToImage-pRegions-00171  
  `srcBuffer` **must** be large enough to contain all buffer locations
  that are accessed according to [Buffer and Image
  Addressing](#copies-buffers-images-addressing), for each element of
  `pRegions`

- <a href="#VUID-vkCmdCopyBufferToImage-pRegions-00173"
  id="VUID-vkCmdCopyBufferToImage-pRegions-00173"></a>
  VUID-vkCmdCopyBufferToImage-pRegions-00173  
  The union of all source regions, and the union of all destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-vkCmdCopyBufferToImage-srcBuffer-00174"
  id="VUID-vkCmdCopyBufferToImage-srcBuffer-00174"></a>
  VUID-vkCmdCopyBufferToImage-srcBuffer-00174  
  `srcBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-01997"
  id="VUID-vkCmdCopyBufferToImage-dstImage-01997"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-01997  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

- <a href="#VUID-vkCmdCopyBufferToImage-srcBuffer-00176"
  id="VUID-vkCmdCopyBufferToImage-srcBuffer-00176"></a>
  VUID-vkCmdCopyBufferToImage-srcBuffer-00176  
  If `srcBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-00177"
  id="VUID-vkCmdCopyBufferToImage-dstImage-00177"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-00177  
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdCopyBufferToImage-dstImageLayout-00180"
  id="VUID-vkCmdCopyBufferToImage-dstImageLayout-00180"></a>
  VUID-vkCmdCopyBufferToImage-dstImageLayout-00180  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImageLayout-01396"
  id="VUID-vkCmdCopyBufferToImage-dstImageLayout-01396"></a>
  VUID-vkCmdCopyBufferToImage-dstImageLayout-01396  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdCopyBufferToImage-pRegions-07931"
  id="VUID-vkCmdCopyBufferToImage-pRegions-07931"></a>
  VUID-vkCmdCopyBufferToImage-pRegions-07931  
  If
  [VK_EXT_depth_range_unrestricted](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_range_unrestricted.html)
  is not enabled, for each element of `pRegions` whose
  `imageSubresource` contains a depth aspect, the data in `srcBuffer`
  **must** be in the range \[0,1\]

<!-- -->

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07979"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07979"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07979  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height`
  **must** be `1`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-09104"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-09104"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and
  (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than
  or equal to `0` and less than or equal to the depth of the specified
  `imageSubresource` of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07980"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07980"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07980  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `imageOffset.z` **must** be `0`
  and `imageExtent.depth` **must** be `1`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07274"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07274"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07274  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `imageOffset.x` **must** be
  a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-10051"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-10051"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-10051  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `imageOffset.x` does not
  equal the width of the subresource specified by `imageSubresource`,
  `imageOffset.x` **must** be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07275"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07275"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07275  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `imageOffset.y` **must** be
  a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-10052"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-10052"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-10052  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `imageOffset.y` does
  not equal the height of the subresource specified by
  `imageSubresource`, `imageOffset.y` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07276"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07276"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-00207"
  id="VUID-vkCmdCopyBufferToImage-dstImage-00207"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-00207  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of
  `imageOffset.x` and `extent.width` does not equal the width of the
  subresource specified by `imageSubresource`, `extent.width` **must**
  be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-10053"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-10053"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-10053  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference
  of `imageOffset.x` and `extent.height` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-10054"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-10054"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-10054  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `imageOffset.x` and `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-10055"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-10055"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-10055  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of
  `imageOffset.x` and `extent.height` does not equal the width of the
  subresource specified by `imageSubresource`, `extent.height` **must**
  be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-00208"
  id="VUID-vkCmdCopyBufferToImage-dstImage-00208"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-00208  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of
  `imageOffset.y` and `extent.height` does not equal the height of the
  subresource specified by `imageSubresource`, `extent.height` **must**
  be a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-10056"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-10056"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-10056  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of
  `imageOffset.y` and `extent.width` does not equal the height of the
  subresource specified by `imageSubresource`, `extent.width` **must**
  be a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-10057"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-10057"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-10057  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `imageOffset.y` and `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageOffset-10058"
  id="VUID-vkCmdCopyBufferToImage-imageOffset-10058"></a>
  VUID-vkCmdCopyBufferToImage-imageOffset-10058  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference
  of `imageOffset.y` and `extent.width` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-00209"
  id="VUID-vkCmdCopyBufferToImage-dstImage-00209"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-imageSubresource-09105"
  id="VUID-vkCmdCopyBufferToImage-imageSubresource-09105"></a>
  VUID-vkCmdCopyBufferToImage-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must**
  specify aspects present in `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07981"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07981"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07981  
  If `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `imageSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07983"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07983"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07983  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, for each element of
  `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and
  `imageSubresource.layerCount` **must** be `1`

<!-- -->

- <a href="#VUID-vkCmdCopyBufferToImage-bufferRowLength-09106"
  id="VUID-vkCmdCopyBufferToImage-bufferRowLength-09106"></a>
  VUID-vkCmdCopyBufferToImage-bufferRowLength-09106  
  For each element of `pRegions`, `bufferRowLength` **must** be a
  multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-bufferImageHeight-09107"
  id="VUID-vkCmdCopyBufferToImage-bufferImageHeight-09107"></a>
  VUID-vkCmdCopyBufferToImage-bufferImageHeight-09107  
  For each element of `pRegions`, `bufferImageHeight` **must** be a
  multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyBufferToImage-bufferRowLength-09108"
  id="VUID-vkCmdCopyBufferToImage-bufferRowLength-09108"></a>
  VUID-vkCmdCopyBufferToImage-bufferRowLength-09108  
  For each element of `pRegions`, `bufferRowLength` divided by the
  [texel block extent width](#formats-compatibility-classes) and then
  multiplied by the texel block size of `dstImage` **must** be less than
  or equal to 2<sup>31</sup>-1

<!-- -->

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07975"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07975"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07975  
  If `dstImage` does not have either a depth/stencil format or a
  [multi-planar format](#formats-requiring-sampler-ycbcr-conversion),
  then for each element of `pRegions`, `bufferOffset` **must** be a
  multiple of the [texel block size](#formats-compatibility-classes)

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07976"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07976"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07976  
  If `dstImage` has a [multi-planar
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `bufferOffset` **must** be a multiple of the
  element size of the compatible format for the format and the
  `aspectMask` of the `imageSubresource` as defined in
  [\[formats-compatible-planes\]](#formats-compatible-planes)

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-07978"
  id="VUID-vkCmdCopyBufferToImage-dstImage-07978"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-07978  
  If `dstImage` has a depth/stencil format, the `bufferOffset` member of
  any element of `pRegions` **must** be a multiple of `4`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyBufferToImage-commandBuffer-parameter"
  id="VUID-vkCmdCopyBufferToImage-commandBuffer-parameter"></a>
  VUID-vkCmdCopyBufferToImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyBufferToImage-srcBuffer-parameter"
  id="VUID-vkCmdCopyBufferToImage-srcBuffer-parameter"></a>
  VUID-vkCmdCopyBufferToImage-srcBuffer-parameter  
  `srcBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdCopyBufferToImage-dstImage-parameter"
  id="VUID-vkCmdCopyBufferToImage-dstImage-parameter"></a>
  VUID-vkCmdCopyBufferToImage-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdCopyBufferToImage-dstImageLayout-parameter"
  id="VUID-vkCmdCopyBufferToImage-dstImageLayout-parameter"></a>
  VUID-vkCmdCopyBufferToImage-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkCmdCopyBufferToImage-pRegions-parameter"
  id="VUID-vkCmdCopyBufferToImage-pRegions-parameter"></a>
  VUID-vkCmdCopyBufferToImage-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html) structures

- <a href="#VUID-vkCmdCopyBufferToImage-commandBuffer-recording"
  id="VUID-vkCmdCopyBufferToImage-commandBuffer-recording"></a>
  VUID-vkCmdCopyBufferToImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyBufferToImage-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyBufferToImage-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyBufferToImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyBufferToImage-renderpass"
  id="VUID-vkCmdCopyBufferToImage-renderpass"></a>
  VUID-vkCmdCopyBufferToImage-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyBufferToImage-videocoding"
  id="VUID-vkCmdCopyBufferToImage-videocoding"></a>
  VUID-vkCmdCopyBufferToImage-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdCopyBufferToImage-regionCount-arraylength"
  id="VUID-vkCmdCopyBufferToImage-regionCount-arraylength"></a>
  VUID-vkCmdCopyBufferToImage-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-vkCmdCopyBufferToImage-commonparent"
  id="VUID-vkCmdCopyBufferToImage-commonparent"></a>
  VUID-vkCmdCopyBufferToImage-commonparent  
  Each of `commandBuffer`, `dstImage`, and `srcBuffer` **must** have
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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyBufferToImage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

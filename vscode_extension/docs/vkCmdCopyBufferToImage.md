# vkCmdCopyBufferToImage(3) Manual Page

## Name

vkCmdCopyBufferToImage - Copy data from a buffer into an image



## [](#_c_specification)C Specification

To copy data from a buffer object to an image object, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdCopyBufferToImage(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    srcBuffer,
    VkImage                                     dstImage,
    VkImageLayout                               dstImageLayout,
    uint32_t                                    regionCount,
    const VkBufferImageCopy*                    pRegions);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `srcBuffer` is the source buffer.
- `dstImage` is the destination image.
- `dstImageLayout` is the layout of the destination image subresources for the copy.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html) structures specifying the regions to copy.

## [](#_description)Description

Each source region specified by `pRegions` is copied from the source buffer to the destination region of the destination image according to the [addressing calculations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers-images-addressing) for each resource. If any of the specified regions in `srcBuffer` overlaps in memory with any of the specified regions in `dstImage`, values read from those overlapping regions are undefined. If any region accesses a depth aspect in `dstImage` and the `VK_EXT_depth_range_unrestricted` extension is not enabled, values copied from `srcBuffer` outside of the range \[0,1] will be written as undefined values to the destination image.

Copy regions for the image **must** be aligned to a multiple of the texel block extent in each dimension, except at the edges of the image, where region extents **must** match the edge of the image.

Valid Usage

- [](#VUID-vkCmdCopyBufferToImage-dstImage-07966)VUID-vkCmdCopyBufferToImage-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdCopyBufferToImage-imageSubresource-07967)VUID-vkCmdCopyBufferToImage-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-vkCmdCopyBufferToImage-imageSubresource-07968)VUID-vkCmdCopyBufferToImage-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07969)VUID-vkCmdCopyBufferToImage-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!--THE END-->

- [](#VUID-vkCmdCopyBufferToImage-imageSubresource-07971)VUID-vkCmdCopyBufferToImage-imageSubresource-07971  
  For each element of `pRegions`, `imageOffset.x` and (`imageExtent.width` + `imageOffset.x`) **must** both be greater than or equal to `0` and less than or equal to the width of the specified `imageSubresource` of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageSubresource-07972)VUID-vkCmdCopyBufferToImage-imageSubresource-07972  
  For each element of `pRegions`, `imageOffset.y` and (`imageExtent.height` + `imageOffset.y`) **must** both be greater than or equal to `0` and less than or equal to the height of the specified `imageSubresource` of `dstImage`

<!--THE END-->

- [](#VUID-vkCmdCopyBufferToImage-dstImage-07973)VUID-vkCmdCopyBufferToImage-dstImage-07973  
  `dstImage` **must** have a sample count equal to `VK_SAMPLE_COUNT_1_BIT`

<!--THE END-->

- [](#VUID-vkCmdCopyBufferToImage-commandBuffer-01828)VUID-vkCmdCopyBufferToImage-commandBuffer-01828  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `srcBuffer` **must** not be a protected buffer
- [](#VUID-vkCmdCopyBufferToImage-commandBuffer-01829)VUID-vkCmdCopyBufferToImage-commandBuffer-01829  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be a protected image
- [](#VUID-vkCmdCopyBufferToImage-commandBuffer-01830)VUID-vkCmdCopyBufferToImage-commandBuffer-01830  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be an unprotected image
- [](#VUID-vkCmdCopyBufferToImage-commandBuffer-07737)VUID-vkCmdCopyBufferToImage-commandBuffer-07737  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`, the `bufferOffset` member of any element of `pRegions` **must** be a multiple of `4`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-07738)VUID-vkCmdCopyBufferToImage-imageOffset-07738  
  The `imageOffset` and `imageExtent` members of each element of `pRegions` **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-vkCmdCopyBufferToImage-commandBuffer-07739)VUID-vkCmdCopyBufferToImage-commandBuffer-07739  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each element of `pRegions`, the `aspectMask` member of `imageSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

<!--THE END-->

- [](#VUID-vkCmdCopyBufferToImage-pRegions-00171)VUID-vkCmdCopyBufferToImage-pRegions-00171  
  `srcBuffer` **must** be large enough to contain all buffer locations that are accessed according to [Buffer and Image Addressing](#copies-buffers-images-addressing), for each element of `pRegions`
- [](#VUID-vkCmdCopyBufferToImage-pRegions-00173)VUID-vkCmdCopyBufferToImage-pRegions-00173  
  The union of all source regions, and the union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory
- [](#VUID-vkCmdCopyBufferToImage-srcBuffer-00174)VUID-vkCmdCopyBufferToImage-srcBuffer-00174  
  `srcBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` usage flag
- [](#VUID-vkCmdCopyBufferToImage-dstImage-01997)VUID-vkCmdCopyBufferToImage-dstImage-01997  
  The [format features](#resources-image-format-features) of `dstImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`
- [](#VUID-vkCmdCopyBufferToImage-srcBuffer-00176)VUID-vkCmdCopyBufferToImage-srcBuffer-00176  
  If `srcBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdCopyBufferToImage-dstImage-00177)VUID-vkCmdCopyBufferToImage-dstImage-00177  
  `dstImage` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdCopyBufferToImage-dstImageLayout-00180)VUID-vkCmdCopyBufferToImage-dstImageLayout-00180  
  `dstImageLayout` **must** specify the layout of the image subresources of `dstImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-vkCmdCopyBufferToImage-dstImageLayout-01396)VUID-vkCmdCopyBufferToImage-dstImageLayout-01396  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-vkCmdCopyBufferToImage-pRegions-07931)VUID-vkCmdCopyBufferToImage-pRegions-07931  
  If [VK\_EXT\_depth\_range\_unrestricted](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_range_unrestricted.html) is not enabled, for each element of `pRegions` whose `imageSubresource` contains a depth aspect, the data in `srcBuffer` **must** be in the range \[0,1]

<!--THE END-->

- [](#VUID-vkCmdCopyBufferToImage-dstImage-07979)VUID-vkCmdCopyBufferToImage-dstImage-07979  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height` **must** be `1`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-09104)VUID-vkCmdCopyBufferToImage-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `imageSubresource` of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07980)VUID-vkCmdCopyBufferToImage-dstImage-07980  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `imageOffset.z` **must** be `0` and `imageExtent.depth` **must** be `1`
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07274)VUID-vkCmdCopyBufferToImage-dstImage-07274  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-10051)VUID-vkCmdCopyBufferToImage-imageOffset-10051  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `imageOffset.x` does not equal the width of the subresource specified by `imageSubresource`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07275)VUID-vkCmdCopyBufferToImage-dstImage-07275  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-10052)VUID-vkCmdCopyBufferToImage-imageOffset-10052  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `imageOffset.y` does not equal the height of the subresource specified by `imageSubresource`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07276)VUID-vkCmdCopyBufferToImage-dstImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-dstImage-00207)VUID-vkCmdCopyBufferToImage-dstImage-00207  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of `imageOffset.x` and `extent.width` does not equal the width of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-10053)VUID-vkCmdCopyBufferToImage-imageOffset-10053  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference of `imageOffset.x` and `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-10054)VUID-vkCmdCopyBufferToImage-imageOffset-10054  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.x` and `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-10055)VUID-vkCmdCopyBufferToImage-imageOffset-10055  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of `imageOffset.x` and `extent.height` does not equal the width of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-dstImage-00208)VUID-vkCmdCopyBufferToImage-dstImage-00208  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of `imageOffset.y` and `extent.height` does not equal the height of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-10056)VUID-vkCmdCopyBufferToImage-imageOffset-10056  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of `imageOffset.y` and `extent.width` does not equal the height of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-10057)VUID-vkCmdCopyBufferToImage-imageOffset-10057  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.y` and `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageOffset-10058)VUID-vkCmdCopyBufferToImage-imageOffset-10058  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference of `imageOffset.y` and `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-dstImage-00209)VUID-vkCmdCopyBufferToImage-dstImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and `extent.depth` does not equal the depth of the subresource specified by `srcSubresource`, `extent.depth` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-imageSubresource-09105)VUID-vkCmdCopyBufferToImage-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must** specify aspects present in `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07981)VUID-vkCmdCopyBufferToImage-dstImage-07981  
  If `dstImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `imageSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07983)VUID-vkCmdCopyBufferToImage-dstImage-07983  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, for each element of `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and `imageSubresource.layerCount` **must** be `1`

<!--THE END-->

- [](#VUID-vkCmdCopyBufferToImage-bufferRowLength-09106)VUID-vkCmdCopyBufferToImage-bufferRowLength-09106  
  For each element of `pRegions`, `bufferRowLength` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-bufferImageHeight-09107)VUID-vkCmdCopyBufferToImage-bufferImageHeight-09107  
  For each element of `pRegions`, `bufferImageHeight` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyBufferToImage-bufferRowLength-09108)VUID-vkCmdCopyBufferToImage-bufferRowLength-09108  
  For each element of `pRegions`, `bufferRowLength` divided by the [texel block extent width](#formats-compatibility-classes) and then multiplied by the texel block size of `dstImage` **must** be less than or equal to 231-1

<!--THE END-->

- [](#VUID-vkCmdCopyBufferToImage-dstImage-07975)VUID-vkCmdCopyBufferToImage-dstImage-07975  
  If `dstImage` does not have either a depth/stencil format or a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `bufferOffset` **must** be a multiple of the [texel block size](#formats-compatibility-classes)
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07976)VUID-vkCmdCopyBufferToImage-dstImage-07976  
  If `dstImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `bufferOffset` **must** be a multiple of the element size of the compatible format for the format and the `aspectMask` of the `imageSubresource` as defined in [\[formats-compatible-planes\]](#formats-compatible-planes)
- [](#VUID-vkCmdCopyBufferToImage-dstImage-07978)VUID-vkCmdCopyBufferToImage-dstImage-07978  
  If `dstImage` has a depth/stencil format, the `bufferOffset` member of any element of `pRegions` **must** be a multiple of `4`

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyBufferToImage-commandBuffer-parameter)VUID-vkCmdCopyBufferToImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyBufferToImage-srcBuffer-parameter)VUID-vkCmdCopyBufferToImage-srcBuffer-parameter  
  `srcBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdCopyBufferToImage-dstImage-parameter)VUID-vkCmdCopyBufferToImage-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkCmdCopyBufferToImage-dstImageLayout-parameter)VUID-vkCmdCopyBufferToImage-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdCopyBufferToImage-pRegions-parameter)VUID-vkCmdCopyBufferToImage-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html) structures
- [](#VUID-vkCmdCopyBufferToImage-commandBuffer-recording)VUID-vkCmdCopyBufferToImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyBufferToImage-commandBuffer-cmdpool)VUID-vkCmdCopyBufferToImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdCopyBufferToImage-renderpass)VUID-vkCmdCopyBufferToImage-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyBufferToImage-videocoding)VUID-vkCmdCopyBufferToImage-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdCopyBufferToImage-regionCount-arraylength)VUID-vkCmdCopyBufferToImage-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-vkCmdCopyBufferToImage-commonparent)VUID-vkCmdCopyBufferToImage-commonparent  
  Each of `commandBuffer`, `dstImage`, and `srcBuffer` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Transfer  
Graphics  
Compute

Action

Conditional Rendering

vkCmdCopyBufferToImage is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyBufferToImage)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
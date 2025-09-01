# vkCmdCopyImageToBuffer(3) Manual Page

## Name

vkCmdCopyImageToBuffer - Copy image data into a buffer



## [](#_c_specification)C Specification

To copy data from an image object to a buffer object, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdCopyImageToBuffer(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     srcImage,
    VkImageLayout                               srcImageLayout,
    VkBuffer                                    dstBuffer,
    uint32_t                                    regionCount,
    const VkBufferImageCopy*                    pRegions);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `srcImage` is the source image.
- `srcImageLayout` is the layout of the source image subresources for the copy.
- `dstBuffer` is the destination buffer.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html) structures specifying the regions to copy.

## [](#_description)Description

Each source region specified by `pRegions` is copied from the source image to the destination region of the destination buffer according to the [addressing calculations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers-images-addressing) for each resource. If any of the specified regions in `srcImage` overlaps in memory with any of the specified regions in `dstBuffer`, values read from those overlapping regions are undefined.

Copy regions for the image **must** be aligned to a multiple of the texel block extent in each dimension, except at the edges of the image, where region extents **must** match the edge of the image.

Valid Usage

- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07966)VUID-vkCmdCopyImageToBuffer-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdCopyImageToBuffer-imageSubresource-07967)VUID-vkCmdCopyImageToBuffer-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-vkCmdCopyImageToBuffer-imageSubresource-07968)VUID-vkCmdCopyImageToBuffer-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07969)VUID-vkCmdCopyImageToBuffer-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!--THE END-->

- [](#VUID-vkCmdCopyImageToBuffer-imageSubresource-07971)VUID-vkCmdCopyImageToBuffer-imageSubresource-07971  
  For each element of `pRegions`, `imageOffset.x` and (`imageExtent.width` + `imageOffset.x`) **must** both be greater than or equal to `0` and less than or equal to the width of the specified `imageSubresource` of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageSubresource-07972)VUID-vkCmdCopyImageToBuffer-imageSubresource-07972  
  For each element of `pRegions`, `imageOffset.y` and (`imageExtent.height` + `imageOffset.y`) **must** both be greater than or equal to `0` and less than or equal to the height of the specified `imageSubresource` of `srcImage`

<!--THE END-->

- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07973)VUID-vkCmdCopyImageToBuffer-srcImage-07973  
  `srcImage` **must** have a sample count equal to `VK_SAMPLE_COUNT_1_BIT`

<!--THE END-->

- [](#VUID-vkCmdCopyImageToBuffer-commandBuffer-01831)VUID-vkCmdCopyImageToBuffer-commandBuffer-01831  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `srcImage` **must** not be a protected image
- [](#VUID-vkCmdCopyImageToBuffer-commandBuffer-01832)VUID-vkCmdCopyImageToBuffer-commandBuffer-01832  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be a protected buffer
- [](#VUID-vkCmdCopyImageToBuffer-commandBuffer-01833)VUID-vkCmdCopyImageToBuffer-commandBuffer-01833  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be an unprotected buffer
- [](#VUID-vkCmdCopyImageToBuffer-commandBuffer-07746)VUID-vkCmdCopyImageToBuffer-commandBuffer-07746  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`, the `bufferOffset` member of any element of `pRegions` **must** be a multiple of `4`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-07747)VUID-vkCmdCopyImageToBuffer-imageOffset-07747  
  The `imageOffset` and `imageExtent` members of each element of `pRegions` **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-vkCmdCopyImageToBuffer-commandBuffer-10216)VUID-vkCmdCopyImageToBuffer-commandBuffer-10216  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each element of `pRegions`, the `aspectMask` member of `imageSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

<!--THE END-->

- [](#VUID-vkCmdCopyImageToBuffer-pRegions-00183)VUID-vkCmdCopyImageToBuffer-pRegions-00183  
  `dstBuffer` **must** be large enough to contain all buffer locations that are accessed according to [Buffer and Image Addressing](#copies-buffers-images-addressing), for each element of `pRegions`
- [](#VUID-vkCmdCopyImageToBuffer-pRegions-00184)VUID-vkCmdCopyImageToBuffer-pRegions-00184  
  The union of all source regions, and the union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-00186)VUID-vkCmdCopyImageToBuffer-srcImage-00186  
  `srcImage` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage flag
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-01998)VUID-vkCmdCopyImageToBuffer-srcImage-01998  
  The [format features](#resources-image-format-features) of `srcImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`
- [](#VUID-vkCmdCopyImageToBuffer-dstBuffer-00191)VUID-vkCmdCopyImageToBuffer-dstBuffer-00191  
  `dstBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdCopyImageToBuffer-dstBuffer-00192)VUID-vkCmdCopyImageToBuffer-dstBuffer-00192  
  If `dstBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdCopyImageToBuffer-srcImageLayout-00189)VUID-vkCmdCopyImageToBuffer-srcImageLayout-00189  
  `srcImageLayout` **must** specify the layout of the image subresources of `srcImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-vkCmdCopyImageToBuffer-srcImageLayout-01397)VUID-vkCmdCopyImageToBuffer-srcImageLayout-01397  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

<!--THE END-->

- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07979)VUID-vkCmdCopyImageToBuffer-srcImage-07979  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height` **must** be `1`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-09104)VUID-vkCmdCopyImageToBuffer-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `imageSubresource` of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07980)VUID-vkCmdCopyImageToBuffer-srcImage-07980  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `imageOffset.z` **must** be `0` and `imageExtent.depth` **must** be `1`
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07274)VUID-vkCmdCopyImageToBuffer-srcImage-07274  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-10051)VUID-vkCmdCopyImageToBuffer-imageOffset-10051  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `imageOffset.x` does not equal the width of the subresource specified by `imageSubresource`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07275)VUID-vkCmdCopyImageToBuffer-srcImage-07275  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-10052)VUID-vkCmdCopyImageToBuffer-imageOffset-10052  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `imageOffset.y` does not equal the height of the subresource specified by `imageSubresource`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07276)VUID-vkCmdCopyImageToBuffer-srcImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-00207)VUID-vkCmdCopyImageToBuffer-srcImage-00207  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of `imageOffset.x` and `extent.width` does not equal the width of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-10053)VUID-vkCmdCopyImageToBuffer-imageOffset-10053  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference of `imageOffset.x` and `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-10054)VUID-vkCmdCopyImageToBuffer-imageOffset-10054  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.x` and `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-10055)VUID-vkCmdCopyImageToBuffer-imageOffset-10055  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of `imageOffset.x` and `extent.height` does not equal the width of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-00208)VUID-vkCmdCopyImageToBuffer-srcImage-00208  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of `imageOffset.y` and `extent.height` does not equal the height of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-10056)VUID-vkCmdCopyImageToBuffer-imageOffset-10056  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of `imageOffset.y` and `extent.width` does not equal the height of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-10057)VUID-vkCmdCopyImageToBuffer-imageOffset-10057  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.y` and `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageOffset-10058)VUID-vkCmdCopyImageToBuffer-imageOffset-10058  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference of `imageOffset.y` and `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-00209)VUID-vkCmdCopyImageToBuffer-srcImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and `extent.depth` does not equal the depth of the subresource specified by `srcSubresource`, `extent.depth` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-imageSubresource-09105)VUID-vkCmdCopyImageToBuffer-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must** specify aspects present in `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07981)VUID-vkCmdCopyImageToBuffer-srcImage-07981  
  If `srcImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `imageSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07983)VUID-vkCmdCopyImageToBuffer-srcImage-07983  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, for each element of `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and `imageSubresource.layerCount` **must** be `1`

<!--THE END-->

- [](#VUID-vkCmdCopyImageToBuffer-bufferRowLength-09106)VUID-vkCmdCopyImageToBuffer-bufferRowLength-09106  
  For each element of `pRegions`, `bufferRowLength` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-bufferImageHeight-09107)VUID-vkCmdCopyImageToBuffer-bufferImageHeight-09107  
  For each element of `pRegions`, `bufferImageHeight` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImageToBuffer-bufferRowLength-09108)VUID-vkCmdCopyImageToBuffer-bufferRowLength-09108  
  For each element of `pRegions`, `bufferRowLength` divided by the [texel block extent width](#formats-compatibility-classes) and then multiplied by the texel block size of `srcImage` **must** be less than or equal to 231-1

<!--THE END-->

- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07975)VUID-vkCmdCopyImageToBuffer-srcImage-07975  
  If `srcImage` does not have either a depth/stencil format or a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `bufferOffset` **must** be a multiple of the [texel block size](#formats-compatibility-classes)
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07976)VUID-vkCmdCopyImageToBuffer-srcImage-07976  
  If `srcImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `bufferOffset` **must** be a multiple of the element size of the compatible format for the format and the `aspectMask` of the `imageSubresource` as defined in [\[formats-compatible-planes\]](#formats-compatible-planes)
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-07978)VUID-vkCmdCopyImageToBuffer-srcImage-07978  
  If `srcImage` has a depth/stencil format, the `bufferOffset` member of any element of `pRegions` **must** be a multiple of `4`

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyImageToBuffer-commandBuffer-parameter)VUID-vkCmdCopyImageToBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyImageToBuffer-srcImage-parameter)VUID-vkCmdCopyImageToBuffer-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkCmdCopyImageToBuffer-srcImageLayout-parameter)VUID-vkCmdCopyImageToBuffer-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdCopyImageToBuffer-dstBuffer-parameter)VUID-vkCmdCopyImageToBuffer-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdCopyImageToBuffer-pRegions-parameter)VUID-vkCmdCopyImageToBuffer-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html) structures
- [](#VUID-vkCmdCopyImageToBuffer-commandBuffer-recording)VUID-vkCmdCopyImageToBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyImageToBuffer-commandBuffer-cmdpool)VUID-vkCmdCopyImageToBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdCopyImageToBuffer-renderpass)VUID-vkCmdCopyImageToBuffer-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyImageToBuffer-videocoding)VUID-vkCmdCopyImageToBuffer-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdCopyImageToBuffer-regionCount-arraylength)VUID-vkCmdCopyImageToBuffer-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-vkCmdCopyImageToBuffer-commonparent)VUID-vkCmdCopyImageToBuffer-commonparent  
  Each of `commandBuffer`, `dstBuffer`, and `srcImage` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

vkCmdCopyImageToBuffer is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyImageToBuffer).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
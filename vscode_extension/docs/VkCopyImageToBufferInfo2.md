# VkCopyImageToBufferInfo2(3) Manual Page

## Name

VkCopyImageToBufferInfo2 - Structure specifying parameters of an image to buffer copy command



## [](#_c_specification)C Specification

The `VkCopyImageToBufferInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkCopyImageToBufferInfo2 {
    VkStructureType              sType;
    const void*                  pNext;
    VkImage                      srcImage;
    VkImageLayout                srcImageLayout;
    VkBuffer                     dstBuffer;
    uint32_t                     regionCount;
    const VkBufferImageCopy2*    pRegions;
} VkCopyImageToBufferInfo2;
```

or the equivalent

```c++
// Provided by VK_KHR_copy_commands2
typedef VkCopyImageToBufferInfo2 VkCopyImageToBufferInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcImage` is the source image.
- `srcImageLayout` is the layout of the source image subresources for the copy.
- `dstBuffer` is the destination buffer.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html) structures specifying the regions to copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyImageToBufferInfo2-pRegions-04566)VUID-VkCopyImageToBufferInfo2-pRegions-04566  
  The image region specified by each element of `pRegions` that does not contain [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain **must** be contained within the specified `imageSubresource` of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2KHR-pRegions-04557)VUID-VkCopyImageToBufferInfo2KHR-pRegions-04557  
  If the image region specified by each element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, the [rotated source region](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers-images-rotation-addressing) **must** be contained within `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2KHR-pRegions-04558)VUID-VkCopyImageToBufferInfo2KHR-pRegions-04558  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `srcImage` **must** have a 1x1x1 [texel block extent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility-classes)
- [](#VUID-VkCopyImageToBufferInfo2KHR-pRegions-06205)VUID-VkCopyImageToBufferInfo2KHR-pRegions-06205  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `srcImage` **must** be of type `VK_IMAGE_TYPE_2D`
- [](#VUID-VkCopyImageToBufferInfo2KHR-pRegions-06206)VUID-VkCopyImageToBufferInfo2KHR-pRegions-06206  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `srcImage` **must** not have a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar)

<!--THE END-->

- [](#VUID-VkCopyImageToBufferInfo2-pRegions-00183)VUID-VkCopyImageToBufferInfo2-pRegions-00183  
  `dstBuffer` **must** be large enough to contain all buffer locations that are accessed according to [Buffer and Image Addressing](#copies-buffers-images-addressing), for each element of `pRegions`
- [](#VUID-VkCopyImageToBufferInfo2-pRegions-00184)VUID-VkCopyImageToBufferInfo2-pRegions-00184  
  The union of all source regions, and the union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-00186)VUID-VkCopyImageToBufferInfo2-srcImage-00186  
  `srcImage` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage flag
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-01998)VUID-VkCopyImageToBufferInfo2-srcImage-01998  
  The [format features](#resources-image-format-features) of `srcImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`
- [](#VUID-VkCopyImageToBufferInfo2-dstBuffer-00191)VUID-VkCopyImageToBufferInfo2-dstBuffer-00191  
  `dstBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-VkCopyImageToBufferInfo2-dstBuffer-00192)VUID-VkCopyImageToBufferInfo2-dstBuffer-00192  
  If `dstBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkCopyImageToBufferInfo2-srcImageLayout-00189)VUID-VkCopyImageToBufferInfo2-srcImageLayout-00189  
  `srcImageLayout` **must** specify the layout of the image subresources of `srcImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-VkCopyImageToBufferInfo2-srcImageLayout-01397)VUID-VkCopyImageToBufferInfo2-srcImageLayout-01397  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

<!--THE END-->

- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07966)VUID-VkCopyImageToBufferInfo2-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkCopyImageToBufferInfo2-imageSubresource-07967)VUID-VkCopyImageToBufferInfo2-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-VkCopyImageToBufferInfo2-imageSubresource-07968)VUID-VkCopyImageToBufferInfo2-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07969)VUID-VkCopyImageToBufferInfo2-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!--THE END-->

- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07973)VUID-VkCopyImageToBufferInfo2-srcImage-07973  
  `srcImage` **must** have a sample count equal to `VK_SAMPLE_COUNT_1_BIT`

<!--THE END-->

- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07979)VUID-VkCopyImageToBufferInfo2-srcImage-07979  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height` **must** be `1`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-09104)VUID-VkCopyImageToBufferInfo2-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `imageSubresource` of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07980)VUID-VkCopyImageToBufferInfo2-srcImage-07980  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `imageOffset.z` **must** be `0` and `imageExtent.depth` **must** be `1`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07274)VUID-VkCopyImageToBufferInfo2-srcImage-07274  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-10051)VUID-VkCopyImageToBufferInfo2-imageOffset-10051  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `imageOffset.x` does not equal the width of the subresource specified by `imageSubresource`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07275)VUID-VkCopyImageToBufferInfo2-srcImage-07275  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-10052)VUID-VkCopyImageToBufferInfo2-imageOffset-10052  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `imageOffset.y` does not equal the height of the subresource specified by `imageSubresource`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07276)VUID-VkCopyImageToBufferInfo2-srcImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-00207)VUID-VkCopyImageToBufferInfo2-srcImage-00207  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of `imageOffset.x` and `extent.width` does not equal the width of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-10053)VUID-VkCopyImageToBufferInfo2-imageOffset-10053  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference of `imageOffset.x` and `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-10054)VUID-VkCopyImageToBufferInfo2-imageOffset-10054  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.x` and `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-10055)VUID-VkCopyImageToBufferInfo2-imageOffset-10055  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of `imageOffset.x` and `extent.height` does not equal the width of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-00208)VUID-VkCopyImageToBufferInfo2-srcImage-00208  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of `imageOffset.y` and `extent.height` does not equal the height of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-10056)VUID-VkCopyImageToBufferInfo2-imageOffset-10056  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of `imageOffset.y` and `extent.width` does not equal the height of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-10057)VUID-VkCopyImageToBufferInfo2-imageOffset-10057  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.y` and `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-10058)VUID-VkCopyImageToBufferInfo2-imageOffset-10058  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference of `imageOffset.y` and `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-00209)VUID-VkCopyImageToBufferInfo2-srcImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and `extent.depth` does not equal the depth of the subresource specified by `srcSubresource`, `extent.depth` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageSubresource-09105)VUID-VkCopyImageToBufferInfo2-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must** specify aspects present in `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07981)VUID-VkCopyImageToBufferInfo2-srcImage-07981  
  If `srcImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `imageSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07983)VUID-VkCopyImageToBufferInfo2-srcImage-07983  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, for each element of `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and `imageSubresource.layerCount` **must** be `1`

<!--THE END-->

- [](#VUID-VkCopyImageToBufferInfo2-bufferRowLength-09106)VUID-VkCopyImageToBufferInfo2-bufferRowLength-09106  
  For each element of `pRegions`, `bufferRowLength` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-bufferImageHeight-09107)VUID-VkCopyImageToBufferInfo2-bufferImageHeight-09107  
  For each element of `pRegions`, `bufferImageHeight` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-bufferRowLength-09108)VUID-VkCopyImageToBufferInfo2-bufferRowLength-09108  
  For each element of `pRegions`, `bufferRowLength` divided by the [texel block extent width](#formats-compatibility-classes) and then multiplied by the texel block size of `srcImage` **must** be less than or equal to 231-1

<!--THE END-->

- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07975)VUID-VkCopyImageToBufferInfo2-srcImage-07975  
  If `srcImage` does not have either a depth/stencil format or a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `bufferOffset` **must** be a multiple of the [texel block size](#formats-compatibility-classes)
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07976)VUID-VkCopyImageToBufferInfo2-srcImage-07976  
  If `srcImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `bufferOffset` **must** be a multiple of the element size of the compatible format for the format and the `aspectMask` of the `imageSubresource` as defined in [\[formats-compatible-planes\]](#formats-compatible-planes)
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-07978)VUID-VkCopyImageToBufferInfo2-srcImage-07978  
  If `srcImage` has a depth/stencil format, the `bufferOffset` member of any element of `pRegions` **must** be a multiple of `4`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-00197)VUID-VkCopyImageToBufferInfo2-imageOffset-00197  
  For each element of `pRegions` not containing `VkCopyCommandTransformInfoQCOM` in its `pNext` chain, `imageOffset.x` and (`imageExtent.width` + `imageOffset.x`) **must** both be greater than or equal to `0` and less than or equal to the width of the specified `imageSubresource` of `srcImage`
- [](#VUID-VkCopyImageToBufferInfo2-imageOffset-00198)VUID-VkCopyImageToBufferInfo2-imageOffset-00198  
  For each element of `pRegions` not containing `VkCopyCommandTransformInfoQCOM` in its `pNext` chain, `imageOffset.y` and (`imageExtent.height` + `imageOffset.y`) **must** both be greater than or equal to `0` and less than or equal to the height of the specified `imageSubresource` of `srcImage`

Valid Usage (Implicit)

- [](#VUID-VkCopyImageToBufferInfo2-sType-sType)VUID-VkCopyImageToBufferInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_BUFFER_INFO_2`
- [](#VUID-VkCopyImageToBufferInfo2-pNext-pNext)VUID-VkCopyImageToBufferInfo2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyImageToBufferInfo2-srcImage-parameter)VUID-VkCopyImageToBufferInfo2-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkCopyImageToBufferInfo2-srcImageLayout-parameter)VUID-VkCopyImageToBufferInfo2-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkCopyImageToBufferInfo2-dstBuffer-parameter)VUID-VkCopyImageToBufferInfo2-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkCopyImageToBufferInfo2-pRegions-parameter)VUID-VkCopyImageToBufferInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html) structures
- [](#VUID-VkCopyImageToBufferInfo2-regionCount-arraylength)VUID-VkCopyImageToBufferInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-VkCopyImageToBufferInfo2-commonparent)VUID-VkCopyImageToBufferInfo2-commonparent  
  Both of `dstBuffer`, and `srcImage` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer2.html), [vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyImageToBufferInfo2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
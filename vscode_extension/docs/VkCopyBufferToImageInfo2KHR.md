# VkCopyBufferToImageInfo2(3) Manual Page

## Name

VkCopyBufferToImageInfo2 - Structure specifying parameters of a buffer to image copy command



## [](#_c_specification)C Specification

The `VkCopyBufferToImageInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkCopyBufferToImageInfo2 {
    VkStructureType              sType;
    const void*                  pNext;
    VkBuffer                     srcBuffer;
    VkImage                      dstImage;
    VkImageLayout                dstImageLayout;
    uint32_t                     regionCount;
    const VkBufferImageCopy2*    pRegions;
} VkCopyBufferToImageInfo2;
```

or the equivalent

```c++
// Provided by VK_KHR_copy_commands2
typedef VkCopyBufferToImageInfo2 VkCopyBufferToImageInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcBuffer` is the source buffer.
- `dstImage` is the destination image.
- `dstImageLayout` is the layout of the destination image subresources for the copy.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html) structures specifying the regions to copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyBufferToImageInfo2-pRegions-04565)VUID-VkCopyBufferToImageInfo2-pRegions-04565  
  The image region specified by each element of `pRegions` that does not contain [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain **must** be contained within the specified `imageSubresource` of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2KHR-pRegions-04554)VUID-VkCopyBufferToImageInfo2KHR-pRegions-04554  
  If the image region specified by each element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, the [rotated destination region](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers-images-rotation-addressing) **must** be contained within `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2KHR-pRegions-04555)VUID-VkCopyBufferToImageInfo2KHR-pRegions-04555  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `dstImage` **must** have a 1x1x1 [texel block extent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility-classes)
- [](#VUID-VkCopyBufferToImageInfo2KHR-pRegions-06203)VUID-VkCopyBufferToImageInfo2KHR-pRegions-06203  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `dstImage` **must** be of type `VK_IMAGE_TYPE_2D`
- [](#VUID-VkCopyBufferToImageInfo2KHR-pRegions-06204)VUID-VkCopyBufferToImageInfo2KHR-pRegions-06204  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `dstImage` **must** not have a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar)

<!--THE END-->

- [](#VUID-VkCopyBufferToImageInfo2-pRegions-00171)VUID-VkCopyBufferToImageInfo2-pRegions-00171  
  `srcBuffer` **must** be large enough to contain all buffer locations that are accessed according to [Buffer and Image Addressing](#copies-buffers-images-addressing), for each element of `pRegions`
- [](#VUID-VkCopyBufferToImageInfo2-pRegions-00173)VUID-VkCopyBufferToImageInfo2-pRegions-00173  
  The union of all source regions, and the union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory
- [](#VUID-VkCopyBufferToImageInfo2-srcBuffer-00174)VUID-VkCopyBufferToImageInfo2-srcBuffer-00174  
  `srcBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` usage flag
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-01997)VUID-VkCopyBufferToImageInfo2-dstImage-01997  
  The [format features](#resources-image-format-features) of `dstImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`
- [](#VUID-VkCopyBufferToImageInfo2-srcBuffer-00176)VUID-VkCopyBufferToImageInfo2-srcBuffer-00176  
  If `srcBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-00177)VUID-VkCopyBufferToImageInfo2-dstImage-00177  
  `dstImage` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-VkCopyBufferToImageInfo2-dstImageLayout-00180)VUID-VkCopyBufferToImageInfo2-dstImageLayout-00180  
  `dstImageLayout` **must** specify the layout of the image subresources of `dstImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-VkCopyBufferToImageInfo2-dstImageLayout-01396)VUID-VkCopyBufferToImageInfo2-dstImageLayout-01396  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-VkCopyBufferToImageInfo2-pRegions-07931)VUID-VkCopyBufferToImageInfo2-pRegions-07931  
  If [VK\_EXT\_depth\_range\_unrestricted](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_range_unrestricted.html) is not enabled, for each element of `pRegions` whose `imageSubresource` contains a depth aspect, the data in `srcBuffer` **must** be in the range \[0,1]

<!--THE END-->

- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07966)VUID-VkCopyBufferToImageInfo2-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkCopyBufferToImageInfo2-imageSubresource-07967)VUID-VkCopyBufferToImageInfo2-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkCopyBufferToImageInfo2-imageSubresource-07968)VUID-VkCopyBufferToImageInfo2-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07969)VUID-VkCopyBufferToImageInfo2-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!--THE END-->

- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07973)VUID-VkCopyBufferToImageInfo2-dstImage-07973  
  `dstImage` **must** have a sample count equal to `VK_SAMPLE_COUNT_1_BIT`

<!--THE END-->

- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07979)VUID-VkCopyBufferToImageInfo2-dstImage-07979  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height` **must** be `1`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-09104)VUID-VkCopyBufferToImageInfo2-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `imageSubresource` of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07980)VUID-VkCopyBufferToImageInfo2-dstImage-07980  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `imageOffset.z` **must** be `0` and `imageExtent.depth` **must** be `1`
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07274)VUID-VkCopyBufferToImageInfo2-dstImage-07274  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-10051)VUID-VkCopyBufferToImageInfo2-imageOffset-10051  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `imageOffset.x` does not equal the width of the subresource specified by `imageSubresource`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07275)VUID-VkCopyBufferToImageInfo2-dstImage-07275  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-10052)VUID-VkCopyBufferToImageInfo2-imageOffset-10052  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `imageOffset.y` does not equal the height of the subresource specified by `imageSubresource`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07276)VUID-VkCopyBufferToImageInfo2-dstImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-00207)VUID-VkCopyBufferToImageInfo2-dstImage-00207  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of `imageOffset.x` and `extent.width` does not equal the width of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-10053)VUID-VkCopyBufferToImageInfo2-imageOffset-10053  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference of `imageOffset.x` and `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-10054)VUID-VkCopyBufferToImageInfo2-imageOffset-10054  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.x` and `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-10055)VUID-VkCopyBufferToImageInfo2-imageOffset-10055  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of `imageOffset.x` and `extent.height` does not equal the width of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-00208)VUID-VkCopyBufferToImageInfo2-dstImage-00208  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of `imageOffset.y` and `extent.height` does not equal the height of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-10056)VUID-VkCopyBufferToImageInfo2-imageOffset-10056  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of `imageOffset.y` and `extent.width` does not equal the height of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-10057)VUID-VkCopyBufferToImageInfo2-imageOffset-10057  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.y` and `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageOffset-10058)VUID-VkCopyBufferToImageInfo2-imageOffset-10058  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference of `imageOffset.y` and `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-00209)VUID-VkCopyBufferToImageInfo2-dstImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and `extent.depth` does not equal the depth of the subresource specified by `srcSubresource`, `extent.depth` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-imageSubresource-09105)VUID-VkCopyBufferToImageInfo2-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must** specify aspects present in `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07981)VUID-VkCopyBufferToImageInfo2-dstImage-07981  
  If `dstImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `imageSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07983)VUID-VkCopyBufferToImageInfo2-dstImage-07983  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, for each element of `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and `imageSubresource.layerCount` **must** be `1`

<!--THE END-->

- [](#VUID-VkCopyBufferToImageInfo2-bufferRowLength-09106)VUID-VkCopyBufferToImageInfo2-bufferRowLength-09106  
  For each element of `pRegions`, `bufferRowLength` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-bufferImageHeight-09107)VUID-VkCopyBufferToImageInfo2-bufferImageHeight-09107  
  For each element of `pRegions`, `bufferImageHeight` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-bufferRowLength-09108)VUID-VkCopyBufferToImageInfo2-bufferRowLength-09108  
  For each element of `pRegions`, `bufferRowLength` divided by the [texel block extent width](#formats-compatibility-classes) and then multiplied by the texel block size of `dstImage` **must** be less than or equal to 231-1

<!--THE END-->

- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07975)VUID-VkCopyBufferToImageInfo2-dstImage-07975  
  If `dstImage` does not have either a depth/stencil format or a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `bufferOffset` **must** be a multiple of the [texel block size](#formats-compatibility-classes)
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07976)VUID-VkCopyBufferToImageInfo2-dstImage-07976  
  If `dstImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `bufferOffset` **must** be a multiple of the element size of the compatible format for the format and the `aspectMask` of the `imageSubresource` as defined in [\[formats-compatible-planes\]](#formats-compatible-planes)
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-07978)VUID-VkCopyBufferToImageInfo2-dstImage-07978  
  If `dstImage` has a depth/stencil format, the `bufferOffset` member of any element of `pRegions` **must** be a multiple of `4`
- [](#VUID-VkCopyBufferToImageInfo2-pRegions-06223)VUID-VkCopyBufferToImageInfo2-pRegions-06223  
  For each element of `pRegions` not containing `VkCopyCommandTransformInfoQCOM` in its `pNext` chain, `imageOffset.x` and (`imageExtent.width` + `imageOffset.x`) **must** both be greater than or equal to `0` and less than or equal to the width of the specified `imageSubresource` of `dstImage`
- [](#VUID-VkCopyBufferToImageInfo2-pRegions-06224)VUID-VkCopyBufferToImageInfo2-pRegions-06224  
  For each element of `pRegions` not containing `VkCopyCommandTransformInfoQCOM` in its `pNext` chain, `imageOffset.y` and (`imageExtent.height` + `imageOffset.y`) **must** both be greater than or equal to `0` and less than or equal to the height of the specified `imageSubresource` of `dstImage`

Valid Usage (Implicit)

- [](#VUID-VkCopyBufferToImageInfo2-sType-sType)VUID-VkCopyBufferToImageInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_BUFFER_TO_IMAGE_INFO_2`
- [](#VUID-VkCopyBufferToImageInfo2-pNext-pNext)VUID-VkCopyBufferToImageInfo2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyBufferToImageInfo2-srcBuffer-parameter)VUID-VkCopyBufferToImageInfo2-srcBuffer-parameter  
  `srcBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkCopyBufferToImageInfo2-dstImage-parameter)VUID-VkCopyBufferToImageInfo2-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkCopyBufferToImageInfo2-dstImageLayout-parameter)VUID-VkCopyBufferToImageInfo2-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkCopyBufferToImageInfo2-pRegions-parameter)VUID-VkCopyBufferToImageInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html) structures
- [](#VUID-VkCopyBufferToImageInfo2-regionCount-arraylength)VUID-VkCopyBufferToImageInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-VkCopyBufferToImageInfo2-commonparent)VUID-VkCopyBufferToImageInfo2-commonparent  
  Both of `dstImage`, and `srcBuffer` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage2.html), [vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyBufferToImageInfo2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
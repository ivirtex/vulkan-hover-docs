# VkCopyMemoryToImageInfo(3) Manual Page

## Name

VkCopyMemoryToImageInfo - Structure specifying parameters of host memory to image copy command



## [](#_c_specification)C Specification

The `VkCopyMemoryToImageInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkCopyMemoryToImageInfo {
    VkStructureType               sType;
    const void*                   pNext;
    VkHostImageCopyFlags          flags;
    VkImage                       dstImage;
    VkImageLayout                 dstImageLayout;
    uint32_t                      regionCount;
    const VkMemoryToImageCopy*    pRegions;
} VkCopyMemoryToImageInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkCopyMemoryToImageInfo VkCopyMemoryToImageInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkHostImageCopyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlagBits.html) values describing additional copy parameters.
- `dstImage` is the destination image.
- `dstImageLayout` is the layout of the destination image subresources for the copy.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkMemoryToImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryToImageCopy.html) structures specifying the regions to copy.

## [](#_description)Description

`vkCopyMemoryToImage` does not check whether the device memory associated with `dstImage` is currently in use before performing the copy. The application **must** guarantee that any previously submitted command that reads from or writes to the copy regions has completed before the host performs the copy.

Copy regions for the image **must** be aligned to a multiple of the texel block extent in each dimension, except at the edges of the image, where region extents **must** match the edge of the image.

Valid Usage

- [](#VUID-VkCopyMemoryToImageInfo-dstImage-09109)VUID-VkCopyMemoryToImageInfo-dstImage-09109  
  If `dstImage` is sparse then all memory ranges accessed by the copy command **must** be bound as described in [Binding Resource Memory](#sparsememory-resource-binding)
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-09111)VUID-VkCopyMemoryToImageInfo-dstImage-09111  
  If the stencil aspect of `dstImage` is accessed, and `dstImage` was not created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `dstImage` **must** have been created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` set in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-09112)VUID-VkCopyMemoryToImageInfo-dstImage-09112  
  If the stencil aspect of `dstImage` is accessed, and `dstImage` was created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `dstImage` **must** have been created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` set in [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-09113)VUID-VkCopyMemoryToImageInfo-dstImage-09113  
  If non-stencil aspects of `dstImage` are accessed, `dstImage` **must** have been created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` set in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-09114)VUID-VkCopyMemoryToImageInfo-imageOffset-09114  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_BIT`, the `x`, `y`, and `z` members of the `imageOffset` member of each element of `pRegions` **must** be `0`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-09115)VUID-VkCopyMemoryToImageInfo-dstImage-09115  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_BIT`, the `imageExtent` member of each element of `pRegions` **must** equal the extents of `dstImage` identified by `imageSubresource`

<!--THE END-->

- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07966)VUID-VkCopyMemoryToImageInfo-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkCopyMemoryToImageInfo-imageSubresource-07967)VUID-VkCopyMemoryToImageInfo-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkCopyMemoryToImageInfo-imageSubresource-07968)VUID-VkCopyMemoryToImageInfo-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07969)VUID-VkCopyMemoryToImageInfo-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!--THE END-->

- [](#VUID-VkCopyMemoryToImageInfo-imageSubresource-07971)VUID-VkCopyMemoryToImageInfo-imageSubresource-07971  
  For each element of `pRegions`, `imageOffset.x` and (`imageExtent.width` + `imageOffset.x`) **must** both be greater than or equal to `0` and less than or equal to the width of the specified `imageSubresource` of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageSubresource-07972)VUID-VkCopyMemoryToImageInfo-imageSubresource-07972  
  For each element of `pRegions`, `imageOffset.y` and (`imageExtent.height` + `imageOffset.y`) **must** both be greater than or equal to `0` and less than or equal to the height of the specified `imageSubresource` of `dstImage`

<!--THE END-->

- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07973)VUID-VkCopyMemoryToImageInfo-dstImage-07973  
  `dstImage` **must** have a sample count equal to `VK_SAMPLE_COUNT_1_BIT`

<!--THE END-->

- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07979)VUID-VkCopyMemoryToImageInfo-dstImage-07979  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height` **must** be `1`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-09104)VUID-VkCopyMemoryToImageInfo-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `imageSubresource` of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07980)VUID-VkCopyMemoryToImageInfo-dstImage-07980  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `imageOffset.z` **must** be `0` and `imageExtent.depth` **must** be `1`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07274)VUID-VkCopyMemoryToImageInfo-dstImage-07274  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-10051)VUID-VkCopyMemoryToImageInfo-imageOffset-10051  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `imageOffset.x` does not equal the width of the subresource specified by `imageSubresource`, `imageOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07275)VUID-VkCopyMemoryToImageInfo-dstImage-07275  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-10052)VUID-VkCopyMemoryToImageInfo-imageOffset-10052  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `imageOffset.y` does not equal the height of the subresource specified by `imageSubresource`, `imageOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07276)VUID-VkCopyMemoryToImageInfo-dstImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-00207)VUID-VkCopyMemoryToImageInfo-dstImage-00207  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of `imageOffset.x` and `extent.width` does not equal the width of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-10053)VUID-VkCopyMemoryToImageInfo-imageOffset-10053  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference of `imageOffset.x` and `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-10054)VUID-VkCopyMemoryToImageInfo-imageOffset-10054  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.x` and `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-10055)VUID-VkCopyMemoryToImageInfo-imageOffset-10055  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of `imageOffset.x` and `extent.height` does not equal the width of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-00208)VUID-VkCopyMemoryToImageInfo-dstImage-00208  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of `imageOffset.y` and `extent.height` does not equal the height of the subresource specified by `imageSubresource`, `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-10056)VUID-VkCopyMemoryToImageInfo-imageOffset-10056  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of `imageOffset.y` and `extent.width` does not equal the height of the subresource specified by `imageSubresource`, `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-10057)VUID-VkCopyMemoryToImageInfo-imageOffset-10057  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference of `imageOffset.y` and `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageOffset-10058)VUID-VkCopyMemoryToImageInfo-imageOffset-10058  
  For each element of `pRegions`, if [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform` is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference of `imageOffset.y` and `extent.width` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-00209)VUID-VkCopyMemoryToImageInfo-dstImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and `extent.depth` does not equal the depth of the subresource specified by `srcSubresource`, `extent.depth` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-imageSubresource-09105)VUID-VkCopyMemoryToImageInfo-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must** specify aspects present in `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07981)VUID-VkCopyMemoryToImageInfo-dstImage-07981  
  If `dstImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `imageSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-07983)VUID-VkCopyMemoryToImageInfo-dstImage-07983  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, for each element of `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and `imageSubresource.layerCount` **must** be `1`

<!--THE END-->

- [](#VUID-VkCopyMemoryToImageInfo-memoryRowLength-09106)VUID-VkCopyMemoryToImageInfo-memoryRowLength-09106  
  For each element of `pRegions`, `memoryRowLength` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-memoryImageHeight-09107)VUID-VkCopyMemoryToImageInfo-memoryImageHeight-09107  
  For each element of `pRegions`, `memoryImageHeight` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyMemoryToImageInfo-memoryRowLength-09108)VUID-VkCopyMemoryToImageInfo-memoryRowLength-09108  
  For each element of `pRegions`, `memoryRowLength` divided by the [texel block extent width](#formats-compatibility-classes) and then multiplied by the texel block size of `dstImage` **must** be less than or equal to 231-1
- [](#VUID-VkCopyMemoryToImageInfo-dstImageLayout-09059)VUID-VkCopyMemoryToImageInfo-dstImageLayout-09059  
  `dstImageLayout` **must** specify the current layout of the image subresources of `dstImage` specified in `pRegions`
- [](#VUID-VkCopyMemoryToImageInfo-dstImageLayout-09060)VUID-VkCopyMemoryToImageInfo-dstImageLayout-09060  
  `dstImageLayout` **must** be one of the image layouts returned in [VkPhysicalDeviceHostImageCopyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostImageCopyProperties.html)::`pCopyDstLayouts`
- [](#VUID-VkCopyMemoryToImageInfo-flags-09393)VUID-VkCopyMemoryToImageInfo-flags-09393  
  If `flags` includes `VK_HOST_IMAGE_COPY_MEMCPY_BIT`, for each region in `pRegions`, `memoryRowLength` and `memoryImageHeight` **must** both be 0

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryToImageInfo-sType-sType)VUID-VkCopyMemoryToImageInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_IMAGE_INFO`
- [](#VUID-VkCopyMemoryToImageInfo-pNext-pNext)VUID-VkCopyMemoryToImageInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyMemoryToImageInfo-flags-parameter)VUID-VkCopyMemoryToImageInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkHostImageCopyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlagBits.html) values
- [](#VUID-VkCopyMemoryToImageInfo-dstImage-parameter)VUID-VkCopyMemoryToImageInfo-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkCopyMemoryToImageInfo-dstImageLayout-parameter)VUID-VkCopyMemoryToImageInfo-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkCopyMemoryToImageInfo-pRegions-parameter)VUID-VkCopyMemoryToImageInfo-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkMemoryToImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryToImageCopy.html) structures
- [](#VUID-VkCopyMemoryToImageInfo-regionCount-arraylength)VUID-VkCopyMemoryToImageInfo-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkHostImageCopyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlags.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkMemoryToImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryToImageCopy.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCopyMemoryToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToImage.html), [vkCopyMemoryToImageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToImageEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryToImageInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
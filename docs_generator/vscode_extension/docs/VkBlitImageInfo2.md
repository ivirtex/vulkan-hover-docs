# VkBlitImageInfo2(3) Manual Page

## Name

VkBlitImageInfo2 - Structure specifying parameters of blit image command



## [](#_c_specification)C Specification

The `VkBlitImageInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkBlitImageInfo2 {
    VkStructureType        sType;
    const void*            pNext;
    VkImage                srcImage;
    VkImageLayout          srcImageLayout;
    VkImage                dstImage;
    VkImageLayout          dstImageLayout;
    uint32_t               regionCount;
    const VkImageBlit2*    pRegions;
    VkFilter               filter;
} VkBlitImageInfo2;
```

or the equivalent

```c++
// Provided by VK_KHR_copy_commands2
typedef VkBlitImageInfo2 VkBlitImageInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcImage` is the source image.
- `srcImageLayout` is the layout of the source image subresources for the blit.
- `dstImage` is the destination image.
- `dstImageLayout` is the layout of the destination image subresources for the blit.
- `regionCount` is the number of regions to blit.
- `pRegions` is a pointer to an array of [VkImageBlit2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit2.html) structures specifying the regions to blit.
- `filter` is a [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html) specifying the filter to apply if the blits require scaling.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBlitImageInfo2-pRegions-00217)VUID-VkBlitImageInfo2-pRegions-00217  
  The union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory with any texel that **may** be sampled during the blit operation
- [](#VUID-VkBlitImageInfo2-srcImage-01999)VUID-VkBlitImageInfo2-srcImage-01999  
  The [format features](#resources-image-format-features) of `srcImage` **must** contain `VK_FORMAT_FEATURE_BLIT_SRC_BIT`
- [](#VUID-VkBlitImageInfo2-srcImage-06421)VUID-VkBlitImageInfo2-srcImage-06421  
  `srcImage` **must** not use a [format that requires a sampler Y′CBCR conversion](#formats-requiring-sampler-ycbcr-conversion)
- [](#VUID-VkBlitImageInfo2-srcImage-00219)VUID-VkBlitImageInfo2-srcImage-00219  
  `srcImage` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage flag
- [](#VUID-VkBlitImageInfo2-srcImage-00220)VUID-VkBlitImageInfo2-srcImage-00220  
  If `srcImage` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkBlitImageInfo2-srcImageLayout-00221)VUID-VkBlitImageInfo2-srcImageLayout-00221  
  `srcImageLayout` **must** specify the layout of the image subresources of `srcImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-VkBlitImageInfo2-srcImageLayout-01398)VUID-VkBlitImageInfo2-srcImageLayout-01398  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-VkBlitImageInfo2-srcImage-09459)VUID-VkBlitImageInfo2-srcImage-09459  
  If `srcImage` and `dstImage` are the same, and an elements of `pRegions` contains the `srcSubresource` and `dstSubresource` with matching `mipLevel` and overlapping array layers, then the `srcImageLayout` and `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`
- [](#VUID-VkBlitImageInfo2-dstImage-02000)VUID-VkBlitImageInfo2-dstImage-02000  
  The [format features](#resources-image-format-features) of `dstImage` **must** contain `VK_FORMAT_FEATURE_BLIT_DST_BIT`
- [](#VUID-VkBlitImageInfo2-dstImage-06422)VUID-VkBlitImageInfo2-dstImage-06422  
  `dstImage` **must** not use a [format that requires a sampler Y′CBCR conversion](#formats-requiring-sampler-ycbcr-conversion)
- [](#VUID-VkBlitImageInfo2-dstImage-00224)VUID-VkBlitImageInfo2-dstImage-00224  
  `dstImage` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-VkBlitImageInfo2-dstImage-00225)VUID-VkBlitImageInfo2-dstImage-00225  
  If `dstImage` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkBlitImageInfo2-dstImageLayout-00226)VUID-VkBlitImageInfo2-dstImageLayout-00226  
  `dstImageLayout` **must** specify the layout of the image subresources of `dstImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-VkBlitImageInfo2-dstImageLayout-01399)VUID-VkBlitImageInfo2-dstImageLayout-01399  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-VkBlitImageInfo2-srcImage-00229)VUID-VkBlitImageInfo2-srcImage-00229  
  If either of `srcImage` or `dstImage` was created with a signed integer [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), the other **must** also have been created with a signed integer [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html)
- [](#VUID-VkBlitImageInfo2-srcImage-00230)VUID-VkBlitImageInfo2-srcImage-00230  
  If either of `srcImage` or `dstImage` was created with an unsigned integer [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), the other **must** also have been created with an unsigned integer [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html)
- [](#VUID-VkBlitImageInfo2-srcImage-00231)VUID-VkBlitImageInfo2-srcImage-00231  
  If either of `srcImage` or `dstImage` was created with a depth/stencil format, the other **must** have exactly the same format
- [](#VUID-VkBlitImageInfo2-srcImage-00232)VUID-VkBlitImageInfo2-srcImage-00232  
  If `srcImage` was created with a depth/stencil format, `filter` **must** be `VK_FILTER_NEAREST`
- [](#VUID-VkBlitImageInfo2-srcImage-00233)VUID-VkBlitImageInfo2-srcImage-00233  
  `srcImage` **must** have been created with a `samples` value of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkBlitImageInfo2-dstImage-00234)VUID-VkBlitImageInfo2-dstImage-00234  
  `dstImage` **must** have been created with a `samples` value of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkBlitImageInfo2-filter-02001)VUID-VkBlitImageInfo2-filter-02001  
  If `filter` is `VK_FILTER_LINEAR`, then the [format features](#resources-image-format-features) of `srcImage` **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`
- [](#VUID-VkBlitImageInfo2-filter-02002)VUID-VkBlitImageInfo2-filter-02002  
  If `filter` is `VK_FILTER_CUBIC_EXT`, then the [format features](#resources-image-format-features) of `srcImage` **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`
- [](#VUID-VkBlitImageInfo2-filter-00237)VUID-VkBlitImageInfo2-filter-00237  
  If `filter` is `VK_FILTER_CUBIC_EXT`, `srcImage` **must** be of type `VK_IMAGE_TYPE_2D`
- [](#VUID-VkBlitImageInfo2-srcSubresource-01705)VUID-VkBlitImageInfo2-srcSubresource-01705  
  The `srcSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-VkBlitImageInfo2-dstSubresource-01706)VUID-VkBlitImageInfo2-dstSubresource-01706  
  The `dstSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkBlitImageInfo2-srcSubresource-01707)VUID-VkBlitImageInfo2-srcSubresource-01707  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-VkBlitImageInfo2-dstSubresource-01708)VUID-VkBlitImageInfo2-dstSubresource-01708  
  If `dstSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkBlitImageInfo2-dstImage-02545)VUID-VkBlitImageInfo2-dstImage-02545  
  `dstImage` and `srcImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`
- [](#VUID-VkBlitImageInfo2-maintenance8-10207)VUID-VkBlitImageInfo2-maintenance8-10207  
  If the [`maintenance8`](#features-maintenance8) feature is enabled and `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `srcSubresource.baseArrayLayer` **must** be `0`, and `srcSubresource.layerCount` and `dstSubresource.layerCount` **must** each be `1`
- [](#VUID-VkBlitImageInfo2-maintenance8-10208)VUID-VkBlitImageInfo2-maintenance8-10208  
  If the [`maintenance8`](#features-maintenance8) feature is enabled and `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `dstSubresource.baseArrayLayer` **must** be `0`, and `srcSubresource.layerCount` and `dstSubresource.layerCount` **must** each be `1`
- [](#VUID-VkBlitImageInfo2-maintenance8-10579)VUID-VkBlitImageInfo2-maintenance8-10579  
  If the [`maintenance8`](#features-maintenance8) feature is enabled, `dstImage` is `VK_IMAGE_TYPE_3D`, and `srcImage` is not of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, the absolute difference of the `z` member of each member of `dstOffsets` **must** equal `srcSubresource.layerCount`
- [](#VUID-VkBlitImageInfo2-maintenance8-10580)VUID-VkBlitImageInfo2-maintenance8-10580  
  If the [`maintenance8`](#features-maintenance8) feature is enabled, `srcImage` is `VK_IMAGE_TYPE_3D`, and `dstImage` is not of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, the absolute difference of the `z` member of each member of `srcOffsets` **must** equal `dstSubresource.layerCount`
- [](#VUID-VkBlitImageInfo2-srcImage-00240)VUID-VkBlitImageInfo2-srcImage-00240  
  If the [`maintenance8`](#features-maintenance8) feature is not enabled and either `srcImage` or `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `srcSubresource.baseArrayLayer` and `dstSubresource.baseArrayLayer` **must** each be `0`, and `srcSubresource.layerCount` and `dstSubresource.layerCount` **must** each be `1`
- [](#VUID-VkBlitImageInfo2-aspectMask-00241)VUID-VkBlitImageInfo2-aspectMask-00241  
  For each element of `pRegions`, `srcSubresource.aspectMask` **must** specify aspects present in `srcImage`
- [](#VUID-VkBlitImageInfo2-aspectMask-00242)VUID-VkBlitImageInfo2-aspectMask-00242  
  For each element of `pRegions`, `dstSubresource.aspectMask` **must** specify aspects present in `dstImage`
- [](#VUID-VkBlitImageInfo2-srcOffset-00243)VUID-VkBlitImageInfo2-srcOffset-00243  
  For each element of `pRegions`, `srcOffsets`\[0].x and `srcOffsets`\[1].x **must** both be greater than or equal to `0` and less than or equal to the width of the specified `srcSubresource` of `srcImage`
- [](#VUID-VkBlitImageInfo2-srcOffset-00244)VUID-VkBlitImageInfo2-srcOffset-00244  
  For each element of `pRegions`, `srcOffsets`\[0].y and `srcOffsets`\[1].y **must** both be greater than or equal to `0` and less than or equal to the height of the specified `srcSubresource` of `srcImage`
- [](#VUID-VkBlitImageInfo2-srcImage-00245)VUID-VkBlitImageInfo2-srcImage-00245  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `srcOffsets`\[0].y **must** be `0` and `srcOffsets`\[1].y **must** be `1`
- [](#VUID-VkBlitImageInfo2-srcOffset-00246)VUID-VkBlitImageInfo2-srcOffset-00246  
  For each element of `pRegions`, `srcOffsets`\[0].z and `srcOffsets`\[1].z **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `srcSubresource` of `srcImage`
- [](#VUID-VkBlitImageInfo2-srcImage-00247)VUID-VkBlitImageInfo2-srcImage-00247  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `srcOffsets`\[0].z **must** be `0` and `srcOffsets`\[1].z **must** be `1`
- [](#VUID-VkBlitImageInfo2-dstOffset-00248)VUID-VkBlitImageInfo2-dstOffset-00248  
  For each element of `pRegions`, `dstOffsets`\[0].x and `dstOffsets`\[1].x **must** both be greater than or equal to `0` and less than or equal to the width of the specified `dstSubresource` of `dstImage`
- [](#VUID-VkBlitImageInfo2-dstOffset-00249)VUID-VkBlitImageInfo2-dstOffset-00249  
  For each element of `pRegions`, `dstOffsets`\[0].y and `dstOffsets`\[1].y **must** both be greater than or equal to `0` and less than or equal to the height of the specified `dstSubresource` of `dstImage`
- [](#VUID-VkBlitImageInfo2-dstImage-00250)VUID-VkBlitImageInfo2-dstImage-00250  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `dstOffsets`\[0].y **must** be `0` and `dstOffsets`\[1].y **must** be `1`
- [](#VUID-VkBlitImageInfo2-dstOffset-00251)VUID-VkBlitImageInfo2-dstOffset-00251  
  For each element of `pRegions`, `dstOffsets`\[0].z and `dstOffsets`\[1].z **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `dstSubresource` of `dstImage`
- [](#VUID-VkBlitImageInfo2-dstImage-00252)VUID-VkBlitImageInfo2-dstImage-00252  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `dstOffsets`\[0].z **must** be `0` and `dstOffsets`\[1].z **must** be `1`
- [](#VUID-VkBlitImageInfo2-pRegions-04561)VUID-VkBlitImageInfo2-pRegions-04561  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `srcImage` and `dstImage` **must** not be block-compressed images
- [](#VUID-VkBlitImageInfo2KHR-pRegions-06207)VUID-VkBlitImageInfo2KHR-pRegions-06207  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `srcImage` **must** be of type `VK_IMAGE_TYPE_2D`
- [](#VUID-VkBlitImageInfo2KHR-pRegions-06208)VUID-VkBlitImageInfo2KHR-pRegions-06208  
  If any element of `pRegions` contains [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html) in its `pNext` chain, then `srcImage` **must** not have a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar)
- [](#VUID-VkBlitImageInfo2-filter-09204)VUID-VkBlitImageInfo2-filter-09204  
  If `filter` is `VK_FILTER_CUBIC_EXT` and if the [selectableCubicWeights](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-selectableCubicWeights) feature is not enabled then the cubic weights **must** be `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`

Valid Usage (Implicit)

- [](#VUID-VkBlitImageInfo2-sType-sType)VUID-VkBlitImageInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BLIT_IMAGE_INFO_2`
- [](#VUID-VkBlitImageInfo2-pNext-pNext)VUID-VkBlitImageInfo2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkBlitImageCubicWeightsInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageCubicWeightsInfoQCOM.html)
- [](#VUID-VkBlitImageInfo2-sType-unique)VUID-VkBlitImageInfo2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkBlitImageInfo2-srcImage-parameter)VUID-VkBlitImageInfo2-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkBlitImageInfo2-srcImageLayout-parameter)VUID-VkBlitImageInfo2-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkBlitImageInfo2-dstImage-parameter)VUID-VkBlitImageInfo2-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkBlitImageInfo2-dstImageLayout-parameter)VUID-VkBlitImageInfo2-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkBlitImageInfo2-pRegions-parameter)VUID-VkBlitImageInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkImageBlit2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit2.html) structures
- [](#VUID-VkBlitImageInfo2-filter-parameter)VUID-VkBlitImageInfo2-filter-parameter  
  `filter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html) value
- [](#VUID-VkBlitImageInfo2-regionCount-arraylength)VUID-VkBlitImageInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-VkBlitImageInfo2-commonparent)VUID-VkBlitImageInfo2-commonparent  
  Both of `dstImage`, and `srcImage` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageBlit2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit2.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBlitImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2.html), [vkCmdBlitImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBlitImageInfo2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
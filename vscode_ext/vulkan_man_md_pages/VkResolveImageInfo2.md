# VkResolveImageInfo2(3) Manual Page

## Name

VkResolveImageInfo2 - Structure specifying parameters of resolve image
command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkResolveImageInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkResolveImageInfo2 {
    VkStructureType           sType;
    const void*               pNext;
    VkImage                   srcImage;
    VkImageLayout             srcImageLayout;
    VkImage                   dstImage;
    VkImageLayout             dstImageLayout;
    uint32_t                  regionCount;
    const VkImageResolve2*    pRegions;
} VkResolveImageInfo2;
```

or the equivalent

``` c
// Provided by VK_KHR_copy_commands2
typedef VkResolveImageInfo2 VkResolveImageInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcImage` is the source image.

- `srcImageLayout` is the layout of the source image subresources for
  the resolve.

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the resolve.

- `regionCount` is the number of regions to resolve.

- `pRegions` is a pointer to an array of
  [VkImageResolve2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve2.html) structures specifying the
  regions to resolve.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkResolveImageInfo2-pRegions-00255"
  id="VUID-VkResolveImageInfo2-pRegions-00255"></a>
  VUID-VkResolveImageInfo2-pRegions-00255  
  The union of all source regions, and the union of all destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-VkResolveImageInfo2-srcImage-00256"
  id="VUID-VkResolveImageInfo2-srcImage-00256"></a>
  VUID-VkResolveImageInfo2-srcImage-00256  
  If `srcImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkResolveImageInfo2-srcImage-00257"
  id="VUID-VkResolveImageInfo2-srcImage-00257"></a>
  VUID-VkResolveImageInfo2-srcImage-00257  
  `srcImage` **must** have a sample count equal to any valid sample
  count value other than `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkResolveImageInfo2-dstImage-00258"
  id="VUID-VkResolveImageInfo2-dstImage-00258"></a>
  VUID-VkResolveImageInfo2-dstImage-00258  
  If `dstImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkResolveImageInfo2-dstImage-00259"
  id="VUID-VkResolveImageInfo2-dstImage-00259"></a>
  VUID-VkResolveImageInfo2-dstImage-00259  
  `dstImage` **must** have a sample count equal to
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkResolveImageInfo2-srcImageLayout-00260"
  id="VUID-VkResolveImageInfo2-srcImageLayout-00260"></a>
  VUID-VkResolveImageInfo2-srcImageLayout-00260  
  `srcImageLayout` **must** specify the layout of the image subresources
  of `srcImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-VkResolveImageInfo2-srcImageLayout-01400"
  id="VUID-VkResolveImageInfo2-srcImageLayout-01400"></a>
  VUID-VkResolveImageInfo2-srcImageLayout-01400  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-VkResolveImageInfo2-dstImageLayout-00262"
  id="VUID-VkResolveImageInfo2-dstImageLayout-00262"></a>
  VUID-VkResolveImageInfo2-dstImageLayout-00262  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-VkResolveImageInfo2-dstImageLayout-01401"
  id="VUID-VkResolveImageInfo2-dstImageLayout-01401"></a>
  VUID-VkResolveImageInfo2-dstImageLayout-01401  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-VkResolveImageInfo2-dstImage-02003"
  id="VUID-VkResolveImageInfo2-dstImage-02003"></a>
  VUID-VkResolveImageInfo2-dstImage-02003  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkResolveImageInfo2-linearColorAttachment-06519"
  id="VUID-VkResolveImageInfo2-linearColorAttachment-06519"></a>
  VUID-VkResolveImageInfo2-linearColorAttachment-06519  
  If the [`linearColorAttachment`](#features-linearColorAttachment)
  feature is enabled and the image is created with
  `VK_IMAGE_TILING_LINEAR`, the [format
  features](#resources-image-format-features) of `dstImage` **must**
  contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

- <a href="#VUID-VkResolveImageInfo2-srcImage-01386"
  id="VUID-VkResolveImageInfo2-srcImage-01386"></a>
  VUID-VkResolveImageInfo2-srcImage-01386  
  `srcImage` and `dstImage` **must** have been created with the same
  image format

- <a href="#VUID-VkResolveImageInfo2-srcSubresource-01709"
  id="VUID-VkResolveImageInfo2-srcSubresource-01709"></a>
  VUID-VkResolveImageInfo2-srcSubresource-01709  
  The `srcSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-VkResolveImageInfo2-dstSubresource-01710"
  id="VUID-VkResolveImageInfo2-dstSubresource-01710"></a>
  VUID-VkResolveImageInfo2-dstSubresource-01710  
  The `dstSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-VkResolveImageInfo2-srcSubresource-01711"
  id="VUID-VkResolveImageInfo2-srcSubresource-01711"></a>
  VUID-VkResolveImageInfo2-srcSubresource-01711  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-VkResolveImageInfo2-dstSubresource-01712"
  id="VUID-VkResolveImageInfo2-dstSubresource-01712"></a>
  VUID-VkResolveImageInfo2-dstSubresource-01712  
  If `dstSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-VkResolveImageInfo2-dstImage-02546"
  id="VUID-VkResolveImageInfo2-dstImage-02546"></a>
  VUID-VkResolveImageInfo2-dstImage-02546  
  `dstImage` and `srcImage` **must** not have been created with `flags`
  containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- <a href="#VUID-VkResolveImageInfo2-srcImage-04446"
  id="VUID-VkResolveImageInfo2-srcImage-04446"></a>
  VUID-VkResolveImageInfo2-srcImage-04446  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `srcSubresource.layerCount` **must** be `1`

- <a href="#VUID-VkResolveImageInfo2-srcImage-04447"
  id="VUID-VkResolveImageInfo2-srcImage-04447"></a>
  VUID-VkResolveImageInfo2-srcImage-04447  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `dstSubresource.baseArrayLayer` **must** be `0` and
  `dstSubresource.layerCount` **must** be `1`

- <a href="#VUID-VkResolveImageInfo2-srcOffset-00269"
  id="VUID-VkResolveImageInfo2-srcOffset-00269"></a>
  VUID-VkResolveImageInfo2-srcOffset-00269  
  For each element of `pRegions`, `srcOffset.x` and (`extent.width` +
  `srcOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkResolveImageInfo2-srcOffset-00270"
  id="VUID-VkResolveImageInfo2-srcOffset-00270"></a>
  VUID-VkResolveImageInfo2-srcOffset-00270  
  For each element of `pRegions`, `srcOffset.y` and (`extent.height` +
  `srcOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkResolveImageInfo2-srcImage-00271"
  id="VUID-VkResolveImageInfo2-srcImage-00271"></a>
  VUID-VkResolveImageInfo2-srcImage-00271  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-VkResolveImageInfo2-srcOffset-00272"
  id="VUID-VkResolveImageInfo2-srcOffset-00272"></a>
  VUID-VkResolveImageInfo2-srcOffset-00272  
  For each element of `pRegions`, `srcOffset.z` and (`extent.depth` +
  `srcOffset.z`) **must** both be greater than or equal to `0` and less
  than or equal to the depth of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkResolveImageInfo2-srcImage-00273"
  id="VUID-VkResolveImageInfo2-srcImage-00273"></a>
  VUID-VkResolveImageInfo2-srcImage-00273  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `srcOffset.z` **must** be `0` and
  `extent.depth` **must** be `1`

- <a href="#VUID-VkResolveImageInfo2-dstOffset-00274"
  id="VUID-VkResolveImageInfo2-dstOffset-00274"></a>
  VUID-VkResolveImageInfo2-dstOffset-00274  
  For each element of `pRegions`, `dstOffset.x` and (`extent.width` +
  `dstOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkResolveImageInfo2-dstOffset-00275"
  id="VUID-VkResolveImageInfo2-dstOffset-00275"></a>
  VUID-VkResolveImageInfo2-dstOffset-00275  
  For each element of `pRegions`, `dstOffset.y` and (`extent.height` +
  `dstOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkResolveImageInfo2-dstImage-00276"
  id="VUID-VkResolveImageInfo2-dstImage-00276"></a>
  VUID-VkResolveImageInfo2-dstImage-00276  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-VkResolveImageInfo2-dstOffset-00277"
  id="VUID-VkResolveImageInfo2-dstOffset-00277"></a>
  VUID-VkResolveImageInfo2-dstOffset-00277  
  For each element of `pRegions`, `dstOffset.z` and (`extent.depth` +
  `dstOffset.z`) **must** both be greater than or equal to `0` and less
  than or equal to the depth of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkResolveImageInfo2-dstImage-00278"
  id="VUID-VkResolveImageInfo2-dstImage-00278"></a>
  VUID-VkResolveImageInfo2-dstImage-00278  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `dstOffset.z` **must** be `0` and
  `extent.depth` **must** be `1`

- <a href="#VUID-VkResolveImageInfo2-srcImage-06762"
  id="VUID-VkResolveImageInfo2-srcImage-06762"></a>
  VUID-VkResolveImageInfo2-srcImage-06762  
  `srcImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-VkResolveImageInfo2-srcImage-06763"
  id="VUID-VkResolveImageInfo2-srcImage-06763"></a>
  VUID-VkResolveImageInfo2-srcImage-06763  
  The [format features](#resources-image-format-features) of `srcImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`

- <a href="#VUID-VkResolveImageInfo2-dstImage-06764"
  id="VUID-VkResolveImageInfo2-dstImage-06764"></a>
  VUID-VkResolveImageInfo2-dstImage-06764  
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-VkResolveImageInfo2-dstImage-06765"
  id="VUID-VkResolveImageInfo2-dstImage-06765"></a>
  VUID-VkResolveImageInfo2-dstImage-06765  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

Valid Usage (Implicit)

- <a href="#VUID-VkResolveImageInfo2-sType-sType"
  id="VUID-VkResolveImageInfo2-sType-sType"></a>
  VUID-VkResolveImageInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RESOLVE_IMAGE_INFO_2`

- <a href="#VUID-VkResolveImageInfo2-pNext-pNext"
  id="VUID-VkResolveImageInfo2-pNext-pNext"></a>
  VUID-VkResolveImageInfo2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkResolveImageInfo2-srcImage-parameter"
  id="VUID-VkResolveImageInfo2-srcImage-parameter"></a>
  VUID-VkResolveImageInfo2-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkResolveImageInfo2-srcImageLayout-parameter"
  id="VUID-VkResolveImageInfo2-srcImageLayout-parameter"></a>
  VUID-VkResolveImageInfo2-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkResolveImageInfo2-dstImage-parameter"
  id="VUID-VkResolveImageInfo2-dstImage-parameter"></a>
  VUID-VkResolveImageInfo2-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkResolveImageInfo2-dstImageLayout-parameter"
  id="VUID-VkResolveImageInfo2-dstImageLayout-parameter"></a>
  VUID-VkResolveImageInfo2-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkResolveImageInfo2-pRegions-parameter"
  id="VUID-VkResolveImageInfo2-pRegions-parameter"></a>
  VUID-VkResolveImageInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkImageResolve2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve2.html) structures

- <a href="#VUID-VkResolveImageInfo2-regionCount-arraylength"
  id="VUID-VkResolveImageInfo2-regionCount-arraylength"></a>
  VUID-VkResolveImageInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-VkResolveImageInfo2-commonparent"
  id="VUID-VkResolveImageInfo2-commonparent"></a>
  VUID-VkResolveImageInfo2-commonparent  
  Both of `dstImage`, and `srcImage` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageResolve2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageResolve2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdResolveImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage2.html),
[vkCmdResolveImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkResolveImageInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

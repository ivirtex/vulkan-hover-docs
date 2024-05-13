# VkCopyMemoryToImageInfoEXT(3) Manual Page

## Name

VkCopyMemoryToImageInfoEXT - Structure specifying parameters of host
memory to image copy command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyMemoryToImageInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_host_image_copy
typedef struct VkCopyMemoryToImageInfoEXT {
    VkStructureType                  sType;
    const void*                      pNext;
    VkHostImageCopyFlagsEXT          flags;
    VkImage                          dstImage;
    VkImageLayout                    dstImageLayout;
    uint32_t                         regionCount;
    const VkMemoryToImageCopyEXT*    pRegions;
} VkCopyMemoryToImageInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkHostImageCopyFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagBitsEXT.html) values
  describing additional copy parameters.

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the copy.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkMemoryToImageCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryToImageCopyEXT.html) structures
  specifying the regions to copy.

## <a href="#_description" class="anchor"></a>Description

`vkCopyMemoryToImageEXT` does not check whether the device memory
associated with `dstImage` is currently in use before performing the
copy. The application **must** guarantee that any previously submitted
command that reads from or writes to the copy regions has completed
before the host performs the copy.

Copy regions for the image **must** be aligned to a multiple of the
texel block extent in each dimension, except at the edges of the image,
where region extents **must** match the edge of the image.

Valid Usage

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-09109"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-09109"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-09109  
  If `dstImage` is sparse then all memory ranges accessed by the copy
  command **must** be bound as described in [Binding Resource
  Memory](#sparsememory-resource-binding)

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-09111"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-09111"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-09111  
  If the stencil aspect of `dstImage` is accessed, and `dstImage` was
  not created with [separate stencil
  usage](#VkImageStencilUsageCreateInfo), `dstImage` **must** have been
  created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-09112"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-09112"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-09112  
  If the stencil aspect of `dstImage` is accessed, and `dstImage` was
  created with [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-09113"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-09113"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-09113  
  If non-stencil aspects of `dstImage` are accessed, `dstImage` **must**
  have been created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-imageOffset-09114"
  id="VUID-VkCopyMemoryToImageInfoEXT-imageOffset-09114"></a>
  VUID-VkCopyMemoryToImageInfoEXT-imageOffset-09114  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, the `x`, `y`, and
  `z` members of the `imageOffset` member of each element of `pRegions`
  **must** be `0`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-09115"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-09115"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-09115  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, the `imageExtent`
  member of each element of `pRegions` **must** equal the extents of
  `dstImage` identified by `imageSubresource`

<!-- -->

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07966"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07966"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07967"
  id="VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07967"></a>
  VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07968"
  id="VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07968"></a>
  VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of
  each element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07969"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07969"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07970"
  id="VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07970"></a>
  VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07970  
  The image region specified by each element of `pRegions` **must** be
  contained within the specified `imageSubresource` of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07971"
  id="VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07971"></a>
  VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07971  
  For each element of `pRegions`, `imageOffset.x` and
  (`imageExtent.width` + `imageOffset.x`) **must** both be greater than
  or equal to `0` and less than or equal to the width of the specified
  `imageSubresource` of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07972"
  id="VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07972"></a>
  VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-07972  
  For each element of `pRegions`, `imageOffset.y` and
  (`imageExtent.height` + `imageOffset.y`) **must** both be greater than
  or equal to `0` and less than or equal to the height of the specified
  `imageSubresource` of `dstImage`

<!-- -->

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07973"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07973"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07973  
  `dstImage` **must** have a sample count equal to
  `VK_SAMPLE_COUNT_1_BIT`

<!-- -->

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07979"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07979"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07979  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height`
  **must** be `1`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-imageOffset-09104"
  id="VUID-VkCopyMemoryToImageInfoEXT-imageOffset-09104"></a>
  VUID-VkCopyMemoryToImageInfoEXT-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and
  (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than
  or equal to `0` and less than or equal to the depth of the specified
  `imageSubresource` of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07980"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07980"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07980  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `imageOffset.z` **must** be `0`
  and `imageExtent.depth` **must** be `1`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07274"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07274"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07274  
  For each element of `pRegions`, `imageOffset.x` **must** be a multiple
  of the [texel block extent width](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07275"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07275"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07275  
  For each element of `pRegions`, `imageOffset.y` **must** be a multiple
  of the [texel block extent height](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07276"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07276"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-00207"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-00207"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-00207  
  For each element of `pRegions`, if the sum of `imageOffset.x` and
  `extent.width` does not equal the width of the subresource specified
  by `srcSubresource`, `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-00208"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-00208"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-00208  
  For each element of `pRegions`, if the sum of `imageOffset.y` and
  `extent.height` does not equal the height of the subresource specified
  by `srcSubresource`, `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-00209"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-00209"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-09105"
  id="VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-09105"></a>
  VUID-VkCopyMemoryToImageInfoEXT-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must**
  specify aspects present in `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07981"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07981"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07981  
  If `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `imageSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-07983"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-07983"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-07983  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, for each element of
  `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and
  `imageSubresource.layerCount` **must** be `1`

<!-- -->

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-memoryRowLength-09106"
  id="VUID-VkCopyMemoryToImageInfoEXT-memoryRowLength-09106"></a>
  VUID-VkCopyMemoryToImageInfoEXT-memoryRowLength-09106  
  For each element of `pRegions`, `memoryRowLength` **must** be a
  multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-memoryImageHeight-09107"
  id="VUID-VkCopyMemoryToImageInfoEXT-memoryImageHeight-09107"></a>
  VUID-VkCopyMemoryToImageInfoEXT-memoryImageHeight-09107  
  For each element of `pRegions`, `memoryImageHeight` **must** be a
  multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-memoryRowLength-09108"
  id="VUID-VkCopyMemoryToImageInfoEXT-memoryRowLength-09108"></a>
  VUID-VkCopyMemoryToImageInfoEXT-memoryRowLength-09108  
  For each element of `pRegions`, `memoryRowLength` divided by the
  [texel block extent width](#formats-compatibility-classes) and then
  multiplied by the texel block size of `dstImage` **must** be less than
  or equal to 2<sup>31</sup>-1

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-09059"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-09059"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-09059  
  `dstImageLayout` **must** specify the current layout of the image
  subresources of `dstImage` specified in `pRegions`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-09060"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-09060"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-09060  
  `dstImageLayout` **must** be one of the image layouts returned in
  [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)::`pCopyDstLayouts`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-flags-09393"
  id="VUID-VkCopyMemoryToImageInfoEXT-flags-09393"></a>
  VUID-VkCopyMemoryToImageInfoEXT-flags-09393  
  If `flags` includes `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, for each region
  in `pRegions`, `memoryRowLength` and `memoryImageHeight` **must** both
  be 0

Valid Usage (Implicit)

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-sType-sType"
  id="VUID-VkCopyMemoryToImageInfoEXT-sType-sType"></a>
  VUID-VkCopyMemoryToImageInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_IMAGE_INFO_EXT`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-pNext-pNext"
  id="VUID-VkCopyMemoryToImageInfoEXT-pNext-pNext"></a>
  VUID-VkCopyMemoryToImageInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-flags-parameter"
  id="VUID-VkCopyMemoryToImageInfoEXT-flags-parameter"></a>
  VUID-VkCopyMemoryToImageInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkHostImageCopyFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagBitsEXT.html) values

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImage-parameter"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImage-parameter"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-parameter"
  id="VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-parameter"></a>
  VUID-VkCopyMemoryToImageInfoEXT-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-pRegions-parameter"
  id="VUID-VkCopyMemoryToImageInfoEXT-pRegions-parameter"></a>
  VUID-VkCopyMemoryToImageInfoEXT-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkMemoryToImageCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryToImageCopyEXT.html) structures

- <a href="#VUID-VkCopyMemoryToImageInfoEXT-regionCount-arraylength"
  id="VUID-VkCopyMemoryToImageInfoEXT-regionCount-arraylength"></a>
  VUID-VkCopyMemoryToImageInfoEXT-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkHostImageCopyFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagsEXT.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkMemoryToImageCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryToImageCopyEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCopyMemoryToImageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToImageEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyMemoryToImageInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

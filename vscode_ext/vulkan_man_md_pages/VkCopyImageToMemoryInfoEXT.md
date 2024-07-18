# VkCopyImageToMemoryInfoEXT(3) Manual Page

## Name

VkCopyImageToMemoryInfoEXT - Structure specifying parameters of an image
to host memory copy command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyImageToMemoryInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_host_image_copy
typedef struct VkCopyImageToMemoryInfoEXT {
    VkStructureType                  sType;
    const void*                      pNext;
    VkHostImageCopyFlagsEXT          flags;
    VkImage                          srcImage;
    VkImageLayout                    srcImageLayout;
    uint32_t                         regionCount;
    const VkImageToMemoryCopyEXT*    pRegions;
} VkCopyImageToMemoryInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkHostImageCopyFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagBitsEXT.html) values
  describing additional copy parameters.

- `srcImage` is the source image.

- `srcImageLayout` is the layout of the source image subresources for
  the copy.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkImageToMemoryCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageToMemoryCopyEXT.html) structures
  specifying the regions to copy.

## <a href="#_description" class="anchor"></a>Description

`vkCopyImageToMemoryEXT` does not check whether the device memory
associated with `srcImage` is currently in use before performing the
copy. The application **must** guarantee that any previously submitted
command that writes to the copy regions has completed before the host
performs the copy.

Copy regions for the image **must** be aligned to a multiple of the
texel block extent in each dimension, except at the edges of the image,
where region extents **must** match the edge of the image.

Valid Usage

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-09109"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-09109"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-09109  
  If `srcImage` is sparse then all memory ranges accessed by the copy
  command **must** be bound as described in [Binding Resource
  Memory](#sparsememory-resource-binding)

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-09111"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-09111"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-09111  
  If the stencil aspect of `srcImage` is accessed, and `srcImage` was
  not created with [separate stencil
  usage](#VkImageStencilUsageCreateInfo), `srcImage` **must** have been
  created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-09112"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-09112"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-09112  
  If the stencil aspect of `srcImage` is accessed, and `srcImage` was
  created with [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `srcImage` **must** have been created with
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-09113"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-09113"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-09113  
  If non-stencil aspects of `srcImage` are accessed, `srcImage` **must**
  have been created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-09114"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-09114"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-09114  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, the `x`, `y`, and
  `z` members of the `imageOffset` member of each element of `pRegions`
  **must** be `0`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-09115"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-09115"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-09115  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, the `imageExtent`
  member of each element of `pRegions` **must** equal the extents of
  `srcImage` identified by `imageSubresource`

<!-- -->

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07966"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07966"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07967"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07967"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07968"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07968"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of
  each element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07969"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07969"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07970"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07970"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07970  
  The image region specified by each element of `pRegions` **must** be
  contained within the specified `imageSubresource` of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07971"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07971"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07971  
  For each element of `pRegions`, `imageOffset.x` and
  (`imageExtent.width` + `imageOffset.x`) **must** both be greater than
  or equal to `0` and less than or equal to the width of the specified
  `imageSubresource` of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07972"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07972"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-07972  
  For each element of `pRegions`, `imageOffset.y` and
  (`imageExtent.height` + `imageOffset.y`) **must** both be greater than
  or equal to `0` and less than or equal to the height of the specified
  `imageSubresource` of `srcImage`

<!-- -->

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07973"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07973"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07973  
  `srcImage` **must** have a sample count equal to
  `VK_SAMPLE_COUNT_1_BIT`

<!-- -->

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07979"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07979"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07979  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height`
  **must** be `1`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-09104"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-09104"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and
  (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than
  or equal to `0` and less than or equal to the depth of the specified
  `imageSubresource` of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07980"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07980"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07980  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `imageOffset.z` **must** be `0`
  and `imageExtent.depth` **must** be `1`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07274"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07274"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07274  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `imageOffset.x` **must** be
  a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10051"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10051"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10051  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `imageOffset.x` does not
  equal the width of the subresource specified by `imageSubresource`,
  `imageOffset.x` **must** be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07275"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07275"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07275  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `imageOffset.y` **must** be
  a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10052"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10052"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10052  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `imageOffset.y` does
  not equal the height of the subresource specified by
  `imageSubresource`, `imageOffset.y` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07276"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07276"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-00207"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-00207"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-00207  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of
  `imageOffset.x` and `extent.width` does not equal the width of the
  subresource specified by `imageSubresource`, `extent.width` **must**
  be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10053"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10053"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10053  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference
  of `imageOffset.x` and `extent.height` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10054"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10054"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10054  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `imageOffset.x` and `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10055"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10055"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10055  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of
  `imageOffset.x` and `extent.height` does not equal the width of the
  subresource specified by `imageSubresource`, `extent.height` **must**
  be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-00208"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-00208"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-00208  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of
  `imageOffset.y` and `extent.height` does not equal the height of the
  subresource specified by `imageSubresource`, `extent.height` **must**
  be a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10056"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10056"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10056  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of
  `imageOffset.y` and `extent.width` does not equal the height of the
  subresource specified by `imageSubresource`, `extent.width` **must**
  be a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10057"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10057"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10057  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `imageOffset.y` and `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10058"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10058"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageOffset-10058  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference
  of `imageOffset.y` and `extent.width` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-00209"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-00209"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-09105"
  id="VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-09105"></a>
  VUID-VkCopyImageToMemoryInfoEXT-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must**
  specify aspects present in `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07981"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07981"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07981  
  If `srcImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `imageSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-07983"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-07983"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-07983  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, for each element of
  `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and
  `imageSubresource.layerCount` **must** be `1`

<!-- -->

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-memoryRowLength-09106"
  id="VUID-VkCopyImageToMemoryInfoEXT-memoryRowLength-09106"></a>
  VUID-VkCopyImageToMemoryInfoEXT-memoryRowLength-09106  
  For each element of `pRegions`, `memoryRowLength` **must** be a
  multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-memoryImageHeight-09107"
  id="VUID-VkCopyImageToMemoryInfoEXT-memoryImageHeight-09107"></a>
  VUID-VkCopyImageToMemoryInfoEXT-memoryImageHeight-09107  
  For each element of `pRegions`, `memoryImageHeight` **must** be a
  multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-memoryRowLength-09108"
  id="VUID-VkCopyImageToMemoryInfoEXT-memoryRowLength-09108"></a>
  VUID-VkCopyImageToMemoryInfoEXT-memoryRowLength-09108  
  For each element of `pRegions`, `memoryRowLength` divided by the
  [texel block extent width](#formats-compatibility-classes) and then
  multiplied by the texel block size of `srcImage` **must** be less than
  or equal to 2<sup>31</sup>-1

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-09064"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-09064"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-09064  
  `srcImageLayout` **must** specify the current layout of the image
  subresources of `srcImage` specified in `pRegions`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-09065"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-09065"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-09065  
  `srcImageLayout` **must** be one of the image layouts returned in
  [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)::`pCopySrcLayouts`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-flags-09394"
  id="VUID-VkCopyImageToMemoryInfoEXT-flags-09394"></a>
  VUID-VkCopyImageToMemoryInfoEXT-flags-09394  
  If `flags` includes `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, for each region
  in `pRegions`, `memoryRowLength` and `memoryImageHeight` **must** both
  be 0

Valid Usage (Implicit)

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-sType-sType"
  id="VUID-VkCopyImageToMemoryInfoEXT-sType-sType"></a>
  VUID-VkCopyImageToMemoryInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_MEMORY_INFO_EXT`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-pNext-pNext"
  id="VUID-VkCopyImageToMemoryInfoEXT-pNext-pNext"></a>
  VUID-VkCopyImageToMemoryInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-flags-parameter"
  id="VUID-VkCopyImageToMemoryInfoEXT-flags-parameter"></a>
  VUID-VkCopyImageToMemoryInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkHostImageCopyFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagBitsEXT.html) values

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImage-parameter"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImage-parameter"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-parameter"
  id="VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-parameter"></a>
  VUID-VkCopyImageToMemoryInfoEXT-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-pRegions-parameter"
  id="VUID-VkCopyImageToMemoryInfoEXT-pRegions-parameter"></a>
  VUID-VkCopyImageToMemoryInfoEXT-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkImageToMemoryCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageToMemoryCopyEXT.html) structures

- <a href="#VUID-VkCopyImageToMemoryInfoEXT-regionCount-arraylength"
  id="VUID-VkCopyImageToMemoryInfoEXT-regionCount-arraylength"></a>
  VUID-VkCopyImageToMemoryInfoEXT-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkHostImageCopyFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagsEXT.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageToMemoryCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageToMemoryCopyEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCopyImageToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyImageToMemoryEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyImageToMemoryInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

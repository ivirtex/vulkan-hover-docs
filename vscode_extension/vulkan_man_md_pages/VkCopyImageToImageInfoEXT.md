# VkCopyImageToImageInfoEXT(3) Manual Page

## Name

VkCopyImageToImageInfoEXT - Structure specifying parameters of an image
to image host copy command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyImageToImageInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_host_image_copy
typedef struct VkCopyImageToImageInfoEXT {
    VkStructureType            sType;
    const void*                pNext;
    VkHostImageCopyFlagsEXT    flags;
    VkImage                    srcImage;
    VkImageLayout              srcImageLayout;
    VkImage                    dstImage;
    VkImageLayout              dstImageLayout;
    uint32_t                   regionCount;
    const VkImageCopy2*        pRegions;
} VkCopyImageToImageInfoEXT;
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

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the copy.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy2.html) structures specifying the regions to
  copy.

## <a href="#_description" class="anchor"></a>Description

`vkCopyImageToImageEXT` does not check whether the device memory
associated with `srcImage` or `dstImage` is currently in use before
performing the copy. The application **must** guarantee that any
previously submitted command that writes to the copy regions has
completed before the host performs the copy.

Valid Usage

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-09069"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-09069"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-09069  
  `srcImage` and `dstImage` **must** have been created with identical
  image creation parameters

<!-- -->

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-09109"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-09109"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-09109  
  If `srcImage` is sparse then all memory ranges accessed by the copy
  command **must** be bound as described in [Binding Resource
  Memory](#sparsememory-resource-binding)

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-09111"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-09111"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-09111  
  If the stencil aspect of `srcImage` is accessed, and `srcImage` was
  not created with [separate stencil
  usage](#VkImageStencilUsageCreateInfo), `srcImage` **must** have been
  created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-09112"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-09112"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-09112  
  If the stencil aspect of `srcImage` is accessed, and `srcImage` was
  created with [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `srcImage` **must** have been created with
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-09113"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-09113"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-09113  
  If non-stencil aspects of `srcImage` are accessed, `srcImage` **must**
  have been created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-09114"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-09114"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-09114  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, the `x`, `y`, and
  `z` members of the `srcOffset` member of each element of `pRegions`
  **must** be `0`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-09115"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-09115"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-09115  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, the `extent`
  member of each element of `pRegions` **must** equal the extents of
  `srcImage` identified by `srcSubresource`

<!-- -->

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07966"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07966"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcSubresource-07967"
  id="VUID-VkCopyImageToImageInfoEXT-srcSubresource-07967"></a>
  VUID-VkCopyImageToImageInfoEXT-srcSubresource-07967  
  The `srcSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcSubresource-07968"
  id="VUID-VkCopyImageToImageInfoEXT-srcSubresource-07968"></a>
  VUID-VkCopyImageToImageInfoEXT-srcSubresource-07968  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07969"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07969"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcSubresource-07970"
  id="VUID-VkCopyImageToImageInfoEXT-srcSubresource-07970"></a>
  VUID-VkCopyImageToImageInfoEXT-srcSubresource-07970  
  The image region specified by each element of `pRegions` **must** be
  contained within the specified `srcSubresource` of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcSubresource-07971"
  id="VUID-VkCopyImageToImageInfoEXT-srcSubresource-07971"></a>
  VUID-VkCopyImageToImageInfoEXT-srcSubresource-07971  
  For each element of `pRegions`, `srcOffset.x` and (`extent.width` +
  `srcOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcSubresource-07972"
  id="VUID-VkCopyImageToImageInfoEXT-srcSubresource-07972"></a>
  VUID-VkCopyImageToImageInfoEXT-srcSubresource-07972  
  For each element of `pRegions`, `srcOffset.y` and (`extent.height` +
  `srcOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `srcSubresource` of
  `srcImage`

<!-- -->

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07979"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07979"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07979  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-09104"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-09104"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-09104  
  For each element of `pRegions`, `srcOffset.z` and (`extent.depth` +
  `srcOffset.z`) **must** both be greater than or equal to `0` and less
  than or equal to the depth of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07980"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07980"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07980  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `srcOffset.z` **must** be `0` and
  `extent.depth` **must** be `1`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07274"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07274"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07274  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `srcOffset.x` **must** be a
  multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-10051"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-10051"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-10051  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `srcOffset.x` does not
  equal the width of the subresource specified by `srcSubresource`,
  `srcOffset.x` **must** be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07275"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07275"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07275  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `srcOffset.y` **must** be a
  multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-10052"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-10052"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-10052  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `srcOffset.y` does not
  equal the height of the subresource specified by `srcSubresource`,
  `srcOffset.y` **must** be a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07276"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07276"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07276  
  For each element of `pRegions`, `srcOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-00207"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-00207"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-00207  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of
  `srcOffset.x` and `extent.width` does not equal the width of the
  subresource specified by `srcSubresource`, `extent.width` **must** be
  a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-10053"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-10053"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-10053  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference
  of `srcOffset.x` and `extent.height` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-10054"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-10054"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-10054  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `srcOffset.x` and `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-10055"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-10055"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-10055  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of
  `srcOffset.x` and `extent.height` does not equal the width of the
  subresource specified by `srcSubresource`, `extent.height` **must** be
  a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-00208"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-00208"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-00208  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of
  `srcOffset.y` and `extent.height` does not equal the height of the
  subresource specified by `srcSubresource`, `extent.height` **must** be
  a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-10056"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-10056"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-10056  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of
  `srcOffset.y` and `extent.width` does not equal the height of the
  subresource specified by `srcSubresource`, `extent.width` **must** be
  a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-10057"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-10057"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-10057  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `srcOffset.y` and `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcOffset-10058"
  id="VUID-VkCopyImageToImageInfoEXT-srcOffset-10058"></a>
  VUID-VkCopyImageToImageInfoEXT-srcOffset-10058  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference
  of `srcOffset.y` and `extent.width` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-00209"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-00209"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-00209  
  For each element of `pRegions`, if the sum of `srcOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcSubresource-09105"
  id="VUID-VkCopyImageToImageInfoEXT-srcSubresource-09105"></a>
  VUID-VkCopyImageToImageInfoEXT-srcSubresource-09105  
  For each element of `pRegions`, `srcSubresource.aspectMask` **must**
  specify aspects present in `srcImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07981"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07981"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07981  
  If `srcImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `srcSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-07983"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-07983"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-07983  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, for each element of
  `pRegions`, `srcSubresource.baseArrayLayer` **must** be `0` and
  `srcSubresource.layerCount` **must** be `1`

<!-- -->

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-09109"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-09109"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-09109  
  If `dstImage` is sparse then all memory ranges accessed by the copy
  command **must** be bound as described in [Binding Resource
  Memory](#sparsememory-resource-binding)

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-09111"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-09111"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-09111  
  If the stencil aspect of `dstImage` is accessed, and `dstImage` was
  not created with [separate stencil
  usage](#VkImageStencilUsageCreateInfo), `dstImage` **must** have been
  created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-09112"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-09112"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-09112  
  If the stencil aspect of `dstImage` is accessed, and `dstImage` was
  created with [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-09113"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-09113"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-09113  
  If non-stencil aspects of `dstImage` are accessed, `dstImage` **must**
  have been created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-09114"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-09114"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-09114  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, the `x`, `y`, and
  `z` members of the `dstOffset` member of each element of `pRegions`
  **must** be `0`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-09115"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-09115"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-09115  
  If `flags` contains `VK_HOST_IMAGE_COPY_MEMCPY_EXT`, the `extent`
  member of each element of `pRegions` **must** equal the extents of
  `dstImage` identified by `dstSubresource`

<!-- -->

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07966"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07966"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstSubresource-07967"
  id="VUID-VkCopyImageToImageInfoEXT-dstSubresource-07967"></a>
  VUID-VkCopyImageToImageInfoEXT-dstSubresource-07967  
  The `dstSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstSubresource-07968"
  id="VUID-VkCopyImageToImageInfoEXT-dstSubresource-07968"></a>
  VUID-VkCopyImageToImageInfoEXT-dstSubresource-07968  
  If `dstSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07969"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07969"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstSubresource-07970"
  id="VUID-VkCopyImageToImageInfoEXT-dstSubresource-07970"></a>
  VUID-VkCopyImageToImageInfoEXT-dstSubresource-07970  
  The image region specified by each element of `pRegions` **must** be
  contained within the specified `dstSubresource` of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstSubresource-07971"
  id="VUID-VkCopyImageToImageInfoEXT-dstSubresource-07971"></a>
  VUID-VkCopyImageToImageInfoEXT-dstSubresource-07971  
  For each element of `pRegions`, `dstOffset.x` and (`extent.width` +
  `dstOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstSubresource-07972"
  id="VUID-VkCopyImageToImageInfoEXT-dstSubresource-07972"></a>
  VUID-VkCopyImageToImageInfoEXT-dstSubresource-07972  
  For each element of `pRegions`, `dstOffset.y` and (`extent.height` +
  `dstOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `dstSubresource` of
  `dstImage`

<!-- -->

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07979"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07979"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07979  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-09104"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-09104"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-09104  
  For each element of `pRegions`, `dstOffset.z` and (`extent.depth` +
  `dstOffset.z`) **must** both be greater than or equal to `0` and less
  than or equal to the depth of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07980"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07980"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07980  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `dstOffset.z` **must** be `0` and
  `extent.depth` **must** be `1`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07274"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07274"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07274  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `dstOffset.x` **must** be a
  multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-10051"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-10051"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-10051  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `dstOffset.x` does not
  equal the width of the subresource specified by `dstSubresource`,
  `dstOffset.x` **must** be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07275"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07275"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07275  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `dstOffset.y` **must** be a
  multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-10052"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-10052"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-10052  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `dstOffset.y` does not
  equal the height of the subresource specified by `dstSubresource`,
  `dstOffset.y` **must** be a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07276"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07276"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07276  
  For each element of `pRegions`, `dstOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-00207"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-00207"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-00207  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of
  `dstOffset.x` and `extent.width` does not equal the width of the
  subresource specified by `dstSubresource`, `extent.width` **must** be
  a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-10053"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-10053"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-10053  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference
  of `dstOffset.x` and `extent.height` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-10054"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-10054"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-10054  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `dstOffset.x` and `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-10055"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-10055"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-10055  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of
  `dstOffset.x` and `extent.height` does not equal the width of the
  subresource specified by `dstSubresource`, `extent.height` **must** be
  a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-00208"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-00208"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-00208  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of
  `dstOffset.y` and `extent.height` does not equal the height of the
  subresource specified by `dstSubresource`, `extent.height` **must** be
  a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-10056"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-10056"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-10056  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of
  `dstOffset.y` and `extent.width` does not equal the height of the
  subresource specified by `dstSubresource`, `extent.width` **must** be
  a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-10057"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-10057"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-10057  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `dstOffset.y` and `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstOffset-10058"
  id="VUID-VkCopyImageToImageInfoEXT-dstOffset-10058"></a>
  VUID-VkCopyImageToImageInfoEXT-dstOffset-10058  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference
  of `dstOffset.y` and `extent.width` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-00209"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-00209"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-00209  
  For each element of `pRegions`, if the sum of `dstOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstSubresource-09105"
  id="VUID-VkCopyImageToImageInfoEXT-dstSubresource-09105"></a>
  VUID-VkCopyImageToImageInfoEXT-dstSubresource-09105  
  For each element of `pRegions`, `dstSubresource.aspectMask` **must**
  specify aspects present in `dstImage`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07981"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07981"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07981  
  If `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `dstSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-07983"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-07983"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-07983  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, for each element of
  `pRegions`, `dstSubresource.baseArrayLayer` **must** be `0` and
  `dstSubresource.layerCount` **must** be `1`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImageLayout-09070"
  id="VUID-VkCopyImageToImageInfoEXT-srcImageLayout-09070"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImageLayout-09070  
  `srcImageLayout` **must** specify the current layout of the image
  subresources of `srcImage` specified in `pRegions`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImageLayout-09071"
  id="VUID-VkCopyImageToImageInfoEXT-dstImageLayout-09071"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImageLayout-09071  
  `dstImageLayout` **must** specify the current layout of the image
  subresources of `dstImage` specified in `pRegions`

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImageLayout-09072"
  id="VUID-VkCopyImageToImageInfoEXT-srcImageLayout-09072"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImageLayout-09072  
  `srcImageLayout` **must** be one of the image layouts returned in
  [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)::`pCopySrcLayouts`

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImageLayout-09073"
  id="VUID-VkCopyImageToImageInfoEXT-dstImageLayout-09073"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImageLayout-09073  
  `dstImageLayout` **must** be one of the image layouts returned in
  [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)::`pCopyDstLayouts`

Valid Usage (Implicit)

- <a href="#VUID-VkCopyImageToImageInfoEXT-sType-sType"
  id="VUID-VkCopyImageToImageInfoEXT-sType-sType"></a>
  VUID-VkCopyImageToImageInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_IMAGE_INFO_EXT`

- <a href="#VUID-VkCopyImageToImageInfoEXT-pNext-pNext"
  id="VUID-VkCopyImageToImageInfoEXT-pNext-pNext"></a>
  VUID-VkCopyImageToImageInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyImageToImageInfoEXT-flags-parameter"
  id="VUID-VkCopyImageToImageInfoEXT-flags-parameter"></a>
  VUID-VkCopyImageToImageInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkHostImageCopyFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagBitsEXT.html) values

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImage-parameter"
  id="VUID-VkCopyImageToImageInfoEXT-srcImage-parameter"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkCopyImageToImageInfoEXT-srcImageLayout-parameter"
  id="VUID-VkCopyImageToImageInfoEXT-srcImageLayout-parameter"></a>
  VUID-VkCopyImageToImageInfoEXT-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImage-parameter"
  id="VUID-VkCopyImageToImageInfoEXT-dstImage-parameter"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkCopyImageToImageInfoEXT-dstImageLayout-parameter"
  id="VUID-VkCopyImageToImageInfoEXT-dstImageLayout-parameter"></a>
  VUID-VkCopyImageToImageInfoEXT-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkCopyImageToImageInfoEXT-pRegions-parameter"
  id="VUID-VkCopyImageToImageInfoEXT-pRegions-parameter"></a>
  VUID-VkCopyImageToImageInfoEXT-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy2.html) structures

- <a href="#VUID-VkCopyImageToImageInfoEXT-regionCount-arraylength"
  id="VUID-VkCopyImageToImageInfoEXT-regionCount-arraylength"></a>
  VUID-VkCopyImageToImageInfoEXT-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-VkCopyImageToImageInfoEXT-commonparent"
  id="VUID-VkCopyImageToImageInfoEXT-commonparent"></a>
  VUID-VkCopyImageToImageInfoEXT-commonparent  
  Both of `dstImage`, and `srcImage` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkHostImageCopyFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagsEXT.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy2.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCopyImageToImageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyImageToImageEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyImageToImageInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

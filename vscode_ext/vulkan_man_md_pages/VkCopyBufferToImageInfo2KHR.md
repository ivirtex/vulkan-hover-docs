# VkCopyBufferToImageInfo2(3) Manual Page

## Name

VkCopyBufferToImageInfo2 - Structure specifying parameters of a buffer
to image copy command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyBufferToImageInfo2` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_copy_commands2
typedef VkCopyBufferToImageInfo2 VkCopyBufferToImageInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcBuffer` is the source buffer.

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the copy.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy2.html) structures specifying
  the regions to copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyBufferToImageInfo2-pRegions-04565"
  id="VUID-VkCopyBufferToImageInfo2-pRegions-04565"></a>
  VUID-VkCopyBufferToImageInfo2-pRegions-04565  
  The image region specified by each element of `pRegions` that does not
  contain
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)
  in its `pNext` chain **must** be contained within the specified
  `imageSubresource` of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2KHR-pRegions-04554"
  id="VUID-VkCopyBufferToImageInfo2KHR-pRegions-04554"></a>
  VUID-VkCopyBufferToImageInfo2KHR-pRegions-04554  
  If the image region specified by each element of `pRegions` contains
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)
  in its `pNext` chain, the rotated destination region as described in
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-buffers-images-rotation-addressing"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-buffers-images-rotation-addressing</a>
  **must** be contained within `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2KHR-pRegions-04555"
  id="VUID-VkCopyBufferToImageInfo2KHR-pRegions-04555"></a>
  VUID-VkCopyBufferToImageInfo2KHR-pRegions-04555  
  If any element of `pRegions` contains
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)
  in its `pNext` chain, then `dstImage` **must** have a 1x1x1 <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility-classes"
  target="_blank" rel="noopener">texel block extent</a>

- <a href="#VUID-VkCopyBufferToImageInfo2KHR-pRegions-06203"
  id="VUID-VkCopyBufferToImageInfo2KHR-pRegions-06203"></a>
  VUID-VkCopyBufferToImageInfo2KHR-pRegions-06203  
  If any element of `pRegions` contains
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)
  in its `pNext` chain, then `dstImage` **must** be of type
  `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkCopyBufferToImageInfo2KHR-pRegions-06204"
  id="VUID-VkCopyBufferToImageInfo2KHR-pRegions-06204"></a>
  VUID-VkCopyBufferToImageInfo2KHR-pRegions-06204  
  If any element of `pRegions` contains
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)
  in its `pNext` chain, then `dstImage` **must** not have a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">multi-planar format</a>

<!-- -->

- <a href="#VUID-VkCopyBufferToImageInfo2-pRegions-00171"
  id="VUID-VkCopyBufferToImageInfo2-pRegions-00171"></a>
  VUID-VkCopyBufferToImageInfo2-pRegions-00171  
  `srcBuffer` **must** be large enough to contain all buffer locations
  that are accessed according to [Buffer and Image
  Addressing](#copies-buffers-images-addressing), for each element of
  `pRegions`

- <a href="#VUID-VkCopyBufferToImageInfo2-pRegions-00173"
  id="VUID-VkCopyBufferToImageInfo2-pRegions-00173"></a>
  VUID-VkCopyBufferToImageInfo2-pRegions-00173  
  The union of all source regions, and the union of all destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-VkCopyBufferToImageInfo2-srcBuffer-00174"
  id="VUID-VkCopyBufferToImageInfo2-srcBuffer-00174"></a>
  VUID-VkCopyBufferToImageInfo2-srcBuffer-00174  
  `srcBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-01997"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-01997"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-01997  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

- <a href="#VUID-VkCopyBufferToImageInfo2-srcBuffer-00176"
  id="VUID-VkCopyBufferToImageInfo2-srcBuffer-00176"></a>
  VUID-VkCopyBufferToImageInfo2-srcBuffer-00176  
  If `srcBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-00177"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-00177"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-00177  
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImageLayout-00180"
  id="VUID-VkCopyBufferToImageInfo2-dstImageLayout-00180"></a>
  VUID-VkCopyBufferToImageInfo2-dstImageLayout-00180  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImageLayout-01396"
  id="VUID-VkCopyBufferToImageInfo2-dstImageLayout-01396"></a>
  VUID-VkCopyBufferToImageInfo2-dstImageLayout-01396  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-VkCopyBufferToImageInfo2-pRegions-07931"
  id="VUID-VkCopyBufferToImageInfo2-pRegions-07931"></a>
  VUID-VkCopyBufferToImageInfo2-pRegions-07931  
  If
  [VK_EXT_depth_range_unrestricted](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_range_unrestricted.html)
  is not enabled, for each element of `pRegions` whose
  `imageSubresource` contains a depth aspect, the data in `srcBuffer`
  **must** be in the range \[0,1\]

<!-- -->

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07966"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07966"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-VkCopyBufferToImageInfo2-imageSubresource-07967"
  id="VUID-VkCopyBufferToImageInfo2-imageSubresource-07967"></a>
  VUID-VkCopyBufferToImageInfo2-imageSubresource-07967  
  The `imageSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-VkCopyBufferToImageInfo2-imageSubresource-07968"
  id="VUID-VkCopyBufferToImageInfo2-imageSubresource-07968"></a>
  VUID-VkCopyBufferToImageInfo2-imageSubresource-07968  
  If `imageSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `imageSubresource.baseArrayLayer` + `imageSubresource.layerCount` of
  each element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07969"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07969"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07973"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07973"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07973  
  `dstImage` **must** have a sample count equal to
  `VK_SAMPLE_COUNT_1_BIT`

<!-- -->

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07979"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07979"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07979  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `imageOffset.y` **must** be `0` and `imageExtent.height`
  **must** be `1`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-09104"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-09104"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-09104  
  For each element of `pRegions`, `imageOffset.z` and
  (`imageExtent.depth` + `imageOffset.z`) **must** both be greater than
  or equal to `0` and less than or equal to the depth of the specified
  `imageSubresource` of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07980"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07980"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07980  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `imageOffset.z` **must** be `0`
  and `imageExtent.depth` **must** be `1`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07274"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07274"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07274  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, `imageOffset.x` **must** be
  a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-10051"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-10051"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-10051  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, and `imageOffset.x` does not
  equal the width of the subresource specified by `imageSubresource`,
  `imageOffset.x` **must** be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07275"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07275"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07275  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `imageOffset.y` **must** be
  a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-10052"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-10052"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-10052  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` or
  `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, and `imageOffset.y` does
  not equal the height of the subresource specified by
  `imageSubresource`, `imageOffset.y` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07276"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07276"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07276  
  For each element of `pRegions`, `imageOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-00207"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-00207"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-00207  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the sum of
  `imageOffset.x` and `extent.width` does not equal the width of the
  subresource specified by `imageSubresource`, `extent.width` **must**
  be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-10053"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-10053"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-10053  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the difference
  of `imageOffset.x` and `extent.height` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-10054"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-10054"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-10054  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `imageOffset.x` and `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-10055"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-10055"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-10055  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the sum of
  `imageOffset.x` and `extent.height` does not equal the width of the
  subresource specified by `imageSubresource`, `extent.height` **must**
  be a multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-00208"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-00208"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-00208  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, and the sum of
  `imageOffset.y` and `extent.height` does not equal the height of the
  subresource specified by `imageSubresource`, `extent.height` **must**
  be a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-10056"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-10056"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-10056  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, the sum of
  `imageOffset.y` and `extent.width` does not equal the height of the
  subresource specified by `imageSubresource`, `extent.width` **must**
  be a multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-10057"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-10057"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-10057  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, the difference
  of `imageOffset.y` and `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageOffset-10058"
  id="VUID-VkCopyBufferToImageInfo2-imageOffset-10058"></a>
  VUID-VkCopyBufferToImageInfo2-imageOffset-10058  
  For each element of `pRegions`, if
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)::`transform`
  is equal to `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, the difference
  of `imageOffset.y` and `extent.width` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-00209"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-00209"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-00209  
  For each element of `pRegions`, if the sum of `imageOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-imageSubresource-09105"
  id="VUID-VkCopyBufferToImageInfo2-imageSubresource-09105"></a>
  VUID-VkCopyBufferToImageInfo2-imageSubresource-09105  
  For each element of `pRegions`, `imageSubresource.aspectMask` **must**
  specify aspects present in `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07981"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07981"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07981  
  If `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `imageSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07983"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07983"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07983  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, for each element of
  `pRegions`, `imageSubresource.baseArrayLayer` **must** be `0` and
  `imageSubresource.layerCount` **must** be `1`

<!-- -->

- <a href="#VUID-VkCopyBufferToImageInfo2-bufferRowLength-09106"
  id="VUID-VkCopyBufferToImageInfo2-bufferRowLength-09106"></a>
  VUID-VkCopyBufferToImageInfo2-bufferRowLength-09106  
  For each element of `pRegions`, `bufferRowLength` **must** be a
  multiple of the [texel block extent
  width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-bufferImageHeight-09107"
  id="VUID-VkCopyBufferToImageInfo2-bufferImageHeight-09107"></a>
  VUID-VkCopyBufferToImageInfo2-bufferImageHeight-09107  
  For each element of `pRegions`, `bufferImageHeight` **must** be a
  multiple of the [texel block extent
  height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-bufferRowLength-09108"
  id="VUID-VkCopyBufferToImageInfo2-bufferRowLength-09108"></a>
  VUID-VkCopyBufferToImageInfo2-bufferRowLength-09108  
  For each element of `pRegions`, `bufferRowLength` divided by the
  [texel block extent width](#formats-compatibility-classes) and then
  multiplied by the texel block size of `dstImage` **must** be less than
  or equal to 2<sup>31</sup>-1

<!-- -->

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07975"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07975"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07975  
  If `dstImage` does not have either a depth/stencil format or a
  [multi-planar format](#formats-requiring-sampler-ycbcr-conversion),
  then for each element of `pRegions`, `bufferOffset` **must** be a
  multiple of the [texel block size](#formats-compatibility-classes)

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07976"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07976"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07976  
  If `dstImage` has a [multi-planar
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `bufferOffset` **must** be a multiple of the
  element size of the compatible format for the format and the
  `aspectMask` of the `imageSubresource` as defined in
  [\[formats-compatible-planes\]](#formats-compatible-planes)

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-07978"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-07978"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-07978  
  If `dstImage` has a depth/stencil format, the `bufferOffset` member of
  any element of `pRegions` **must** be a multiple of `4`

- <a href="#VUID-VkCopyBufferToImageInfo2-pRegions-06223"
  id="VUID-VkCopyBufferToImageInfo2-pRegions-06223"></a>
  VUID-VkCopyBufferToImageInfo2-pRegions-06223  
  For each element of `pRegions` not containing
  `VkCopyCommandTransformInfoQCOM` in its `pNext` chain, `imageOffset.x`
  and (`imageExtent.width` + `imageOffset.x`) **must** both be greater
  than or equal to `0` and less than or equal to the width of the
  specified `imageSubresource` of `dstImage`

- <a href="#VUID-VkCopyBufferToImageInfo2-pRegions-06224"
  id="VUID-VkCopyBufferToImageInfo2-pRegions-06224"></a>
  VUID-VkCopyBufferToImageInfo2-pRegions-06224  
  For each element of `pRegions` not containing
  `VkCopyCommandTransformInfoQCOM` in its `pNext` chain, `imageOffset.y`
  and (`imageExtent.height` + `imageOffset.y`) **must** both be greater
  than or equal to `0` and less than or equal to the height of the
  specified `imageSubresource` of `dstImage`

Valid Usage (Implicit)

- <a href="#VUID-VkCopyBufferToImageInfo2-sType-sType"
  id="VUID-VkCopyBufferToImageInfo2-sType-sType"></a>
  VUID-VkCopyBufferToImageInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_BUFFER_TO_IMAGE_INFO_2`

- <a href="#VUID-VkCopyBufferToImageInfo2-pNext-pNext"
  id="VUID-VkCopyBufferToImageInfo2-pNext-pNext"></a>
  VUID-VkCopyBufferToImageInfo2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyBufferToImageInfo2-srcBuffer-parameter"
  id="VUID-VkCopyBufferToImageInfo2-srcBuffer-parameter"></a>
  VUID-VkCopyBufferToImageInfo2-srcBuffer-parameter  
  `srcBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImage-parameter"
  id="VUID-VkCopyBufferToImageInfo2-dstImage-parameter"></a>
  VUID-VkCopyBufferToImageInfo2-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkCopyBufferToImageInfo2-dstImageLayout-parameter"
  id="VUID-VkCopyBufferToImageInfo2-dstImageLayout-parameter"></a>
  VUID-VkCopyBufferToImageInfo2-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkCopyBufferToImageInfo2-pRegions-parameter"
  id="VUID-VkCopyBufferToImageInfo2-pRegions-parameter"></a>
  VUID-VkCopyBufferToImageInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy2.html) structures

- <a href="#VUID-VkCopyBufferToImageInfo2-regionCount-arraylength"
  id="VUID-VkCopyBufferToImageInfo2-regionCount-arraylength"></a>
  VUID-VkCopyBufferToImageInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-VkCopyBufferToImageInfo2-commonparent"
  id="VUID-VkCopyBufferToImageInfo2-commonparent"></a>
  VUID-VkCopyBufferToImageInfo2-commonparent  
  Both of `dstImage`, and `srcBuffer` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy2.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage2.html),
[vkCmdCopyBufferToImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyBufferToImageInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

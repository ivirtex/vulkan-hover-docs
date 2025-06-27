# VkBlitImageInfo2(3) Manual Page

## Name

VkBlitImageInfo2 - Structure specifying parameters of blit image command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBlitImageInfo2` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_copy_commands2
typedef VkBlitImageInfo2 VkBlitImageInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcImage` is the source image.

- `srcImageLayout` is the layout of the source image subresources for
  the blit.

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the blit.

- `regionCount` is the number of regions to blit.

- `pRegions` is a pointer to an array of
  [VkImageBlit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit2.html) structures specifying the regions to
  blit.

- `filter` is a [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) specifying the filter to apply
  if the blits require scaling.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBlitImageInfo2-pRegions-00215"
  id="VUID-VkBlitImageInfo2-pRegions-00215"></a>
  VUID-VkBlitImageInfo2-pRegions-00215  
  The source region specified by each element of `pRegions` **must** be
  a region that is contained within `srcImage`

- <a href="#VUID-VkBlitImageInfo2-pRegions-00216"
  id="VUID-VkBlitImageInfo2-pRegions-00216"></a>
  VUID-VkBlitImageInfo2-pRegions-00216  
  The destination region specified by each element of `pRegions`
  **must** be a region that is contained within `dstImage`

- <a href="#VUID-VkBlitImageInfo2-pRegions-00217"
  id="VUID-VkBlitImageInfo2-pRegions-00217"></a>
  VUID-VkBlitImageInfo2-pRegions-00217  
  The union of all destination regions, specified by the elements of
  `pRegions`, **must** not overlap in memory with any texel that **may**
  be sampled during the blit operation

- <a href="#VUID-VkBlitImageInfo2-srcImage-01999"
  id="VUID-VkBlitImageInfo2-srcImage-01999"></a>
  VUID-VkBlitImageInfo2-srcImage-01999  
  The [format features](#resources-image-format-features) of `srcImage`
  **must** contain `VK_FORMAT_FEATURE_BLIT_SRC_BIT`

- <a href="#VUID-VkBlitImageInfo2-srcImage-06421"
  id="VUID-VkBlitImageInfo2-srcImage-06421"></a>
  VUID-VkBlitImageInfo2-srcImage-06421  
  `srcImage` **must** not use a [format that requires a sampler
  Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#formats-requiring-sampler-ycbcr-conversion)

- <a href="#VUID-VkBlitImageInfo2-srcImage-00219"
  id="VUID-VkBlitImageInfo2-srcImage-00219"></a>
  VUID-VkBlitImageInfo2-srcImage-00219  
  `srcImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-VkBlitImageInfo2-srcImage-00220"
  id="VUID-VkBlitImageInfo2-srcImage-00220"></a>
  VUID-VkBlitImageInfo2-srcImage-00220  
  If `srcImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkBlitImageInfo2-srcImageLayout-00221"
  id="VUID-VkBlitImageInfo2-srcImageLayout-00221"></a>
  VUID-VkBlitImageInfo2-srcImageLayout-00221  
  `srcImageLayout` **must** specify the layout of the image subresources
  of `srcImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-VkBlitImageInfo2-srcImageLayout-01398"
  id="VUID-VkBlitImageInfo2-srcImageLayout-01398"></a>
  VUID-VkBlitImageInfo2-srcImageLayout-01398  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-VkBlitImageInfo2-srcImage-09459"
  id="VUID-VkBlitImageInfo2-srcImage-09459"></a>
  VUID-VkBlitImageInfo2-srcImage-09459  
  If `srcImage` and `dstImage` are the same, and an elements of
  `pRegions` contains the `srcSubresource` and `dstSubresource` with
  matching `mipLevel` and overlapping array layers, then the
  `srcImageLayout` and `dstImageLayout` **must** be
  `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`

- <a href="#VUID-VkBlitImageInfo2-dstImage-02000"
  id="VUID-VkBlitImageInfo2-dstImage-02000"></a>
  VUID-VkBlitImageInfo2-dstImage-02000  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_BLIT_DST_BIT`

- <a href="#VUID-VkBlitImageInfo2-dstImage-06422"
  id="VUID-VkBlitImageInfo2-dstImage-06422"></a>
  VUID-VkBlitImageInfo2-dstImage-06422  
  `dstImage` **must** not use a [format that requires a sampler
  Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#formats-requiring-sampler-ycbcr-conversion)

- <a href="#VUID-VkBlitImageInfo2-dstImage-00224"
  id="VUID-VkBlitImageInfo2-dstImage-00224"></a>
  VUID-VkBlitImageInfo2-dstImage-00224  
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-VkBlitImageInfo2-dstImage-00225"
  id="VUID-VkBlitImageInfo2-dstImage-00225"></a>
  VUID-VkBlitImageInfo2-dstImage-00225  
  If `dstImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkBlitImageInfo2-dstImageLayout-00226"
  id="VUID-VkBlitImageInfo2-dstImageLayout-00226"></a>
  VUID-VkBlitImageInfo2-dstImageLayout-00226  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-VkBlitImageInfo2-dstImageLayout-01399"
  id="VUID-VkBlitImageInfo2-dstImageLayout-01399"></a>
  VUID-VkBlitImageInfo2-dstImageLayout-01399  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-VkBlitImageInfo2-srcImage-00229"
  id="VUID-VkBlitImageInfo2-srcImage-00229"></a>
  VUID-VkBlitImageInfo2-srcImage-00229  
  If either of `srcImage` or `dstImage` was created with a signed
  integer [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), the other **must** also have been
  created with a signed integer [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)

- <a href="#VUID-VkBlitImageInfo2-srcImage-00230"
  id="VUID-VkBlitImageInfo2-srcImage-00230"></a>
  VUID-VkBlitImageInfo2-srcImage-00230  
  If either of `srcImage` or `dstImage` was created with an unsigned
  integer [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), the other **must** also have been
  created with an unsigned integer [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)

- <a href="#VUID-VkBlitImageInfo2-srcImage-00231"
  id="VUID-VkBlitImageInfo2-srcImage-00231"></a>
  VUID-VkBlitImageInfo2-srcImage-00231  
  If either of `srcImage` or `dstImage` was created with a depth/stencil
  format, the other **must** have exactly the same format

- <a href="#VUID-VkBlitImageInfo2-srcImage-00232"
  id="VUID-VkBlitImageInfo2-srcImage-00232"></a>
  VUID-VkBlitImageInfo2-srcImage-00232  
  If `srcImage` was created with a depth/stencil format, `filter`
  **must** be `VK_FILTER_NEAREST`

- <a href="#VUID-VkBlitImageInfo2-srcImage-00233"
  id="VUID-VkBlitImageInfo2-srcImage-00233"></a>
  VUID-VkBlitImageInfo2-srcImage-00233  
  `srcImage` **must** have been created with a `samples` value of
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkBlitImageInfo2-dstImage-00234"
  id="VUID-VkBlitImageInfo2-dstImage-00234"></a>
  VUID-VkBlitImageInfo2-dstImage-00234  
  `dstImage` **must** have been created with a `samples` value of
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkBlitImageInfo2-filter-02001"
  id="VUID-VkBlitImageInfo2-filter-02001"></a>
  VUID-VkBlitImageInfo2-filter-02001  
  If `filter` is `VK_FILTER_LINEAR`, then the [format
  features](#resources-image-format-features) of `srcImage` **must**
  contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <a href="#VUID-VkBlitImageInfo2-filter-02002"
  id="VUID-VkBlitImageInfo2-filter-02002"></a>
  VUID-VkBlitImageInfo2-filter-02002  
  If `filter` is `VK_FILTER_CUBIC_EXT`, then the [format
  features](#resources-image-format-features) of `srcImage` **must**
  contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`

- <a href="#VUID-VkBlitImageInfo2-filter-00237"
  id="VUID-VkBlitImageInfo2-filter-00237"></a>
  VUID-VkBlitImageInfo2-filter-00237  
  If `filter` is `VK_FILTER_CUBIC_EXT`, `srcImage` **must** be of type
  `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkBlitImageInfo2-srcSubresource-01705"
  id="VUID-VkBlitImageInfo2-srcSubresource-01705"></a>
  VUID-VkBlitImageInfo2-srcSubresource-01705  
  The `srcSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-VkBlitImageInfo2-dstSubresource-01706"
  id="VUID-VkBlitImageInfo2-dstSubresource-01706"></a>
  VUID-VkBlitImageInfo2-dstSubresource-01706  
  The `dstSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-VkBlitImageInfo2-srcSubresource-01707"
  id="VUID-VkBlitImageInfo2-srcSubresource-01707"></a>
  VUID-VkBlitImageInfo2-srcSubresource-01707  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-VkBlitImageInfo2-dstSubresource-01708"
  id="VUID-VkBlitImageInfo2-dstSubresource-01708"></a>
  VUID-VkBlitImageInfo2-dstSubresource-01708  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-VkBlitImageInfo2-dstImage-02545"
  id="VUID-VkBlitImageInfo2-dstImage-02545"></a>
  VUID-VkBlitImageInfo2-dstImage-02545  
  `dstImage` and `srcImage` **must** not have been created with `flags`
  containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- <a href="#VUID-VkBlitImageInfo2-srcImage-00240"
  id="VUID-VkBlitImageInfo2-srcImage-00240"></a>
  VUID-VkBlitImageInfo2-srcImage-00240  
  If either `srcImage` or `dstImage` is of type `VK_IMAGE_TYPE_3D`, then
  for each element of `pRegions`, `srcSubresource.baseArrayLayer` and
  `dstSubresource.baseArrayLayer` **must** each be `0`, and
  `srcSubresource.layerCount` and `dstSubresource.layerCount` **must**
  each be `1`

- <a href="#VUID-VkBlitImageInfo2-aspectMask-00241"
  id="VUID-VkBlitImageInfo2-aspectMask-00241"></a>
  VUID-VkBlitImageInfo2-aspectMask-00241  
  For each element of `pRegions`, `srcSubresource.aspectMask` **must**
  specify aspects present in `srcImage`

- <a href="#VUID-VkBlitImageInfo2-aspectMask-00242"
  id="VUID-VkBlitImageInfo2-aspectMask-00242"></a>
  VUID-VkBlitImageInfo2-aspectMask-00242  
  For each element of `pRegions`, `dstSubresource.aspectMask` **must**
  specify aspects present in `dstImage`

- <a href="#VUID-VkBlitImageInfo2-srcOffset-00243"
  id="VUID-VkBlitImageInfo2-srcOffset-00243"></a>
  VUID-VkBlitImageInfo2-srcOffset-00243  
  For each element of `pRegions`, `srcOffsets`\[0\].x and
  `srcOffsets`\[1\].x **must** both be greater than or equal to `0` and
  less than or equal to the width of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkBlitImageInfo2-srcOffset-00244"
  id="VUID-VkBlitImageInfo2-srcOffset-00244"></a>
  VUID-VkBlitImageInfo2-srcOffset-00244  
  For each element of `pRegions`, `srcOffsets`\[0\].y and
  `srcOffsets`\[1\].y **must** both be greater than or equal to `0` and
  less than or equal to the height of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkBlitImageInfo2-srcImage-00245"
  id="VUID-VkBlitImageInfo2-srcImage-00245"></a>
  VUID-VkBlitImageInfo2-srcImage-00245  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffsets`\[0\].y **must** be `0` and
  `srcOffsets`\[1\].y **must** be `1`

- <a href="#VUID-VkBlitImageInfo2-srcOffset-00246"
  id="VUID-VkBlitImageInfo2-srcOffset-00246"></a>
  VUID-VkBlitImageInfo2-srcOffset-00246  
  For each element of `pRegions`, `srcOffsets`\[0\].z and
  `srcOffsets`\[1\].z **must** both be greater than or equal to `0` and
  less than or equal to the depth of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkBlitImageInfo2-srcImage-00247"
  id="VUID-VkBlitImageInfo2-srcImage-00247"></a>
  VUID-VkBlitImageInfo2-srcImage-00247  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `srcOffsets`\[0\].z **must** be
  `0` and `srcOffsets`\[1\].z **must** be `1`

- <a href="#VUID-VkBlitImageInfo2-dstOffset-00248"
  id="VUID-VkBlitImageInfo2-dstOffset-00248"></a>
  VUID-VkBlitImageInfo2-dstOffset-00248  
  For each element of `pRegions`, `dstOffsets`\[0\].x and
  `dstOffsets`\[1\].x **must** both be greater than or equal to `0` and
  less than or equal to the width of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkBlitImageInfo2-dstOffset-00249"
  id="VUID-VkBlitImageInfo2-dstOffset-00249"></a>
  VUID-VkBlitImageInfo2-dstOffset-00249  
  For each element of `pRegions`, `dstOffsets`\[0\].y and
  `dstOffsets`\[1\].y **must** both be greater than or equal to `0` and
  less than or equal to the height of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkBlitImageInfo2-dstImage-00250"
  id="VUID-VkBlitImageInfo2-dstImage-00250"></a>
  VUID-VkBlitImageInfo2-dstImage-00250  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffsets`\[0\].y **must** be `0` and
  `dstOffsets`\[1\].y **must** be `1`

- <a href="#VUID-VkBlitImageInfo2-dstOffset-00251"
  id="VUID-VkBlitImageInfo2-dstOffset-00251"></a>
  VUID-VkBlitImageInfo2-dstOffset-00251  
  For each element of `pRegions`, `dstOffsets`\[0\].z and
  `dstOffsets`\[1\].z **must** both be greater than or equal to `0` and
  less than or equal to the depth of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkBlitImageInfo2-dstImage-00252"
  id="VUID-VkBlitImageInfo2-dstImage-00252"></a>
  VUID-VkBlitImageInfo2-dstImage-00252  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `dstOffsets`\[0\].z **must** be
  `0` and `dstOffsets`\[1\].z **must** be `1`

- <a href="#VUID-VkBlitImageInfo2-pRegions-04561"
  id="VUID-VkBlitImageInfo2-pRegions-04561"></a>
  VUID-VkBlitImageInfo2-pRegions-04561  
  If any element of `pRegions` contains
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)
  in its `pNext` chain, then `srcImage` and `dstImage` **must** not be
  block-compressed images

- <a href="#VUID-VkBlitImageInfo2KHR-pRegions-06207"
  id="VUID-VkBlitImageInfo2KHR-pRegions-06207"></a>
  VUID-VkBlitImageInfo2KHR-pRegions-06207  
  If any element of `pRegions` contains
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)
  in its `pNext` chain, then `srcImage` **must** be of type
  `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkBlitImageInfo2KHR-pRegions-06208"
  id="VUID-VkBlitImageInfo2KHR-pRegions-06208"></a>
  VUID-VkBlitImageInfo2KHR-pRegions-06208  
  If any element of `pRegions` contains
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)
  in its `pNext` chain, then `srcImage` **must** not have a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">multi-planar format</a>

- <a href="#VUID-VkBlitImageInfo2-filter-09204"
  id="VUID-VkBlitImageInfo2-filter-09204"></a>
  VUID-VkBlitImageInfo2-filter-09204  
  If `filter` is `VK_FILTER_CUBIC_EXT` and if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-filter-cubic-weight-selection"
  target="_blank" rel="noopener">selectableCubicWeights</a> feature is
  not enabled then the cubic weights **must** be
  `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`

Valid Usage (Implicit)

- <a href="#VUID-VkBlitImageInfo2-sType-sType"
  id="VUID-VkBlitImageInfo2-sType-sType"></a>
  VUID-VkBlitImageInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BLIT_IMAGE_INFO_2`

- <a href="#VUID-VkBlitImageInfo2-pNext-pNext"
  id="VUID-VkBlitImageInfo2-pNext-pNext"></a>
  VUID-VkBlitImageInfo2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkBlitImageCubicWeightsInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageCubicWeightsInfoQCOM.html)

- <a href="#VUID-VkBlitImageInfo2-sType-unique"
  id="VUID-VkBlitImageInfo2-sType-unique"></a>
  VUID-VkBlitImageInfo2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkBlitImageInfo2-srcImage-parameter"
  id="VUID-VkBlitImageInfo2-srcImage-parameter"></a>
  VUID-VkBlitImageInfo2-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkBlitImageInfo2-srcImageLayout-parameter"
  id="VUID-VkBlitImageInfo2-srcImageLayout-parameter"></a>
  VUID-VkBlitImageInfo2-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkBlitImageInfo2-dstImage-parameter"
  id="VUID-VkBlitImageInfo2-dstImage-parameter"></a>
  VUID-VkBlitImageInfo2-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkBlitImageInfo2-dstImageLayout-parameter"
  id="VUID-VkBlitImageInfo2-dstImageLayout-parameter"></a>
  VUID-VkBlitImageInfo2-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkBlitImageInfo2-pRegions-parameter"
  id="VUID-VkBlitImageInfo2-pRegions-parameter"></a>
  VUID-VkBlitImageInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkImageBlit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit2.html) structures

- <a href="#VUID-VkBlitImageInfo2-filter-parameter"
  id="VUID-VkBlitImageInfo2-filter-parameter"></a>
  VUID-VkBlitImageInfo2-filter-parameter  
  `filter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) value

- <a href="#VUID-VkBlitImageInfo2-regionCount-arraylength"
  id="VUID-VkBlitImageInfo2-regionCount-arraylength"></a>
  VUID-VkBlitImageInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-VkBlitImageInfo2-commonparent"
  id="VUID-VkBlitImageInfo2-commonparent"></a>
  VUID-VkBlitImageInfo2-commonparent  
  Both of `dstImage`, and `srcImage` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageBlit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit2.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBlitImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage2.html),
[vkCmdBlitImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBlitImageInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# vkCmdBlitImage(3) Manual Page

## Name

vkCmdBlitImage - Copy regions of an image, potentially performing format
conversion,



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy regions of a source image into a destination image, potentially
performing format conversion, arbitrary scaling, and filtering, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdBlitImage(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     srcImage,
    VkImageLayout                               srcImageLayout,
    VkImage                                     dstImage,
    VkImageLayout                               dstImageLayout,
    uint32_t                                    regionCount,
    const VkImageBlit*                          pRegions,
    VkFilter                                    filter);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `srcImage` is the source image.

- `srcImageLayout` is the layout of the source image subresources for
  the blit.

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the blit.

- `regionCount` is the number of regions to blit.

- `pRegions` is a pointer to an array of [VkImageBlit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit.html)
  structures specifying the regions to blit.

- `filter` is a [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) specifying the filter to apply
  if the blits require scaling.

## <a href="#_description" class="anchor"></a>Description

`vkCmdBlitImage` **must** not be used for multisampled source or
destination images. Use [vkCmdResolveImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage.html) for
this purpose.

As the sizes of the source and destination extents **can** differ in any
dimension, texels in the source extent are scaled and filtered to the
destination extent. Scaling occurs via the following operations:

- For each destination texel, the integer coordinate of that texel is
  converted to an unnormalized texture coordinate, using the effective
  inverse of the equations described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-unnormalized-to-integer"
  target="_blank" rel="noopener">unnormalized to integer conversion</a>:

    
  u<sub>base</sub> = i + ½

    
  v<sub>base</sub> = j + ½

    
  w<sub>base</sub> = k + ½

- These base coordinates are then offset by the first destination
  offset:

    
  u<sub>offset</sub> = u<sub>base</sub> - x<sub>dst0</sub>

    
  v<sub>offset</sub> = v<sub>base</sub> - y<sub>dst0</sub>

    
  w<sub>offset</sub> = w<sub>base</sub> - z<sub>dst0</sub>

    
  a<sub>offset</sub> = a - `baseArrayCount`<sub>dst</sub>

- The scale is determined from the source and destination regions, and
  applied to the offset coordinates:

    
  scale<sub>u</sub> = (x<sub>src1</sub> - x<sub>src0</sub>) /
  (x<sub>dst1</sub> - x<sub>dst0</sub>)

    
  scale<sub>v</sub> = (y<sub>src1</sub> - y<sub>src0</sub>) /
  (y<sub>dst1</sub> - y<sub>dst0</sub>)

    
  scale<sub>w</sub> = (z<sub>src1</sub> - z<sub>src0</sub>) /
  (z<sub>dst1</sub> - z<sub>dst0</sub>)

    
  u<sub>scaled</sub> = u<sub>offset</sub> × scale<sub>u</sub>

    
  v<sub>scaled</sub> = v<sub>offset</sub> × scale<sub>v</sub>

    
  w<sub>scaled</sub> = w<sub>offset</sub> × scale<sub>w</sub>

- Finally the source offset is added to the scaled coordinates, to
  determine the final unnormalized coordinates used to sample from
  `srcImage`:

    
  u = u<sub>scaled</sub> + x<sub>src0</sub>

    
  v = v<sub>scaled</sub> + y<sub>src0</sub>

    
  w = w<sub>scaled</sub> + z<sub>src0</sub>

    
  q = `mipLevel`

    
  a = a<sub>offset</sub> + `baseArrayCount`<sub>src</sub>

These coordinates are used to sample from the source image, as described
in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures"
target="_blank" rel="noopener">Image Operations chapter</a>, with the
filter mode equal to that of `filter`, a mipmap mode of
`VK_SAMPLER_MIPMAP_MODE_NEAREST` and an address mode of
`VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`. Implementations **must** clamp
at the edge of the source image, and **may** additionally clamp to the
edge of the source region.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Due to allowable rounding errors in the generation of the source
texture coordinates, it is not always possible to guarantee exactly
which source texels will be sampled for a given blit. As rounding errors
are implementation-dependent, the exact results of a blitting operation
are also implementation-dependent.</p></td>
</tr>
</tbody>
</table>

Blits are done layer by layer starting with the `baseArrayLayer` member
of `srcSubresource` for the source and `dstSubresource` for the
destination. `layerCount` layers are blitted to the destination image.

When blitting 3D textures, slices in the destination region bounded by
`dstOffsets`\[0\].z and `dstOffsets`\[1\].z are sampled from slices in
the source region bounded by `srcOffsets`\[0\].z and
`srcOffsets`\[1\].z. If the `filter` parameter is `VK_FILTER_LINEAR`
then the value sampled from the source image is taken by doing linear
filtering using the interpolated **z** coordinate represented by **w**
in the previous equations. If the `filter` parameter is
`VK_FILTER_NEAREST` then the value sampled from the source image is
taken from the single nearest slice, with an implementation-dependent
arithmetic rounding mode.

The following filtering and conversion rules apply:

- Integer formats **can** only be converted to other integer formats
  with the same signedness.

- No format conversion is supported between depth/stencil images. The
  formats **must** match.

- Format conversions on unorm, snorm, scaled and packed float formats of
  the copied aspect of the image are performed by first converting the
  pixels to float values.

- For sRGB source formats, nonlinear RGB values are converted to linear
  representation prior to filtering.

- After filtering, the float values are first clamped and then cast to
  the destination image format. In case of sRGB destination format,
  linear RGB values are converted to nonlinear representation before
  writing the pixel to the image.

Signed and unsigned integers are converted by first clamping to the
representable range of the destination format, then casting the value.

Valid Usage

- <a href="#VUID-vkCmdBlitImage-commandBuffer-01834"
  id="VUID-vkCmdBlitImage-commandBuffer-01834"></a>
  VUID-vkCmdBlitImage-commandBuffer-01834  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcImage` **must** not be a protected image

- <a href="#VUID-vkCmdBlitImage-commandBuffer-01835"
  id="VUID-vkCmdBlitImage-commandBuffer-01835"></a>
  VUID-vkCmdBlitImage-commandBuffer-01835  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be a protected image

- <a href="#VUID-vkCmdBlitImage-commandBuffer-01836"
  id="VUID-vkCmdBlitImage-commandBuffer-01836"></a>
  VUID-vkCmdBlitImage-commandBuffer-01836  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be an unprotected image

<!-- -->

- <a href="#VUID-vkCmdBlitImage-pRegions-00215"
  id="VUID-vkCmdBlitImage-pRegions-00215"></a>
  VUID-vkCmdBlitImage-pRegions-00215  
  The source region specified by each element of `pRegions` **must** be
  a region that is contained within `srcImage`

- <a href="#VUID-vkCmdBlitImage-pRegions-00216"
  id="VUID-vkCmdBlitImage-pRegions-00216"></a>
  VUID-vkCmdBlitImage-pRegions-00216  
  The destination region specified by each element of `pRegions`
  **must** be a region that is contained within `dstImage`

- <a href="#VUID-vkCmdBlitImage-pRegions-00217"
  id="VUID-vkCmdBlitImage-pRegions-00217"></a>
  VUID-vkCmdBlitImage-pRegions-00217  
  The union of all destination regions, specified by the elements of
  `pRegions`, **must** not overlap in memory with any texel that **may**
  be sampled during the blit operation

- <a href="#VUID-vkCmdBlitImage-srcImage-01999"
  id="VUID-vkCmdBlitImage-srcImage-01999"></a>
  VUID-vkCmdBlitImage-srcImage-01999  
  The [format features](#resources-image-format-features) of `srcImage`
  **must** contain `VK_FORMAT_FEATURE_BLIT_SRC_BIT`

- <a href="#VUID-vkCmdBlitImage-srcImage-06421"
  id="VUID-vkCmdBlitImage-srcImage-06421"></a>
  VUID-vkCmdBlitImage-srcImage-06421  
  `srcImage` **must** not use a [format that requires a sampler
  Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#formats-requiring-sampler-ycbcr-conversion)

- <a href="#VUID-vkCmdBlitImage-srcImage-00219"
  id="VUID-vkCmdBlitImage-srcImage-00219"></a>
  VUID-vkCmdBlitImage-srcImage-00219  
  `srcImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-vkCmdBlitImage-srcImage-00220"
  id="VUID-vkCmdBlitImage-srcImage-00220"></a>
  VUID-vkCmdBlitImage-srcImage-00220  
  If `srcImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdBlitImage-srcImageLayout-00221"
  id="VUID-vkCmdBlitImage-srcImageLayout-00221"></a>
  VUID-vkCmdBlitImage-srcImageLayout-00221  
  `srcImageLayout` **must** specify the layout of the image subresources
  of `srcImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdBlitImage-srcImageLayout-01398"
  id="VUID-vkCmdBlitImage-srcImageLayout-01398"></a>
  VUID-vkCmdBlitImage-srcImageLayout-01398  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdBlitImage-srcImage-09459"
  id="VUID-vkCmdBlitImage-srcImage-09459"></a>
  VUID-vkCmdBlitImage-srcImage-09459  
  If `srcImage` and `dstImage` are the same, and an elements of
  `pRegions` contains the `srcSubresource` and `dstSubresource` with
  matching `mipLevel` and overlapping array layers, then the
  `srcImageLayout` and `dstImageLayout` **must** be
  `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`

- <a href="#VUID-vkCmdBlitImage-dstImage-02000"
  id="VUID-vkCmdBlitImage-dstImage-02000"></a>
  VUID-vkCmdBlitImage-dstImage-02000  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_BLIT_DST_BIT`

- <a href="#VUID-vkCmdBlitImage-dstImage-06422"
  id="VUID-vkCmdBlitImage-dstImage-06422"></a>
  VUID-vkCmdBlitImage-dstImage-06422  
  `dstImage` **must** not use a [format that requires a sampler
  Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#formats-requiring-sampler-ycbcr-conversion)

- <a href="#VUID-vkCmdBlitImage-dstImage-00224"
  id="VUID-vkCmdBlitImage-dstImage-00224"></a>
  VUID-vkCmdBlitImage-dstImage-00224  
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdBlitImage-dstImage-00225"
  id="VUID-vkCmdBlitImage-dstImage-00225"></a>
  VUID-vkCmdBlitImage-dstImage-00225  
  If `dstImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdBlitImage-dstImageLayout-00226"
  id="VUID-vkCmdBlitImage-dstImageLayout-00226"></a>
  VUID-vkCmdBlitImage-dstImageLayout-00226  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdBlitImage-dstImageLayout-01399"
  id="VUID-vkCmdBlitImage-dstImageLayout-01399"></a>
  VUID-vkCmdBlitImage-dstImageLayout-01399  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdBlitImage-srcImage-00229"
  id="VUID-vkCmdBlitImage-srcImage-00229"></a>
  VUID-vkCmdBlitImage-srcImage-00229  
  If either of `srcImage` or `dstImage` was created with a signed
  integer [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), the other **must** also have been
  created with a signed integer [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)

- <a href="#VUID-vkCmdBlitImage-srcImage-00230"
  id="VUID-vkCmdBlitImage-srcImage-00230"></a>
  VUID-vkCmdBlitImage-srcImage-00230  
  If either of `srcImage` or `dstImage` was created with an unsigned
  integer [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), the other **must** also have been
  created with an unsigned integer [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)

- <a href="#VUID-vkCmdBlitImage-srcImage-00231"
  id="VUID-vkCmdBlitImage-srcImage-00231"></a>
  VUID-vkCmdBlitImage-srcImage-00231  
  If either of `srcImage` or `dstImage` was created with a depth/stencil
  format, the other **must** have exactly the same format

- <a href="#VUID-vkCmdBlitImage-srcImage-00232"
  id="VUID-vkCmdBlitImage-srcImage-00232"></a>
  VUID-vkCmdBlitImage-srcImage-00232  
  If `srcImage` was created with a depth/stencil format, `filter`
  **must** be `VK_FILTER_NEAREST`

- <a href="#VUID-vkCmdBlitImage-srcImage-00233"
  id="VUID-vkCmdBlitImage-srcImage-00233"></a>
  VUID-vkCmdBlitImage-srcImage-00233  
  `srcImage` **must** have been created with a `samples` value of
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-vkCmdBlitImage-dstImage-00234"
  id="VUID-vkCmdBlitImage-dstImage-00234"></a>
  VUID-vkCmdBlitImage-dstImage-00234  
  `dstImage` **must** have been created with a `samples` value of
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-vkCmdBlitImage-filter-02001"
  id="VUID-vkCmdBlitImage-filter-02001"></a>
  VUID-vkCmdBlitImage-filter-02001  
  If `filter` is `VK_FILTER_LINEAR`, then the [format
  features](#resources-image-format-features) of `srcImage` **must**
  contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <a href="#VUID-vkCmdBlitImage-filter-02002"
  id="VUID-vkCmdBlitImage-filter-02002"></a>
  VUID-vkCmdBlitImage-filter-02002  
  If `filter` is `VK_FILTER_CUBIC_EXT`, then the [format
  features](#resources-image-format-features) of `srcImage` **must**
  contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`

- <a href="#VUID-vkCmdBlitImage-filter-00237"
  id="VUID-vkCmdBlitImage-filter-00237"></a>
  VUID-vkCmdBlitImage-filter-00237  
  If `filter` is `VK_FILTER_CUBIC_EXT`, `srcImage` **must** be of type
  `VK_IMAGE_TYPE_2D`

- <a href="#VUID-vkCmdBlitImage-srcSubresource-01705"
  id="VUID-vkCmdBlitImage-srcSubresource-01705"></a>
  VUID-vkCmdBlitImage-srcSubresource-01705  
  The `srcSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-vkCmdBlitImage-dstSubresource-01706"
  id="VUID-vkCmdBlitImage-dstSubresource-01706"></a>
  VUID-vkCmdBlitImage-dstSubresource-01706  
  The `dstSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-vkCmdBlitImage-srcSubresource-01707"
  id="VUID-vkCmdBlitImage-srcSubresource-01707"></a>
  VUID-vkCmdBlitImage-srcSubresource-01707  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-vkCmdBlitImage-dstSubresource-01708"
  id="VUID-vkCmdBlitImage-dstSubresource-01708"></a>
  VUID-vkCmdBlitImage-dstSubresource-01708  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-vkCmdBlitImage-dstImage-02545"
  id="VUID-vkCmdBlitImage-dstImage-02545"></a>
  VUID-vkCmdBlitImage-dstImage-02545  
  `dstImage` and `srcImage` **must** not have been created with `flags`
  containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- <a href="#VUID-vkCmdBlitImage-srcImage-00240"
  id="VUID-vkCmdBlitImage-srcImage-00240"></a>
  VUID-vkCmdBlitImage-srcImage-00240  
  If either `srcImage` or `dstImage` is of type `VK_IMAGE_TYPE_3D`, then
  for each element of `pRegions`, `srcSubresource.baseArrayLayer` and
  `dstSubresource.baseArrayLayer` **must** each be `0`, and
  `srcSubresource.layerCount` and `dstSubresource.layerCount` **must**
  each be `1`

- <a href="#VUID-vkCmdBlitImage-aspectMask-00241"
  id="VUID-vkCmdBlitImage-aspectMask-00241"></a>
  VUID-vkCmdBlitImage-aspectMask-00241  
  For each element of `pRegions`, `srcSubresource.aspectMask` **must**
  specify aspects present in `srcImage`

- <a href="#VUID-vkCmdBlitImage-aspectMask-00242"
  id="VUID-vkCmdBlitImage-aspectMask-00242"></a>
  VUID-vkCmdBlitImage-aspectMask-00242  
  For each element of `pRegions`, `dstSubresource.aspectMask` **must**
  specify aspects present in `dstImage`

- <a href="#VUID-vkCmdBlitImage-srcOffset-00243"
  id="VUID-vkCmdBlitImage-srcOffset-00243"></a>
  VUID-vkCmdBlitImage-srcOffset-00243  
  For each element of `pRegions`, `srcOffsets`\[0\].x and
  `srcOffsets`\[1\].x **must** both be greater than or equal to `0` and
  less than or equal to the width of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-vkCmdBlitImage-srcOffset-00244"
  id="VUID-vkCmdBlitImage-srcOffset-00244"></a>
  VUID-vkCmdBlitImage-srcOffset-00244  
  For each element of `pRegions`, `srcOffsets`\[0\].y and
  `srcOffsets`\[1\].y **must** both be greater than or equal to `0` and
  less than or equal to the height of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-vkCmdBlitImage-srcImage-00245"
  id="VUID-vkCmdBlitImage-srcImage-00245"></a>
  VUID-vkCmdBlitImage-srcImage-00245  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffsets`\[0\].y **must** be `0` and
  `srcOffsets`\[1\].y **must** be `1`

- <a href="#VUID-vkCmdBlitImage-srcOffset-00246"
  id="VUID-vkCmdBlitImage-srcOffset-00246"></a>
  VUID-vkCmdBlitImage-srcOffset-00246  
  For each element of `pRegions`, `srcOffsets`\[0\].z and
  `srcOffsets`\[1\].z **must** both be greater than or equal to `0` and
  less than or equal to the depth of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-vkCmdBlitImage-srcImage-00247"
  id="VUID-vkCmdBlitImage-srcImage-00247"></a>
  VUID-vkCmdBlitImage-srcImage-00247  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `srcOffsets`\[0\].z **must** be
  `0` and `srcOffsets`\[1\].z **must** be `1`

- <a href="#VUID-vkCmdBlitImage-dstOffset-00248"
  id="VUID-vkCmdBlitImage-dstOffset-00248"></a>
  VUID-vkCmdBlitImage-dstOffset-00248  
  For each element of `pRegions`, `dstOffsets`\[0\].x and
  `dstOffsets`\[1\].x **must** both be greater than or equal to `0` and
  less than or equal to the width of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-vkCmdBlitImage-dstOffset-00249"
  id="VUID-vkCmdBlitImage-dstOffset-00249"></a>
  VUID-vkCmdBlitImage-dstOffset-00249  
  For each element of `pRegions`, `dstOffsets`\[0\].y and
  `dstOffsets`\[1\].y **must** both be greater than or equal to `0` and
  less than or equal to the height of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-vkCmdBlitImage-dstImage-00250"
  id="VUID-vkCmdBlitImage-dstImage-00250"></a>
  VUID-vkCmdBlitImage-dstImage-00250  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffsets`\[0\].y **must** be `0` and
  `dstOffsets`\[1\].y **must** be `1`

- <a href="#VUID-vkCmdBlitImage-dstOffset-00251"
  id="VUID-vkCmdBlitImage-dstOffset-00251"></a>
  VUID-vkCmdBlitImage-dstOffset-00251  
  For each element of `pRegions`, `dstOffsets`\[0\].z and
  `dstOffsets`\[1\].z **must** both be greater than or equal to `0` and
  less than or equal to the depth of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-vkCmdBlitImage-dstImage-00252"
  id="VUID-vkCmdBlitImage-dstImage-00252"></a>
  VUID-vkCmdBlitImage-dstImage-00252  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D` or `VK_IMAGE_TYPE_2D`,
  then for each element of `pRegions`, `dstOffsets`\[0\].z **must** be
  `0` and `dstOffsets`\[1\].z **must** be `1`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBlitImage-commandBuffer-parameter"
  id="VUID-vkCmdBlitImage-commandBuffer-parameter"></a>
  VUID-vkCmdBlitImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBlitImage-srcImage-parameter"
  id="VUID-vkCmdBlitImage-srcImage-parameter"></a>
  VUID-vkCmdBlitImage-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdBlitImage-srcImageLayout-parameter"
  id="VUID-vkCmdBlitImage-srcImageLayout-parameter"></a>
  VUID-vkCmdBlitImage-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkCmdBlitImage-dstImage-parameter"
  id="VUID-vkCmdBlitImage-dstImage-parameter"></a>
  VUID-vkCmdBlitImage-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdBlitImage-dstImageLayout-parameter"
  id="VUID-vkCmdBlitImage-dstImageLayout-parameter"></a>
  VUID-vkCmdBlitImage-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkCmdBlitImage-pRegions-parameter"
  id="VUID-vkCmdBlitImage-pRegions-parameter"></a>
  VUID-vkCmdBlitImage-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkImageBlit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit.html) structures

- <a href="#VUID-vkCmdBlitImage-filter-parameter"
  id="VUID-vkCmdBlitImage-filter-parameter"></a>
  VUID-vkCmdBlitImage-filter-parameter  
  `filter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) value

- <a href="#VUID-vkCmdBlitImage-commandBuffer-recording"
  id="VUID-vkCmdBlitImage-commandBuffer-recording"></a>
  VUID-vkCmdBlitImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBlitImage-commandBuffer-cmdpool"
  id="VUID-vkCmdBlitImage-commandBuffer-cmdpool"></a>
  VUID-vkCmdBlitImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBlitImage-renderpass"
  id="VUID-vkCmdBlitImage-renderpass"></a>
  VUID-vkCmdBlitImage-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBlitImage-videocoding"
  id="VUID-vkCmdBlitImage-videocoding"></a>
  VUID-vkCmdBlitImage-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBlitImage-regionCount-arraylength"
  id="VUID-vkCmdBlitImage-regionCount-arraylength"></a>
  VUID-vkCmdBlitImage-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-vkCmdBlitImage-commonparent"
  id="VUID-vkCmdBlitImage-commonparent"></a>
  VUID-vkCmdBlitImage-commonparent  
  Each of `commandBuffer`, `dstImage`, and `srcImage` **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageBlit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBlitImage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

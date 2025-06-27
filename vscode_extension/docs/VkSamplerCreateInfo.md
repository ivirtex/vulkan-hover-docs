# VkSamplerCreateInfo(3) Manual Page

## Name

VkSamplerCreateInfo - Structure specifying parameters of a newly created
sampler



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSamplerCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSamplerCreateInfo {
    VkStructureType         sType;
    const void*             pNext;
    VkSamplerCreateFlags    flags;
    VkFilter                magFilter;
    VkFilter                minFilter;
    VkSamplerMipmapMode     mipmapMode;
    VkSamplerAddressMode    addressModeU;
    VkSamplerAddressMode    addressModeV;
    VkSamplerAddressMode    addressModeW;
    float                   mipLodBias;
    VkBool32                anisotropyEnable;
    float                   maxAnisotropy;
    VkBool32                compareEnable;
    VkCompareOp             compareOp;
    float                   minLod;
    float                   maxLod;
    VkBorderColor           borderColor;
    VkBool32                unnormalizedCoordinates;
} VkSamplerCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlagBits.html) describing
  additional parameters of the sampler.

- `magFilter` is a [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) value specifying the
  magnification filter to apply to lookups.

- `minFilter` is a [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) value specifying the
  minification filter to apply to lookups.

- `mipmapMode` is a [VkSamplerMipmapMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerMipmapMode.html)
  value specifying the mipmap filter to apply to lookups.

- `addressModeU` is a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html)
  value specifying the addressing mode for U coordinates outside \[0,1).

- `addressModeV` is a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html)
  value specifying the addressing mode for V coordinates outside \[0,1).

- `addressModeW` is a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html)
  value specifying the addressing mode for W coordinates outside \[0,1).

- <span id="samplers-mipLodBias"></span> `mipLodBias` is the bias to be
  added to mipmap LOD calculation and bias provided by image sampling
  functions in SPIR-V, as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-level-of-detail-operation"
  target="_blank" rel="noopener">LOD Operation</a> section.

- <span id="samplers-maxAnisotropy"></span> `anisotropyEnable` is
  `VK_TRUE` to enable anisotropic filtering, as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-anisotropic-filtering"
  target="_blank" rel="noopener">Texel Anisotropic Filtering</a>
  section, or `VK_FALSE` otherwise.

- `maxAnisotropy` is the anisotropy value clamp used by the sampler when
  `anisotropyEnable` is `VK_TRUE`. If `anisotropyEnable` is `VK_FALSE`,
  `maxAnisotropy` is ignored.

- `compareEnable` is `VK_TRUE` to enable comparison against a reference
  value during lookups, or `VK_FALSE` otherwise.

  - Note: Some implementations will default to shader state if this
    member does not match.

- `compareOp` is a [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html) value specifying the
  comparison operator to apply to fetched data before filtering as
  described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-depth-compare-operation"
  target="_blank" rel="noopener">Depth Compare Operation</a> section.

- `minLod` is used to clamp the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-level-of-detail-operation"
  target="_blank" rel="noopener">minimum of the computed LOD value</a>.

- `maxLod` is used to clamp the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-level-of-detail-operation"
  target="_blank" rel="noopener">maximum of the computed LOD value</a>.
  To avoid clamping the maximum value, set `maxLod` to the constant
  `VK_LOD_CLAMP_NONE`.

- `borderColor` is a [VkBorderColor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html) value
  specifying the predefined border color to use.

- <span id="samplers-unnormalizedCoordinates"></span>
  `unnormalizedCoordinates` controls whether to use unnormalized or
  normalized texel coordinates to address texels of the image. When set
  to `VK_TRUE`, the range of the image coordinates used to lookup the
  texel is in the range of zero to the image size in each dimension.
  When set to `VK_FALSE` the range of image coordinates is zero to one.

  When `unnormalizedCoordinates` is `VK_TRUE`, images the sampler is
  used with in the shader have the following requirements:

  - The `viewType` **must** be either `VK_IMAGE_VIEW_TYPE_1D` or
    `VK_IMAGE_VIEW_TYPE_2D`.

  - The image view **must** have a single layer and a single mip level.

    When `unnormalizedCoordinates` is `VK_TRUE`, image built-in
    functions in the shader that use the sampler have the following
    requirements:

  - The functions **must** not use projection.

  - The functions **must** not use offsets.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Mapping of OpenGL to Vulkan filter modes
<p><code>magFilter</code> values of <code>VK_FILTER_NEAREST</code> and
<code>VK_FILTER_LINEAR</code> directly correspond to
<code>GL_NEAREST</code> and <code>GL_LINEAR</code> magnification
filters. <code>minFilter</code> and <code>mipmapMode</code> combine to
correspond to the similarly named OpenGL minification filter of
<code>GL_minFilter_MIPMAP_mipmapMode</code> (e.g. <code>minFilter</code>
of <code>VK_FILTER_LINEAR</code> and <code>mipmapMode</code> of
<code>VK_SAMPLER_MIPMAP_MODE_NEAREST</code> correspond to
<code>GL_LINEAR_MIPMAP_NEAREST</code>).</p>
<p>There are no Vulkan filter modes that directly correspond to OpenGL
minification filters of <code>GL_LINEAR</code> or
<code>GL_NEAREST</code>, but they <strong>can</strong> be emulated using
<code>VK_SAMPLER_MIPMAP_MODE_NEAREST</code>, <code>minLod</code> = 0,
and <code>maxLod</code> = 0.25, and using <code>minFilter</code> =
<code>VK_FILTER_LINEAR</code> or <code>minFilter</code> =
<code>VK_FILTER_NEAREST</code>, respectively.</p>
<p>Note that using a <code>maxLod</code> of zero would cause <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-filtering"
target="_blank" rel="noopener">magnification</a> to always be performed,
and the <code>magFilter</code> to always be used. This is valid, just
not an exact match for OpenGL behavior. Clamping the maximum LOD to 0.25
allows the λ value to be non-zero and minification to be performed,
while still always rounding down to the base level. If the
<code>minFilter</code> and <code>magFilter</code> are equal, then using
a <code>maxLod</code> of zero also works.</p></td>
</tr>
</tbody>
</table>

The maximum number of sampler objects which **can** be simultaneously
created on a device is implementation-dependent and specified by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxSamplerAllocationCount"
target="_blank"
rel="noopener"><code>maxSamplerAllocationCount</code></a> member of the
[VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html) structure.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For historical reasons, if <code>maxSamplerAllocationCount</code> is
exceeded, some implementations may return
<code>VK_ERROR_TOO_MANY_OBJECTS</code>. Exceeding this limit will result
in undefined behavior, and an application should not rely on the use of
the returned error code in order to identify when the limit is
reached.</p></td>
</tr>
</tbody>
</table>

Since [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) is a non-dispatchable handle type,
implementations **may** return the same handle for sampler state vectors
that are identical. In such cases, all such objects would only count
once against the `maxSamplerAllocationCount` limit.

Valid Usage

- <a href="#VUID-VkSamplerCreateInfo-mipLodBias-01069"
  id="VUID-VkSamplerCreateInfo-mipLodBias-01069"></a>
  VUID-VkSamplerCreateInfo-mipLodBias-01069  
  The absolute value of `mipLodBias` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxSamplerLodBias`

- <a href="#VUID-VkSamplerCreateInfo-samplerMipLodBias-04467"
  id="VUID-VkSamplerCreateInfo-samplerMipLodBias-04467"></a>
  VUID-VkSamplerCreateInfo-samplerMipLodBias-04467  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`samplerMipLodBias`
  is `VK_FALSE`, `mipLodBias` **must** be zero

- <a href="#VUID-VkSamplerCreateInfo-maxLod-01973"
  id="VUID-VkSamplerCreateInfo-maxLod-01973"></a>
  VUID-VkSamplerCreateInfo-maxLod-01973  
  `maxLod` **must** be greater than or equal to `minLod`

- <a href="#VUID-VkSamplerCreateInfo-anisotropyEnable-01070"
  id="VUID-VkSamplerCreateInfo-anisotropyEnable-01070"></a>
  VUID-VkSamplerCreateInfo-anisotropyEnable-01070  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-samplerAnisotropy"
  target="_blank" rel="noopener"><code>samplerAnisotropy</code></a>
  feature is not enabled, `anisotropyEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-anisotropyEnable-01071"
  id="VUID-VkSamplerCreateInfo-anisotropyEnable-01071"></a>
  VUID-VkSamplerCreateInfo-anisotropyEnable-01071  
  If `anisotropyEnable` is `VK_TRUE`, `maxAnisotropy` **must** be
  between `1.0` and `VkPhysicalDeviceLimits`::`maxSamplerAnisotropy`,
  inclusive

- <a href="#VUID-VkSamplerCreateInfo-minFilter-01645"
  id="VUID-VkSamplerCreateInfo-minFilter-01645"></a>
  VUID-VkSamplerCreateInfo-minFilter-01645  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> is enabled and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> of the
  sampler Y′C<sub>B</sub>C<sub>R</sub> conversion do not support
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT`,
  `minFilter` and `magFilter` **must** be equal to the sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion’s `chromaFilter`

- <a href="#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01072"
  id="VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01072"></a>
  VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01072  
  If `unnormalizedCoordinates` is `VK_TRUE`, `minFilter` and `magFilter`
  **must** be equal

- <a href="#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01073"
  id="VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01073"></a>
  VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01073  
  If `unnormalizedCoordinates` is `VK_TRUE`, `mipmapMode` **must** be
  `VK_SAMPLER_MIPMAP_MODE_NEAREST`

- <a href="#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01074"
  id="VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01074"></a>
  VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01074  
  If `unnormalizedCoordinates` is `VK_TRUE`, `minLod` and `maxLod`
  **must** be zero

- <a href="#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01075"
  id="VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01075"></a>
  VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01075  
  If `unnormalizedCoordinates` is `VK_TRUE`, `addressModeU` and
  `addressModeV` **must** each be either
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE` or
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`

- <a href="#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01076"
  id="VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01076"></a>
  VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01076  
  If `unnormalizedCoordinates` is `VK_TRUE`, `anisotropyEnable` **must**
  be `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01077"
  id="VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01077"></a>
  VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01077  
  If `unnormalizedCoordinates` is `VK_TRUE`, `compareEnable` **must** be
  `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-addressModeU-01078"
  id="VUID-VkSamplerCreateInfo-addressModeU-01078"></a>
  VUID-VkSamplerCreateInfo-addressModeU-01078  
  If any of `addressModeU`, `addressModeV` or `addressModeW` are
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`, `borderColor` **must** be a
  valid [VkBorderColor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html) value

- <a href="#VUID-VkSamplerCreateInfo-addressModeU-01646"
  id="VUID-VkSamplerCreateInfo-addressModeU-01646"></a>
  VUID-VkSamplerCreateInfo-addressModeU-01646  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> is enabled, `addressModeU`, `addressModeV`, and
  `addressModeW` **must** be `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`,
  `anisotropyEnable` **must** be `VK_FALSE`, and
  `unnormalizedCoordinates` **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-None-01647"
  id="VUID-VkSamplerCreateInfo-None-01647"></a>
  VUID-VkSamplerCreateInfo-None-01647  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> is enabled and the `pNext` chain includes a
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)
  structure, then the sampler reduction mode **must** be set to
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`

- <a href="#VUID-VkSamplerCreateInfo-pNext-06726"
  id="VUID-VkSamplerCreateInfo-pNext-06726"></a>
  VUID-VkSamplerCreateInfo-pNext-06726  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-samplerFilterMinmax"
  target="_blank" rel="noopener"><code>samplerFilterMinmax</code></a> is
  not enabled and the `pNext` chain includes a
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)
  structure, then the sampler reduction mode **must** be set to
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`

- <a href="#VUID-VkSamplerCreateInfo-addressModeU-01079"
  id="VUID-VkSamplerCreateInfo-addressModeU-01079"></a>
  VUID-VkSamplerCreateInfo-addressModeU-01079  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-samplerMirrorClampToEdge"
  target="_blank" rel="noopener"><code>samplerMirrorClampToEdge</code></a>
  is not enabled, and if the
  [`VK_KHR_sampler_mirror_clamp_to_edge`](VK_KHR_sampler_mirror_clamp_to_edge.html)
  extension is not enabled, `addressModeU`, `addressModeV` and
  `addressModeW` **must** not be
  `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE`

- <a href="#VUID-VkSamplerCreateInfo-compareEnable-01080"
  id="VUID-VkSamplerCreateInfo-compareEnable-01080"></a>
  VUID-VkSamplerCreateInfo-compareEnable-01080  
  If `compareEnable` is `VK_TRUE`, `compareOp` **must** be a valid
  [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html) value

- <a href="#VUID-VkSamplerCreateInfo-magFilter-01081"
  id="VUID-VkSamplerCreateInfo-magFilter-01081"></a>
  VUID-VkSamplerCreateInfo-magFilter-01081  
  If either `magFilter` or `minFilter` is `VK_FILTER_CUBIC_EXT`,
  `anisotropyEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-magFilter-07911"
  id="VUID-VkSamplerCreateInfo-magFilter-07911"></a>
  VUID-VkSamplerCreateInfo-magFilter-07911  
  If the [VK_EXT_filter_cubic](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_filter_cubic.html) extension is
  not enabled and either `magFilter` or `minFilter` is
  `VK_FILTER_CUBIC_IMG`, the `reductionMode` member of
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)
  **must** be `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`

- <a href="#VUID-VkSamplerCreateInfo-compareEnable-01423"
  id="VUID-VkSamplerCreateInfo-compareEnable-01423"></a>
  VUID-VkSamplerCreateInfo-compareEnable-01423  
  If `compareEnable` is `VK_TRUE`, the `reductionMode` member of
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)
  **must** be `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`

- <a href="#VUID-VkSamplerCreateInfo-flags-02574"
  id="VUID-VkSamplerCreateInfo-flags-02574"></a>
  VUID-VkSamplerCreateInfo-flags-02574  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then
  `minFilter` and `magFilter` **must** be equal

- <a href="#VUID-VkSamplerCreateInfo-flags-02575"
  id="VUID-VkSamplerCreateInfo-flags-02575"></a>
  VUID-VkSamplerCreateInfo-flags-02575  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then
  `mipmapMode` **must** be `VK_SAMPLER_MIPMAP_MODE_NEAREST`

- <a href="#VUID-VkSamplerCreateInfo-flags-02576"
  id="VUID-VkSamplerCreateInfo-flags-02576"></a>
  VUID-VkSamplerCreateInfo-flags-02576  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then
  `minLod` and `maxLod` **must** be zero

- <a href="#VUID-VkSamplerCreateInfo-flags-02577"
  id="VUID-VkSamplerCreateInfo-flags-02577"></a>
  VUID-VkSamplerCreateInfo-flags-02577  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then
  `addressModeU` and `addressModeV` **must** each be either
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE` or
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`

- <a href="#VUID-VkSamplerCreateInfo-flags-02578"
  id="VUID-VkSamplerCreateInfo-flags-02578"></a>
  VUID-VkSamplerCreateInfo-flags-02578  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then
  `anisotropyEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-flags-02579"
  id="VUID-VkSamplerCreateInfo-flags-02579"></a>
  VUID-VkSamplerCreateInfo-flags-02579  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then
  `compareEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-flags-02580"
  id="VUID-VkSamplerCreateInfo-flags-02580"></a>
  VUID-VkSamplerCreateInfo-flags-02580  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then
  `unnormalizedCoordinates` **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-nonSeamlessCubeMap-06788"
  id="VUID-VkSamplerCreateInfo-nonSeamlessCubeMap-06788"></a>
  VUID-VkSamplerCreateInfo-nonSeamlessCubeMap-06788  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nonSeamlessCubeMap"
  target="_blank" rel="noopener"><code>nonSeamlessCubeMap</code></a>
  feature is not enabled, `flags` **must** not include
  `VK_SAMPLER_CREATE_NON_SEAMLESS_CUBE_MAP_BIT_EXT`

- <a href="#VUID-VkSamplerCreateInfo-borderColor-04011"
  id="VUID-VkSamplerCreateInfo-borderColor-04011"></a>
  VUID-VkSamplerCreateInfo-borderColor-04011  
  If `borderColor` is one of `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` or
  `VK_BORDER_COLOR_INT_CUSTOM_EXT`, then a
  [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)
  **must** be included in the `pNext` chain

- <a href="#VUID-VkSamplerCreateInfo-customBorderColors-04085"
  id="VUID-VkSamplerCreateInfo-customBorderColors-04085"></a>
  VUID-VkSamplerCreateInfo-customBorderColors-04085  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-customBorderColors"
  target="_blank" rel="noopener"><code>customBorderColors</code></a>
  feature is not enabled, `borderColor` **must** not be
  `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` or `VK_BORDER_COLOR_INT_CUSTOM_EXT`

- <a href="#VUID-VkSamplerCreateInfo-borderColor-04442"
  id="VUID-VkSamplerCreateInfo-borderColor-04442"></a>
  VUID-VkSamplerCreateInfo-borderColor-04442  
  If `borderColor` is one of `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` or
  `VK_BORDER_COLOR_INT_CUSTOM_EXT`, and
  [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)::`format`
  is not `VK_FORMAT_UNDEFINED`,
  [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)::`customBorderColor`
  **must** be within the range of values representable in `format`

- <a href="#VUID-VkSamplerCreateInfo-None-04012"
  id="VUID-VkSamplerCreateInfo-None-04012"></a>
  VUID-VkSamplerCreateInfo-None-04012  
  The maximum number of samplers with custom border colors which **can**
  be simultaneously created on a device is implementation-dependent and
  specified by the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxCustomBorderColorSamplers"
  target="_blank"
  rel="noopener"><code>maxCustomBorderColorSamplers</code></a> member of
  the
  [VkPhysicalDeviceCustomBorderColorPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCustomBorderColorPropertiesEXT.html)
  structure

- <a href="#VUID-VkSamplerCreateInfo-flags-08110"
  id="VUID-VkSamplerCreateInfo-flags-08110"></a>
  VUID-VkSamplerCreateInfo-flags-08110  
  If `flags` includes
  `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBufferCaptureReplay"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a href="#VUID-VkSamplerCreateInfo-pNext-08111"
  id="VUID-VkSamplerCreateInfo-pNext-08111"></a>
  VUID-VkSamplerCreateInfo-pNext-08111  
  If the `pNext` chain includes a
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  structure, `flags` **must** contain
  `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`

- <a href="#VUID-VkSamplerCreateInfo-flags-06964"
  id="VUID-VkSamplerCreateInfo-flags-06964"></a>
  VUID-VkSamplerCreateInfo-flags-06964  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`,
  then `minFilter` and `magFilter` **must** be `VK_FILTER_NEAREST`

- <a href="#VUID-VkSamplerCreateInfo-flags-06965"
  id="VUID-VkSamplerCreateInfo-flags-06965"></a>
  VUID-VkSamplerCreateInfo-flags-06965  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`,
  then `mipmapMode` **must** be `VK_SAMPLER_MIPMAP_MODE_NEAREST`

- <a href="#VUID-VkSamplerCreateInfo-flags-06966"
  id="VUID-VkSamplerCreateInfo-flags-06966"></a>
  VUID-VkSamplerCreateInfo-flags-06966  
  \[If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`,
  then `minLod` and `maxLod` **must** be zero

- <a href="#VUID-VkSamplerCreateInfo-flags-06967"
  id="VUID-VkSamplerCreateInfo-flags-06967"></a>
  VUID-VkSamplerCreateInfo-flags-06967  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`,
  then `addressModeU` and `addressModeV` **must** each be either
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE` or
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`

- <a href="#VUID-VkSamplerCreateInfo-flags-06968"
  id="VUID-VkSamplerCreateInfo-flags-06968"></a>
  VUID-VkSamplerCreateInfo-flags-06968  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`, and
  if `addressModeU` or `addressModeV` is
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`, then `borderColor` **must**
  be `VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK`

- <a href="#VUID-VkSamplerCreateInfo-flags-06969"
  id="VUID-VkSamplerCreateInfo-flags-06969"></a>
  VUID-VkSamplerCreateInfo-flags-06969  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`,
  then `anisotropyEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerCreateInfo-flags-06970"
  id="VUID-VkSamplerCreateInfo-flags-06970"></a>
  VUID-VkSamplerCreateInfo-flags-06970  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`,
  then `compareEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- <a href="#VUID-VkSamplerCreateInfo-sType-sType"
  id="VUID-VkSamplerCreateInfo-sType-sType"></a>
  VUID-VkSamplerCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO`

- <a href="#VUID-VkSamplerCreateInfo-pNext-pNext"
  id="VUID-VkSamplerCreateInfo-pNext-pNext"></a>
  VUID-VkSamplerCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html),
  [VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html),
  [VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html),
  [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html),
  [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html),
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html),
  or [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)

- <a href="#VUID-VkSamplerCreateInfo-sType-unique"
  id="VUID-VkSamplerCreateInfo-sType-unique"></a>
  VUID-VkSamplerCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkSamplerCreateInfo-flags-parameter"
  id="VUID-VkSamplerCreateInfo-flags-parameter"></a>
  VUID-VkSamplerCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlagBits.html) values

- <a href="#VUID-VkSamplerCreateInfo-magFilter-parameter"
  id="VUID-VkSamplerCreateInfo-magFilter-parameter"></a>
  VUID-VkSamplerCreateInfo-magFilter-parameter  
  `magFilter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) value

- <a href="#VUID-VkSamplerCreateInfo-minFilter-parameter"
  id="VUID-VkSamplerCreateInfo-minFilter-parameter"></a>
  VUID-VkSamplerCreateInfo-minFilter-parameter  
  `minFilter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) value

- <a href="#VUID-VkSamplerCreateInfo-mipmapMode-parameter"
  id="VUID-VkSamplerCreateInfo-mipmapMode-parameter"></a>
  VUID-VkSamplerCreateInfo-mipmapMode-parameter  
  `mipmapMode` **must** be a valid
  [VkSamplerMipmapMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerMipmapMode.html) value

- <a href="#VUID-VkSamplerCreateInfo-addressModeU-parameter"
  id="VUID-VkSamplerCreateInfo-addressModeU-parameter"></a>
  VUID-VkSamplerCreateInfo-addressModeU-parameter  
  `addressModeU` **must** be a valid
  [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html) value

- <a href="#VUID-VkSamplerCreateInfo-addressModeV-parameter"
  id="VUID-VkSamplerCreateInfo-addressModeV-parameter"></a>
  VUID-VkSamplerCreateInfo-addressModeV-parameter  
  `addressModeV` **must** be a valid
  [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html) value

- <a href="#VUID-VkSamplerCreateInfo-addressModeW-parameter"
  id="VUID-VkSamplerCreateInfo-addressModeW-parameter"></a>
  VUID-VkSamplerCreateInfo-addressModeW-parameter  
  `addressModeW` **must** be a valid
  [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkBorderColor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html), [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html),
[VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html),
[VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html),
[VkSamplerCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlags.html),
[VkSamplerMipmapMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerMipmapMode.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSampler.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

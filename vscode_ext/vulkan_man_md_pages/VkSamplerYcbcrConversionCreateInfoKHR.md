# VkSamplerYcbcrConversionCreateInfo(3) Manual Page

## Name

VkSamplerYcbcrConversionCreateInfo - Structure specifying the parameters
of the newly created conversion



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSamplerYcbcrConversionCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkSamplerYcbcrConversionCreateInfo {
    VkStructureType                  sType;
    const void*                      pNext;
    VkFormat                         format;
    VkSamplerYcbcrModelConversion    ycbcrModel;
    VkSamplerYcbcrRange              ycbcrRange;
    VkComponentMapping               components;
    VkChromaLocation                 xChromaOffset;
    VkChromaLocation                 yChromaOffset;
    VkFilter                         chromaFilter;
    VkBool32                         forceExplicitReconstruction;
} VkSamplerYcbcrConversionCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrConversionCreateInfo VkSamplerYcbcrConversionCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `format` is the format of the image from which color information will
  be retrieved.

- `ycbcrModel` describes the color matrix for conversion between color
  models.

- `ycbcrRange` describes whether the encoded values have headroom and
  foot room, or whether the encoding uses the full numerical range.

- `components` applies a *swizzle* based on
  [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html) enums prior to range
  expansion and color model conversion.

- `xChromaOffset` describes the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-chroma-reconstruction"
  target="_blank" rel="noopener">sample location</a> associated with
  downsampled chroma components in the x dimension. `xChromaOffset` has
  no effect for formats in which chroma components are not downsampled
  horizontally.

- `yChromaOffset` describes the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-chroma-reconstruction"
  target="_blank" rel="noopener">sample location</a> associated with
  downsampled chroma components in the y dimension. `yChromaOffset` has
  no effect for formats in which the chroma components are not
  downsampled vertically.

- `chromaFilter` is the filter for chroma reconstruction.

- `forceExplicitReconstruction` **can** be used to ensure that
  reconstruction is done explicitly, if supported.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Setting <code>forceExplicitReconstruction</code> to
<code>VK_TRUE</code> <strong>may</strong> have a performance penalty on
implementations where explicit reconstruction is not the default mode of
operation.</p>
<p>If <code>format</code> supports
<code>VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT</code>
the <code>forceExplicitReconstruction</code> value behaves as if it was
set to <code>VK_TRUE</code>.</p></td>
</tr>
</tbody>
</table>

If the `pNext` chain includes a
[VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure with
non-zero `externalFormat` member, the sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion object represents an *external
format conversion*, and `format` **must** be `VK_FORMAT_UNDEFINED`. Such
conversions **must** only be used to sample image views with a matching
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-external-formats"
target="_blank" rel="noopener">external format</a>. When creating an
external format conversion, the value of `components` is ignored.

Valid Usage

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-format-01904"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-format-01904"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-format-01904  
  If an external format conversion is being created, `format` **must**
  be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-format-04061"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-format-04061"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-format-04061  
  If an external format conversion is not being created, `format`
  **must** represent unsigned normalized values (i.e. the format
  **must** be a `UNORM` format)

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-format-01650"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-format-01650"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-format-01650  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> of the
  sampler Y′C<sub>B</sub>C<sub>R</sub> conversion **must** support
  `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT` or
  `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01651"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01651"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01651  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> of the
  sampler Y′C<sub>B</sub>C<sub>R</sub> conversion do not support
  `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT`, `xChromaOffset` and
  `yChromaOffset` **must** not be `VK_CHROMA_LOCATION_COSITED_EVEN` if
  the corresponding components are <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-chroma-reconstruction"
  target="_blank" rel="noopener">downsampled</a>

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01652"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01652"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01652  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> of the
  sampler Y′C<sub>B</sub>C<sub>R</sub> conversion do not support
  `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT`, `xChromaOffset` and
  `yChromaOffset` **must** not be `VK_CHROMA_LOCATION_MIDPOINT` if the
  corresponding components are <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-chroma-reconstruction"
  target="_blank" rel="noopener">downsampled</a>

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-components-02581"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-components-02581"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-components-02581  
  If the format has a `_422` or `_420` suffix, then `components.g`
  **must** be the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-components-02582"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-components-02582"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-components-02582  
  If the format has a `_422` or `_420` suffix, then `components.a`
  **must** be the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>,
  `VK_COMPONENT_SWIZZLE_ONE`, or `VK_COMPONENT_SWIZZLE_ZERO`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-components-02583"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-components-02583"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-components-02583  
  If the format has a `_422` or `_420` suffix, then `components.r`
  **must** be the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a> or
  `VK_COMPONENT_SWIZZLE_B`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-components-02584"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-components-02584"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-components-02584  
  If the format has a `_422` or `_420` suffix, then `components.b`
  **must** be the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a> or
  `VK_COMPONENT_SWIZZLE_R`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-components-02585"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-components-02585"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-components-02585  
  If the format has a `_422` or `_420` suffix, and if either
  `components.r` or `components.b` is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>, both values
  **must** be the identity swizzle

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-01655"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-01655"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-01655  
  If `ycbcrModel` is not
  `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY`, then `components.r`,
  `components.g`, and `components.b` **must** correspond to components
  of the `format`; that is, `components.r`, `components.g`, and
  `components.b` **must** not be `VK_COMPONENT_SWIZZLE_ZERO` or
  `VK_COMPONENT_SWIZZLE_ONE`, and **must** not correspond to a component
  containing zero or one as a consequence of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-conversion-to-rgba"
  target="_blank" rel="noopener">conversion to RGBA</a>

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-02748"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-02748"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-02748  
  If `ycbcrRange` is `VK_SAMPLER_YCBCR_RANGE_ITU_NARROW` then the R, G
  and B components obtained by applying the `component` swizzle to
  `format` **must** each have a bit-depth greater than or equal to 8

- <a
  href="#VUID-VkSamplerYcbcrConversionCreateInfo-forceExplicitReconstruction-01656"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-forceExplicitReconstruction-01656"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-forceExplicitReconstruction-01656  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> of the
  sampler Y′C<sub>B</sub>C<sub>R</sub> conversion do not support
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT`
  `forceExplicitReconstruction` **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-01657"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-01657"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-01657  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> of the
  sampler Y′C<sub>B</sub>C<sub>R</sub> conversion do not support
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT`,
  `chromaFilter` **must** not be `VK_FILTER_LINEAR`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09207"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09207"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09207  
  If the `pNext` chain includes a
  [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)
  structure, and if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-ycbcr-degamma"
  target="_blank" rel="noopener"><code>ycbcrDegamma</code></a> feature
  is not enabled, then
  [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)::`enableYDegamma`
  **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09208"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09208"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09208  
  If the `pNext` chain includes a
  [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)
  structure, and if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-ycbcr-degamma"
  target="_blank" rel="noopener"><code>ycbcrDegamma</code></a> feature
  is not enabled, then
  [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)::`enableCbCrDegamma`
  **must** be `VK_FALSE`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09209"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09209"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09209  
  If the `pNext` chain includes a
  [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)
  structure, `format` **must** be a format with 8-bit R, G, and B
  components

Valid Usage (Implicit)

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-sType-sType"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-sType-sType"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO`

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-pNext-pNext"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-pNext-pNext"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html),
  [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html), or
  [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-sType-unique"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-sType-unique"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-format-parameter"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-format-parameter"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-parameter"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-parameter"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-parameter  
  `ycbcrModel` **must** be a valid
  [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html)
  value

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-parameter"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-parameter"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-parameter  
  `ycbcrRange` **must** be a valid
  [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html) value

- <a href="#VUID-VkSamplerYcbcrConversionCreateInfo-components-parameter"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-components-parameter"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-components-parameter  
  `components` **must** be a valid
  [VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html) structure

- <a
  href="#VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-parameter"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-parameter"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-parameter  
  `xChromaOffset` **must** be a valid
  [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html) value

- <a
  href="#VUID-VkSamplerYcbcrConversionCreateInfo-yChromaOffset-parameter"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-yChromaOffset-parameter"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-yChromaOffset-parameter  
  `yChromaOffset` **must** be a valid
  [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html) value

- <a
  href="#VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-parameter"
  id="VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-parameter"></a>
  VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-parameter  
  `chromaFilter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html) value

If `chromaFilter` is `VK_FILTER_NEAREST`, chroma samples are
reconstructed to luma component resolution using nearest-neighbour
sampling. Otherwise, chroma samples are reconstructed using
interpolation. More details can be found in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-sampler-YCbCr-conversion"
target="_blank" rel="noopener">the description of sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion</a> in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures"
target="_blank" rel="noopener">Image Operations</a> chapter.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html),
[VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html),
[VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html),
[VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversion.html),
[vkCreateSamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversionKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerYcbcrConversionCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

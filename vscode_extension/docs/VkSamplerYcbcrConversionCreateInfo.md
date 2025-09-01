# VkSamplerYcbcrConversionCreateInfo(3) Manual Page

## Name

VkSamplerYcbcrConversionCreateInfo - Structure specifying the parameters of the newly created conversion



## [](#_c_specification)C Specification

The `VkSamplerYcbcrConversionCreateInfo` structure is defined as:

```c++
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

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrConversionCreateInfo VkSamplerYcbcrConversionCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `format` is the format of the image from which color information will be retrieved.
- `ycbcrModel` describes the color matrix for conversion between color models.
- `ycbcrRange` describes whether the encoded values have headroom and foot room, or whether the encoding uses the full numerical range.
- `components` applies a *swizzle* based on [VkComponentSwizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentSwizzle.html) enums prior to range expansion and color model conversion.
- `xChromaOffset` describes the [sample location](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-chroma-reconstruction) associated with downsampled chroma components in the x dimension. `xChromaOffset` has no effect for formats in which chroma components are not downsampled horizontally.
- `yChromaOffset` describes the [sample location](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-chroma-reconstruction) associated with downsampled chroma components in the y dimension. `yChromaOffset` has no effect for formats in which the chroma components are not downsampled vertically.
- `chromaFilter` is the filter for chroma reconstruction.
- `forceExplicitReconstruction` **can** be used to ensure that reconstruction is done explicitly, if supported.

## [](#_description)Description

Note

Setting `forceExplicitReconstruction` to `VK_TRUE` **may** have a performance penalty on implementations where explicit reconstruction is not the default mode of operation.

If `format` supports `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT` the `forceExplicitReconstruction` value behaves as if it were `VK_TRUE`.

If the `pNext` chain includes a [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html) structure with non-zero `externalFormat` member, the sampler Y′CBCR conversion object represents an *external format conversion*, and `format` **must** be `VK_FORMAT_UNDEFINED`. Such conversions **must** only be used to sample image views with a matching [external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-android-hardware-buffer-external-formats). When creating an external format conversion, the value of `components` is ignored.

Valid Usage

- [](#VUID-VkSamplerYcbcrConversionCreateInfo-format-01904)VUID-VkSamplerYcbcrConversionCreateInfo-format-01904  
  If an external format conversion is being created, `format` **must** be `VK_FORMAT_UNDEFINED`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-format-04061)VUID-VkSamplerYcbcrConversionCreateInfo-format-04061  
  If an external format conversion is not being created, `format` **must** represent unsigned normalized values (i.e. the format **must** be a `UNORM` format)
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-format-01650)VUID-VkSamplerYcbcrConversionCreateInfo-format-01650  
  The [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) of the sampler Y′CBCR conversion **must** support `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT` or `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01651)VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01651  
  If the [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) of the sampler Y′CBCR conversion do not support `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT`, `xChromaOffset` and `yChromaOffset` **must** not be `VK_CHROMA_LOCATION_COSITED_EVEN` if the corresponding components are [downsampled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-chroma-reconstruction)
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01652)VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-01652  
  If the [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) of the sampler Y′CBCR conversion do not support `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT`, `xChromaOffset` and `yChromaOffset` **must** not be `VK_CHROMA_LOCATION_MIDPOINT` if the corresponding components are [downsampled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-chroma-reconstruction)
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-components-02581)VUID-VkSamplerYcbcrConversionCreateInfo-components-02581  
  If the format has a `_422` or `_420` suffix, then `components.g` **must** be the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings)
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-components-02582)VUID-VkSamplerYcbcrConversionCreateInfo-components-02582  
  If the format has a `_422` or `_420` suffix, then `components.a` **must** be the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings), `VK_COMPONENT_SWIZZLE_ONE`, or `VK_COMPONENT_SWIZZLE_ZERO`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-components-02583)VUID-VkSamplerYcbcrConversionCreateInfo-components-02583  
  If the format has a `_422` or `_420` suffix, then `components.r` **must** be the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings) or `VK_COMPONENT_SWIZZLE_B`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-components-02584)VUID-VkSamplerYcbcrConversionCreateInfo-components-02584  
  If the format has a `_422` or `_420` suffix, then `components.b` **must** be the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings) or `VK_COMPONENT_SWIZZLE_R`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-components-02585)VUID-VkSamplerYcbcrConversionCreateInfo-components-02585  
  If the format has a `_422` or `_420` suffix, and if either `components.r` or `components.b` is the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings), both values **must** be the identity swizzle
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-01655)VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-01655  
  If `ycbcrModel` is not `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY`, then `components.r`, `components.g`, and `components.b` **must** correspond to components of the `format`; that is, `components.r`, `components.g`, and `components.b` **must** not be `VK_COMPONENT_SWIZZLE_ZERO` or `VK_COMPONENT_SWIZZLE_ONE`, and **must** not correspond to a component containing zero or one as a consequence of [conversion to RGBA](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-conversion-to-rgba)
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-02748)VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-02748  
  If `ycbcrRange` is `VK_SAMPLER_YCBCR_RANGE_ITU_NARROW` then the R, G and B components obtained by applying the `component` swizzle to `format` **must** each have a bit-depth greater than or equal to 8
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-forceExplicitReconstruction-01656)VUID-VkSamplerYcbcrConversionCreateInfo-forceExplicitReconstruction-01656  
  If the [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) of the sampler Y′CBCR conversion do not support `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT` `forceExplicitReconstruction` **must** be `VK_FALSE`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-01657)VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-01657  
  If the [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) of the sampler Y′CBCR conversion do not support `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT`, `chromaFilter` **must** not be `VK_FILTER_LINEAR`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09207)VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09207  
  If the `pNext` chain includes a [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html) structure, and if the [`ycbcrDegamma`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-ycbcrDegamma) feature is not enabled, then [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)::`enableYDegamma` **must** be `VK_FALSE`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09208)VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09208  
  If the `pNext` chain includes a [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html) structure, and if the [`ycbcrDegamma`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-ycbcrDegamma) feature is not enabled, then [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)::`enableCbCrDegamma` **must** be `VK_FALSE`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09209)VUID-VkSamplerYcbcrConversionCreateInfo-pNext-09209  
  If the `pNext` chain includes a [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html) structure, `format` **must** be a format with 8-bit R, G, and B components

Valid Usage (Implicit)

- [](#VUID-VkSamplerYcbcrConversionCreateInfo-sType-sType)VUID-VkSamplerYcbcrConversionCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO`
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-pNext-pNext)VUID-VkSamplerYcbcrConversionCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html), [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatQNX.html), or [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-sType-unique)VUID-VkSamplerYcbcrConversionCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-format-parameter)VUID-VkSamplerYcbcrConversionCreateInfo-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-parameter)VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrModel-parameter  
  `ycbcrModel` **must** be a valid [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversion.html) value
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-parameter)VUID-VkSamplerYcbcrConversionCreateInfo-ycbcrRange-parameter  
  `ycbcrRange` **must** be a valid [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html) value
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-components-parameter)VUID-VkSamplerYcbcrConversionCreateInfo-components-parameter  
  `components` **must** be a valid [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html) structure
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-parameter)VUID-VkSamplerYcbcrConversionCreateInfo-xChromaOffset-parameter  
  `xChromaOffset` **must** be a valid [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html) value
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-yChromaOffset-parameter)VUID-VkSamplerYcbcrConversionCreateInfo-yChromaOffset-parameter  
  `yChromaOffset` **must** be a valid [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html) value
- [](#VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-parameter)VUID-VkSamplerYcbcrConversionCreateInfo-chromaFilter-parameter  
  `chromaFilter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html) value

If `chromaFilter` is `VK_FILTER_NEAREST`, chroma samples are reconstructed to luma component resolution using nearest-neighbour sampling. Otherwise, chroma samples are reconstructed using interpolation. More details can be found in [the description of sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-sampler-YCbCr-conversion) in the [Image Operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures) chapter.

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html), [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html), [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversion.html), [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSamplerYcbcrConversion.html), [vkCreateSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSamplerYcbcrConversion.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerYcbcrConversionCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
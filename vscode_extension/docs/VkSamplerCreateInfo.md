# VkSamplerCreateInfo(3) Manual Page

## Name

VkSamplerCreateInfo - Structure specifying parameters of a newly created sampler



## [](#_c_specification)C Specification

The `VkSamplerCreateInfo` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateFlagBits.html) describing additional parameters of the sampler.
- `magFilter` is a [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html) value specifying the magnification filter to apply to lookups.
- `minFilter` is a [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html) value specifying the minification filter to apply to lookups.
- `mipmapMode` is a [VkSamplerMipmapMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerMipmapMode.html) value specifying the mipmap filter to apply to lookups.
- `addressModeU` is a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) value specifying the addressing mode for U coordinates outside [0,1).
- `addressModeV` is a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) value specifying the addressing mode for V coordinates outside [0,1).
- `addressModeW` is a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) value specifying the addressing mode for W coordinates outside [0,1).
- []()`mipLodBias` is the bias to be added to mipmap LOD calculation and bias provided by image sampling functions in SPIR-V, as described in the [LOD Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-level-of-detail-operation) section.
- []()`anisotropyEnable` is `VK_TRUE` to enable anisotropic filtering, as described in the [Texel Anisotropic Filtering](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-texel-anisotropic-filtering) section, or `VK_FALSE` otherwise.
- `maxAnisotropy` is the anisotropy value clamp used by the sampler when `anisotropyEnable` is `VK_TRUE`. If `anisotropyEnable` is `VK_FALSE`, `maxAnisotropy` is ignored.
- `compareEnable` is `VK_TRUE` to enable comparison against a reference value during lookups, or `VK_FALSE` otherwise.
  
  - Note: Some implementations will default to shader state if this member does not match.
- `compareOp` is a [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html) value specifying the comparison operator to apply to fetched data before filtering as described in the [Depth Compare Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-depth-compare-operation) section.
- `minLod` is used to clamp the [minimum of the computed LOD value](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-level-of-detail-operation).
- `maxLod` is used to clamp the [maximum of the computed LOD value](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-level-of-detail-operation). To avoid clamping the maximum value, set `maxLod` to the constant `VK_LOD_CLAMP_NONE`.
- `borderColor` is a [VkBorderColor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBorderColor.html) value specifying the predefined border color to use.
- []()`unnormalizedCoordinates` controls whether to use unnormalized or normalized texel coordinates to address texels of the image. When `unnormalizedCoordinates` is `VK_TRUE`, the range of the image coordinates used to lookup the texel is in the range of zero to the image size in each dimension. When `unnormalizedCoordinates` is `VK_FALSE`, the range of image coordinates is zero to one.
  
  When `unnormalizedCoordinates` is `VK_TRUE`, images the sampler is used with in the shader have the following requirements:
  
  - The `viewType` **must** be either `VK_IMAGE_VIEW_TYPE_1D` or `VK_IMAGE_VIEW_TYPE_2D`.
  - The image view **must** have a single layer and a single mip level.
    
    When `unnormalizedCoordinates` is `VK_TRUE`, image built-in functions in the shader that use the sampler have the following requirements:
  - The functions **must** not use projection.
  - The functions **must** not use offsets.

## [](#_description)Description

Note

Mapping of OpenGL to Vulkan Filter Modes

`magFilter` values of `VK_FILTER_NEAREST` and `VK_FILTER_LINEAR` directly correspond to `GL_NEAREST` and `GL_LINEAR` magnification filters. `minFilter` and `mipmapMode` combine to correspond to the similarly named OpenGL minification filter of `GL_minFilter_MIPMAP_mipmapMode` (e.g. `minFilter` of `VK_FILTER_LINEAR` and `mipmapMode` of `VK_SAMPLER_MIPMAP_MODE_NEAREST` correspond to `GL_LINEAR_MIPMAP_NEAREST`).

There are no Vulkan filter modes that directly correspond to OpenGL minification filters of `GL_LINEAR` or `GL_NEAREST`, but they **can** be emulated using `VK_SAMPLER_MIPMAP_MODE_NEAREST`, `minLod` = 0, and `maxLod` = 0.25, and using `minFilter` = `VK_FILTER_LINEAR` or `minFilter` = `VK_FILTER_NEAREST`, respectively.

Note that using a `maxLod` of zero would cause [magnification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-texel-filtering) to always be performed, and the `magFilter` to always be used. This is valid, just not an exact match for OpenGL behavior. Clamping the maximum LOD to 0.25 allows the λ value to be non-zero and minification to be performed, while still always rounding down to the base level. If the `minFilter` and `magFilter` are equal, then using a `maxLod` of zero also works.

The maximum number of sampler objects which **can** be simultaneously created on a device is implementation-dependent and specified by the [`maxSamplerAllocationCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSamplerAllocationCount) member of the [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html) structure.

Note

For historical reasons, if `maxSamplerAllocationCount` is exceeded, some implementations may return `VK_ERROR_TOO_MANY_OBJECTS`. Exceeding this limit will result in undefined behavior, and an application should not rely on the use of the returned error code in order to identify when the limit is reached.

Since [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) is a non-dispatchable handle type, implementations **may** return the same handle for sampler state vectors that are identical. In such cases, all such objects would only count once against the `maxSamplerAllocationCount` limit.

Valid Usage

- [](#VUID-VkSamplerCreateInfo-mipLodBias-01069)VUID-VkSamplerCreateInfo-mipLodBias-01069  
  The absolute value of `mipLodBias` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxSamplerLodBias`
- [](#VUID-VkSamplerCreateInfo-samplerMipLodBias-04467)VUID-VkSamplerCreateInfo-samplerMipLodBias-04467  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`samplerMipLodBias` is `VK_FALSE`, `mipLodBias` **must** be zero
- [](#VUID-VkSamplerCreateInfo-maxLod-01973)VUID-VkSamplerCreateInfo-maxLod-01973  
  `maxLod` **must** be greater than or equal to `minLod`
- [](#VUID-VkSamplerCreateInfo-anisotropyEnable-01070)VUID-VkSamplerCreateInfo-anisotropyEnable-01070  
  If the [`samplerAnisotropy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerAnisotropy) feature is not enabled, `anisotropyEnable` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-anisotropyEnable-01071)VUID-VkSamplerCreateInfo-anisotropyEnable-01071  
  If `anisotropyEnable` is `VK_TRUE`, `maxAnisotropy` **must** be between `1.0` and `VkPhysicalDeviceLimits`::`maxSamplerAnisotropy`, inclusive
- [](#VUID-VkSamplerCreateInfo-minFilter-01645)VUID-VkSamplerCreateInfo-minFilter-01645  
  If [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) is enabled and the [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) of the sampler Y′CBCR conversion do not support `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT`, `minFilter` and `magFilter` **must** be equal to the sampler Y′CBCR conversion’s `chromaFilter`
- [](#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01072)VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01072  
  If `unnormalizedCoordinates` is `VK_TRUE`, `minFilter` and `magFilter` **must** be equal
- [](#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01073)VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01073  
  If `unnormalizedCoordinates` is `VK_TRUE`, `mipmapMode` **must** be `VK_SAMPLER_MIPMAP_MODE_NEAREST`
- [](#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01074)VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01074  
  If `unnormalizedCoordinates` is `VK_TRUE`, `minLod` and `maxLod` **must** be zero
- [](#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01075)VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01075  
  If `unnormalizedCoordinates` is `VK_TRUE`, `addressModeU` and `addressModeV` **must** each be either `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE` or `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`
- [](#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01076)VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01076  
  If `unnormalizedCoordinates` is `VK_TRUE`, `anisotropyEnable` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01077)VUID-VkSamplerCreateInfo-unnormalizedCoordinates-01077  
  If `unnormalizedCoordinates` is `VK_TRUE`, `compareEnable` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-addressModeU-01078)VUID-VkSamplerCreateInfo-addressModeU-01078  
  If any of `addressModeU`, `addressModeV` or `addressModeW` are `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`, `borderColor` **must** be a valid [VkBorderColor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBorderColor.html) value
- [](#VUID-VkSamplerCreateInfo-addressModeU-01646)VUID-VkSamplerCreateInfo-addressModeU-01646  
  If [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) is enabled, `addressModeU`, `addressModeV`, and `addressModeW` **must** be `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`, `anisotropyEnable` **must** be `VK_FALSE`, and `unnormalizedCoordinates` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-None-01647)VUID-VkSamplerCreateInfo-None-01647  
  If [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) is enabled and the `pNext` chain includes a [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html) structure, then the sampler reduction mode **must** be `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`
- [](#VUID-VkSamplerCreateInfo-pNext-06726)VUID-VkSamplerCreateInfo-pNext-06726  
  If the [`samplerFilterMinmax`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerFilterMinmax) feature is not enabled and the `pNext` chain includes a [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html) structure, then the sampler reduction mode **must** be `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`
- [](#VUID-VkSamplerCreateInfo-addressModeU-01079)VUID-VkSamplerCreateInfo-addressModeU-01079  
  If the [`samplerMirrorClampToEdge`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerMirrorClampToEdge) feature is not enabled, and if the `VK_KHR_sampler_mirror_clamp_to_edge` extension is not enabled, `addressModeU`, `addressModeV` and `addressModeW` **must** not be `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE`
- [](#VUID-VkSamplerCreateInfo-compareEnable-01080)VUID-VkSamplerCreateInfo-compareEnable-01080  
  If `compareEnable` is `VK_TRUE`, `compareOp` **must** be a valid [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html) value
- [](#VUID-VkSamplerCreateInfo-magFilter-01081)VUID-VkSamplerCreateInfo-magFilter-01081  
  If either `magFilter` or `minFilter` is `VK_FILTER_CUBIC_EXT`, `anisotropyEnable` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-magFilter-07911)VUID-VkSamplerCreateInfo-magFilter-07911  
  If the [VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html) extension is not enabled and either `magFilter` or `minFilter` is `VK_FILTER_CUBIC_IMG`, the `reductionMode` member of [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html) **must** be `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`
- [](#VUID-VkSamplerCreateInfo-compareEnable-01423)VUID-VkSamplerCreateInfo-compareEnable-01423  
  If `compareEnable` is `VK_TRUE`, the `reductionMode` member of [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html) **must** be `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`
- [](#VUID-VkSamplerCreateInfo-flags-02574)VUID-VkSamplerCreateInfo-flags-02574  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then `minFilter` and `magFilter` **must** be equal
- [](#VUID-VkSamplerCreateInfo-flags-02575)VUID-VkSamplerCreateInfo-flags-02575  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then `mipmapMode` **must** be `VK_SAMPLER_MIPMAP_MODE_NEAREST`
- [](#VUID-VkSamplerCreateInfo-flags-02576)VUID-VkSamplerCreateInfo-flags-02576  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then `minLod` and `maxLod` **must** be zero
- [](#VUID-VkSamplerCreateInfo-flags-02577)VUID-VkSamplerCreateInfo-flags-02577  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then `addressModeU` and `addressModeV` **must** each be either `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE` or `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`
- [](#VUID-VkSamplerCreateInfo-flags-02578)VUID-VkSamplerCreateInfo-flags-02578  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then `anisotropyEnable` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-flags-02579)VUID-VkSamplerCreateInfo-flags-02579  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then `compareEnable` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-flags-02580)VUID-VkSamplerCreateInfo-flags-02580  
  If `flags` includes `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`, then `unnormalizedCoordinates` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-nonSeamlessCubeMap-06788)VUID-VkSamplerCreateInfo-nonSeamlessCubeMap-06788  
  If the [`nonSeamlessCubeMap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nonSeamlessCubeMap) feature is not enabled, `flags` **must** not include `VK_SAMPLER_CREATE_NON_SEAMLESS_CUBE_MAP_BIT_EXT`
- [](#VUID-VkSamplerCreateInfo-borderColor-04011)VUID-VkSamplerCreateInfo-borderColor-04011  
  If `borderColor` is one of `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` or `VK_BORDER_COLOR_INT_CUSTOM_EXT`, then a [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html) **must** be included in the `pNext` chain
- [](#VUID-VkSamplerCreateInfo-customBorderColors-04085)VUID-VkSamplerCreateInfo-customBorderColors-04085  
  If the [`customBorderColors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-customBorderColors) feature is not enabled, `borderColor` **must** not be `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` or `VK_BORDER_COLOR_INT_CUSTOM_EXT`
- [](#VUID-VkSamplerCreateInfo-borderColor-04442)VUID-VkSamplerCreateInfo-borderColor-04442  
  If `borderColor` is one of `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` or `VK_BORDER_COLOR_INT_CUSTOM_EXT`, and [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)::`format` is not `VK_FORMAT_UNDEFINED`, [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)::`customBorderColor` **must** be within the range of values representable in `format`
- [](#VUID-VkSamplerCreateInfo-None-04012)VUID-VkSamplerCreateInfo-None-04012  
  The maximum number of samplers with custom border colors which **can** be simultaneously created on a device is implementation-dependent and specified by the [`maxCustomBorderColorSamplers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxCustomBorderColorSamplers) member of the [VkPhysicalDeviceCustomBorderColorPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCustomBorderColorPropertiesEXT.html) structure
- [](#VUID-VkSamplerCreateInfo-flags-08110)VUID-VkSamplerCreateInfo-flags-08110  
  If `flags` includes `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`, the [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-VkSamplerCreateInfo-pNext-08111)VUID-VkSamplerCreateInfo-pNext-08111  
  If the `pNext` chain includes a [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) structure, `flags` **must** contain `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
- [](#VUID-VkSamplerCreateInfo-flags-06964)VUID-VkSamplerCreateInfo-flags-06964  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`, then `minFilter` and `magFilter` **must** be `VK_FILTER_NEAREST`
- [](#VUID-VkSamplerCreateInfo-flags-06965)VUID-VkSamplerCreateInfo-flags-06965  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`, then `mipmapMode` **must** be `VK_SAMPLER_MIPMAP_MODE_NEAREST`
- [](#VUID-VkSamplerCreateInfo-flags-06966)VUID-VkSamplerCreateInfo-flags-06966  
  [If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`, then `minLod` and `maxLod` **must** be zero
- [](#VUID-VkSamplerCreateInfo-flags-06967)VUID-VkSamplerCreateInfo-flags-06967  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`, then `addressModeU` and `addressModeV` **must** each be either `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE` or `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`
- [](#VUID-VkSamplerCreateInfo-flags-06968)VUID-VkSamplerCreateInfo-flags-06968  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`, and if `addressModeU` or `addressModeV` is `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER`, then `borderColor` **must** be `VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK`
- [](#VUID-VkSamplerCreateInfo-flags-06969)VUID-VkSamplerCreateInfo-flags-06969  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`, then `anisotropyEnable` **must** be `VK_FALSE`
- [](#VUID-VkSamplerCreateInfo-flags-06970)VUID-VkSamplerCreateInfo-flags-06970  
  If `flags` includes `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`, then `compareEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkSamplerCreateInfo-sType-sType)VUID-VkSamplerCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO`
- [](#VUID-VkSamplerCreateInfo-pNext-pNext)VUID-VkSamplerCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html), [VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html), [VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html), [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html), [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html), [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html), or [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html)
- [](#VUID-VkSamplerCreateInfo-sType-unique)VUID-VkSamplerCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkSamplerCreateInfo-flags-parameter)VUID-VkSamplerCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateFlagBits.html) values
- [](#VUID-VkSamplerCreateInfo-magFilter-parameter)VUID-VkSamplerCreateInfo-magFilter-parameter  
  `magFilter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html) value
- [](#VUID-VkSamplerCreateInfo-minFilter-parameter)VUID-VkSamplerCreateInfo-minFilter-parameter  
  `minFilter` **must** be a valid [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html) value
- [](#VUID-VkSamplerCreateInfo-mipmapMode-parameter)VUID-VkSamplerCreateInfo-mipmapMode-parameter  
  `mipmapMode` **must** be a valid [VkSamplerMipmapMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerMipmapMode.html) value
- [](#VUID-VkSamplerCreateInfo-addressModeU-parameter)VUID-VkSamplerCreateInfo-addressModeU-parameter  
  `addressModeU` **must** be a valid [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) value
- [](#VUID-VkSamplerCreateInfo-addressModeV-parameter)VUID-VkSamplerCreateInfo-addressModeV-parameter  
  `addressModeV` **must** be a valid [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) value
- [](#VUID-VkSamplerCreateInfo-addressModeW-parameter)VUID-VkSamplerCreateInfo-addressModeW-parameter  
  `addressModeW` **must** be a valid [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkBorderColor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBorderColor.html), [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html), [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html), [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html), [VkSamplerCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateFlags.html), [VkSamplerMipmapMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerMipmapMode.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSampler.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
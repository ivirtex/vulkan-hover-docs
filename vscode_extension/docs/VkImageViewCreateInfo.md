# VkImageViewCreateInfo(3) Manual Page

## Name

VkImageViewCreateInfo - Structure specifying parameters of a newly created image view



## [](#_c_specification)C Specification

The `VkImageViewCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkImageViewCreateInfo {
    VkStructureType            sType;
    const void*                pNext;
    VkImageViewCreateFlags     flags;
    VkImage                    image;
    VkImageViewType            viewType;
    VkFormat                   format;
    VkComponentMapping         components;
    VkImageSubresourceRange    subresourceRange;
} VkImageViewCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateFlagBits.html) specifying additional parameters of the image view.
- `image` is a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) on which the view will be created.
- `viewType` is a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) value specifying the type of the image view.
- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) specifying the format and type used to interpret texel blocks of the image.
- `components` is a [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html) structure specifying a remapping of color components (or of depth or stencil components after they have been converted into color components).
- `subresourceRange` is a [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structure selecting the set of mipmap levels and array layers to be accessible to the view.

## [](#_description)Description

Some of the `image` creation parameters are inherited by the view. In particular, image view creation inherits the implicit parameter `usage` specifying the allowed usages of the image view that, by default, takes the value of the corresponding `usage` parameter specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) at image creation time. The implicit `usage` **can** be overridden by adding a [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewUsageCreateInfo.html) structure to the `pNext` chain, but the view usage **must** be a subset of the image usage. If `image` has a depth-stencil format and was created with a [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html) structure included in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), the usage is calculated based on the `subresource.aspectMask` provided:

- If `aspectMask` includes only `VK_IMAGE_ASPECT_STENCIL_BIT`, the implicit `usage` is equal to [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`.
- If `aspectMask` includes only `VK_IMAGE_ASPECT_DEPTH_BIT`, the implicit `usage` is equal to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage`.
- If both aspects are included in `aspectMask`, the implicit `usage` is equal to the intersection of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` and [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`.

If `image` is a 3D image, its Z range **can** be restricted to a subset by adding a [VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSlicedCreateInfoEXT.html) to the `pNext` chain.

If `image` was created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` flag, and if the `format` of the image is not [multi-planar](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar) `format` **can** be different from the image’s format, but if `image` was created without the `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag and they are not equal they **must** be *compatible*. Image format compatibility is defined in the [Format Compatibility Classes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility-classes) section. Views of compatible formats will have the same mapping between texel coordinates and memory locations irrespective of the `format`, with only the interpretation of the bit pattern changing.

If `image` was created with a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), and the image view’s `aspectMask` is one of `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT` or `VK_IMAGE_ASPECT_PLANE_2_BIT`, the view’s aspect mask is considered to be equivalent to `VK_IMAGE_ASPECT_COLOR_BIT` when used as a framebuffer attachment.

Note

Values intended to be used with one view format **may** not be exactly preserved when written or read through a different format. For example, an integer value that happens to have the bit pattern of a floating-point denorm or NaN **may** be flushed or canonicalized when written or read through a view with a floating-point format. Similarly, a value written through a signed normalized format that has a bit pattern exactly equal to -2b **may** be changed to -2b + 1 as described in [Conversion from Normalized Fixed-Point to Floating-Point](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-fixedfpconv).

If `image` was created with the `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag, `format` **must** be *compatible* with the image’s format as described above; or **must** be an uncompressed format, in which case it **must** be [*size-compatible*](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-size-compatibility) with the image’s format. In this case, the resulting image view’s texel dimensions equal the dimensions of the selected mip level divided by the compressed texel block size and rounded up.

The [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html) `components` member describes a remapping from components of the image to components of the vector returned by shader image instructions. This remapping **must** be the identity swizzle for storage image descriptors, input attachment descriptors, framebuffer attachments, and any `VkImageView` used with a combined image sampler that enables [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion).

If the image view is to be used with a sampler which supports [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion), an *identically defined object* of type [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) to that used to create the sampler **must** be passed to [vkCreateImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImageView.html) in a [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html) included in the `pNext` chain of [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html). Conversely, if a [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) object is passed to [vkCreateImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImageView.html), an [identically defined](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-identically-defined) [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) object **must** be used when sampling the image.

If the image has a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), `subresourceRange.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, and `usage` includes `VK_IMAGE_USAGE_SAMPLED_BIT`, then the `format` **must** be identical to the image `format` and the sampler to be used with the image view **must** enable [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion).

When such an image is used in a [video coding](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding) operation, the [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) has no effect.

If `image` was created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` and the image has a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), and if `subresourceRange.aspectMask` is `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, or `VK_IMAGE_ASPECT_PLANE_2_BIT`, `format` **must** be [compatible](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatible-planes) with the corresponding plane of the image, and the sampler to be used with the image view **must** not enable [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion). The `width` and `height` of the single-plane image view **must** be derived from the multi-planar image’s dimensions in the manner listed for [plane compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatible-planes) for the plane.

Any view of an image plane will have the same mapping between texel coordinates and memory locations as used by the components of the color aspect, subject to the formulae relating texel coordinates to lower-resolution planes as described in [Chroma Reconstruction](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-chroma-reconstruction). That is, if an R or B plane has a reduced resolution relative to the G plane of the multi-planar image, the image view operates using the (*uplane*, *vplane*) unnormalized coordinates of the reduced-resolution plane, and these coordinates access the same memory locations as the (*ucolor*, *vcolor*) unnormalized coordinates of the color aspect for which chroma reconstruction operations operate on the same (*uplane*, *vplane*) or (*iplane*, *jplane*) coordinates.

Table 1. Image Type and Image View Type Compatibility Requirements   Image View Type Compatible Image Types

`VK_IMAGE_VIEW_TYPE_1D`

`VK_IMAGE_TYPE_1D`

`VK_IMAGE_VIEW_TYPE_1D_ARRAY`

`VK_IMAGE_TYPE_1D`

`VK_IMAGE_VIEW_TYPE_2D`

`VK_IMAGE_TYPE_2D` , `VK_IMAGE_TYPE_3D`

`VK_IMAGE_VIEW_TYPE_2D_ARRAY`

`VK_IMAGE_TYPE_2D` , `VK_IMAGE_TYPE_3D`

`VK_IMAGE_VIEW_TYPE_CUBE`

`VK_IMAGE_TYPE_2D`

`VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`

`VK_IMAGE_TYPE_2D`

`VK_IMAGE_VIEW_TYPE_3D`

`VK_IMAGE_TYPE_3D`

Valid Usage

- [](#VUID-VkImageViewCreateInfo-image-01003)VUID-VkImageViewCreateInfo-image-01003  
  If `image` was not created with `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT` then `viewType` **must** not be `VK_IMAGE_VIEW_TYPE_CUBE` or `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`
- [](#VUID-VkImageViewCreateInfo-viewType-01004)VUID-VkImageViewCreateInfo-viewType-01004  
  If the [`imageCubeArray`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-imageCubeArray) feature is not enabled, `viewType` **must** not be `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`
- [](#VUID-VkImageViewCreateInfo-image-06723)VUID-VkImageViewCreateInfo-image-06723  
  If `image` was created with `VK_IMAGE_TYPE_3D` but without `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set then `viewType` **must** not be `VK_IMAGE_VIEW_TYPE_2D_ARRAY`
- [](#VUID-VkImageViewCreateInfo-image-06728)VUID-VkImageViewCreateInfo-image-06728  
  If `image` was created with `VK_IMAGE_TYPE_3D` but without `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` or `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT` set, then `viewType` **must** not be `VK_IMAGE_VIEW_TYPE_2D`
- [](#VUID-VkImageViewCreateInfo-image-04970)VUID-VkImageViewCreateInfo-image-04970  
  If `image` was created with `VK_IMAGE_TYPE_3D` and `viewType` is `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY` then `subresourceRange.levelCount` **must** be 1
- [](#VUID-VkImageViewCreateInfo-image-04972)VUID-VkImageViewCreateInfo-image-04972  
  If `image` was created with a `samples` value not equal to `VK_SAMPLE_COUNT_1_BIT` then `viewType` **must** be either `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`
- [](#VUID-VkImageViewCreateInfo-image-04441)VUID-VkImageViewCreateInfo-image-04441  
  `image` **must** have been created with a `usage` value containing at least one of the usages defined in the [valid image usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#valid-imageview-imageusage) list for image views
- [](#VUID-VkImageViewCreateInfo-None-02273)VUID-VkImageViewCreateInfo-None-02273  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) of the resultant image view **must** contain at least one bit
- [](#VUID-VkImageViewCreateInfo-usage-02274)VUID-VkImageViewCreateInfo-usage-02274  
  If `usage` contains `VK_IMAGE_USAGE_SAMPLED_BIT`, then the [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) of the resultant image view **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`
- [](#VUID-VkImageViewCreateInfo-usage-02275)VUID-VkImageViewCreateInfo-usage-02275  
  If `usage` contains `VK_IMAGE_USAGE_STORAGE_BIT`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT`
- [](#VUID-VkImageViewCreateInfo-usage-08931)VUID-VkImageViewCreateInfo-usage-08931  
  If `usage` contains `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`
- [](#VUID-VkImageViewCreateInfo-usage-02277)VUID-VkImageViewCreateInfo-usage-02277  
  If `usage` contains `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageViewCreateInfo-image-08333)VUID-VkImageViewCreateInfo-image-08333  
  If `usage` contains `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_VIDEO_DECODE_OUTPUT_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-image-08334)VUID-VkImageViewCreateInfo-image-08334  
  If `usage` contains `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_VIDEO_DECODE_DPB_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-image-08335)VUID-VkImageViewCreateInfo-image-08335  
  `usage` **must** not include `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-image-08336)VUID-VkImageViewCreateInfo-image-08336  
  If `usage` contains `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_VIDEO_ENCODE_INPUT_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-image-08337)VUID-VkImageViewCreateInfo-image-08337  
  If `usage` contains `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_VIDEO_ENCODE_DPB_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-image-08338)VUID-VkImageViewCreateInfo-image-08338  
  `usage` **must** not include `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-usage-10259)VUID-VkImageViewCreateInfo-usage-10259  
  If `usage` contains `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-usage-10260)VUID-VkImageViewCreateInfo-usage-10260  
  If `usage` contains `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-usage-08932)VUID-VkImageViewCreateInfo-usage-08932  
  If `usage` contains `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`, and any of the following is true:
  
  - the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is not enabled
  - the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) property is `VK_FALSE`
  - `image` was created with an [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)::`externalFormat` value of 0
  
  then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain at least one of `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT` or `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`
- [](#VUID-VkImageViewCreateInfo-subresourceRange-01478)VUID-VkImageViewCreateInfo-subresourceRange-01478  
  `subresourceRange.baseMipLevel` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageViewCreateInfo-subresourceRange-01718)VUID-VkImageViewCreateInfo-subresourceRange-01718  
  If `subresourceRange.levelCount` is not `VK_REMAINING_MIP_LEVELS`, `subresourceRange.baseMipLevel` + `subresourceRange.levelCount` **must** be less than or equal to the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageViewCreateInfo-image-02571)VUID-VkImageViewCreateInfo-image-02571  
  If `image` was created with `usage` containing `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`, `subresourceRange.levelCount` **must** be `1`
- [](#VUID-VkImageViewCreateInfo-image-06724)VUID-VkImageViewCreateInfo-image-06724  
  If `image` is not a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` or `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT` set, or `viewType` is not `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, `subresourceRange.baseArrayLayer` **must** be less than the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageViewCreateInfo-subresourceRange-06725)VUID-VkImageViewCreateInfo-subresourceRange-06725  
  If `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `image` is not a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` or `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT` set, or `viewType` is not `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, `subresourceRange.layerCount` **must** be non-zero and `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageViewCreateInfo-image-02724)VUID-VkImageViewCreateInfo-image-02724  
  If `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, and `viewType` is `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, `subresourceRange.baseArrayLayer` **must** be less than the depth computed from `baseMipLevel` and `extent.depth` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created, according to the formula defined in [Image Mip Level Sizing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-mip-level-sizing)
- [](#VUID-VkImageViewCreateInfo-subresourceRange-02725)VUID-VkImageViewCreateInfo-subresourceRange-02725  
  If `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, and `viewType` is `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, `subresourceRange.layerCount` **must** be non-zero and `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount` **must** be less than or equal to the depth computed from `baseMipLevel` and `extent.depth` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created, according to the formula defined in [Image Mip Level Sizing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-mip-level-sizing)
- [](#VUID-VkImageViewCreateInfo-image-01761)VUID-VkImageViewCreateInfo-image-01761  
  If `image` was created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` flag, but without the `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag, and if the `format` of the `image` is not a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), `format` **must** be compatible with the `format` used to create `image`, as defined in [Format Compatibility Classes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility-classes)
- [](#VUID-VkImageViewCreateInfo-image-01583)VUID-VkImageViewCreateInfo-image-01583  
  If `image` was created with the `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag, `format` **must** be compatible with, or **must** be an uncompressed format that is [size-compatible](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-size-compatibility) with, the `format` used to create `image`
- [](#VUID-VkImageViewCreateInfo-image-07072)VUID-VkImageViewCreateInfo-image-07072  
  If `image` was created with the `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag and `format` is a non-compressed format, the `levelCount` member of `subresourceRange` **must** be `1`
- [](#VUID-VkImageViewCreateInfo-image-09487)VUID-VkImageViewCreateInfo-image-09487  
  If `image` was created with the `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag, the `VkPhysicalDeviceMaintenance6Properties`::`blockTexelViewCompatibleMultipleLayers` property is not `VK_TRUE`, and `format` is a non-compressed format, then the `layerCount` member of `subresourceRange` **must** be `1`
- [](#VUID-VkImageViewCreateInfo-pNext-01585)VUID-VkImageViewCreateInfo-pNext-01585  
  If a [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html) structure was included in the `pNext` chain of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure used when creating `image` and [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount` is not zero then `format` **must** be one of the formats in [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats`
- [](#VUID-VkImageViewCreateInfo-image-01586)VUID-VkImageViewCreateInfo-image-01586  
  If `image` was created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` flag, if the `format` of the `image` is a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), and if `subresourceRange.aspectMask` is one of the [multi-planar aspect mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar-image-aspect) bits, then `format` **must** be compatible with the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) for the plane of the `image` `format` indicated by `subresourceRange.aspectMask`, as defined in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatible-planes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatible-planes)
- [](#VUID-VkImageViewCreateInfo-subresourceRange-07818)VUID-VkImageViewCreateInfo-subresourceRange-07818  
  `subresourceRange.aspectMask` **must** only have at most 1 valid [multi-planar aspect mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar-image-aspect) bit
- [](#VUID-VkImageViewCreateInfo-image-01762)VUID-VkImageViewCreateInfo-image-01762  
  If `image` was not created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` flag, or if the `format` of the `image` is a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar) and if `subresourceRange.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, `format` **must** be identical to the `format` used to create `image`
- [](#VUID-VkImageViewCreateInfo-format-06415)VUID-VkImageViewCreateInfo-format-06415  
  If the image view [requires a sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#image-views-requiring-sampler-ycbcr-conversion) and `usage` contains `VK_IMAGE_USAGE_SAMPLED_BIT`, then the `pNext` chain **must** include a [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html) structure with a conversion value other than [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkImageViewCreateInfo-format-04714)VUID-VkImageViewCreateInfo-format-04714  
  If `format` has a `_422` or `_420` suffix then `image` **must** have been created with a width that is a multiple of 2
- [](#VUID-VkImageViewCreateInfo-format-04715)VUID-VkImageViewCreateInfo-format-04715  
  If `format` has a `_420` suffix then `image` **must** have been created with a height that is a multiple of 2
- [](#VUID-VkImageViewCreateInfo-pNext-01970)VUID-VkImageViewCreateInfo-pNext-01970  
  If the `pNext` chain includes a [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html) structure with a `conversion` value other than [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), all members of `components` **must** have the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings)
- [](#VUID-VkImageViewCreateInfo-pNext-06658)VUID-VkImageViewCreateInfo-pNext-06658  
  If the `pNext` chain includes a [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html) structure with a `conversion` value other than [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `format` **must** be the same used in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html)::`format`
- [](#VUID-VkImageViewCreateInfo-image-01020)VUID-VkImageViewCreateInfo-image-01020  
  If `image` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkImageViewCreateInfo-subResourceRange-01021)VUID-VkImageViewCreateInfo-subResourceRange-01021  
  `viewType` **must** be compatible with the type of `image` as shown in the [view type compatibility table](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-compatibility)
- [](#VUID-VkImageViewCreateInfo-image-02399)VUID-VkImageViewCreateInfo-image-02399  
  If `image` has an [Android external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-android-hardware-buffer-external-formats), `format` **must** be `VK_FORMAT_UNDEFINED`
- [](#VUID-VkImageViewCreateInfo-image-02400)VUID-VkImageViewCreateInfo-image-02400  
  If `image` has an [Android external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-android-hardware-buffer-external-formats), the `pNext` chain **must** include a [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html) structure with a `conversion` object created with the same external format as `image`
- [](#VUID-VkImageViewCreateInfo-image-02401)VUID-VkImageViewCreateInfo-image-02401  
  If `image` has an [Android external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-android-hardware-buffer-external-formats), all members of `components` **must** be the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings)
- [](#VUID-VkImageViewCreateInfo-image-08957)VUID-VkImageViewCreateInfo-image-08957  
  If `image` has an [QNX Screen external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-screen-buffer-external-formats), `format` **must** be `VK_FORMAT_UNDEFINED`
- [](#VUID-VkImageViewCreateInfo-image-08958)VUID-VkImageViewCreateInfo-image-08958  
  If `image` has an [QNX Screen external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-screen-buffer-external-formats), the `pNext` chain **must** include a [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html) structure with a `conversion` object created with the same external format as `image`
- [](#VUID-VkImageViewCreateInfo-image-08959)VUID-VkImageViewCreateInfo-image-08959  
  If `image` has an [QNX Screen external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-screen-buffer-external-formats), all members of `components` **must** be the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings)
- [](#VUID-VkImageViewCreateInfo-image-02086)VUID-VkImageViewCreateInfo-image-02086  
  If `image` was created with `usage` containing `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, `viewType` **must** be `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`
- [](#VUID-VkImageViewCreateInfo-image-02087)VUID-VkImageViewCreateInfo-image-02087  
  If the [`shadingRateImage`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shadingRateImage) feature is enabled, and If `image` was created with `usage` containing `VK_IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV`, `format` **must** be `VK_FORMAT_R8_UINT`
- [](#VUID-VkImageViewCreateInfo-usage-04550)VUID-VkImageViewCreateInfo-usage-04550  
  If the [`attachmentFragmentShadingRate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-attachmentFragmentShadingRate) feature is enabled, and the `usage` for the image view includes `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, then the image view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkImageViewCreateInfo-usage-04551)VUID-VkImageViewCreateInfo-usage-04551  
  If the [`attachmentFragmentShadingRate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-attachmentFragmentShadingRate) feature is enabled, the `usage` for the image view includes `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, and [`layeredShadingRateAttachments`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-layeredShadingRateAttachments) is `VK_FALSE`, `subresourceRange.layerCount` **must** be `1`
- [](#VUID-VkImageViewCreateInfo-flags-02572)VUID-VkImageViewCreateInfo-flags-02572  
  If the [`fragmentDensityMapDynamic`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fragmentDensityMapDynamic) feature is not enabled, `flags` **must** not contain `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT`
- [](#VUID-VkImageViewCreateInfo-flags-03567)VUID-VkImageViewCreateInfo-flags-03567  
  If the [`fragmentDensityMapDeferred`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fragmentDensityMapDeferred) feature is not enabled, `flags` **must** not contain `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT`
- [](#VUID-VkImageViewCreateInfo-flags-03568)VUID-VkImageViewCreateInfo-flags-03568  
  If `flags` contains `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT`, `flags` **must** not contain `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT`
- [](#VUID-VkImageViewCreateInfo-image-03569)VUID-VkImageViewCreateInfo-image-03569  
  If `image` was created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT` and `usage` containing `VK_IMAGE_USAGE_SAMPLED_BIT`, `subresourceRange.layerCount` **must** be less than or equal to [`VkPhysicalDeviceFragmentDensityMap2PropertiesEXT`::`maxSubsampledArrayLayers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSubsampledArrayLayers)
- [](#VUID-VkImageViewCreateInfo-invocationMask-04993)VUID-VkImageViewCreateInfo-invocationMask-04993  
  If the [`invocationMask`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-invocationMask) feature is enabled, and if `image` was created with `usage` containing `VK_IMAGE_USAGE_INVOCATION_MASK_BIT_HUAWEI`, `format` **must** be `VK_FORMAT_R8_UINT`
- [](#VUID-VkImageViewCreateInfo-flags-04116)VUID-VkImageViewCreateInfo-flags-04116  
  If `flags` does not contain `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT` and `image` was created with `usage` containing `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`, its `flags` **must** not contain any of `VK_IMAGE_CREATE_PROTECTED_BIT`, `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`
- [](#VUID-VkImageViewCreateInfo-pNext-02662)VUID-VkImageViewCreateInfo-pNext-02662  
  If the `pNext` chain includes a [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewUsageCreateInfo.html) structure, and `image` was not created with a [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html) structure included in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), its `usage` member **must** not include any bits that were not set in the `usage` member of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure used to create `image`
- [](#VUID-VkImageViewCreateInfo-pNext-02663)VUID-VkImageViewCreateInfo-pNext-02663  
  If the `pNext` chain includes a [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewUsageCreateInfo.html) structure, `image` was created with a [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html) structure included in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), and `subresourceRange.aspectMask` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, the `usage` member of the [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewUsageCreateInfo.html) structure **must** not include any bits that were not set in the `usage` member of the [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html) structure used to create `image`
- [](#VUID-VkImageViewCreateInfo-pNext-02664)VUID-VkImageViewCreateInfo-pNext-02664  
  If the `pNext` chain includes a [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewUsageCreateInfo.html) structure, `image` was created with a [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html) structure included in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), and `subresourceRange.aspectMask` includes bits other than `VK_IMAGE_ASPECT_STENCIL_BIT`, the `usage` member of the [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewUsageCreateInfo.html) structure **must** not include any bits that were not set in the `usage` member of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure used to create `image`
- [](#VUID-VkImageViewCreateInfo-imageViewType-04973)VUID-VkImageViewCreateInfo-imageViewType-04973  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_1D`, `VK_IMAGE_VIEW_TYPE_2D`, or `VK_IMAGE_VIEW_TYPE_3D`; and `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, then `subresourceRange.layerCount` **must** be 1
- [](#VUID-VkImageViewCreateInfo-imageViewType-04974)VUID-VkImageViewCreateInfo-imageViewType-04974  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_1D`, `VK_IMAGE_VIEW_TYPE_2D`, or `VK_IMAGE_VIEW_TYPE_3D`; and `subresourceRange.layerCount` is `VK_REMAINING_ARRAY_LAYERS`, then the remaining number of layers **must** be 1
- [](#VUID-VkImageViewCreateInfo-viewType-02960)VUID-VkImageViewCreateInfo-viewType-02960  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_CUBE` and `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `subresourceRange.layerCount` **must** be `6`
- [](#VUID-VkImageViewCreateInfo-viewType-02961)VUID-VkImageViewCreateInfo-viewType-02961  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY` and `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `subresourceRange.layerCount` **must** be a multiple of `6`
- [](#VUID-VkImageViewCreateInfo-viewType-02962)VUID-VkImageViewCreateInfo-viewType-02962  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_CUBE` and `subresourceRange.layerCount` is `VK_REMAINING_ARRAY_LAYERS`, the remaining number of layers **must** be `6`
- [](#VUID-VkImageViewCreateInfo-viewType-02963)VUID-VkImageViewCreateInfo-viewType-02963  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY` and `subresourceRange.layerCount` is `VK_REMAINING_ARRAY_LAYERS`, the remaining number of layers **must** be a multiple of `6`
- [](#VUID-VkImageViewCreateInfo-imageViewFormatSwizzle-04465)VUID-VkImageViewCreateInfo-imageViewFormatSwizzle-04465  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`imageViewFormatSwizzle` is `VK_FALSE`, all elements of `components` **must** have the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings)
- [](#VUID-VkImageViewCreateInfo-imageViewFormatReinterpretation-04466)VUID-VkImageViewCreateInfo-imageViewFormatReinterpretation-04466  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`imageViewFormatReinterpretation` is `VK_FALSE`, the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) in `format` **must** not contain a different number of components, or a different number of bits in each component, than the format of the `VkImage` in `image`
- [](#VUID-VkImageViewCreateInfo-image-04817)VUID-VkImageViewCreateInfo-image-04817  
  If `image` was created with `usage` containing `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`, or `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, then the `viewType` **must** be `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`
- [](#VUID-VkImageViewCreateInfo-image-04818)VUID-VkImageViewCreateInfo-image-04818  
  If `image` was created with `usage` containing `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`, `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`, or `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`, then the `viewType` **must** be `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`
- [](#VUID-VkImageViewCreateInfo-image-10261)VUID-VkImageViewCreateInfo-image-10261  
  If `image` was created with `usage` containing `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`, then `viewType` **must** be `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`
- [](#VUID-VkImageViewCreateInfo-flags-08106)VUID-VkImageViewCreateInfo-flags-08106  
  If `flags` includes `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`, the [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-VkImageViewCreateInfo-pNext-08107)VUID-VkImageViewCreateInfo-pNext-08107  
  If the `pNext` chain includes a [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) structure, `flags` **must** contain `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
- [](#VUID-VkImageViewCreateInfo-pNext-06787)VUID-VkImageViewCreateInfo-pNext-06787  
  If the `pNext` chain includes a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure, its `exportObjectType` member **must** be `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT`
- [](#VUID-VkImageViewCreateInfo-pNext-06944)VUID-VkImageViewCreateInfo-pNext-06944  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure, then [`textureSampleWeighted`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-textureSampleWeighted) feature **must** be enabled
- [](#VUID-VkImageViewCreateInfo-pNext-06945)VUID-VkImageViewCreateInfo-pNext-06945  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure, then `image` **must** have been created with `usage` containing `VK_IMAGE_USAGE_SAMPLE_WEIGHT_BIT_QCOM`
- [](#VUID-VkImageViewCreateInfo-pNext-06946)VUID-VkImageViewCreateInfo-pNext-06946  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure, then `components` **must** be `VK_COMPONENT_SWIZZLE_IDENTITY` for all components
- [](#VUID-VkImageViewCreateInfo-pNext-06947)VUID-VkImageViewCreateInfo-pNext-06947  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure, then `subresourceRange.aspectMask` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkImageViewCreateInfo-pNext-06948)VUID-VkImageViewCreateInfo-pNext-06948  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure, then `subresourceRange.levelCount` **must** be `1`
- [](#VUID-VkImageViewCreateInfo-pNext-06949)VUID-VkImageViewCreateInfo-pNext-06949  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure, then `viewType` **must** be `VK_IMAGE_VIEW_TYPE_1D_ARRAY` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`
- [](#VUID-VkImageViewCreateInfo-pNext-06950)VUID-VkImageViewCreateInfo-pNext-06950  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure and if `viewType` is `VK_IMAGE_VIEW_TYPE_1D_ARRAY`, then `image` **must** have been created with `imageType` `VK_IMAGE_TYPE_1D`
- [](#VUID-VkImageViewCreateInfo-pNext-06951)VUID-VkImageViewCreateInfo-pNext-06951  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure and `viewType` is `VK_IMAGE_VIEW_TYPE_1D_ARRAY`, then `subresourceRange.layerCount` **must** be equal to `2`
- [](#VUID-VkImageViewCreateInfo-pNext-06952)VUID-VkImageViewCreateInfo-pNext-06952  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure and `viewType` is `VK_IMAGE_VIEW_TYPE_1D_ARRAY`, then `image` **must** have been created with `width` equal to or greater than (numPhases×max(align(filterSize.width,4),filterSize.height))
- [](#VUID-VkImageViewCreateInfo-pNext-06953)VUID-VkImageViewCreateInfo-pNext-06953  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure and if `viewType` is `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, then `image` **must** have been created with `imageType` `VK_IMAGE_TYPE_2D`
- [](#VUID-VkImageViewCreateInfo-pNext-06954)VUID-VkImageViewCreateInfo-pNext-06954  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure and `viewType` is `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, then `subresourceRange.layerCount` **must** be equal or greater than numPhases
- [](#VUID-VkImageViewCreateInfo-pNext-06955)VUID-VkImageViewCreateInfo-pNext-06955  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure and `viewType` is `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, then `image` **must** have been created with `width` equal to or greater than `filterSize.width`
- [](#VUID-VkImageViewCreateInfo-pNext-06956)VUID-VkImageViewCreateInfo-pNext-06956  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure and `viewType` is `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, then `image` **must** have been created with `height` equal to or greater than `filterSize.height`
- [](#VUID-VkImageViewCreateInfo-pNext-06957)VUID-VkImageViewCreateInfo-pNext-06957  
  If the `pNext` chain includes [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html) structure then [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)::`filterSize.height` **must** be less than or equal to [`VkPhysicalDeviceImageProcessingPropertiesQCOM`::`maxWeightFilterDimension.height`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-weightfilter-maxdimension)
- [](#VUID-VkImageViewCreateInfo-subresourceRange-09594)VUID-VkImageViewCreateInfo-subresourceRange-09594  
  `subresourceRange.aspectMask` **must** be valid for the `format` the `image` was created with

Valid Usage (Implicit)

- [](#VUID-VkImageViewCreateInfo-sType-sType)VUID-VkImageViewCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO`
- [](#VUID-VkImageViewCreateInfo-pNext-pNext)VUID-VkImageViewCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html), [VkImageViewASTCDecodeModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewASTCDecodeModeEXT.html), [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewMinLodCreateInfoEXT.html), [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html), [VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSlicedCreateInfoEXT.html), [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewUsageCreateInfo.html), [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html), or [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfo.html)
- [](#VUID-VkImageViewCreateInfo-sType-unique)VUID-VkImageViewCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique, with the exception of structures of type [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html)
- [](#VUID-VkImageViewCreateInfo-flags-parameter)VUID-VkImageViewCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateFlagBits.html) values
- [](#VUID-VkImageViewCreateInfo-image-parameter)VUID-VkImageViewCreateInfo-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkImageViewCreateInfo-viewType-parameter)VUID-VkImageViewCreateInfo-viewType-parameter  
  `viewType` **must** be a valid [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) value
- [](#VUID-VkImageViewCreateInfo-format-parameter)VUID-VkImageViewCreateInfo-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkImageViewCreateInfo-components-parameter)VUID-VkImageViewCreateInfo-components-parameter  
  `components` **must** be a valid [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html) structure
- [](#VUID-VkImageViewCreateInfo-subresourceRange-parameter)VUID-VkImageViewCreateInfo-subresourceRange-parameter  
  `subresourceRange` **must** be a valid [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html), [VkImageViewCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateFlags.html), [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImageView.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageViewCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
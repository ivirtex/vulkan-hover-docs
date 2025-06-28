# VkFormatFeatureFlagBits2(3) Manual Page

## Name

VkFormatFeatureFlagBits2 - Bitmask specifying features supported by a buffer



## [](#_c_specification)C Specification

Bits which **can** be set in the [VkFormatProperties3](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3.html) features `linearTilingFeatures`, `optimalTilingFeatures`, and `bufferFeatures` are:

```c++
// Provided by VK_VERSION_1_3
// Flag bits for VkFormatFeatureFlagBits2
typedef VkFlags64 VkFormatFeatureFlagBits2;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_BIT = 0x00000001ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_IMAGE_BIT = 0x00000002ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_IMAGE_ATOMIC_BIT = 0x00000004ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_UNIFORM_TEXEL_BUFFER_BIT = 0x00000008ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_TEXEL_BUFFER_BIT = 0x00000010ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_TEXEL_BUFFER_ATOMIC_BIT = 0x00000020ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_VERTEX_BUFFER_BIT = 0x00000040ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_COLOR_ATTACHMENT_BIT = 0x00000080ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_COLOR_ATTACHMENT_BLEND_BIT = 0x00000100ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_DEPTH_STENCIL_ATTACHMENT_BIT = 0x00000200ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_BLIT_SRC_BIT = 0x00000400ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_BLIT_DST_BIT = 0x00000800ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_LINEAR_BIT = 0x00001000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_TRANSFER_SRC_BIT = 0x00004000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_TRANSFER_DST_BIT = 0x00008000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_MINMAX_BIT = 0x00010000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_MIDPOINT_CHROMA_SAMPLES_BIT = 0x00020000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT = 0x00040000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT = 0x00080000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT = 0x00100000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT = 0x00200000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_DISJOINT_BIT = 0x00400000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_COSITED_CHROMA_SAMPLES_BIT = 0x00800000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT = 0x80000000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT = 0x100000000ULL;
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT = 0x200000000ULL;
// Provided by VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_CUBIC_BIT = 0x00002000ULL;
// Provided by VK_VERSION_1_4
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_HOST_IMAGE_TRANSFER_BIT = 0x400000000000ULL;
// Provided by VK_KHR_video_decode_queue with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_VIDEO_DECODE_OUTPUT_BIT_KHR = 0x02000000ULL;
// Provided by VK_KHR_video_decode_queue with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_VIDEO_DECODE_DPB_BIT_KHR = 0x04000000ULL;
// Provided by VK_KHR_acceleration_structure with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR = 0x20000000ULL;
// Provided by VK_EXT_fragment_density_map with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_FRAGMENT_DENSITY_MAP_BIT_EXT = 0x01000000ULL;
// Provided by VK_KHR_fragment_shading_rate with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR = 0x40000000ULL;
// Provided by VK_EXT_host_image_copy
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_HOST_IMAGE_TRANSFER_BIT_EXT = 0x400000000000ULL;
// Provided by VK_KHR_video_encode_queue with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_VIDEO_ENCODE_INPUT_BIT_KHR = 0x08000000ULL;
// Provided by VK_KHR_video_encode_queue with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_VIDEO_ENCODE_DPB_BIT_KHR = 0x10000000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_BIT_KHR = 0x00000001ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_IMAGE_BIT_KHR = 0x00000002ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_IMAGE_ATOMIC_BIT_KHR = 0x00000004ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_UNIFORM_TEXEL_BUFFER_BIT_KHR = 0x00000008ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_TEXEL_BUFFER_BIT_KHR = 0x00000010ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_TEXEL_BUFFER_ATOMIC_BIT_KHR = 0x00000020ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_VERTEX_BUFFER_BIT_KHR = 0x00000040ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_COLOR_ATTACHMENT_BIT_KHR = 0x00000080ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_COLOR_ATTACHMENT_BLEND_BIT_KHR = 0x00000100ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_DEPTH_STENCIL_ATTACHMENT_BIT_KHR = 0x00000200ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_BLIT_SRC_BIT_KHR = 0x00000400ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_BLIT_DST_BIT_KHR = 0x00000800ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_LINEAR_BIT_KHR = 0x00001000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_TRANSFER_SRC_BIT_KHR = 0x00004000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_TRANSFER_DST_BIT_KHR = 0x00008000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_MIDPOINT_CHROMA_SAMPLES_BIT_KHR = 0x00020000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR = 0x00040000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR = 0x00080000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR = 0x00100000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR = 0x00200000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_DISJOINT_BIT_KHR = 0x00400000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_COSITED_CHROMA_SAMPLES_BIT_KHR = 0x00800000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT_KHR = 0x80000000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT_KHR = 0x100000000ULL;
// Provided by VK_KHR_format_feature_flags2
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT_KHR = 0x200000000ULL;
// Provided by VK_KHR_format_feature_flags2 with VK_VERSION_1_2 or VK_EXT_sampler_filter_minmax
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_MINMAX_BIT_KHR = 0x00010000ULL;
// Provided by VK_KHR_format_feature_flags2 with VK_EXT_filter_cubic or VK_IMG_filter_cubic
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT = 0x00002000ULL;
// Provided by VK_NV_ray_tracing_linear_swept_spheres
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_ACCELERATION_STRUCTURE_RADIUS_BUFFER_BIT_NV = 0x8000000000000ULL;
// Provided by VK_NV_linear_color_attachment with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV = 0x4000000000ULL;
// Provided by VK_QCOM_image_processing with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM = 0x400000000ULL;
// Provided by VK_QCOM_image_processing with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM = 0x800000000ULL;
// Provided by VK_QCOM_image_processing with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM = 0x1000000000ULL;
// Provided by VK_QCOM_image_processing with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM = 0x2000000000ULL;
// Provided by VK_ARM_tensors
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_TENSOR_SHADER_BIT_ARM = 0x8000000000ULL;
// Provided by VK_ARM_tensors
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_TENSOR_IMAGE_ALIASING_BIT_ARM = 0x80000000000ULL;
// Provided by VK_NV_optical_flow
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_OPTICAL_FLOW_IMAGE_BIT_NV = 0x10000000000ULL;
// Provided by VK_NV_optical_flow
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_OPTICAL_FLOW_VECTOR_BIT_NV = 0x20000000000ULL;
// Provided by VK_NV_optical_flow
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_OPTICAL_FLOW_COST_BIT_NV = 0x40000000000ULL;
// Provided by VK_ARM_data_graph
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_TENSOR_DATA_GRAPH_BIT_ARM = 0x1000000000000ULL;
// Provided by VK_KHR_video_encode_quantization_map
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR = 0x2000000000000ULL;
// Provided by VK_KHR_video_encode_quantization_map
static const VkFormatFeatureFlagBits2 VK_FORMAT_FEATURE_2_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR = 0x4000000000000ULL;
```

or the equivalent

```c++
// Provided by VK_KHR_format_feature_flags2
typedef VkFormatFeatureFlagBits2 VkFormatFeatureFlagBits2KHR;
```

## [](#_description)Description

The following bits **may** be set in `linearTilingFeatures` and `optimalTilingFeatures`, specifying that the features are supported by [images](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) or [image views](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) or [sampler Y′CBCR conversion objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) created with the queried [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2.html)::`format`:

- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_BIT` specifies that an image view **can** be [sampled from](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-sampledimage).
- `VK_FORMAT_FEATURE_2_STORAGE_IMAGE_BIT` specifies that an image view **can** be used as a [storage image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage).
- `VK_FORMAT_FEATURE_2_STORAGE_IMAGE_ATOMIC_BIT` specifies that an image view **can** be used as storage image that supports atomic operations.
- `VK_FORMAT_FEATURE_2_COLOR_ATTACHMENT_BIT` specifies that an image view **can** be used as a framebuffer color attachment and as an input attachment.
- `VK_FORMAT_FEATURE_2_COLOR_ATTACHMENT_BLEND_BIT` specifies that an image view **can** be used as a framebuffer color attachment that supports blending.
- `VK_FORMAT_FEATURE_2_DEPTH_STENCIL_ATTACHMENT_BIT` specifies that an image view **can** be used as a framebuffer depth/stencil attachment and as an input attachment.
- `VK_FORMAT_FEATURE_2_BLIT_SRC_BIT` specifies that an image **can** be used as the `srcImage` for [vkCmdBlitImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2.html) and [vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage.html).
- `VK_FORMAT_FEATURE_2_BLIT_DST_BIT` specifies that an image **can** be used as the `dstImage` for [vkCmdBlitImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2.html) and [vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage.html).
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_LINEAR_BIT` specifies that if `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_BIT` is also set, an image view **can** be used with a sampler that has either of `magFilter` or `minFilter` set to `VK_FILTER_LINEAR`, or `mipmapMode` set to `VK_SAMPLER_MIPMAP_MODE_LINEAR`. If `VK_FORMAT_FEATURE_2_BLIT_SRC_BIT` is also set, an image can be used as the `srcImage` for [vkCmdBlitImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2.html) and `vkCmdBlitImage` with a `filter` of `VK_FILTER_LINEAR`. This bit **must** only be exposed for formats that also support the `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_BIT` or `VK_FORMAT_FEATURE_2_BLIT_SRC_BIT`.
  
  If the format being queried is a depth/stencil format, this bit only specifies that the depth aspect (not the stencil aspect) of an image of this format supports linear filtering. Where depth comparison is supported it **may** be linear filtered whether this bit is present or not, but where this bit is not present the filtered value **may** be computed in an implementation-dependent manner which differs from the normal rules of linear filtering. The resulting value **must** be in the range \[0,1] and **should** be proportional to, or a weighted average of, the number of comparison passes or failures.
- `VK_FORMAT_FEATURE_2_TRANSFER_SRC_BIT` specifies that an image **can** be used as a source image for [copy commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies).
- `VK_FORMAT_FEATURE_2_TRANSFER_DST_BIT` specifies that an image **can** be used as a destination image for [copy commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies) and [clear commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears).
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_MINMAX_BIT` specifies `VkImage` **can** be used as a sampled image with a min or max [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionMode.html). This bit **must** only be exposed for formats that also support the `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_BIT`.
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_CUBIC_BIT` specifies that `VkImage` **can** be used with a sampler that has either of `magFilter` or `minFilter` set to `VK_FILTER_CUBIC_EXT`, or be the source image for a blit with `filter` set to `VK_FILTER_CUBIC_EXT`. This bit **must** only be exposed for formats that also support the `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_BIT`. If the format being queried is a depth/stencil format, this only specifies that the depth aspect is cubic filterable.
- `VK_FORMAT_FEATURE_2_MIDPOINT_CHROMA_SAMPLES_BIT` specifies that an application **can** define a [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) using this format as a source, and that an image of this format **can** be used with a [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html) `xChromaOffset` and/or `yChromaOffset` of `VK_CHROMA_LOCATION_MIDPOINT`. Otherwise both `xChromaOffset` and `yChromaOffset` **must** be `VK_CHROMA_LOCATION_COSITED_EVEN`. If a format does not incorporate chroma downsampling (it is not a “422” or “420” format) but the implementation supports sampler Y′CBCR conversion for this format, the implementation **must** set `VK_FORMAT_FEATURE_2_MIDPOINT_CHROMA_SAMPLES_BIT`.
- `VK_FORMAT_FEATURE_2_COSITED_CHROMA_SAMPLES_BIT` specifies that an application **can** define a [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) using this format as a source, and that an image of this format **can** be used with a [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html) `xChromaOffset` and/or `yChromaOffset` of `VK_CHROMA_LOCATION_COSITED_EVEN`. Otherwise both `xChromaOffset` and `yChromaOffset` **must** be `VK_CHROMA_LOCATION_MIDPOINT`. If neither `VK_FORMAT_FEATURE_2_COSITED_CHROMA_SAMPLES_BIT` nor `VK_FORMAT_FEATURE_2_MIDPOINT_CHROMA_SAMPLES_BIT` is set, the application **must** not define a [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) using this format as a source.
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT` specifies that an application **can** define a [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) using this format as a source with `chromaFilter` set to `VK_FILTER_LINEAR`.
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT` specifies that the format can have different chroma, min, and mag filters.
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT` specifies that reconstruction is explicit, as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-chroma-reconstruction](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-chroma-reconstruction). If this bit is not present, reconstruction is implicit by default.
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT` specifies that reconstruction **can** be forcibly made explicit by setting [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html)::`forceExplicitReconstruction` to `VK_TRUE`. If the format being queried supports `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT` it **must** also support `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT`.
- `VK_FORMAT_FEATURE_2_DISJOINT_BIT` specifies that a multi-planar image **can** have the `VK_IMAGE_CREATE_DISJOINT_BIT` set during image creation. An implementation **must** not set `VK_FORMAT_FEATURE_2_DISJOINT_BIT` for *single-plane formats*.
- `VK_FORMAT_FEATURE_2_FRAGMENT_DENSITY_MAP_BIT_EXT` specifies that an image view **can** be used as a [fragment density map attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-fragmentdensitymapattachment).
- `VK_FORMAT_FEATURE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` specifies that an image view **can** be used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment). An implementation **must** not set this feature for formats with a [numeric format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-numericformat) other than `UINT`, or set it as a buffer feature.
- `VK_FORMAT_FEATURE_2_VIDEO_DECODE_OUTPUT_BIT_KHR` specifies that an image view with this format **can** be used as a [decode output picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-output-picture) in [video decode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-decode-operations).
- `VK_FORMAT_FEATURE_2_VIDEO_DECODE_DPB_BIT_KHR` specifies that an image view with this format **can** be used as an output [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#reconstructed-picture) or an input [reference picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#reference-picture) in [video decode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-decode-operations).
- `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_INPUT_BIT_KHR` specifies that an image view with this format **can** be used as an [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture) in [video encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations).
- `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_DPB_BIT_KHR` specifies that an image view with this format **can** be used as an output [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#reconstructed-picture) or an input [reference picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#reference-picture) in [video encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations).
  
  Note
  
  Specific [video profiles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles) **may** have additional restrictions on the format and other image creation parameters corresponding to image views used by video coding operations that **can** be enumerated using the [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html) command.
- `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT` specifies that image views or buffer views created with this format **can** be used as [storage images](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage) or [storage texel buffers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer) respectively for read operations without specifying a format.
- `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT` specifies that image views or buffer views created with this format **can** be used as [storage images](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage) or [storage texel buffers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer) respectively for write operations without specifying a format.
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT` specifies that image views created with this format **can** be used for depth comparison performed by `OpImage*Dref*` instructions.
- `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV` specifies that the format is supported as a renderable [Linear Color Attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-linear-color-attachment). This bit will be set for renderable color formats in the `linearTilingFeatures`. This **must** not be set in the `optimalTilingFeatures` or `bufferFeatures` members.
- `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM` specifies that image views created with this format **can** be used as the [weight image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-weightimage) input to [weight image sampling](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-weightimage) operations.
- `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM` specifies that image views created with this format **can** be sampled in [weight image sampling](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-weightimage) operations.
- `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM` specifies that image views created with this format **can** be used in [block matching](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-blockmatch) operations.
- `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM` specifies that image views created with this format **can** be sampled in [box filter sampling](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-boxfilter) operations.
- `VK_FORMAT_FEATURE_2_HOST_IMAGE_TRANSFER_BIT` specifies that an image **can** be created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT`.
- `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` specifies that an image view with this format **can** be used as a [quantization delta map](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-delta-map) in [video encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations).
- `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR` specifies that an image view with this format **can** be used as an [emphasis map](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-emphasis-map) in [video encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations).
  
  Note
  
  [Video encode quantization maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map) have additional restrictions specific to the [video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles) they are used with that **can** be enumerated using the [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html) command.

The following bits **may** be set in `bufferFeatures`, specifying that the features are supported by [buffers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) or [buffer views](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) created with the queried [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2.html)::`format`:

- `VK_FORMAT_FEATURE_2_UNIFORM_TEXEL_BUFFER_BIT` specifies that the format **can** be used to create a buffer view that **can** be bound to a `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` descriptor.
- `VK_FORMAT_FEATURE_2_STORAGE_TEXEL_BUFFER_BIT` specifies that the format **can** be used to create a buffer view that **can** be bound to a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor.
- `VK_FORMAT_FEATURE_2_STORAGE_TEXEL_BUFFER_ATOMIC_BIT` specifies that atomic operations are supported on `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` with this format.
- `VK_FORMAT_FEATURE_2_VERTEX_BUFFER_BIT` specifies that the format **can** be used as a vertex attribute format (`VkVertexInputAttributeDescription`::`format`).
- `VK_FORMAT_FEATURE_2_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR` specifies that the format **can** be used as the vertex format when creating an [acceleration structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure) (`VkAccelerationStructureGeometryTrianglesDataKHR`::`vertexFormat`). This format **can** also be used as the vertex format in host memory when doing [host acceleration structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#host-acceleration-structure) builds.
- `VK_FORMAT_FEATURE_2_ACCELERATION_STRUCTURE_RADIUS_BUFFER_BIT_NV` specifies that the format **can** be used as the radius format when creating an [acceleration structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure) (`VkAccelerationStructureGeometryLinearSweptSpheresDataNV`::`radiusFormat` or `VkAccelerationStructureGeometrySpheresDataNV`::`radiusFormat`). This format **can** also be used as the radius format in host memory when doing [host acceleration structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#host-acceleration-structure) builds.
- `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT` specifies that buffer views created with this format **can** be used as [storage texel buffers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer) for read operations without specifying a format.
- `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT` specifies that buffer views created with this format **can** be used as [storage texel buffers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer) for write operations without specifying a format.
- `VK_FORMAT_FEATURE_2_OPTICAL_FLOW_IMAGE_BIT_NV` specifies that an image view with this format **can** be used as an input or reference to [optical flow operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#opticalflow-operations)
- `VK_FORMAT_FEATURE_2_OPTICAL_FLOW_VECTOR_BIT_NV` specifies that an image view with this format **can** be used as a flow vector map (either as hint, output or global flow) for [optical flow operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#opticalflow-operations)
- `VK_FORMAT_FEATURE_2_OPTICAL_FLOW_COST_BIT_NV` specifies that an image view with this format **can** be used as an output cost map for [optical flow operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#opticalflow-operations)

## [](#_see_also)See Also

[VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFormatFeatureFlagBits2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
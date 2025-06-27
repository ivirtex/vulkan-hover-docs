# VkFormatFeatureFlagBits(3) Manual Page

## Name

VkFormatFeatureFlagBits - Bitmask specifying features supported by a
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in the
[VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html) features
`linearTilingFeatures`, `optimalTilingFeatures`,
[VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierTilingFeatures`,
and `bufferFeatures` are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkFormatFeatureFlagBits {
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT = 0x00000001,
    VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT = 0x00000002,
    VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT = 0x00000004,
    VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT = 0x00000008,
    VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT = 0x00000010,
    VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT = 0x00000020,
    VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT = 0x00000040,
    VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT = 0x00000080,
    VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT = 0x00000100,
    VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT = 0x00000200,
    VK_FORMAT_FEATURE_BLIT_SRC_BIT = 0x00000400,
    VK_FORMAT_FEATURE_BLIT_DST_BIT = 0x00000800,
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT = 0x00001000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_TRANSFER_SRC_BIT = 0x00004000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_TRANSFER_DST_BIT = 0x00008000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT = 0x00020000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT = 0x00040000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT = 0x00080000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT = 0x00100000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT = 0x00200000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_DISJOINT_BIT = 0x00400000,
  // Provided by VK_VERSION_1_1
    VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT = 0x00800000,
  // Provided by VK_VERSION_1_2
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT = 0x00010000,
  // Provided by VK_KHR_video_decode_queue
    VK_FORMAT_FEATURE_VIDEO_DECODE_OUTPUT_BIT_KHR = 0x02000000,
  // Provided by VK_KHR_video_decode_queue
    VK_FORMAT_FEATURE_VIDEO_DECODE_DPB_BIT_KHR = 0x04000000,
  // Provided by VK_KHR_acceleration_structure
    VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR = 0x20000000,
  // Provided by VK_EXT_filter_cubic
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT = 0x00002000,
  // Provided by VK_EXT_fragment_density_map
    VK_FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT = 0x01000000,
  // Provided by VK_KHR_fragment_shading_rate
    VK_FORMAT_FEATURE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR = 0x40000000,
  // Provided by VK_KHR_video_encode_queue
    VK_FORMAT_FEATURE_VIDEO_ENCODE_INPUT_BIT_KHR = 0x08000000,
  // Provided by VK_KHR_video_encode_queue
    VK_FORMAT_FEATURE_VIDEO_ENCODE_DPB_BIT_KHR = 0x10000000,
  // Provided by VK_IMG_filter_cubic
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG = VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT,
  // Provided by VK_KHR_maintenance1
    VK_FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR = VK_FORMAT_FEATURE_TRANSFER_SRC_BIT,
  // Provided by VK_KHR_maintenance1
    VK_FORMAT_FEATURE_TRANSFER_DST_BIT_KHR = VK_FORMAT_FEATURE_TRANSFER_DST_BIT,
  // Provided by VK_EXT_sampler_filter_minmax
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT = VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT_KHR = VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR = VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR = VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR = VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR = VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_FORMAT_FEATURE_DISJOINT_BIT_KHR = VK_FORMAT_FEATURE_DISJOINT_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT_KHR = VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT,
} VkFormatFeatureFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

These values all have the same meaning as the equivalently named values
for [VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags2.html) and **may** be
set in `linearTilingFeatures`, `optimalTilingFeatures`, and
[VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierTilingFeatures`,
specifying that the features are supported by
<a href="VkImage.html" target="_blank" rel="noopener">images</a> or
<a href="VkImageView.html" target="_blank" rel="noopener">image
views</a> or <a href="VkSamplerYcbcrConversion.html" target="_blank"
rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub> conversion
objects</a> created with the queried
[vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)::`format`:

- `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT` specifies that an image view
  **can** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampledimage"
  target="_blank" rel="noopener">sampled from</a>.

- `VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT` specifies that an image view
  **can** be used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage"
  target="_blank" rel="noopener">storage image</a>.

- `VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT` specifies that an image
  view **can** be used as storage image that supports atomic operations.

- `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` specifies that an image view
  **can** be used as a framebuffer color attachment and as an input
  attachment.

- `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT` specifies that an image
  view **can** be used as a framebuffer color attachment that supports
  blending.

- `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT` specifies that an
  image view **can** be used as a framebuffer depth/stencil attachment
  and as an input attachment.

- `VK_FORMAT_FEATURE_BLIT_SRC_BIT` specifies that an image **can** be
  used as `srcImage` for the `vkCmdBlitImage2` and `vkCmdBlitImage`
  commands.

- `VK_FORMAT_FEATURE_BLIT_DST_BIT` specifies that an image **can** be
  used as `dstImage` for the `vkCmdBlitImage2` and `vkCmdBlitImage`
  commands.

- `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT` specifies that if
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT` is also set, an image view
  **can** be used with a sampler that has either of `magFilter` or
  `minFilter` set to `VK_FILTER_LINEAR`, or `mipmapMode` set to
  `VK_SAMPLER_MIPMAP_MODE_LINEAR`. If `VK_FORMAT_FEATURE_BLIT_SRC_BIT`
  is also set, an image can be used as the `srcImage` to
  `vkCmdBlitImage2` and `vkCmdBlitImage` with a `filter` of
  `VK_FILTER_LINEAR`. This bit **must** only be exposed for formats that
  also support the `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT` or
  `VK_FORMAT_FEATURE_BLIT_SRC_BIT`.

  If the format being queried is a depth/stencil format, this bit only
  specifies that the depth aspect (not the stencil aspect) of an image
  of this format supports linear filtering, and that linear filtering of
  the depth aspect is supported whether depth compare is enabled in the
  sampler or not. Where depth comparison is supported it **may** be
  linear filtered whether this bit is present or not, but where this bit
  is not present the filtered value **may** be computed in an
  implementation-dependent manner which differs from the normal rules of
  linear filtering. The resulting value **must** be in the range \[0,1\]
  and **should** be proportional to, or a weighted average of, the
  number of comparison passes or failures.

- `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT` specifies that an image **can**
  be used as a source image for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies"
  target="_blank" rel="noopener">copy commands</a>. If the application
  `apiVersion` is Vulkan 1.0 and
  [`VK_KHR_maintenance1`](VK_KHR_maintenance1.html) is not supported,
  `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT` is implied to be set when the
  format feature flag is not 0.

- `VK_FORMAT_FEATURE_TRANSFER_DST_BIT` specifies that an image **can**
  be used as a destination image for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies"
  target="_blank" rel="noopener">copy commands</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears"
  target="_blank" rel="noopener">clear commands</a>. If the application
  `apiVersion` is Vulkan 1.0 and
  [`VK_KHR_maintenance1`](VK_KHR_maintenance1.html) is not supported,
  `VK_FORMAT_FEATURE_TRANSFER_DST_BIT` is implied to be set when the
  format feature flag is not 0.

- `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT` specifies
  `VkImage` **can** be used as a sampled image with a min or max
  [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html). This bit
  **must** only be exposed for formats that also support the
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`.

- `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT` specifies that
  `VkImage` **can** be used with a sampler that has either of
  `magFilter` or `minFilter` set to `VK_FILTER_CUBIC_EXT`, or be the
  source image for a blit with `filter` set to `VK_FILTER_CUBIC_EXT`.
  This bit **must** only be exposed for formats that also support the
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`. If the format being queried is
  a depth/stencil format, this only specifies that the depth aspect is
  cubic filterable.

- `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT` specifies that an
  application **can** define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> using this format as a source, and that an image of
  this format **can** be used with a
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
  `xChromaOffset` and/or `yChromaOffset` of
  `VK_CHROMA_LOCATION_MIDPOINT`. Otherwise both `xChromaOffset` and
  `yChromaOffset` **must** be `VK_CHROMA_LOCATION_COSITED_EVEN`. If a
  format does not incorporate chroma downsampling (it is not a “422” or
  “420” format) but the implementation supports sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion for this format, the
  implementation **must** set
  `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT`.

- `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT` specifies that an
  application **can** define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> using this format as a source, and that an image of
  this format **can** be used with a
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
  `xChromaOffset` and/or `yChromaOffset` of
  `VK_CHROMA_LOCATION_COSITED_EVEN`. Otherwise both `xChromaOffset` and
  `yChromaOffset` **must** be `VK_CHROMA_LOCATION_MIDPOINT`. If neither
  `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT` nor
  `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT` is set, the
  application **must** not define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> using this format as a source.

- `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT`
  specifies that an application **can** define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> using this format as a source with `chromaFilter` set
  to `VK_FILTER_LINEAR`.

- `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT`
  specifies that the format can have different chroma, min, and mag
  filters.

- `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT`
  specifies that reconstruction is explicit, as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-chroma-reconstruction"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-chroma-reconstruction</a>.
  If this bit is not present, reconstruction is implicit by default.

- `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT`
  specifies that reconstruction **can** be forcibly made explicit by
  setting
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)::`forceExplicitReconstruction`
  to `VK_TRUE`. If the format being queried supports
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT`
  it **must** also support
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT`.

- `VK_FORMAT_FEATURE_DISJOINT_BIT` specifies that a multi-planar image
  **can** have the `VK_IMAGE_CREATE_DISJOINT_BIT` set during image
  creation. An implementation **must** not set
  `VK_FORMAT_FEATURE_DISJOINT_BIT` for *single-plane formats*.

- `VK_FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT` specifies that an
  image view **can** be used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-fragmentdensitymapattachment"
  target="_blank" rel="noopener">fragment density map attachment</a>.

- `VK_FORMAT_FEATURE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` specifies
  that an image view **can** be used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a>.
  An implementation **must** not set this feature for formats with a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-numericformat"
  target="_blank" rel="noopener">numeric format</a> other than `UINT`,
  or set it as a buffer feature.

- `VK_FORMAT_FEATURE_VIDEO_DECODE_OUTPUT_BIT_KHR` specifies that an
  image view with this format **can** be used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture"
  target="_blank" rel="noopener">decode output picture</a> in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operations</a>.

- `VK_FORMAT_FEATURE_VIDEO_DECODE_DPB_BIT_KHR` specifies that an image
  view with this format **can** be used as an output <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
  target="_blank" rel="noopener">reconstructed picture</a> or an input
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
  target="_blank" rel="noopener">reference picture</a> in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operations</a>.

- `VK_FORMAT_FEATURE_VIDEO_ENCODE_INPUT_BIT_KHR` specifies that an image
  view with this format **can** be used as an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
  target="_blank" rel="noopener">encode input picture</a> in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operations</a>.

- `VK_FORMAT_FEATURE_VIDEO_ENCODE_DPB_BIT_KHR` specifies that an image
  view with this format **can** be used as an output <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
  target="_blank" rel="noopener">reconstructed picture</a> or an input
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
  target="_blank" rel="noopener">reference picture</a> in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operations</a>.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>Specific <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
  target="_blank" rel="noopener">video profiles</a> <strong>may</strong>
  have additional restrictions on the format and other image creation
  parameters corresponding to image views used by video coding operations
  that <strong>can</strong> be enumerated using the <a
  href="vkGetPhysicalDeviceVideoFormatPropertiesKHR.html">vkGetPhysicalDeviceVideoFormatPropertiesKHR</a>
  command.</p></td>
  </tr>
  </tbody>
  </table>

The following bits **may** be set in `bufferFeatures`, specifying that
the features are supported by
<a href="VkBuffer.html" target="_blank" rel="noopener">buffers</a> or
<a href="VkBufferView.html" target="_blank" rel="noopener">buffer
views</a> created with the queried
[vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)::`format`:

- `VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT` specifies that the format
  **can** be used to create a buffer view that **can** be bound to a
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` descriptor.

- `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT` specifies that the format
  **can** be used to create a buffer view that **can** be bound to a
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor.

- `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT` specifies that
  atomic operations are supported on
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` with this format.

- `VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT` specifies that the format
  **can** be used as a vertex attribute format
  (`VkVertexInputAttributeDescription`::`format`).

- `VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR`
  specifies that the format **can** be used as the vertex format when
  creating an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure"
  target="_blank" rel="noopener">acceleration structure</a>
  (`VkAccelerationStructureGeometryTrianglesDataKHR`::`vertexFormat`).
  This format **can** also be used as the vertex format in host memory
  when doing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#host-acceleration-structure"
  target="_blank" rel="noopener">host acceleration structure</a> builds.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT</code> and
<code>VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT</code> are only
intended to be advertised for single-component formats, since SPIR-V
atomic operations require a scalar type.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFormatFeatureFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

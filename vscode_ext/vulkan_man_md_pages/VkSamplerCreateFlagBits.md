# VkSamplerCreateFlagBits(3) Manual Page

## Name

VkSamplerCreateFlagBits - Bitmask specifying additional parameters of
sampler



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`flags`, specifying
additional parameters of a sampler, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkSamplerCreateFlagBits {
  // Provided by VK_EXT_fragment_density_map
    VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT = 0x00000001,
  // Provided by VK_EXT_fragment_density_map
    VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT = 0x00000002,
  // Provided by VK_EXT_descriptor_buffer
    VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT = 0x00000008,
  // Provided by VK_EXT_non_seamless_cube_map
    VK_SAMPLER_CREATE_NON_SEAMLESS_CUBE_MAP_BIT_EXT = 0x00000004,
  // Provided by VK_QCOM_image_processing
    VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM = 0x00000010,
} VkSamplerCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- <span id="samplers-subsamplesampler"></span>
  `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT` specifies that the sampler will
  read from an image created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`.

- `VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT` specifies
  that the implementation **may** use approximations when reconstructing
  a full color value for texture access from a subsampled image.

- `VK_SAMPLER_CREATE_NON_SEAMLESS_CUBE_MAP_BIT_EXT` specifies that <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-cubemapedge"
  target="_blank" rel="noopener">cube map edge handling</a> is not
  performed.

- <span id="samplers-imageprocessingsampler"></span>
  `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM` specifies that the
  sampler will read from images using only `OpImageWeightedSampleQCOM`,
  `OpImageBoxFilterQCOM`, `OpImageBlockMatchGatherSSDQCOM`,
  `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchWindowSSDQCOM`,
  `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchSSDQCOM`, or
  `OpImageBlockMatchSADQCOM`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The approximations used when
<code>VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT</code>
is specified are implementation defined. Some implementations
<strong>may</strong> interpolate between fragment density levels in a
subsampled image. In that case, this bit <strong>may</strong> be used to
decide whether the interpolation factors are calculated per fragment or
at a coarser granularity.</p></td>
</tr>
</tbody>
</table>

- `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` specifies
  that the sampler **can** be used with descriptor buffers when
  capturing and replaying (e.g. for trace capture and replay), see
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  for more detail.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkSamplerCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

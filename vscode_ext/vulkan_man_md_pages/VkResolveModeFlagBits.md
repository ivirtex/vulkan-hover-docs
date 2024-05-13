# VkResolveModeFlagBits(3) Manual Page

## Name

VkResolveModeFlagBits - Bitmask indicating supported depth and stencil
resolve modes



## <a href="#_c_specification" class="anchor"></a>C Specification

Multisample values in a multisample attachment are combined according to
the resolve mode used:

``` c
// Provided by VK_VERSION_1_2
typedef enum VkResolveModeFlagBits {
    VK_RESOLVE_MODE_NONE = 0,
    VK_RESOLVE_MODE_SAMPLE_ZERO_BIT = 0x00000001,
    VK_RESOLVE_MODE_AVERAGE_BIT = 0x00000002,
    VK_RESOLVE_MODE_MIN_BIT = 0x00000004,
    VK_RESOLVE_MODE_MAX_BIT = 0x00000008,
  // Provided by VK_ANDROID_external_format_resolve with VK_KHR_dynamic_rendering or VK_VERSION_1_3
    VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID = 0x00000010,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_NONE_KHR = VK_RESOLVE_MODE_NONE,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR = VK_RESOLVE_MODE_SAMPLE_ZERO_BIT,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_AVERAGE_BIT_KHR = VK_RESOLVE_MODE_AVERAGE_BIT,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_MIN_BIT_KHR = VK_RESOLVE_MODE_MIN_BIT,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_MAX_BIT_KHR = VK_RESOLVE_MODE_MAX_BIT,
} VkResolveModeFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_depth_stencil_resolve
typedef VkResolveModeFlagBits VkResolveModeFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_RESOLVE_MODE_NONE` indicates that no resolve operation is done.

- `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT` indicates that result of the resolve
  operation is equal to the value of sample 0.

- `VK_RESOLVE_MODE_AVERAGE_BIT` indicates that result of the resolve
  operation is the average of the sample values.

- `VK_RESOLVE_MODE_MIN_BIT` indicates that result of the resolve
  operation is the minimum of the sample values.

- `VK_RESOLVE_MODE_MAX_BIT` indicates that result of the resolve
  operation is the maximum of the sample values.

- `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID` indicates that
  rather than a multisample resolve, a single sampled color attachment
  will be downsampled into a Y′C<sub>B</sub>C<sub>R</sub> format image
  specified by an external Android format. Unlike other resolve modes,
  implementations can resolve multiple times during rendering, or even
  bypass writing to the color attachment altogether, as long as the
  final value is resolved to the resolve attachment. Values in the G, B,
  and R channels of the color attachment will be written to the Y,
  C<sub>B</sub>, and C<sub>R</sub> channels of the external format
  image, respectively. Chroma values are calculated as if sampling with
  a linear filter from the color attachment at full rate, at the
  location the chroma values sit according to
  [VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html)::`chromaOffsetX`,
  [VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html)::`chromaOffsetY`,
  and the chroma sample rate of the resolved image.

If no resolve mode is otherwise specified, `VK_RESOLVE_MODE_AVERAGE_BIT`
is used.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>No range compression or Y′C<sub>B</sub>C<sub>R</sub> model conversion
is performed by
<code>VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID</code>;
applications have to do these conversions themselves. Value outputs are
expected to match those that would be read through a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-sampler-YCbCr-conversion-modelconversion"
target="_blank" rel="noopener">Y′C<sub>B</sub>C<sub>R</sub> sampler
using <code>VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY</code></a>.
The color space that the values should be in is defined by the platform
and is not exposed via Vulkan.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_depth_stencil_resolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_depth_stencil_resolve.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html),
[VkResolveModeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlags.html),
[VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionDepthStencilResolve.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkResolveModeFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

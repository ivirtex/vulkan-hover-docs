# VkColorSpaceKHR(3) Manual Page

## Name

VkColorSpaceKHR - Supported color space of the presentation engine



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html)::`colorSpace`, specifying
the color spaces that a presentation engine can accept, are:

``` c
// Provided by VK_KHR_surface
typedef enum VkColorSpaceKHR {
    VK_COLOR_SPACE_SRGB_NONLINEAR_KHR = 0,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT = 1000104001,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT = 1000104002,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT = 1000104003,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_DCI_P3_NONLINEAR_EXT = 1000104004,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_BT709_LINEAR_EXT = 1000104005,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_BT709_NONLINEAR_EXT = 1000104006,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_BT2020_LINEAR_EXT = 1000104007,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_HDR10_ST2084_EXT = 1000104008,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_DOLBYVISION_EXT = 1000104009,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_HDR10_HLG_EXT = 1000104010,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_ADOBERGB_LINEAR_EXT = 1000104011,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_ADOBERGB_NONLINEAR_EXT = 1000104012,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_PASS_THROUGH_EXT = 1000104013,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT = 1000104014,
  // Provided by VK_AMD_display_native_hdr
    VK_COLOR_SPACE_DISPLAY_NATIVE_AMD = 1000213000,
    VK_COLORSPACE_SRGB_NONLINEAR_KHR = VK_COLOR_SPACE_SRGB_NONLINEAR_KHR,
  // Provided by VK_EXT_swapchain_colorspace
    VK_COLOR_SPACE_DCI_P3_LINEAR_EXT = VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT,
} VkColorSpaceKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COLOR_SPACE_SRGB_NONLINEAR_KHR` specifies support for the images
  in sRGB color space, encoded according to the sRGB specification.

- `VK_COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT` specifies support for the
  images in Display-P3 color space, encoded using a Display-P3 transfer
  function.

- `VK_COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT` specifies support for the
  images in extended sRGB color space, encoded using a linear transfer
  function.

- `VK_COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT` specifies support for the
  images in extended sRGB color space, encoded according to the scRGB
  specification.

- `VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT` specifies support for the
  images in Display-P3 color space, encoded using a linear transfer
  function.

- `VK_COLOR_SPACE_DCI_P3_NONLINEAR_EXT` specifies support for the images
  in DCI-P3 color space, encoded according to the DCI-P3 specification.
  Note that values in such an image are interpreted as XYZ encoded color
  data by the presentation engine.

- `VK_COLOR_SPACE_BT709_LINEAR_EXT` specifies support for the images in
  BT709 color space, encoded using a linear transfer function.

- `VK_COLOR_SPACE_BT709_NONLINEAR_EXT` specifies support for the images
  in BT709 color space, encoded according to the BT709 specification.

- `VK_COLOR_SPACE_BT2020_LINEAR_EXT` specifies support for the images in
  BT2020 color space, encoded using a linear transfer function.

- `VK_COLOR_SPACE_HDR10_ST2084_EXT` specifies support for the images in
  HDR10 (BT2020) color space, encoded according to SMPTE ST2084
  Perceptual Quantizer (PQ) specification.

- `VK_COLOR_SPACE_DOLBYVISION_EXT` specifies support for the images in
  Dolby Vision (BT2020) color space, encoded according to SMPTE ST2084
  Perceptual Quantizer (PQ) specification. The presentation engine is
  expected to use Dolby’s proprietary techniques to display the image.

- `VK_COLOR_SPACE_HDR10_HLG_EXT` specifies support for the images in
  HDR10 (BT2020) color space, encoded according to the Hybrid Log Gamma
  (HLG) specification.

- `VK_COLOR_SPACE_ADOBERGB_LINEAR_EXT` specifies support for images in
  Adobe RGB color space, encoded using a linear transfer function.

- `VK_COLOR_SPACE_ADOBERGB_NONLINEAR_EXT` specifies support for the
  images in Adobe RGB color space, encoded according to the Adobe RGB
  specification (approximately Gamma 2.2).

- `VK_COLOR_SPACE_PASS_THROUGH_EXT` specifies that color components are
  used “as is”. This is intended to allow applications to supply data
  for color spaces not described here.

- `VK_COLOR_SPACE_DISPLAY_NATIVE_AMD` specifies support for the
  display’s native color space. This matches the color space
  expectations of AMD’s FreeSync2 standard, for displays supporting it.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>In the initial release of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html"><code>VK_KHR_surface</code></a> and <a
href="VK_KHR_swapchain.html"><code>VK_KHR_swapchain</code></a>
extensions, the token <code>VK_COLORSPACE_SRGB_NONLINEAR_KHR</code> was
used. Starting in the 2016-05-13 updates to the extension branches,
matching release 1.0.13 of the core API specification,
<code>VK_COLOR_SPACE_SRGB_NONLINEAR_KHR</code> is used instead for
consistency with Vulkan naming rules. The older enum is still available
for backwards compatibility.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>In older versions of this extension
<code>VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT</code> was misnamed
<code>VK_COLOR_SPACE_DCI_P3_LINEAR_EXT</code>. This has been updated to
indicate that it uses RGB color encoding, not XYZ. The old name is
deprecated but is maintained for backwards compatibility.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For a traditional “Linear” or non-gamma transfer function color space
use <code>VK_COLOR_SPACE_PASS_THROUGH_EXT</code>.</p></td>
</tr>
</tbody>
</table>

The presentation engine interprets the pixel values of the R, G, and B
components as having been encoded using an appropriate transfer
function. Applications **should** ensure that the appropriate transfer
function has been applied. <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-output-format-conversion"
target="_blank" rel="noopener">Textures Output Format Conversion</a>
requires that all implementations implicitly apply the sRGB
EOTF<sup>-1</sup> on R, G, and B components when shaders write to an
sRGB pixel format image, which is useful for sRGB color spaces. For sRGB
color spaces with other pixel formats, or other non-linear color spaces,
applications **can** apply the transfer function explicitly in a shader.
The A channel is always interpreted as linearly encoded.

This extension defines enums for [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html)
that correspond to the following color spaces:

| Name | Red Primary | Green Primary | Blue Primary | White-point | Transfer function |
|----|----|----|----|----|----|
| DCI-P3 | 1.000, 0.000 | 0.000, 1.000 | 0.000, 0.000 | 0.3333, 0.3333 | DCI P3 |
| Display-P3 | 0.680, 0.320 | 0.265, 0.690 | 0.150, 0.060 | 0.3127, 0.3290 (D65) | Display-P3 |
| BT709 | 0.640, 0.330 | 0.300, 0.600 | 0.150, 0.060 | 0.3127, 0.3290 (D65) | BT709 |
| sRGB | 0.640, 0.330 | 0.300, 0.600 | 0.150, 0.060 | 0.3127, 0.3290 (D65) | sRGB |
| extended sRGB | 0.640, 0.330 | 0.300, 0.600 | 0.150, 0.060 | 0.3127, 0.3290 (D65) | scRGB |
| HDR10_ST2084 | 0.708, 0.292 | 0.170, 0.797 | 0.131, 0.046 | 0.3127, 0.3290 (D65) | ST2084 PQ |
| DOLBYVISION | 0.708, 0.292 | 0.170, 0.797 | 0.131, 0.046 | 0.3127, 0.3290 (D65) | ST2084 PQ |
| HDR10_HLG | 0.708, 0.292 | 0.170, 0.797 | 0.131, 0.046 | 0.3127, 0.3290 (D65) | HLG |
| Adobe RGB | 0.640, 0.330 | 0.210, 0.710 | 0.150, 0.060 | 0.3127, 0.3290 (D65) | Adobe RGB |

Table 1. Color Spaces and Attributes
{#VK_EXT_swapchain_colorspace-table}

The transfer functions are described in the “Transfer Functions” chapter
of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#data-format"
target="_blank" rel="noopener">Khronos Data Format Specification</a>.

Except Display-P3 OETF, which is:

E​={1.055×L2.41​−0.05512.92×L​for 0.0030186≤L≤1for 0≤L\<0.0030186​​

where L is the linear value of a color component and E is the encoded
value (as stored in the image in memory).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For most uses, the sRGB OETF is equivalent.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkColorSpaceKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

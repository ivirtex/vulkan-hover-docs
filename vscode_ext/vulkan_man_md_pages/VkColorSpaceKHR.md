# VkColorSpaceKHR(3) Manual Page

## Name

VkColorSpaceKHR - Supported color space of the presentation engine



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html)::`colorSpace`, specifying
supported color spaces of a presentation engine, are:

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

- `VK_COLOR_SPACE_SRGB_NONLINEAR_KHR` specifies support for the sRGB
  color space.

- `VK_COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT` specifies support for the
  Display-P3 color space to be displayed using an sRGB-like EOTF
  (defined below).

- `VK_COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT` specifies support for the
  extended sRGB color space to be displayed using a linear EOTF.

- `VK_COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT` specifies support for the
  extended sRGB color space to be displayed using an sRGB EOTF.

- `VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT` specifies support for the
  Display-P3 color space to be displayed using a linear EOTF.

- `VK_COLOR_SPACE_DCI_P3_NONLINEAR_EXT` specifies support for the DCI-P3
  color space to be displayed using the DCI-P3 EOTF. Note that values in
  such an image are interpreted as XYZ encoded color data by the
  presentation engine.

- `VK_COLOR_SPACE_BT709_LINEAR_EXT` specifies support for the BT709
  color space to be displayed using a linear EOTF.

- `VK_COLOR_SPACE_BT709_NONLINEAR_EXT` specifies support for the BT709
  color space to be displayed using the SMPTE 170M EOTF.

- `VK_COLOR_SPACE_BT2020_LINEAR_EXT` specifies support for the BT2020
  color space to be displayed using a linear EOTF.

- `VK_COLOR_SPACE_HDR10_ST2084_EXT` specifies support for the HDR10
  (BT2020 color) space to be displayed using the SMPTE ST2084 Perceptual
  Quantizer (PQ) EOTF.

- `VK_COLOR_SPACE_DOLBYVISION_EXT` specifies support for the Dolby
  Vision (BT2020 color space), proprietary encoding, to be displayed
  using the SMPTE ST2084 EOTF.

- `VK_COLOR_SPACE_HDR10_HLG_EXT` specifies support for the HDR10 (BT2020
  color space) to be displayed using the Hybrid Log Gamma (HLG) EOTF.

- `VK_COLOR_SPACE_ADOBERGB_LINEAR_EXT` specifies support for the
  AdobeRGB color space to be displayed using a linear EOTF.

- `VK_COLOR_SPACE_ADOBERGB_NONLINEAR_EXT` specifies support for the
  AdobeRGB color space to be displayed using the Gamma 2.2 EOTF.

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
<tr class="odd">
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
<tr class="odd">
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
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>For a traditional “Linear” or non-gamma transfer function color space
use <code>VK_COLOR_SPACE_PASS_THROUGH_EXT</code>.</p></td>
</tr>
</tbody>
</table>

The color components of non-linear color space swapchain images **must**
have had the appropriate transfer function applied. The color space
selected for the swapchain image will not affect the processing of data
written into the image by the implementation. Vulkan requires that all
implementations support the sRGB transfer function by use of an SRGB
pixel format. Other transfer functions, such as SMPTE 170M or SMPTE2084,
**can** be performed by the application shader. This extension defines
enums for [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html) that correspond to the
following color spaces:

| Name          | Red Primary  | Green Primary | Blue Primary | White-point          | Transfer function |
|---------------|--------------|---------------|--------------|----------------------|-------------------|
| DCI-P3        | 1.000, 0.000 | 0.000, 1.000  | 0.000, 0.000 | 0.3333, 0.3333       | DCI P3            |
| Display-P3    | 0.680, 0.320 | 0.265, 0.690  | 0.150, 0.060 | 0.3127, 0.3290 (D65) | Display-P3        |
| BT709         | 0.640, 0.330 | 0.300, 0.600  | 0.150, 0.060 | 0.3127, 0.3290 (D65) | ITU (SMPTE 170M)  |
| sRGB          | 0.640, 0.330 | 0.300, 0.600  | 0.150, 0.060 | 0.3127, 0.3290 (D65) | sRGB              |
| extended sRGB | 0.640, 0.330 | 0.300, 0.600  | 0.150, 0.060 | 0.3127, 0.3290 (D65) | extended sRGB     |
| HDR10_ST2084  | 0.708, 0.292 | 0.170, 0.797  | 0.131, 0.046 | 0.3127, 0.3290 (D65) | ST2084 PQ         |
| DOLBYVISION   | 0.708, 0.292 | 0.170, 0.797  | 0.131, 0.046 | 0.3127, 0.3290 (D65) | ST2084 PQ         |
| HDR10_HLG     | 0.708, 0.292 | 0.170, 0.797  | 0.131, 0.046 | 0.3127, 0.3290 (D65) | HLG               |
| AdobeRGB      | 0.640, 0.330 | 0.210, 0.710  | 0.150, 0.060 | 0.3127, 0.3290 (D65) | AdobeRGB          |

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
<tr class="odd">
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

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

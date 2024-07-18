# VK_EXT_swapchain_colorspace(3) Manual Page

## Name

VK_EXT_swapchain_colorspace - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

105

## <a href="#_revision" class="anchor"></a>Revision

5

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Courtney Goeltzenleuchter <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_swapchain_colorspace%5D%20@courtney-g%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_swapchain_colorspace%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>courtney-g</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-04-26

**IP Status**  
No known IP claims.

**Contributors**  
- Courtney Goeltzenleuchter, Google

## <a href="#_description" class="anchor"></a>Description

This extension expands [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html) to add
support for most standard color spaces beyond
`VK_COLOR_SPACE_SRGB_NONLINEAR_KHR`. This extension also adds support
for `VK_COLOR_SPACE_PASS_THROUGH_EXT` which allows applications to use
color spaces not explicitly enumerated in
[VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html).

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SWAPCHAIN_COLOR_SPACE_EXTENSION_NAME`

- `VK_EXT_SWAPCHAIN_COLOR_SPACE_SPEC_VERSION`

- Extending [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html):

  - `VK_COLOR_SPACE_ADOBERGB_LINEAR_EXT`

  - `VK_COLOR_SPACE_ADOBERGB_NONLINEAR_EXT`

  - `VK_COLOR_SPACE_BT2020_LINEAR_EXT`

  - `VK_COLOR_SPACE_BT709_LINEAR_EXT`

  - `VK_COLOR_SPACE_BT709_NONLINEAR_EXT`

  - `VK_COLOR_SPACE_DCI_P3_LINEAR_EXT`

  - `VK_COLOR_SPACE_DCI_P3_NONLINEAR_EXT`

  - `VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT`

  - `VK_COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT`

  - `VK_COLOR_SPACE_DOLBYVISION_EXT`

  - `VK_COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT`

  - `VK_COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT`

  - `VK_COLOR_SPACE_HDR10_HLG_EXT`

  - `VK_COLOR_SPACE_HDR10_ST2084_EXT`

  - `VK_COLOR_SPACE_PASS_THROUGH_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does the spec need to specify which kinds of image formats support
the color spaces?

**RESOLVED**: Pixel format is independent of color space (though some
color spaces really want / need floating-point color components to be
useful). Therefore, do not plan on documenting what formats support
which color spaces. An application **can** call
[vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html)
to query what a particular implementation supports.

2\) How does application determine if HW supports appropriate transfer
function for a color space?

**RESOLVED**: Extension indicates that implementation **must** not do
the OETF encoding if it is not sRGB. That responsibility falls to the
application shaders. Any other native OETF / EOTF functions supported by
an implementation can be described by separate extension.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-12-27 (Courtney Goeltzenleuchter)

  - Initial version

- Revision 2, 2017-01-19 (Courtney Goeltzenleuchter)

  - Add pass through and multiple options for BT2020.

  - Clean up some issues with equations not displaying properly.

- Revision 3, 2017-06-23 (Courtney Goeltzenleuchter)

  - Add extended sRGB non-linear enum.

- Revision 4, 2019-04-26 (Graeme Leese)

  - Clarify color space transfer function usage.

  - Refer to normative definitions in the Data Format Specification.

  - Clarify DCI-P3 and Display P3 usage.

- Revision 5, 2024-03-16 (Zehui Lin)

  - Fix interchanged concepts of EOTF and OETF.

  - Clarify that the presentation engine can accept the color spaces.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_swapchain_colorspace"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

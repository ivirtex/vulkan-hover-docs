# VK\_EXT\_swapchain\_colorspace(3) Manual Page

## Name

VK\_EXT\_swapchain\_colorspace - instance extension



## [](#_registered_extension_number)Registered Extension Number

105

## [](#_revision)Revision

5

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Courtney Goeltzenleuchter [\[GitHub\]courtney-g](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_swapchain_colorspace%5D%20%40courtney-g%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_swapchain_colorspace%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-04-26

**IP Status**

No known IP claims.

**Contributors**

- Courtney Goeltzenleuchter, Google

## [](#_description)Description

This extension expands [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html) to add support for most standard color spaces beyond `VK_COLOR_SPACE_SRGB_NONLINEAR_KHR`. This extension also adds support for `VK_COLOR_SPACE_PASS_THROUGH_EXT` which allows applications to use color spaces not explicitly enumerated in [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html).

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SWAPCHAIN_COLOR_SPACE_EXTENSION_NAME`
- `VK_EXT_SWAPCHAIN_COLOR_SPACE_SPEC_VERSION`
- Extending [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html):
  
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

## [](#_issues)Issues

1\) Does the spec need to specify which kinds of image formats support the color spaces?

**RESOLVED**: Pixel format is independent of color space (though some color spaces really want / need floating-point color components to be useful). Therefore, do not plan on documenting what formats support which color spaces. An application **can** call [vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html) to query what a particular implementation supports.

2\) How does application determine if HW supports appropriate transfer function for a color space?

**RESOLVED**: Extension indicates that implementation **must** not do the OETF encoding if it is not sRGB. That responsibility falls to the application shaders. Any other native OETF / EOTF functions supported by an implementation can be described by separate extension.

## [](#_version_history)Version History

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

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_swapchain_colorspace).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
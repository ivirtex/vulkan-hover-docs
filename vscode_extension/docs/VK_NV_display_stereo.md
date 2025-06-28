# VK\_NV\_display\_stereo(3) Manual Page

## Name

VK\_NV\_display\_stereo - instance extension



## [](#_registered_extension_number)Registered Extension Number

552

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html)  
and  
[VK\_KHR\_get\_display\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_display_properties2.html)

## [](#_contact)Contact

- Russell Chou [\[GitHub\]russellcnv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_display_stereo%5D%20%40russellcnv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_display_stereo%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_NV\_display\_stereo](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_display_stereo.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-11-20

**Contributors**

- Russell Chou, NVIDIA
- Jeff Juliano, NVIDIA
- James Jones, NVIDIA

## [](#_description)Description

This extension allows the application to choose which type of 3D stereo hardware it wants to use so the driver can configure it properly. This configuration is useful for swapchains created from display surfaces because some environments do not have an intermediate windowing system available for easy configuration. This extension will override any stereo type configuration in the windowing system.

For HDMI 3D, only some display modes support stereo rendering, and a new structure is needed to expose that information to the application.

## [](#_new_structures)New Structures

- Extending [VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeProperties2KHR.html):
  
  - [VkDisplayModeStereoPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeStereoPropertiesNV.html)
- Extending [VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceCreateInfoKHR.html):
  
  - [VkDisplaySurfaceStereoCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceStereoCreateInfoNV.html)

## [](#_new_enums)New Enums

- [VkDisplaySurfaceStereoTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceStereoTypeNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_DISPLAY_STEREO_EXTENSION_NAME`
- `VK_NV_DISPLAY_STEREO_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DISPLAY_MODE_STEREO_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_DISPLAY_SURFACE_STEREO_CREATE_INFO_NV`

## [](#_version_history)Version History

- Revision 1, 2024-11-20 (Russell Chou)
  
  - Initial release

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_display_stereo)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
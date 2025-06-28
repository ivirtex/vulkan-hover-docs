# VK\_EXT\_direct\_mode\_display(3) Manual Page

## Name

VK\_EXT\_direct\_mode\_display - instance extension



## [](#_registered_extension_number)Registered Extension Number

89

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_direct_mode_display%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_direct_mode_display%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-12-13

**IP Status**

No known IP claims.

**Contributors**

- Pierre Boudier, NVIDIA
- James Jones, NVIDIA
- Damien Leone, NVIDIA
- Pierre-Loup Griffais, Valve
- Liam Middlebrook, NVIDIA

## [](#_description)Description

This is extension, along with related platform extensions, allows applications to take exclusive control of displays associated with a native windowing system. This is especially useful for virtual reality applications that wish to hide HMDs (head mounted displays) from the native platform’s display management system, desktop, and/or other applications.

## [](#_new_commands)New Commands

- [vkReleaseDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseDisplayEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DIRECT_MODE_DISPLAY_EXTENSION_NAME`
- `VK_EXT_DIRECT_MODE_DISPLAY_SPEC_VERSION`

## [](#_issues)Issues

1\) Should this extension and its related platform-specific extensions leverage `VK_KHR_display`, or provide separate equivalent interfaces.

**RESOLVED**: Use `VK_KHR_display` concepts and objects. `VK_KHR_display` can be used to enumerate all displays on the system, including those attached to/in use by a window system or native platform, but `VK_KHR_display_swapchain` will fail to create a swapchain on in-use displays. This extension and its platform-specific children will allow applications to grab in-use displays away from window systems and/or native platforms, allowing them to be used with `VK_KHR_display_swapchain`.

2\) Are separate calls needed to acquire displays and enable direct mode?

**RESOLVED**: No, these operations happen in one combined command. Acquiring a display puts it into direct mode.

## [](#_version_history)Version History

- Revision 1, 2016-12-13 (James Jones)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_direct_mode_display)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
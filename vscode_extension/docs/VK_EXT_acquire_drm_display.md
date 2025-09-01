# VK\_EXT\_acquire\_drm\_display(3) Manual Page

## Name

VK\_EXT\_acquire\_drm\_display - instance extension



## [](#_registered_extension_number)Registered Extension Number

286

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_direct\_mode\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_direct_mode_display.html)

## [](#_contact)Contact

- Drew DeVault [sir@cmpwn.com](mailto:sir@cmpwn.com)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-06-09

**IP Status**

No known IP claims.

**Contributors**

- Simon Zeni, Status Holdings, Ltd.

## [](#_description)Description

This extension allows an application to take exclusive control of a display using the Direct Rendering Manager (DRM) interface. When acquired, the display will be under full control of the application until the display is either released or the connector is unplugged.

## [](#_new_commands)New Commands

- [vkAcquireDrmDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireDrmDisplayEXT.html)
- [vkGetDrmDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDrmDisplayEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_ACQUIRE_DRM_DISPLAY_EXTENSION_NAME`
- `VK_EXT_ACQUIRE_DRM_DISPLAY_SPEC_VERSION`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2021-05-11 (Simon Zeni)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_acquire_drm_display).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
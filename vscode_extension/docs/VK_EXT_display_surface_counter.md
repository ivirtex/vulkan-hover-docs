# VK\_EXT\_display\_surface\_counter(3) Manual Page

## Name

VK\_EXT\_display\_surface\_counter - instance extension



## [](#_registered_extension_number)Registered Extension Number

91

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_display_surface_counter%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_display_surface_counter%20extension%2A)

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
- Daniel Vetter, Intel

## [](#_description)Description

This extension defines a vertical blanking period counter associated with display surfaces. It provides a mechanism to query support for such a counter from a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2EXT.html)

## [](#_new_structures)New Structures

- [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2EXT.html)

## [](#_new_enums)New Enums

- [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkSurfaceCounterFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DISPLAY_SURFACE_COUNTER_EXTENSION_NAME`
- `VK_EXT_DISPLAY_SURFACE_COUNTER_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES2_EXT`
  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT`

## [](#_version_history)Version History

- Revision 1, 2016-12-13 (James Jones)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_display_surface_counter).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
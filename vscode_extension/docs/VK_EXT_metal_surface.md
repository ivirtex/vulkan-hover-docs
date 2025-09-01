# VK\_EXT\_metal\_surface(3) Manual Page

## Name

VK\_EXT\_metal\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

218

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Dzmitry Malyshau [\[GitHub\]kvark](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_metal_surface%5D%20%40kvark%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_metal_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-10-01

**IP Status**

No known IP claims.

**Contributors**

- Dzmitry Malyshau, Mozilla Corp.

## [](#_description)Description

The `VK_EXT_metal_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) from [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html), which is the native rendering surface of Apple’s Metal framework.

## [](#_new_base_types)New Base Types

- [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html)

## [](#_new_commands)New Commands

- [vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMetalSurfaceEXT.html)

## [](#_new_structures)New Structures

- [VkMetalSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMetalSurfaceCreateInfoEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkMetalSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMetalSurfaceCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_METAL_SURFACE_EXTENSION_NAME`
- `VK_EXT_METAL_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2018-10-01 (Dzmitry Malyshau)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_metal_surface).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
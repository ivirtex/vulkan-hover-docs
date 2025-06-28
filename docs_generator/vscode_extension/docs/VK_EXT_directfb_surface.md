# VK\_EXT\_directfb\_surface(3) Manual Page

## Name

VK\_EXT\_directfb\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

347

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Nicolas Caramelli [\[GitHub\]caramelli](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_directfb_surface%5D%20%40caramelli%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_directfb_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-06-16

**IP Status**

No known IP claims.

**Contributors**

- Nicolas Caramelli

## [](#_description)Description

The `VK_EXT_directfb_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to a DirectFB `IDirectFBSurface`, as well as a query to determine support for rendering via DirectFB.

## [](#_new_commands)New Commands

- [vkCreateDirectFBSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDirectFBSurfaceEXT.html)
- [vkGetPhysicalDeviceDirectFBPresentationSupportEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDirectFBPresentationSupportEXT.html)

## [](#_new_structures)New Structures

- [VkDirectFBSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectFBSurfaceCreateInfoEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDirectFBSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectFBSurfaceCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DIRECTFB_SURFACE_EXTENSION_NAME`
- `VK_EXT_DIRECTFB_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DIRECTFB_SURFACE_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2020-06-16 (Nicolas Caramelli)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_directfb_surface)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
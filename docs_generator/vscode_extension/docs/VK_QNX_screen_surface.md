# VK\_QNX\_screen\_surface(3) Manual Page

## Name

VK\_QNX\_screen\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

379

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Mike Gorchak [\[GitHub\]mgorchak-blackberry](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QNX_screen_surface%5D%20%40mgorchak-blackberry%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QNX_screen_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-01-11

**IP Status**

No known IP claims.

**Contributors**

- Mike Gorchak, BlackBerry Limited

## [](#_description)Description

The `VK_QNX_screen_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to a QNX Screen `window`, as well as a query to determine support for rendering to a QNX Screen compositor.

## [](#_new_commands)New Commands

- [vkCreateScreenSurfaceQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateScreenSurfaceQNX.html)
- [vkGetPhysicalDeviceScreenPresentationSupportQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceScreenPresentationSupportQNX.html)

## [](#_new_structures)New Structures

- [VkScreenSurfaceCreateInfoQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenSurfaceCreateInfoQNX.html)

## [](#_new_bitmasks)New Bitmasks

- [VkScreenSurfaceCreateFlagsQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenSurfaceCreateFlagsQNX.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QNX_SCREEN_SURFACE_EXTENSION_NAME`
- `VK_QNX_SCREEN_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SCREEN_SURFACE_CREATE_INFO_QNX`

## [](#_version_history)Version History

- Revision 1, 2021-01-11 (Mike Gorchak)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QNX_screen_surface)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
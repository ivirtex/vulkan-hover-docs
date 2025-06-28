# VK\_EXT\_headless\_surface(3) Manual Page

## Name

VK\_EXT\_headless\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

257

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Lisa Wu [\[GitHub\]chengtianww](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_headless_surface%5D%20%40chengtianww%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_headless_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-03-21

**IP Status**

No known IP claims.

**Contributors**

- Ray Smith, Arm

## [](#_description)Description

The `VK_EXT_headless_surface` extension is an instance extension. It provides a mechanism to create [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) objects independently of any window system or display device. The presentation operation for a swapchain created from a headless surface is by default a no-op, resulting in no externally-visible result.

Because there is no real presentation target, future extensions can layer on top of the headless surface to introduce arbitrary or customizable sets of restrictions or features. These could include features like saving to a file or restrictions to emulate a particular presentation target.

This functionality is expected to be useful for application and driver development because it allows any platform to expose an arbitrary or customizable set of restrictions and features of a presentation engine. This makes it a useful portable test target for applications targeting a wide range of presentation engines where the actual target presentation engines might be scarce, unavailable or otherwise undesirable or inconvenient to use for general Vulkan application development.

## [](#_new_commands)New Commands

- [vkCreateHeadlessSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateHeadlessSurfaceEXT.html)

## [](#_new_structures)New Structures

- [VkHeadlessSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHeadlessSurfaceCreateInfoEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkHeadlessSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHeadlessSurfaceCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_HEADLESS_SURFACE_EXTENSION_NAME`
- `VK_EXT_HEADLESS_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2019-03-21 (Ray Smith)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_headless_surface)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VK\_KHR\_android\_surface(3) Manual Page

## Name

VK\_KHR\_android\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

9

## [](#_revision)Revision

6

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_android_surface%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_android_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-01-14

**IP Status**

No known IP claims.

**Contributors**

- Patrick Doane, Blizzard
- Faith Ekstrand, Intel
- Ian Elliott, LunarG
- Courtney Goeltzenleuchter, LunarG
- Jesse Hall, Google
- James Jones, NVIDIA
- Antoine Labour, Google
- Jon Leech, Khronos
- David Mao, AMD
- Norbert Nopper, Freescale
- Alon Or-bach, Samsung
- Daniel Rakos, AMD
- Graham Sellers, AMD
- Ray Smith, ARM
- Jeff Vigil, Qualcomm
- Chia-I Wu, LunarG

## [](#_description)Description

The `VK_KHR_android_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to an [ANativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/ANativeWindow.html), Android’s native surface type. The [ANativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/ANativeWindow.html) represents the producer endpoint of any buffer queue, regardless of consumer endpoint. Common consumer endpoints for `ANativeWindows` are the system window compositor, video encoders, and application-specific compositors importing the images through a `SurfaceTexture`.

## [](#_new_base_types)New Base Types

- [ANativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/ANativeWindow.html)

## [](#_new_commands)New Commands

- [vkCreateAndroidSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateAndroidSurfaceKHR.html)

## [](#_new_structures)New Structures

- [VkAndroidSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidSurfaceCreateInfoKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkAndroidSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidSurfaceCreateFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_ANDROID_SURFACE_EXTENSION_NAME`
- `VK_KHR_ANDROID_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR`

## [](#_issues)Issues

1\) Does Android need a way to query for compatibility between a particular physical device (and queue family?) and a specific Android display?

**RESOLVED**: No. Currently on Android, any physical device is expected to be able to present to the system compositor, and all queue families must support the necessary image layout transitions and synchronization operations.

## [](#_version_history)Version History

- Revision 1, 2015-09-23 (Jesse Hall)
  
  - Initial draft.
- Revision 2, 2015-10-26 (Ian Elliott)
  
  - Renamed from VK\_EXT\_KHR\_android\_surface to VK\_KHR\_android\_surface.
- Revision 3, 2015-11-03 (Daniel Rakos)
  
  - Added allocation callbacks to surface creation function.
- Revision 4, 2015-11-10 (Jesse Hall)
  
  - Removed VK\_ERROR\_INVALID\_ANDROID\_WINDOW\_KHR.
- Revision 5, 2015-11-28 (Daniel Rakos)
  
  - Updated the surface create function to take a pCreateInfo structure.
- Revision 6, 2016-01-14 (James Jones)
  
  - Moved VK\_ERROR\_NATIVE\_WINDOW\_IN\_USE\_KHR from the VK\_KHR\_android\_surface to the VK\_KHR\_surface extension.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_android_surface).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
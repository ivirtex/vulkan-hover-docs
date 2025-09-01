# VK\_KHR\_xcb\_surface(3) Manual Page

## Name

VK\_KHR\_xcb\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

6

## [](#_revision)Revision

6

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_xcb_surface%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_xcb_surface%20extension%2A)
- Ian Elliott [\[GitHub\]ianelliottus](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_xcb_surface%5D%20%40ianelliottus%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_xcb_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2015-11-28

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

The `VK_KHR_xcb_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to an X11 `Window`, using the XCB client-side library, as well as a query to determine support for rendering via XCB.

## [](#_new_commands)New Commands

- [vkCreateXcbSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateXcbSurfaceKHR.html)
- [vkGetPhysicalDeviceXcbPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceXcbPresentationSupportKHR.html)

## [](#_new_structures)New Structures

- [VkXcbSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkXcbSurfaceCreateInfoKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkXcbSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkXcbSurfaceCreateFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_XCB_SURFACE_EXTENSION_NAME`
- `VK_KHR_XCB_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR`

## [](#_issues)Issues

1\) Does XCB need a way to query for compatibility between a particular physical device and a specific screen? This would be a more general query than [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html): If it returned `VK_TRUE`, then the physical device could be assumed to support presentation to any window on that screen.

**RESOLVED**: Yes, this is needed for toolkits that want to create a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) before creating a window. To ensure the query is reliable, it must be made against a particular X visual rather than the screen in general.

## [](#_version_history)Version History

- Revision 1, 2015-09-23 (Jesse Hall)
  
  - Initial draft, based on the previous contents of VK\_EXT\_KHR\_swapchain (later renamed VK\_EXT\_KHR\_surface).
- Revision 2, 2015-10-02 (James Jones)
  
  - Added presentation support query for an (xcb\_connection\_t\*, xcb\_visualid\_t) pair.
  - Removed “root” parameter from CreateXcbSurfaceKHR(), as it is redundant when a window on the same screen is specified as well.
  - Adjusted wording of issue #1 and added agreed upon resolution.
- Revision 3, 2015-10-14 (Ian Elliott)
  
  - Removed “root” parameter from CreateXcbSurfaceKHR() in one more place.
- Revision 4, 2015-10-26 (Ian Elliott)
  
  - Renamed from VK\_EXT\_KHR\_xcb\_surface to VK\_KHR\_xcb\_surface.
- Revision 5, 2015-10-23 (Daniel Rakos)
  
  - Added allocation callbacks to vkCreateXcbSurfaceKHR.
- Revision 6, 2015-11-28 (Daniel Rakos)
  
  - Updated the surface create function to take a pCreateInfo structure.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_xcb_surface).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
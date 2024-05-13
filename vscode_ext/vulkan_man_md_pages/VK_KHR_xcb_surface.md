# VK_KHR_xcb_surface(3) Manual Page

## Name

VK_KHR_xcb_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

6

## <a href="#_revision" class="anchor"></a>Revision

6

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_xcb_surface%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_xcb_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>critsec</a>

- Ian Elliott <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_xcb_surface%5D%20@ianelliottus%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_xcb_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ianelliottus</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_xcb_surface` extension is an instance extension. It provides
a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) object
(defined by the [`VK_KHR_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) extension) that
refers to an X11 `Window`, using the XCB client-side library, as well as
a query to determine support for rendering via XCB.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateXcbSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateXcbSurfaceKHR.html)

- [vkGetPhysicalDeviceXcbPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceXcbPresentationSupportKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkXcbSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXcbSurfaceCreateInfoKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkXcbSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXcbSurfaceCreateFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_XCB_SURFACE_EXTENSION_NAME`

- `VK_KHR_XCB_SURFACE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does XCB need a way to query for compatibility between a particular
physical device and a specific screen? This would be a more general
query than
[vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html):
If it returned `VK_TRUE`, then the physical device could be assumed to
support presentation to any window on that screen.

**RESOLVED**: Yes, this is needed for toolkits that want to create a
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) before creating a window. To ensure the query
is reliable, it must be made against a particular X visual rather than
the screen in general.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2015-09-23 (Jesse Hall)

  - Initial draft, based on the previous contents of
    VK_EXT_KHR_swapchain (later renamed VK_EXT_KHR_surface).

- Revision 2, 2015-10-02 (James Jones)

  - Added presentation support query for an (xcb_connection_t\*,
    xcb_visualid_t) pair.

  - Removed “root” parameter from CreateXcbSurfaceKHR(), as it is
    redundant when a window on the same screen is specified as well.

  - Adjusted wording of issue \#1 and added agreed upon resolution.

- Revision 3, 2015-10-14 (Ian Elliott)

  - Removed “root” parameter from CreateXcbSurfaceKHR() in one more
    place.

- Revision 4, 2015-10-26 (Ian Elliott)

  - Renamed from VK_EXT_KHR_xcb_surface to VK_KHR_xcb_surface.

- Revision 5, 2015-10-23 (Daniel Rakos)

  - Added allocation callbacks to vkCreateXcbSurfaceKHR.

- Revision 6, 2015-11-28 (Daniel Rakos)

  - Updated the surface create function to take a pCreateInfo structure.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_xcb_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

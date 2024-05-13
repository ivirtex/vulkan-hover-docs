# VK_KHR_wayland_surface(3) Manual Page

## Name

VK_KHR_wayland_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

7

## <a href="#_revision" class="anchor"></a>Revision

6

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_wayland_surface%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_wayland_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>critsec</a>

- Ian Elliott <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_wayland_surface%5D%20@ianelliottus%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_wayland_surface%20extension*"
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

The `VK_KHR_wayland_surface` extension is an instance extension. It
provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)
object (defined by the [`VK_KHR_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)
extension) that refers to a Wayland `wl_surface`, as well as a query to
determine support for rendering to a Wayland compositor.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWaylandSurfaceKHR.html)

- [vkGetPhysicalDeviceWaylandPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceWaylandPresentationSupportKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkWaylandSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWaylandSurfaceCreateInfoKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkWaylandSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWaylandSurfaceCreateFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_WAYLAND_SURFACE_EXTENSION_NAME`

- `VK_KHR_WAYLAND_SURFACE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does Wayland need a way to query for compatibility between a
particular physical device and a specific Wayland display? This would be
a more general query than
[vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html):
if the Wayland-specific query returned `VK_TRUE` for a
([VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html), `struct wl_display*`) pair,
then the physical device could be assumed to support presentation to any
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) for surfaces on the display.

**RESOLVED**: Yes.
[vkGetPhysicalDeviceWaylandPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceWaylandPresentationSupportKHR.html)
was added to address this issue.

2\) Should we require surfaces created with
[vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWaylandSurfaceKHR.html) to support
the `VK_PRESENT_MODE_MAILBOX_KHR` present mode?

**RESOLVED**: Yes. Wayland is an inherently mailbox window system and
mailbox support is required for some Wayland compositor interactions to
work as expected. While handling these interactions may be possible with
`VK_PRESENT_MODE_FIFO_KHR`, it is much more difficult to do without
deadlock and requiring all Wayland applications to be able to support
implementations which only support `VK_PRESENT_MODE_FIFO_KHR` would be
an onerous restriction on application developers.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2015-09-23 (Jesse Hall)

  - Initial draft, based on the previous contents of
    VK_EXT_KHR_swapchain (later renamed VK_EXT_KHR_surface).

- Revision 2, 2015-10-02 (James Jones)

  - Added vkGetPhysicalDeviceWaylandPresentationSupportKHR() to resolve
    issue \#1.

  - Adjusted wording of issue \#1 to match the agreed-upon solution.

  - Renamed “window” parameters to “surface” to match Wayland
    conventions.

- Revision 3, 2015-10-26 (Ian Elliott)

  - Renamed from VK_EXT_KHR_wayland_surface to VK_KHR_wayland_surface.

- Revision 4, 2015-11-03 (Daniel Rakos)

  - Added allocation callbacks to vkCreateWaylandSurfaceKHR.

- Revision 5, 2015-11-28 (Daniel Rakos)

  - Updated the surface create function to take a pCreateInfo structure.

- Revision 6, 2017-02-08 (Faith Ekstrand)

  - Added the requirement that implementations support
    `VK_PRESENT_MODE_MAILBOX_KHR`.

  - Added wording about interactions between
    [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) and the Wayland requests
    sent to the compositor.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_wayland_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

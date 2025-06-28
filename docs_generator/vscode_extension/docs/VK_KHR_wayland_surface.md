# VK\_KHR\_wayland\_surface(3) Manual Page

## Name

VK\_KHR\_wayland\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

7

## [](#_revision)Revision

6

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_wayland_surface%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_wayland_surface%20extension%2A)
- Ian Elliott [\[GitHub\]ianelliottus](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_wayland_surface%5D%20%40ianelliottus%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_wayland_surface%20extension%2A)

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

The `VK_KHR_wayland_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to a Wayland `wl_surface`, as well as a query to determine support for rendering to a Wayland compositor.

## [](#_new_commands)New Commands

- [vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWaylandSurfaceKHR.html)
- [vkGetPhysicalDeviceWaylandPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceWaylandPresentationSupportKHR.html)

## [](#_new_structures)New Structures

- [VkWaylandSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWaylandSurfaceCreateInfoKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkWaylandSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWaylandSurfaceCreateFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_WAYLAND_SURFACE_EXTENSION_NAME`
- `VK_KHR_WAYLAND_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR`

## [](#_issues)Issues

1\) Does Wayland need a way to query for compatibility between a particular physical device and a specific Wayland display? This would be a more general query than [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html): if the Wayland-specific query returned `VK_TRUE` for a ([VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), `struct wl_display*`) pair, then the physical device could be assumed to support presentation to any [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) for surfaces on the display.

**RESOLVED**: Yes. [vkGetPhysicalDeviceWaylandPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceWaylandPresentationSupportKHR.html) was added to address this issue.

2\) Should we require surfaces created with [vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWaylandSurfaceKHR.html) to support the `VK_PRESENT_MODE_MAILBOX_KHR` present mode?

**RESOLVED**: Yes. Wayland is an inherently mailbox window system and mailbox support is required for some Wayland compositor interactions to work as expected. While handling these interactions may be possible with `VK_PRESENT_MODE_FIFO_KHR`, it is much more difficult to do without deadlock and requiring all Wayland applications to be able to support implementations which only support `VK_PRESENT_MODE_FIFO_KHR` would be an onerous restriction on application developers.

## [](#_version_history)Version History

- Revision 1, 2015-09-23 (Jesse Hall)
  
  - Initial draft, based on the previous contents of VK\_EXT\_KHR\_swapchain (later renamed VK\_EXT\_KHR\_surface).
- Revision 2, 2015-10-02 (James Jones)
  
  - Added vkGetPhysicalDeviceWaylandPresentationSupportKHR() to resolve issue #1.
  - Adjusted wording of issue #1 to match the agreed-upon solution.
  - Renamed “window” parameters to “surface” to match Wayland conventions.
- Revision 3, 2015-10-26 (Ian Elliott)
  
  - Renamed from VK\_EXT\_KHR\_wayland\_surface to VK\_KHR\_wayland\_surface.
- Revision 4, 2015-11-03 (Daniel Rakos)
  
  - Added allocation callbacks to vkCreateWaylandSurfaceKHR.
- Revision 5, 2015-11-28 (Daniel Rakos)
  
  - Updated the surface create function to take a pCreateInfo structure.
- Revision 6, 2017-02-08 (Faith Ekstrand)
  
  - Added the requirement that implementations support `VK_PRESENT_MODE_MAILBOX_KHR`.
  - Added wording about interactions between [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) and the Wayland requests sent to the compositor.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_wayland_surface)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
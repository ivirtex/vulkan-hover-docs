# VK\_KHR\_win32\_surface(3) Manual Page

## Name

VK\_KHR\_win32\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

10

## [](#_revision)Revision

6

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_win32_surface%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_win32_surface%20extension%2A)
- Ian Elliott [\[GitHub\]ianelliottus](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_win32_surface%5D%20%40ianelliottus%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_win32_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-04-24

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

The `VK_KHR_win32_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to a Win32 `HWND`, as well as a query to determine support for rendering to the windows desktop.

## [](#_new_commands)New Commands

- [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWin32SurfaceKHR.html)
- [vkGetPhysicalDeviceWin32PresentationSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceWin32PresentationSupportKHR.html)

## [](#_new_structures)New Structures

- [VkWin32SurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32SurfaceCreateInfoKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkWin32SurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32SurfaceCreateFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_WIN32_SURFACE_EXTENSION_NAME`
- `VK_KHR_WIN32_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR`

## [](#_issues)Issues

1\) Does Win32 need a way to query for compatibility between a particular physical device and a specific screen? Compatibility between a physical device and a window generally only depends on what screen the window is on. However, there is not an obvious way to identify a screen without already having a window on the screen.

**RESOLVED**: No. While it may be useful, there is not a clear way to do this on Win32. However, a method was added to query support for presenting to the windows desktop as a whole.

2\) If a native window object (`HWND`) is used by one graphics API, and then is later used by a different graphics API (one of which is Vulkan), can these uses interfere with each other?

**RESOLVED**: Yes.

Uses of a window object by multiple graphics APIs results in undefined behavior. Such behavior may succeed when using one Vulkan implementation but fail when using a different Vulkan implementation. Potential failures include:

- Creating then destroying a flip presentation model DXGI swapchain on a window object can prevent [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html) from succeeding on the same window object.
- Creating then destroying a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) on a window object can prevent creation of a bitblt model DXGI swapchain on the same window object.
- Creating then destroying a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) on a window object can effectively `SetPixelFormat` to a different format than the format chosen by an OpenGL application.
- Creating then destroying a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) on a window object on one [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) can prevent [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html) from succeeding on the same window object, but on a different [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) that is associated with a different Vulkan ICD.

In all cases the problem can be worked around by creating a new window object.

Technical details include:

- Creating a DXGI swapchain over a window object can alter the object for the remainder of its lifetime. The alteration persists even after the DXGI swapchain has been destroyed. This alteration can make it impossible for a conformant Vulkan implementation to create a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) over the same window object. Mention of this alteration can be found in the remarks section of the MSDN documentation for `DXGI_SWAP_EFFECT`.
- Calling GDI’s `SetPixelFormat` (needed by OpenGL’s WGL layer) on a window object alters the object for the remainder of its lifetime. The MSDN documentation for `SetPixelFormat` explains that a window object’s pixel format can be set only one time.
- Creating a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) over a window object can alter the object for its remaining lifetime. Either of the above alterations may occur as a side effect of [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).

## [](#_version_history)Version History

- Revision 1, 2015-09-23 (Jesse Hall)
  
  - Initial draft, based on the previous contents of VK\_EXT\_KHR\_swapchain (later renamed VK\_EXT\_KHR\_surface).
- Revision 2, 2015-10-02 (James Jones)
  
  - Added presentation support query for win32 desktops.
- Revision 3, 2015-10-26 (Ian Elliott)
  
  - Renamed from VK\_EXT\_KHR\_win32\_surface to VK\_KHR\_win32\_surface.
- Revision 4, 2015-11-03 (Daniel Rakos)
  
  - Added allocation callbacks to vkCreateWin32SurfaceKHR.
- Revision 5, 2015-11-28 (Daniel Rakos)
  
  - Updated the surface create function to take a pCreateInfo structure.
- Revision 6, 2017-04-24 (Jeff Juliano)
  
  - Add issue 2 addressing reuse of a native window object in a different Graphics API, or by a different Vulkan ICD.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_win32_surface).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
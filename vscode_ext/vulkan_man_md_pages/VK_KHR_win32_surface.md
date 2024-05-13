# VK_KHR_win32_surface(3) Manual Page

## Name

VK_KHR_win32_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

10

## <a href="#_revision" class="anchor"></a>Revision

6

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_win32_surface%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_win32_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>critsec</a>

- Ian Elliott <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_win32_surface%5D%20@ianelliottus%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_win32_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ianelliottus</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_win32_surface` extension is an instance extension. It
provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)
object (defined by the [`VK_KHR_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)
extension) that refers to a Win32 `HWND`, as well as a query to
determine support for rendering to the windows desktop.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWin32SurfaceKHR.html)

- [vkGetPhysicalDeviceWin32PresentationSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceWin32PresentationSupportKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkWin32SurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32SurfaceCreateInfoKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkWin32SurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32SurfaceCreateFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_WIN32_SURFACE_EXTENSION_NAME`

- `VK_KHR_WIN32_SURFACE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does Win32 need a way to query for compatibility between a
particular physical device and a specific screen? Compatibility between
a physical device and a window generally only depends on what screen the
window is on. However, there is not an obvious way to identify a screen
without already having a window on the screen.

**RESOLVED**: No. While it may be useful, there is not a clear way to do
this on Win32. However, a method was added to query support for
presenting to the windows desktop as a whole.

2\) If a native window object (`HWND`) is used by one graphics API, and
then is later used by a different graphics API (one of which is Vulkan),
can these uses interfere with each other?

**RESOLVED**: Yes.

Uses of a window object by multiple graphics APIs results in undefined
behavior. Such behavior may succeed when using one Vulkan implementation
but fail when using a different Vulkan implementation. Potential
failures include:

- Creating then destroying a flip presentation model DXGI swapchain on a
  window object can prevent
  [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html) from succeeding on
  the same window object.

- Creating then destroying a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) on a
  window object can prevent creation of a bitblt model DXGI swapchain on
  the same window object.

- Creating then destroying a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) on a
  window object can effectively `SetPixelFormat` to a different format
  than the format chosen by an OpenGL application.

- Creating then destroying a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) on a
  window object on one [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) can
  prevent [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html) from
  succeeding on the same window object, but on a different
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) that is associated with a
  different Vulkan ICD.

In all cases the problem can be worked around by creating a new window
object.

Technical details include:

- Creating a DXGI swapchain over a window object can alter the object
  for the remainder of its lifetime. The alteration persists even after
  the DXGI swapchain has been destroyed. This alteration can make it
  impossible for a conformant Vulkan implementation to create a
  [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) over the same window object.
  Mention of this alteration can be found in the remarks section of the
  MSDN documentation for `DXGI_SWAP_EFFECT`.

- Calling GDI’s `SetPixelFormat` (needed by OpenGL’s WGL layer) on a
  window object alters the object for the remainder of its lifetime. The
  MSDN documentation for `SetPixelFormat` explains that a window
  object’s pixel format can be set only one time.

- Creating a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) over a window object
  can alter the object for its remaining lifetime. Either of the above
  alterations may occur as a side effect of
  [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2015-09-23 (Jesse Hall)

  - Initial draft, based on the previous contents of
    VK_EXT_KHR_swapchain (later renamed VK_EXT_KHR_surface).

- Revision 2, 2015-10-02 (James Jones)

  - Added presentation support query for win32 desktops.

- Revision 3, 2015-10-26 (Ian Elliott)

  - Renamed from VK_EXT_KHR_win32_surface to VK_KHR_win32_surface.

- Revision 4, 2015-11-03 (Daniel Rakos)

  - Added allocation callbacks to vkCreateWin32SurfaceKHR.

- Revision 5, 2015-11-28 (Daniel Rakos)

  - Updated the surface create function to take a pCreateInfo structure.

- Revision 6, 2017-04-24 (Jeff Juliano)

  - Add issue 2 addressing reuse of a native window object in a
    different Graphics API, or by a different Vulkan ICD.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_win32_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

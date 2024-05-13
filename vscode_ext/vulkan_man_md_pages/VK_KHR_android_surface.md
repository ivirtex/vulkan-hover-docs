# VK_KHR_android_surface(3) Manual Page

## Name

VK_KHR_android_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

9

## <a href="#_revision" class="anchor"></a>Revision

6

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_android_surface%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_android_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>critsec</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_android_surface` extension is an instance extension. It
provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)
object (defined by the [`VK_KHR_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)
extension) that refers to an [ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html),
Android’s native surface type. The [ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html)
represents the producer endpoint of any buffer queue, regardless of
consumer endpoint. Common consumer endpoints for `ANativeWindows` are
the system window compositor, video encoders, and application-specific
compositors importing the images through a `SurfaceTexture`.

## <a href="#_new_base_types" class="anchor"></a>New Base Types

- [ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateAndroidSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAndroidSurfaceKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkAndroidSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidSurfaceCreateInfoKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkAndroidSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidSurfaceCreateFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_ANDROID_SURFACE_EXTENSION_NAME`

- `VK_KHR_ANDROID_SURFACE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does Android need a way to query for compatibility between a
particular physical device (and queue family?) and a specific Android
display?

**RESOLVED**: No. Currently on Android, any physical device is expected
to be able to present to the system compositor, and all queue families
must support the necessary image layout transitions and synchronization
operations.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2015-09-23 (Jesse Hall)

  - Initial draft.

- Revision 2, 2015-10-26 (Ian Elliott)

  - Renamed from VK_EXT_KHR_android_surface to VK_KHR_android_surface.

- Revision 3, 2015-11-03 (Daniel Rakos)

  - Added allocation callbacks to surface creation function.

- Revision 4, 2015-11-10 (Jesse Hall)

  - Removed VK_ERROR_INVALID_ANDROID_WINDOW_KHR.

- Revision 5, 2015-11-28 (Daniel Rakos)

  - Updated the surface create function to take a pCreateInfo structure.

- Revision 6, 2016-01-14 (James Jones)

  - Moved VK_ERROR_NATIVE_WINDOW_IN_USE_KHR from the
    VK_KHR_android_surface to the VK_KHR_surface extension.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_android_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# VK\_KHR\_surface(3) Manual Page

## Name

VK\_KHR\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

1

## [](#_revision)Revision

25

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_surface%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_surface%20extension%2A)
- Ian Elliott [\[GitHub\]ianelliottus](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_surface%5D%20%40ianelliottus%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-08-25

**IP Status**

No known IP claims.

**Contributors**

- Patrick Doane, Blizzard
- Ian Elliott, LunarG
- Jesse Hall, Google
- James Jones, NVIDIA
- David Mao, AMD
- Norbert Nopper, Freescale
- Alon Or-bach, Samsung
- Daniel Rakos, AMD
- Graham Sellers, AMD
- Jeff Vigil, Qualcomm
- Chia-I Wu, LunarG
- Faith Ekstrand, Intel

## [](#_description)Description

The `VK_KHR_surface` extension is an instance extension. It introduces [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) objects, which abstract native platform surface or window objects for use with Vulkan. It also provides a way to determine whether a queue family in a physical device supports presenting to particular surface.

Separate extensions for each platform provide the mechanisms for creating [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) objects, but once created they may be used in this and other platform-independent extensions, in particular the `VK_KHR_swapchain` extension.

## [](#_new_object_types)New Object Types

- [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_new_commands)New Commands

- [vkDestroySurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySurfaceKHR.html)
- [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html)
- [vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html)
- [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html)
- [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)

## [](#_new_structures)New Structures

- [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)
- [VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFormatKHR.html)

## [](#_new_enums)New Enums

- [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html)
- [VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompositeAlphaFlagBitsKHR.html)
- [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html)
- [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkCompositeAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompositeAlphaFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SURFACE_EXTENSION_NAME`
- `VK_KHR_SURFACE_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_SURFACE_KHR`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR`
  - `VK_ERROR_SURFACE_LOST_KHR`

## [](#_examples)Examples

Note

The example code for the `VK_KHR_surface` and `VK_KHR_swapchain` extensions was removed from the appendix after revision 1.0.29. This WSI example code was ported to the cube demo that is shipped with the official Khronos SDK, and is being kept up-to-date in that location (see: [https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c](https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c)).

## [](#_issues)Issues

1\) Should this extension include a method to query whether a physical device supports presenting to a specific window or native surface on a given platform?

**RESOLVED**: Yes. Without this, applications would need to create a device instance to determine whether a particular window can be presented to. Knowing that a device supports presentation to a platform in general is not sufficient, as a single machine might support multiple seats, or instances of the platform that each use different underlying physical devices. Additionally, on some platforms, such as the X Window System, different drivers and devices might be used for different windows depending on which section of the desktop they exist on.

2\) Should the [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html), [vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html), and [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html) functions be in this extension and operate on physical devices, rather than being in `VK_KHR_swapchain` (i.e. device extension) and being dependent on [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)?

**RESOLVED**: Yes. While it might be useful to depend on `VkDevice` (and therefore on enabled extensions and features) for the queries, Vulkan was released only with the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) versions. Many cases can be resolved by a Valid Usage statement, and/or by a separate `pNext` chain version of the query structure specific to a given extension or parameters, via extensible versions of the queries: [vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html), [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html), and [vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html).

3\) Should Vulkan support Xlib or XCB as the API for accessing the X Window System platform?

**RESOLVED**: Both. XCB is a more modern and efficient API, but Xlib usage is deeply ingrained in many applications and likely will remain in use for the foreseeable future. Not all drivers necessarily need to support both, but including both as options in the core specification will probably encourage support, which should in turn ease adoption of the Vulkan API in older codebases. Additionally, the performance improvements possible with XCB likely will not have a measurable impact on the performance of Vulkan presentation and other minimal window system interactions defined here.

4\) Should the GBM platform be included in the list of platform enums?

**RESOLVED**: Deferred, and will be addressed with a platform-specific extension to be written in the future.

## [](#_version_history)Version History

- Revision 1, 2015-05-20 (James Jones)
  
  - Initial draft, based on LunarG KHR spec, other KHR specs, patches attached to bugs.
- Revision 2, 2015-05-22 (Ian Elliott)
  
  - Created initial Description section.
  - Removed query for whether a platform requires the use of a queue for presentation, since it was decided that presentation will always be modeled as being part of the queue.
  - Fixed typos and other minor mistakes.
- Revision 3, 2015-05-26 (Ian Elliott)
  
  - Improved the Description section.
- Revision 4, 2015-05-27 (James Jones)
  
  - Fixed compilation errors in example code.
- Revision 5, 2015-06-01 (James Jones)
  
  - Added issues 1 and 2 and made related spec updates.
- Revision 6, 2015-06-01 (James Jones)
  
  - Merged the platform type mappings table previously removed from VK\_KHR\_swapchain with the platform description table in this spec.
  - Added issues 3 and 4 documenting choices made when building the initial list of native platforms supported.
- Revision 7, 2015-06-11 (Ian Elliott)
  
  - Updated table 1 per input from the KHR TSG.
  - Updated issue 4 (GBM) per discussion with Daniel Stone. He will create a platform-specific extension sometime in the future.
- Revision 8, 2015-06-17 (James Jones)
  
  - Updated enum-extending values using new convention.
  - Fixed the value of VK\_SURFACE\_PLATFORM\_INFO\_TYPE\_SUPPORTED\_KHR.
- Revision 9, 2015-06-17 (James Jones)
  
  - Rebased on Vulkan API version 126.
- Revision 10, 2015-06-18 (James Jones)
  
  - Marked issues 2 and 3 resolved.
- Revision 11, 2015-06-23 (Ian Elliott)
  
  - Examples now show use of function pointers for extension functions.
  - Eliminated extraneous whitespace.
- Revision 12, 2015-07-07 (Daniel Rakos)
  
  - Added error section describing when each error is expected to be reported.
  - Replaced the term “queue node index” with “queue family index” in the spec as that is the agreed term to be used in the latest version of the core header and spec.
  - Replaced bool32\_t with VkBool32.
- Revision 13, 2015-08-06 (Daniel Rakos)
  
  - Updated spec against latest core API header version.
- Revision 14, 2015-08-20 (Ian Elliott)
  
  - Renamed this extension and all of its enumerations, types, functions, etc. This makes it compliant with the proposed standard for Vulkan extensions.
  - Switched from “revision” to “version”, including use of the VK\_MAKE\_VERSION macro in the header file.
  - Did miscellaneous cleanup, etc.
- Revision 15, 2015-08-20 (Ian Elliott—​porting a 2015-07-29 change from James Jones)
  
  - Moved the surface transform enums here from VK\_WSI\_swapchain so they could be reused by VK\_WSI\_display.
- Revision 16, 2015-09-01 (James Jones)
  
  - Restore single-field revision number.
- Revision 17, 2015-09-01 (James Jones)
  
  - Fix example code compilation errors.
- Revision 18, 2015-09-26 (Jesse Hall)
  
  - Replaced VkSurfaceDescriptionKHR with the VkSurfaceKHR object, which is created via layered extensions. Added VkDestroySurfaceKHR.
- Revision 19, 2015-09-28 (Jesse Hall)
  
  - Renamed from VK\_EXT\_KHR\_swapchain to VK\_EXT\_KHR\_surface.
- Revision 20, 2015-09-30 (Jeff Vigil)
  
  - Add error result VK\_ERROR\_SURFACE\_LOST\_KHR.
- Revision 21, 2015-10-15 (Daniel Rakos)
  
  - Updated the resolution of issue #2 and include the surface capability queries in this extension.
  - Renamed SurfaceProperties to SurfaceCapabilities as it better reflects that the values returned are the capabilities of the surface on a particular device.
  - Other minor cleanup and consistency changes.
- Revision 22, 2015-10-26 (Ian Elliott)
  
  - Renamed from VK\_EXT\_KHR\_surface to VK\_KHR\_surface.
- Revision 23, 2015-11-03 (Daniel Rakos)
  
  - Added allocation callbacks to vkDestroySurfaceKHR.
- Revision 24, 2015-11-10 (Jesse Hall)
  
  - Removed VkSurfaceTransformKHR. Use VkSurfaceTransformFlagBitsKHR instead.
  - Rename VkSurfaceCapabilitiesKHR member maxImageArraySize to maxImageArrayLayers.
- Revision 25, 2016-01-14 (James Jones)
  
  - Moved VK\_ERROR\_NATIVE\_WINDOW\_IN\_USE\_KHR from the VK\_KHR\_android\_surface to the VK\_KHR\_surface extension.
- 2016-08-23 (Ian Elliott)
  
  - Update the example code, to not have so many characters per line, and to split out a new example to show how to obtain function pointers.
- 2016-08-25 (Ian Elliott)
  
  - A note was added at the beginning of the example code, stating that it will be removed from future versions of the appendix.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_surface)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VK_KHR_surface(3) Manual Page

## Name

VK_KHR_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

1

## <a href="#_revision" class="anchor"></a>Revision

25

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_surface%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

- Ian Elliott <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_surface%5D%20@ianelliottus%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ianelliottus</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_surface` extension is an instance extension. It introduces
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) objects, which abstract native
platform surface or window objects for use with Vulkan. It also provides
a way to determine whether a queue family in a physical device supports
presenting to particular surface.

Separate extensions for each platform provide the mechanisms for
creating [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) objects, but once created
they may be used in this and other platform-independent extensions, in
particular the [`VK_KHR_swapchain`](VK_KHR_swapchain.html) extension.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkDestroySurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySurfaceKHR.html)

- [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html)

- [vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html)

- [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html)

- [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html)

- [VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html)

- [VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html)

- [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html)

- [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkCompositeAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SURFACE_EXTENSION_NAME`

- `VK_KHR_SURFACE_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_SURFACE_KHR`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR`

  - `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_examples" class="anchor"></a>Examples

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The example code for the <code>VK_KHR_surface</code> and <a
href="VK_KHR_swapchain.html"><code>VK_KHR_swapchain</code></a>
extensions was removed from the appendix after revision 1.0.29. This WSI
example code was ported to the cube demo that is shipped with the
official Khronos SDK, and is being kept up-to-date in that location
(see: <a
href="https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c"
class="bare">https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c</a>).</p></td>
</tr>
</tbody>
</table>

## <a href="#_issues" class="anchor"></a>Issues

1\) Should this extension include a method to query whether a physical
device supports presenting to a specific window or native surface on a
given platform?

**RESOLVED**: Yes. Without this, applications would need to create a
device instance to determine whether a particular window can be
presented to. Knowing that a device supports presentation to a platform
in general is not sufficient, as a single machine might support multiple
seats, or instances of the platform that each use different underlying
physical devices. Additionally, on some platforms, such as the X Window
System, different drivers and devices might be used for different
windows depending on which section of the desktop they exist on.

2\) Should the
[vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html),
[vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html),
and
[vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html)
functions be in this extension and operate on physical devices, rather
than being in [`VK_KHR_swapchain`](VK_KHR_swapchain.html) (i.e. device
extension) and being dependent on [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)?

**RESOLVED**: Yes. While it might be useful to depend on `VkDevice` (and
therefore on enabled extensions and features) for the queries, Vulkan
was released only with the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)
versions. Many cases can be resolved by a Valid Usage statement, and/or
by a separate `pNext` chain version of the query struct specific to a
given extension or parameters, via extensible versions of the queries:
[vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html),
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html),
and
[vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html).

3\) Should Vulkan support Xlib or XCB as the API for accessing the X
Window System platform?

**RESOLVED**: Both. XCB is a more modern and efficient API, but Xlib
usage is deeply ingrained in many applications and likely will remain in
use for the foreseeable future. Not all drivers necessarily need to
support both, but including both as options in the core specification
will probably encourage support, which should in turn ease adoption of
the Vulkan API in older codebases. Additionally, the performance
improvements possible with XCB likely will not have a measurable impact
on the performance of Vulkan presentation and other minimal window
system interactions defined here.

4\) Should the GBM platform be included in the list of platform enums?

**RESOLVED**: Deferred, and will be addressed with a platform-specific
extension to be written in the future.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2015-05-20 (James Jones)

  - Initial draft, based on LunarG KHR spec, other KHR specs, patches
    attached to bugs.

- Revision 2, 2015-05-22 (Ian Elliott)

  - Created initial Description section.

  - Removed query for whether a platform requires the use of a queue for
    presentation, since it was decided that presentation will always be
    modeled as being part of the queue.

  - Fixed typos and other minor mistakes.

- Revision 3, 2015-05-26 (Ian Elliott)

  - Improved the Description section.

- Revision 4, 2015-05-27 (James Jones)

  - Fixed compilation errors in example code.

- Revision 5, 2015-06-01 (James Jones)

  - Added issues 1 and 2 and made related spec updates.

- Revision 6, 2015-06-01 (James Jones)

  - Merged the platform type mappings table previously removed from
    VK_KHR_swapchain with the platform description table in this spec.

  - Added issues 3 and 4 documenting choices made when building the
    initial list of native platforms supported.

- Revision 7, 2015-06-11 (Ian Elliott)

  - Updated table 1 per input from the KHR TSG.

  - Updated issue 4 (GBM) per discussion with Daniel Stone. He will
    create a platform-specific extension sometime in the future.

- Revision 8, 2015-06-17 (James Jones)

  - Updated enum-extending values using new convention.

  - Fixed the value of VK_SURFACE_PLATFORM_INFO_TYPE_SUPPORTED_KHR.

- Revision 9, 2015-06-17 (James Jones)

  - Rebased on Vulkan API version 126.

- Revision 10, 2015-06-18 (James Jones)

  - Marked issues 2 and 3 resolved.

- Revision 11, 2015-06-23 (Ian Elliott)

  - Examples now show use of function pointers for extension functions.

  - Eliminated extraneous whitespace.

- Revision 12, 2015-07-07 (Daniel Rakos)

  - Added error section describing when each error is expected to be
    reported.

  - Replaced the term “queue node index” with “queue family index” in
    the spec as that is the agreed term to be used in the latest version
    of the core header and spec.

  - Replaced bool32_t with VkBool32.

- Revision 13, 2015-08-06 (Daniel Rakos)

  - Updated spec against latest core API header version.

- Revision 14, 2015-08-20 (Ian Elliott)

  - Renamed this extension and all of its enumerations, types,
    functions, etc. This makes it compliant with the proposed standard
    for Vulkan extensions.

  - Switched from “revision” to “version”, including use of the
    VK_MAKE_VERSION macro in the header file.

  - Did miscellaneous cleanup, etc.

- Revision 15, 2015-08-20 (Ian Elliott—​porting a 2015-07-29 change from
  James Jones)

  - Moved the surface transform enums here from VK_WSI_swapchain so they
    could be reused by VK_WSI_display.

- Revision 16, 2015-09-01 (James Jones)

  - Restore single-field revision number.

- Revision 17, 2015-09-01 (James Jones)

  - Fix example code compilation errors.

- Revision 18, 2015-09-26 (Jesse Hall)

  - Replaced VkSurfaceDescriptionKHR with the VkSurfaceKHR object, which
    is created via layered extensions. Added VkDestroySurfaceKHR.

- Revision 19, 2015-09-28 (Jesse Hall)

  - Renamed from VK_EXT_KHR_swapchain to VK_EXT_KHR_surface.

- Revision 20, 2015-09-30 (Jeff Vigil)

  - Add error result VK_ERROR_SURFACE_LOST_KHR.

- Revision 21, 2015-10-15 (Daniel Rakos)

  - Updated the resolution of issue \#2 and include the surface
    capability queries in this extension.

  - Renamed SurfaceProperties to SurfaceCapabilities as it better
    reflects that the values returned are the capabilities of the
    surface on a particular device.

  - Other minor cleanup and consistency changes.

- Revision 22, 2015-10-26 (Ian Elliott)

  - Renamed from VK_EXT_KHR_surface to VK_KHR_surface.

- Revision 23, 2015-11-03 (Daniel Rakos)

  - Added allocation callbacks to vkDestroySurfaceKHR.

- Revision 24, 2015-11-10 (Jesse Hall)

  - Removed VkSurfaceTransformKHR. Use VkSurfaceTransformFlagBitsKHR
    instead.

  - Rename VkSurfaceCapabilitiesKHR member maxImageArraySize to
    maxImageArrayLayers.

- Revision 25, 2016-01-14 (James Jones)

  - Moved VK_ERROR_NATIVE_WINDOW_IN_USE_KHR from the
    VK_KHR_android_surface to the VK_KHR_surface extension.

- 2016-08-23 (Ian Elliott)

  - Update the example code, to not have so many characters per line,
    and to split out a new example to show how to obtain function
    pointers.

- 2016-08-25 (Ian Elliott)

  - A note was added at the beginning of the example code, stating that
    it will be removed from future versions of the appendix.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# VK_EXT_full_screen_exclusive(3) Manual Page

## Name

VK_EXT_full_screen_exclusive - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

256

## <a href="#_revision" class="anchor"></a>Revision

4

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  
and  
[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  
and  
[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_1

- Interacts with VK_KHR_device_group

- Interacts with VK_KHR_win32_surface

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_full_screen_exclusive%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_full_screen_exclusive%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-03-12

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- Interacts with Vulkan 1.1

- Interacts with [`VK_KHR_device_group`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html)

- Interacts with [`VK_KHR_win32_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_win32_surface.html)

**Contributors**  
- Hans-Kristian Arntzen, ARM

- Slawomir Grajewski, Intel

- Tobias Hector, AMD

- James Jones, NVIDIA

- Daniel Rakos, AMD

- Jeff Juliano, NVIDIA

- Joshua Schnarr, NVIDIA

- Aaron Hagan, AMD

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to set the policy for swapchain
creation and presentation mechanisms relating to full-screen access.
Implementations may be able to acquire exclusive access to a particular
display for an application window that covers the whole screen. This can
increase performance on some systems by bypassing composition, however
it can also result in disruptive or expensive transitions in the
underlying windowing system when a change occurs.

Applications can choose between explicitly disallowing or allowing this
behavior, letting the implementation decide, or managing this mode of
operation directly using the new
[vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireFullScreenExclusiveModeEXT.html)
and
[vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseFullScreenExclusiveModeEXT.html)
commands.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireFullScreenExclusiveModeEXT.html)

- [vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html)

- [vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseFullScreenExclusiveModeEXT.html)

If [VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html) is supported:

- [vkGetDeviceGroupSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModes2EXT.html)

If [Version 1.1](#versions-1.1) is supported:

- [vkGetDeviceGroupSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModes2EXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html),
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html):

  - [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html):

  - [VkSurfaceCapabilitiesFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesFullScreenExclusiveEXT.html)

If [VK_KHR_win32_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_win32_surface.html) is supported:

- Extending
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html),
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html):

  - [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFullScreenExclusiveEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_FULL_SCREEN_EXCLUSIVE_EXTENSION_NAME`

- `VK_EXT_FULL_SCREEN_EXCLUSIVE_SPEC_VERSION`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT`

  - `VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT`

If [VK_KHR_win32_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_win32_surface.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) What should the extension & flag be called?

**RESOLVED**: VK_EXT_full_screen_exclusive.

Other options considered (prior to the app-controlled mode) were:

- VK_EXT_smooth_fullscreen_transition

- VK_EXT_fullscreen_behavior

- VK_EXT_fullscreen_preference

- VK_EXT_fullscreen_hint

- VK_EXT_fast_fullscreen_transition

- VK_EXT_avoid_fullscreen_exclusive

2\) Do we need more than a boolean toggle?

**RESOLVED**: Yes.

Using an enum with default/allowed/disallowed/app-controlled enables
applications to accept driver default behavior, specifically override it
in either direction without implying the driver is ever required to use
full-screen exclusive mechanisms, or manage this mode explicitly.

3\) Should this be a KHR or EXT extension?

**RESOLVED**: EXT, in order to allow it to be shipped faster.

4\) Can the fullscreen hint affect the surface capabilities, and if so,
should the hint also be specified as input when querying the surface
capabilities?

**RESOLVED**: Yes on both accounts.

While the hint does not guarantee a particular fullscreen mode will be
used when the swapchain is created, it can sometimes imply particular
modes will NOT be used. If the driver determines that it will opt-out of
using a particular mode based on the policy, and knows it can only
support certain capabilities if that mode is used, it would be confusing
at best to the application to report those capabilities in such cases.
Not allowing implementations to report this state to applications could
result in situations where applications are unable to determine why
swapchain creation fails when they specify certain hint values, which
could result in never- terminating surface creation loops.

5\) Should full-screen be one word or two?

**RESOLVED**: Two words.

"Fullscreen" is not in my dictionary, and web searches did not turn up
definitive proof that it is a colloquially accepted compound word.
Documentation for the corresponding Windows API mechanisms dithers. The
text consistently uses a hyphen, but none-the-less, there is a
SetFullscreenState method in the DXGI swapchain object. Given this
inconclusive external guidance, it is best to adhere to the Vulkan style
guidelines and avoid inventing new compound words.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 4, 2019-03-12 (Tobias Hector)

  - Added application-controlled mode, and related functions

  - Tidied up appendix

- Revision 3, 2019-01-03 (James Jones)

  - Renamed to VK_EXT_full_screen_exclusive

  - Made related adjustments to the tri-state enumerant names.

- Revision 2, 2018-11-27 (James Jones)

  - Renamed to VK_KHR_fullscreen_behavior

  - Switched from boolean flag to tri-state enum

- Revision 1, 2018-11-06 (James Jones)

  - Internal revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_full_screen_exclusive"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

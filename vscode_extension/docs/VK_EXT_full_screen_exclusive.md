# VK\_EXT\_full\_screen\_exclusive(3) Manual Page

## Name

VK\_EXT\_full\_screen\_exclusive - device extension



## [](#_registered_extension_number)Registered Extension Number

256

## [](#_revision)Revision

4

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)  
and  
[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_1
- Interacts with VK\_KHR\_device\_group
- Interacts with VK\_KHR\_win32\_surface

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_full_screen_exclusive%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_full_screen_exclusive%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-03-12

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Interacts with Vulkan 1.1
- Interacts with `VK_KHR_device_group`
- Interacts with `VK_KHR_win32_surface`

**Contributors**

- Hans-Kristian Arntzen, ARM
- Slawomir Grajewski, Intel
- Tobias Hector, AMD
- James Jones, NVIDIA
- Daniel Rakos, AMD
- Jeff Juliano, NVIDIA
- Joshua Schnarr, NVIDIA
- Aaron Hagan, AMD

## [](#_description)Description

This extension allows applications to set the policy for swapchain creation and presentation mechanisms relating to full-screen access. Implementations may be able to acquire exclusive access to a particular display for an application window that covers the whole screen. This can increase performance on some systems by bypassing composition, however it can also result in disruptive or expensive transitions in the underlying windowing system when a change occurs.

Applications can choose between explicitly disallowing or allowing this behavior, letting the implementation decide, or managing this mode of operation directly using the new [vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireFullScreenExclusiveModeEXT.html) and [vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseFullScreenExclusiveModeEXT.html) commands.

## [](#_new_commands)New Commands

- [vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireFullScreenExclusiveModeEXT.html)
- [vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html)
- [vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseFullScreenExclusiveModeEXT.html)

If [VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html) or [Vulkan Version 1.1](#versions-1.1) is supported:

- [vkGetDeviceGroupSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupSurfacePresentModes2EXT.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html):
  
  - [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)
- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkSurfaceCapabilitiesFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesFullScreenExclusiveEXT.html)

If [VK\_KHR\_win32\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_win32_surface.html) is supported:

- Extending [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html):
  
  - [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html)

## [](#_new_enums)New Enums

- [VkFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFullScreenExclusiveEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_FULL_SCREEN_EXCLUSIVE_EXTENSION_NAME`
- `VK_EXT_FULL_SCREEN_EXCLUSIVE_SPEC_VERSION`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT`
  - `VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT`

If [VK\_KHR\_win32\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_win32_surface.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT`

## [](#_issues)Issues

1\) What should the extension &amp; flag be called?

**RESOLVED**: VK\_EXT\_full\_screen\_exclusive.

Other options considered (prior to the app-controlled mode) were:

- VK\_EXT\_smooth\_fullscreen\_transition
- VK\_EXT\_fullscreen\_behavior
- VK\_EXT\_fullscreen\_preference
- VK\_EXT\_fullscreen\_hint
- VK\_EXT\_fast\_fullscreen\_transition
- VK\_EXT\_avoid\_fullscreen\_exclusive

2\) Do we need more than a boolean toggle?

**RESOLVED**: Yes.

Using an enum with default/allowed/disallowed/app-controlled enables applications to accept driver default behavior, specifically override it in either direction without implying the driver is ever required to use full-screen exclusive mechanisms, or manage this mode explicitly.

3\) Should this be a KHR or EXT extension?

**RESOLVED**: EXT, in order to allow it to be shipped faster.

4\) Can the fullscreen hint affect the surface capabilities, and if so, should the hint also be specified as input when querying the surface capabilities?

**RESOLVED**: Yes on both accounts.

While the hint does not guarantee a particular fullscreen mode will be used when the swapchain is created, it can sometimes imply particular modes will NOT be used. If the driver determines that it will opt-out of using a particular mode based on the policy, and knows it can only support certain capabilities if that mode is used, it would be confusing at best to the application to report those capabilities in such cases. Not allowing implementations to report this state to applications could result in situations where applications are unable to determine why swapchain creation fails when they specify certain hint values, which could result in never- terminating surface creation loops.

5\) Should full-screen be one word or two?

**RESOLVED**: Two words.

"Fullscreen" is not in my dictionary, and web searches did not turn up definitive proof that it is a colloquially accepted compound word. Documentation for the corresponding Windows API mechanisms dithers. The text consistently uses a hyphen, but none-the-less, there is a SetFullscreenState method in the DXGI swapchain object. Given this inconclusive external guidance, it is best to adhere to the Vulkan style guidelines and avoid inventing new compound words.

## [](#_version_history)Version History

- Revision 4, 2019-03-12 (Tobias Hector)
  
  - Added application-controlled mode, and related functions
  - Tidied up appendix
- Revision 3, 2019-01-03 (James Jones)
  
  - Renamed to VK\_EXT\_full\_screen\_exclusive
  - Made related adjustments to the tri-state enumerant names.
- Revision 2, 2018-11-27 (James Jones)
  
  - Renamed to VK\_KHR\_fullscreen\_behavior
  - Switched from boolean flag to tri-state enum
- Revision 1, 2018-11-06 (James Jones)
  
  - Internal revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_full_screen_exclusive).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
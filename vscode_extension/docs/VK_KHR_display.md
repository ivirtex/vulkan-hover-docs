# VK\_KHR\_display(3) Manual Page

## Name

VK\_KHR\_display - instance extension



## [](#_registered_extension_number)Registered Extension Number

3

## [](#_revision)Revision

23

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_display%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_display%20extension%2A)
- Norbert Nopper [\[GitHub\]FslNopper](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_display%5D%20%40FslNopper%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_display%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-03-13

**IP Status**

No known IP claims.

**Contributors**

- James Jones, NVIDIA
- Norbert Nopper, Freescale
- Jeff Vigil, Qualcomm
- Daniel Rakos, AMD

## [](#_description)Description

This extension provides the API to enumerate displays and available modes on a given device.

## [](#_new_object_types)New Object Types

- [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html)
- [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html)

## [](#_new_commands)New Commands

- [vkCreateDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDisplayModeKHR.html)
- [vkCreateDisplayPlaneSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDisplayPlaneSurfaceKHR.html)
- [vkGetDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayModePropertiesKHR.html)
- [vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneCapabilitiesKHR.html)
- [vkGetDisplayPlaneSupportedDisplaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneSupportedDisplaysKHR.html)
- [vkGetPhysicalDeviceDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDisplayPlanePropertiesKHR.html)
- [vkGetPhysicalDeviceDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDisplayPropertiesKHR.html)

## [](#_new_structures)New Structures

- [VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeCreateInfoKHR.html)
- [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeParametersKHR.html)
- [VkDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModePropertiesKHR.html)
- [VkDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneCapabilitiesKHR.html)
- [VkDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlanePropertiesKHR.html)
- [VkDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPropertiesKHR.html)
- [VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceCreateInfoKHR.html)

## [](#_new_enums)New Enums

- [VkDisplayPlaneAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneAlphaFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDisplayModeCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeCreateFlagsKHR.html)
- [VkDisplayPlaneAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneAlphaFlagsKHR.html)
- [VkDisplaySurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceCreateFlagsKHR.html)
- [VkSurfaceTransformFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DISPLAY_EXTENSION_NAME`
- `VK_KHR_DISPLAY_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_DISPLAY_KHR`
  - `VK_OBJECT_TYPE_DISPLAY_MODE_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR`

## [](#_issues)Issues

1\) Which properties of a mode should be fixed in the mode information vs. settable in some other function when setting the mode? E.g., do we need to double the size of the mode pool to include both stereo and non-stereo modes? YUV and RGB scanout even if they both take RGB input images? BGR vs. RGB input? etc.

**RESOLVED**: Many modern displays support at most a handful of resolutions and timings natively. Other “modes” are expected to be supported using scaling hardware on the display engine or GPU. Other properties, such as rotation and mirroring should not require duplicating hardware modes just to express all combinations. Further, these properties may be implemented on a per-display or per-overlay granularity.

To avoid the exponential growth of modes as mutable properties are added, as was the case with `EGLConfig`/WGL pixel formats/`GLXFBConfig`, this specification should separate out hardware properties and configurable state into separate objects. Modes and overlay planes will express capabilities of the hardware, while a separate structure will allow applications to configure scaling, rotation, mirroring, color keys, LUT values, alpha masks, etc. for a given swapchain independent of the mode in use. Constraints on these settings will be established by properties of the immutable objects.

Note the resolution of this issue may affect issue 5 as well.

2\) What properties of a display itself are useful?

**RESOLVED**: This issue is too broad. It was meant to prompt general discussion, but resolving this issue amounts to completing this specification. All interesting properties should be included. The issue will remain as a placeholder since removing it would make it hard to parse existing discussion notes that refer to issues by number.

3\) How are multiple overlay planes within a display or mode enumerated?

**RESOLVED**: They are referred to by an index. Each display will report the number of overlay planes it contains.

4\) Should swapchains be created relative to a mode or a display?

**RESOLVED**: When using this extension, swapchains are created relative to a mode and a plane. The mode implies the display object the swapchain will present to. If the specified mode is not the display’s current mode, the new mode will be applied when the first image is presented to the swapchain, and the default operating system mode, if any, will be restored when the swapchain is destroyed.

5\) Should users query generic ranges from displays and construct their own modes explicitly using those constraints rather than querying a fixed set of modes (Most monitors only have one real “mode” these days, even though many support relatively arbitrary scaling, either on the monitor side or in the GPU display engine, making “modes” something of a relic/compatibility construct).

**RESOLVED**: Expose both. Display information structures will expose a set of predefined modes, as well as any attributes necessary to construct a customized mode.

6\) Is it fine if we return the display and display mode handles in the structure used to query their properties?

**RESOLVED**: Yes.

7\) Is there a possibility that not all displays of a device work with all of the present queues of a device? If yes, how do we determine which displays work with which present queues?

**RESOLVED**: No known hardware has such limitations, but determining such limitations is supported automatically using the existing `VK_KHR_surface` and `VK_KHR_swapchain` query mechanisms.

8\) Should all presentation need to be done relative to an overlay plane, or can a display mode + display be used alone to target an output?

**RESOLVED**: Require specifying a plane explicitly.

9\) Should displays have an associated window system display, such as an `HDC` or `Display*`?

**RESOLVED**: No. Displays are independent of any windowing system in use on the system. Further, neither `HDC` nor `Display*` refer to a physical display object.

10\) Are displays queried from a physical GPU or from a device instance?

**RESOLVED**: Developers prefer to query modes directly from the physical GPU so they can use display information as an input to their device selection algorithms prior to device creation. This avoids the need to create placeholder device instances to enumerate displays.

This preference must be weighed against the extra initialization that must be done by driver vendors prior to device instance creation to support this usage.

11\) Should displays and/or modes be dispatchable objects? If functions are to take displays, overlays, or modes as their first parameter, they must be dispatchable objects as defined in Khronos bug 13529. If they are not added to the list of dispatchable objects, functions operating on them must take some higher-level object as their first parameter. There is no performance case against making them dispatchable objects, but they would be the first extension objects to be dispatchable.

**RESOLVED**: Do not make displays or modes dispatchable. They will dispatch based on their associated physical device.

12\) Should hardware cursor capabilities be exposed?

**RESOLVED**: Defer. This could be a separate extension on top of the base WSI specs.

13\) How many display objects should be enumerated for "tiled" display devices? There are ongoing design discussions among lower-level display API authors regarding how to expose displays if they are one physical display device to an end user, but may internally be implemented as two side-by-side displays using the same display engine (and sometimes cabling) resources as two physically separate display devices.

**RESOLVED**: Tiled displays will appear as a single display object in this API.

14\) Should the raw EDID data be included in the display information?

**RESOLVED**: No. A future extension could be added which reports the EDID if necessary. This may be complicated by the outcome of issue 13.

15\) Should min and max scaling factor capabilities of overlays be exposed?

**RESOLVED**: Yes. This is exposed indirectly by allowing applications to query the min/max position and extent of the source and destination regions from which image contents are fetched by the display engine when using a particular mode and overlay pair.

16\) Should devices be able to expose planes that can be moved between displays? If so, how?

**RESOLVED**: Yes. Applications can determine which displays a given plane supports using [vkGetDisplayPlaneSupportedDisplaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneSupportedDisplaysKHR.html).

17\) Should there be a way to destroy display modes? If so, does it support destroying “built in” modes?

**RESOLVED**: Not in this extension. A future extension could add this functionality.

18\) What should the lifetime of display and built-in display mode objects be?

**RESOLVED**: The lifetime of the instance. These objects cannot be destroyed. A future extension may be added to expose a way to destroy these objects and/or support display hotplug.

19\) Should persistent mode for smart panels be enabled/disabled at swapchain creation time, or on a per-present basis.

**RESOLVED**: On a per-present basis.

## [](#_examples)Examples

Note

The example code for the `VK_KHR_display` and `VK_KHR_display_swapchain` extensions was removed from the appendix after revision 1.0.43. The display enumeration example code was ported to the cube demo that is shipped with the official Khronos SDK, and is being kept up-to-date in that location (see: [https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c](https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c)).

## [](#_version_history)Version History

- Revision 1, 2015-02-24 (James Jones)
  
  - Initial draft
- Revision 2, 2015-03-12 (Norbert Nopper)
  
  - Added overlay enumeration for a display.
- Revision 3, 2015-03-17 (Norbert Nopper)
  
  - Fixed typos and namings as discussed in Bugzilla.
  - Reordered and grouped functions.
  - Added functions to query count of display, mode and overlay.
  - Added native display handle, which may be needed on some platforms to create a native Window.
- Revision 4, 2015-03-18 (Norbert Nopper)
  
  - Removed primary and virtualPostion members (see comment of James Jones in Bugzilla).
  - Added native overlay handle to information structure.
  - Replaced , with ; in struct.
- Revision 6, 2015-03-18 (Daniel Rakos)
  
  - Added WSI extension suffix to all items.
  - Made the whole API more “Vulkanish”.
  - Replaced all functions with a single vkGetDisplayInfoKHR function to better match the rest of the API.
  - Made the display, display mode, and overlay objects be first class objects, not subclasses of VkBaseObject as they do not support the common functions anyways.
  - Renamed \*Info structures to \*Properties.
  - Removed overlayIndex field from VkOverlayProperties as there is an implicit index already as a result of moving to a “Vulkanish” API.
  - Displays are not get through device, but through physical GPU to match the rest of the Vulkan API. Also this is something ISVs explicitly requested.
  - Added issue (6) and (7).
- Revision 7, 2015-03-25 (James Jones)
  
  - Added an issues section
  - Added rotation and mirroring flags
- Revision 8, 2015-03-25 (James Jones)
  
  - Combined the duplicate issues sections introduced in last change.
  - Added proposed resolutions to several issues.
- Revision 9, 2015-04-01 (Daniel Rakos)
  
  - Rebased extension against Vulkan 0.82.0
- Revision 10, 2015-04-01 (James Jones)
  
  - Added issues (10) and (11).
  - Added more straw-man issue resolutions, and cleaned up the proposed resolution for issue (4).
  - Updated the rotation and mirroring enums to have proper bitmask semantics.
- Revision 11, 2015-04-15 (James Jones)
  
  - Added proposed resolution for issues (1) and (2).
  - Added issues (12), (13), (14), and (15)
  - Removed pNativeHandle field from overlay structure.
  - Fixed small compilation errors in example code.
- Revision 12, 2015-07-29 (James Jones)
  
  - Rewrote the guts of the extension against the latest WSI swapchain specifications and the latest Vulkan API.
  - Address overlay planes by their index rather than an object handle and refer to them as “planes” rather than “overlays” to make it slightly clearer that even a display with no “overlays” still has at least one base “plane” that images can be displayed on.
  - Updated most of the issues.
  - Added an “extension type” section to the specification header.
  - Reused the VK\_EXT\_KHR\_surface surface transform enumerations rather than redefining them here.
  - Updated the example code to use the new semantics.
- Revision 13, 2015-08-21 (Ian Elliott)
  
  - Renamed this extension and all of its enumerations, types, functions, etc. This makes it compliant with the proposed standard for Vulkan extensions.
  - Switched from “revision” to “version”, including use of the VK\_MAKE\_VERSION macro in the header file.
- Revision 14, 2015-09-01 (James Jones)
  
  - Restore single-field revision number.
- Revision 15, 2015-09-08 (James Jones)
  
  - Added alpha flags enum.
  - Added premultiplied alpha support.
- Revision 16, 2015-09-08 (James Jones)
  
  - Added description section to the spec.
  - Added issues 16 - 18.
- Revision 17, 2015-10-02 (James Jones)
  
  - Planes are now a property of the entire device rather than individual displays. This allows planes to be moved between multiple displays on devices that support it.
  - Added a function to create a VkSurfaceKHR object describing a display plane and mode to align with the new per-platform surface creation conventions.
  - Removed detailed mode timing data. It was agreed that the mode extents and refresh rate are sufficient for current use cases. Other information could be added back in as an extension if it is needed in the future.
  - Added support for smart/persistent/buffered display devices.
- Revision 18, 2015-10-26 (Ian Elliott)
  
  - Renamed from VK\_EXT\_KHR\_display to VK\_KHR\_display.
- Revision 19, 2015-11-02 (James Jones)
  
  - Updated example code to match revision 17 changes.
- Revision 20, 2015-11-03 (Daniel Rakos)
  
  - Added allocation callbacks to creation functions.
- Revision 21, 2015-11-10 (Jesse Hall)
  
  - Added VK\_DISPLAY\_PLANE\_ALPHA\_OPAQUE\_BIT\_KHR, and use VkDisplayPlaneAlphaFlagBitsKHR for VkDisplayPlanePropertiesKHR::alphaMode instead of VkDisplayPlaneAlphaFlagsKHR, since it only represents one mode.
  - Added reserved flags bitmask to VkDisplayPlanePropertiesKHR.
  - Use VkSurfaceTransformFlagBitsKHR instead of obsolete VkSurfaceTransformKHR.
  - Renamed vkGetDisplayPlaneSupportedDisplaysKHR parameters for clarity.
- Revision 22, 2015-12-18 (James Jones)
  
  - Added missing “planeIndex” parameter to vkGetDisplayPlaneSupportedDisplaysKHR()
- Revision 23, 2017-03-13 (James Jones)
  
  - Closed all remaining issues. The specification and implementations have been shipping with the proposed resolutions for some time now.
  - Removed the sample code and noted it has been integrated into the official Vulkan SDK cube demo.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_display)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VK\_GOOGLE\_display\_timing(3) Manual Page

## Name

VK\_GOOGLE\_display\_timing - device extension



## [](#_registered_extension_number)Registered Extension Number

93

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)

## [](#_contact)Contact

- Ian Elliott [\[GitHub\]ianelliottus](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_GOOGLE_display_timing%5D%20%40ianelliottus%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_GOOGLE_display_timing%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-02-14

**IP Status**

No known IP claims.

**Contributors**

- Ian Elliott, Google
- Jesse Hall, Google

## [](#_description)Description

This device extension allows an application that uses the `VK_KHR_swapchain` extension to obtain information about the presentation engine’s display, to obtain timing information about each present, and to schedule a present to happen no earlier than a desired time. An application can use this to minimize various visual anomalies (e.g. stuttering).

Traditional game and real-time animation applications need to correctly position their geometry for when the presentable image will be presented to the user. To accomplish this, applications need various timing information about the presentation engine’s display. They need to know when presentable images were actually presented, and when they could have been presented. Applications also need to tell the presentation engine to display an image no sooner than a given time. This allows the application to avoid stuttering, so the animation looks smooth to the user.

This extension treats variable-refresh-rate (VRR) displays as if they are fixed-refresh-rate (FRR) displays.

## [](#_new_commands)New Commands

- [vkGetPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPastPresentationTimingGOOGLE.html)
- [vkGetRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRefreshCycleDurationGOOGLE.html)

## [](#_new_structures)New Structures

- [VkPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPastPresentationTimingGOOGLE.html)
- [VkPresentTimeGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentTimeGOOGLE.html)
- [VkRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRefreshCycleDurationGOOGLE.html)
- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkPresentTimesInfoGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentTimesInfoGOOGLE.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_GOOGLE_DISPLAY_TIMING_EXTENSION_NAME`
- `VK_GOOGLE_DISPLAY_TIMING_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE`

## [](#_examples)Examples

Note

The example code for the this extension (like the `VK_KHR_surface` and `VK_GOOGLE_display_timing` extensions) is contained in the cube demo that is shipped with the official Khronos SDK, and is being kept up-to-date in that location (see: [https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c](https://github.com/KhronosGroup/Vulkan-Tools/blob/main/cube/cube.c) ).

## [](#_version_history)Version History

- Revision 1, 2017-02-14 (Ian Elliott)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_GOOGLE_display_timing)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VK\_KHR\_surface\_maintenance1(3) Manual Page

## Name

VK\_KHR\_surface\_maintenance1 - instance extension



## [](#_registered_extension_number)Registered Extension Number

487

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)  
or  
[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_surface_maintenance1%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_surface_maintenance1%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_surface\_maintenance1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_surface_maintenance1.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-03-31

**Contributors**

- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Lionel Landwerlin, Intel
- Shahbaz Youssefi, Google
- Chris Forbes, Google
- Ian Elliott, Google
- Hans-Kristian Arntzen, Valve
- Daniel Stone, Collabora

## [](#_description)Description

This extension is based off the `VK_EXT_surface_maintenance1` extension.

`VK_KHR_surface_maintenance1` adds a collection of window system integration features that were intentionally left out or overlooked in the original `VK_KHR_surface` extension.

The new features are as follows:

- Allow querying number of min/max images from a surface for a particular presentation mode.
- Allow querying a surface’s scaled presentation capabilities.
- Allow querying a surface for the set of presentation modes which can be easily switched between without requiring swapchain recreation.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html):
  
  - [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkSurfacePresentModeCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityKHR.html)
  - [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)

## [](#_new_enums)New Enums

- [VkPresentGravityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsKHR.html)
- [VkPresentScalingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPresentGravityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagsKHR.html)
- [VkPresentScalingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SURFACE_MAINTENANCE_1_EXTENSION_NAME`
- `VK_KHR_SURFACE_MAINTENANCE_1_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_COMPATIBILITY_KHR`
  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_KHR`
  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_SCALING_CAPABILITIES_KHR`

## [](#_version_history)Version History

- Revision 1, 2025-03-31 (Shahbaz Youssefi)
  
  - Based on VK\_EXT\_surface\_maintenance1

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_surface_maintenance1).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
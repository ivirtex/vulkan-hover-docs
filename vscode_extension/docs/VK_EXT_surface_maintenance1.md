# VK\_EXT\_surface\_maintenance1(3) Manual Page

## Name

VK\_EXT\_surface\_maintenance1 - instance extension



## [](#_registered_extension_number)Registered Extension Number

275

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)  
and  
[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface_maintenance1.html) extension

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_surface_maintenance1%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_surface_maintenance1%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_surface\_maintenance1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_surface_maintenance1.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-12-16

**Interactions and External Dependencies**

- Promoted to `VK_KHR_surface_maintenance1`

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

`VK_EXT_surface_maintenance1` adds a collection of window system integration features that were intentionally left out or overlooked in the original `VK_KHR_surface` extension.

The new features are as follows:

- Allow querying number of min/max images from a surface for a particular presentation mode.
- Allow querying a surface’s scaled presentation capabilities.
- Allow querying a surface for the set of presentation modes which can be easily switched between without requiring swapchain recreation.

## [](#_promotion_to_vk_khr_surface_maintenance1)Promotion to `VK_KHR_surface_maintenance1`

All functionality in this extension is included in `VK_KHR_surface_maintenance1`, with the suffix changed to KHR. The original type, enum and command names are still available as aliases of the KHR functionality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html):
  
  - [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html)
- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityEXT.html)
  - [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)

## [](#_new_enums)New Enums

- [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsEXT.html)
- [VkPresentScalingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPresentGravityFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagsEXT.html)
- [VkPresentScalingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SURFACE_MAINTENANCE_1_EXTENSION_NAME`
- `VK_EXT_SURFACE_MAINTENANCE_1_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_COMPATIBILITY_EXT`
  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_EXT`
  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_SCALING_CAPABILITIES_EXT`

## [](#_version_history)Version History

- Revision 0, 2019-02-27 (Lionel Landwerlin)
  
  - Internal revisions
- Revision 0, 2020-06-15 (James Jones)
  
  - Internal revisions
- Revision 1, 2022-11-09 (Shahbaz Youssefi)
  
  - Add functionality and complete spec

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_surface_maintenance1).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
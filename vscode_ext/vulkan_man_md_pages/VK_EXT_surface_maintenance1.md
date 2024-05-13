# VK_EXT_surface_maintenance1(3) Manual Page

## Name

VK_EXT_surface_maintenance1 - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

275

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  
and  
[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_surface_maintenance1%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_surface_maintenance1%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_surface_maintenance1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_surface_maintenance1.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-11-09

**Contributors**  
- Jeff Juliano, NVIDIA

- Lionel Landwerlin, Intel

- Shahbaz Youssefi, Google

- Chris Forbes, Google

- Ian Elliott, Google

- Hans-Kristian Arntzen, Valve

- Daniel Stone, Collabora

## <a href="#_description" class="anchor"></a>Description

[`VK_EXT_surface_maintenance1`](VK_EXT_surface_maintenance1.html) adds a
collection of window system integration features that were intentionally
left out or overlooked in the original
[`VK_KHR_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) extension.

The new features are as follows:

- Allow querying number of min/max images from a surface for a
  particular presentation mode.

- Allow querying a surface’s scaled presentation capabilities.

- Allow querying a surface for the set of presentation modes which can
  be easily switched between without requiring swapchain recreation.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html):

  - [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html):

  - [VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeCompatibilityEXT.html)

  - [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagBitsEXT.html)

- [VkPresentScalingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPresentGravityFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagsEXT.html)

- [VkPresentScalingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SURFACE_MAINTENANCE_1_EXTENSION_NAME`

- `VK_EXT_SURFACE_MAINTENANCE_1_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_COMPATIBILITY_EXT`

  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_EXT`

  - `VK_STRUCTURE_TYPE_SURFACE_PRESENT_SCALING_CAPABILITIES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2019-02-27 (Lionel Landwerlin)

  - Internal revisions

- Revision 1, 2022-11-09 (Shahbaz Youssefi)

  - Add functionality and complete spec

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_surface_maintenance1"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

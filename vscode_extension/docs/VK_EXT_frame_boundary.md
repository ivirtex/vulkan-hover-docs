# VK\_EXT\_frame\_boundary(3) Manual Page

## Name

VK\_EXT\_frame\_boundary - device extension



## [](#_registered_extension_number)Registered Extension Number

376

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- James Fitzpatrick [\[GitHub\]jamesfitzpatrick](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_frame_boundary%5D%20%40jamesfitzpatrick%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_frame_boundary%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_frame\_boundary](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_frame_boundary.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-06-14

**Contributors**

- James Fitzpatrick, Imagination Technologies
- Hugues Evrard, Google
- Melih Yasin Yalcin, Google
- Andrew Garrard, Imagination Technologies
- Jan-Harald Fredriksen, Arm
- Vassili Nikolaev, NVIDIA
- Ting Wei, Huawei

## [](#_description)Description

[VK\_EXT\_frame\_boundary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_frame_boundary.html) is a device extension that helps **tools** (such as debuggers) to group queue submissions per frames in non-trivial scenarios, typically when [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) is not a relevant frame boundary delimiter.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFrameBoundaryFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFrameBoundaryFeaturesEXT.html)
- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html), [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html), [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html):
  
  - [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html)

## [](#_new_enums)New Enums

- [VkFrameBoundaryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkFrameBoundaryFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_FRAME_BOUNDARY_EXTENSION_NAME`
- `VK_EXT_FRAME_BOUNDARY_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_FRAME_BOUNDARY_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAME_BOUNDARY_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 0, 2022-01-14 (Hugues Evard)
  
  - Initial proposal
- Revision 1, 2023-06-14 (James Fitzpatrick)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_frame_boundary)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
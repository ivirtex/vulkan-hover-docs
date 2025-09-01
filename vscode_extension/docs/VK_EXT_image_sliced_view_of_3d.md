# VK\_EXT\_image\_sliced\_view\_of\_3d(3) Manual Page

## Name

VK\_EXT\_image\_sliced\_view\_of\_3d - device extension



## [](#_registered_extension_number)Registered Extension Number

419

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html)  
     and  
     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_sliced_view_of_3d%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_sliced_view_of_3d%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_image\_sliced\_view\_of\_3d](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_image_sliced_view_of_3d.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-01-24

**IP Status**

No known IP claims.

**Contributors**

- Mike Blumenkrantz, Valve
- Hans-Kristian Arntzen, Valve
- Ricardo Garcia, Igalia
- Shahbaz Youssefi, Google
- Piers Daniell, NVIDIA

## [](#_description)Description

This extension allows creating 3D views of 3D images such that the views contain a subset of the slices in the image, using a Z offset and range, for the purpose of using the views as storage image descriptors. This matches functionality in D3D12 and is primarily intended to support D3D12 emulation.

## [](#_new_structures)New Structures

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html):
  
  - [VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSlicedCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_IMAGE_SLICED_VIEW_OF_3D_EXTENSION_NAME`
- `VK_EXT_IMAGE_SLICED_VIEW_OF_3D_SPEC_VERSION`
- `VK_REMAINING_3D_SLICES_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_SLICED_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_SLICED_VIEW_OF_3D_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2022-10-21 (Mike Blumenkrantz)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_image_sliced_view_of_3d).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
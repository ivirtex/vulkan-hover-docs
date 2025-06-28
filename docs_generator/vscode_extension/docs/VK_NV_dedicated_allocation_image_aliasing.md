# VK\_NV\_dedicated\_allocation\_image\_aliasing(3) Manual Page

## Name

VK\_NV\_dedicated\_allocation\_image\_aliasing - device extension



## [](#_registered_extension_number)Registered Extension Number

241

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_dedicated\_allocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dedicated_allocation.html)  
     and  
     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Nuno Subtil [\[GitHub\]nsubtil](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_dedicated_allocation_image_aliasing%5D%20%40nsubtil%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_dedicated_allocation_image_aliasing%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-01-04

**Contributors**

- Nuno Subtil, NVIDIA
- Jeff Bolz, NVIDIA
- Eric Werness, NVIDIA
- Axel Gneiting, id Software

## [](#_description)Description

This extension allows applications to alias images on dedicated allocations, subject to specific restrictions: the extent and the number of layers in the image being aliased must be smaller than or equal to those of the original image for which the allocation was created, and every other image parameter must match.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_DEDICATED_ALLOCATION_IMAGE_ALIASING_EXTENSION_NAME`
- `VK_NV_DEDICATED_ALLOCATION_IMAGE_ALIASING_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEDICATED_ALLOCATION_IMAGE_ALIASING_FEATURES_NV`

## [](#_version_history)Version History

- Revision 1, 2019-01-04 (Nuno Subtil)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_dedicated_allocation_image_aliasing)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VK\_EXT\_ycbcr\_image\_arrays(3) Manual Page

## Name

VK\_EXT\_ycbcr\_image\_arrays - device extension



## [](#_registered_extension_number)Registered Extension Number

253

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_ycbcr_image_arrays%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_ycbcr_image_arrays%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-01-15

**Contributors**

- Piers Daniell, NVIDIA

## [](#_description)Description

This extension allows images of a format that requires [Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion) to be created with multiple array layers, which is otherwise restricted.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceYcbcrImageArraysFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceYcbcrImageArraysFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_YCBCR_IMAGE_ARRAYS_EXTENSION_NAME`
- `VK_EXT_YCBCR_IMAGE_ARRAYS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_IMAGE_ARRAYS_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2019-01-15 (Piers Daniell)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_ycbcr_image_arrays)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
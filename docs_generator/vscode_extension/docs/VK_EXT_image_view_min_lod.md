# VK\_EXT\_image\_view\_min\_lod(3) Manual Page

## Name

VK\_EXT\_image\_view\_min\_lod - device extension



## [](#_registered_extension_number)Registered Extension Number

392

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Joshua Ashton [\[GitHub\]Joshua-Ashton](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_view_min_lod%5D%20%40Joshua-Ashton%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_view_min_lod%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-11-09

**IP Status**

No known IP claims.

**Contributors**

- Joshua Ashton, Valve
- Hans-Kristian Arntzen, Valve
- Samuel Iglesias Gonsalvez, Igalia
- Tobias Hector, AMD
- Faith Ekstrand, Intel
- Tom Olson, ARM

## [](#_description)Description

This extension allows applications to clamp the minimum LOD value during [Image Level(s) Selection](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-image-level-selection), [Texel Gathering](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-gather) and [Integer Texel Coordinate Operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-integer-coordinate-operations) with a given [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) by [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewMinLodCreateInfoEXT.html)::`minLod`.

This extension may be useful to restrict a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) to only mips which have been uploaded, and the use of fractional `minLod` can be useful for smoothly introducing new mip levels when using linear mipmap filtering.

## [](#_new_structures)New Structures

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html):
  
  - [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewMinLodCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImageViewMinLodFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageViewMinLodFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_IMAGE_VIEW_MIN_LOD_EXTENSION_NAME`
- `VK_EXT_IMAGE_VIEW_MIN_LOD_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_MIN_LOD_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_MIN_LOD_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2021-07-06 (Joshua Ashton)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_image_view_min_lod)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
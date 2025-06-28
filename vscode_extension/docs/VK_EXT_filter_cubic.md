# VK\_EXT\_filter\_cubic(3) Manual Page

## Name

VK\_EXT\_filter\_cubic - device extension



## [](#_registered_extension_number)Registered Extension Number

171

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_filter_cubic%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_filter_cubic%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-12-13

**Contributors**

- Bill Licea-Kane, Qualcomm Technologies, Inc.
- Andrew Garrard, Samsung
- Daniel Koch, NVIDIA
- Donald Scorgie, Imagination Technologies
- Graeme Leese, Broadcom
- Jan-Harald Fredriksen, ARM
- Jeff Leger, Qualcomm Technologies, Inc.
- Tobias Hector, AMD
- Tom Olson, ARM
- Stuart Smith, Imagination Technologies

## [](#_description)Description

`VK_EXT_filter_cubic` extends `VK_IMG_filter_cubic`.

It documents cubic filtering of other image view types. It adds new structures that **can** be added to the `pNext` chain of [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) and [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) that **can** be used to determine which image types and which image view types support cubic filtering.

## [](#_new_structures)New Structures

- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html):
  
  - [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)
- Extending [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html):
  
  - [VkPhysicalDeviceImageViewImageFormatInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageViewImageFormatInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_FILTER_CUBIC_EXTENSION_NAME`
- `VK_EXT_FILTER_CUBIC_SPEC_VERSION`
- Extending [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html):
  
  - `VK_FILTER_CUBIC_EXT`
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_FILTER_CUBIC_IMAGE_VIEW_IMAGE_FORMAT_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_IMAGE_FORMAT_INFO_EXT`

## [](#_version_history)Version History

- Revision 3, 2019-12-13 (wwlk)
  
  - Delete requirement to cubic filter the formats USCALED\_PACKED32, SSCALED\_PACKED32, UINT\_PACK32, and SINT\_PACK32 (cut/paste error)
- Revision 2, 2019-06-05 (wwlk)
  
  - Clarify 1D optional
- Revision 1, 2019-01-24 (wwlk)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_filter_cubic)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
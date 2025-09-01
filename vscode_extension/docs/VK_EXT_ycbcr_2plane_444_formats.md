# VK\_EXT\_ycbcr\_2plane\_444\_formats(3) Manual Page

## Name

VK\_EXT\_ycbcr\_2plane\_444\_formats - device extension



## [](#_registered_extension_number)Registered Extension Number

331

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Tony Zlatinski [\[GitHub\]tzlatinski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_ycbcr_2plane_444_formats%5D%20%40tzlatinski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_ycbcr_2plane_444_formats%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-07-28

**IP Status**

No known IP claims.

**Contributors**

- Piers Daniell, NVIDIA
- Ping Liu, Intel

## [](#_description)Description

This extension adds some Y′CBCR formats that are in common use for video encode and decode, but were not part of the `VK_KHR_sampler_ycbcr_conversion` extension.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_YCBCR_2PLANE_444_FORMATS_EXTENSION_NAME`
- `VK_EXT_YCBCR_2PLANE_444_FORMATS_SPEC_VERSION`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_G10X6_B10X6R10X6_2PLANE_444_UNORM_3PACK16_EXT`
  - `VK_FORMAT_G12X4_B12X4R12X4_2PLANE_444_UNORM_3PACK16_EXT`
  - `VK_FORMAT_G16_B16R16_2PLANE_444_UNORM_EXT`
  - `VK_FORMAT_G8_B8R8_2PLANE_444_UNORM_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_2_PLANE_444_FORMATS_FEATURES_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

The format enumerants introduced by the extension are included in core Vulkan 1.3, with the EXT suffix omitted. However, runtime support for these formats is optional in core Vulkan 1.3, while if this extension is supported, runtime support is mandatory. The feature structure is not promoted. The original enum names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2020-03-08 (Piers Daniell)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_ycbcr_2plane_444_formats).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VK\_EXT\_rgba10x6\_formats(3) Manual Page

## Name

VK\_EXT\_rgba10x6\_formats - device extension



## [](#_registered_extension_number)Registered Extension Number

345

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_rgba10x6_formats%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_rgba10x6_formats%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-09-29

**IP Status**

No known IP claims.

**Contributors**

- Jan-Harald Fredriksen, Arm
- Graeme Leese, Broadcom
- Spencer Fricke, Samsung

## [](#_description)Description

This extension enables the `VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16` format to be used without a [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) enabled.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_RGBA10X6_FORMATS_EXTENSION_NAME`
- `VK_EXT_RGBA10X6_FORMATS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RGBA10X6_FORMATS_FEATURES_EXT`

## [](#_issues)Issues

1\) Should we reuse the existing format enumeration or introduce a new one?

**RESOLVED**: We reuse an existing format enumeration, `VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16`, that was previously exclusively used for YCbCr and therefore had a set of limitations related to that usage. The alternative was to introduce a new format token with exactly the same bit representation as the existing token, but without the limitations.

2\) Should we only introduce `VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16` or also 1-3 component variations?

**RESOLVED**: Only the 4-component format is introduced because the 1- and 2- component variations are already not exclusive to YCbCr, and the 3-component variation is not a good match for hardware capabilities.

## [](#_version_history)Version History

- Revision 1, 2021-09-29 (Jan-Harald Fredriksen)
  
  - Initial EXT version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_rgba10x6_formats).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0